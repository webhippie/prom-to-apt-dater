{
  description = "Nix flake for development";

  inputs = {
    nixpkgs = {
      url = "github:nixos/nixpkgs/nixpkgs-unstable";
    };

    utils = {
      url = "github:numtide/flake-utils";
    };
  };

  outputs = { self, nixpkgs, utils, ... }@inputs:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            gnumake
            go_1_19
          ];

          shellHook = ''
            export PROM_TO_APT_DATER_LOG_LEVEL=debug
            export PROM_TO_APT_DATER_LOG_PRETTY=true
            export PROM_TO_APT_DATER_LOG_COLOR=true
          '';
        };
      }
    );
}
