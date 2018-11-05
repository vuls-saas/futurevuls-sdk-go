package fvuls

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTaskDetail get TaskDetail
// https://doc.vuls.biz/#/task
func (c *Client) GetTaskDetail(prm GetTaskDetailParam) (*Task, error) {
	req, err := http.NewRequest("GET", c.urlFor(fmt.Sprintf("/v1/task/%d", prm.TaskID)).String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var task Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, err
}

// UpdateTask updates Task
// https://doc.vuls.biz/#/task
func (c *Client) UpdateTask(prm UpdateTaskParam) (*Task, error) {
	resp, err := c.PutJSON(fmt.Sprintf("/v1/task/%d", prm.TaskID), prm)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var task Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, err
}

// GetTaskList get a list of tasks
// https://doc.vuls.biz/#/task
func (c *Client) GetTaskList(prm GetTaskListParam) (*PagingTasks, error) {
	req, err := http.NewRequest("GET", c.urlFor("/v1/tasks").String(), nil)
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
	for _, p := range prm.FilterStatus {
		q.Set("filterStatus", p)
	}
	for _, p := range prm.FilterPriority {
		q.Set("filterPriority", p)
	}
	if prm.FilterIgnore != nil {
		q.Set("filterIgnore", c.toJSON(prm.FilterIgnore))
	}
	for _, p := range prm.FilterMainUserIDs {
		q.Set("filterMainUserIDs", fmt.Sprint(p))
	}
	for _, p := range prm.FilterSubUserIDs {
		// q.Set("filterSubUserIDs", c.toJSON(prm.FilterSubUserIDs))
		q.Set("filterSubUserIDs", fmt.Sprint(p))
	}
	if prm.FilterCveID != nil {
		q.Set("filterCveID", c.toJSON(*prm.FilterCveID))
	}
	if prm.FilterServerID != nil {
		q.Set("filterServerID", fmt.Sprint(*prm.FilterServerID))
	}
	if prm.FilterRoleID != nil {
		q.Set("filterRoleID", fmt.Sprint(*prm.FilterRoleID))
	}
	if prm.FilterPkgID != nil {
		q.Set("filterPkgID", fmt.Sprint(*prm.FilterPkgID))
	}
	if prm.FilterCpeID != nil {
		q.Set("filterCpeID", fmt.Sprint(*prm.FilterCpeID))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}
	var tasks PagingTasks
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return &tasks, err
}

// GetTaskListParam is the payload type of the task service getTaskList
// method.
type GetTaskListParam struct {
	// api key auth
	Key *string
	// Page Number
	Page uint
	// Limit
	Limit uint
	// Offset
	Offset uint
	// Status filter
	FilterStatus []string
	// Priority filter
	FilterPriority []string
	// Ignore filter
	FilterIgnore *bool
	// MainUserIDs filter
	FilterMainUserIDs []int
	// MainSubIDs filter
	FilterSubUserIDs []int
	// CveID filter
	FilterCveID *string
	// ServerID filter
	FilterServerID *int
	// ServerRoleID filter
	FilterRoleID *int
	// PackageID filter
	FilterPkgID *int
	// CpeID filter
	FilterCpeID *int
}

// PagingTasks is the result type of the task service getTaskList method.
type PagingTasks struct {
	// Paging
	Paging *Paging
	// Task list
	Tasks []*Task
}

// GetTaskDetailParam is the payload type of the task service getTaskDetail
// method.
type GetTaskDetailParam struct {
	// api key auth
	Key *string
	// Task ID
	TaskID int
}

// Task is the result type of the task service getTaskDetail method.
type Task struct {
	// ID of task
	ID int
	// CVE ID of task
	CveID string
	// MainUserID of task
	MainUserID *int `json:",omitempty"`
	// MainUserName of task
	MainUserName *string `json:",omitempty"`
	// SubUserID of task
	SubUserID *int `json:",omitempty"`
	// SubUserName of task
	SubUserName *string `json:",omitempty"`
	// ServerID of task
	ServerID int
	// ServerName of task
	ServerName string
	Server     *Server
	// ServerRoleID of task
	RoleID *int `json:",omitempty"`
	// ServerRoleName of task
	RoleName *string `json:",omitempty"`
	// Pcakge And Cpe list of task
	PkgCpes []*PkgCpe `json:",omitempty"`
	// Package And CPE Names of task
	PkgCpeNames []string `json:",omitempty"`
	// packageStatus of task
	PackageStatuses map[string]string `json:",omitempty"`
	// Flag of Not Fixed Yet of task
	PkgNotFixedYet *bool `json:",omitempty"`
	// Key Value of CveID and Cvss of task
	Cvss map[string]interface{}
	// ApplyingPatchOn of task
	ApplyingPatchOn *string
	// Status of task
	Status string
	// Priority of task
	Priority string
	// Ignore of task
	Ignore bool
	// Ignore until of task
	IgnoreUntil *string `json:",omitempty"`
	// Comment of task
	Comments []*TaskComment
	// DetectionMethod of task
	DetectionMethods []*DetectionMethod `json:",omitempty"`
	// DetectionTools of task
	DetectionTools []*DetectionTool
	// created time of task
	CreatedAt string
	// updated time of task
	UpdatedAt string
}

// UpdateTaskParam is the payload type of the task service updateTask method.
type UpdateTaskParam struct {
	// api key auth
	Key *string
	// Task ID
	TaskID int
	// mainUserID of task
	MainUserID *int
	// subUserID of task
	SubUserID *int
	// Status of task
	Status *string
	// Priority of task
	Priority *string
	// applyingPatchOn (YYYY-MM-DD) UTC
	ApplyingPatchOn *string
}

// ChildTask describes a child task
type ChildTask struct {
	// ID of task
	ID int
	// CVE ID of task
	CveID string
	// MainUserID of task
	MainUserID *int `json:",omitempty"`
	// MainUserName of task
	MainUserName *string `json:",omitempty"`
	// SubUserID of task
	SubUserID *int `json:",omitempty"`
	// SubUserName of task
	SubUserName *string `json:",omitempty"`
	// ServerID of task
	ServerID int
	// ApplyingPatchOn of task
	ApplyingPatchOn *string
	// Status of task
	Status string
	// Priority of task
	Priority string
	// Ignore of task
	Ignore bool
	// Ignore until of task
	IgnoreUntil *string `json:",omitempty"`
	// created time of task
	CreatedAt string
	// updated time of task
	UpdatedAt string
}

// PkgCpe describes a pkgCpe
type PkgCpe struct {
	// ID of package or cpe
	ID int
	// Package ID of package
	PkgID *int `json:",omitempty"`
	// CpeID of cpe
	CpeID *int `json:",omitempty"`
	// Name of package or cpe
	Name string
	// Version of package or cpe
	Version string
	// Cpe URI of cpe
	CpeURI *string `json:",omitempty"`
	// New Version of package
	NewVersion *string `json:",omitempty"`
	// New Release of package
	NewRelease *string `json:",omitempty"`
	// Release of package
	Release *string
	// Flag of Not fixed yet of package
	NotFixedYet *bool `json:",omitempty"`
	// package status of package
	PackageStatuses map[string]string `json:",omitempty"`
	// ServerID of package or cpe
	ServerID int
	// ServerUUID of package or cpe
	ServerUUID string
	// ServerName of package or cpe
	ServerName string
	Server     *Server `json:",omitempty"`
	// updated time of server
	Tasks []*ChildTask `json:",omitempty"`
	// AffectedProcess list of package
	AffectedProcs []*AffectedProc `json:",omitempty"`
	// NeedRestartProcess list of package
	NeedRestartProcs []*NeedRestartProc `json:",omitempty"`
	// crated time of package or cpe
	CreatedAt string
	// updated time of package or cpe
	UpdatedAt string
}

// AffectedProc describes a AffectedProc
type AffectedProc struct {
	// PID
	Pid string
	// AffectedProc Name
	Name string
}

// NeedRestartProc NEedescribes a NeedRestartProc
type NeedRestartProc struct {
	// PID
	Pid string
	// Path of NeedRestartProc
	Path string
	// ServiceName of NeedRestartProc
	ServiceName string
	// InitSystem of NeedRestartProc
	InitSystem string
}

// TaskComment describes a TaskComment
type TaskComment struct {
	// ID of TaskComment
	ID int
	// Comment content of TaskComment
	Comment string
	// Type of TaskComment
	Type string
	// UserID of TaskComment
	UserID int
	// UserName of TaskComment
	UserName string
	// created time of TaskComment
	CreatedAt string
	// updated time of TaskComment
	UpdatedAt string
}

// DetectionMethod describes a DetectionMethod
type DetectionMethod struct {
	// Detection Method Name
	Name string
	// ReliabilityScore
	ReliabilityScore int
}

// DetectionTool describes a DetectionTool
type DetectionTool struct {
	// Detection Tool Name
	Name string
	// PatchAppliedAt
	PatchAppliedAt *string
}
