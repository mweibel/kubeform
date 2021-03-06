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

// FakeNodebalancerNodes implements NodebalancerNodeInterface
type FakeNodebalancerNodes struct {
	Fake *FakeLinodeV1alpha1
	ns   string
}

var nodebalancernodesResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "nodebalancernodes"}

var nodebalancernodesKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "NodebalancerNode"}

// Get takes name of the nodebalancerNode, and returns the corresponding nodebalancerNode object, and an error if there is any.
func (c *FakeNodebalancerNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodebalancerNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodebalancernodesResource, c.ns, name), &v1alpha1.NodebalancerNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerNode), err
}

// List takes label and field selectors, and returns the list of NodebalancerNodes that match those selectors.
func (c *FakeNodebalancerNodes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodebalancerNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodebalancernodesResource, nodebalancernodesKind, c.ns, opts), &v1alpha1.NodebalancerNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodebalancerNodeList{ListMeta: obj.(*v1alpha1.NodebalancerNodeList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodebalancerNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodebalancerNodes.
func (c *FakeNodebalancerNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodebalancernodesResource, c.ns, opts))

}

// Create takes the representation of a nodebalancerNode and creates it.  Returns the server's representation of the nodebalancerNode, and an error, if there is any.
func (c *FakeNodebalancerNodes) Create(ctx context.Context, nodebalancerNode *v1alpha1.NodebalancerNode, opts v1.CreateOptions) (result *v1alpha1.NodebalancerNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodebalancernodesResource, c.ns, nodebalancerNode), &v1alpha1.NodebalancerNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerNode), err
}

// Update takes the representation of a nodebalancerNode and updates it. Returns the server's representation of the nodebalancerNode, and an error, if there is any.
func (c *FakeNodebalancerNodes) Update(ctx context.Context, nodebalancerNode *v1alpha1.NodebalancerNode, opts v1.UpdateOptions) (result *v1alpha1.NodebalancerNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodebalancernodesResource, c.ns, nodebalancerNode), &v1alpha1.NodebalancerNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerNode), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodebalancerNodes) UpdateStatus(ctx context.Context, nodebalancerNode *v1alpha1.NodebalancerNode, opts v1.UpdateOptions) (*v1alpha1.NodebalancerNode, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodebalancernodesResource, "status", c.ns, nodebalancerNode), &v1alpha1.NodebalancerNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerNode), err
}

// Delete takes name of the nodebalancerNode and deletes it. Returns an error if one occurs.
func (c *FakeNodebalancerNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodebalancernodesResource, c.ns, name), &v1alpha1.NodebalancerNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodebalancerNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodebalancernodesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodebalancerNodeList{})
	return err
}

// Patch applies the patch and returns the patched nodebalancerNode.
func (c *FakeNodebalancerNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodebalancerNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodebalancernodesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NodebalancerNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerNode), err
}
