package main

import (
    "testing"

    "github.com/onsi/ginkgo/v2"
    "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
    gomega.RegisterFailHandler(ginkgo.Fail)
    ginkgo.RunSpecs(m, "Cloud Provider Interface E2E Suite")
}

