package client

import (
	"context"
	"time"

	"github.com/canonical/microcluster/internal/rest/types"
	"github.com/lxc/lxd/shared/api"
)

// AddClusterMember records a new cluster member in the trust store of each current cluster member.
func (c *Client) AddClusterMember(ctx context.Context, args types.ClusterMember) (*types.TokenResponse, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	tokenResponse := types.TokenResponse{}
	err := c.QueryStruct(queryCtx, "POST", PublicEndpoint, api.NewURL().Path("cluster"), args, &tokenResponse)
	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

// GetClusterMembers returns the database record of cluster members.
func (c *Client) GetClusterMembers(ctx context.Context) ([]types.ClusterMember, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	clusterMembers := []types.ClusterMember{}
	err := c.QueryStruct(queryCtx, "GET", PublicEndpoint, api.NewURL().Path("cluster"), nil, &clusterMembers)

	return clusterMembers, err
}

// DeleteClusterMember deletes the cluster member with the given name.
func (c *Client) DeleteClusterMember(ctx context.Context, name string, force bool) error {
	queryCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	endpoint := api.NewURL().Path("cluster", name)
	if force {
		endpoint = endpoint.WithQuery("force", "1")
	}

	return c.QueryStruct(queryCtx, "DELETE", PublicEndpoint, endpoint, nil, nil)
}

// ResetClusterMember clears the state directory of the cluster member, and re-execs its daemon.
func (c *Client) ResetClusterMember(ctx context.Context, name string) error {
	queryCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	return c.QueryStruct(queryCtx, "PUT", PublicEndpoint, api.NewURL().Path("cluster", name), nil, nil)
}
