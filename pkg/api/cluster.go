package api

import (
	"context"
	"fmt"
	"log"

	"github.com/endzyme/slipway/pkg/slipway"
	pb "github.com/endzyme/slipway/protobuf/slipway"
)

type ClusterHandler struct {
	slipwayCluster slipway.Cluster
}

// ListMembers will call to the slipway cluster to list the gossip members of the cluster
func (c ClusterHandler) ListMembers(ctx context.Context, request *pb.MembersRequest) (*pb.MembersResponse, error) {
	var response pb.MembersResponse

	members, err := c.slipwayCluster.GetMembers(nil, nil)
	if err != nil {
		log.Print("Error getting members")
		return &response, err
	}

	response.Result = transformClusterMemberToMembersResponse(members)

	return &response, nil
}

func transformClusterMemberToMembersResponse(clusterMembers []slipway.ClusterMember) []*pb.ClusterMember {
	var response []*pb.ClusterMember
	for _, clusterMember := range clusterMembers {
		mem := pb.ClusterMember{
			Id:        clusterMember.ID,
			Hostname:  clusterMember.Hostname,
			IpAddress: clusterMember.IPAddress.String(),
		}

		for key, val := range clusterMember.Tags {
			mem.Labels = append(mem.Labels, fmt.Sprintf("%v=%v", key, val))
		}
		response = append(response, &mem)
	}

	return response
}
