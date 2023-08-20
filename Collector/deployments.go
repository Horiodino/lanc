package collector

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (CLIENT *K8sclient) PullDeploymentInfo() (DeploymentInfoStruct, error) {

	deployments, err := CLIENT.clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	DeploymentInfo := &DeploymentInfoStruct{}

	for _, deployment := range deployments.Items {
		DeploymentInfo.Deployments = append(DeploymentInfo.Deployments, DeploymentStructure{
			Name:                    deployment.Name,
			Namespace:               deployment.Namespace,
			CreationTimestamp:       deployment.CreationTimestamp.Time,
			Replicas:                *deployment.Spec.Replicas,
			AvailableRep:            deployment.Status.AvailableReplicas,
			ReadyRep:                deployment.Status.ReadyReplicas,
			UpToDateRep:             deployment.Status.UpdatedReplicas,
			AvailableUpToDateRep:    deployment.Status.AvailableReplicas,
			Age:                     deployment.Status.Conditions[0].LastTransitionTime.String(),
			Spec:                    deployment.Spec.Template.Spec.Containers[0].Image,
			Labels:                  deployment.GetLabels()[deployment.Name],
			Selector:                deployment.Spec.Selector.MatchLabels[deployment.Name],
			Strategy:                string(deployment.Spec.Strategy.Type),
			MinReadySeconds:         deployment.Spec.MinReadySeconds,
			RevisionHistoryLimit:    *deployment.Spec.RevisionHistoryLimit,
			Paused:                  deployment.Spec.Paused,
			ProgressDeadlineSeconds: *deployment.Spec.ProgressDeadlineSeconds,
			ReplicaSet:              deployment.Status.Replicas,
			Conditions:              string(deployment.Status.Conditions[0].Type),
		})
	}
	return *DeploymentInfo, nil
}
