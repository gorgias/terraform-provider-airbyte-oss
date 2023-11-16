// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectionStreamRequestBody struct {
	ConnectionID string             `json:"connectionId"`
	Streams      []ConnectionStream `json:"streams"`
}

func (o *ConnectionStreamRequestBody) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ConnectionStreamRequestBody) GetStreams() []ConnectionStream {
	if o == nil {
		return []ConnectionStream{}
	}
	return o.Streams
}
