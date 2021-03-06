package fake

import (
	v1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeSubjectAccessReviews implements SubjectAccessReviewInterface
type FakeSubjectAccessReviews struct {
	Fake *FakeAuthorizationV1
}

var subjectaccessreviewsResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "v1", Resource: "subjectaccessreviews"}

var subjectaccessreviewsKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "v1", Kind: "SubjectAccessReview"}

// Create takes the representation of a subjectAccessReview and creates it.  Returns the server's representation of the subjectAccessReviewResponse, and an error, if there is any.
func (c *FakeSubjectAccessReviews) Create(subjectAccessReview *v1.SubjectAccessReview) (result *v1.SubjectAccessReviewResponse, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(subjectaccessreviewsResource, subjectAccessReview), &v1.SubjectAccessReviewResponse{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.SubjectAccessReviewResponse), err
}
