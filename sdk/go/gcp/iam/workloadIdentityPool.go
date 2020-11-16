// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type WorkloadIdentityPool struct {
	pulumi.CustomResourceState

	// A description of the pool. Cannot exceed 256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
	// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	// access again.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// A display name for the pool. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The resource name of the pool as
	// 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The state of the pool. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The pool is active, and may be used in Google
	// Cloud policies. * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30
	// days. You can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted
	// pool until it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing
	// tokens to access resources. If the pool is undeleted, existing tokens grant access again.
	State pulumi.StringOutput `pulumi:"state"`
	// The ID to use for the pool, which becomes the final component of the resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId pulumi.StringOutput `pulumi:"workloadIdentityPoolId"`
}

// NewWorkloadIdentityPool registers a new resource with the given unique name, arguments, and options.
func NewWorkloadIdentityPool(ctx *pulumi.Context,
	name string, args *WorkloadIdentityPoolArgs, opts ...pulumi.ResourceOption) (*WorkloadIdentityPool, error) {
	if args == nil || args.WorkloadIdentityPoolId == nil {
		return nil, errors.New("missing required argument 'WorkloadIdentityPoolId'")
	}
	if args == nil {
		args = &WorkloadIdentityPoolArgs{}
	}
	var resource WorkloadIdentityPool
	err := ctx.RegisterResource("gcp:iam/workloadIdentityPool:WorkloadIdentityPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadIdentityPool gets an existing WorkloadIdentityPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadIdentityPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadIdentityPoolState, opts ...pulumi.ResourceOption) (*WorkloadIdentityPool, error) {
	var resource WorkloadIdentityPool
	err := ctx.ReadResource("gcp:iam/workloadIdentityPool:WorkloadIdentityPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadIdentityPool resources.
type workloadIdentityPoolState struct {
	// A description of the pool. Cannot exceed 256 characters.
	Description *string `pulumi:"description"`
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
	// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	// access again.
	Disabled *bool `pulumi:"disabled"`
	// A display name for the pool. Cannot exceed 32 characters.
	DisplayName *string `pulumi:"displayName"`
	// The resource name of the pool as
	// 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The state of the pool. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The pool is active, and may be used in Google
	// Cloud policies. * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30
	// days. You can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted
	// pool until it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing
	// tokens to access resources. If the pool is undeleted, existing tokens grant access again.
	State *string `pulumi:"state"`
	// The ID to use for the pool, which becomes the final component of the resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId *string `pulumi:"workloadIdentityPoolId"`
}

type WorkloadIdentityPoolState struct {
	// A description of the pool. Cannot exceed 256 characters.
	Description pulumi.StringPtrInput
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
	// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	// access again.
	Disabled pulumi.BoolPtrInput
	// A display name for the pool. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrInput
	// The resource name of the pool as
	// 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The state of the pool. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The pool is active, and may be used in Google
	// Cloud policies. * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30
	// days. You can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted
	// pool until it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing
	// tokens to access resources. If the pool is undeleted, existing tokens grant access again.
	State pulumi.StringPtrInput
	// The ID to use for the pool, which becomes the final component of the resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId pulumi.StringPtrInput
}

func (WorkloadIdentityPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadIdentityPoolState)(nil)).Elem()
}

type workloadIdentityPoolArgs struct {
	// A description of the pool. Cannot exceed 256 characters.
	Description *string `pulumi:"description"`
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
	// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	// access again.
	Disabled *bool `pulumi:"disabled"`
	// A display name for the pool. Cannot exceed 32 characters.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The ID to use for the pool, which becomes the final component of the resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId string `pulumi:"workloadIdentityPoolId"`
}

// The set of arguments for constructing a WorkloadIdentityPool resource.
type WorkloadIdentityPoolArgs struct {
	// A description of the pool. Cannot exceed 256 characters.
	Description pulumi.StringPtrInput
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
	// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	// access again.
	Disabled pulumi.BoolPtrInput
	// A display name for the pool. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The ID to use for the pool, which becomes the final component of the resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId pulumi.StringInput
}

func (WorkloadIdentityPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadIdentityPoolArgs)(nil)).Elem()
}

type WorkloadIdentityPoolInput interface {
	pulumi.Input

	ToWorkloadIdentityPoolOutput() WorkloadIdentityPoolOutput
	ToWorkloadIdentityPoolOutputWithContext(ctx context.Context) WorkloadIdentityPoolOutput
}

func (WorkloadIdentityPool) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadIdentityPool)(nil)).Elem()
}

func (i WorkloadIdentityPool) ToWorkloadIdentityPoolOutput() WorkloadIdentityPoolOutput {
	return i.ToWorkloadIdentityPoolOutputWithContext(context.Background())
}

func (i WorkloadIdentityPool) ToWorkloadIdentityPoolOutputWithContext(ctx context.Context) WorkloadIdentityPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadIdentityPoolOutput)
}

type WorkloadIdentityPoolOutput struct {
	*pulumi.OutputState
}

func (WorkloadIdentityPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadIdentityPoolOutput)(nil)).Elem()
}

func (o WorkloadIdentityPoolOutput) ToWorkloadIdentityPoolOutput() WorkloadIdentityPoolOutput {
	return o
}

func (o WorkloadIdentityPoolOutput) ToWorkloadIdentityPoolOutputWithContext(ctx context.Context) WorkloadIdentityPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkloadIdentityPoolOutput{})
}
