/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package unversioned

import (
	"k8s.io/kubernetes/pkg/apis/authentication"
)

// TokenReviews has methods to work with TokenReview resources in a namespace
type TokenReviewsInterface interface {
	TokenReviews() TokenReviewInterface
}

// TokenReviewInterface has methods to work with TokenReview resources.
type TokenReviewInterface interface {
	Create(tokenReview *authentication.TokenReview) (*authentication.TokenReview, error)
}

// tokenReviews implements TokenReviewsNamespacer interface
type tokenReviews struct {
	client *AuthenticationClient
}

// newTokenReviews returns a tokenReviews
func newTokenReviews(c *AuthenticationClient) *tokenReviews {
	return &tokenReviews{
		client: c,
	}
}

func (c *tokenReviews) Create(obj *authentication.TokenReview) (result *authentication.TokenReview, err error) {
	result = &authentication.TokenReview{}
	err = c.client.Post().Resource("tokenreviews").Body(obj).Do().Into(result)
	return
}
