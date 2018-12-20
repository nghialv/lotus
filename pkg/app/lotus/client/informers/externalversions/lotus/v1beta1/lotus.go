// Copyright (c) 2018 Lotus Load
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	lotusv1beta1 "github.com/lotusload/lotus/pkg/app/lotus/apis/lotus/v1beta1"
	versioned "github.com/lotusload/lotus/pkg/app/lotus/client/clientset/versioned"
	internalinterfaces "github.com/lotusload/lotus/pkg/app/lotus/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/lotusload/lotus/pkg/app/lotus/client/listers/lotus/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LotusInformer provides access to a shared informer and lister for
// Lotuses.
type LotusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.LotusLister
}

type lotusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLotusInformer constructs a new informer for Lotus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLotusInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLotusInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLotusInformer constructs a new informer for Lotus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLotusInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LotusV1beta1().Lotuses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LotusV1beta1().Lotuses(namespace).Watch(options)
			},
		},
		&lotusv1beta1.Lotus{},
		resyncPeriod,
		indexers,
	)
}

func (f *lotusInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLotusInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *lotusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&lotusv1beta1.Lotus{}, f.defaultInformer)
}

func (f *lotusInformer) Lister() v1beta1.LotusLister {
	return v1beta1.NewLotusLister(f.Informer().GetIndexer())
}
