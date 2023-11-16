// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobRetryStateRequestBody struct {
	ID                         *string `json:"id,omitempty"`
	ConnectionID               string  `json:"connectionId"`
	JobID                      int64   `json:"jobId"`
	SuccessiveCompleteFailures int64   `json:"successiveCompleteFailures"`
	TotalCompleteFailures      int64   `json:"totalCompleteFailures"`
	SuccessivePartialFailures  int64   `json:"successivePartialFailures"`
	TotalPartialFailures       int64   `json:"totalPartialFailures"`
}

func (o *JobRetryStateRequestBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JobRetryStateRequestBody) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *JobRetryStateRequestBody) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

func (o *JobRetryStateRequestBody) GetSuccessiveCompleteFailures() int64 {
	if o == nil {
		return 0
	}
	return o.SuccessiveCompleteFailures
}

func (o *JobRetryStateRequestBody) GetTotalCompleteFailures() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCompleteFailures
}

func (o *JobRetryStateRequestBody) GetSuccessivePartialFailures() int64 {
	if o == nil {
		return 0
	}
	return o.SuccessivePartialFailures
}

func (o *JobRetryStateRequestBody) GetTotalPartialFailures() int64 {
	if o == nil {
		return 0
	}
	return o.TotalPartialFailures
}
