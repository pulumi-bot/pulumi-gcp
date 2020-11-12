// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
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
// ## Import
//
// Cluster IAM resources can be imported using the project, region, cluster name, role and/or member.
//
// ```sh
//  $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster}"
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ClusterIAMPolicy struct {
	pulumi.CustomResourceState

	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewClusterIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewClusterIAMPolicy(ctx *pulumi.Context,
	name string, args *ClusterIAMPolicyArgs, opts ...pulumi.ResourceOption) (*ClusterIAMPolicy, error) {
	if args == nil || args.Cluster == nil {
		return nil, errors.New("missing required argument 'Cluster'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &ClusterIAMPolicyArgs{}
	}
	var resource ClusterIAMPolicy
	err := ctx.RegisterResource("gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterIAMPolicy gets an existing ClusterIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterIAMPolicyState, opts ...pulumi.ResourceOption) (*ClusterIAMPolicy, error) {
	var resource ClusterIAMPolicy
	err := ctx.ReadResource("gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterIAMPolicy resources.
type clusterIAMPolicyState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster *string `pulumi:"cluster"`
	// (Computed) The etag of the clusters's IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
}

type ClusterIAMPolicyState struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster pulumi.StringPtrInput
	// (Computed) The etag of the clusters's IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
}

func (ClusterIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMPolicyState)(nil)).Elem()
}

type clusterIAMPolicyArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster string `pulumi:"cluster"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ClusterIAMPolicy resource.
type ClusterIAMPolicyArgs struct {
	// The name or relative resource id of the cluster to manage IAM policies for.
	Cluster pulumi.StringInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The project in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the cluster belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
}

func (ClusterIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterIAMPolicyArgs)(nil)).Elem()
}

type ClusterIAMPolicyInput interface {
	pulumi.Input

	ToClusterIAMPolicyOutput() ClusterIAMPolicyOutput
	ToClusterIAMPolicyOutputWithContext(ctx context.Context) ClusterIAMPolicyOutput
}

func (ClusterIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIAMPolicy)(nil)).Elem()
}

func (i ClusterIAMPolicy) ToClusterIAMPolicyOutput() ClusterIAMPolicyOutput {
	return i.ToClusterIAMPolicyOutputWithContext(context.Background())
}

func (i ClusterIAMPolicy) ToClusterIAMPolicyOutputWithContext(ctx context.Context) ClusterIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterIAMPolicyOutput)
}

type ClusterIAMPolicyOutput struct {
	*pulumi.OutputState
}

func (ClusterIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterIAMPolicyOutput)(nil)).Elem()
}

func (o ClusterIAMPolicyOutput) ToClusterIAMPolicyOutput() ClusterIAMPolicyOutput {
	return o
}

func (o ClusterIAMPolicyOutput) ToClusterIAMPolicyOutputWithContext(ctx context.Context) ClusterIAMPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterIAMPolicyOutput{})
}
