// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDefinitionManifestResourceModel) ToCreateSDKType() *shared.DeclarativeSourceDefinitionCreateManifestRequestBody {
	workspaceID := r.WorkspaceID.ValueString()
	sourceDefinitionID := r.SourceDefinitionID.ValueString()
	setAsActiveManifest := r.SetAsActiveManifest.ValueBool()
	description := r.DeclarativeManifest.Description.ValueString()
	manifest := shared.DeclarativeManifest{}
	spec := shared.SourceDefinitionSpecification{}
	version := r.DeclarativeManifest.Version.ValueInt64()
	declarativeManifest := shared.DeclarativeSourceManifest{
		Description: description,
		Manifest:    manifest,
		Spec:        spec,
		Version:     version,
	}
	out := shared.DeclarativeSourceDefinitionCreateManifestRequestBody{
		WorkspaceID:         workspaceID,
		SourceDefinitionID:  sourceDefinitionID,
		SetAsActiveManifest: setAsActiveManifest,
		DeclarativeManifest: declarativeManifest,
	}
	return &out
}

func (r *SourceDefinitionManifestResourceModel) ToUpdateSDKType() *shared.UpdateActiveManifestRequestBody {
	workspaceID := r.WorkspaceID.ValueString()
	sourceDefinitionID := r.SourceDefinitionID.ValueString()
	out := shared.UpdateActiveManifestRequestBody{
		WorkspaceID:        workspaceID,
		SourceDefinitionID: sourceDefinitionID,
	}
	return &out
}

func (r *SourceDefinitionManifestResourceModel) RefreshFromCreateResponse(resp *shared.NotFoundKnownExceptionInfo) {
	if resp.ExceptionClassName != nil {
		r.ExceptionClassName = types.StringValue(*resp.ExceptionClassName)
	} else {
		r.ExceptionClassName = types.StringNull()
	}
	r.ExceptionStack = nil
	for _, v := range resp.ExceptionStack {
		r.ExceptionStack = append(r.ExceptionStack, types.StringValue(v))
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	r.Message = types.StringValue(resp.Message)
	if resp.RootCauseExceptionClassName != nil {
		r.RootCauseExceptionClassName = types.StringValue(*resp.RootCauseExceptionClassName)
	} else {
		r.RootCauseExceptionClassName = types.StringNull()
	}
	r.RootCauseExceptionStack = nil
	for _, v := range resp.RootCauseExceptionStack {
		r.RootCauseExceptionStack = append(r.RootCauseExceptionStack, types.StringValue(v))
	}
}

func (r *SourceDefinitionManifestResourceModel) RefreshFromUpdateResponse(resp *shared.NotFoundKnownExceptionInfo) {
	r.RefreshFromCreateResponse(resp)
}