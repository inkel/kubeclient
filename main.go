package kubeclient

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func New(inCluster bool, context string) (*kubernetes.Clientset, *rest.Config, error) {
	cfg, err := NewClientConfig(inCluster, context)
	if err != nil {
		return nil, nil, err
	}

	c, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}

	return c, cfg, nil
}

func NewClientConfig(inCluster bool, context string) (*rest.Config, error) {
	if inCluster {
		return rest.InClusterConfig()
	}

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{
		CurrentContext: context,
	}

	cfg, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides).
		ClientConfig()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
