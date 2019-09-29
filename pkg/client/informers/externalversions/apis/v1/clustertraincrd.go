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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	apisv1 "finupgroup.com/decision/traincrd/pkg/apis/v1"
	versioned "finupgroup.com/decision/traincrd/pkg/client/clientset/versioned"
	internalinterfaces "finupgroup.com/decision/traincrd/pkg/client/informers/externalversions/internalinterfaces"
	v1 "finupgroup.com/decision/traincrd/pkg/client/listers/apis/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterTraincrdInformer provides access to a shared informer and lister for
// ClusterTraincrds.
type ClusterTraincrdInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterTraincrdLister
}

type clusterTraincrdInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterTraincrdInformer constructs a new informer for ClusterTraincrd type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterTraincrdInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterTraincrdInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterTraincrdInformer constructs a new informer for ClusterTraincrd type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterTraincrdInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DecisionV1().ClusterTraincrds().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DecisionV1().ClusterTraincrds().Watch(options)
			},
		},
		&apisv1.ClusterTraincrd{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterTraincrdInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterTraincrdInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterTraincrdInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisv1.ClusterTraincrd{}, f.defaultInformer)
}

func (f *clusterTraincrdInformer) Lister() v1.ClusterTraincrdLister {
	return v1.NewClusterTraincrdLister(f.Informer().GetIndexer())
}
