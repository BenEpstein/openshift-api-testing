# Install Kubernetes client-go matching your cluster version
go get k8s.io/client-go@v0.25.12
go get k8s.io/apimachinery@v0.25.12
go get k8s.io/api@v0.25.12

# Install OpenShift client libraries matching your OpenShift version
go get github.com/openshift/api@release-4.12
go get github.com/openshift/client-go@release-4.12

# Install Ginkgo and Gomega
go get github.com/onsi/ginkgo/v2
go get github.com/onsi/gomega

# Install the Ginkgo CLI tool
go install github.com/onsi/ginkgo/v2/ginkgo@latest

#
go get k8s.io/apimachinery/pkg/runtime/serializer/json@v0.25.12
go get k8s.io/client-go/plugin/pkg/client/auth/exec@v0.25.12
go get k8s.io/client-go/tools/clientcmd@v0.25.12
