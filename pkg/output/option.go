package output

// Option configures a client option.
type Option func(*Client)

// WithFilter returns an option to set filter.
func WithFilter(val string) Option {
	return func(c *Client) {
		c.Filter = val
	}
}

// WithGroup returns an option to set group.
func WithGroup(val string) Option {
	return func(c *Client) {
		c.Group = val
	}
}

// WithName returns an option to set name.
func WithName(val string) Option {
	return func(c *Client) {
		c.Name = val
	}
}

// WithHost returns an option to set host.
func WithHost(val string) Option {
	return func(c *Client) {
		c.Host = val
	}
}

// WithFile returns an option to set file.
func WithFile(val string) Option {
	return func(c *Client) {
		c.File = val
	}
}

// WithTemplate returns an option to set template.
func WithTemplate(val string) Option {
	return func(c *Client) {
		c.Template = val
	}
}

// WithTargets returns an option to set targets.
func WithTargets(val []map[string]interface{}) Option {
	return func(c *Client) {
		c.Targets = val
	}
}
