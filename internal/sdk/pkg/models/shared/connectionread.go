// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type ConnectionRead struct {
	ConnectionID string `json:"connectionId"`
	Name         string `json:"name"`
	// Method used for computing final namespace in destination
	NamespaceDefinition *NamespaceDefinitionType `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `default:"null" json:"namespaceFormat"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination.
	Prefix        *string  `json:"prefix,omitempty"`
	SourceID      string   `json:"sourceId"`
	DestinationID string   `json:"destinationId"`
	OperationIds  []string `json:"operationIds,omitempty"`
	// describes the available schema (catalog).
	SyncCatalog AirbyteCatalog `json:"syncCatalog"`
	// if null, then no schedule is set.
	Schedule *ConnectionSchedule `json:"schedule,omitempty"`
	// determine how the schedule data should be interpreted
	ScheduleType *ConnectionScheduleType `json:"scheduleType,omitempty"`
	// schedule for when the the connection should run, per the schedule type
	ScheduleData *ConnectionScheduleData `json:"scheduleData,omitempty"`
	// Active means that data is flowing through the connection. Inactive means it is not. Deprecated means the connection is off and cannot be re-activated. the schema field describes the elements of the schema that will be synced.
	Status ConnectionStatus `json:"status"`
	// optional resource requirements to run workers (blank for unbounded allocations)
	ResourceRequirements         *ResourceRequirements         `json:"resourceRequirements,omitempty"`
	SourceCatalogID              *string                       `json:"sourceCatalogId,omitempty"`
	Geography                    *Geography                    `json:"geography,omitempty"`
	BreakingChange               bool                          `json:"breakingChange"`
	NotifySchemaChanges          *bool                         `json:"notifySchemaChanges,omitempty"`
	NotifySchemaChangesByEmail   *bool                         `json:"notifySchemaChangesByEmail,omitempty"`
	NonBreakingChangesPreference *NonBreakingChangesPreference `json:"nonBreakingChangesPreference,omitempty"`
	WorkspaceID                  *string                       `json:"workspaceId,omitempty"`
}

func (c ConnectionRead) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionRead) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionRead) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ConnectionRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionRead) GetNamespaceDefinition() *NamespaceDefinitionType {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *ConnectionRead) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *ConnectionRead) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *ConnectionRead) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *ConnectionRead) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *ConnectionRead) GetOperationIds() []string {
	if o == nil {
		return nil
	}
	return o.OperationIds
}

func (o *ConnectionRead) GetSyncCatalog() AirbyteCatalog {
	if o == nil {
		return AirbyteCatalog{}
	}
	return o.SyncCatalog
}

func (o *ConnectionRead) GetSchedule() *ConnectionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *ConnectionRead) GetScheduleType() *ConnectionScheduleType {
	if o == nil {
		return nil
	}
	return o.ScheduleType
}

func (o *ConnectionRead) GetScheduleData() *ConnectionScheduleData {
	if o == nil {
		return nil
	}
	return o.ScheduleData
}

func (o *ConnectionRead) GetStatus() ConnectionStatus {
	if o == nil {
		return ConnectionStatus("")
	}
	return o.Status
}

func (o *ConnectionRead) GetResourceRequirements() *ResourceRequirements {
	if o == nil {
		return nil
	}
	return o.ResourceRequirements
}

func (o *ConnectionRead) GetSourceCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.SourceCatalogID
}

func (o *ConnectionRead) GetGeography() *Geography {
	if o == nil {
		return nil
	}
	return o.Geography
}

func (o *ConnectionRead) GetBreakingChange() bool {
	if o == nil {
		return false
	}
	return o.BreakingChange
}

func (o *ConnectionRead) GetNotifySchemaChanges() *bool {
	if o == nil {
		return nil
	}
	return o.NotifySchemaChanges
}

func (o *ConnectionRead) GetNotifySchemaChangesByEmail() *bool {
	if o == nil {
		return nil
	}
	return o.NotifySchemaChangesByEmail
}

func (o *ConnectionRead) GetNonBreakingChangesPreference() *NonBreakingChangesPreference {
	if o == nil {
		return nil
	}
	return o.NonBreakingChangesPreference
}

func (o *ConnectionRead) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}
