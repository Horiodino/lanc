package api

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
	Name                       []string
	Phase                      []string
	StartTime                  []string
	PodIP                      []string
	HostIP                     []string
	QOSClass                   []string
	Message                    []string
	Reason                     []string
	ContainerStatuses          []string
	InitContainerStatuses      []string
	NominatedNodeName          []string
	Conditions                 []string
	EphemeralContainerStatuses []string
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
