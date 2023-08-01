package dota2

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	Client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return &Client{
		Client: &http.Client{
			Timeout: timeout,
			Transport: loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) GetHeroes() ([]Hero, error) {
	resp, err := c.Client.Get("https://api.opendota.com/api/heroes")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r heroesResponce
	if err := json.Unmarshal(body, &r.heroes); err != nil {
		return nil, err
	}
	return r.heroes, err
}
