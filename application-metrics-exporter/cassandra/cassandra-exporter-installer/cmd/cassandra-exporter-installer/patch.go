package main

import (
	"log"

	"github.com/mayadata.io/monitoring-tools/application-metrics-exporter/cassandra/cassandra-exporter-installer/pkg/kubernetes"
)

func main() {
	// Create rest config
	restConfig, err := kubernetes.NewRestConfig()
	if err != nil {
		log.Fatal(err)
	}
	kubernetes.RestConfig = restConfig

	// Create k8s clientset
	k8sClientSet, err := kubernetes.NewClientSet(kubernetes.RestConfig)
	if err != nil {
		log.Fatal(err)
	}
	kubernetes.ClientSet = k8sClientSet

	envMap := kubernetes.Getenv()

	//patch the resource
	err = kubernetes.StatefulsetPatch(envMap)
	if err != nil {
		log.Fatal(err)
	}

}
