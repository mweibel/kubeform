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

	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLkeClusters implements LkeClusterInterface
type FakeLkeClusters struct {
	Fake *FakeLinodeV1alpha1
	ns   string
}

var lkeclustersResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "lkeclusters"}

var lkeclustersKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "LkeCluster"}

// Get takes name of the lkeCluster, and returns the corresponding lkeCluster object, and an error if there is any.
func (c *FakeLkeClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LkeCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lkeclustersResource, c.ns, name), &v1alpha1.LkeCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LkeCluster), err
}

// List takes label and field selectors, and returns the list of LkeClusters that match those selectors.
func (c *FakeLkeClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LkeClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lkeclustersResource, lkeclustersKind, c.ns, opts), &v1alpha1.LkeClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LkeClusterList{ListMeta: obj.(*v1alpha1.LkeClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.LkeClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lkeClusters.
func (c *FakeLkeClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lkeclustersResource, c.ns, opts))

}

// Create takes the representation of a lkeCluster and creates it.  Returns the server's representation of the lkeCluster, and an error, if there is any.
func (c *FakeLkeClusters) Create(ctx context.Context, lkeCluster *v1alpha1.LkeCluster, opts v1.CreateOptions) (result *v1alpha1.LkeCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lkeclustersResource, c.ns, lkeCluster), &v1alpha1.LkeCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LkeCluster), err
}

// Update takes the representation of a lkeCluster and updates it. Returns the server's representation of the lkeCluster, and an error, if there is any.
func (c *FakeLkeClusters) Update(ctx context.Context, lkeCluster *v1alpha1.LkeCluster, opts v1.UpdateOptions) (result *v1alpha1.LkeCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lkeclustersResource, c.ns, lkeCluster), &v1alpha1.LkeCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LkeCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLkeClusters) UpdateStatus(ctx context.Context, lkeCluster *v1alpha1.LkeCluster, opts v1.UpdateOptions) (*v1alpha1.LkeCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lkeclustersResource, "status", c.ns, lkeCluster), &v1alpha1.LkeCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LkeCluster), err
}

// Delete takes name of the lkeCluster and deletes it. Returns an error if one occurs.
func (c *FakeLkeClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lkeclustersResource, c.ns, name), &v1alpha1.LkeCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLkeClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lkeclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LkeClusterList{})
	return err
}

// Patch applies the patch and returns the patched lkeCluster.
func (c *FakeLkeClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LkeCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lkeclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.LkeCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LkeCluster), err
}
