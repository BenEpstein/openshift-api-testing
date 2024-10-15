package utils

import (
	"context"
	"fmt"

	"api-tests/consts"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreateJobInZone creates a job with a node selector for a specific availability zone
func CreateJobInZone(clientset *kubernetes.Clientset, namespace, jobName, azValue string) (*batchv1.Job, error) {
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					NodeSelector: map[string]string{
						consts.AZLabel: azValue,
					},
					RestartPolicy: corev1.RestartPolicyNever,
					Containers: []corev1.Container{
						{
							Name:    "test-container",
							Image:   consts.ClientImage,
							Command: []string{"sh", "-c", "sleep 10"},
							Resources: SetResourceRequirements(
								"500m",  // CPU request
								"2",     // CPU limit
								"400Mi", // Memory request
								"400Mi", // Memory limit
							),
						},
					},
				},
			},
		},
	}

	// Create the Job
	return clientset.BatchV1().Jobs(namespace).Create(context.TODO(), job, metav1.CreateOptions{})
}

// CheckJobCompletion verifies if the job has completed successfully
func CheckJobCompletion(clientset *kubernetes.Clientset, namespace, jobName string) (batchv1.JobConditionType, error) {
	job, err := clientset.BatchV1().Jobs(namespace).Get(context.TODO(), jobName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get job: %w", err)
	}

	// Check if the job has succeeded
	for _, condition := range job.Status.Conditions {
		if condition.Type == batchv1.JobComplete && condition.Status == corev1.ConditionTrue {
			return batchv1.JobComplete, nil // Job completed successfully
		} else if condition.Type == batchv1.JobFailed && condition.Status == corev1.ConditionTrue {
			return batchv1.JobFailed, fmt.Errorf("job %s failed", jobName) // Job failed
		}
	}
	return "", nil // Job not complete yet
}
