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

// GetServerList get a list of servers
// https://doc.vuls.biz/#/servers
func (c *Client) GetServerList(prm GetServerListParam) (*PagingServers, error) {
	req, err := http.NewRequest("GET", c.urlFor("/v1/servers").String(), nil)
	if err != nil {
		return nil, err
	}
	q := c.BaseURL.Query()
	if 0 < prm.Page {
		q.Set("page", fmt.Sprint(prm.Page))
	}
	if 0 < prm.Limit {
		q.Set("limit", fmt.Sprint(prm.Limit))
	}
	if 0 < prm.Offset {
		q.Set("offset", fmt.Sprint(prm.Offset))
	}
	if prm.FilterCveID != nil {
		q.Set("filterCveID", c.toJSON(*prm.FilterCveID))
	}
	if prm.FilterRoleID != nil {
		q.Set("filterRoleID", fmt.Sprint(*prm.FilterRoleID))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var servers PagingServers
	err = json.NewDecoder(resp.Body).Decode(&servers)
	if err != nil {
		return nil, err
	}
	return &servers, err
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

// Server is the result type of the server service getServerDetail method.
type Server struct {
	// ID of server
	ID int `json:"id"`
	// UUID of server
	ServerUUID string
	// UUID of server
	HostUUID string
	// Name of server
	ServerName string
	// ID of server role
	ServerroleID int
	// Name of server role
	ServerroleName string
	// OS Name of server
	OsFamily string
	// OS Version of server
	OsVersion string
	// Whether server needs kernel restart
	NeedKernelRestart bool
	// default user ID of server
	DefaultUserID *int
	// default user name of server
	DefaultUserName *string
	// last scanned time of server
	LastScannedAt *string `json:"lastScannedAt,omitempty"`
	// last uploaded time of server
	LastUploadedAt *string `json:"lastUploadedAt,omitempty"`
	// tags is list of server tag
	Tags []*ServerTag `json:"tags,omitempty"`
	// tasks of server
	Tasks []*ChildTask `json:"tasks,omitempty"`
	// crated time of server
	CreatedAt string
	// updated time of server
	UpdatedAt string
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

// ServerTag describes a server tag
type ServerTag struct {
	// ID of server tag
	ID int
	// Name of server tag
	Name string
}
