package utils

import (
	"context"
	"fmt"

	mcfgv1 "github.com/openshift/api/machineconfiguration/v1"
	machineconfigv1 "github.com/openshift/client-go/machineconfiguration/clientset/versioned/typed/machineconfiguration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetMCP retrieves MachineConfigPool objects
func GetMCP(mcpClient machineconfigv1.MachineconfigurationV1Client, mcpNames ...string) ([]mcfgv1.MachineConfigPool, error) {
	var mcpList []mcfgv1.MachineConfigPool

	if len(mcpNames) == 0 {
		mcpObjects, err := mcpClient.MachineConfigPools().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to list MCPs: %w", err)
		}
		mcpList = append(mcpList, mcpObjects.Items...)
	} else {
		for _, name := range mcpNames {
			mcp, err := mcpClient.MachineConfigPools().Get(context.TODO(), name, metav1.GetOptions{})
			if err != nil {
				return nil, fmt.Errorf("failed to get MCP %s: %w", name, err)
			}
			mcpList = append(mcpList, *mcp)
		}
	}

	return mcpList, nil
}

// Helper function to get a specific condition from MCP status
func GetMCPCondition(conditions []mcfgv1.MachineConfigPoolCondition, conditionType mcfgv1.MachineConfigPoolConditionType) *mcfgv1.MachineConfigPoolCondition {
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return &condition
		}
	}
	return nil
}
