def main(ctx):
  before = testing(ctx)

  stages = [
    docker(ctx, 'amd64'),
    docker(ctx, 'arm64v8'),
    docker(ctx, 'arm32v6'),
    binary(ctx, 'linux'),
    binary(ctx, 'darwin'),
    binary(ctx, 'windows'),
  ]

  after = [
    manifest(ctx),
    documentation(ctx)
  ]

  for b in before:
    for s in stages:
      s['depends_on'].append(b['name'])

  for s in stages:
    for a in after:
      a['depends_on'].append(s['name'])

  return before + stages + after

def testing(ctx):
  return [{
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'testing',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'generate',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make generate'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'vet',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make vet'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'staticcheck',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make staticcheck'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'lint',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make lint'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'build',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make build'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'test',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make test'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'codacy',
        'image': 'plugins/codacy:1',
        'pull': 'always',
        'settings': {
          'token': {
            'from_secret': 'codacy_token'
          }
        }
      }
    ],
    'volumes': [
      {
        'name': 'gopath',
        'temp': {}
      }
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**'
      ]
    }
  }]

def docker(ctx, arch):
  agent = "amd64"

  if arch == "arm32v6":
    agent = "arm"

  if arch == "arm64v8":
    agent = "arm64"

  if ctx.build.event == "pull_request":
    settings = {
      'dry_run': True,
      'tags': 'linux-%s' % arch,
      'auto_tag_suffix': 'linux-%s' % arch,
      'dockerfile': 'docker/Dockerfile.linux.%s' % arch,
      'repo': 'tboerger/prom-to-apt-dater',
    }
  else:
    settings = {
      'username': {
        'from_secret': 'docker_username'
      },
      'password': {
        'from_secret': 'docker_password'
      },
      'auto_tag': True,
      'auto_tag_suffix': 'linux-%s' % arch,
      'dockerfile': 'docker/Dockerfile.linux.%s' % arch,
      'repo': 'tboerger/prom-to-apt-dater',
    }

  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': arch,
    'platform': {
      'os': 'linux',
      'arch': agent,
    },
    'steps': [
      {
        'name': 'generate',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make generate'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'build',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make build'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'docker',
        'image': 'plugins/docker:18.09',
        'pull': 'always',
        'settings': settings
      },
    ],
    'volumes': [
      {
        'name': 'gopath',
        'temp': {}
      }
    ],
    'depends_on': [],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**'
      ]
    }
  }

def binary(ctx, name):
  if ctx.build.event == "tag":
    settings = {
      'endpoint': {
        'from_secret': 's3_endpoint'
      },
      'access_key': {
        'from_secret': 'aws_access_key_id'
      },
      'secret_key': {
        'from_secret': 'aws_secret_access_key'
      },
      'bucket': {
        'from_secret': 's3_bucket'
      },
      'path_style': True,
      'strip_prefix': 'dist/release/',
      'source': 'dist/release/*',
      'target': '/prom-to-apt-dater/%s' % ctx.build.ref.replace("refs/tags/v", "")
    }
  else:
    settings = {
      'endpoint': {
        'from_secret': 's3_endpoint'
      },
      'access_key': {
        'from_secret': 'aws_access_key_id'
      },
      'secret_key': {
        'from_secret': 'aws_secret_access_key'
      },
      'bucket': {
        'from_secret': 's3_bucket'
      },
      'path_style': True,
      'strip_prefix': 'dist/release/',
      'source': 'dist/release/*',
      'target': '/prom-to-apt-dater/testing'
    }

  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': name,
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'generate',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make generate'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'build',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make release-%s' % name
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'finish',
        'image': 'webhippie/golang:1.12',
        'pull': 'always',
        'commands': [
          'make release-finish'
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app'
          }
        ]
      },
      {
        'name': 'gpgsign',
        'image': 'plugins/gpgsign:1',
        'pull': 'always',
        'settings': {
          'key': {
            'from_secret': 'gpgsign_key'
          },
          'passphrase': {
            'from_secret': 'gpgsign_passphrase'
          },
          'files': [
            'dist/release/*'
          ],
          'excludes': [
            'dist/release/*.sha256'
          ],
          'detach_sign': True
        },
        'when': {
          'event': [
            'push',
            'tag'
          ]
        }
      },
      {
        'name': 'upload',
        'image': 'plugins/s3:1',
        'pull': 'always',
        'settings': settings,
        'when': {
          'event': [
            'push',
            'tag'
          ]
        }
      },
      {
        'name': 'changelog',
        'image': 'toolhippie/calens:latest',
        'pull': 'always',
        'commands': [
          'calens --version %s | tee dist/CHANGELOG.md' % ctx.build.ref.replace("refs/tags/v", "")
        ],
        'when': {
          'event': [
            'tag'
          ]
        }
      },
      {
        'name': 'release',
        'image': 'plugins/github-release:1',
        'pull': 'always',
        'settings': {
          'api_key': {
            'from_secret': 'github_token'
          },
          'files': [
            'dist/release/*'
          ],
          'title': ctx.build.ref.replace("refs/tags/", ""),
          'note': 'dist/CHANGELOG.md',
          'overwrite': True
        },
        'when': {
          'event': [
            'tag'
          ]
        }
      }
    ],
    'volumes': [
      {
        'name': 'gopath',
        'temp': {}
      }
    ],
    'depends_on': [],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**'
      ]
    }
  }

def manifest(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'manifest',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'manifest',
        'image': 'plugins/manifest',
        'pull': 'always',
        'settings': {
          'username': {
            'from_secret': 'docker_username'
          },
          'password': {
            'from_secret': 'docker_password'
          },
          'spec': 'docker/manifest.tmpl',
          'auto_tag': 'true',
          'ignore_missing': 'true',
        },
      },
      {
        'name': 'microbadger',
        'image': 'plugins/webhook',
        'pull': 'always',
        'settings': {
          'urls': {
            'from_secret': 'microbadger_url'
          }
        },
      }
    ],
    'depends_on': [],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**'
      ]
    }
  }

def documentation(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'documentation',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'generate',
        'image': 'webhippie/hugo:latest',
        'pull': 'always',
        'commands': [
          'make docs'
        ]
      },
      {
        'name': 'publish',
        'image': 'plugins/gh-pages:1',
        'pull': 'always',
        'settings': {
          'username': {
            'from_secret': 'github_username'
          },
          'password': {
            'from_secret': 'github_token'
          },
          'pages_directory': 'docs/public/',
          'temporary_base': 'tmp/',
        },
        'when': {
          'event': {
            'exclude': [
              'pull_request'
            ]
          }
        }
      }
    ],
    'depends_on': [],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/pull/**'
      ]
    }
  }
