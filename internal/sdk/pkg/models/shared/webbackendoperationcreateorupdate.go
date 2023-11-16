// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WebBackendOperationCreateOrUpdate struct {
	OperationID           *string               `json:"operationId,omitempty"`
	WorkspaceID           string                `json:"workspaceId"`
	Name                  string                `json:"name"`
	OperatorConfiguration OperatorConfiguration `json:"operatorConfiguration"`
}

func (o *WebBackendOperationCreateOrUpdate) GetOperationID() *string {
	if o == nil {
		return nil
	}
	return o.OperationID
}

func (o *WebBackendOperationCreateOrUpdate) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *WebBackendOperationCreateOrUpdate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WebBackendOperationCreateOrUpdate) GetOperatorConfiguration() OperatorConfiguration {
	if o == nil {
		return OperatorConfiguration{}
	}
	return o.OperatorConfiguration
}
