// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DestinationDefinitionSpecificationRead - Successful operation
type DestinationDefinitionSpecificationRead struct {
	AdvancedAuth *AdvancedAuth `json:"advancedAuth,omitempty"`
	// The specification for what values are required to configure the destinationDefinition.
	ConnectionSpecification       *DestinationDefinitionSpecification `json:"connectionSpecification,omitempty"`
	DestinationDefinitionID       string                              `json:"destinationDefinitionId"`
	DocumentationURL              *string                             `json:"documentationUrl,omitempty"`
	JobInfo                       SynchronousJobRead                  `json:"jobInfo"`
	SupportedDestinationSyncModes []DestinationSyncMode               `json:"supportedDestinationSyncModes,omitempty"`
}
