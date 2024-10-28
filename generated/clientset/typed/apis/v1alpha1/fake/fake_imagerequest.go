/*
  Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/srl-labs/clabernetes/apis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImageRequests implements ImageRequestInterface
type FakeImageRequests struct {
	Fake *FakeClabernetesV1alpha1
	ns   string
}

var imagerequestsResource = v1alpha1.SchemeGroupVersion.WithResource("imagerequests")

var imagerequestsKind = v1alpha1.SchemeGroupVersion.WithKind("ImageRequest")

// Get takes name of the imageRequest, and returns the corresponding imageRequest object, and an error if there is any.
func (c *FakeImageRequests) Get(
	ctx context.Context,
	name string,
	options v1.GetOptions,
) (result *v1alpha1.ImageRequest, err error) {
	emptyResult := &v1alpha1.ImageRequest{}
	obj, err := c.Fake.
		Invokes(
			testing.NewGetActionWithOptions(imagerequestsResource, c.ns, name, options),
			emptyResult,
		)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ImageRequest), err
}

// List takes label and field selectors, and returns the list of ImageRequests that match those selectors.
func (c *FakeImageRequests) List(
	ctx context.Context,
	opts v1.ListOptions,
) (result *v1alpha1.ImageRequestList, err error) {
	emptyResult := &v1alpha1.ImageRequestList{}
	obj, err := c.Fake.
		Invokes(
			testing.NewListActionWithOptions(imagerequestsResource, imagerequestsKind, c.ns, opts),
			emptyResult,
		)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImageRequestList{ListMeta: obj.(*v1alpha1.ImageRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImageRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageRequests.
func (c *FakeImageRequests) Watch(
	ctx context.Context,
	opts v1.ListOptions,
) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(imagerequestsResource, c.ns, opts))

}

// Create takes the representation of a imageRequest and creates it.  Returns the server's representation of the imageRequest, and an error, if there is any.
func (c *FakeImageRequests) Create(
	ctx context.Context,
	imageRequest *v1alpha1.ImageRequest,
	opts v1.CreateOptions,
) (result *v1alpha1.ImageRequest, err error) {
	emptyResult := &v1alpha1.ImageRequest{}
	obj, err := c.Fake.
		Invokes(
			testing.NewCreateActionWithOptions(imagerequestsResource, c.ns, imageRequest, opts),
			emptyResult,
		)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ImageRequest), err
}

// Update takes the representation of a imageRequest and updates it. Returns the server's representation of the imageRequest, and an error, if there is any.
func (c *FakeImageRequests) Update(
	ctx context.Context,
	imageRequest *v1alpha1.ImageRequest,
	opts v1.UpdateOptions,
) (result *v1alpha1.ImageRequest, err error) {
	emptyResult := &v1alpha1.ImageRequest{}
	obj, err := c.Fake.
		Invokes(
			testing.NewUpdateActionWithOptions(imagerequestsResource, c.ns, imageRequest, opts),
			emptyResult,
		)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ImageRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageRequests) UpdateStatus(
	ctx context.Context,
	imageRequest *v1alpha1.ImageRequest,
	opts v1.UpdateOptions,
) (result *v1alpha1.ImageRequest, err error) {
	emptyResult := &v1alpha1.ImageRequest{}
	obj, err := c.Fake.
		Invokes(
			testing.NewUpdateSubresourceActionWithOptions(
				imagerequestsResource,
				"status",
				c.ns,
				imageRequest,
				opts,
			),
			emptyResult,
		)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ImageRequest), err
}

// Delete takes name of the imageRequest and deletes it. Returns an error if one occurs.
func (c *FakeImageRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(
			testing.NewDeleteActionWithOptions(imagerequestsResource, c.ns, name, opts),
			&v1alpha1.ImageRequest{},
		)

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageRequests) DeleteCollection(
	ctx context.Context,
	opts v1.DeleteOptions,
	listOpts v1.ListOptions,
) error {
	action := testing.NewDeleteCollectionActionWithOptions(
		imagerequestsResource,
		c.ns,
		opts,
		listOpts,
	)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImageRequestList{})
	return err
}

// Patch applies the patch and returns the patched imageRequest.
func (c *FakeImageRequests) Patch(
	ctx context.Context,
	name string,
	pt types.PatchType,
	data []byte,
	opts v1.PatchOptions,
	subresources ...string,
) (result *v1alpha1.ImageRequest, err error) {
	emptyResult := &v1alpha1.ImageRequest{}
	obj, err := c.Fake.
		Invokes(
			testing.NewPatchSubresourceActionWithOptions(
				imagerequestsResource,
				c.ns,
				name,
				pt,
				data,
				opts,
				subresources...),
			emptyResult,
		)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ImageRequest), err
}
