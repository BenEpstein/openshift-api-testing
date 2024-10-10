package tests

import (
	"api-tests/utils"
	"fmt"

	configv1 "github.com/openshift/api/config/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClusterOperator Status Check", func() {
	It("should ensure all ClusterOperators are Available=True, Progressing=False, Degraded=False", func() {
		client, err := utils.InitializeClusterOperatorClient()
		if err != nil {
			fmt.Printf("Error creating MCP client: %v\n", err)
			return
		}

		coList, err := utils.GetClusterOperators(client)
		Expect(err).NotTo(HaveOccurred(), "Failed to list MachineConfigPools")

		for _, co := range coList {

			By(fmt.Sprintf("Validating ClusterOperator: %s", co.Name))

			// Check Available condition
			available := utils.GetCOConditionStatus(co.Status.Conditions, configv1.OperatorAvailable)
			Expect(available).To(Equal(configv1.ConditionTrue),
				fmt.Sprintf("ClusterOperator '%s' is not Available (status: %s)", co.Name, available))

			// Check Progressing condition
			progressing := utils.GetCOConditionStatus(co.Status.Conditions, configv1.OperatorProgressing)
			Expect(progressing).To(Equal(configv1.ConditionFalse),
				fmt.Sprintf("ClusterOperator '%s' is Progressing (status: %s)", co.Name, progressing))

			// Check Degraded condition
			degraded := utils.GetCOConditionStatus(co.Status.Conditions, configv1.OperatorDegraded)
			Expect(degraded).To(Equal(configv1.ConditionFalse),
				fmt.Sprintf("ClusterOperator '%s' is Degraded (status: %s)", co.Name, degraded))
		}
	})
})
