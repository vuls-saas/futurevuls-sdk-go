package fvuls

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCpeDetail get pkgCpe
// https://doc.vuls.biz/#/task
func (c *Client) GetCpeDetail(prm GetCpeDetailParam) (*PkgCpe, error) {
	req, err := http.NewRequest("GET", c.urlFor(fmt.Sprintf("/v1/pkgCpe/cpe/%d", prm.CpeID)).String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var pkgCpe PkgCpe
	err = json.NewDecoder(resp.Body).Decode(&pkgCpe)
	if err != nil {
		return nil, err
	}
	return &pkgCpe, err
}

// GetPkgDetail get pkgCpe
// https://doc.vuls.biz/#/pkgCpe/pkgCpe#getPkgDetail
func (c *Client) GetPkgDetail(prm GetPkgDetailParam) (*PkgCpe, error) {
	req, err := http.NewRequest("GET", c.urlFor(fmt.Sprintf("/v1/pkgCpe/pkg/%d", prm.PkgID)).String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var pkgCpe PkgCpe
	err = json.NewDecoder(resp.Body).Decode(&pkgCpe)
	if err != nil {
		return nil, err
	}
	return &pkgCpe, err
}

// GetPkgCpeList get a list of PkgCpeList
// https://doc.vuls.biz/#/pkgCpe/pkgCpe#getPkgCpeList
func (c *Client) GetPkgCpeList(prm GetPkgCpeListParam) (*PagingPkgCpes, error) {
	req, err := http.NewRequest("GET", c.urlFor("/v1/pkgCpes").String(), nil)
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
		q.Set("filterCveID", *prm.FilterCveID)
	}
	if prm.FilterTaskID != nil {
		q.Set("filterTaskID", fmt.Sprint(*prm.FilterTaskID))
	}
	if prm.FilterServerID != nil {
		q.Set("filterServerID", fmt.Sprint(*prm.FilterServerID))
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
	var pkgCpes PagingPkgCpes
	err = json.NewDecoder(resp.Body).Decode(&pkgCpes)
	if err != nil {
		return nil, err
	}
	return &pkgCpes, err
}

// GetAllPkgCpeList get a list of pkgCpes
// https://doc.vuls.biz/#/pkgCpe/pkgCpe#getPkgCpeList
func (c *Client) GetAllPkgCpeList(prm GetPkgCpeListParam) ([]*PkgCpe, error) {
	req, err := http.NewRequest("GET", c.urlFor("/v1/pkgCpes").String(), nil)
	if err != nil {
		return nil, err
	}
	q := c.BaseURL.Query()
	if prm.FilterCveID != nil {
		q.Set("filterCveID", *prm.FilterCveID)
	}
	if prm.FilterTaskID != nil {
		q.Set("filterTaskID", fmt.Sprint(*prm.FilterTaskID))
	}
	if prm.FilterServerID != nil {
		q.Set("filterServerID", fmt.Sprint(*prm.FilterServerID))
	}
	if prm.FilterRoleID != nil {
		q.Set("filterRoleID", fmt.Sprint(*prm.FilterRoleID))
	}

	pkgCpes := []*PkgCpe{}
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
		var res PagingPkgCpes
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}
		pkgCpes = append(pkgCpes, res.PkgCpes...)
		if uint(i) == res.Paging.TotalPage {
			break
		}
	}
	return pkgCpes, nil
}

// GetPkgCpeListParam is the payload type of the pkgCpe service getPkgCpeList
// method.
type GetPkgCpeListParam struct {
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
	// TaskID filter
	FilterTaskID *int
	// ServerID filter
	FilterServerID *int
	// ServerRoleID filter
	FilterRoleID *int
}

// PagingPkgCpes is the result type of the pkgCpe service getPkgCpeList method.
type PagingPkgCpes struct {
	// Paging
	Paging *Paging
	// PkgCpes list
	PkgCpes []*PkgCpe
}

// GetPkgDetailParam is the payload type of the pkgCpe service getPkgDetail
// method.
type GetPkgDetailParam struct {
	// api key auth
	Key *string
	// PackageID
	PkgID int
}

// GetCpeDetailParam is the payload type of the pkgCpe service getCpeDetail
// method.
type GetCpeDetailParam struct {
	// api key auth
	Key *string
	// cpe ID
	CpeID int
}

// Server describes a server
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

// ServerTag describes a server tag
type ServerTag struct {
	// ID of server tag
	ID int
	// Name of server tag
	Name string
}
