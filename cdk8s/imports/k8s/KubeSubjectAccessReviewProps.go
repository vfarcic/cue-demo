package k8s


// SubjectAccessReview checks whether or not a user or group can perform an action.
type KubeSubjectAccessReviewProps struct {
	// Spec holds information about the request being evaluated.
	Spec *SubjectAccessReviewSpec `field:"required" json:"spec" yaml:"spec"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

