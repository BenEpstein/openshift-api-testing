package tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApiTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApiTests Suite")
}
