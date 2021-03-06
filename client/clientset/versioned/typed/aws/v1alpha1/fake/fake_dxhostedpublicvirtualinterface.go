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

// FakeDxHostedPublicVirtualInterfaces implements DxHostedPublicVirtualInterfaceInterface
type FakeDxHostedPublicVirtualInterfaces struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dxhostedpublicvirtualinterfacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dxhostedpublicvirtualinterfaces"}

var dxhostedpublicvirtualinterfacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DxHostedPublicVirtualInterface"}

// Get takes name of the dxHostedPublicVirtualInterface, and returns the corresponding dxHostedPublicVirtualInterface object, and an error if there is any.
func (c *FakeDxHostedPublicVirtualInterfaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dxhostedpublicvirtualinterfacesResource, c.ns, name), &v1alpha1.DxHostedPublicVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPublicVirtualInterface), err
}

// List takes label and field selectors, and returns the list of DxHostedPublicVirtualInterfaces that match those selectors.
func (c *FakeDxHostedPublicVirtualInterfaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DxHostedPublicVirtualInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dxhostedpublicvirtualinterfacesResource, dxhostedpublicvirtualinterfacesKind, c.ns, opts), &v1alpha1.DxHostedPublicVirtualInterfaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DxHostedPublicVirtualInterfaceList{ListMeta: obj.(*v1alpha1.DxHostedPublicVirtualInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DxHostedPublicVirtualInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dxHostedPublicVirtualInterfaces.
func (c *FakeDxHostedPublicVirtualInterfaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dxhostedpublicvirtualinterfacesResource, c.ns, opts))

}

// Create takes the representation of a dxHostedPublicVirtualInterface and creates it.  Returns the server's representation of the dxHostedPublicVirtualInterface, and an error, if there is any.
func (c *FakeDxHostedPublicVirtualInterfaces) Create(ctx context.Context, dxHostedPublicVirtualInterface *v1alpha1.DxHostedPublicVirtualInterface, opts v1.CreateOptions) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dxhostedpublicvirtualinterfacesResource, c.ns, dxHostedPublicVirtualInterface), &v1alpha1.DxHostedPublicVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPublicVirtualInterface), err
}

// Update takes the representation of a dxHostedPublicVirtualInterface and updates it. Returns the server's representation of the dxHostedPublicVirtualInterface, and an error, if there is any.
func (c *FakeDxHostedPublicVirtualInterfaces) Update(ctx context.Context, dxHostedPublicVirtualInterface *v1alpha1.DxHostedPublicVirtualInterface, opts v1.UpdateOptions) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dxhostedpublicvirtualinterfacesResource, c.ns, dxHostedPublicVirtualInterface), &v1alpha1.DxHostedPublicVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPublicVirtualInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDxHostedPublicVirtualInterfaces) UpdateStatus(ctx context.Context, dxHostedPublicVirtualInterface *v1alpha1.DxHostedPublicVirtualInterface, opts v1.UpdateOptions) (*v1alpha1.DxHostedPublicVirtualInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dxhostedpublicvirtualinterfacesResource, "status", c.ns, dxHostedPublicVirtualInterface), &v1alpha1.DxHostedPublicVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPublicVirtualInterface), err
}

// Delete takes name of the dxHostedPublicVirtualInterface and deletes it. Returns an error if one occurs.
func (c *FakeDxHostedPublicVirtualInterfaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dxhostedpublicvirtualinterfacesResource, c.ns, name), &v1alpha1.DxHostedPublicVirtualInterface{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDxHostedPublicVirtualInterfaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dxhostedpublicvirtualinterfacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DxHostedPublicVirtualInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched dxHostedPublicVirtualInterface.
func (c *FakeDxHostedPublicVirtualInterfaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DxHostedPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dxhostedpublicvirtualinterfacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DxHostedPublicVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPublicVirtualInterface), err
}
