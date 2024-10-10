package utils

import (
	"context"
	"fmt"

	configv1 "github.com/openshift/api/config/v1"
	configclientv1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InitializeClusterOperatorClient returns a ClusterOperatorsGetter client for ClusterOperator operations
func InitializeClusterOperatorClient() (configclientv1.ClusterOperatorsGetter, error) {
	config, err := Authenticate()
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	return configclientv1.NewForConfig(config)
}

// GetClusterOperators retrieves one, multiple, or all ClusterOperator objects by name
func GetClusterOperators(coClient configclientv1.ClusterOperatorsGetter, coNames ...string) ([]configv1.ClusterOperator, error) {
	var coList []configv1.ClusterOperator

	if len(coNames) == 0 {
		// Retrieve all ClusterOperators
		coObjects, err := coClient.ClusterOperators().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to list ClusterOperators: %w", err)
		}
		coList = append(coList, coObjects.Items...)
	} else {
		// Retrieve specific ClusterOperators
		for _, name := range coNames {
			co, err := coClient.ClusterOperators().Get(context.TODO(), name, metav1.GetOptions{})
			if err != nil {
				return nil, fmt.Errorf("failed to get ClusterOperator %s: %w", name, err)
			}
			coList = append(coList, *co)
		}
	}

	return coList, nil
}

// Helper function to get the status of a condition
func GetCOConditionStatus(conditions []configv1.ClusterOperatorStatusCondition, conditionType configv1.ClusterStatusConditionType) configv1.ConditionStatus {
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return condition.Status
		}
	}
	return configv1.ConditionUnknown
}
