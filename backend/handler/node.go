package handler

import (
	"context"
	"fmt"

	"github.com/hisamouna/node_describe_dashboard/pkg/kube"
	"github.com/hisamouna/node_describe_dashboard/pkg/server/node"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NodeHandler struct {
	node.UnimplementedNodeServiceServer
}

func (handler *NodeHandler) DescribeNodes(ctx context.Context, in *emptypb.Empty) (*node.NodesResponse, error) {
	kc, err := kube.NewClient("kind_config")
	if err != nil {
		return nil, fmt.Errorf("DescribeNode(kueb.NewClient): %v", err)
	}
	nodes, err := kc.GetNodes()
	if err != nil {
		return nil, fmt.Errorf("DescribeNode(kueb.GetNodes): %v", err)
	}

	var response = &node.NodesResponse{}
	for _, nd := range nodes {
		dn, err := kc.DescribeNode(nd)
		if err != nil {
			return nil, fmt.Errorf("Describe(kube.DescribeNode): %v", err)
		}
		var np = []*node.Pod{}
		for _, p := range dn.Pods {
			np = append(np, &node.Pod{
				Namespace: p.Namespace,
				Name:      p.Name,
			})
		}
		n := &node.Node{
			Name: dn.Name,
			Pods: np,
		}
		response.Nodes = append(response.Nodes, n)
	}
	return response, nil
}
