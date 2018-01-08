/*
Copyright 2018 caicloud authors. All rights reserved.
*/

package v1alpha1

import (
	"github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/config/v1alpha1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ConfigV1alpha1Interface interface {
	RESTClient() rest.Interface
	ConfigClaimsGetter
	ConfigReferencesGetter
}

// ConfigV1alpha1Client is used to interact with features provided by the config.caicloud.io group.
type ConfigV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ConfigV1alpha1Client) ConfigClaims(namespace string) ConfigClaimInterface {
	return newConfigClaims(c, namespace)
}

func (c *ConfigV1alpha1Client) ConfigReferences(namespace string) ConfigReferenceInterface {
	return newConfigReferences(c, namespace)
}

// NewForConfig creates a new ConfigV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ConfigV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ConfigV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ConfigV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ConfigV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ConfigV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ConfigV1alpha1Client {
	return &ConfigV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ConfigV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
