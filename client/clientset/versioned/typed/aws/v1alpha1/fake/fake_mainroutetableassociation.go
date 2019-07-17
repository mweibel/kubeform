/*
Copyright 2019 The Kubeform Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeMainRouteTableAssociations implements MainRouteTableAssociationInterface
type FakeMainRouteTableAssociations struct {
	Fake *FakeAwsV1alpha1
}

var mainroutetableassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "mainroutetableassociations"}

var mainroutetableassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "MainRouteTableAssociation"}

// Get takes name of the mainRouteTableAssociation, and returns the corresponding mainRouteTableAssociation object, and an error if there is any.
func (c *FakeMainRouteTableAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.MainRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(mainroutetableassociationsResource, name), &v1alpha1.MainRouteTableAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MainRouteTableAssociation), err
}

// List takes label and field selectors, and returns the list of MainRouteTableAssociations that match those selectors.
func (c *FakeMainRouteTableAssociations) List(opts v1.ListOptions) (result *v1alpha1.MainRouteTableAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(mainroutetableassociationsResource, mainroutetableassociationsKind, opts), &v1alpha1.MainRouteTableAssociationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MainRouteTableAssociationList{ListMeta: obj.(*v1alpha1.MainRouteTableAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.MainRouteTableAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mainRouteTableAssociations.
func (c *FakeMainRouteTableAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(mainroutetableassociationsResource, opts))
}

// Create takes the representation of a mainRouteTableAssociation and creates it.  Returns the server's representation of the mainRouteTableAssociation, and an error, if there is any.
func (c *FakeMainRouteTableAssociations) Create(mainRouteTableAssociation *v1alpha1.MainRouteTableAssociation) (result *v1alpha1.MainRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(mainroutetableassociationsResource, mainRouteTableAssociation), &v1alpha1.MainRouteTableAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MainRouteTableAssociation), err
}

// Update takes the representation of a mainRouteTableAssociation and updates it. Returns the server's representation of the mainRouteTableAssociation, and an error, if there is any.
func (c *FakeMainRouteTableAssociations) Update(mainRouteTableAssociation *v1alpha1.MainRouteTableAssociation) (result *v1alpha1.MainRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(mainroutetableassociationsResource, mainRouteTableAssociation), &v1alpha1.MainRouteTableAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MainRouteTableAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMainRouteTableAssociations) UpdateStatus(mainRouteTableAssociation *v1alpha1.MainRouteTableAssociation) (*v1alpha1.MainRouteTableAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(mainroutetableassociationsResource, "status", mainRouteTableAssociation), &v1alpha1.MainRouteTableAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MainRouteTableAssociation), err
}

// Delete takes name of the mainRouteTableAssociation and deletes it. Returns an error if one occurs.
func (c *FakeMainRouteTableAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(mainroutetableassociationsResource, name), &v1alpha1.MainRouteTableAssociation{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMainRouteTableAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(mainroutetableassociationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MainRouteTableAssociationList{})
	return err
}

// Patch applies the patch and returns the patched mainRouteTableAssociation.
func (c *FakeMainRouteTableAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MainRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(mainroutetableassociationsResource, name, pt, data, subresources...), &v1alpha1.MainRouteTableAssociation{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MainRouteTableAssociation), err
}