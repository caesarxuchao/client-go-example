package main

import (
	"flag"
	"time"

	"github.com/golang/glog"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/rest"
)

func main() {
	flag.Parse()
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//	eventBroadcaster := record.NewBroadcaster()
	//	eventBroadcaster.StartLogging(glog.Infof)
	//	eventBroadcaster.StartRecordingToSink(clientset.Core().Events(""))
	//	recorder := eventBroadcaster.NewRecorder(v1.EventSource{Component: "client-go-example"})
	for {
		pods, err := clientset.Core().Pods("").List(api.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		glog.Infof("len(pods)=%d", len(pods.Items))
		// recorder.Eventf(nil, v1.EventTypeNormal, "sample event", "there are %d pods", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
