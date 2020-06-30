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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSecurityCenterContacts implements SecurityCenterContactInterface
type FakeSecurityCenterContacts struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var securitycentercontactsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "securitycentercontacts"}

var securitycentercontactsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SecurityCenterContact"}

// Get takes name of the securityCenterContact, and returns the corresponding securityCenterContact object, and an error if there is any.
func (c *FakeSecurityCenterContacts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecurityCenterContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(securitycentercontactsResource, c.ns, name), &v1alpha1.SecurityCenterContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterContact), err
}

// List takes label and field selectors, and returns the list of SecurityCenterContacts that match those selectors.
func (c *FakeSecurityCenterContacts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecurityCenterContactList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(securitycentercontactsResource, securitycentercontactsKind, c.ns, opts), &v1alpha1.SecurityCenterContactList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecurityCenterContactList{ListMeta: obj.(*v1alpha1.SecurityCenterContactList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecurityCenterContactList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested securityCenterContacts.
func (c *FakeSecurityCenterContacts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(securitycentercontactsResource, c.ns, opts))

}

// Create takes the representation of a securityCenterContact and creates it.  Returns the server's representation of the securityCenterContact, and an error, if there is any.
func (c *FakeSecurityCenterContacts) Create(ctx context.Context, securityCenterContact *v1alpha1.SecurityCenterContact, opts v1.CreateOptions) (result *v1alpha1.SecurityCenterContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(securitycentercontactsResource, c.ns, securityCenterContact), &v1alpha1.SecurityCenterContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterContact), err
}

// Update takes the representation of a securityCenterContact and updates it. Returns the server's representation of the securityCenterContact, and an error, if there is any.
func (c *FakeSecurityCenterContacts) Update(ctx context.Context, securityCenterContact *v1alpha1.SecurityCenterContact, opts v1.UpdateOptions) (result *v1alpha1.SecurityCenterContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(securitycentercontactsResource, c.ns, securityCenterContact), &v1alpha1.SecurityCenterContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterContact), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecurityCenterContacts) UpdateStatus(ctx context.Context, securityCenterContact *v1alpha1.SecurityCenterContact, opts v1.UpdateOptions) (*v1alpha1.SecurityCenterContact, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(securitycentercontactsResource, "status", c.ns, securityCenterContact), &v1alpha1.SecurityCenterContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterContact), err
}

// Delete takes name of the securityCenterContact and deletes it. Returns an error if one occurs.
func (c *FakeSecurityCenterContacts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(securitycentercontactsResource, c.ns, name), &v1alpha1.SecurityCenterContact{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecurityCenterContacts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(securitycentercontactsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecurityCenterContactList{})
	return err
}

// Patch applies the patch and returns the patched securityCenterContact.
func (c *FakeSecurityCenterContacts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecurityCenterContact, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(securitycentercontactsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SecurityCenterContact{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterContact), err
}
