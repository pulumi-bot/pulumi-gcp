// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:
//
// * `spanner.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
//
// > **Warning:** It's entirely possibly to lock yourself out of your instance using `spanner.InstanceIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
//
// * `spanner.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `spanner.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `spanner.InstanceIAMPolicy` **cannot** be used in conjunction with `spanner.InstanceIAMBinding` and `spanner.InstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `spanner.InstanceIAMBinding` resources **can be** used in conjunction with `spanner.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
type InstanceIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the instance's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the instance.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewInstanceIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewInstanceIAMPolicy(ctx *pulumi.Context,
	name string, args *InstanceIAMPolicyArgs, opts ...pulumi.ResourceOption) (*InstanceIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource InstanceIAMPolicy
	err := ctx.RegisterResource("gcp:spanner/instanceIAMPolicy:InstanceIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIAMPolicy gets an existing InstanceIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIAMPolicyState, opts ...pulumi.ResourceOption) (*InstanceIAMPolicy, error) {
	var resource InstanceIAMPolicy
	err := ctx.ReadResource("gcp:spanner/instanceIAMPolicy:InstanceIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIAMPolicy resources.
type instanceIAMPolicyState struct {
	// (Computed) The etag of the instance's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the instance.
	Instance *string `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type InstanceIAMPolicyState struct {
	// (Computed) The etag of the instance's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the instance.
	Instance pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (InstanceIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMPolicyState)(nil)).Elem()
}

type instanceIAMPolicyArgs struct {
	// The name of the instance.
	Instance string `pulumi:"instance"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a InstanceIAMPolicy resource.
type InstanceIAMPolicyArgs struct {
	// The name of the instance.
	Instance pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (InstanceIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMPolicyArgs)(nil)).Elem()
}

type InstanceIAMPolicyInput interface {
	pulumi.Input

	ToInstanceIAMPolicyOutput() InstanceIAMPolicyOutput
	ToInstanceIAMPolicyOutputWithContext(ctx context.Context) InstanceIAMPolicyOutput
}

func (InstanceIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIAMPolicy)(nil)).Elem()
}

func (i InstanceIAMPolicy) ToInstanceIAMPolicyOutput() InstanceIAMPolicyOutput {
	return i.ToInstanceIAMPolicyOutputWithContext(context.Background())
}

func (i InstanceIAMPolicy) ToInstanceIAMPolicyOutputWithContext(ctx context.Context) InstanceIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMPolicyOutput)
}

type InstanceIAMPolicyOutput struct {
	*pulumi.OutputState
}

func (InstanceIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIAMPolicyOutput)(nil)).Elem()
}

func (o InstanceIAMPolicyOutput) ToInstanceIAMPolicyOutput() InstanceIAMPolicyOutput {
	return o
}

func (o InstanceIAMPolicyOutput) ToInstanceIAMPolicyOutputWithContext(ctx context.Context) InstanceIAMPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceIAMPolicyOutput{})
}
