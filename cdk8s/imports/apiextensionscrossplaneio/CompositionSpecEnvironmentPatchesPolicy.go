package apiextensionscrossplaneio


// Policy configures the specifics of patching behaviour.
type CompositionSpecEnvironmentPatchesPolicy struct {
	// FromFieldPath specifies how to patch from a field path.
	//
	// The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.
	FromFieldPath CompositionSpecEnvironmentPatchesPolicyFromFieldPath `field:"optional" json:"fromFieldPath" yaml:"fromFieldPath"`
	// MergeOptions Specifies merge options on a field path.
	MergeOptions *CompositionSpecEnvironmentPatchesPolicyMergeOptions `field:"optional" json:"mergeOptions" yaml:"mergeOptions"`
}

