// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SaveStatsRequestBody struct {
	JobID         int64                `json:"jobId"`
	AttemptNumber int                  `json:"attemptNumber"`
	Stats         AttemptStats         `json:"stats"`
	StreamStats   []AttemptStreamStats `json:"streamStats,omitempty"`
}

func (o *SaveStatsRequestBody) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

func (o *SaveStatsRequestBody) GetAttemptNumber() int {
	if o == nil {
		return 0
	}
	return o.AttemptNumber
}

func (o *SaveStatsRequestBody) GetStats() AttemptStats {
	if o == nil {
		return AttemptStats{}
	}
	return o.Stats
}

func (o *SaveStatsRequestBody) GetStreamStats() []AttemptStreamStats {
	if o == nil {
		return nil
	}
	return o.StreamStats
}
