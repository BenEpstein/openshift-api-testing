# Openshift API Testing

This project is designed to test the Kubernetes and OpenShift API server using Go, Ginkgo, and Gomega. It includes several tests to verify the health and status of various components within a KuberneOpenshift cluster.

## Project Structure

- `auth.go`: Contains authentication logic for connecting to the Kubernetes/OpenShift cluster.
- `tests/`: Directory containing the test files using Ginkgo and Gomega.
  - `api_tests_suite_test.go`: Sets up the test suite using Ginkgo.
  - `mcp_health_test.go`: Tests related to the health of Machine Config Pools.
  - `cluster_operator_status_test.go`: Tests for checking the status of Cluster Operators.
- `environment/`: Shell scripts to set up the Go environment and install dependencies.
  - `create-go-module.sh`: Creates the Go module for the project.
  - `install-go-libraries.sh`: Installs the required Go libraries, including Kubernetes and OpenShift client libraries, Ginkgo, and Gomega.
  - `install-golang.sh`: Script to download and install Go.
- `utils/`: folder which includes basic utility functions to help simplify test writing. 
  - `auth.go`: Used to authenticate to any Openshift/k8s client API.

## Getting Started

### Prerequisites

- Go 1.23.1 or later
- Kubernetes Cluster (OpenShift 4.12.35)
- Access to the Kubernetes/OpenShift API server

### Installation

1. **Set up Go environment in Fedora**

   Run the `install-golang.sh` script to install Go and set up the environment variables.

   ```bash
   chmod +x scripts/install-golang.sh
   ./scripts/install-golang.sh

2. **Create the Go module**

   Execute the `create-go-module.sh` script to create the Go module.

   ```bash
   chmod +x scripts/create-go-module.sh
   ./scripts/create-go-module.sh
   ```

3. **Install dependencies**

   Run the `install-go-libraries.sh` script to install the necessary Go libraries.

   ```bash
   chmod +x scripts/install-go-libraries.sh
   ./scripts/install-go-libraries.sh
   ```

### Running Tests

To execute the tests, use the Ginkgo CLI tool. The tests are located in the `tests/` directory.

```bash
ginkgo -v tests/
```

## Project Dependencies

- [Ginkgo](https://onsi.github.io/ginkgo/) - BDD Testing Framework for Go
- [Gomega](https://onsi.github.io/gomega/) - Matcher library for Go
- Kubernetes client-go
- OpenShift client-go

## Notes

- Ensure you have the correct access to your Kubernetes/OpenShift cluster.
- The Ginkgo CLI tool is required to run the tests.

