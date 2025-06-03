package core

import (
	"context"

	"github.com/miyadav/ccm-conformance-tests/tests/e2e/framework"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = ginkgo.Describe("CCM Routes Management", func() {
	var f *framework.Framework

	ginkgo.BeforeEach(func() {
		f = framework.NewDefaultFramework("routes")
	})

	ginkgo.Context("When nodes are present in the cluster", func() {
		ginkgo.It("should ensure each node has a PodCIDR allocated ", func() {

			nodes, err := f.ClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			gomega.Expect(nodes.Items).NotTo(gomega.BeEmpty())

			for _, node := range nodes.Items {
				// A very basic check.
				gomega.Expect(node.Spec.PodCIDR).NotTo(gomega.BeEmpty(), "Node %s should have a podcidr", node.Name)

			}
		})

	})
})
