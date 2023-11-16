// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeclarativeManifestRead struct {
	// Low code CDK manifest JSON object
	Manifest    *DeclarativeManifest `json:"manifest,omitempty"`
	IsDraft     *bool                `json:"isDraft,omitempty"`
	Version     *int64               `json:"version,omitempty"`
	Description *string              `json:"description,omitempty"`
}

func (o *DeclarativeManifestRead) GetManifest() *DeclarativeManifest {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *DeclarativeManifestRead) GetIsDraft() *bool {
	if o == nil {
		return nil
	}
	return o.IsDraft
}

func (o *DeclarativeManifestRead) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *DeclarativeManifestRead) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
