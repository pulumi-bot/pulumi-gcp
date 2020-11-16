// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type IAMPolicy struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput `pulumi:"etag"`
	Folder     pulumi.StringOutput `pulumi:"folder"`
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	if args == nil || args.Folder == nil {
		return nil, errors.New("missing required argument 'Folder'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &IAMPolicyArgs{}
	}
	var resource IAMPolicy
	err := ctx.RegisterResource("gcp:folder/iAMPolicy:IAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMPolicy gets an existing IAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMPolicyState, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	var resource IAMPolicy
	err := ctx.ReadResource("gcp:folder/iAMPolicy:IAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMPolicy resources.
type iampolicyState struct {
	Etag       *string `pulumi:"etag"`
	Folder     *string `pulumi:"folder"`
	PolicyData *string `pulumi:"policyData"`
}

type IAMPolicyState struct {
	Etag       pulumi.StringPtrInput
	Folder     pulumi.StringPtrInput
	PolicyData pulumi.StringPtrInput
}

func (IAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyState)(nil)).Elem()
}

type iampolicyArgs struct {
	Folder     string `pulumi:"folder"`
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	Folder     pulumi.StringInput
	PolicyData pulumi.StringInput
}

func (IAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyArgs)(nil)).Elem()
}

type IAMPolicyInput interface {
	pulumi.Input

	ToIAMPolicyOutput() IAMPolicyOutput
	ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput
}

func (IAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMPolicy)(nil)).Elem()
}

func (i IAMPolicy) ToIAMPolicyOutput() IAMPolicyOutput {
	return i.ToIAMPolicyOutputWithContext(context.Background())
}

func (i IAMPolicy) ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyOutput)
}

type IAMPolicyOutput struct {
	*pulumi.OutputState
}

func (IAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMPolicyOutput)(nil)).Elem()
}

func (o IAMPolicyOutput) ToIAMPolicyOutput() IAMPolicyOutput {
	return o
}

func (o IAMPolicyOutput) ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IAMPolicyOutput{})
}
