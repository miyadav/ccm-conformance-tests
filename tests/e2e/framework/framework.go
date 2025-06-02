package framework

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Framework struct {
	ClientSet *kubernetes.Clientset
	Namespace *corev1.Namespace
}

var (
	KubeConfig string
	F          *Framework
)

func init() {
	flag.StringVar(&KubeConfig, "kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "path to kubeconfig")
}

// NewDefaultFramework is used for Ginkgo-style test isolation
func NewDefaultFramework(name string) *Framework {
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", KubeConfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	ns, err := clientset.CoreV1().Namespaces().Create(
		context.TODO(),
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: name + "-",
			},
		},
		metav1.CreateOptions{},
	)
	if err != nil {
		panic(err)
	}

	return &Framework{
		ClientSet: clientset,
		Namespace: ns,
	}
}
