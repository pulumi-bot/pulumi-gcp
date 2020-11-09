// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gameservices

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A game server cluster resource.
//
// To get more information about GameServerCluster, see:
//
// * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms.gameServerClusters)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/game-servers/docs)
//
// ## Example Usage
//
// ## Import
//
// GameServerCluster can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default {{project}}/{{location}}/{{realm_id}}/{{cluster_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default {{location}}/{{realm_id}}/{{cluster_id}}
// ```
type GameServerCluster struct {
	pulumi.CustomResourceState

	// Required. The resource name of the game server cluster
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo GameServerClusterConnectionInfoOutput `pulumi:"connectionInfo"`
	// Human readable description of the cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location of the Cluster.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource id of the game server cluster, eg:
	// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The realm id of the game server realm.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGameServerCluster registers a new resource with the given unique name, arguments, and options.
func NewGameServerCluster(ctx *pulumi.Context,
	name string, args *GameServerClusterArgs, opts ...pulumi.ResourceOption) (*GameServerCluster, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil || args.ConnectionInfo == nil {
		return nil, errors.New("missing required argument 'ConnectionInfo'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &GameServerClusterArgs{}
	}
	var resource GameServerCluster
	err := ctx.RegisterResource("gcp:gameservices/gameServerCluster:GameServerCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameServerCluster gets an existing GameServerCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameServerCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameServerClusterState, opts ...pulumi.ResourceOption) (*GameServerCluster, error) {
	var resource GameServerCluster
	err := ctx.ReadResource("gcp:gameservices/gameServerCluster:GameServerCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameServerCluster resources.
type gameServerClusterState struct {
	// Required. The resource name of the game server cluster
	ClusterId *string `pulumi:"clusterId"`
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo *GameServerClusterConnectionInfo `pulumi:"connectionInfo"`
	// Human readable description of the cluster.
	Description *string `pulumi:"description"`
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Cluster.
	Location *string `pulumi:"location"`
	// The resource id of the game server cluster, eg:
	// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The realm id of the game server realm.
	RealmId *string `pulumi:"realmId"`
}

type GameServerClusterState struct {
	// Required. The resource name of the game server cluster
	ClusterId pulumi.StringPtrInput
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo GameServerClusterConnectionInfoPtrInput
	// Human readable description of the cluster.
	Description pulumi.StringPtrInput
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels pulumi.StringMapInput
	// Location of the Cluster.
	Location pulumi.StringPtrInput
	// The resource id of the game server cluster, eg:
	// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
	// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The realm id of the game server realm.
	RealmId pulumi.StringPtrInput
}

func (GameServerClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerClusterState)(nil)).Elem()
}

type gameServerClusterArgs struct {
	// Required. The resource name of the game server cluster
	ClusterId string `pulumi:"clusterId"`
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo GameServerClusterConnectionInfo `pulumi:"connectionInfo"`
	// Human readable description of the cluster.
	Description *string `pulumi:"description"`
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Cluster.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The realm id of the game server realm.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a GameServerCluster resource.
type GameServerClusterArgs struct {
	// Required. The resource name of the game server cluster
	ClusterId pulumi.StringInput
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	// Structure is documented below.
	ConnectionInfo GameServerClusterConnectionInfoInput
	// Human readable description of the cluster.
	Description pulumi.StringPtrInput
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	Labels pulumi.StringMapInput
	// Location of the Cluster.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The realm id of the game server realm.
	RealmId pulumi.StringInput
}

func (GameServerClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerClusterArgs)(nil)).Elem()
}
