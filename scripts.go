package paperspace

import (
	"net/http"
)

// ScriptsService manages Scripts for the Paperspace API.
// Paperspace API docs: https://paperspace.github.io/paperspace-node/scripts.html
type ScriptsService struct {
	client *Client
}

// Scripts represents a Paperspace Script.
type Scripts struct {
	ID                string `json:"id,omitempty"`
	ScriptName        string `json:"scriptName,omitempty"`
	ScriptDescription string `json:"scriptDescription,omitempty"`
	ScriptFile        string `json:"scriptFile,omitempty"`
	ScriptText        string `json:"scriptText,omitempty"`
	IsEnabled         bool   `json:"isEnabled,omitempty"`
	RunOnce           bool   `json:"runOnce,omitempty"`
}

// Create creates a new Paperspace Script based on given parameters.
// https://paperspace.github.io/paperspace-node/scripts.html#.create
func (s *ScriptsService) Create(options *Scripts) (*Scripts, *Response, error) {
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

// List gets all scripts from Paperspace.
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
