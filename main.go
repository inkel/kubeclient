package kubeclient

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func New(inCluster bool, context string) (*kubernetes.Clientset, error) {
	cfg, err := NewClientConfig(inCluster, context)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(cfg)
}

func NewClientConfig(inCluster bool, context string) (clientcmd.ClientConfig, error) {
	if inCluster {
		return rest.InClusterConfig()
	}

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{
		CurrentContext: context,
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides), nil
}
