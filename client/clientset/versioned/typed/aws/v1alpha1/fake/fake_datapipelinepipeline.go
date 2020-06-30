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

// FakeDatapipelinePipelines implements DatapipelinePipelineInterface
type FakeDatapipelinePipelines struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var datapipelinepipelinesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "datapipelinepipelines"}

var datapipelinepipelinesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DatapipelinePipeline"}

// Get takes name of the datapipelinePipeline, and returns the corresponding datapipelinePipeline object, and an error if there is any.
func (c *FakeDatapipelinePipelines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatapipelinePipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datapipelinepipelinesResource, c.ns, name), &v1alpha1.DatapipelinePipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatapipelinePipeline), err
}

// List takes label and field selectors, and returns the list of DatapipelinePipelines that match those selectors.
func (c *FakeDatapipelinePipelines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatapipelinePipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datapipelinepipelinesResource, datapipelinepipelinesKind, c.ns, opts), &v1alpha1.DatapipelinePipelineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatapipelinePipelineList{ListMeta: obj.(*v1alpha1.DatapipelinePipelineList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatapipelinePipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datapipelinePipelines.
func (c *FakeDatapipelinePipelines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datapipelinepipelinesResource, c.ns, opts))

}

// Create takes the representation of a datapipelinePipeline and creates it.  Returns the server's representation of the datapipelinePipeline, and an error, if there is any.
func (c *FakeDatapipelinePipelines) Create(ctx context.Context, datapipelinePipeline *v1alpha1.DatapipelinePipeline, opts v1.CreateOptions) (result *v1alpha1.DatapipelinePipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datapipelinepipelinesResource, c.ns, datapipelinePipeline), &v1alpha1.DatapipelinePipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatapipelinePipeline), err
}

// Update takes the representation of a datapipelinePipeline and updates it. Returns the server's representation of the datapipelinePipeline, and an error, if there is any.
func (c *FakeDatapipelinePipelines) Update(ctx context.Context, datapipelinePipeline *v1alpha1.DatapipelinePipeline, opts v1.UpdateOptions) (result *v1alpha1.DatapipelinePipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datapipelinepipelinesResource, c.ns, datapipelinePipeline), &v1alpha1.DatapipelinePipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatapipelinePipeline), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatapipelinePipelines) UpdateStatus(ctx context.Context, datapipelinePipeline *v1alpha1.DatapipelinePipeline, opts v1.UpdateOptions) (*v1alpha1.DatapipelinePipeline, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datapipelinepipelinesResource, "status", c.ns, datapipelinePipeline), &v1alpha1.DatapipelinePipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatapipelinePipeline), err
}

// Delete takes name of the datapipelinePipeline and deletes it. Returns an error if one occurs.
func (c *FakeDatapipelinePipelines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datapipelinepipelinesResource, c.ns, name), &v1alpha1.DatapipelinePipeline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatapipelinePipelines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datapipelinepipelinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatapipelinePipelineList{})
	return err
}

// Patch applies the patch and returns the patched datapipelinePipeline.
func (c *FakeDatapipelinePipelines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatapipelinePipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datapipelinepipelinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatapipelinePipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatapipelinePipeline), err
}
