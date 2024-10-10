package utils

import (
	"k8s.io/client-go/kubernetes"
)

// GetKubernetesClientset creates a Kubernetes clientset.
func GetKubernetesClientset() (*kubernetes.Clientset, error) {
	config, err := Authenticate()
	if err != nil {
		return nil, err
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func GetDynamicClientset() (*kubernetes.Clientset, error) {
	config, err := Authenticate()
	if err != nil {
		return nil, err
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
