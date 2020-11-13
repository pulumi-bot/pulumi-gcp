// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Replace all existing Access Levels in an Access Policy with the Access Levels provided. This is done atomically.
// This is a bulk edit of all Access Levels and may override existing Access Levels created by `accesscontextmanager.AccessLevel`,
// thus causing a permadiff if used alongside `accesscontextmanager.AccessLevel` on the same parent.
//
// To get more information about AccessLevels, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
// * How-to Guides
//     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
//
// ## Example Usage
type AccessLevels struct {
	pulumi.CustomResourceState

	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels AccessLevelsAccessLevelArrayOutput `pulumi:"accessLevels"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
}

// NewAccessLevels registers a new resource with the given unique name, arguments, and options.
func NewAccessLevels(ctx *pulumi.Context,
	name string, args *AccessLevelsArgs, opts ...pulumi.ResourceOption) (*AccessLevels, error) {
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil {
		args = &AccessLevelsArgs{}
	}
	var resource AccessLevels
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessLevels:AccessLevels", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessLevels gets an existing AccessLevels resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessLevels(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessLevelsState, opts ...pulumi.ResourceOption) (*AccessLevels, error) {
	var resource AccessLevels
	err := ctx.ReadResource("gcp:accesscontextmanager/accessLevels:AccessLevels", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessLevels resources.
type accessLevelsState struct {
	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels []AccessLevelsAccessLevel `pulumi:"accessLevels"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `pulumi:"parent"`
}

type AccessLevelsState struct {
	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels AccessLevelsAccessLevelArrayInput
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringPtrInput
}

func (AccessLevelsState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelsState)(nil)).Elem()
}

type accessLevelsArgs struct {
	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels []AccessLevelsAccessLevel `pulumi:"accessLevels"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent string `pulumi:"parent"`
}

// The set of arguments for constructing a AccessLevels resource.
type AccessLevelsArgs struct {
	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// Structure is documented below.
	AccessLevels AccessLevelsAccessLevelArrayInput
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringInput
}

func (AccessLevelsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelsArgs)(nil)).Elem()
}

type AccessLevelsInput interface {
	pulumi.Input

	ToAccessLevelsOutput() AccessLevelsOutput
	ToAccessLevelsOutputWithContext(ctx context.Context) AccessLevelsOutput
}

func (AccessLevels) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessLevels)(nil)).Elem()
}

func (i AccessLevels) ToAccessLevelsOutput() AccessLevelsOutput {
	return i.ToAccessLevelsOutputWithContext(context.Background())
}

func (i AccessLevels) ToAccessLevelsOutputWithContext(ctx context.Context) AccessLevelsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLevelsOutput)
}

type AccessLevelsOutput struct {
	*pulumi.OutputState
}

func (AccessLevelsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessLevelsOutput)(nil)).Elem()
}

func (o AccessLevelsOutput) ToAccessLevelsOutput() AccessLevelsOutput {
	return o
}

func (o AccessLevelsOutput) ToAccessLevelsOutputWithContext(ctx context.Context) AccessLevelsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccessLevelsOutput{})
}
