// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AirbyteStreamAndConfiguration - each stream is split in two parts; the immutable schema from source and mutable configuration for destination
type AirbyteStreamAndConfiguration struct {
	// the immutable schema defined by the source
	Stream *AirbyteStream `json:"stream,omitempty"`
	// the mutable part of the stream to configure the destination
	Config *AirbyteStreamConfiguration `json:"config,omitempty"`
}

func (o *AirbyteStreamAndConfiguration) GetStream() *AirbyteStream {
	if o == nil {
		return nil
	}
	return o.Stream
}

func (o *AirbyteStreamAndConfiguration) GetConfig() *AirbyteStreamConfiguration {
	if o == nil {
		return nil
	}
	return o.Config
}
