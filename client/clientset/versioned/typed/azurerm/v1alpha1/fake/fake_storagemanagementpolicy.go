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

// FakeStorageManagementPolicies implements StorageManagementPolicyInterface
type FakeStorageManagementPolicies struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var storagemanagementpoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "storagemanagementpolicies"}

var storagemanagementpoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StorageManagementPolicy"}

// Get takes name of the storageManagementPolicy, and returns the corresponding storageManagementPolicy object, and an error if there is any.
func (c *FakeStorageManagementPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageManagementPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagemanagementpoliciesResource, c.ns, name), &v1alpha1.StorageManagementPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageManagementPolicy), err
}

// List takes label and field selectors, and returns the list of StorageManagementPolicies that match those selectors.
func (c *FakeStorageManagementPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageManagementPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagemanagementpoliciesResource, storagemanagementpoliciesKind, c.ns, opts), &v1alpha1.StorageManagementPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageManagementPolicyList{ListMeta: obj.(*v1alpha1.StorageManagementPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageManagementPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageManagementPolicies.
func (c *FakeStorageManagementPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagemanagementpoliciesResource, c.ns, opts))

}

// Create takes the representation of a storageManagementPolicy and creates it.  Returns the server's representation of the storageManagementPolicy, and an error, if there is any.
func (c *FakeStorageManagementPolicies) Create(ctx context.Context, storageManagementPolicy *v1alpha1.StorageManagementPolicy, opts v1.CreateOptions) (result *v1alpha1.StorageManagementPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagemanagementpoliciesResource, c.ns, storageManagementPolicy), &v1alpha1.StorageManagementPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageManagementPolicy), err
}

// Update takes the representation of a storageManagementPolicy and updates it. Returns the server's representation of the storageManagementPolicy, and an error, if there is any.
func (c *FakeStorageManagementPolicies) Update(ctx context.Context, storageManagementPolicy *v1alpha1.StorageManagementPolicy, opts v1.UpdateOptions) (result *v1alpha1.StorageManagementPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagemanagementpoliciesResource, c.ns, storageManagementPolicy), &v1alpha1.StorageManagementPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageManagementPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageManagementPolicies) UpdateStatus(ctx context.Context, storageManagementPolicy *v1alpha1.StorageManagementPolicy, opts v1.UpdateOptions) (*v1alpha1.StorageManagementPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagemanagementpoliciesResource, "status", c.ns, storageManagementPolicy), &v1alpha1.StorageManagementPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageManagementPolicy), err
}

// Delete takes name of the storageManagementPolicy and deletes it. Returns an error if one occurs.
func (c *FakeStorageManagementPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagemanagementpoliciesResource, c.ns, name), &v1alpha1.StorageManagementPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageManagementPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagemanagementpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageManagementPolicyList{})
	return err
}

// Patch applies the patch and returns the patched storageManagementPolicy.
func (c *FakeStorageManagementPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageManagementPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagemanagementpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageManagementPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageManagementPolicy), err
}