package kube

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

type KubeClient struct {
	ClientSet *kubernetes.Clientset
}

func NewClient(fileName string) (*KubeClient, error) {
	kubeconfig := fmt.Sprintf("%s/.kube/%s", homedir.HomeDir(), fileName)

	config, err := clientcmd.BuildConfigFromFlags("", *&kubeconfig)
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubeClient{client}, err
}

func (kc *KubeClient) GetNodes() ([]string, error) {
	nodes, err := kc.ClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var nodeNames []string
	for _, item := range nodes.Items {
		nodeNames = append(nodeNames, item.Name)
	}
	return nodeNames, nil
}

type Node struct {
	Name string
	Pods []Pod
}

type Pod struct {
	Namespace string
	Name      string
}

func (kc *KubeClient) DescribeNode(name string) (*Node, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + name + ",status.phase!=" + string(corev1.PodSucceeded) + ",status.phase!=" + string(corev1.PodFailed))
	if err != nil {
		return nil, err
	}
	nodeNonTerminatedPodsList, err := kc.ClientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{FieldSelector: fieldSelector.String()})
	if err != nil {
		return nil, err
	}
	node := &Node{Name: name}
	for _, pod := range nodeNonTerminatedPodsList.Items {
		node.Pods = append(node.Pods, Pod{Namespace: pod.Namespace, Name: pod.Name})
	}
	return node, nil
}
