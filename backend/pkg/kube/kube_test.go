package kube

import "testing"

func TestGetNodes(t *testing.T) {
	kc, err := NewClient("kind_config")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	nodes, err := kc.GetNodes()
	if err != nil {
		t.Fatalf("GetNodes: %v", err)
	}

	for _, node := range nodes {
		_, err = kc.DescribeNode(node)
		if err != nil {
			t.Fatalf("DescribeNode: %v", err)
		}
	}
}
