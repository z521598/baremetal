package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/util/workqueue"

	"baremetal/pkg/apis/baremetal"
	v1 "baremetal/pkg/apis/baremetal/v1"
	"baremetal/pkg/client/clientset/versioned"
	informers "baremetal/pkg/client/informers/externalversions"
)

const namespace string = "langshiquan-namespace"

func main() {
	fmt.Println("hello,world")
	//listAndWatchNamespacedPods()
	listAndWatchBareMetalJobs()
}

func listAndWatchBareMetalJobs() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/langshiquan/.kube/config")
	stopCh := make(chan struct{})
	if err != nil {
		panic(err)
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	informerFactory := informers.NewSharedInformerFactoryWithOptions(clientset, time.Minute, informers.WithNamespace(namespace))

	informerFac, err := informerFactory.ForResource(schema.GroupVersionResource{
		Group:    baremetal.GroupName,
		Version:  baremetal.Version,
		Resource: "baremetaljobs",
	})
	if err != nil {
		panic(err)
	}
	lister := informerFac.Lister().ByNamespace(namespace)
	informer := informerFac.Informer()
	wq := workqueue.New()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{

		AddFunc: func(obj interface{}) {
			bmj := obj.(*v1.BareMetalJob)
			fmt.Println(fmt.Sprintf("put %s[add] in workqueue", bmj.Name))
			m := Delta{
				Name:   bmj.Name,
				Action: "add",
			}
			wq.Add(m)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			bmj := newObj.(*v1.BareMetalJob)
			fmt.Println(fmt.Sprintf("put %s[update] in workqueue", bmj.Name))
			m := Delta{
				Name:   bmj.Name,
				Action: "update",
			}
			wq.Add(m)
		},
		DeleteFunc: func(obj interface{}) {
			bmj := obj.(*v1.BareMetalJob)
			fmt.Println(fmt.Sprintf("put %s[delete] in workqueue", bmj.Name))
			m := Delta{
				Name:   bmj.Name,
				Action: "delete",
			}
			wq.Add(m)
		},
	})
	go func() {
		for {
			item, shutdown := wq.Get()
			if shutdown {
				break
			}
			m := item.(Delta)
			fmt.Println(fmt.Sprintf("poll name=%s action=%s from workqueue", m.Name, m.Action))
			bmj, _ := lister.Get(m.Name)
			fmt.Println(fmt.Sprintf("get name=%s from lister contenet as follow: ", m.Name))
			fmt.Println(bmj)

		}

	}()
	informer.Run(stopCh)
}

type Delta struct {
	Action string
	Name   string
}

//func listAndWatchNamespacedPods() {
//	config, err := clientcmd.BuildConfigFromFlags("", "/Users/langshiquan/.kube/config")
//
//	if err != nil {
//		panic(err)
//	}
//
//	clientset, err := kubernetes.NewForConfig(config)
//
//	if err != nil {
//		panic(err)
//	}
//
//	stopCh := make(chan struct{})
//	defer close(stopCh)
//	sharedInformers := informers.NewSharedInformerFactoryWithOptions(clientset, time.Minute, informers.WithNamespace("langshiquan-namespace"))
//
//	informer := sharedInformers.Core().V1().Pods().Informer()
//
//	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
//		AddFunc: func(obj interface{}) {
//			fmt.Println("add")
//			fmt.Println(obj)
//		},
//		UpdateFunc: func(oldObj, newObj interface{}) {
//			fmt.Println("update")
//			fmt.Println(oldObj)
//			fmt.Println(newObj)
//		},
//		DeleteFunc: func(obj interface{}) {
//			fmt.Println("delete")
//			fmt.Println(obj)
//		},
//	})
//	informer.Run(stopCh)
//
//}
