package collector

import (
	"context"
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Pullinfo interface {
	PullClusterInfo(CLIENT *K8sclient) (string, error)
	PullNodeInfo(CLIENT *K8sclient) (string, error)
	PullPodInfo(CLIENT *K8sclient) (string, error)
	PullNamespaceInfo(CLIENT *K8sclient) (string, error)
	PullDeploymentInfo(CLIENT *K8sclient) (string, error)
	PullServiceInfo(CLIENT *K8sclient) (string, error)
	PullIngressInfo(CLIENT *K8sclient) (string, error)
	PullConfigMapInfo(CLIENT *K8sclient) (string, error)
	PullSecretInfo(CLIENT *K8sclient) (string, error)
	PullPersistentVolumeInfo(CLIENT *K8sclient) (string, error)
	PullPersistentVolumeClaimInfo(CLIENT *K8sclient) (string, error)
	PullStorageClassInfo(CLIENT *K8sclient) (string, error)
	PullRoleInfo(CLIENT *K8sclient) (string, error)
	PullRoleBindingInfo(CLIENT *K8sclient) (string, error)
	PullClusterRoleInfo(CLIENT *K8sclient) (string, error)
	PullClusterRoleBindingInfo(CLIENT *K8sclient) (string, error)
	PullServiceAccountInfo(CLIENT *K8sclient) (string, error)
	PullJobInfo(CLIENT *K8sclient) (string, error)
	PullCronJobInfo(CLIENT *K8sclient) (string, error)
	PullStatefulSetInfo(CLIENT *K8sclient) (string, error)
	PullDaemonSetInfo(CLIENT *K8sclient) (string, error)
	PullReplicaSetInfo(CLIENT *K8sclient) (string, error)
	PullReplicationControllerInfo(CLIENT *K8sclient) (string, error)
	PullHorizontalPodAutoscalerInfo(CLIENT *K8sclient) (string, error)
	PullPodDisruptionBudPullInfo(CLIENT *K8sclient) (string, error)
	PullNetworkPolicyInfo(CLIENT *K8sclient) (string, error)
	PullPodSecurityPolicyInfo(CLIENT *K8sclient) (string, error)
	PullLimitRangeInfo(CLIENT *K8sclient) (string, error)
	PullResourceQuotaInfo(CLIENT *K8sclient) (string, error)
	PullEndpointsInfo(CLIENT *K8sclient) (string, error)
	PullEventInfo(CLIENT *K8sclient) (string, error)
	PullComponentStatusInfo(CLIENT *K8sclient) (string, error)
	PullCustomResourceDefinitionInfo(CLIENT *K8sclient) (string, error)
	PullMutatingWebhookConfigurationInfo(CLIENT *K8sclient) (string, error)
	PullValidatingWebhookConfigurationInfo(CLIENT *K8sclient) (string, error)
	PullPodTemplateInfo(CLIENT *K8sclient) (string, error)
	PullPodSecurityPolicyTemplateInfo(CLIENT *K8sclient) (string, error)
	PullLeaseInfo(CLIENT *K8sclient) (string, error)
	PullPriorityClassInfo(CLIENT *K8sclient) (string, error)
	PullRuntimeClassInfo(CLIENT *K8sclient) (string, error)
	PullCertificateSigningRequestInfo(CLIENT *K8sclient) (string, error)
	PullTokenReviewInfo(CLIENT *K8sclient) (string, error)
	PullSelfSubjectAccessReviewInfo(CLIENT *K8sclient) (string, error)
	PullSelfSubjectRulesReviewInfo(CLIENT *K8sclient) (string, error)
	PullSubjectAccessReviewInfo(CLIENT *K8sclient) (string, error)
	PullHorizontalPodAutoscalerInfoV2Beta1(CLIENT *K8sclient) (string, error)
	PullHorizontalPodAutoscalerInfoV1(CLIENT *K8sclient) (string, error)
	PullPodDisruptionBudgetInfoV1Beta1(CLIENT *K8sclient) (string, error)
}

type NICINFO interface {
	Nicinfo(context context.Context) (string, error)
	GetEndpoints(context context.Context) (string, error)
	GetNodeIP(context context.Context) (string, error)
	AcceptRequest(context context.Context) (string, error)
}

type K8sclient struct {
	clientset *kubernetes.Clientset
	config    *rest.Config
}

var KCLIENT *K8sclient

func NewClient() K8sclient {
	KCLIENT = &K8sclient{}
	os.Getenv("OS")
	if os.Getenv("OS") == "Windows_NT" {
		os.Setenv("KUBECONFIG", os.Getenv("USERPROFILE")+"/.kube/config")
	}

	k8sconfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(k8sconfig)
	if err != nil {
		log.Fatal(err)
	}

	KCLIENT.clientset = clientset
	KCLIENT.config = k8sconfig
	KCLIENT = &K8sclient{
		clientset: KCLIENT.clientset,
		config:    KCLIENT.config,
	}
	return *KCLIENT
}
