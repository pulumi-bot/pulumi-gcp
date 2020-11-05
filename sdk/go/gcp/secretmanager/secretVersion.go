// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretmanager

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A secret version resource.
//
// > **Warning:** All arguments including `payload.secret_data` will be stored in the raw
// state as plain-text.
//
// ## Example Usage
type SecretVersion struct {
	pulumi.CustomResourceState

	// The time at which the Secret was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time at which the Secret was destroyed. Only present if state is DESTROYED.
	DestroyTime pulumi.StringOutput `pulumi:"destroyTime"`
	// The current state of the SecretVersion.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The resource name of the SecretVersion. Format: 'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'
	Name pulumi.StringOutput `pulumi:"name"`
	// Secret Manager secret resource
	Secret pulumi.StringOutput `pulumi:"secret"`
	// The secret data. Must be no larger than 64KiB.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SecretData pulumi.StringPtrOutput `pulumi:"secretData"`
}

// NewSecretVersion registers a new resource with the given unique name, arguments, and options.
func NewSecretVersion(ctx *pulumi.Context,
	name string, args *SecretVersionArgs, opts ...pulumi.ResourceOption) (*SecretVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Secret == nil {
		return nil, errors.New("invalid value for required argument 'Secret'")
	}
	var resource SecretVersion
	err := ctx.RegisterResource("gcp:secretmanager/secretVersion:SecretVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretVersion gets an existing SecretVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretVersionState, opts ...pulumi.ResourceOption) (*SecretVersion, error) {
	var resource SecretVersion
	err := ctx.ReadResource("gcp:secretmanager/secretVersion:SecretVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretVersion resources.
type secretVersionState struct {
	// The time at which the Secret was created.
	CreateTime *string `pulumi:"createTime"`
	// The time at which the Secret was destroyed. Only present if state is DESTROYED.
	DestroyTime *string `pulumi:"destroyTime"`
	// The current state of the SecretVersion.
	Enabled *bool `pulumi:"enabled"`
	// The resource name of the SecretVersion. Format: 'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'
	Name *string `pulumi:"name"`
	// Secret Manager secret resource
	Secret *string `pulumi:"secret"`
	// The secret data. Must be no larger than 64KiB.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SecretData *string `pulumi:"secretData"`
}

type SecretVersionState struct {
	// The time at which the Secret was created.
	CreateTime pulumi.StringPtrInput
	// The time at which the Secret was destroyed. Only present if state is DESTROYED.
	DestroyTime pulumi.StringPtrInput
	// The current state of the SecretVersion.
	Enabled pulumi.BoolPtrInput
	// The resource name of the SecretVersion. Format: 'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'
	Name pulumi.StringPtrInput
	// Secret Manager secret resource
	Secret pulumi.StringPtrInput
	// The secret data. Must be no larger than 64KiB.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SecretData pulumi.StringPtrInput
}

func (SecretVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretVersionState)(nil)).Elem()
}

type secretVersionArgs struct {
	// The current state of the SecretVersion.
	Enabled *bool `pulumi:"enabled"`
	// Secret Manager secret resource
	Secret string `pulumi:"secret"`
	// The secret data. Must be no larger than 64KiB.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SecretData *string `pulumi:"secretData"`
}

// The set of arguments for constructing a SecretVersion resource.
type SecretVersionArgs struct {
	// The current state of the SecretVersion.
	Enabled pulumi.BoolPtrInput
	// Secret Manager secret resource
	Secret pulumi.StringInput
	// The secret data. Must be no larger than 64KiB.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SecretData pulumi.StringPtrInput
}

func (SecretVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretVersionArgs)(nil)).Elem()
}
