package tests

import (
	"context"
	"fmt"

	"api-tests/utils"

	mcfgv1 "github.com/openshift/api/machineconfiguration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MachineConfigPool Health Check", func() {

	It("should ensure all MachineConfigPools are updated and healthy", func() {
		configClient, err := utils.GetResourceClient("machineconfiguration.openshift.io", "v1", "machineconfigpools")
		if err != nil {
			fmt.Printf("Error creating MCP client: %v\n", err)
			return
		}

		mcpList, err := configClient.List(context.TODO(), metav1.ListOptions{})
		Expect(err).NotTo(HaveOccurred(), "Failed to list MachineConfigPools")

		for _, mcpUnstructured := range mcpList.Items {
			var mcp mcfgv1.MachineConfigPool
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(mcpUnstructured.Object, &mcp)
			Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("Failed to convert Unstructured to api resource for %s", mcpUnstructured.GetName()))

			By(fmt.Sprintf("Validating MachineConfigPool: %s", mcp.Name))
			//By(fmt.Sprintf("MCP '%s' Conditions: %+v", mcp.Name, mcp.Status.Conditions))

			// Check that UpdatedMachineCount equals MachineCount
			Expect(mcp.Status.UpdatedMachineCount).To(Equal(mcp.Status.MachineCount),
				fmt.Sprintf("MCP '%s' has mismatched UpdatedMachineCount (%d) and MachineCount (%d)",
					mcp.Name, mcp.Status.UpdatedMachineCount, mcp.Status.MachineCount))

			// Check that ReadyMachineCount equals MachineCount
			Expect(mcp.Status.ReadyMachineCount).To(Equal(mcp.Status.MachineCount),
				fmt.Sprintf("MCP '%s' has mismatched ReadyMachineCount (%d) and MachineCount (%d)",
					mcp.Name, mcp.Status.ReadyMachineCount, mcp.Status.MachineCount))

			// Check that DegradedMachineCount is zero
			Expect(mcp.Status.DegradedMachineCount).To(BeZero(),
				fmt.Sprintf("MCP '%s' has DegradedMachineCount of %d",
					mcp.Name, mcp.Status.DegradedMachineCount))

			// Check that the MCP is Updated and not Degraded
			updatedCondition := getMCPCondition(mcp.Status.Conditions, mcfgv1.MachineConfigPoolUpdated)
			Expect(updatedCondition).NotTo(BeNil(),
				fmt.Sprintf("MCP '%s' does not have an 'Updated' condition", mcp.Name))
			Expect(string(updatedCondition.Status)).To(Equal(string(metav1.ConditionTrue)),
				fmt.Sprintf("MCP '%s' is not in 'Updated' state", mcp.Name))

			degradedCondition := getMCPCondition(mcp.Status.Conditions, mcfgv1.MachineConfigPoolDegraded)
			Expect(degradedCondition).NotTo(BeNil(),
				fmt.Sprintf("MCP '%s' does not have a 'Degraded' condition", mcp.Name))
			Expect(string(degradedCondition.Status)).To(Equal(string(metav1.ConditionFalse)),
				fmt.Sprintf("MCP '%s' is in 'Degraded' state", mcp.Name))
		}
	})
})

// Helper function to get a specific condition from MCP status
func getMCPCondition(conditions []mcfgv1.MachineConfigPoolCondition, conditionType mcfgv1.MachineConfigPoolConditionType) *mcfgv1.MachineConfigPoolCondition {
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return &condition
		}
	}
	return nil
}
