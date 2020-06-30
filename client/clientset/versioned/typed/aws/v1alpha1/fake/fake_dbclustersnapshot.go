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

// FakeDbClusterSnapshots implements DbClusterSnapshotInterface
type FakeDbClusterSnapshots struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dbclustersnapshotsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dbclustersnapshots"}

var dbclustersnapshotsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DbClusterSnapshot"}

// Get takes name of the dbClusterSnapshot, and returns the corresponding dbClusterSnapshot object, and an error if there is any.
func (c *FakeDbClusterSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dbclustersnapshotsResource, c.ns, name), &v1alpha1.DbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbClusterSnapshot), err
}

// List takes label and field selectors, and returns the list of DbClusterSnapshots that match those selectors.
func (c *FakeDbClusterSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DbClusterSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dbclustersnapshotsResource, dbclustersnapshotsKind, c.ns, opts), &v1alpha1.DbClusterSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DbClusterSnapshotList{ListMeta: obj.(*v1alpha1.DbClusterSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.DbClusterSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dbClusterSnapshots.
func (c *FakeDbClusterSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dbclustersnapshotsResource, c.ns, opts))

}

// Create takes the representation of a dbClusterSnapshot and creates it.  Returns the server's representation of the dbClusterSnapshot, and an error, if there is any.
func (c *FakeDbClusterSnapshots) Create(ctx context.Context, dbClusterSnapshot *v1alpha1.DbClusterSnapshot, opts v1.CreateOptions) (result *v1alpha1.DbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dbclustersnapshotsResource, c.ns, dbClusterSnapshot), &v1alpha1.DbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbClusterSnapshot), err
}

// Update takes the representation of a dbClusterSnapshot and updates it. Returns the server's representation of the dbClusterSnapshot, and an error, if there is any.
func (c *FakeDbClusterSnapshots) Update(ctx context.Context, dbClusterSnapshot *v1alpha1.DbClusterSnapshot, opts v1.UpdateOptions) (result *v1alpha1.DbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dbclustersnapshotsResource, c.ns, dbClusterSnapshot), &v1alpha1.DbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbClusterSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDbClusterSnapshots) UpdateStatus(ctx context.Context, dbClusterSnapshot *v1alpha1.DbClusterSnapshot, opts v1.UpdateOptions) (*v1alpha1.DbClusterSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dbclustersnapshotsResource, "status", c.ns, dbClusterSnapshot), &v1alpha1.DbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbClusterSnapshot), err
}

// Delete takes name of the dbClusterSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeDbClusterSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dbclustersnapshotsResource, c.ns, name), &v1alpha1.DbClusterSnapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDbClusterSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dbclustersnapshotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DbClusterSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched dbClusterSnapshot.
func (c *FakeDbClusterSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dbclustersnapshotsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbClusterSnapshot), err
}
