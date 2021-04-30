package client

import (
	"errors"
	"net/http"
)

type Scanner struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) ListScanners(arg string) ([]Scanner, error) {
	scanners := make([]Scanner, 0)

	req, err := c.newRequest("GET", "/api/v1/scanners", nil)
	if err != nil {
		return scanners, err
	}

	a := make(map[string]string, 0)

	resp, err := c.do(req, &a)
	if err != nil {
		return scanners, err
	}

	for id, name := range a {
		scanners = append(scanners, Scanner{
			ID:   id,
			Name: name,
		})
	}

	if resp.StatusCode != http.StatusOK {
		return scanners, errors.New("response not ok")
	}

	return scanners, nil
}
