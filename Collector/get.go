package collector

import (
	"fmt"
	"log"

	"k8s.io/metrics/pkg/client/clientset/versioned"

	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (CLIENT *K8sclient) PullClusterInfo() (ClusterInfoStruct, error) {
	ClusterInfo := &ClusterInfoStruct{}

	MetricresClient, err := versioned.NewForConfig(CLIENT.config)
	if err != nil {
		fmt.Println(err)
	}

	metrics, err := MetricresClient.MetricsV1beta1().NodeMetricses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range metrics.Items {
		ClusterInfo.Cpu += float64(node.Usage.Cpu().MilliValue())
		ClusterInfo.Usedmemory += float64((node.Usage.Memory().Value()) / 1000000)
		ClusterInfo.Disk += float64(node.Usage.StorageEphemeral().Value() / 1000000)
	}

	nodes, err := CLIENT.clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes.Items {
		ClusterInfo.Nodes += (node.Status.Capacity.Cpu().Value())
		ClusterInfo.Cores += (node.Status.Capacity.Cpu().Value())
		ClusterInfo.Totaldisk += float64(((node.Status.Capacity.StorageEphemeral().Value()) / 1000000))
		ClusterInfo.Totalmemory += float64((node.Status.Capacity.Memory().Value()) / 1000000)
	}

	ClusterInfo = &ClusterInfoStruct{

		Cpu:         ClusterInfo.Cpu,
		Usedmemory:  ClusterInfo.Usedmemory,
		Disk:        ClusterInfo.Disk,
		Nodes:       ClusterInfo.Nodes,
		Cores:       ClusterInfo.Cores,
		Totaldisk:   ClusterInfo.Totaldisk,
		Totalmemory: ClusterInfo.Totalmemory,
	}

	return *ClusterInfo, nil
}

//func (CLIENT *K8sclient) PullNodeInfo() (string, error) {
//
//	return "NodeInfo", nil
//}

/*
func (CLIENT *K8sclient) PullDeploymentInfo() (DeploymentInfoStruct, error) {

	Deployments, err := CLIENT.clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	Deploy := &DeploymentInfoStruct{}
	for i := 0; i < len(Deployments.Items); i++ {
		Deploy.Name = append(Deploy.Name, Deployments.Items[i].Name)
		Deploy.Namespace = append(Deploy.Namespace, Deployments.Items[i].Namespace)
		Deploy.CreationTimestamp = append(Deploy.CreationTimestamp, Deployments.Items[i].CreationTimestamp.Format("2006-01-02 15:04:05"))
		// #TODO ADD MORE INFO
	}

	return *Deploy, nil
}
*/

/*
func (CLIENT *K8sclient) PullStatefulSetInfo() (StatefulSetInfoStruct, error) {

	ST, err := CLIENT.clientset.AppsV1().StatefulSets("").List(context.Background(), metav1.ListOptions{})
	if err != nil { //vscode-file://vscode-app/c:/Users/pkkha/AppData/Local/Programs/Microsoft%20VS%20Code/resources/app/out/vs/code/electron-sandbox/workbench/workbench.html
		log.Fatal(err)
	}
	Statefulsets := &StatefulSetInfoStruct{}

	for i := 0; i < len(ST.Items); i++ {
		Statefulsets.Apiver = append(Statefulsets.Apiver, ST.Items[i].APIVersion)
		Statefulsets.Name = append(Statefulsets.Name, ST.Items[i].Name)
		// Statefulsets.MatchLabels = append(Statefulsets.MatchLabels, ST.Items[i].Spec.Selector.MatchLabels)
		Statefulsets.ServiceName = append(Statefulsets.ServiceName, ST.Items[i].Spec.ServiceName)
		Statefulsets.Replicas = append(Statefulsets.Replicas, *ST.Items[i].Spec.Replicas)
		// Statefulsets.ReadyRep = append(Statefulsets.ReadyRep, ST.Items[i].Spec.)
		Statefulsets.Age = append(Statefulsets.Age, ST.Items[i].CreationTimestamp.Format("2006-01-02 15:04:05"))
		Statefulsets.MinReadySec = append(Statefulsets.MinReadySec, ST.Items[i].Spec.MinReadySeconds)
		// Statefulsets.Label = append(Statefulsets.Label, ST.Items[i].Spec.Template.GetLabels()[i])
		Statefulsets.PodName = append(Statefulsets.PodName, ST.Items[i].Spec.Template.Name)
		Statefulsets.PodImage = append(Statefulsets.PodImage, ST.Items[i].Spec.Template.Spec.Containers[0].Image, ST.Items[i].Spec.Template.Spec.Containers[1].Image)
		Statefulsets.VolmountName = append(Statefulsets.VolmountName, ST.Items[i].Spec.Template.Spec.Containers[i].VolumeMounts[i].Name)
		Statefulsets.AccessModes = append(Statefulsets.AccessModes, string(ST.Items[i].Spec.VolumeClaimTemplates[i].Spec.AccessModes[i]))
		Statefulsets.StorageClassName = append(Statefulsets.StorageClassName, string(*ST.Items[i].Spec.VolumeClaimTemplates[i].Spec.StorageClassName))
		Statefulsets.Request = append(Statefulsets.Request, ST.Items[i].Spec.VolumeClaimTemplates[i].Spec.Resources.Requests.Storage().String())
		Statefulsets.Storage = append(Statefulsets.Storage, ST.Items[i].Spec.VolumeClaimTemplates[i].Spec.Resources.Requests.Storage().String())
	}

	return *Statefulsets, nil
}
*/

/*
func (CLIENT *K8sclient) PullDaemonSetInfo() (DeploymentInfoStruct, error) {

	DS, err := CLIENT.clientset.AppsV1().DaemonSets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	Daemonset := &DaemonSetInfoStruct{}

	for i := 0; i < len(DS.Items); i++ {
		Daemonset.Apiver = append(Daemonset.Apiver, DS.Items[i].APIVersion)
		Daemonset.Name = append(Daemonset.Name, DS.Items[i].Name)
		Daemonset.Namespace = append(Daemonset.Namespace, DS.Items[i].Namespace)
		// Daemonset.Label = append(Daemonset.Label, DS.Items[i].Spec.Template.GetLabels()[i])
		Daemonset.Key = append(Daemonset.Key, DS.Items[i].Spec.Template.Spec.Containers[i].Env[i].Name)
		Daemonset.Oprator = append(Daemonset.Oprator, DS.Items[i].Spec.Template.Spec.Containers[i].Env[i].Value)
		Daemonset.Effect = append(Daemonset.Effect, DS.Items[i].Spec.Template.Spec.Containers[i].Env[i].ValueFrom.FieldRef.FieldPath)
		Daemonset.ContainersName = append(Daemonset.ContainersName, DS.Items[i].Spec.Template.Spec.Containers[i].Name)
		Daemonset.ContainersImage = append(Daemonset.ContainersImage, DS.Items[i].Spec.Template.Spec.Containers[i].Image)
		Daemonset.TerminationGracePeriodSeconds = append(Daemonset.TerminationGracePeriodSeconds, *DS.Items[i].Spec.Template.Spec.TerminationGracePeriodSeconds)
		Daemonset.VolumeMountsName = append(Daemonset.VolumeMountsName, DS.Items[i].Spec.Template.Spec.Containers[i].VolumeMounts[i].Name)
		Daemonset.VolumeMountsMountPath = append(Daemonset.VolumeMountsMountPath, DS.Items[i].Spec.Template.Spec.Containers[i].VolumeMounts[i].MountPath)
		Daemonset.CreationTimestamp = append(Daemonset.CreationTimestamp, DS.Items[i].CreationTimestamp.Format("2006-01-02 15:04:05"))
	}

	return DeploymentInfoStruct{}, nil
}
*/
//func (CLIENT *K8sclient) PullReplicaSetInfo() (string, error) {
//	return "ReplicaSetInfo", nil
//}
//
//func (CLIENT *K8sclient) PullReplicationControllerInfo() (string, error) {
//	return "ReplicationControllerInfo", nil
//}
//
//func (CLIENT *K8sclient) PullHorizontalPodAutoscalerInfo() (string, error) {
//	return "HorizontalPodAutoscalerInfo", nil
//}
//
//func (CLIENT *K8sclient) PullPodDisruptionBudgetInfo() (string, error) {
//	return "PodDisruptionBudgetInfo", nil
//}
//
//func (CLIENT *K8sclient) PullNetworkPolicyInfo() (string, error) {
//	return "NetworkPolicyInfo", nil
//}
//
//func (CLIENT *K8sclient) PullPodSecurityPolicyInfo() (string, error) {
//	return "PodSecurityPolicyInfo", nil
//}
//
//func (CLIENT *K8sclient) PullLimitRangeInfo() (string, error) {
//	return "LimitRangeInfo", nil
//}
