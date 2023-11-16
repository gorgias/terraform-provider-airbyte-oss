// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &WorkspaceResource{}
var _ resource.ResourceWithImportState = &WorkspaceResource{}

func NewWorkspaceResource() resource.Resource {
	return &WorkspaceResource{}
}

// WorkspaceResource defines the resource implementation.
type WorkspaceResource struct {
	client *sdk.SDK
}

// WorkspaceResourceModel describes the resource data model.
type WorkspaceResourceModel struct {
	AnonymousDataCollection types.Bool            `tfsdk:"anonymous_data_collection"`
	CustomerID              types.String          `tfsdk:"customer_id"`
	DefaultGeography        types.String          `tfsdk:"default_geography"`
	DisplaySetupWizard      types.Bool            `tfsdk:"display_setup_wizard"`
	Email                   types.String          `tfsdk:"email"`
	FeedbackDone            types.Bool            `tfsdk:"feedback_done"`
	FirstCompletedSync      types.Bool            `tfsdk:"first_completed_sync"`
	InitialSetupComplete    types.Bool            `tfsdk:"initial_setup_complete"`
	Name                    types.String          `tfsdk:"name"`
	News                    types.Bool            `tfsdk:"news"`
	Notifications           []Notification        `tfsdk:"notifications"`
	NotificationSettings    *NotificationSettings `tfsdk:"notification_settings"`
	SecurityUpdates         types.Bool            `tfsdk:"security_updates"`
	Slug                    types.String          `tfsdk:"slug"`
	WebhookConfigs          []WebhookConfigRead   `tfsdk:"webhook_configs"`
	WorkspaceID             types.String          `tfsdk:"workspace_id"`
}

func (r *WorkspaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace"
}

func (r *WorkspaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Workspace Resource",

		Attributes: map[string]schema.Attribute{
			"anonymous_data_collection": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"customer_id": schema.StringAttribute{
				Computed: true,
			},
			"default_geography": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `must be one of ["auto", "us", "eu"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"auto",
						"us",
						"eu",
					),
				},
			},
			"display_setup_wizard": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"email": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"feedback_done": schema.BoolAttribute{
				Computed: true,
			},
			"first_completed_sync": schema.BoolAttribute{
				Computed: true,
			},
			"initial_setup_complete": schema.BoolAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"news": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"notifications": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"customerio_configuration": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							Attributes: map[string]schema.Attribute{},
						},
						"notification_type": schema.StringAttribute{
							Required:    true,
							Description: `must be one of ["slack", "customerio"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"slack",
									"customerio",
								),
							},
						},
						"send_on_failure": schema.BoolAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Default: true`,
						},
						"send_on_success": schema.BoolAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Default: false`,
						},
						"slack_configuration": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"webhook": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
				},
			},
			"notification_settings": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"send_on_connection_update": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"customerio_configuration": schema.SingleNestedAttribute{
								Computed:   true,
								Optional:   true,
								Attributes: map[string]schema.Attribute{},
							},
							"notification_type": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"slack_configuration": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"webhook": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
					},
					"send_on_connection_update_action_required": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"customerio_configuration": schema.SingleNestedAttribute{
								Computed:   true,
								Optional:   true,
								Attributes: map[string]schema.Attribute{},
							},
							"notification_type": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"slack_configuration": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"webhook": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
					},
					"send_on_failure": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"customerio_configuration": schema.SingleNestedAttribute{
								Computed:   true,
								Optional:   true,
								Attributes: map[string]schema.Attribute{},
							},
							"notification_type": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"slack_configuration": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"webhook": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
					},
					"send_on_success": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"customerio_configuration": schema.SingleNestedAttribute{
								Computed:   true,
								Optional:   true,
								Attributes: map[string]schema.Attribute{},
							},
							"notification_type": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"slack_configuration": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"webhook": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
					},
					"send_on_sync_disabled": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"customerio_configuration": schema.SingleNestedAttribute{
								Computed:   true,
								Optional:   true,
								Attributes: map[string]schema.Attribute{},
							},
							"notification_type": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"slack_configuration": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"webhook": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
					},
					"send_on_sync_disabled_warning": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"customerio_configuration": schema.SingleNestedAttribute{
								Computed:   true,
								Optional:   true,
								Attributes: map[string]schema.Attribute{},
							},
							"notification_type": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"slack_configuration": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"webhook": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"security_updates": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"slug": schema.StringAttribute{
				Computed: true,
			},
			"webhook_configs": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"auth_token": schema.StringAttribute{
							Optional:    true,
							Description: `an auth token, to be passed as the value for an HTTP Authorization header.`,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `human readable name for this webhook e.g. for UI display.`,
						},
						"validation_url": schema.StringAttribute{
							Optional:    true,
							Description: `if supplied, the webhook config will be validated by checking that this URL returns a 2xx response.`,
						},
					},
				},
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *WorkspaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WorkspaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkspaceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToCreateSDKType()
	res, err := r.client.Workspace.CreateWorkspace(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.WorkspaceRead == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.WorkspaceRead)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkspaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkspaceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkspaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkspaceResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToUpdateSDKType()
	res, err := r.client.Workspace.UpdateWorkspace(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.WorkspaceRead == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.WorkspaceRead)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkspaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkspaceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToDeleteSDKType()
	res, err := r.client.Workspace.DeleteWorkspace(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *WorkspaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource workspace.")
}
