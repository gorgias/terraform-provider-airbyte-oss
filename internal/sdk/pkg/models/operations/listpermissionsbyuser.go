// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListPermissionsByUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	PermissionReadList *shared.PermissionReadList
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
}

func (o *ListPermissionsByUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPermissionsByUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPermissionsByUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPermissionsByUserResponse) GetPermissionReadList() *shared.PermissionReadList {
	if o == nil {
		return nil
	}
	return o.PermissionReadList
}

func (o *ListPermissionsByUserResponse) GetInvalidInputExceptionInfo() *shared.InvalidInputExceptionInfo {
	if o == nil {
		return nil
	}
	return o.InvalidInputExceptionInfo
}
