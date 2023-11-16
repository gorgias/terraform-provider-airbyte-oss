// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type ListResourcesForWorkspacesRequestBody struct {
	WorkspaceIds   []string   `json:"workspaceIds"`
	IncludeDeleted *bool      `default:"false" json:"includeDeleted"`
	Pagination     Pagination `json:"pagination"`
	NameContains   *string    `json:"nameContains,omitempty"`
}

func (l ListResourcesForWorkspacesRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListResourcesForWorkspacesRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListResourcesForWorkspacesRequestBody) GetWorkspaceIds() []string {
	if o == nil {
		return []string{}
	}
	return o.WorkspaceIds
}

func (o *ListResourcesForWorkspacesRequestBody) GetIncludeDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeleted
}

func (o *ListResourcesForWorkspacesRequestBody) GetPagination() Pagination {
	if o == nil {
		return Pagination{}
	}
	return o.Pagination
}

func (o *ListResourcesForWorkspacesRequestBody) GetNameContains() *string {
	if o == nil {
		return nil
	}
	return o.NameContains
}
