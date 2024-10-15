package framework

import (
	"api-tests/utils"

	. "github.com/onsi/gomega"
	machineconfigv1 "github.com/openshift/client-go/machineconfiguration/clientset/versioned/typed/machineconfiguration/v1"
	route "github.com/openshift/client-go/route/clientset/versioned"

	configclientv1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// TestContext holds reusable values across tests, including a random name for resources
type TestContext struct {
	Config      *rest.Config
	KubeClient  *kubernetes.Clientset
	RouteClient *route.Clientset
	MCPClient   *machineconfigv1.MachineconfigurationV1Client
	COClient    *configclientv1.ConfigV1Client
	Suffix      string
}

// Setup initializes the environment (e.g., auth, logging) and sets the random name for each test
func Setup() *TestContext {

	config, err := utils.Authenticate()
	Expect(err).ToNot(HaveOccurred(), "Failed to authenticate with Kubernetes")

	kubeClient, err := kubernetes.NewForConfig(config)
	Expect(err).ToNot(HaveOccurred(), "Failed to authenticate with kubernetes")

	routeClient, err := route.NewForConfig(config)
	Expect(err).ToNot(HaveOccurred(), "Failed to create Route client")

	mcpClient, err := machineconfigv1.NewForConfig(config)
	Expect(err).ToNot(HaveOccurred(), "Failed to create MCP client")

	coClient, err := configclientv1.NewForConfig(config)
	Expect(err).ToNot(HaveOccurred(), "Failed to create ClusterOperator client")

	suffix := utils.GenerateRandomName()

	return &TestContext{
		KubeClient:  kubeClient,
		Config:      config,
		RouteClient: routeClient,
		MCPClient:   mcpClient,
		COClient:    coClient,
		Suffix:      suffix,
	}
}
