package main

import (
	"context"
	"flag"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/utils/pointer"
	"path/filepath"
)

const (
	NAMESPACE       = "python-jf"
	DEPLOYMENT_NAME = "jf-test-deploy"
	SERVICE_NAME    = "jf-test-service"
)

func main() {
	var kubeconfig *string

	// home是家目录，如果能取得家目录的值，就可以用来做默认值
	if home := homedir.HomeDir(); home != "" {
		// 如果输入了kubeconfig参数，该参数的值就是kubeconfig文件的绝对路径，
		// 如果没有输入kubeconfig参数，就用默认路径~/.kube/config
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		// 如果取不到当前用户的家目录，就没办法设置kubeconfig的默认目录了，只能从入参中取
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	operate := flag.String("opertate", "create", "operate type: create or clean")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("operation is %v\n", *operate)

	if *operate == "clean" {
		clean(clientset)
	} else {
		createNamespace(clientset)

		createDeployment(clientset)

		createService(clientset)
	}
}

func createService(clientset *kubernetes.Clientset) {
	serviceClient := clientset.CoreV1().Services(NAMESPACE)

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: SERVICE_NAME,
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Name:     "http",
					Port:     8080,
					NodePort: 30080,
				},
			},
			Selector: map[string]string{
				"app": "tomcat",
			},
			Type: apiv1.ServiceTypeNodePort,
		},
	}

	result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Create service %s \n", result.GetName())
}

func createDeployment(clientset *kubernetes.Clientset) {
	deploymentClient := clientset.AppsV1().Deployments(NAMESPACE)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: DEPLOYMENT_NAME,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "tomcat",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "tomcat",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            "tomcat",
							Image:           "tomcat:8.0.18-jre8",
							ImagePullPolicy: "IfNotPresent",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Create deployment %s \n", result.GetName())

}

func createNamespace(clientset *kubernetes.Clientset) {
	namespaceClient := clientset.CoreV1().Namespaces()

	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: NAMESPACE,
		},
	}

	result, err := namespaceClient.Create(context.TODO(), namespace, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Create namespace %s \n", result.GetName())
}

func clean(clientset *kubernetes.Clientset) {
	emptyDeleteOptions := metav1.DeleteOptions{}

	if err := clientset.CoreV1().Services(NAMESPACE).Delete(context.TODO(), SERVICE_NAME, emptyDeleteOptions); err != nil {
		panic(err.Error())
	}

	if err := clientset.AppsV1().Deployments(NAMESPACE).Delete(context.TODO(), DEPLOYMENT_NAME, emptyDeleteOptions); err != nil {
		panic(err.Error())
	}

	if err := clientset.CoreV1().Namespaces().Delete(context.TODO(), NAMESPACE, emptyDeleteOptions); err != nil {
		panic(err.Error())
	}
}
