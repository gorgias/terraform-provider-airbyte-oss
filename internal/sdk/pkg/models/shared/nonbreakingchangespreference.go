// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type NonBreakingChangesPreference string

const (
	NonBreakingChangesPreferenceIgnore           NonBreakingChangesPreference = "ignore"
	NonBreakingChangesPreferenceDisable          NonBreakingChangesPreference = "disable"
	NonBreakingChangesPreferencePropagateColumns NonBreakingChangesPreference = "propagate_columns"
	NonBreakingChangesPreferencePropagateFully   NonBreakingChangesPreference = "propagate_fully"
)

func (e NonBreakingChangesPreference) ToPointer() *NonBreakingChangesPreference {
	return &e
}

func (e *NonBreakingChangesPreference) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ignore":
		fallthrough
	case "disable":
		fallthrough
	case "propagate_columns":
		fallthrough
	case "propagate_fully":
		*e = NonBreakingChangesPreference(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NonBreakingChangesPreference: %v", v)
	}
}
