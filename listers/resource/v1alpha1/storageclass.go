/*
Copyright 2017 caicloud authors. All rights reserved.
*/

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/resource/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageClassLister helps list StorageClasses.
type StorageClassLister interface {
	// List lists all StorageClasses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageClass, err error)
	// Get retrieves the StorageClass from the index for a given name.
	Get(name string) (*v1alpha1.StorageClass, error)
	StorageClassListerExpansion
}

// storageClassLister implements the StorageClassLister interface.
type storageClassLister struct {
	indexer cache.Indexer
}

// NewStorageClassLister returns a new StorageClassLister.
func NewStorageClassLister(indexer cache.Indexer) StorageClassLister {
	return &storageClassLister{indexer: indexer}
}

// List lists all StorageClasses in the indexer.
func (s *storageClassLister) List(selector labels.Selector) (ret []*v1alpha1.StorageClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageClass))
	})
	return ret, err
}

// Get retrieves the StorageClass from the index for a given name.
func (s *storageClassLister) Get(name string) (*v1alpha1.StorageClass, error) {
	key := &v1alpha1.StorageClass{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storageclass"), name)
	}
	return obj.(*v1alpha1.StorageClass), nil
}
