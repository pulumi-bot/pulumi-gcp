// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get info about a GKE cluster from its name and location.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/container"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myCluster, err := container.LookupCluster(ctx, &container.LookupClusterArgs{
// 			Name:     "my-cluster",
// 			Location: "us-east1-a",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("clusterUsername", myCluster.MasterAuths[0].Username)
// 		ctx.Export("clusterPassword", myCluster.MasterAuths[0].Password)
// 		ctx.Export("endpoint", myCluster.Endpoint)
// 		ctx.Export("instanceGroupUrls", myCluster.InstanceGroupUrls)
// 		ctx.Export("nodeConfig", myCluster.NodeConfigs)
// 		ctx.Export("nodePools", myCluster.NodePools)
// 		return nil
// 	})
// }
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("gcp:container/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// The location (zone or region) this cluster has been
	// created in. One of `location`, `region`, `zone`, or a provider-level `zone` must
	// be specified.
	Location *string `pulumi:"location"`
	// The name of the cluster.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region this cluster has been created in. Deprecated
	// in favour of `location`.
	Region *string `pulumi:"region"`
	// The zone this cluster has been created in. Deprecated in
	// favour of `location`.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	AdditionalZones            []string                              `pulumi:"additionalZones"`
	AddonsConfigs              []GetClusterAddonsConfig              `pulumi:"addonsConfigs"`
	AuthenticatorGroupsConfigs []GetClusterAuthenticatorGroupsConfig `pulumi:"authenticatorGroupsConfigs"`
	ClusterAutoscalings        []GetClusterClusterAutoscaling        `pulumi:"clusterAutoscalings"`
	ClusterIpv4Cidr            string                                `pulumi:"clusterIpv4Cidr"`
	ClusterTelemetries         []GetClusterClusterTelemetry          `pulumi:"clusterTelemetries"`
	DatabaseEncryptions        []GetClusterDatabaseEncryption        `pulumi:"databaseEncryptions"`
	DefaultMaxPodsPerNode      int                                   `pulumi:"defaultMaxPodsPerNode"`
	Description                string                                `pulumi:"description"`
	EnableBinaryAuthorization  bool                                  `pulumi:"enableBinaryAuthorization"`
	EnableIntranodeVisibility  bool                                  `pulumi:"enableIntranodeVisibility"`
	EnableKubernetesAlpha      bool                                  `pulumi:"enableKubernetesAlpha"`
	EnableLegacyAbac           bool                                  `pulumi:"enableLegacyAbac"`
	EnableShieldedNodes        bool                                  `pulumi:"enableShieldedNodes"`
	EnableTpu                  bool                                  `pulumi:"enableTpu"`
	Endpoint                   string                                `pulumi:"endpoint"`
	// The provider-assigned unique ID for this managed resource.
	Id                              string                                     `pulumi:"id"`
	InitialNodeCount                int                                        `pulumi:"initialNodeCount"`
	InstanceGroupUrls               []string                                   `pulumi:"instanceGroupUrls"`
	IpAllocationPolicies            []GetClusterIpAllocationPolicy             `pulumi:"ipAllocationPolicies"`
	LabelFingerprint                string                                     `pulumi:"labelFingerprint"`
	Location                        *string                                    `pulumi:"location"`
	LoggingService                  string                                     `pulumi:"loggingService"`
	MaintenancePolicies             []GetClusterMaintenancePolicy              `pulumi:"maintenancePolicies"`
	MasterAuthorizedNetworksConfigs []GetClusterMasterAuthorizedNetworksConfig `pulumi:"masterAuthorizedNetworksConfigs"`
	MasterAuths                     []GetClusterMasterAuth                     `pulumi:"masterAuths"`
	MasterVersion                   string                                     `pulumi:"masterVersion"`
	MinMasterVersion                string                                     `pulumi:"minMasterVersion"`
	MonitoringService               string                                     `pulumi:"monitoringService"`
	Name                            string                                     `pulumi:"name"`
	Network                         string                                     `pulumi:"network"`
	NetworkPolicies                 []GetClusterNetworkPolicy                  `pulumi:"networkPolicies"`
	NodeConfigs                     []GetClusterNodeConfig                     `pulumi:"nodeConfigs"`
	NodeLocations                   []string                                   `pulumi:"nodeLocations"`
	NodePools                       []GetClusterNodePool                       `pulumi:"nodePools"`
	NodeVersion                     string                                     `pulumi:"nodeVersion"`
	Operation                       string                                     `pulumi:"operation"`
	PodSecurityPolicyConfigs        []GetClusterPodSecurityPolicyConfig        `pulumi:"podSecurityPolicyConfigs"`
	PrivateClusterConfigs           []GetClusterPrivateClusterConfig           `pulumi:"privateClusterConfigs"`
	Project                         *string                                    `pulumi:"project"`
	Region                          *string                                    `pulumi:"region"`
	ReleaseChannels                 []GetClusterReleaseChannel                 `pulumi:"releaseChannels"`
	RemoveDefaultNodePool           bool                                       `pulumi:"removeDefaultNodePool"`
	ResourceLabels                  map[string]string                          `pulumi:"resourceLabels"`
	ResourceUsageExportConfigs      []GetClusterResourceUsageExportConfig      `pulumi:"resourceUsageExportConfigs"`
	ServicesIpv4Cidr                string                                     `pulumi:"servicesIpv4Cidr"`
	Subnetwork                      string                                     `pulumi:"subnetwork"`
	TpuIpv4CidrBlock                string                                     `pulumi:"tpuIpv4CidrBlock"`
	VerticalPodAutoscalings         []GetClusterVerticalPodAutoscaling         `pulumi:"verticalPodAutoscalings"`
	WorkloadIdentityConfigs         []GetClusterWorkloadIdentityConfig         `pulumi:"workloadIdentityConfigs"`
	Zone                            *string                                    `pulumi:"zone"`
}
