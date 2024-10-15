package tests

import (
	"api-tests/utils"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
)

var _ = Describe("Job Scheduling Tests", func() {
	It("should create and complete a job in a specific availability zone", func() {
		// Define job parameters
		jobBaseName := "general-cluster-scheduling-test"
		namespace := "core"

		// Define the az variable for specific AZ values
		availabilityZones := []string{"az-a", "az-b", "az-c"} // Replace these with valid AZ values for your cluster

		for _, az := range availabilityZones {
			jobName := fmt.Sprintf("%s-%s-%s", jobBaseName, az, ctx.Suffix)

			// Create a job in the specified AZ with the unique name
			job, err := utils.CreateJobInZone(ctx.KubeClient, namespace, jobName, az)
			Expect(err).ToNot(HaveOccurred(), "Failed to create job in specific AZ")

			fmt.Printf("Job %s created in AZ %s\n", job.Name, az)

			// Use Eventually to verify that the job completes successfully
			Eventually(func() (batchv1.JobConditionType, error) {
				return utils.CheckJobCompletion(ctx.KubeClient, namespace, job.Name)
			}, 60*time.Second, 5*time.Second).Should(Equal(batchv1.JobComplete), "Job did not reach the Completed state")
		}
	})
})
