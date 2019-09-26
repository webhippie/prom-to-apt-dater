package output

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/Knetic/govaluate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Client defines an output client.
type Client struct {
	filter           string
	filterExpression *govaluate.EvaluableExpression
	group            string
	groupExpression  *govaluate.EvaluableExpression
	name             string
	nameExpression   *govaluate.EvaluableExpression
	user             string
	userExpression   *govaluate.EvaluableExpression
	host             string
	hostExpression   *govaluate.EvaluableExpression
	port             string
	portExpression   *govaluate.EvaluableExpression
	file             string
	template         string
	targets          []map[string]interface{}
}

// Generate builds the hosts config and writes it to the file.
func (c *Client) Generate() error {
	groups := Groups{}

	for _, target := range c.targets {
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

		user, err := c.UserExpression(target)

		if err != nil {
			log.Error().
				Err(err).
				Fields(target).
				Msg("Failed to detect user")

			return err
		}

		port, err := c.PortExpression(target)

		if err != nil {
			log.Error().
				Err(err).
				Fields(target).
				Msg("Failed to detect port")

			return err
		}

		g := groups.Append(Group{
			Name: group,
		})

		h := g.Hosts.Append(Host{
			Name:    name,
			User:    user,
			Address: host,
			Port:    port,
		})

		log.Debug().
			Str("group", g.Name).
			Str("name", h.Name).
			Str("address", h.Address).
			Msg("Adding host to config")
	}

	result, err := c.Template(groups)

	if err != nil {
		return err
	}

	return c.Writer(result)
}

// Writer writes the generated content into a file.
func (c *Client) Writer(content string) error {
	if c.file == "-" {
		fmt.Fprintf(
			os.Stdout,
			content,
		)
	} else {
		if err := ioutil.WriteFile(
			c.file,
			[]byte(content),
			0644,
		); err != nil {
			return errors.Wrap(err, "failed to write to file")
		}
	}

	return nil
}

// Template renders the template content.
func (c *Client) Template(groups Groups) (string, error) {
	result := bytes.NewBufferString("")
	tpl := DefaultTemplate

	if c.template != "" {
		if _, err := os.Stat(c.template); os.IsNotExist(err) {
			return "", fmt.Errorf("template file does not exist")
		}

		content, err := ioutil.ReadFile(c.template)

		if err != nil {
			return "", fmt.Errorf("failed to read template content")
		}

		tpl = string(content)
	}

	t := template.Must(
		template.New(
			"",
		).Parse(
			tpl,
		),
	)

	if err := t.Execute(result, struct {
		Groups Groups
	}{
		Groups: groups,
	}); err != nil {
		return "", fmt.Errorf("failed to process template. %s", err)
	}

	return result.String(), nil
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

// UserExpression evaluates the user expression.
func (c *Client) UserExpression(attrs map[string]interface{}) (string, error) {
	res, err := c.userExpression.Evaluate(attrs)

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

// PortExpression evaluates the port expression.
func (c *Client) PortExpression(attrs map[string]interface{}) (string, error) {
	res, err := c.portExpression.Evaluate(attrs)

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
		expr, err := govaluate.NewEvaluableExpressionWithFunctions(
			c.filter,
			Helpers(),
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to parse filter expression")
		}

		c.filterExpression = expr
	}

	{
		expr, err := govaluate.NewEvaluableExpressionWithFunctions(
			c.group,
			Helpers(),
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to parse group expression")
		}

		c.groupExpression = expr
	}

	{
		expr, err := govaluate.NewEvaluableExpressionWithFunctions(
			c.name,
			Helpers(),
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to parse name expression")
		}

		c.nameExpression = expr
	}

	{
		expr, err := govaluate.NewEvaluableExpressionWithFunctions(
			c.user,
			Helpers(),
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to parse user expression")
		}

		c.userExpression = expr
	}

	{
		expr, err := govaluate.NewEvaluableExpressionWithFunctions(
			c.host,
			Helpers(),
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to parse host expression")
		}

		c.hostExpression = expr
	}

	{
		expr, err := govaluate.NewEvaluableExpressionWithFunctions(
			c.port,
			Helpers(),
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to parse port expression")
		}

		c.portExpression = expr
	}

	return c, nil
}
