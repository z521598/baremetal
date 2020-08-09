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
	baremetalv1 "baremetal/pkg/apis/baremetal/v1"
	versioned "baremetal/pkg/client/clientset/versioned"
	internalinterfaces "baremetal/pkg/client/informers/externalversions/internalinterfaces"
	v1 "baremetal/pkg/client/listers/baremetal/v1"
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BareMetalJobInformer provides access to a shared informer and lister for
// BareMetalJobs.
type BareMetalJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BareMetalJobLister
}

type bareMetalJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBareMetalJobInformer constructs a new informer for BareMetalJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBareMetalJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBareMetalJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBareMetalJobInformer constructs a new informer for BareMetalJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBareMetalJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MagellanV1().BareMetalJobs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MagellanV1().BareMetalJobs(namespace).Watch(context.TODO(), options)
			},
		},
		&baremetalv1.BareMetalJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *bareMetalJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBareMetalJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bareMetalJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&baremetalv1.BareMetalJob{}, f.defaultInformer)
}

func (f *bareMetalJobInformer) Lister() v1.BareMetalJobLister {
	return v1.NewBareMetalJobLister(f.Informer().GetIndexer())
}
