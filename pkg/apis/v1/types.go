/*
Copyright 2015 The Kubernetes Authors.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Train is a top-level type. A client is created for it.
type Traincrd struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec   TraincrdSpec   `json:"spec"`
	Status TraincrdStatus `json:"status,omitempty"`
}

type TraincrdSpec struct {
	Image     string `json:"image"`
	Cpu string `json:"cpu"`
	Memory string `json:"memory"`
	ReqCpu string `json:"reqcpu,omitempty"`  //如果
	ReqMemory string `json:"reqmemory,omitempty"`
	Replicas int `json:"replicas,omitempty"`
	Capacity string `json:"capacity,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TraincrdList is a top-level list type. The client methods for lists are automatically created.
// You are not supposed to create a separated client for this one.
type TraincrdList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Traincrd `json:"items"`
}

type TraincrdStatus struct {
	Blah string
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterTraincrdList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ClusterTraincrd
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:method=GetScale,verb=get,subresource=scale,result=k8s.io/kubernetes/pkg/apis/autoscaling.Scale
// +genclient:method=UpdateScale,verb=update,subresource=scale,input=k8s.io/kubernetes/pkg/apis/autoscaling.Scale,result=k8s.io/kubernetes/pkg/apis/autoscaling.Scale

type ClusterTraincrd struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Status ClusterTraincrdStatus `json:"status,omitempty"`
}

type ClusterTraincrdStatus struct {
	Blah string
}
