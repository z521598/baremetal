package main

import (
	"fmt"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println("hello,world")
	listAndWatchNamespacedPods()
}

func listAndWatchNamespacedPods() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/langshiquan/.kube/config")

	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	sharedInformers := informers.NewSharedInformerFactoryWithOptions(clientset, time.Minute, informers.WithNamespace("langshiquan-namespace"))

	informer := sharedInformers.Core().V1().Pods().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("add")
			fmt.Println(obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("update")
			fmt.Println(oldObj)
			fmt.Println(newObj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete")
			fmt.Println(obj)
		},
	})
	informer.Run(stopCh)

}
