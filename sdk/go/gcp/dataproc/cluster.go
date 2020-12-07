// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Cloud Dataproc cluster resource within GCP. For more information see
// [the official dataproc documentation](https://cloud.google.com/dataproc/).
//
// !> **Warning:** Due to limitations of the API, all arguments except
// `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updatable. Changing others will cause recreation of the
// whole cluster!
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewCluster(ctx, "simplecluster", &dataproc.ClusterArgs{
// 			Region: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Advanced
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewCluster(ctx, "mycluster", &dataproc.ClusterArgs{
// 			ClusterConfig: &dataproc.ClusterClusterConfigArgs{
// 				GceClusterConfig: &dataproc.ClusterClusterConfigGceClusterConfigArgs{
// 					ServiceAccountScopes: pulumi.StringArray{
// 						pulumi.String("https://www.googleapis.com/auth/monitoring"),
// 						pulumi.String("useraccounts-ro"),
// 						pulumi.String("storage-rw"),
// 						pulumi.String("logging-write"),
// 					},
// 					Tags: pulumi.StringArray{
// 						pulumi.String("foo"),
// 						pulumi.String("bar"),
// 					},
// 				},
// 				InitializationActions: dataproc.ClusterClusterConfigInitializationActionArray{
// 					&dataproc.ClusterClusterConfigInitializationActionArgs{
// 						Script:     pulumi.String("gs://dataproc-initialization-actions/stackdriver/stackdriver.sh"),
// 						TimeoutSec: pulumi.Int(500),
// 					},
// 				},
// 				MasterConfig: &dataproc.ClusterClusterConfigMasterConfigArgs{
// 					DiskConfig: &dataproc.ClusterClusterConfigMasterConfigDiskConfigArgs{
// 						BootDiskSizeGb: pulumi.Int(15),
// 						BootDiskType:   pulumi.String("pd-ssd"),
// 					},
// 					MachineType:  pulumi.String("e2-medium"),
// 					NumInstances: pulumi.Int(1),
// 				},
// 				PreemptibleWorkerConfig: &dataproc.ClusterClusterConfigPreemptibleWorkerConfigArgs{
// 					NumInstances: pulumi.Int(0),
// 				},
// 				SoftwareConfig: &dataproc.ClusterClusterConfigSoftwareConfigArgs{
// 					ImageVersion: pulumi.String("1.3.7-deb9"),
// 					OverrideProperties: pulumi.StringMap{
// 						"dataproc:dataproc.allow.zero.workers": pulumi.String("true"),
// 					},
// 				},
// 				StagingBucket: pulumi.String("dataproc-staging-bucket"),
// 				WorkerConfig: &dataproc.ClusterClusterConfigWorkerConfigArgs{
// 					DiskConfig: &dataproc.ClusterClusterConfigWorkerConfigDiskConfigArgs{
// 						BootDiskSizeGb: pulumi.Int(15),
// 						NumLocalSsds:   pulumi.Int(1),
// 					},
// 					MachineType:    pulumi.String("e2-medium"),
// 					MinCpuPlatform: pulumi.String("Intel Skylake"),
// 					NumInstances:   pulumi.Int(2),
// 				},
// 			},
// 			GracefulDecommissionTimeout: pulumi.String("120s"),
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			Region: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using A GPU Accelerator
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewCluster(ctx, "acceleratedCluster", &dataproc.ClusterArgs{
// 			ClusterConfig: &dataproc.ClusterClusterConfigArgs{
// 				GceClusterConfig: &dataproc.ClusterClusterConfigGceClusterConfigArgs{
// 					Zone: pulumi.String("us-central1-a"),
// 				},
// 				MasterConfig: &dataproc.ClusterClusterConfigMasterConfigArgs{
// 					Accelerators: dataproc.ClusterClusterConfigMasterConfigAcceleratorArray{
// 						&dataproc.ClusterClusterConfigMasterConfigAcceleratorArgs{
// 							AcceleratorCount: pulumi.Int(1),
// 							AcceleratorType:  pulumi.String("nvidia-tesla-k80"),
// 						},
// 					},
// 				},
// 			},
// 			Region: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// This resource does not support import.
type Cluster struct {
	pulumi.CustomResourceState

	// Allows you to configure various aspects of the cluster.
	// Structure defined below.
	ClusterConfig ClusterClusterConfigOutput `pulumi:"clusterConfig"`
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout pulumi.StringPtrOutput `pulumi:"gracefulDecommissionTimeout"`
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
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout *string `pulumi:"gracefulDecommissionTimeout"`
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
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout pulumi.StringPtrInput
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
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout *string `pulumi:"gracefulDecommissionTimeout"`
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
	// The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	// terraform apply
	GracefulDecommissionTimeout pulumi.StringPtrInput
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

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil)).Elem()
}

func (i Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct {
	*pulumi.OutputState
}

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterOutput)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
