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

// TraincrdInformer provides access to a shared informer and lister for
// Traincrds.
type TraincrdInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TraincrdLister
}

type traincrdInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTraincrdInformer constructs a new informer for Traincrd type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTraincrdInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTraincrdInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTraincrdInformer constructs a new informer for Traincrd type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTraincrdInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DecisionV1().Traincrds(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DecisionV1().Traincrds(namespace).Watch(options)
			},
		},
		&apisv1.Traincrd{},
		resyncPeriod,
		indexers,
	)
}

func (f *traincrdInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTraincrdInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *traincrdInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisv1.Traincrd{}, f.defaultInformer)
}

func (f *traincrdInformer) Lister() v1.TraincrdLister {
	return v1.NewTraincrdLister(f.Informer().GetIndexer())
}
