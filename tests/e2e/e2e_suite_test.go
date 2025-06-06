package e2e

import (
	"testing"

	"github.com/miyadav/ccm-conformance-tests/tests/e2e/framework"
	_ "github.com/miyadav/ccm-conformance-tests/tests/e2e/suites/core"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "E2E Suite")
}

var F *framework.Framework

var _ = ginkgo.BeforeSuite(func() {
	F = framework.NewDefaultFramework("e2e")
})
