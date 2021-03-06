/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// This file was automatically generated by informer-gen

package v1beta1

import (
	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1beta1 "github.com/caicloud/clientset/listers/resource/v1beta1"
	resource_v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	client_go_kubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// StorageServiceInformer provides access to a shared informer and lister for
// StorageServices.
type StorageServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.StorageServiceLister
}

type storageServiceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newStorageServiceInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ResourceV1beta1().StorageServices().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ResourceV1beta1().StorageServices().Watch(options)
			},
		},
		&resource_v1beta1.StorageService{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *storageServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resource_v1beta1.StorageService{}, func(client client_go_kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
		// panic if client is not *kubernetes.Clientset
		return newStorageServiceInformer(client.(kubernetes.Interface), resyncPeriod)
	})
}

func (f *storageServiceInformer) Lister() v1beta1.StorageServiceLister {
	return v1beta1.NewStorageServiceLister(f.Informer().GetIndexer())
}
