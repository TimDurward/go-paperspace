package paperspace

import (
	"fmt"
	"net/http"
	"time"
)

// NetworksService manages Networks for the Paperspace API.
// https://paperspace.github.io/paperspace-node/networks.html
type NetworksService struct {
	client *Client
}

// Network represents a Paperspace Network.
type Network struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Region    string    `json:"region"`
	DtCreated time.Time `json:"dtCreated"`
	Network   string    `json:"network"`
	Netmask   string    `json:"netmask"`
	TeamID    string    `json:"teamId"`
}

// List lists information about all machines available to the authenticated user or team.
// https://paperspace.github.io/paperspace-node/machines.html#.list
func (s *NetworksService) List() ([]*Network, *Response, error) {
	apiEndpoint := "networks/getNetworks"

	req, err := s.client.NewRequest(http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var networks []*Network
	resp, err := s.client.Do(req, &networks)
	if err != nil {
		fmt.Print(err)
		return nil, resp, err
	}

	return networks, resp, nil
}
