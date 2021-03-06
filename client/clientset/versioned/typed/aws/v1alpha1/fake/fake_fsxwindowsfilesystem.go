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

// FakeFsxWindowsFileSystems implements FsxWindowsFileSystemInterface
type FakeFsxWindowsFileSystems struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var fsxwindowsfilesystemsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "fsxwindowsfilesystems"}

var fsxwindowsfilesystemsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "FsxWindowsFileSystem"}

// Get takes name of the fsxWindowsFileSystem, and returns the corresponding fsxWindowsFileSystem object, and an error if there is any.
func (c *FakeFsxWindowsFileSystems) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FsxWindowsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fsxwindowsfilesystemsResource, c.ns, name), &v1alpha1.FsxWindowsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FsxWindowsFileSystem), err
}

// List takes label and field selectors, and returns the list of FsxWindowsFileSystems that match those selectors.
func (c *FakeFsxWindowsFileSystems) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FsxWindowsFileSystemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fsxwindowsfilesystemsResource, fsxwindowsfilesystemsKind, c.ns, opts), &v1alpha1.FsxWindowsFileSystemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FsxWindowsFileSystemList{ListMeta: obj.(*v1alpha1.FsxWindowsFileSystemList).ListMeta}
	for _, item := range obj.(*v1alpha1.FsxWindowsFileSystemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fsxWindowsFileSystems.
func (c *FakeFsxWindowsFileSystems) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fsxwindowsfilesystemsResource, c.ns, opts))

}

// Create takes the representation of a fsxWindowsFileSystem and creates it.  Returns the server's representation of the fsxWindowsFileSystem, and an error, if there is any.
func (c *FakeFsxWindowsFileSystems) Create(ctx context.Context, fsxWindowsFileSystem *v1alpha1.FsxWindowsFileSystem, opts v1.CreateOptions) (result *v1alpha1.FsxWindowsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fsxwindowsfilesystemsResource, c.ns, fsxWindowsFileSystem), &v1alpha1.FsxWindowsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FsxWindowsFileSystem), err
}

// Update takes the representation of a fsxWindowsFileSystem and updates it. Returns the server's representation of the fsxWindowsFileSystem, and an error, if there is any.
func (c *FakeFsxWindowsFileSystems) Update(ctx context.Context, fsxWindowsFileSystem *v1alpha1.FsxWindowsFileSystem, opts v1.UpdateOptions) (result *v1alpha1.FsxWindowsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fsxwindowsfilesystemsResource, c.ns, fsxWindowsFileSystem), &v1alpha1.FsxWindowsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FsxWindowsFileSystem), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFsxWindowsFileSystems) UpdateStatus(ctx context.Context, fsxWindowsFileSystem *v1alpha1.FsxWindowsFileSystem, opts v1.UpdateOptions) (*v1alpha1.FsxWindowsFileSystem, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fsxwindowsfilesystemsResource, "status", c.ns, fsxWindowsFileSystem), &v1alpha1.FsxWindowsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FsxWindowsFileSystem), err
}

// Delete takes name of the fsxWindowsFileSystem and deletes it. Returns an error if one occurs.
func (c *FakeFsxWindowsFileSystems) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fsxwindowsfilesystemsResource, c.ns, name), &v1alpha1.FsxWindowsFileSystem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFsxWindowsFileSystems) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fsxwindowsfilesystemsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FsxWindowsFileSystemList{})
	return err
}

// Patch applies the patch and returns the patched fsxWindowsFileSystem.
func (c *FakeFsxWindowsFileSystems) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FsxWindowsFileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fsxwindowsfilesystemsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FsxWindowsFileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FsxWindowsFileSystem), err
}
