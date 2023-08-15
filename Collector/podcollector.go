package collector

import (
	"context"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (CLIENT *K8sclient) PullPodInfo() (PodInfoStruct, error) {
	Pods, err := CLIENT.clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	PodInfo := &PodInfoStruct{}

	PodInfo.Timestamp = time.Now()
	for _, pod := range Pods.Items {

		PodInfo.PODS = append(PodInfo.PODS, PodStruct{
			PodName:                       pod.Name,
			PodCreation:                   pod.CreationTimestamp.Time,
			PodDeletionGracePeriodSeconds: pod.DeletionGracePeriodSeconds,
			PodPhase:                      pod.Status.Phase,
			PodRunningOn:                  pod.Spec.NodeName,
			PodLabels:                     pod.Labels,
			PodNamespace:                  pod.Namespace,
			PodAnnotation:                 pod.Annotations,
			PodResourceVersion:            pod.ResourceVersion,
			PodUID:                        pod.UID,
			PodDNSConfig:                  pod.Spec.DNSConfig,
			PodDNSPolicy:                  pod.Spec.DNSPolicy,
			PodEnableServiceLinks:         pod.Spec.EnableServiceLinks,
			PodHostIpc:                    pod.Spec.HostIPC,
			PodHostNetwork:                pod.Spec.HostNetwork,
			PodHostPID:                    pod.Spec.HostPID,
			PodRestartPolicy:              pod.Spec.RestartPolicy,
			PodRuntimeClassName:           pod.Spec.RuntimeClassName,
		})
	}

	return *PodInfo, nil
}
