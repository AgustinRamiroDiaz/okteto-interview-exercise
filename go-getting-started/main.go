package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	fmt.Println("Starting hello-world server...")
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/pods", podsInCurrentNamespace)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Oktetos!")
}

// Made with help from https://github.com/kubernetes/client-go/blob/v0.23.1/examples/in-cluster-client-configuration/main.go
func podsInCurrentNamespace(w http.ResponseWriter, r *http.Request) {
	// Prefetch the kubernetes config and create a client from it so we can talk to the API
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := os.Getenv("K8S_NAMESPACE")

	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprint(w, "There are ", len(pods.Items), " pods in the cluster\n")

	for podIndex := range pods.Items {
		pod := pods.Items[podIndex]
		fmt.Fprint(w, "Pod: ", pod.Name, "\n")
	}
}
