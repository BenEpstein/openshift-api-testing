package tests

import (
	"api-tests/framework"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	ctx *framework.TestContext
)

var _ = BeforeSuite(func() {
	// Run Setup() once before any tests
	ctx = framework.Setup()
	Expect(ctx).ToNot(BeNil(), "Test context should be initialized in BeforeSuite")
})
