package collector

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

	for _, pod := range Pods.Items {

		PodInfo.PodName = append(PodInfo.PodName, pod.Name)
		PodInfo.PodCreation = append(PodInfo.PodCreation, pod.CreationTimestamp.Time)
		PodInfo.PodDeletionGracePeriodSeconds = append(PodInfo.PodDeletionGracePeriodSeconds, pod.DeletionGracePeriodSeconds)
		PodInfo.PodPhase = append(PodInfo.PodPhase, pod.Status.Phase)
		PodInfo.PodRunningOn = append(PodInfo.PodRunningOn, pod.Spec.NodeName)
		PodInfo.PodLabels = pod.GetLabels()
		PodInfo.PodNamespace = append(PodInfo.PodNamespace, pod.GetNamespace())
		PodInfo.PodAnnotation = pod.ObjectMeta.Annotations
		// PodInfo.PodOwnerReferences = append(PodInfo.PodOwnerReferences, pod.ObjectMeta.GetObjectMeta().GetOwnerReferences())
		PodInfo.PodResourceVersion = append(PodInfo.PodResourceVersion, pod.ObjectMeta.GetObjectMeta().GetResourceVersion())
		PodInfo.PodUID = append(PodInfo.PodUID, pod.ObjectMeta.GetObjectMeta().GetUID())
		PodInfo.PodDNSConfig = append(PodInfo.PodDNSConfig, pod.Spec.DNSConfig)
		PodInfo.PodDNSPolicy = append(PodInfo.PodDNSPolicy, pod.Spec.DNSPolicy)
		PodInfo.PodEnableServiceLinks = append(PodInfo.PodEnableServiceLinks, pod.Spec.EnableServiceLinks)
		// PodInfo.PodEphemeralContainers = append(PodInfo.PodEphemeralContainers, pod.Spec.EphemeralContainers)
		PodInfo.PodHostIpc = append(PodInfo.PodHostIpc, pod.Spec.HostIPC)
		PodInfo.PodHostNetwork = append(PodInfo.PodHostNetwork, pod.Spec.HostNetwork)
		PodInfo.PodHostPID = append(PodInfo.PodHostPID, pod.Spec.HostPID)
		PodInfo.PodHostUsers = append(PodInfo.PodHostUsers, pod.Spec.HostUsers)
		// PodInfo.PodImagePullSecrets = append(PodInfo.PodImagePullSecrets, pod.Spec.ImagePullSecrets)

		// PodInfo.PodInitContainers = append(PodInfo.PodInitContainers, pod.Spec.InitContainers)
		PodInfo.PodRestartPolicy = append(PodInfo.PodRestartPolicy, pod.Spec.RestartPolicy)
		PodInfo.PodRuntimeClassName = append(PodInfo.PodRuntimeClassName, pod.Spec.RuntimeClassName)

	}

	return *PodInfo, nil
}
