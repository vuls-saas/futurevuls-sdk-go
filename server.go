package fvuls

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetServerDetail get TaskDetail
// https://doc.vuls.biz/#/server
func (c *Client) GetServerDetail(prm GetServerDetailParam) (*Server, error) {
	req, err := http.NewRequest("GET", c.urlFor(fmt.Sprintf("/v1/server/%s", prm.ServerID)).String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var server Server
	err = json.NewDecoder(resp.Body).Decode(&server)
	if err != nil {
		return nil, err
	}
	return &server, err
}

// GetAllServerList get a list of servers
// https://doc.vuls.biz/#/servers
func (c *Client) GetAllServerList(prm GetServerListParam) ([]*Server, error) {
	req, err := http.NewRequest("GET", c.urlFor("/v1/servers").String(), nil)
	if err != nil {
		return nil, err
	}
	q := c.BaseURL.Query()
	if prm.FilterCveID != nil {
		q.Set("filterCveID", *prm.FilterCveID)
	}
	if prm.FilterRoleID != nil {
		q.Set("filterRoleID", fmt.Sprint(*prm.FilterRoleID))
	}
	req.URL.RawQuery = q.Encode()

	servers := []*Server{}
	for i := 1; ; i++ {
		if prm.Limit == 0 {
			prm.Limit = 1000
		}
		q.Set("limit", fmt.Sprint(prm.Limit))
		q.Set("page", fmt.Sprint(i))

		req.URL.RawQuery = q.Encode()
		resp, err := c.Request(req)
		defer closeResponse(resp)
		if err != nil {
			return nil, err
		}
		var res PagingServers
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}
		servers = append(servers, res.Servers...)
		if uint(i) == res.Paging.TotalPage {
			break
		}
	}
	return servers, nil

}

// UpdateServer updates Server
// https://doc.vuls.biz/#/server
func (c *Client) UpdateServer(prm UpdateServerParam) (*Server, error) {
	resp, err := c.PutJSON(fmt.Sprintf("/v1/server/%d", prm.ServerID), prm)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var server Server
	err = json.NewDecoder(resp.Body).Decode(&server)
	if err != nil {
		return nil, err
	}
	return &server, err
}

// DeleteServer delete server
// https://doc.vuls.biz/#/server
func (c *Client) DeleteServer(prm DeleteServerParam) error {
	req, err := http.NewRequest("DELETE", c.urlFor(fmt.Sprintf("/v1/server/%d", prm.ServerID)).String(), nil)
	if err != nil {
		return err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	return err
}

// GetServerListParam is the payload type of the server service getServerList
// method.
type GetServerListParam struct {
	// api key auth
	Key *string
	// Page Number
	Page uint
	// Limit
	Limit uint
	// Offset
	Offset uint
	// CveID filter
	FilterCveID *string
	// ServerRoleID filter
	FilterRoleID *int
}

// PagingServers is the result type of the server service getServerList method.
type PagingServers struct {
	// Paging
	Paging *Paging
	// Servers list
	Servers []*Server
}

// GetServerDetailParam is the payload type of the server service
// getServerDetail method.
type GetServerDetailParam struct {
	// api key auth
	Key *string
	// Server ID or Server UUID
	ServerID string
}

// UpdateServerParam is the payload type of the server service updateServer
// method.
type UpdateServerParam struct {
	// api key auth
	Key *string
	// Server ID
	ServerID int
	// ServerName of server
	ServerName *string
	// ServerRoleID of server
	RoleID *int
	// DefaultUserID of server
	DefaultUserID *int
}

// DeleteServerParam is the payload type of the server service deleteServer
// method.
type DeleteServerParam struct {
	// api key auth
	Key *string
	// Server ID
	ServerID int
}
