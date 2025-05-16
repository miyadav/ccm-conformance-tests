package framework

import (
    "flag"
    "os"
    "path/filepath"
    "context"
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
    F *Framework
)

func init() {
    flag.StringVar(&KubeConfig, "kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "path to kubeconfig")
}

func SetupFramework() {
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
                GenerateName: "ccm-test-",
            },
        },
        metav1.CreateOptions{},
    )
    if err != nil {
        panic(err)
    }

    F = &Framework{
        ClientSet: clientset,
        Namespace: ns,
    }
}

func TeardownFramework() {
    if F != nil && F.Namespace != nil {
        _ = F.ClientSet.CoreV1().Namespaces().Delete(context.TODO(), F.Namespace.Name, metav1.DeleteOptions{})
    }
}

