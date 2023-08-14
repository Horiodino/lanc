package collector

import (
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type GlobalDATA struct {
	ClusterInfo    ClusterInfoStruct
	PodInfo        PodInfoStruct
	DeploymentInfo DeploymentInfoStruct
}
type DataBase struct {
	clustermetrices ClusterInfoStruct
}

var ClusterInfo *ClusterInfoStruct

type ClusterInfoStruct struct {
	ClusterName string
	Cpu         float64
	Cores       int64
	Nodes       int64
	Totalmemory float64
	Usedmemory  float64
	Disk        float64
	Totaldisk   float64
	Billing     float64
}

var NodeInfoList []NodeInfoStruct

type NodeInfoStruct struct {
	Name    []string
	Memory  []float64
	CPU     []float64
	Disk    []float64
	CpuTemp []float64
	IP      []string
}

type PodInfoStruct struct {
	PodName                       []string
	PodCreation                   []time.Time
	PodDeletionGracePeriodSeconds []*int64
	PodPhase                      []v1.PodPhase
	PodRunningOn                  []string
	PodLabels                     map[string]string
	PodNamespace                  []string
	PodAnnotation                 map[string]string

	PodOwnerReferences     []string
	PodResourceVersion     []string
	PodUID                 []types.UID
	PodDNSConfig           []*v1.PodDNSConfig
	PodDNSPolicy           []v1.DNSPolicy
	PodEnableServiceLinks  []*bool
	PodEphemeralContainers []v1.EphemeralContainer
	PodHostIpc             []bool
	PodHostNetwork         []bool
	PodHostPID             []bool
	PodHostUsers           []*bool
	PodImagePullSecrets    []string
	PodInitContainers      []string
	PodRestartPolicy       []v1.RestartPolicy
	PodRuntimeClassName    []*string
}

type DeploymentInfoStruct struct {
	Name                    []string
	Namespace               []string
	CreationTimestamp       []string
	Replicas                []int32
	AvailableRep            []int32
	ReadyRep                []int32
	UpToDateRep             []int32
	AvailableUpToDateRep    []int32
	Age                     []string
	Spec                    []string
	Labels                  []string
	Selector                []string
	Strategy                []string
	MinReadySeconds         []int32
	RevisionHistoryLimit    []int32
	Paused                  []bool
	ProgressDeadlineSeconds []int32
	ReplicaSet              []string
	Conditions              []string
}

type StatefulSetInfoStruct struct {
	Apiver []string
	Name   []string
	// Selector         []string
	// MatchLabels []string
	ServiceName []string
	Replicas    []int32
	// ReadyRep         []int32
	CurrentRep       []int32
	Age              []string
	MinReadySec      []int32
	Label            []string
	PodName          []string
	PodImage         []string
	VolmountName     []string
	AccessModes      []string
	StorageClassName []string
	Request          []string
	Storage          []string
}

type DaemonSetInfoStruct struct {
	Apiver                        []string
	Name                          []string
	Namespace                     []string
	Label                         []string
	MatchLabels                   []string
	Key                           []string
	Oprator                       []string
	Effect                        []string
	ContainersName                []string
	ContainersImage               []string
	TerminationGracePeriodSeconds []int64
	VolumeMountsName              []string
	VolumeMountsMountPath         []string
	CreationTimestamp             []string
}

type ReplicaSetInfoStruct struct {
}

type ReplicationControllerInfoStruct struct {
}

type HorizontalPodAutoscalerInfoStruct struct {
}

type PodDisruptionBudgetInfoStruct struct {
}

type NetworkPolicyInfoStruct struct {
}

type PodSecurityPolicyInfoStruct struct {
}

type LimitRangeInfoStruct struct {
}
