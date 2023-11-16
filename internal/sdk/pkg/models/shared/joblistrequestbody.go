// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobListRequestBody struct {
	ConfigTypes    []JobConfigType `json:"configTypes"`
	ConfigID       string          `json:"configId"`
	IncludingJobID *int64          `json:"includingJobId,omitempty"`
	Pagination     *Pagination     `json:"pagination,omitempty"`
}

func (o *JobListRequestBody) GetConfigTypes() []JobConfigType {
	if o == nil {
		return []JobConfigType{}
	}
	return o.ConfigTypes
}

func (o *JobListRequestBody) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}

func (o *JobListRequestBody) GetIncludingJobID() *int64 {
	if o == nil {
		return nil
	}
	return o.IncludingJobID
}

func (o *JobListRequestBody) GetPagination() *Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
