package prometheus

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/jackspirou/syscerts"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

// Client defines a Prometheus client.
type Client struct {
	URL      string
	Username string
	Password string

	api v1.API
}

// Targets fetches and parses the targets from Prometheus API endpoint.
func (c *Client) Targets() ([]map[string]interface{}, error) {
	targets, err := c.api.Targets(context.Background())

	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)

	for _, target := range targets.Active {
		attrs := make(map[string]interface{})

		for key, val := range target.Labels {
			attrs[string(key)] = string(val)
		}

		for key, val := range target.DiscoveredLabels {
			attrs[string(key)] = string(val)
		}

		result = append(result, attrs)
	}

	return result, nil
}

// RoundTrip implements the RoundTripper interface.
func (c *Client) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			RootCAs: syscerts.SystemRootsPool(),
		},
	}

	if c.Username != "" && c.Password != "" {
		req.SetBasicAuth(
			c.Username,
			c.Password,
		)
	}

	return transport.RoundTrip(req)
}

// New initializes a new Prometheus client.
func New(opts ...Option) (*Client, error) {
	c := new(Client)

	for _, opt := range opts {
		opt(c)
	}

	apiClient, err := api.NewClient(
		api.Config{
			Address:      c.URL,
			RoundTripper: c,
		},
	)

	if err != nil {
		return nil, err
	}

	c.api = v1.NewAPI(
		apiClient,
	)

	return c, nil
}
