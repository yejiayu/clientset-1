/*
Copyright 2017 caicloud authors. All rights reserved.
*/

// This file was automatically generated by informer-gen

package v1beta1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CustomResourceDefinitions returns a CustomResourceDefinitionInformer.
	CustomResourceDefinitions() CustomResourceDefinitionInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// CustomResourceDefinitions returns a CustomResourceDefinitionInformer.
func (v *version) CustomResourceDefinitions() CustomResourceDefinitionInformer {
	return &customResourceDefinitionInformer{factory: v.SharedInformerFactory}
}
