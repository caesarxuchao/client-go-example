package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
)

func main() {
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		glog.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatal(err)
	}
	//	eventBroadcaster := record.NewBroadcaster()
	//	eventBroadcaster.StartLogging(glog.Infof)
	//	eventBroadcaster.StartRecordingToSink(clientset.Core().Events(""))
	//	recorder := eventBroadcaster.NewRecorder(v1.EventSource{Component: "client-go-example"})
	for {
		pods, err := clientset.Core().Pods("").List(api.ListOptions{})
		if err != nil {
			glog.Fatal(err)
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		// recorder.Eventf(nil, v1.EventTypeNormal, "sample event", "there are %d pods", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
