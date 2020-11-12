// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An AccessLevel is a label that can be applied to requests to GCP services,
// along with a list of requirements necessary for the label to be applied.
//
// To get more information about AccessLevel, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
// * How-to Guides
//     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the ACM API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
//
// ## Import
//
// AccessLevel can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/accessLevel:AccessLevel default {{name}}
// ```
type AccessLevel struct {
	pulumi.CustomResourceState

	// A set of predefined conditions for the access level and a combining function.
	// Structure is documented below.
	Basic AccessLevelBasicPtrOutput `pulumi:"basic"`
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// Structure is documented below.
	Custom AccessLevelCustomPtrOutput `pulumi:"custom"`
	// Description of the expression
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name for the Access Level. The shortName component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringOutput `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Title for the expression, i.e. a short string describing its purpose.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewAccessLevel registers a new resource with the given unique name, arguments, and options.
func NewAccessLevel(ctx *pulumi.Context,
	name string, args *AccessLevelArgs, opts ...pulumi.ResourceOption) (*AccessLevel, error) {
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	if args == nil {
		args = &AccessLevelArgs{}
	}
	var resource AccessLevel
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessLevel:AccessLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessLevel gets an existing AccessLevel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessLevelState, opts ...pulumi.ResourceOption) (*AccessLevel, error) {
	var resource AccessLevel
	err := ctx.ReadResource("gcp:accesscontextmanager/accessLevel:AccessLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessLevel resources.
type accessLevelState struct {
	// A set of predefined conditions for the access level and a combining function.
	// Structure is documented below.
	Basic *AccessLevelBasic `pulumi:"basic"`
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// Structure is documented below.
	Custom *AccessLevelCustom `pulumi:"custom"`
	// Description of the expression
	Description *string `pulumi:"description"`
	// Resource name for the Access Level. The shortName component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `pulumi:"parent"`
	// Title for the expression, i.e. a short string describing its purpose.
	Title *string `pulumi:"title"`
}

type AccessLevelState struct {
	// A set of predefined conditions for the access level and a combining function.
	// Structure is documented below.
	Basic AccessLevelBasicPtrInput
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// Structure is documented below.
	Custom AccessLevelCustomPtrInput
	// Description of the expression
	Description pulumi.StringPtrInput
	// Resource name for the Access Level. The shortName component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringPtrInput
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringPtrInput
	// Title for the expression, i.e. a short string describing its purpose.
	Title pulumi.StringPtrInput
}

func (AccessLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelState)(nil)).Elem()
}

type accessLevelArgs struct {
	// A set of predefined conditions for the access level and a combining function.
	// Structure is documented below.
	Basic *AccessLevelBasic `pulumi:"basic"`
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// Structure is documented below.
	Custom *AccessLevelCustom `pulumi:"custom"`
	// Description of the expression
	Description *string `pulumi:"description"`
	// Resource name for the Access Level. The shortName component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent string `pulumi:"parent"`
	// Title for the expression, i.e. a short string describing its purpose.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a AccessLevel resource.
type AccessLevelArgs struct {
	// A set of predefined conditions for the access level and a combining function.
	// Structure is documented below.
	Basic AccessLevelBasicPtrInput
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// Structure is documented below.
	Custom AccessLevelCustomPtrInput
	// Description of the expression
	Description pulumi.StringPtrInput
	// Resource name for the Access Level. The shortName component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringPtrInput
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringInput
	// Title for the expression, i.e. a short string describing its purpose.
	Title pulumi.StringInput
}

func (AccessLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelArgs)(nil)).Elem()
}

type AccessLevelInput interface {
	pulumi.Input

	ToAccessLevelOutput() AccessLevelOutput
	ToAccessLevelOutputWithContext(ctx context.Context) AccessLevelOutput
}

func (AccessLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessLevel)(nil)).Elem()
}

func (i AccessLevel) ToAccessLevelOutput() AccessLevelOutput {
	return i.ToAccessLevelOutputWithContext(context.Background())
}

func (i AccessLevel) ToAccessLevelOutputWithContext(ctx context.Context) AccessLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLevelOutput)
}

type AccessLevelOutput struct {
	*pulumi.OutputState
}

func (AccessLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessLevelOutput)(nil)).Elem()
}

func (o AccessLevelOutput) ToAccessLevelOutput() AccessLevelOutput {
	return o
}

func (o AccessLevelOutput) ToAccessLevelOutputWithContext(ctx context.Context) AccessLevelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccessLevelOutput{})
}
