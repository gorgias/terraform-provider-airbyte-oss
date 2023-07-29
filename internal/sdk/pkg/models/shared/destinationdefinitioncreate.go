// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationDefinitionCreate struct {
	DockerImageTag   string  `json:"dockerImageTag"`
	DockerRepository string  `json:"dockerRepository"`
	DocumentationURL string  `json:"documentationUrl"`
	Icon             *string `json:"icon,omitempty"`
	Name             string  `json:"name"`
	// actor definition specific resource requirements. if default is set, these are the requirements that should be set for ALL jobs run for this actor definition. it is overriden by the job type specific configurations. if not set, the platform will use defaults. these values will be overriden by configuration at the connection level.
	ResourceRequirements *ActorDefinitionResourceRequirements `json:"resourceRequirements,omitempty"`
}
