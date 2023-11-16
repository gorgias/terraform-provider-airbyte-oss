// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceRead struct {
	SourceDefinitionID string `json:"sourceDefinitionId"`
	SourceID           string `json:"sourceId"`
	WorkspaceID        string `json:"workspaceId"`
	// The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	Name                    string      `json:"name"`
	SourceName              string      `json:"sourceName"`
	Icon                    *string     `json:"icon,omitempty"`
}

func (o *SourceRead) GetSourceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.SourceDefinitionID
}

func (o *SourceRead) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *SourceRead) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceRead) GetConnectionConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.ConnectionConfiguration
}

func (o *SourceRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceRead) GetSourceName() string {
	if o == nil {
		return ""
	}
	return o.SourceName
}

func (o *SourceRead) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}
