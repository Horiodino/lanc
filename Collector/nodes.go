package collector

import (
	"context"
	"log"

	// Kubernetes API client libraries and packages
	// ".mongodb.org/mongo-driver/mongo"go

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

func (CLIENT *K8sclient) PullNodeInfo() (NodeInfoStruct, error) {

	nodes, err := CLIENT.clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	info, err := versioned.NewForConfig(CLIENT.config)
	if err != nil {
		log.Fatal(err)
	}

	NodeInfo := &NodeInfoStruct{}

	for _, node := range nodes.Items {
		metrics, err := info.MetricsV1beta1().NodeMetricses().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		var cpuUsage int64
		for _, node := range metrics.Items {
			cpuUsage += (node.Usage.Cpu().MilliValue())
		}
		nodes, err := KCLIENT.clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		var cpuCores int64
		for _, node := range nodes.Items {
			cpuCores += (node.Status.Capacity.Cpu().Value())
		}
		var memoryUsage int64
		for _, node := range metrics.Items {
			memoryUsage += (node.Usage.Memory().Value()) / 1000000
		}
		var memory int64
		for _, node := range nodes.Items {
			memory += (node.Status.Capacity.Memory().Value()) / 1000000
		}
		var diskUsage int64
		for _, node := range metrics.Items {
			diskUsage += (node.Usage.StorageEphemeral().Value()) / 1000000
		}
		var disk int64
		for _, node := range nodes.Items {
			disk += (node.Status.Capacity.StorageEphemeral().Value()) / 1000000
		}
		var ip string
		for _, node := range nodes.Items {
			ip += node.Status.Addresses[0].Address + " "
		}

		NodeInfo.Nodes = append(NodeInfo.Nodes, NodeStruct{
			Name:        node.Name,
			CPUUses:     cpuUsage,
			TotalCPU:    cpuCores,
			MemoryUses:  memoryUsage,
			TotalMemory: memory,
			Disk:        diskUsage,
			TotalDisk:   disk,
			IP:          ip,
		})
	}

	return *NodeInfo, nil
}
