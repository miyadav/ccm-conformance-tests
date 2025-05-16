package core

import (
	"context"
	"github.com/miyadav/ccm-conformance-tests/tests/e2e/framework"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = ginkgo.Describe("CCM Node Management", func() {
	ginkgo.Context("When a new node is present", func() {
		ginkgo.It("should be registered and have valid addresses", func() {
			nodes, err := framework.F.ClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			gomega.Expect(nodes.Items).NotTo(gomega.BeEmpty())

			for _, node := range nodes.Items {
				gomega.Expect(node.Status.Addresses).To(gomega.ContainElement(gomega.Satisfy(func(addr corev1.NodeAddress) bool {
					return addr.Type == corev1.NodeInternalIP && addr.Address != ""
				})))
			}
		})
	})
})
