// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

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
var _ resource.Resource = &ConnectionResource{}
var _ resource.ResourceWithImportState = &ConnectionResource{}

func NewConnectionResource() resource.Resource {
	return &ConnectionResource{}
}

// ConnectionResource defines the resource implementation.
type ConnectionResource struct {
	client *sdk.Airbyte
}

// ConnectionResourceModel describes the resource data model.
type ConnectionResourceModel struct {
	BreakingChange               types.Bool              `tfsdk:"breaking_change"`
	ConnectionID                 types.String            `tfsdk:"connection_id"`
	DestinationID                types.String            `tfsdk:"destination_id"`
	Geography                    types.String            `tfsdk:"geography"`
	Name                         types.String            `tfsdk:"name"`
	NamespaceDefinition          types.String            `tfsdk:"namespace_definition"`
	NamespaceFormat              types.String            `tfsdk:"namespace_format"`
	NonBreakingChangesPreference types.String            `tfsdk:"non_breaking_changes_preference"`
	NotifySchemaChanges          types.Bool              `tfsdk:"notify_schema_changes"`
	NotifySchemaChangesByEmail   types.Bool              `tfsdk:"notify_schema_changes_by_email"`
	OperationIds                 []types.String          `tfsdk:"operation_ids"`
	Prefix                       types.String            `tfsdk:"prefix"`
	ResourceRequirements         *ResourceRequirements   `tfsdk:"resource_requirements"`
	Schedule                     *ConnectionSchedule     `tfsdk:"schedule"`
	ScheduleData                 *ConnectionScheduleData `tfsdk:"schedule_data"`
	ScheduleType                 types.String            `tfsdk:"schedule_type"`
	SourceCatalogID              types.String            `tfsdk:"source_catalog_id"`
	SourceID                     types.String            `tfsdk:"source_id"`
	Status                       types.String            `tfsdk:"status"`
	SyncCatalog                  *AirbyteCatalog         `tfsdk:"sync_catalog"`
	WorkspaceID                  types.String            `tfsdk:"workspace_id"`
}

func (r *ConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection"
}

func (r *ConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Connection Resource",

		Attributes: map[string]schema.Attribute{
			"breaking_change": schema.BoolAttribute{
				Computed: true,
			},
			"connection_id": schema.StringAttribute{
				Computed: true,
			},
			"destination_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"geography": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"auto",
						"us",
						"eu",
					),
				},
				Description: `must be one of [auto, us, eu]`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Optional name of the connection`,
			},
			"namespace_definition": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"source",
						"destination",
						"customformat",
					),
				},
				MarkdownDescription: `must be one of [source, destination, customformat]` + "\n" +
					`Method used for computing final namespace in destination`,
			},
			"namespace_format": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'.`,
			},
			"non_breaking_changes_preference": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ignore",
						"disable",
						"propagate_columns",
						"propagate_fully",
					),
				},
				Description: `must be one of [ignore, disable, propagate_columns, propagate_fully]`,
			},
			"notify_schema_changes": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"notify_schema_changes_by_email": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"operation_ids": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"prefix": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Prefix that will be prepended to the name of each stream when it is written to the destination.`,
			},
			"resource_requirements": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"cpu_limit": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"cpu_request": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"memory_limit": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"memory_request": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `optional resource requirements to run workers (blank for unbounded allocations)`,
			},
			"schedule": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"time_unit": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"minutes",
								"hours",
								"days",
								"weeks",
								"months",
							),
						},
						Description: `must be one of [minutes, hours, days, weeks, months]`,
					},
					"units": schema.Int64Attribute{
						Required: true,
					},
				},
				Description: `if null, then no schedule is set.`,
			},
			"schedule_data": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"basic_schedule": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"time_unit": schema.StringAttribute{
								Required: true,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"minutes",
										"hours",
										"days",
										"weeks",
										"months",
									),
								},
								Description: `must be one of [minutes, hours, days, weeks, months]`,
							},
							"units": schema.Int64Attribute{
								Required: true,
							},
						},
					},
					"cron": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"cron_expression": schema.StringAttribute{
								Required: true,
							},
							"cron_time_zone": schema.StringAttribute{
								Required: true,
							},
						},
					},
				},
				Description: `schedule for when the the connection should run, per the schedule type`,
			},
			"schedule_type": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"manual",
						"basic",
						"cron",
					),
				},
				MarkdownDescription: `must be one of [manual, basic, cron]` + "\n" +
					`determine how the schedule data should be interpreted`,
			},
			"source_catalog_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"source_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"status": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"active",
						"inactive",
						"deprecated",
					),
				},
				MarkdownDescription: `must be one of [active, inactive, deprecated]` + "\n" +
					`Active means that data is flowing through the connection. Inactive means it is not. Deprecated means the connection is off and cannot be re-activated. the schema field describes the elements of the schema that will be synced.`,
			},
			"sync_catalog": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"streams": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config": schema.SingleNestedAttribute{
									Computed: true,
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"alias_name": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Alias name to the stream to be used in the destination`,
										},
										"cursor_field": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
											Description: `Path to the field that will be used to determine if a record is new or modified since the last sync. This field is REQUIRED if ` + "`" + `sync_mode` + "`" + ` is ` + "`" + `incremental` + "`" + `. Otherwise it is ignored.`,
										},
										"destination_sync_mode": schema.StringAttribute{
											Required: true,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"append",
													"overwrite",
													"append_dedup",
												),
											},
											Description: `must be one of [append, overwrite, append_dedup]`,
										},
										"field_selection_enabled": schema.BoolAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Whether field selection should be enabled. If this is true, only the properties in ` + "`" + `selectedFields` + "`" + ` will be included.`,
										},
										"primary_key": schema.ListAttribute{
											Computed: true,
											Optional: true,
											ElementType: types.ListType{
												ElemType: types.StringType,
											},
											Description: `Paths to the fields that will be used as primary key. This field is REQUIRED if ` + "`" + `destination_sync_mode` + "`" + ` is ` + "`" + `*_dedup` + "`" + `. Otherwise it is ignored.`,
										},
										"selected": schema.BoolAttribute{
											Computed:    true,
											Optional:    true,
											Description: `If this is true, the stream is selected with all of its properties. For new connections, this considers if the stream is suggested or not`,
										},
										"selected_fields": schema.ListNestedAttribute{
											Computed: true,
											Optional: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"field_path": schema.ListAttribute{
														Computed:    true,
														Optional:    true,
														ElementType: types.StringType,
													},
												},
											},
											Description: `Paths to the fields that will be included in the configured catalog. This must be set if ` + "`" + `fieldSelectedEnabled` + "`" + ` is set. An empty list indicates that no properties will be included.`,
										},
										"suggested": schema.BoolAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Does the connector suggest that this stream be enabled by default?`,
										},
										"sync_mode": schema.StringAttribute{
											Required: true,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"full_refresh",
													"incremental",
												),
											},
											Description: `must be one of [full_refresh, incremental]`,
										},
									},
									Description: `the mutable part of the stream to configure the destination`,
								},
								"stream": schema.SingleNestedAttribute{
									Computed: true,
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"default_cursor_field": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
											Description: `Path to the field that will be used to determine if a record is new or modified since the last sync. If not provided by the source, the end user will have to specify the comparable themselves.`,
										},
										"json_schema": schema.SingleNestedAttribute{
											Computed:    true,
											Optional:    true,
											Attributes:  map[string]schema.Attribute{},
											Description: `Stream schema using Json Schema specs.`,
										},
										"name": schema.StringAttribute{
											Required:    true,
											Description: `Stream's name.`,
										},
										"namespace": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Optional Source-defined namespace. Airbyte streams from the same sources should have the same namespace. Currently only used by JDBC destinations to determine what schema to write to.`,
										},
										"source_defined_cursor": schema.BoolAttribute{
											Computed:    true,
											Optional:    true,
											Description: `If the source defines the cursor field, then any other cursor field inputs will be ignored. If it does not, either the user_provided one is used, or the default one is used as a backup.`,
										},
										"source_defined_primary_key": schema.ListAttribute{
											Computed: true,
											Optional: true,
											ElementType: types.ListType{
												ElemType: types.StringType,
											},
											Description: `If the source defines the primary key, paths to the fields that will be used as a primary key. If not provided by the source, the end user will have to specify the primary key themselves.`,
										},
										"supported_sync_modes": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
									},
									Description: `the immutable schema defined by the source`,
								},
							},
						},
					},
				},
				Description: `describes the available schema (catalog).`,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Airbyte)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Airbyte, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ConnectionResourceModel
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
	res, err := r.client.Connection.CreateConnection(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	if res.ConnectionRead == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.ConnectionRead)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ConnectionResourceModel
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

func (r *ConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ConnectionResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToUpdateSDKType()
	res, err := r.client.Connection.UpdateConnection(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	if res.ConnectionRead == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.ConnectionRead)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ConnectionResourceModel
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
	res, err := r.client.Connection.DeleteConnection(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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

func (r *ConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource connection.")
}