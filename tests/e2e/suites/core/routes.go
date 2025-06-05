package core

import (
	"context"
	"fmt"

	"github.com/miyadav/ccm-conformance-tests/tests/e2e/framework"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = ginkgo.Describe("CCM Routes controller", func() {
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

	var _ = ginkgo.Describe("Route/Zones Controller", func() {
		ginkgo.It("should get zone/region labels from all nodes", func() {
			nodes, err := f.ClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(nodes.Items)).To(gomega.BeNumerically(">", 0))
			for _, node := range nodes.Items {
				zone := node.Labels["topology.kubernetes.io/zone"]
				region := node.Labels["topology.kubernetes.io/region"]
				gomega.Expect(zone).NotTo(gomega.BeEmpty(), fmt.Sprintf("Node %s missing zone label", node.Name))
				gomega.Expect(region).NotTo(gomega.BeEmpty(), fmt.Sprintf("Node %s missing region label", node.Name))
			}
		})
	})
})
