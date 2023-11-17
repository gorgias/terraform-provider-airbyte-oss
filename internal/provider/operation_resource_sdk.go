// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *OperationResourceModel) ToCreateSDKType() *shared.OperationCreate {
	workspaceID := r.WorkspaceID.ValueString()
	name := r.Name.ValueString()
	operatorType := shared.OperatorType(r.OperatorConfiguration.OperatorType.ValueString())
	var normalization *shared.OperatorNormalization
	if r.OperatorConfiguration.Normalization != nil {
		option := new(shared.Option)
		if !r.OperatorConfiguration.Normalization.Option.IsUnknown() && !r.OperatorConfiguration.Normalization.Option.IsNull() {
			*option = shared.Option(r.OperatorConfiguration.Normalization.Option.ValueString())
		} else {
			option = nil
		}
		normalization = &shared.OperatorNormalization{
			Option: option,
		}
	}
	var dbt *shared.OperatorDbt
	if r.OperatorConfiguration.Dbt != nil {
		gitRepoURL := r.OperatorConfiguration.Dbt.GitRepoURL.ValueString()
		gitRepoBranch := new(string)
		if !r.OperatorConfiguration.Dbt.GitRepoBranch.IsUnknown() && !r.OperatorConfiguration.Dbt.GitRepoBranch.IsNull() {
			*gitRepoBranch = r.OperatorConfiguration.Dbt.GitRepoBranch.ValueString()
		} else {
			gitRepoBranch = nil
		}
		dockerImage := new(string)
		if !r.OperatorConfiguration.Dbt.DockerImage.IsUnknown() && !r.OperatorConfiguration.Dbt.DockerImage.IsNull() {
			*dockerImage = r.OperatorConfiguration.Dbt.DockerImage.ValueString()
		} else {
			dockerImage = nil
		}
		dbtArguments := new(string)
		if !r.OperatorConfiguration.Dbt.DbtArguments.IsUnknown() && !r.OperatorConfiguration.Dbt.DbtArguments.IsNull() {
			*dbtArguments = r.OperatorConfiguration.Dbt.DbtArguments.ValueString()
		} else {
			dbtArguments = nil
		}
		dbt = &shared.OperatorDbt{
			GitRepoURL:    gitRepoURL,
			GitRepoBranch: gitRepoBranch,
			DockerImage:   dockerImage,
			DbtArguments:  dbtArguments,
		}
	}
	var webhook *shared.OperatorWebhook
	if r.OperatorConfiguration.Webhook != nil {
		webhookConfigID := new(string)
		if !r.OperatorConfiguration.Webhook.WebhookConfigID.IsUnknown() && !r.OperatorConfiguration.Webhook.WebhookConfigID.IsNull() {
			*webhookConfigID = r.OperatorConfiguration.Webhook.WebhookConfigID.ValueString()
		} else {
			webhookConfigID = nil
		}
		webhookType := new(shared.WebhookType)
		if !r.OperatorConfiguration.Webhook.WebhookType.IsUnknown() && !r.OperatorConfiguration.Webhook.WebhookType.IsNull() {
			*webhookType = shared.WebhookType(r.OperatorConfiguration.Webhook.WebhookType.ValueString())
		} else {
			webhookType = nil
		}
		var dbtCloud *shared.DbtCloud
		if r.OperatorConfiguration.Webhook.DbtCloud != nil {
			accountID := r.OperatorConfiguration.Webhook.DbtCloud.AccountID.ValueInt64()
			jobID := r.OperatorConfiguration.Webhook.DbtCloud.JobID.ValueInt64()
			dbtCloud = &shared.DbtCloud{
				AccountID: accountID,
				JobID:     jobID,
			}
		}
		executionURL := new(string)
		if !r.OperatorConfiguration.Webhook.ExecutionURL.IsUnknown() && !r.OperatorConfiguration.Webhook.ExecutionURL.IsNull() {
			*executionURL = r.OperatorConfiguration.Webhook.ExecutionURL.ValueString()
		} else {
			executionURL = nil
		}
		executionBody := new(string)
		if !r.OperatorConfiguration.Webhook.ExecutionBody.IsUnknown() && !r.OperatorConfiguration.Webhook.ExecutionBody.IsNull() {
			*executionBody = r.OperatorConfiguration.Webhook.ExecutionBody.ValueString()
		} else {
			executionBody = nil
		}
		webhook = &shared.OperatorWebhook{
			WebhookConfigID: webhookConfigID,
			WebhookType:     webhookType,
			DbtCloud:        dbtCloud,
			ExecutionURL:    executionURL,
			ExecutionBody:   executionBody,
		}
	}
	operatorConfiguration := shared.OperatorConfiguration{
		OperatorType:  operatorType,
		Normalization: normalization,
		Dbt:           dbt,
		Webhook:       webhook,
	}
	out := shared.OperationCreate{
		WorkspaceID:           workspaceID,
		Name:                  name,
		OperatorConfiguration: operatorConfiguration,
	}
	return &out
}

func (r *OperationResourceModel) ToGetSDKType() *shared.OperationCreate {
	out := r.ToCreateSDKType()
	return out
}

func (r *OperationResourceModel) ToUpdateSDKType() *shared.OperationUpdate {
	operationID := r.OperationID.ValueString()
	name := r.Name.ValueString()
	operatorType := shared.OperatorType(r.OperatorConfiguration.OperatorType.ValueString())
	var normalization *shared.OperatorNormalization
	if r.OperatorConfiguration.Normalization != nil {
		option := new(shared.Option)
		if !r.OperatorConfiguration.Normalization.Option.IsUnknown() && !r.OperatorConfiguration.Normalization.Option.IsNull() {
			*option = shared.Option(r.OperatorConfiguration.Normalization.Option.ValueString())
		} else {
			option = nil
		}
		normalization = &shared.OperatorNormalization{
			Option: option,
		}
	}
	var dbt *shared.OperatorDbt
	if r.OperatorConfiguration.Dbt != nil {
		gitRepoURL := r.OperatorConfiguration.Dbt.GitRepoURL.ValueString()
		gitRepoBranch := new(string)
		if !r.OperatorConfiguration.Dbt.GitRepoBranch.IsUnknown() && !r.OperatorConfiguration.Dbt.GitRepoBranch.IsNull() {
			*gitRepoBranch = r.OperatorConfiguration.Dbt.GitRepoBranch.ValueString()
		} else {
			gitRepoBranch = nil
		}
		dockerImage := new(string)
		if !r.OperatorConfiguration.Dbt.DockerImage.IsUnknown() && !r.OperatorConfiguration.Dbt.DockerImage.IsNull() {
			*dockerImage = r.OperatorConfiguration.Dbt.DockerImage.ValueString()
		} else {
			dockerImage = nil
		}
		dbtArguments := new(string)
		if !r.OperatorConfiguration.Dbt.DbtArguments.IsUnknown() && !r.OperatorConfiguration.Dbt.DbtArguments.IsNull() {
			*dbtArguments = r.OperatorConfiguration.Dbt.DbtArguments.ValueString()
		} else {
			dbtArguments = nil
		}
		dbt = &shared.OperatorDbt{
			GitRepoURL:    gitRepoURL,
			GitRepoBranch: gitRepoBranch,
			DockerImage:   dockerImage,
			DbtArguments:  dbtArguments,
		}
	}
	var webhook *shared.OperatorWebhook
	if r.OperatorConfiguration.Webhook != nil {
		webhookConfigID := new(string)
		if !r.OperatorConfiguration.Webhook.WebhookConfigID.IsUnknown() && !r.OperatorConfiguration.Webhook.WebhookConfigID.IsNull() {
			*webhookConfigID = r.OperatorConfiguration.Webhook.WebhookConfigID.ValueString()
		} else {
			webhookConfigID = nil
		}
		webhookType := new(shared.WebhookType)
		if !r.OperatorConfiguration.Webhook.WebhookType.IsUnknown() && !r.OperatorConfiguration.Webhook.WebhookType.IsNull() {
			*webhookType = shared.WebhookType(r.OperatorConfiguration.Webhook.WebhookType.ValueString())
		} else {
			webhookType = nil
		}
		var dbtCloud *shared.DbtCloud
		if r.OperatorConfiguration.Webhook.DbtCloud != nil {
			accountID := r.OperatorConfiguration.Webhook.DbtCloud.AccountID.ValueInt64()
			jobID := r.OperatorConfiguration.Webhook.DbtCloud.JobID.ValueInt64()
			dbtCloud = &shared.DbtCloud{
				AccountID: accountID,
				JobID:     jobID,
			}
		}
		executionURL := new(string)
		if !r.OperatorConfiguration.Webhook.ExecutionURL.IsUnknown() && !r.OperatorConfiguration.Webhook.ExecutionURL.IsNull() {
			*executionURL = r.OperatorConfiguration.Webhook.ExecutionURL.ValueString()
		} else {
			executionURL = nil
		}
		executionBody := new(string)
		if !r.OperatorConfiguration.Webhook.ExecutionBody.IsUnknown() && !r.OperatorConfiguration.Webhook.ExecutionBody.IsNull() {
			*executionBody = r.OperatorConfiguration.Webhook.ExecutionBody.ValueString()
		} else {
			executionBody = nil
		}
		webhook = &shared.OperatorWebhook{
			WebhookConfigID: webhookConfigID,
			WebhookType:     webhookType,
			DbtCloud:        dbtCloud,
			ExecutionURL:    executionURL,
			ExecutionBody:   executionBody,
		}
	}
	operatorConfiguration := shared.OperatorConfiguration{
		OperatorType:  operatorType,
		Normalization: normalization,
		Dbt:           dbt,
		Webhook:       webhook,
	}
	out := shared.OperationUpdate{
		OperationID:           operationID,
		Name:                  name,
		OperatorConfiguration: operatorConfiguration,
	}
	return &out
}

func (r *OperationResourceModel) ToDeleteSDKType() *shared.OperationCreate {
	out := r.ToCreateSDKType()
	return out
}

func (r *OperationResourceModel) RefreshFromGetResponse(resp *shared.OperationRead) {
	r.Name = types.StringValue(resp.Name)
	r.OperationID = types.StringValue(resp.OperationID)
	if resp.OperatorConfiguration.Dbt == nil {
		r.OperatorConfiguration.Dbt = nil
	} else {
		r.OperatorConfiguration.Dbt = &OperatorDbt{}
		if resp.OperatorConfiguration.Dbt.DbtArguments != nil {
			r.OperatorConfiguration.Dbt.DbtArguments = types.StringValue(*resp.OperatorConfiguration.Dbt.DbtArguments)
		} else {
			r.OperatorConfiguration.Dbt.DbtArguments = types.StringNull()
		}
		if resp.OperatorConfiguration.Dbt.DockerImage != nil {
			r.OperatorConfiguration.Dbt.DockerImage = types.StringValue(*resp.OperatorConfiguration.Dbt.DockerImage)
		} else {
			r.OperatorConfiguration.Dbt.DockerImage = types.StringNull()
		}
		if resp.OperatorConfiguration.Dbt.GitRepoBranch != nil {
			r.OperatorConfiguration.Dbt.GitRepoBranch = types.StringValue(*resp.OperatorConfiguration.Dbt.GitRepoBranch)
		} else {
			r.OperatorConfiguration.Dbt.GitRepoBranch = types.StringNull()
		}
		r.OperatorConfiguration.Dbt.GitRepoURL = types.StringValue(resp.OperatorConfiguration.Dbt.GitRepoURL)
	}
	if resp.OperatorConfiguration.Normalization == nil {
		r.OperatorConfiguration.Normalization = nil
	} else {
		r.OperatorConfiguration.Normalization = &OperatorNormalization{}
		if resp.OperatorConfiguration.Normalization.Option != nil {
			r.OperatorConfiguration.Normalization.Option = types.StringValue(string(*resp.OperatorConfiguration.Normalization.Option))
		} else {
			r.OperatorConfiguration.Normalization.Option = types.StringNull()
		}
	}
	r.OperatorConfiguration.OperatorType = types.StringValue(string(resp.OperatorConfiguration.OperatorType))
	if resp.OperatorConfiguration.Webhook == nil {
		r.OperatorConfiguration.Webhook = nil
	} else {
		r.OperatorConfiguration.Webhook = &OperatorWebhook{}
		if resp.OperatorConfiguration.Webhook.DbtCloud == nil {
			r.OperatorConfiguration.Webhook.DbtCloud = nil
		} else {
			r.OperatorConfiguration.Webhook.DbtCloud = &DbtCloud{}
			r.OperatorConfiguration.Webhook.DbtCloud.AccountID = types.Int64Value(resp.OperatorConfiguration.Webhook.DbtCloud.AccountID)
			r.OperatorConfiguration.Webhook.DbtCloud.JobID = types.Int64Value(resp.OperatorConfiguration.Webhook.DbtCloud.JobID)
		}
		if resp.OperatorConfiguration.Webhook.ExecutionBody != nil {
			r.OperatorConfiguration.Webhook.ExecutionBody = types.StringValue(*resp.OperatorConfiguration.Webhook.ExecutionBody)
		} else {
			r.OperatorConfiguration.Webhook.ExecutionBody = types.StringNull()
		}
		if resp.OperatorConfiguration.Webhook.ExecutionURL != nil {
			r.OperatorConfiguration.Webhook.ExecutionURL = types.StringValue(*resp.OperatorConfiguration.Webhook.ExecutionURL)
		} else {
			r.OperatorConfiguration.Webhook.ExecutionURL = types.StringNull()
		}
		if resp.OperatorConfiguration.Webhook.WebhookConfigID != nil {
			r.OperatorConfiguration.Webhook.WebhookConfigID = types.StringValue(*resp.OperatorConfiguration.Webhook.WebhookConfigID)
		} else {
			r.OperatorConfiguration.Webhook.WebhookConfigID = types.StringNull()
		}
		if resp.OperatorConfiguration.Webhook.WebhookType != nil {
			r.OperatorConfiguration.Webhook.WebhookType = types.StringValue(string(*resp.OperatorConfiguration.Webhook.WebhookType))
		} else {
			r.OperatorConfiguration.Webhook.WebhookType = types.StringNull()
		}
	}
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *OperationResourceModel) RefreshFromCreateResponse(resp *shared.OperationRead) {
	r.RefreshFromGetResponse(resp)
}

func (r *OperationResourceModel) RefreshFromUpdateResponse(resp *shared.OperationRead) {
	r.RefreshFromGetResponse(resp)
}
