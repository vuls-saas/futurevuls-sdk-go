package fvuls

import (
	"fmt"
	"net/http"
)

// Health check health status of API
// https://doc.vuls.biz/#/health/health#health
func (c *Client) Health() error {
	req, err := http.NewRequest("GET", c.urlFor(fmt.Sprintf("/health")).String(), nil)
	if err != nil {
		return err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	return err
}
