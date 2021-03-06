package paperspace

import (
	"net/http"
	"time"
)

// ScriptsService manages Scripts for the Paperspace API.
// Paperspace API docs: https://paperspace.github.io/paperspace-node/scripts.html
type ScriptsService struct {
	client *Client
}

// Scripts represents a Paperspace Script.
type Scripts struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	OwnerType   string    `json:"ownerType,omitempty"`
	OwnerID     string    `json:"ownerId,omitempty"`
	Description string    `json:"description,omitempty"`
	DtCreated   time.Time `json:"dtCreated,omitempty"`
	IsEnabled   bool      `json:"isEnabled,omitempty"`
	RunOnce     bool      `json:"runOnce,omitempty"`
}

// ScriptsRequest represents a request to create a Paperspace Script.
type ScriptsRequest struct {
	MachineID         string `json:"machineId,omitempty"`
	ScriptName        string `json:"scriptName,omitempty"`
	ScriptFile        string `json:"scriptFile,omitempty"`
	ScriptText        string `json:"scriptText,omitempty"`
	ScriptDescription string `json:"scriptDescription,omitempty"`
	IsEnabled         bool   `json:"isEnabled,omitempty"`
	RunOnce           bool   `json:"runOnce,omitempty"`
}

// Create creates a new Paperspace Script based on given parameters.
// https://paperspace.github.io/paperspace-node/scripts.html#.create
func (s *ScriptsService) Create(options *ScriptsRequest) (*Scripts, *Response, error) {
	apiEndpoint := "scripts/createScript"
	req, err := s.client.NewRequest(http.MethodPost, apiEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	scripts := new(Scripts)
	resp, err := s.client.Do(req, scripts)
	if err != nil {
		return nil, resp, err
	}

	return scripts, resp, nil
}

// List lists all scripts from Paperspace.
// https://paperspace.github.io/paperspace-node/scripts.html#.list
func (s *ScriptsService) List() (*Scripts, *Response, error) {
	apiEndpoint := "scripts/getScripts"
	req, err := s.client.NewRequest(http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	scriptsList := new(Scripts)
	resp, err := s.client.Do(req, &scriptsList)
	if err != nil {
		return nil, resp, err
	}
	return scriptsList, resp, nil
}
