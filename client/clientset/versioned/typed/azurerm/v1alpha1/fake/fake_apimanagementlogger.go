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

// FakeApiManagementLoggers implements ApiManagementLoggerInterface
type FakeApiManagementLoggers struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementloggersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementloggers"}

var apimanagementloggersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementLogger"}

// Get takes name of the apiManagementLogger, and returns the corresponding apiManagementLogger object, and an error if there is any.
func (c *FakeApiManagementLoggers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementLogger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementloggersResource, c.ns, name), &v1alpha1.ApiManagementLogger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementLogger), err
}

// List takes label and field selectors, and returns the list of ApiManagementLoggers that match those selectors.
func (c *FakeApiManagementLoggers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementLoggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementloggersResource, apimanagementloggersKind, c.ns, opts), &v1alpha1.ApiManagementLoggerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementLoggerList{ListMeta: obj.(*v1alpha1.ApiManagementLoggerList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementLoggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementLoggers.
func (c *FakeApiManagementLoggers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementloggersResource, c.ns, opts))

}

// Create takes the representation of a apiManagementLogger and creates it.  Returns the server's representation of the apiManagementLogger, and an error, if there is any.
func (c *FakeApiManagementLoggers) Create(ctx context.Context, apiManagementLogger *v1alpha1.ApiManagementLogger, opts v1.CreateOptions) (result *v1alpha1.ApiManagementLogger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementloggersResource, c.ns, apiManagementLogger), &v1alpha1.ApiManagementLogger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementLogger), err
}

// Update takes the representation of a apiManagementLogger and updates it. Returns the server's representation of the apiManagementLogger, and an error, if there is any.
func (c *FakeApiManagementLoggers) Update(ctx context.Context, apiManagementLogger *v1alpha1.ApiManagementLogger, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementLogger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementloggersResource, c.ns, apiManagementLogger), &v1alpha1.ApiManagementLogger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementLogger), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementLoggers) UpdateStatus(ctx context.Context, apiManagementLogger *v1alpha1.ApiManagementLogger, opts v1.UpdateOptions) (*v1alpha1.ApiManagementLogger, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementloggersResource, "status", c.ns, apiManagementLogger), &v1alpha1.ApiManagementLogger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementLogger), err
}

// Delete takes name of the apiManagementLogger and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementLoggers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementloggersResource, c.ns, name), &v1alpha1.ApiManagementLogger{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementLoggers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementloggersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementLoggerList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementLogger.
func (c *FakeApiManagementLoggers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementLogger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementloggersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementLogger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementLogger), err
}
