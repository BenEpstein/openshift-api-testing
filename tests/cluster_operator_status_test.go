package tests

import (
	"api-tests/utils"
	"context"
	"fmt"

	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClusterOperator Status Check", func() {
	It("should ensure all ClusterOperators are Available=True, Progressing=False, Degraded=False", func() {
		configClient, err := utils.GetResourceClient("config.openshift.io", "v1", "clusteroperators")
		if err != nil {
			fmt.Printf("Error creating COs client: %v\n", err)
			return
		}
		coList, err := configClient.List(context.TODO(), metav1.ListOptions{})
		Expect(err).NotTo(HaveOccurred(), "Failed to list ClusterOperators")

		for _, coUnstructured := range coList.Items {
			var co configv1.ClusterOperator
			err := runtime.DefaultUnstructuredConverter.FromUnstructured(coUnstructured.Object, &co)
			Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("Failed to convert Unstructured to api resource for %s", coUnstructured.GetName()))

			By(fmt.Sprintf("Validating ClusterOperator: %s", co.Name))

			// Check Available condition
			available := getConditionStatus(co.Status.Conditions, configv1.OperatorAvailable)
			Expect(available).To(Equal(configv1.ConditionTrue),
				fmt.Sprintf("ClusterOperator '%s' is not Available (status: %s)", co.Name, available))

			// Check Progressing condition
			progressing := getConditionStatus(co.Status.Conditions, configv1.OperatorProgressing)
			Expect(progressing).To(Equal(configv1.ConditionFalse),
				fmt.Sprintf("ClusterOperator '%s' is Progressing (status: %s)", co.Name, progressing))

			// Check Degraded condition
			degraded := getConditionStatus(co.Status.Conditions, configv1.OperatorDegraded)
			Expect(degraded).To(Equal(configv1.ConditionFalse),
				fmt.Sprintf("ClusterOperator '%s' is Degraded (status: %s)", co.Name, degraded))
		}
	})
})

// Helper function to get the status of a condition
func getConditionStatus(conditions []configv1.ClusterOperatorStatusCondition, conditionType configv1.ClusterStatusConditionType) configv1.ConditionStatus {
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return condition.Status
		}
	}
	return configv1.ConditionUnknown
}
