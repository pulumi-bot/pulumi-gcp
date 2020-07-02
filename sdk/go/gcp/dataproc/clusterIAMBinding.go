// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
//
// * `dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
// * `dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
// * `dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
//
// > **Note:** `dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `dataproc.ClusterIAMBinding` and `dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the cluster as `dataproc.ClusterIAMPolicy` replaces the entire policy.
//
// > **Note:** `dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_pubsub\_subscription\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/editor",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataproc.NewClusterIAMPolicy(ctx, "editor", &dataproc.ClusterIAMPolicyArgs{
// 			Project:    pulumi.String("your-project"),
// 			Region:     pulumi.String("your-region"),
// 			Cluster:    pulumi.String("your-dataproc-cluster"),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_subscription\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewClusterIAMBinding(ctx, "editor", &dataproc.ClusterIAMBindingArgs{
// 			Cluster: pulumi.String("your-dataproc-cluster"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_subscription\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewClusterIAMMember(ctx, "editor", &dataproc.ClusterIAMMemberArgs{
// 			Cluster: pulumi.String("your-dataproc-cluster"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Role:    pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ClusterIAMBinding struct {
	pulumi.CustomResourceState

	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   pulumi.StringOutput                 `pulumi:"cluster"`
	Condition ClusterIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewClusterIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewClusterIAMBinding(ctx *pulumi.Context,
	name string, args *ClusterIAMBindingArgs, opts ...pulumi.ResourceOption) (*ClusterIAMBinding, error) {
	if args == nil || args.Cluster == nil {
		return nil, errors.New("missing required argument 'Cluster'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &ClusterIAMBindingArgs{}
	}
	var resource ClusterIAMBinding
	err := ctx.RegisterResource("gcp:dataproc/clusterIAMBinding:ClusterIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterIAMBinding gets an existing ClusterIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterIAMBindingState, opts ...pulumi.ResourceOption) (*ClusterIAMBinding, error) {
	var resource ClusterIAMBinding
	err := ctx.ReadResource("gcp:dataproc/clusterIAMBinding:ClusterIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterIAMBinding resources.
type clusterIAMBindingState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   *string                     `pulumi:"cluster"`
	Condition *ClusterIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ClusterIAMBindingState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   pulumi.StringPtrInput
	Condition ClusterIAMBindingConditionPtrInput
	// (Computed) The etag of the clusters's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ClusterIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMBindingState)(nil)).Elem()
}

type clusterIAMBindingArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   string                      `pulumi:"cluster"`
	Condition *ClusterIAMBindingCondition `pulumi:"condition"`
	Members   []string                    `pulumi:"members"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ClusterIAMBinding resource.
type ClusterIAMBindingArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster   pulumi.StringInput
	Condition ClusterIAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ClusterIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMBindingArgs)(nil)).Elem()
}
