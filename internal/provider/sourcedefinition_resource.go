// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk"

	"github.com/aballiet/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceDefinitionResource{}
var _ resource.ResourceWithImportState = &SourceDefinitionResource{}

func NewSourceDefinitionResource() resource.Resource {
	return &SourceDefinitionResource{}
}

// SourceDefinitionResource defines the resource implementation.
type SourceDefinitionResource struct {
	client *sdk.SDK
}

// SourceDefinitionResourceModel describes the resource data model.
type SourceDefinitionResourceModel struct {
	Custom                    types.Bool                           `tfsdk:"custom"`
	DockerImageTag            types.String                         `tfsdk:"docker_image_tag"`
	DockerRepository          types.String                         `tfsdk:"docker_repository"`
	DocumentationURL          types.String                         `tfsdk:"documentation_url"`
	Icon                      types.String                         `tfsdk:"icon"`
	MaxSecondsBetweenMessages types.Int64                          `tfsdk:"max_seconds_between_messages"`
	Name                      types.String                         `tfsdk:"name"`
	ProtocolVersion           types.String                         `tfsdk:"protocol_version"`
	ReleaseDate               types.String                         `tfsdk:"release_date"`
	ReleaseStage              types.String                         `tfsdk:"release_stage"`
	ResourceRequirements      *ActorDefinitionResourceRequirements `tfsdk:"resource_requirements"`
	ScopeID                   types.String                         `tfsdk:"scope_id"`
	ScopeType                 types.String                         `tfsdk:"scope_type"`
	SourceDefinition          SourceDefinitionCreate               `tfsdk:"source_definition"`
	SourceDefinitionID        types.String                         `tfsdk:"source_definition_id"`
	SourceType                types.String                         `tfsdk:"source_type"`
	SupportLevel              types.String                         `tfsdk:"support_level"`
	WorkspaceID               types.String                         `tfsdk:"workspace_id"`
}

func (r *SourceDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_definition"
}

func (r *SourceDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceDefinition Resource",

		Attributes: map[string]schema.Attribute{
			"custom": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: `Default: false` + "\n" +
					`Whether the connector is custom or not`,
			},
			"docker_image_tag": schema.StringAttribute{
				Computed: true,
			},
			"docker_repository": schema.StringAttribute{
				Computed: true,
			},
			"documentation_url": schema.StringAttribute{
				Computed: true,
			},
			"icon": schema.StringAttribute{
				Computed: true,
			},
			"max_seconds_between_messages": schema.Int64Attribute{
				Computed:    true,
				Description: `Number of seconds allowed between 2 airbyte protocol messages. The source will timeout if this delay is reach`,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"protocol_version": schema.StringAttribute{
				Computed:    true,
				Description: `The Airbyte Protocol version supported by the connector`,
			},
			"release_date": schema.StringAttribute{
				Computed:    true,
				Description: `The date when this connector was first released, in yyyy-mm-dd format.`,
				Validators: []validator.String{
					validators.IsValidDate(),
				},
			},
			"release_stage": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["alpha", "beta", "generally_available", "custom"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"alpha",
						"beta",
						"generally_available",
						"custom",
					),
				},
			},
			"resource_requirements": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"cpu_limit": schema.StringAttribute{
								Computed: true,
							},
							"cpu_request": schema.StringAttribute{
								Computed: true,
							},
							"memory_limit": schema.StringAttribute{
								Computed: true,
							},
							"memory_request": schema.StringAttribute{
								Computed: true,
							},
						},
						Description: `optional resource requirements to run workers (blank for unbounded allocations)`,
					},
					"job_specific": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"job_type": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: `must be one of ["get_spec", "check_connection", "discover_schema", "sync", "reset_connection", "connection_updater", "replicate"]` + "\n" +
										`enum that describes the different types of jobs that the platform runs.`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"get_spec",
											"check_connection",
											"discover_schema",
											"sync",
											"reset_connection",
											"connection_updater",
											"replicate",
										),
									},
								},
								"resource_requirements": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"cpu_limit": schema.StringAttribute{
											Computed: true,
										},
										"cpu_request": schema.StringAttribute{
											Computed: true,
										},
										"memory_limit": schema.StringAttribute{
											Computed: true,
										},
										"memory_request": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `optional resource requirements to run workers (blank for unbounded allocations)`,
								},
							},
						},
					},
				},
				Description: `actor definition specific resource requirements. if default is set, these are the requirements that should be set for ALL jobs run for this actor definition. it is overriden by the job type specific configurations. if not set, the platform will use defaults. these values will be overriden by configuration at the connection level.`,
			},
			"scope_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional: true,
			},
			"scope_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `must be one of ["workspace", "organization"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"workspace",
						"organization",
					),
				},
			},
			"source_definition": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"docker_repository": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"docker_image_tag": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"documentation_url": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"icon": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Optional: true,
					},
					"resource_requirements": schema.SingleNestedAttribute{
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplace(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"default": schema.SingleNestedAttribute{
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"cpu_request": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"cpu_limit": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"memory_request": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"memory_limit": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
								},
								Description: `optional resource requirements to run workers (blank for unbounded allocations)`,
							},
							"job_specific": schema.ListNestedAttribute{
								PlanModifiers: []planmodifier.List{
									listplanmodifier.RequiresReplace(),
								},
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"job_type": schema.StringAttribute{
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
											Required: true,
											MarkdownDescription: `must be one of ["get_spec", "check_connection", "discover_schema", "sync", "reset_connection", "connection_updater", "replicate"]` + "\n" +
												`enum that describes the different types of jobs that the platform runs.`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"get_spec",
													"check_connection",
													"discover_schema",
													"sync",
													"reset_connection",
													"connection_updater",
													"replicate",
												),
											},
										},
										"resource_requirements": schema.SingleNestedAttribute{
											PlanModifiers: []planmodifier.Object{
												objectplanmodifier.RequiresReplace(),
											},
											Required: true,
											Attributes: map[string]schema.Attribute{
												"cpu_request": schema.StringAttribute{
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.RequiresReplace(),
													},
													Optional: true,
												},
												"cpu_limit": schema.StringAttribute{
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.RequiresReplace(),
													},
													Optional: true,
												},
												"memory_request": schema.StringAttribute{
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.RequiresReplace(),
													},
													Optional: true,
												},
												"memory_limit": schema.StringAttribute{
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.RequiresReplace(),
													},
													Optional: true,
												},
											},
											Description: `optional resource requirements to run workers (blank for unbounded allocations)`,
										},
									},
								},
							},
						},
						Description: `actor definition specific resource requirements. if default is set, these are the requirements that should be set for ALL jobs run for this actor definition. it is overriden by the job type specific configurations. if not set, the platform will use defaults. these values will be overriden by configuration at the connection level.`,
					},
				},
			},
			"source_definition_id": schema.StringAttribute{
				Computed: true,
			},
			"source_type": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["api", "file", "database", "custom"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"api",
						"file",
						"database",
						"custom",
					),
				},
			},
			"support_level": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["community", "certified", "none"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"community",
						"certified",
						"none",
					),
				},
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional: true,
			},
		},
	}
}

func (r *SourceDefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceDefinitionResourceModel
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

	request := data.ToCreateSDKType()
	res, err := r.client.SourceDefinition.CreateCustomSourceDefinition(ctx, request)
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
	if res.SourceDefinitionRead == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SourceDefinitionRead)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceDefinitionResourceModel
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

func (r *SourceDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceDefinitionResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToUpdateSDKType()
	res, err := r.client.SourceDefinition.UpdateSourceDefinition(ctx, request)
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
	if res.SourceDefinitionRead == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.SourceDefinitionRead)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceDefinitionResourceModel
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
	res, err := r.client.SourceDefinition.DeleteSourceDefinition(ctx, request)
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

func (r *SourceDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource source_definition.")
}
