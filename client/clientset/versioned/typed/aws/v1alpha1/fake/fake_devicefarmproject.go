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

// FakeDevicefarmProjects implements DevicefarmProjectInterface
type FakeDevicefarmProjects struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var devicefarmprojectsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "devicefarmprojects"}

var devicefarmprojectsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DevicefarmProject"}

// Get takes name of the devicefarmProject, and returns the corresponding devicefarmProject object, and an error if there is any.
func (c *FakeDevicefarmProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DevicefarmProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devicefarmprojectsResource, c.ns, name), &v1alpha1.DevicefarmProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevicefarmProject), err
}

// List takes label and field selectors, and returns the list of DevicefarmProjects that match those selectors.
func (c *FakeDevicefarmProjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DevicefarmProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devicefarmprojectsResource, devicefarmprojectsKind, c.ns, opts), &v1alpha1.DevicefarmProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevicefarmProjectList{ListMeta: obj.(*v1alpha1.DevicefarmProjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevicefarmProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devicefarmProjects.
func (c *FakeDevicefarmProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devicefarmprojectsResource, c.ns, opts))

}

// Create takes the representation of a devicefarmProject and creates it.  Returns the server's representation of the devicefarmProject, and an error, if there is any.
func (c *FakeDevicefarmProjects) Create(ctx context.Context, devicefarmProject *v1alpha1.DevicefarmProject, opts v1.CreateOptions) (result *v1alpha1.DevicefarmProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devicefarmprojectsResource, c.ns, devicefarmProject), &v1alpha1.DevicefarmProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevicefarmProject), err
}

// Update takes the representation of a devicefarmProject and updates it. Returns the server's representation of the devicefarmProject, and an error, if there is any.
func (c *FakeDevicefarmProjects) Update(ctx context.Context, devicefarmProject *v1alpha1.DevicefarmProject, opts v1.UpdateOptions) (result *v1alpha1.DevicefarmProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devicefarmprojectsResource, c.ns, devicefarmProject), &v1alpha1.DevicefarmProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevicefarmProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevicefarmProjects) UpdateStatus(ctx context.Context, devicefarmProject *v1alpha1.DevicefarmProject, opts v1.UpdateOptions) (*v1alpha1.DevicefarmProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devicefarmprojectsResource, "status", c.ns, devicefarmProject), &v1alpha1.DevicefarmProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevicefarmProject), err
}

// Delete takes name of the devicefarmProject and deletes it. Returns an error if one occurs.
func (c *FakeDevicefarmProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devicefarmprojectsResource, c.ns, name), &v1alpha1.DevicefarmProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevicefarmProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devicefarmprojectsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevicefarmProjectList{})
	return err
}

// Patch applies the patch and returns the patched devicefarmProject.
func (c *FakeDevicefarmProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DevicefarmProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devicefarmprojectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DevicefarmProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevicefarmProject), err
}
