// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type HealthCheckRead struct {
	Available bool `json:"available"`
}

func (o *HealthCheckRead) GetAvailable() bool {
	if o == nil {
		return false
	}
	return o.Available
}
