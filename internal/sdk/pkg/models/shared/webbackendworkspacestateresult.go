// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WebBackendWorkspaceStateResult struct {
	HasConnections  bool `json:"hasConnections"`
	HasSources      bool `json:"hasSources"`
	HasDestinations bool `json:"hasDestinations"`
}

func (o *WebBackendWorkspaceStateResult) GetHasConnections() bool {
	if o == nil {
		return false
	}
	return o.HasConnections
}

func (o *WebBackendWorkspaceStateResult) GetHasSources() bool {
	if o == nil {
		return false
	}
	return o.HasSources
}

func (o *WebBackendWorkspaceStateResult) GetHasDestinations() bool {
	if o == nil {
		return false
	}
	return o.HasDestinations
}
