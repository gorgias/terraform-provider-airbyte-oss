// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceUpdate struct {
	SourceID string `json:"sourceId"`
	// The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	Name                    string      `json:"name"`
	SecretID                *string     `json:"secretId,omitempty"`
}

func (o *SourceUpdate) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *SourceUpdate) GetConnectionConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.ConnectionConfiguration
}

func (o *SourceUpdate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceUpdate) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}
