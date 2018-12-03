package fvuls

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetRoleDetail get RoleDetail
// https://doc.vuls.biz/#/role/role#getRoleDetail
func (c *Client) GetRoleDetail(prm GetRoleDetailParam) (*Role, error) {
	req, err := http.NewRequest("GET", c.urlFor(fmt.Sprintf("/v1/role/%d", prm.RoleID)).String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var role Role
	err = json.NewDecoder(resp.Body).Decode(&role)
	if err != nil {
		return nil, err
	}
	return &role, err
}

// GetAllRoleList get a list of Roles
// https://doc.vuls.biz/#/role/role#getRoleList
func (c *Client) GetAllRoleList(prm GetRoleListParam) ([]*Role, error) {
	req, err := http.NewRequest("GET", c.urlFor("/v1/roles").String(), nil)
	if err != nil {
		return nil, err
	}
	q := c.BaseURL.Query()
	if prm.FilterCveID != nil {
		q.Set("filterCveID", *prm.FilterCveID)
	}
	req.URL.RawQuery = q.Encode()

	roles := []*Role{}
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
		var res PagingRoles
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}
		roles = append(roles, res.Roles...)
		if uint(i) == res.Paging.TotalPage {
			break
		}
	}
	return roles, nil

}

// UpdateRole updates Role
// https://doc.vuls.biz/#/role/role#updateRole
func (c *Client) UpdateRole(prm UpdateRoleParam) (*Role, error) {
	resp, err := c.PutJSON(fmt.Sprintf("/v1/role/%d", prm.RoleID), prm)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var role Role
	err = json.NewDecoder(resp.Body).Decode(&role)
	if err != nil {
		return nil, err
	}
	return &role, err
}

// DeleteRole delete role
// https://doc.vuls.biz/#/role/role#deleteRole
func (c *Client) DeleteRole(prm DeleteRoleParam) error {
	req, err := http.NewRequest("DELETE", c.urlFor(fmt.Sprintf("/v1/role/%d", prm.RoleID)).String(), nil)
	if err != nil {
		return err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	return err
}

// Role is the result type of the role service getRoleDetail method.
type Role struct {
	// ID of server role
	ID int
	// Name of server role
	Name string
	// SecMetric of server role
	SecMetric *SecMetric `json:"secMetric,omitempty"`
	// envMetricV2s of server role
	EnvMetricV2s []*EnvMetricV2 `json:"envMetricV2s,omitempty"`
	// envMetricV3s of server role
	EnvMetricV3s []*EnvMetricV3 `json:"envMetricV3s,omitempty"`
	// Server Count of server role
	ServerCount *int `json:"serverCount,omitempty"`
	// Servers of server role
	Servers []*Server `json:"servers,omitempty"`
	// NewTaskCount of server role
	NewTaskCount *int `json:"newTaskCount,omitempty"`
	// AllTaskCount of server role
	AllTaskCount *int `json:"allTaskCount,omitempty"`
	// created time of server role
	CreatedAt string
	// updated time of server role
	UpdatedAt string
}

// GetRoleListParam is the payload type of the role service getRoleList
// method.
type GetRoleListParam struct {
	// api key auth
	Key *string
	// Page Number (default: 1)
	Page uint
	// Limit (default: 20, max: 100)
	Limit uint
	// Offset
	Offset uint
	// CveID filter
	FilterCveID *string
}

// PagingRoles is the result type of the role service getRoleList method.
type PagingRoles struct {
	// Paging
	Paging *Paging
	// ServerRole list
	Roles []*Role
}

// GetRoleDetailParam is the payload type of the role service getRoleDetail
// method.
type GetRoleDetailParam struct {
	// api key auth
	Key *string
	// Role ID
	RoleID int
}

// UpdateRoleParam is the payload type of the role service updateRole method.
type UpdateRoleParam struct {
	// api key auth
	Key *string
	// Role ID
	RoleID int
	// RoleName of role
	RoleName *string
}

// DeleteRoleParam is the payload type of the role service deleteRole method.
type DeleteRoleParam struct {
	// api key auth
	Key *string
	// Role ID
	RoleID int
}
