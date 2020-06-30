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

// FakeSsmPatchGroups implements SsmPatchGroupInterface
type FakeSsmPatchGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ssmpatchgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ssmpatchgroups"}

var ssmpatchgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SsmPatchGroup"}

// Get takes name of the ssmPatchGroup, and returns the corresponding ssmPatchGroup object, and an error if there is any.
func (c *FakeSsmPatchGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SsmPatchGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ssmpatchgroupsResource, c.ns, name), &v1alpha1.SsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmPatchGroup), err
}

// List takes label and field selectors, and returns the list of SsmPatchGroups that match those selectors.
func (c *FakeSsmPatchGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SsmPatchGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ssmpatchgroupsResource, ssmpatchgroupsKind, c.ns, opts), &v1alpha1.SsmPatchGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SsmPatchGroupList{ListMeta: obj.(*v1alpha1.SsmPatchGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.SsmPatchGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ssmPatchGroups.
func (c *FakeSsmPatchGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ssmpatchgroupsResource, c.ns, opts))

}

// Create takes the representation of a ssmPatchGroup and creates it.  Returns the server's representation of the ssmPatchGroup, and an error, if there is any.
func (c *FakeSsmPatchGroups) Create(ctx context.Context, ssmPatchGroup *v1alpha1.SsmPatchGroup, opts v1.CreateOptions) (result *v1alpha1.SsmPatchGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ssmpatchgroupsResource, c.ns, ssmPatchGroup), &v1alpha1.SsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmPatchGroup), err
}

// Update takes the representation of a ssmPatchGroup and updates it. Returns the server's representation of the ssmPatchGroup, and an error, if there is any.
func (c *FakeSsmPatchGroups) Update(ctx context.Context, ssmPatchGroup *v1alpha1.SsmPatchGroup, opts v1.UpdateOptions) (result *v1alpha1.SsmPatchGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ssmpatchgroupsResource, c.ns, ssmPatchGroup), &v1alpha1.SsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmPatchGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSsmPatchGroups) UpdateStatus(ctx context.Context, ssmPatchGroup *v1alpha1.SsmPatchGroup, opts v1.UpdateOptions) (*v1alpha1.SsmPatchGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ssmpatchgroupsResource, "status", c.ns, ssmPatchGroup), &v1alpha1.SsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmPatchGroup), err
}

// Delete takes name of the ssmPatchGroup and deletes it. Returns an error if one occurs.
func (c *FakeSsmPatchGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ssmpatchgroupsResource, c.ns, name), &v1alpha1.SsmPatchGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSsmPatchGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ssmpatchgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SsmPatchGroupList{})
	return err
}

// Patch applies the patch and returns the patched ssmPatchGroup.
func (c *FakeSsmPatchGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SsmPatchGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ssmpatchgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SsmPatchGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmPatchGroup), err
}
