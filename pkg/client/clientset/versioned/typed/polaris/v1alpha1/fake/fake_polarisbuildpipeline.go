/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "github.com/synthesis-labs/polaris-operator/pkg/apis/polaris/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolarisBuildPipelines implements PolarisBuildPipelineInterface
type FakePolarisBuildPipelines struct {
	Fake *FakePolarisV1alpha1
	ns   string
}

var polarisbuildpipelinesResource = schema.GroupVersionResource{Group: "polaris.synthesis.co.za", Version: "v1alpha1", Resource: "polarisbuildpipelines"}

var polarisbuildpipelinesKind = schema.GroupVersionKind{Group: "polaris.synthesis.co.za", Version: "v1alpha1", Kind: "PolarisBuildPipeline"}

// Get takes name of the polarisBuildPipeline, and returns the corresponding polarisBuildPipeline object, and an error if there is any.
func (c *FakePolarisBuildPipelines) Get(name string, options v1.GetOptions) (result *v1alpha1.PolarisBuildPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(polarisbuildpipelinesResource, c.ns, name), &v1alpha1.PolarisBuildPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolarisBuildPipeline), err
}

// List takes label and field selectors, and returns the list of PolarisBuildPipelines that match those selectors.
func (c *FakePolarisBuildPipelines) List(opts v1.ListOptions) (result *v1alpha1.PolarisBuildPipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(polarisbuildpipelinesResource, polarisbuildpipelinesKind, c.ns, opts), &v1alpha1.PolarisBuildPipelineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolarisBuildPipelineList{ListMeta: obj.(*v1alpha1.PolarisBuildPipelineList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolarisBuildPipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested polarisBuildPipelines.
func (c *FakePolarisBuildPipelines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(polarisbuildpipelinesResource, c.ns, opts))

}

// Create takes the representation of a polarisBuildPipeline and creates it.  Returns the server's representation of the polarisBuildPipeline, and an error, if there is any.
func (c *FakePolarisBuildPipelines) Create(polarisBuildPipeline *v1alpha1.PolarisBuildPipeline) (result *v1alpha1.PolarisBuildPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(polarisbuildpipelinesResource, c.ns, polarisBuildPipeline), &v1alpha1.PolarisBuildPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolarisBuildPipeline), err
}

// Update takes the representation of a polarisBuildPipeline and updates it. Returns the server's representation of the polarisBuildPipeline, and an error, if there is any.
func (c *FakePolarisBuildPipelines) Update(polarisBuildPipeline *v1alpha1.PolarisBuildPipeline) (result *v1alpha1.PolarisBuildPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(polarisbuildpipelinesResource, c.ns, polarisBuildPipeline), &v1alpha1.PolarisBuildPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolarisBuildPipeline), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolarisBuildPipelines) UpdateStatus(polarisBuildPipeline *v1alpha1.PolarisBuildPipeline) (*v1alpha1.PolarisBuildPipeline, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(polarisbuildpipelinesResource, "status", c.ns, polarisBuildPipeline), &v1alpha1.PolarisBuildPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolarisBuildPipeline), err
}

// Delete takes name of the polarisBuildPipeline and deletes it. Returns an error if one occurs.
func (c *FakePolarisBuildPipelines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(polarisbuildpipelinesResource, c.ns, name), &v1alpha1.PolarisBuildPipeline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolarisBuildPipelines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(polarisbuildpipelinesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolarisBuildPipelineList{})
	return err
}

// Patch applies the patch and returns the patched polarisBuildPipeline.
func (c *FakePolarisBuildPipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PolarisBuildPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(polarisbuildpipelinesResource, c.ns, name, data, subresources...), &v1alpha1.PolarisBuildPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolarisBuildPipeline), err
}
