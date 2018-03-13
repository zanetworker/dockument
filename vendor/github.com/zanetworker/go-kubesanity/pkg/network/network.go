package network

import (
	"fmt"
	"os"

	"github.com/zanetworker/go-kubesanity/pkg/kubesanityutils"
	"github.com/zanetworker/go-kubesanity/pkg/log"
	typev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

//KubernetesClient is the client object we will use to connect to our cluster
type KubernetesClient struct {
	Client v1.CoreV1Interface
}

//NewKubeClient initializes the kubernetes client
func NewKubeClient() *KubernetesClient {

	kubeconfigPath, ok := os.LookupEnv("KUBECONFIG_PATH")
	if !ok {
		log.FatalS("KUBECONFIG_PATH was not set")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	kubeclient := clientset.CoreV1()

	return &KubernetesClient{kubeclient}
}

//CheckDuplicatePodIP checks if two pods have the same IP in all namespaces
func (kc *KubernetesClient) CheckDuplicatePodIP() (bool, error) {

	podIPs := make(map[string]typev1.Pod)
	podList, err := kc.Client.Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Error(err.Error())
	}

	for _, pod := range podList.Items {
		otherPod, ipExists := podIPs[pod.Status.PodIP]
		if ipExists {
			return true, kubesanityutils.NewDuplicaPodIPError(fmt.Errorf("Duplicate Pod IP Address: %s/%s -> %s and %s/%s -> %s",
				pod.ObjectMeta.Namespace,
				pod.ObjectMeta.Name,
				pod.Status.PodIP,
				otherPod.ObjectMeta.Namespace,
				otherPod.ObjectMeta.Name,
				otherPod.Status.PodIP,
			).Error())
		}
		podIPs[pod.Status.PodIP] = pod
	}
	log.Info("No duplicate pod IPs found!")
	return false, nil
}

//CheckDuplicateServiceIP checks if two services have the same IP in all namespaces
func (kc *KubernetesClient) CheckDuplicateServiceIP() (bool, error) {

	serviceIPs := make(map[string]typev1.Service)
	serviceList, err := kc.Client.Services("").List(metav1.ListOptions{})

	if err != nil {
		log.Error(err.Error())
	}

	for _, service := range serviceList.Items {
		otherService, ipExists := serviceIPs[service.Spec.ClusterIP]
		if ipExists {
			return true, fmt.Errorf("Duplicate Service IP Address: %s/%s -> %s and %s/%s -> %s",
				service.ObjectMeta.Namespace,
				service.ObjectMeta.Name,
				service.Spec.ClusterIP,
				otherService.ObjectMeta.Namespace,
				otherService.ObjectMeta.Name,
				otherService.Spec.ClusterIP,
			)
		}
		serviceIPs[service.Spec.ClusterIP] = service
	}
	log.Info("No duplicate service IPs found!")
	return false, nil
}
