package api

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (CLIENT *K8sclient) PullPodInfo() (PodInfoStruct, error) {
	Pods, err := CLIENT.clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	PodInfo := &PodInfoStruct{}
	for i := 0; i < len(Pods.Items); i++ {
		PodInfo.Name = append(PodInfo.Name, Pods.Items[i].Name)
		PodInfo.Phase = append(PodInfo.Phase, string(Pods.Items[i].Status.Phase))
		PodInfo.StartTime = append(PodInfo.StartTime, Pods.Items[i].Status.StartTime.Format("2006-01-02 15:04:05"))
		PodInfo.PodIP = append(PodInfo.PodIP, Pods.Items[i].Status.PodIP)
		PodInfo.HostIP = append(PodInfo.HostIP, Pods.Items[i].Status.HostIP)
		PodInfo.QOSClass = append(PodInfo.QOSClass, string(Pods.Items[i].Status.QOSClass))
		PodInfo.Message = append(PodInfo.Message, Pods.Items[i].Status.Message)
		PodInfo.Reason = append(PodInfo.Reason, Pods.Items[i].Status.Reason)
		// PodInfo.ContainerStatuses = append(PodInfo.ContainerStatuses, Pods.Items[i].Status.ContainerStatuses)
		// PodInfo.InitContainerStatuses = append(PodInfo.InitContainerStatuses, Pods.Items[i].Status.InitContainerStatuses)
		PodInfo.NominatedNodeName = append(PodInfo.NominatedNodeName, Pods.Items[i].Status.NominatedNodeName)
		// PodInfo.Conditions = append(PodInfo.Conditions, Pods.Items[i].Status.Conditions)
		PodInfo.EphemeralContainerStatuses = append(PodInfo.EphemeralContainerStatuses, Pods.Items[i].Status.NominatedNodeName)
	}

	return *PodInfo, nil
}
