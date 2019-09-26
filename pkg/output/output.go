package output

import (
	"github.com/Knetic/govaluate"
	"github.com/rs/zerolog/log"
)

// Client defines an output client.
type Client struct {
	Filter           string
	filterExpression *govaluate.EvaluableExpression
	Group            string
	groupExpression  *govaluate.EvaluableExpression
	Name             string
	nameExpression   *govaluate.EvaluableExpression
	Host             string
	hostExpression   *govaluate.EvaluableExpression
	File             string
	Template         string
	Targets          []map[string]interface{}
}

// Generate builds the hosts config and writes it to the file.
func (c *Client) Generate() error {
	for _, target := range c.Targets {
		match, err := c.FilterExpression(target)

		if err != nil {
			return err
		}

		if !match {
			log.Debug().
				Fields(target).
				Msg("Does not match filter criteria")

			continue
		}

		group, err := c.GroupExpression(target)

		if err != nil {
			log.Error().
				Err(err).
				Fields(target).
				Msg("Failed to detect group")

			return err
		}

		name, err := c.NameExpression(target)

		if err != nil {
			log.Error().
				Err(err).
				Fields(target).
				Msg("Failed to detect name")

			return err
		}

		host, err := c.HostExpression(target)

		if err != nil {
			log.Error().
				Err(err).
				Fields(target).
				Msg("Failed to detect host")

			return err
		}

		log.Info().
			Str("group", group).
			Str("name", name).
			Str("host", host).
			Msg("")

	}

	return nil
}

// FilterExpression evaluates the filter expression.
func (c *Client) FilterExpression(attrs map[string]interface{}) (bool, error) {
	res, err := c.filterExpression.Evaluate(attrs)

	if err != nil {
		return false, err
	}

	return res.(bool), nil
}

// GroupExpression evaluates the group expression.
func (c *Client) GroupExpression(attrs map[string]interface{}) (string, error) {
	res, err := c.groupExpression.Evaluate(attrs)

	if err != nil {
		return "", err
	}

	return res.(string), nil
}

// NameExpression evaluates the name expression.
func (c *Client) NameExpression(attrs map[string]interface{}) (string, error) {
	res, err := c.nameExpression.Evaluate(attrs)

	if err != nil {
		return "", err
	}

	return res.(string), nil
}

// HostExpression evaluates the host expression.
func (c *Client) HostExpression(attrs map[string]interface{}) (string, error) {
	res, err := c.hostExpression.Evaluate(attrs)

	if err != nil {
		return "", err
	}

	return res.(string), nil
}

// New initializes a new output client.
func New(opts ...Option) (*Client, error) {
	c := new(Client)

	for _, opt := range opts {
		opt(c)
	}

	{
		expr, err := govaluate.NewEvaluableExpression(
			c.Filter,
		)

		if err != nil {
			return nil, err
		}

		c.filterExpression = expr
	}

	{
		expr, err := govaluate.NewEvaluableExpression(
			c.Group,
		)

		if err != nil {
			return nil, err
		}

		c.groupExpression = expr
	}

	{
		expr, err := govaluate.NewEvaluableExpression(
			c.Name,
		)

		if err != nil {
			return nil, err
		}

		c.nameExpression = expr
	}

	{
		expr, err := govaluate.NewEvaluableExpression(
			c.Host,
		)

		if err != nil {
			return nil, err
		}

		c.hostExpression = expr
	}

	return c, nil
}
