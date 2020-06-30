/*
Copyright The Kubeform Authors.

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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeRouters implements ComputeRouterInterface
type FakeComputeRouters struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computeroutersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computerouters"}

var computeroutersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeRouter"}

// Get takes name of the computeRouter, and returns the corresponding computeRouter object, and an error if there is any.
func (c *FakeComputeRouters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeroutersResource, c.ns, name), &v1alpha1.ComputeRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouter), err
}

// List takes label and field selectors, and returns the list of ComputeRouters that match those selectors.
func (c *FakeComputeRouters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeRouterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeroutersResource, computeroutersKind, c.ns, opts), &v1alpha1.ComputeRouterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeRouterList{ListMeta: obj.(*v1alpha1.ComputeRouterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeRouterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeRouters.
func (c *FakeComputeRouters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeroutersResource, c.ns, opts))

}

// Create takes the representation of a computeRouter and creates it.  Returns the server's representation of the computeRouter, and an error, if there is any.
func (c *FakeComputeRouters) Create(ctx context.Context, computeRouter *v1alpha1.ComputeRouter, opts v1.CreateOptions) (result *v1alpha1.ComputeRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeroutersResource, c.ns, computeRouter), &v1alpha1.ComputeRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouter), err
}

// Update takes the representation of a computeRouter and updates it. Returns the server's representation of the computeRouter, and an error, if there is any.
func (c *FakeComputeRouters) Update(ctx context.Context, computeRouter *v1alpha1.ComputeRouter, opts v1.UpdateOptions) (result *v1alpha1.ComputeRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeroutersResource, c.ns, computeRouter), &v1alpha1.ComputeRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeRouters) UpdateStatus(ctx context.Context, computeRouter *v1alpha1.ComputeRouter, opts v1.UpdateOptions) (*v1alpha1.ComputeRouter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeroutersResource, "status", c.ns, computeRouter), &v1alpha1.ComputeRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouter), err
}

// Delete takes name of the computeRouter and deletes it. Returns an error if one occurs.
func (c *FakeComputeRouters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeroutersResource, c.ns, name), &v1alpha1.ComputeRouter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeRouters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeroutersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeRouterList{})
	return err
}

// Patch applies the patch and returns the patched computeRouter.
func (c *FakeComputeRouters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeroutersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouter), err
}
