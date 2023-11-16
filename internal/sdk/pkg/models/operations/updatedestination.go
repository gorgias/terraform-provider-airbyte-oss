// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	DestinationRead *shared.DestinationRead
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
}

func (o *UpdateDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDestinationResponse) GetDestinationRead() *shared.DestinationRead {
	if o == nil {
		return nil
	}
	return o.DestinationRead
}

func (o *UpdateDestinationResponse) GetInvalidInputExceptionInfo() *shared.InvalidInputExceptionInfo {
	if o == nil {
		return nil
	}
	return o.InvalidInputExceptionInfo
}
