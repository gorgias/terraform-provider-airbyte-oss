// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceDiscoverSchemaRequestBody struct {
	SourceID           string  `json:"sourceId"`
	ConnectionID       *string `json:"connectionId,omitempty"`
	DisableCache       *bool   `json:"disable_cache,omitempty"`
	NotifySchemaChange *bool   `json:"notifySchemaChange,omitempty"`
}

func (o *SourceDiscoverSchemaRequestBody) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *SourceDiscoverSchemaRequestBody) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *SourceDiscoverSchemaRequestBody) GetDisableCache() *bool {
	if o == nil {
		return nil
	}
	return o.DisableCache
}

func (o *SourceDiscoverSchemaRequestBody) GetNotifySchemaChange() *bool {
	if o == nil {
		return nil
	}
	return o.NotifySchemaChange
}
