package prometheus

// Option configures a client option.
type Option func(*Client)

// WithURL returns an option to set url.
func WithURL(val string) Option {
	return func(c *Client) {
		c.URL = val
	}
}

// WithUsername returns an option to set username.
func WithUsername(val string) Option {
	return func(c *Client) {
		c.Username = val
	}
}

// WithPassword returns an option to set password.
func WithPassword(val string) Option {
	return func(c *Client) {
		c.Password = val
	}
}
