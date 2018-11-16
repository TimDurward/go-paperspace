package paperspace

import (
	"fmt"
	"net/http"
)

// MachinesService manages Machines for the Paperspace API.
// Paperspace API docs: https://paperspace.github.io/paperspace-node/machines.html
type MachinesService struct {
	client *Client
}

// Machines represents a Paperspace Machine.
type Machines struct {
	ID                     string `json:"id,omitempty"`
	MachineID              string `json:"machineId,omityempty"`
	Name                   string `json:"name,omitempty"`
	OS                     string `json:"os,omitempty"`
	RAM                    string `json:"ram,omitempty"`
	GPU                    string `json:"gpu,omitempty"`
	Cpus                   int    `json:"cpus,omitempty"`
	StorageTotal           string `json:"storageTotal,omitempty"`
	StorageUsed            string `json:"storageUsed,omitempty"`
	UsageRate              string `json:"usageRate,omitempty"`
	ShutdownTimeoutInHours int    `json:"shutdownTimeoutInHours,omitempty"`
	ShutdownTimeoutForces  bool   `json:"shutdownTimeoutForces,omitempty"`
	PerformAutoSnapshot    bool   `json:"performAutoSnapshot,omitempty"`
	AutoSnapshotFrequency  int    `json:"autoSnapshotFrequency,omitempty"`
	AutoSnapshotSaveCount  int    `json:"autoSnapshotSaveCount,omitempty"`
	AgentType              string `json:"agentType,omitempty"`
	State                  string `json:"state,omitempty"`
	UpdatesPending         bool   `json:"updatesPending,omitempty"`
	NetworkID              string `json:"networkId,omitempty"`
	PrivateIPAddress       string `json:"privateIpAddress,omitempty"`
	PublicIPAddress        string `json:"publicIpAddress,omitempty"`
	DynamicPublicIP        bool   `json:"dynamicPublicIp,omitempty"`
	Region                 string `json:"region,omitempty"`
	UserID                 string `json:"userId,omitempty"`
	TeamID                 string `json:"teamId,omitempty"`
	ScriptID               string `json:"scriptId,omitempty"`
	DtCreated              string `json:"dtCreated,omitempty"`
	DtLastRun              string `json:"dtLastRun,omitempty"`
	Available              bool   `json:"available,omitempty"`
}

// MachinesRequest represents a request to create a Paperspace Machine.
type MachinesRequest struct {
	Region            string `json:"region,omitempty"`
	MachineType       string `json:"machineType,omitempty"`
	Size              int    `json:"size,omitempty"`
	BillingType       string `json:"billingType,omitempty"`
	MachineName       string `json:"machineName,omitempty"`
	TemplateID        string `json:"templateId,omitempty"`
	AssignPublicID    bool   `json:"assignPublicIp,omitempty"`
	DynamicPublicID   bool   `json:"dynamicPublicIp,omitempty"`
	NetworkID         string `json:"networkId,omitempty"`
	TeamID            string `json:"teamId,omitempty"`
	UserID            string `json:"userId,omitempty"`
	Email             string `json:"email,omitempty"`
	Password          string `json:"password,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	NotificationEmail string `json:"notificationEmail,omitempty"`
	ScriptID          string `json:"scriptId,omitempty"`
}

// MachineOptions specifies the optional parameters to the
// MachinesService.Availability method.
type MachineOptions struct {
	Region      string `url:"region,omitempty"`
	MachineType string `url:"machineType,omitempty"`
}

// Availability gets machine availability for the given region and machine type.
// https://paperspace.github.io/paperspace-node/machines.html#.availability
func (s *MachinesService) Availability(opt *MachineOptions) (*Machines, *Response, error) {
	apiEndpoint := "machines/getAvailability"

	url, err := addOptions(apiEndpoint, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	machine := new(Machines)
	resp, err := s.client.Do(req, &machine)
	if err != nil {
		return nil, resp, err
	}

	return machine, resp, nil
}

// Create creates new a virtual machine.
// https://paperspace.github.io/paperspace-node/machines.html#.create
func (s *MachinesService) Create(mr *MachinesRequest) (*Machines, *Response, error) {
	apiEndpoint := "machines/createSingleMachinePublic"

	req, err := s.client.NewRequest(http.MethodPost, apiEndpoint, mr)
	if err != nil {
		return nil, nil, err
	}

	machine := new(Machines)
	resp, err := s.client.Do(req, &machine)
	if err != nil {
		fmt.Print(err)
		return nil, resp, err
	}

	return machine, resp, nil
}

// Destroy destroys an existing virtual machine.
// https://paperspace.github.io/paperspace-node/machines.html#.destroy
func (s *MachinesService) Destroy(machineID string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("machines/%s/destroyMachine", machineID)

	req, err := s.client.NewRequest(http.MethodPost, apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		fmt.Print(err)
		return resp, err
	}

	return resp, nil
}

// List lists information about all machines available to the authenticated user or team.
// https://paperspace.github.io/paperspace-node/machines.html#.list
func (s *MachinesService) List(filter ...*Machines) ([]*Machines, *Response, error) {
	apiEndpoint := "machines/getMachines"

	req, err := s.client.NewRequest(http.MethodGet, apiEndpoint, filter)
	if err != nil {
		return nil, nil, err
	}

	var machines []*Machines
	resp, err := s.client.Do(req, &machines)
	if err != nil {
		fmt.Print(err)
		return nil, resp, err
	}

	return machines, resp, nil
}
