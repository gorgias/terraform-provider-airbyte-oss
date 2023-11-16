// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationDefinitionUpdate struct {
	DestinationDefinitionID string  `json:"destinationDefinitionId"`
	DockerImageTag          *string `json:"dockerImageTag,omitempty"`
	// actor definition specific resource requirements. if default is set, these are the requirements that should be set for ALL jobs run for this actor definition. it is overriden by the job type specific configurations. if not set, the platform will use defaults. these values will be overriden by configuration at the connection level.
	ResourceRequirements *ActorDefinitionResourceRequirements `json:"resourceRequirements,omitempty"`
}

func (o *DestinationDefinitionUpdate) GetDestinationDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DestinationDefinitionID
}

func (o *DestinationDefinitionUpdate) GetDockerImageTag() *string {
	if o == nil {
		return nil
	}
	return o.DockerImageTag
}

func (o *DestinationDefinitionUpdate) GetResourceRequirements() *ActorDefinitionResourceRequirements {
	if o == nil {
		return nil
	}
	return o.ResourceRequirements
}
