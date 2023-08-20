package collector

import (
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type GlobalDATA struct {
	Timestamp   time.Time
	Clsuterinfo ClusterInfoStruct
	Nodeinfo    NodeInfoStruct
	PodInfo     PodInfoStruct
	DeployInfo  DeploymentInfoStruct
}

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

type NodeInfoStruct struct {
	Timestamp time.Time
	Nodes     []NodeStruct
}

type NodeStruct struct {
	Name        string
	MemoryUses  int64
	TotalMemory int64
	CPUUses     int64
	TotalCPU    int64
	Disk        int64
	TotalDisk   int64
	IP          string
}

type PodInfoStruct struct {
	Timestamp time.Time
	PODS      []PodStruct
}

type DeploymentInfoStruct struct {
	Timestamp   time.Time
	Deployments []DeploymentStructure
}

type PodStruct struct {
	PodName                       string
	PodCreation                   time.Time
	PodDeletionGracePeriodSeconds *int64
	PodPhase                      v1.PodPhase
	PodRunningOn                  string
	PodLabels                     map[string]string
	PodNamespace                  string
	PodAnnotation                 map[string]string

	PodOwnerReferences     string
	PodResourceVersion     string
	PodUID                 types.UID
	PodDNSConfig           *v1.PodDNSConfig
	PodDNSPolicy           v1.DNSPolicy
	PodEnableServiceLinks  *bool
	PodEphemeralContainers v1.EphemeralContainer
	PodHostIpc             bool
	PodHostNetwork         bool
	PodHostPID             bool
	PodHostUsers           *bool
	PodImagePullSecrets    string
	PodInitContainers      string
	PodRestartPolicy       v1.RestartPolicy
	PodRuntimeClassName    *string
}

type DeploymentStructure struct {
	Name                    string
	Namespace               string
	CreationTimestamp       time.Time
	Replicas                int32
	AvailableRep            int32
	ReadyRep                int32
	UpToDateRep             int32
	AvailableUpToDateRep    int32
	Age                     string
	Spec                    string
	Labels                  string
	Selector                string
	Strategy                string
	MinReadySeconds         int32
	RevisionHistoryLimit    int32
	Paused                  bool
	ProgressDeadlineSeconds int32
	ReplicaSet              int32
	Conditions              string
}
