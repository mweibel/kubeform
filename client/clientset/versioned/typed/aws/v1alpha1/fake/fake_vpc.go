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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVpcs implements VpcInterface
type FakeVpcs struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpcsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpcs"}

var vpcsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Vpc"}

// Get takes name of the vpc, and returns the corresponding vpc object, and an error if there is any.
func (c *FakeVpcs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Vpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcsResource, c.ns, name), &v1alpha1.Vpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vpc), err
}

// List takes label and field selectors, and returns the list of Vpcs that match those selectors.
func (c *FakeVpcs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpcList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcsResource, vpcsKind, c.ns, opts), &v1alpha1.VpcList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcList{ListMeta: obj.(*v1alpha1.VpcList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcs.
func (c *FakeVpcs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcsResource, c.ns, opts))

}

// Create takes the representation of a vpc and creates it.  Returns the server's representation of the vpc, and an error, if there is any.
func (c *FakeVpcs) Create(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.CreateOptions) (result *v1alpha1.Vpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcsResource, c.ns, vpc), &v1alpha1.Vpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vpc), err
}

// Update takes the representation of a vpc and updates it. Returns the server's representation of the vpc, and an error, if there is any.
func (c *FakeVpcs) Update(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.UpdateOptions) (result *v1alpha1.Vpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcsResource, c.ns, vpc), &v1alpha1.Vpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vpc), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcs) UpdateStatus(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.UpdateOptions) (*v1alpha1.Vpc, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcsResource, "status", c.ns, vpc), &v1alpha1.Vpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vpc), err
}

// Delete takes name of the vpc and deletes it. Returns an error if one occurs.
func (c *FakeVpcs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcsResource, c.ns, name), &v1alpha1.Vpc{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcList{})
	return err
}

// Patch applies the patch and returns the patched vpc.
func (c *FakeVpcs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vpc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Vpc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vpc), err
}
