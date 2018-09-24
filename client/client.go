package client

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Client struct {
	c *http.Client
	t *transport
}

func New() *Client {
	t := newTransport()
	return &Client{
		c: &http.Client{
			Transport: t,
			Timeout:   15 * time.Second,
		},
		t: t,
	}
}

func (c *Client) HTTP() *http.Client {
	return c.c
}

func (c *Client) Describe(ch chan<- *prometheus.Desc) {
	c.t.mRequests.Describe(ch)
}

func (c *Client) Collect(ch chan<- prometheus.Metric) {
	c.t.mRequests.Collect(ch)
}

// check interfaces
var (
	_ prometheus.Collector = (*Client)(nil)
)
