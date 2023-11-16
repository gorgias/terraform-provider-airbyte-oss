// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OperationRead struct {
	WorkspaceID           string                `json:"workspaceId"`
	OperationID           string                `json:"operationId"`
	Name                  string                `json:"name"`
	OperatorConfiguration OperatorConfiguration `json:"operatorConfiguration"`
}

func (o *OperationRead) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *OperationRead) GetOperationID() string {
	if o == nil {
		return ""
	}
	return o.OperationID
}

func (o *OperationRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OperationRead) GetOperatorConfiguration() OperatorConfiguration {
	if o == nil {
		return OperatorConfiguration{}
	}
	return o.OperatorConfiguration
}
