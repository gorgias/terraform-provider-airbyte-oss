// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type NotificationTrigger string

const (
	NotificationTriggerSyncSuccess                    NotificationTrigger = "sync_success"
	NotificationTriggerSyncFailure                    NotificationTrigger = "sync_failure"
	NotificationTriggerSyncDisabled                   NotificationTrigger = "sync_disabled"
	NotificationTriggerSyncDisabledWarning            NotificationTrigger = "sync_disabled_warning"
	NotificationTriggerConnectionUpdate               NotificationTrigger = "connection_update"
	NotificationTriggerConnectionUpdateActionRequired NotificationTrigger = "connection_update_action_required"
	NotificationTriggerBreakingChangeWarning          NotificationTrigger = "breaking_change_warning"
	NotificationTriggerBreakingChangeSyncsDisabled    NotificationTrigger = "breaking_change_syncs_disabled"
)

func (e NotificationTrigger) ToPointer() *NotificationTrigger {
	return &e
}

func (e *NotificationTrigger) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sync_success":
		fallthrough
	case "sync_failure":
		fallthrough
	case "sync_disabled":
		fallthrough
	case "sync_disabled_warning":
		fallthrough
	case "connection_update":
		fallthrough
	case "connection_update_action_required":
		fallthrough
	case "breaking_change_warning":
		fallthrough
	case "breaking_change_syncs_disabled":
		*e = NotificationTrigger(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTrigger: %v", v)
	}
}