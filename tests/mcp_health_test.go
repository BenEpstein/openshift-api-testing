package tests

import (
	"fmt"

	"api-tests/utils"

	mcfgv1 "github.com/openshift/api/machineconfiguration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MachineConfigPool Health Check", func() {

	It("should ensure all MachineConfigPools are updated and healthy", func() {
		client, err := utils.InitializeMCPClient()
		if err != nil {
			fmt.Printf("Error creating MCP client: %v\n", err)
			return
		}

		mcpList, err := utils.GetMCP(client)
		Expect(err).NotTo(HaveOccurred(), "Failed to list MachineConfigPools")

		for _, mcp := range mcpList {

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
			updatedCondition := utils.GetMCPCondition(mcp.Status.Conditions, mcfgv1.MachineConfigPoolUpdated)
			Expect(updatedCondition).NotTo(BeNil(),
				fmt.Sprintf("MCP '%s' does not have an 'Updated' condition", mcp.Name))
			Expect(string(updatedCondition.Status)).To(Equal(string(metav1.ConditionTrue)),
				fmt.Sprintf("MCP '%s' is not in 'Updated' state", mcp.Name))

			degradedCondition := utils.GetMCPCondition(mcp.Status.Conditions, mcfgv1.MachineConfigPoolDegraded)
			Expect(degradedCondition).NotTo(BeNil(),
				fmt.Sprintf("MCP '%s' does not have a 'Degraded' condition", mcp.Name))
			Expect(string(degradedCondition.Status)).To(Equal(string(metav1.ConditionFalse)),
				fmt.Sprintf("MCP '%s' is in 'Degraded' state", mcp.Name))
		}
	})
})
