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
	apisv1 "finupgroup.com/decision/traincrd/pkg/apis/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTraincrds implements TraincrdInterface
type FakeTraincrds struct {
	Fake *FakeDecisionV1
	ns   string
}

var traincrdsResource = schema.GroupVersionResource{Group: "decision.finupgroup.com", Version: "v1", Resource: "traincrds"}

var traincrdsKind = schema.GroupVersionKind{Group: "decision.finupgroup.com", Version: "v1", Kind: "Traincrd"}

// Get takes name of the traincrd, and returns the corresponding traincrd object, and an error if there is any.
func (c *FakeTraincrds) Get(name string, options v1.GetOptions) (result *apisv1.Traincrd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(traincrdsResource, c.ns, name), &apisv1.Traincrd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.Traincrd), err
}

// List takes label and field selectors, and returns the list of Traincrds that match those selectors.
func (c *FakeTraincrds) List(opts v1.ListOptions) (result *apisv1.TraincrdList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(traincrdsResource, traincrdsKind, c.ns, opts), &apisv1.TraincrdList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apisv1.TraincrdList{ListMeta: obj.(*apisv1.TraincrdList).ListMeta}
	for _, item := range obj.(*apisv1.TraincrdList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested traincrds.
func (c *FakeTraincrds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(traincrdsResource, c.ns, opts))

}

// Create takes the representation of a traincrd and creates it.  Returns the server's representation of the traincrd, and an error, if there is any.
func (c *FakeTraincrds) Create(traincrd *apisv1.Traincrd) (result *apisv1.Traincrd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(traincrdsResource, c.ns, traincrd), &apisv1.Traincrd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.Traincrd), err
}

// Update takes the representation of a traincrd and updates it. Returns the server's representation of the traincrd, and an error, if there is any.
func (c *FakeTraincrds) Update(traincrd *apisv1.Traincrd) (result *apisv1.Traincrd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(traincrdsResource, c.ns, traincrd), &apisv1.Traincrd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.Traincrd), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTraincrds) UpdateStatus(traincrd *apisv1.Traincrd) (*apisv1.Traincrd, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(traincrdsResource, "status", c.ns, traincrd), &apisv1.Traincrd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.Traincrd), err
}

// Delete takes name of the traincrd and deletes it. Returns an error if one occurs.
func (c *FakeTraincrds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(traincrdsResource, c.ns, name), &apisv1.Traincrd{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTraincrds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(traincrdsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &apisv1.TraincrdList{})
	return err
}

// Patch applies the patch and returns the patched traincrd.
func (c *FakeTraincrds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apisv1.Traincrd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(traincrdsResource, c.ns, name, pt, data, subresources...), &apisv1.Traincrd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.Traincrd), err
}
