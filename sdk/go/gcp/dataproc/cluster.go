// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dataproc

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Cloud Dataproc cluster resource within GCP. For more information see
// [the official dataproc documentation](https://cloud.google.com/dataproc/).
// 
// 
// !> **Warning:** Due to limitations of the API, all arguments except
// `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updatable. Changing others will cause recreation of the
// whole cluster!
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster.html.markdown.
type Cluster struct {
	pulumi.CustomResourceState

	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig ClusterClusterConfigOutput `pulumi:"clusterConfig"`
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the cluster, unique within the project and
	// zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region pulumi.StringPtrOutput `pulumi:"region"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("gcp:dataproc/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("gcp:dataproc/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig *ClusterClusterConfig `pulumi:"clusterConfig"`
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels map[string]string `pulumi:"labels"`
	// The name of the cluster, unique within the project and
	// zone.
	Name *string `pulumi:"name"`
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region *string `pulumi:"region"`
}

type ClusterState struct {
	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig ClusterClusterConfigPtrInput
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels pulumi.StringMapInput
	// The name of the cluster, unique within the project and
	// zone.
	Name pulumi.StringPtrInput
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig *ClusterClusterConfig `pulumi:"clusterConfig"`
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels map[string]string `pulumi:"labels"`
	// The name of the cluster, unique within the project and
	// zone.
	Name *string `pulumi:"name"`
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig ClusterClusterConfigPtrInput
	// The list of labels (key/value pairs) to be applied to
	// instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
	// which is the name of the cluster.
	Labels pulumi.StringMapInput
	// The name of the cluster, unique within the project and
	// zone.
	Name pulumi.StringPtrInput
	// The ID of the project in which the `cluster` will exist. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the cluster and associated nodes will be created in.
	// Defaults to `global`.
	Region pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

