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

// FakeEc2ClientVPNNetworkAssociations implements Ec2ClientVPNNetworkAssociationInterface
type FakeEc2ClientVPNNetworkAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ec2clientvpnnetworkassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ec2clientvpnnetworkassociations"}

var ec2clientvpnnetworkassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Ec2ClientVPNNetworkAssociation"}

// Get takes name of the ec2ClientVPNNetworkAssociation, and returns the corresponding ec2ClientVPNNetworkAssociation object, and an error if there is any.
func (c *FakeEc2ClientVPNNetworkAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ec2ClientVPNNetworkAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ec2clientvpnnetworkassociationsResource, c.ns, name), &v1alpha1.Ec2ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2ClientVPNNetworkAssociation), err
}

// List takes label and field selectors, and returns the list of Ec2ClientVPNNetworkAssociations that match those selectors.
func (c *FakeEc2ClientVPNNetworkAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Ec2ClientVPNNetworkAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ec2clientvpnnetworkassociationsResource, ec2clientvpnnetworkassociationsKind, c.ns, opts), &v1alpha1.Ec2ClientVPNNetworkAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Ec2ClientVPNNetworkAssociationList{ListMeta: obj.(*v1alpha1.Ec2ClientVPNNetworkAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.Ec2ClientVPNNetworkAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ec2ClientVPNNetworkAssociations.
func (c *FakeEc2ClientVPNNetworkAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ec2clientvpnnetworkassociationsResource, c.ns, opts))

}

// Create takes the representation of a ec2ClientVPNNetworkAssociation and creates it.  Returns the server's representation of the ec2ClientVPNNetworkAssociation, and an error, if there is any.
func (c *FakeEc2ClientVPNNetworkAssociations) Create(ctx context.Context, ec2ClientVPNNetworkAssociation *v1alpha1.Ec2ClientVPNNetworkAssociation, opts v1.CreateOptions) (result *v1alpha1.Ec2ClientVPNNetworkAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ec2clientvpnnetworkassociationsResource, c.ns, ec2ClientVPNNetworkAssociation), &v1alpha1.Ec2ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2ClientVPNNetworkAssociation), err
}

// Update takes the representation of a ec2ClientVPNNetworkAssociation and updates it. Returns the server's representation of the ec2ClientVPNNetworkAssociation, and an error, if there is any.
func (c *FakeEc2ClientVPNNetworkAssociations) Update(ctx context.Context, ec2ClientVPNNetworkAssociation *v1alpha1.Ec2ClientVPNNetworkAssociation, opts v1.UpdateOptions) (result *v1alpha1.Ec2ClientVPNNetworkAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ec2clientvpnnetworkassociationsResource, c.ns, ec2ClientVPNNetworkAssociation), &v1alpha1.Ec2ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2ClientVPNNetworkAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEc2ClientVPNNetworkAssociations) UpdateStatus(ctx context.Context, ec2ClientVPNNetworkAssociation *v1alpha1.Ec2ClientVPNNetworkAssociation, opts v1.UpdateOptions) (*v1alpha1.Ec2ClientVPNNetworkAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ec2clientvpnnetworkassociationsResource, "status", c.ns, ec2ClientVPNNetworkAssociation), &v1alpha1.Ec2ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2ClientVPNNetworkAssociation), err
}

// Delete takes name of the ec2ClientVPNNetworkAssociation and deletes it. Returns an error if one occurs.
func (c *FakeEc2ClientVPNNetworkAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ec2clientvpnnetworkassociationsResource, c.ns, name), &v1alpha1.Ec2ClientVPNNetworkAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEc2ClientVPNNetworkAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ec2clientvpnnetworkassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.Ec2ClientVPNNetworkAssociationList{})
	return err
}

// Patch applies the patch and returns the patched ec2ClientVPNNetworkAssociation.
func (c *FakeEc2ClientVPNNetworkAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ec2ClientVPNNetworkAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ec2clientvpnnetworkassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ec2ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2ClientVPNNetworkAssociation), err
}
