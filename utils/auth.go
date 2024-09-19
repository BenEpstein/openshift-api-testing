package utils

import (
	"os"
	"path/filepath"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Authenticate returns a Kubernetes rest.Config based on in-cluster config or kubeconfig file.
func Authenticate() (*rest.Config, error) {
	// Try in-cluster config
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	// Fallback to kubeconfig
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}

	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// GetClient returns a dynamic client for interacting with arbitrary Kubernetes resources
func GetClient() (dynamic.Interface, error) {
	config, err := Authenticate()
	if err != nil {
		return nil, err
	}
	return dynamic.NewForConfig(config)
}

// GetResourceClient returns a dynamic ResourceInterface for a specific resource
func GetResourceClient(group, version, resource string) (dynamic.NamespaceableResourceInterface, error) {
	dynClient, err := GetClient()
	if err != nil {
		return nil, err
	}

	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}

	return dynClient.Resource(gvr), nil
}
