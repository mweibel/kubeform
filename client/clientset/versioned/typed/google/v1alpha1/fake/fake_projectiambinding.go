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

// FakeProjectIamBindings implements ProjectIamBindingInterface
type FakeProjectIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var projectiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "projectiambindings"}

var projectiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ProjectIamBinding"}

// Get takes name of the projectIamBinding, and returns the corresponding projectIamBinding object, and an error if there is any.
func (c *FakeProjectIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProjectIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(projectiambindingsResource, c.ns, name), &v1alpha1.ProjectIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamBinding), err
}

// List takes label and field selectors, and returns the list of ProjectIamBindings that match those selectors.
func (c *FakeProjectIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProjectIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(projectiambindingsResource, projectiambindingsKind, c.ns, opts), &v1alpha1.ProjectIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProjectIamBindingList{ListMeta: obj.(*v1alpha1.ProjectIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProjectIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectIamBindings.
func (c *FakeProjectIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(projectiambindingsResource, c.ns, opts))

}

// Create takes the representation of a projectIamBinding and creates it.  Returns the server's representation of the projectIamBinding, and an error, if there is any.
func (c *FakeProjectIamBindings) Create(ctx context.Context, projectIamBinding *v1alpha1.ProjectIamBinding, opts v1.CreateOptions) (result *v1alpha1.ProjectIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(projectiambindingsResource, c.ns, projectIamBinding), &v1alpha1.ProjectIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamBinding), err
}

// Update takes the representation of a projectIamBinding and updates it. Returns the server's representation of the projectIamBinding, and an error, if there is any.
func (c *FakeProjectIamBindings) Update(ctx context.Context, projectIamBinding *v1alpha1.ProjectIamBinding, opts v1.UpdateOptions) (result *v1alpha1.ProjectIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(projectiambindingsResource, c.ns, projectIamBinding), &v1alpha1.ProjectIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjectIamBindings) UpdateStatus(ctx context.Context, projectIamBinding *v1alpha1.ProjectIamBinding, opts v1.UpdateOptions) (*v1alpha1.ProjectIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(projectiambindingsResource, "status", c.ns, projectIamBinding), &v1alpha1.ProjectIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamBinding), err
}

// Delete takes name of the projectIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeProjectIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(projectiambindingsResource, c.ns, name), &v1alpha1.ProjectIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(projectiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProjectIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched projectIamBinding.
func (c *FakeProjectIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProjectIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(projectiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProjectIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamBinding), err
}
