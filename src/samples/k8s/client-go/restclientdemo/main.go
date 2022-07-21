package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	// 参考path : /api/v1/namespaces/{namespace}/pods
	config.APIPath = "api"
	// pod的group是空字符串
	config.GroupVersion = &corev1.SchemeGroupVersion
	// 指定序列化工具
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)

	if err != nil {
		panic(err.Error())
	}

	result := &corev1.PodList{}

	namespace := "python"

	err = restClient.Get().Namespace(namespace).
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("namespace\t status\t\t name\n")

	for _, d := range result.Items {
		fmt.Printf("%v\t %v\t %v\n",
			d.Namespace,
			d.Status.Phase,
			d.Name)
	}
}
