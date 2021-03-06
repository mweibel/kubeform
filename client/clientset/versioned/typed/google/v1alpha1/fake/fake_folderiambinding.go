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

// FakeFolderIamBindings implements FolderIamBindingInterface
type FakeFolderIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var folderiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "folderiambindings"}

var folderiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "FolderIamBinding"}

// Get takes name of the folderIamBinding, and returns the corresponding folderIamBinding object, and an error if there is any.
func (c *FakeFolderIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FolderIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(folderiambindingsResource, c.ns, name), &v1alpha1.FolderIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamBinding), err
}

// List takes label and field selectors, and returns the list of FolderIamBindings that match those selectors.
func (c *FakeFolderIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FolderIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(folderiambindingsResource, folderiambindingsKind, c.ns, opts), &v1alpha1.FolderIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FolderIamBindingList{ListMeta: obj.(*v1alpha1.FolderIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.FolderIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested folderIamBindings.
func (c *FakeFolderIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(folderiambindingsResource, c.ns, opts))

}

// Create takes the representation of a folderIamBinding and creates it.  Returns the server's representation of the folderIamBinding, and an error, if there is any.
func (c *FakeFolderIamBindings) Create(ctx context.Context, folderIamBinding *v1alpha1.FolderIamBinding, opts v1.CreateOptions) (result *v1alpha1.FolderIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(folderiambindingsResource, c.ns, folderIamBinding), &v1alpha1.FolderIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamBinding), err
}

// Update takes the representation of a folderIamBinding and updates it. Returns the server's representation of the folderIamBinding, and an error, if there is any.
func (c *FakeFolderIamBindings) Update(ctx context.Context, folderIamBinding *v1alpha1.FolderIamBinding, opts v1.UpdateOptions) (result *v1alpha1.FolderIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(folderiambindingsResource, c.ns, folderIamBinding), &v1alpha1.FolderIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFolderIamBindings) UpdateStatus(ctx context.Context, folderIamBinding *v1alpha1.FolderIamBinding, opts v1.UpdateOptions) (*v1alpha1.FolderIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(folderiambindingsResource, "status", c.ns, folderIamBinding), &v1alpha1.FolderIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamBinding), err
}

// Delete takes name of the folderIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeFolderIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(folderiambindingsResource, c.ns, name), &v1alpha1.FolderIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFolderIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(folderiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FolderIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched folderIamBinding.
func (c *FakeFolderIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FolderIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(folderiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FolderIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamBinding), err
}
