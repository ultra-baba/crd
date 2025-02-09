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
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
)

// FakeClusterTraincrds implements ClusterTraincrdInterface
type FakeClusterTraincrds struct {
	Fake *FakeDecisionV1
}

var clustertraincrdsResource = schema.GroupVersionResource{Group: "decision.finupgroup.com", Version: "v1", Resource: "clustertraincrds"}

var clustertraincrdsKind = schema.GroupVersionKind{Group: "decision.finupgroup.com", Version: "v1", Kind: "ClusterTraincrd"}

// Get takes name of the clusterTraincrd, and returns the corresponding clusterTraincrd object, and an error if there is any.
func (c *FakeClusterTraincrds) Get(name string, options v1.GetOptions) (result *apisv1.ClusterTraincrd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustertraincrdsResource, name), &apisv1.ClusterTraincrd{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.ClusterTraincrd), err
}

// List takes label and field selectors, and returns the list of ClusterTraincrds that match those selectors.
func (c *FakeClusterTraincrds) List(opts v1.ListOptions) (result *apisv1.ClusterTraincrdList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustertraincrdsResource, clustertraincrdsKind, opts), &apisv1.ClusterTraincrdList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apisv1.ClusterTraincrdList{ListMeta: obj.(*apisv1.ClusterTraincrdList).ListMeta}
	for _, item := range obj.(*apisv1.ClusterTraincrdList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTraincrds.
func (c *FakeClusterTraincrds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustertraincrdsResource, opts))
}

// Create takes the representation of a clusterTraincrd and creates it.  Returns the server's representation of the clusterTraincrd, and an error, if there is any.
func (c *FakeClusterTraincrds) Create(clusterTraincrd *apisv1.ClusterTraincrd) (result *apisv1.ClusterTraincrd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustertraincrdsResource, clusterTraincrd), &apisv1.ClusterTraincrd{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.ClusterTraincrd), err
}

// Update takes the representation of a clusterTraincrd and updates it. Returns the server's representation of the clusterTraincrd, and an error, if there is any.
func (c *FakeClusterTraincrds) Update(clusterTraincrd *apisv1.ClusterTraincrd) (result *apisv1.ClusterTraincrd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustertraincrdsResource, clusterTraincrd), &apisv1.ClusterTraincrd{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.ClusterTraincrd), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterTraincrds) UpdateStatus(clusterTraincrd *apisv1.ClusterTraincrd) (*apisv1.ClusterTraincrd, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustertraincrdsResource, "status", clusterTraincrd), &apisv1.ClusterTraincrd{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.ClusterTraincrd), err
}

// Delete takes name of the clusterTraincrd and deletes it. Returns an error if one occurs.
func (c *FakeClusterTraincrds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clustertraincrdsResource, name), &apisv1.ClusterTraincrd{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTraincrds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustertraincrdsResource, listOptions)

	_, err := c.Fake.Invokes(action, &apisv1.ClusterTraincrdList{})
	return err
}

// Patch applies the patch and returns the patched clusterTraincrd.
func (c *FakeClusterTraincrds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apisv1.ClusterTraincrd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustertraincrdsResource, name, pt, data, subresources...), &apisv1.ClusterTraincrd{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.ClusterTraincrd), err
}

// GetScale takes name of the clusterTraincrd, and returns the corresponding scale object, and an error if there is any.
func (c *FakeClusterTraincrds) GetScale(clusterTraincrdName string, options v1.GetOptions) (result *autoscaling.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetSubresourceAction(clustertraincrdsResource, "scale", clusterTraincrdName), &autoscaling.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaling.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeClusterTraincrds) UpdateScale(clusterTraincrdName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustertraincrdsResource, "scale", scale), &autoscaling.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaling.Scale), err
}
