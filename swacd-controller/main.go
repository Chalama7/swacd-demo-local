package main

import (
	"context"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := os.ExpandEnv("${HOME}/kcp-local-demo/kcp/.kcp/admin.kubeconfig")
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("🔍 Checking SWACD resources...")
		listResources(dynamicClient)
		time.Sleep(10 * time.Second)
	}
}

func listResources(client dynamic.Interface) {
	listAndPrint(client, "swacd.io", "v1alpha1", "tenants", "Tenant")
	listAndPrint(client, "swacd.io", "v1alpha1", "edgeroutes", "EdgeRoute")
	listAndPrint(client, "swacd.io", "v1alpha1", "originservices", "OriginService")
}

func listAndPrint(client dynamic.Interface, group, version, resource, kind string) {
	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}
	list, err := client.Resource(gvr).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("❌ Error listing %s: %v\n", kind, err)
		return
	}
	fmt.Printf("✅ Reconciled %d %s objects\n", len(list.Items), kind)
}
