package output

// Option configures a client option.
type Option func(*Client)

// WithFilter returns an option to set filter.
func WithFilter(val string) Option {
	return func(c *Client) {
		c.filter = val
	}
}

// WithGroup returns an option to set group.
func WithGroup(val string) Option {
	return func(c *Client) {
		c.group = val
	}
}

// WithName returns an option to set name.
func WithName(val string) Option {
	return func(c *Client) {
		c.name = val
	}
}

// WithUser returns an option to set user.
func WithUser(val string) Option {
	return func(c *Client) {
		c.user = val
	}
}

// WithHost returns an option to set host.
func WithHost(val string) Option {
	return func(c *Client) {
		c.host = val
	}
}

// WithPort returns an option to set port.
func WithPort(val string) Option {
	return func(c *Client) {
		c.port = val
	}
}

// WithFile returns an option to set file.
func WithFile(val string) Option {
	return func(c *Client) {
		c.file = val
	}
}

// WithTemplate returns an option to set template.
func WithTemplate(val string) Option {
	return func(c *Client) {
		c.template = val
	}
}

// WithTargets returns an option to set targets.
func WithTargets(val []map[string]interface{}) Option {
	return func(c *Client) {
		c.targets = val
	}
}
