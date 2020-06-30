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

// FakeApiManagementProperties implements ApiManagementPropertyInterface
type FakeApiManagementProperties struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementpropertiesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementproperties"}

var apimanagementpropertiesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementProperty"}

// Get takes name of the apiManagementProperty, and returns the corresponding apiManagementProperty object, and an error if there is any.
func (c *FakeApiManagementProperties) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementProperty, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementpropertiesResource, c.ns, name), &v1alpha1.ApiManagementProperty{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProperty), err
}

// List takes label and field selectors, and returns the list of ApiManagementProperties that match those selectors.
func (c *FakeApiManagementProperties) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementPropertyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementpropertiesResource, apimanagementpropertiesKind, c.ns, opts), &v1alpha1.ApiManagementPropertyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementPropertyList{ListMeta: obj.(*v1alpha1.ApiManagementPropertyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementPropertyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementProperties.
func (c *FakeApiManagementProperties) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementpropertiesResource, c.ns, opts))

}

// Create takes the representation of a apiManagementProperty and creates it.  Returns the server's representation of the apiManagementProperty, and an error, if there is any.
func (c *FakeApiManagementProperties) Create(ctx context.Context, apiManagementProperty *v1alpha1.ApiManagementProperty, opts v1.CreateOptions) (result *v1alpha1.ApiManagementProperty, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementpropertiesResource, c.ns, apiManagementProperty), &v1alpha1.ApiManagementProperty{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProperty), err
}

// Update takes the representation of a apiManagementProperty and updates it. Returns the server's representation of the apiManagementProperty, and an error, if there is any.
func (c *FakeApiManagementProperties) Update(ctx context.Context, apiManagementProperty *v1alpha1.ApiManagementProperty, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementProperty, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementpropertiesResource, c.ns, apiManagementProperty), &v1alpha1.ApiManagementProperty{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProperty), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementProperties) UpdateStatus(ctx context.Context, apiManagementProperty *v1alpha1.ApiManagementProperty, opts v1.UpdateOptions) (*v1alpha1.ApiManagementProperty, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementpropertiesResource, "status", c.ns, apiManagementProperty), &v1alpha1.ApiManagementProperty{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProperty), err
}

// Delete takes name of the apiManagementProperty and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementProperties) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementpropertiesResource, c.ns, name), &v1alpha1.ApiManagementProperty{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementProperties) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementpropertiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementPropertyList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementProperty.
func (c *FakeApiManagementProperties) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementProperty, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementpropertiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementProperty{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProperty), err
}
