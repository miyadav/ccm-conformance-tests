package core

import (
	"context"
	"fmt"
	"time"

	"github.com/miyadav/ccm-conformance-tests/tests/e2e/framework"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var _ = ginkgo.Describe("LoadBalancer Service", func() {
	f := framework.NewDefaultFramework("loadbalancer")
	var svc *corev1.Service

	ginkgo.BeforeEach(func() {
		svc = &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-lb-service",
			},
			Spec: corev1.ServiceSpec{
				Type: corev1.ServiceTypeLoadBalancer,
				Ports: []corev1.ServicePort{
					{
						Port:       8080, // non-privileged port
						TargetPort: intstr.FromInt(8080),
						Protocol:   corev1.ProtocolTCP,
					},
				},
				Selector: map[string]string{
					"app": "dummy",
				},
			},
		}
	})

	ginkgo.It("should provision an external IP if LoadBalancer is supported", func(ctx context.Context) {
		ginkgo.By("Creating a LoadBalancer type service")
		createdSvc, err := f.ClientSet.CoreV1().Services(f.Namespace.Name).Create(ctx, svc, metav1.CreateOptions{})
		gomega.Expect(err).NotTo(gomega.HaveOccurred(), "failed to create LoadBalancer service")
		svc = createdSvc

		ginkgo.By("Waiting for external IP allocation")
		gomega.Eventually(func() error {
			updatedSvc, err := f.ClientSet.CoreV1().Services(f.Namespace.Name).Get(ctx, svc.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
			if len(updatedSvc.Status.LoadBalancer.Ingress) == 0 {
				return fmt.Errorf("no external IP allocated")
			}
			ip := updatedSvc.Status.LoadBalancer.Ingress[0].IP
			fmt.Fprintf(ginkgo.GinkgoWriter, "LoadBalancer external IP allocated: %s\n", ip)
			return nil
		}, 2*time.Minute, 5*time.Second).Should(gomega.Succeed(), "expected LoadBalancer external IP to be allocated")
	})

	ginkgo.AfterEach(func(ctx context.Context) {
		if svc != nil {
			_ = f.ClientSet.CoreV1().Services(f.Namespace.Name).Delete(ctx, svc.Name, metav1.DeleteOptions{})
		}
	})
})
