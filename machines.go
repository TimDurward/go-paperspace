package paperspace

import (
	"net/http"
)

// MachinesService manages Machines for the Paperspace API.
// Paperspace API docs: https://paperspace.github.io/paperspace-node/machines.html
type MachinesService struct {
	client *Client
}

// Machines represents a Paperspace Machine.
type Machines struct {
	Available bool `json:"available,omitempty"`
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

	machines := new(Machines)
	resp, err := s.client.Do(req, &machines)
	if err != nil {
		return nil, resp, err
	}

	return machines, resp, err
}
