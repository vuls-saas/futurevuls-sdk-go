package fvuls

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCveListPayload is the payload type of the cve service getCveList method.
type GetCveListPayload struct {
	// Page Number
	Page uint
	// Limit
	Limit uint
	// flag of onlyActiveCve
	OnlyActiveCve bool
	// ServerID filter
	FilterServerID *int
	// ServerRoleID filter
	FilterRoleID *int
	// PackageID filter
	FilterPkgID *int
	// CpeID filter
	FilterCpeID *int
}

// PagingCves is the result type of the cve service getCveList method.
type PagingCves struct {
	// Paging
	Paging *Paging
	// Cves list
	Cves []*Cve
}

// GetCveDetailPayload is the payload type of the cve service getCveDetail
// method.
type GetCveDetailPayload struct {
	// Cve ID
	CveID string
}

// Cve is the result type of the cve service getCveDetail method.
type Cve struct {
	// Cve ID string of cve
	CveID string
	// cvss v2 scores of cve
	ScoreV2s map[string]float64
	// cvss v3 scores of cve
	ScoreV3s map[string]float64
	// cvss v2 vectors of cve
	VectorV2s map[string]string
	// cvss v3 vectors of cve
	VectorV3s map[string]string
	// maxV2 of cve
	MaxV2 float64
	// maxV3 of cve
	MaxV3 float64
	// cwes of cve
	Cwes []*Cwe
	// Title of cve
	Title string
	// Flag of active cve
	IsNotActive bool
	// NewTaskCount of cve
	NewTaskCount int
	// AllTaskCount of cve
	AllTaskCount int
	// isOwaspTopTen2017 of cve
	IsOwaspTopTen2017 bool
	// tmpMetricV2 of cve
	TmpMetricV2 *TmpMetric
	// tmpMetricV3 of cve
	TmpMetricV3 *TmpMetric
	// secMetric of cve
	SecMetrics []*SecMetric
	// envMetricV2 of cve
	EnvMetricV2s []*EnvMetricV2
	// envMetricV3 of cve
	EnvMetricV3s []*EnvMetricV3
	// serverOsFamilies of cve
	ServerOsFamilies []string
	// cvss of cve
	Cvss interface{}
	// references of cve
	References map[string]string
	// topicCount of cve
	TopicCount int
	// topicLastUpdatedAt of cve
	TopicLastUpdatedAt string
	// created time
	CreatedAt string
	// updated time
	UpdatedAt string
}

// Paging describes a paging object
type Paging struct {
	// Total Page Size
	TotalPage uint
	// Offset
	Offset uint
	// Page
	Page uint
	// Limit
	Limit uint
	// TotalCount
	TotalCount uint
}

// Cwe describes a cwe
type Cwe struct {
	// sourceType of cwe
	SourceType string
	// cweID of cwe
	CweID string
	// english summary of cwe
	English string
	// japanese summary of cwe
	Japanese string
	// owaspTopTen2017 of cwe
	OwaspTopTen2017 *string
}

// TmpMetric describes a tmpMetric
type TmpMetric struct {
	// E of tmpMetric
	E string
	// RL of tmpMetric
	Rl string
	// RC of tmpMetric
	Rc string
	// created time
	CreatedAt string
	// updated time
	UpdatedAt string
}

// SecMetric describes a secMetric
type SecMetric struct {
	// ServerRoleID of secMetric
	RoleID int
	// ServerRoleName of secMetric
	RoleName string
	// CR of secMetric
	Cr string
	// IR of secMetric
	Ir string
	// AR of secMetric
	Ar string
	// created time
	CreatedAt string
	// updated time
	UpdatedAt string
}

// EnvMetricV2 describes a envMetricV2
type EnvMetricV2 struct {
	// CveID of envMetricV2
	CveID string
	// ServerRoleID of envMetricV2
	RoleID int
	// ServerRoleName of envMetricV2
	RoleName string
	// TD of envMetricV2
	Td string
	// CDP of envMetricV2
	Cdp string
	// created time
	CreatedAt string
	// updated time
	UpdatedAt string
}

// EnvMetricV3 describes a envMetricV3
type EnvMetricV3 struct {
	// CveID of envMetricV3
	CveID string
	// ServerRoleID of envMetricV3
	RoleID int
	// ServerRoleName of envMetricV3
	RoleName string
	// MAV of envMetricV3
	Mav string
	// MAC of envMetricV3
	Mac string
	// MPR of envMetricV3
	Mpr string
	// MUI of envMetricV3
	Mui string
	// MS of envMetricV3
	Ms string
	// MC of envMetricV3
	Mc string
	// MA of envMetricV3
	Ma string
	// created time
	CreatedAt string
	// updated time
	UpdatedAt string
}

// GetCveDetail get CveDetail
func (c *Client) GetCveDetail(prm GetCveDetailPayload) (*Cve, error) {
	req, err := http.NewRequest("GET", c.urlFor(fmt.Sprintf("/v1/cve/"+prm.CveID)).String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var cve Cve
	err = json.NewDecoder(resp.Body).Decode(&cve)
	if err != nil {
		return nil, err
	}
	return &cve, err
}
