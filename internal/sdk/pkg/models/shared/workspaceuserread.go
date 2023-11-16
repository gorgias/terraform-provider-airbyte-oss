// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WorkspaceUserRead struct {
	// Caption name for the user
	Name *string `json:"name,omitempty"`
	// Internal Airbyte user ID
	UserID             string `json:"userId"`
	IsDefaultWorkspace *bool  `json:"isDefaultWorkspace,omitempty"`
	Email              string `json:"email"`
	PermissionID       string `json:"permissionId"`
	// Describes what actions/endpoints the permission entitles to
	PermissionType PermissionType `json:"permissionType"`
	WorkspaceID    string         `json:"workspaceId"`
}

func (o *WorkspaceUserRead) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WorkspaceUserRead) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *WorkspaceUserRead) GetIsDefaultWorkspace() *bool {
	if o == nil {
		return nil
	}
	return o.IsDefaultWorkspace
}

func (o *WorkspaceUserRead) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *WorkspaceUserRead) GetPermissionID() string {
	if o == nil {
		return ""
	}
	return o.PermissionID
}

func (o *WorkspaceUserRead) GetPermissionType() PermissionType {
	if o == nil {
		return PermissionType("")
	}
	return o.PermissionType
}

func (o *WorkspaceUserRead) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}