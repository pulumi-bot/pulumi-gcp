// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* {{policy_tag}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor "{{policy_tag}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor "{{policy_tag}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor {{policy_tag}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type PolicyTagIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringOutput `pulumi:"policyTag"`
}

// NewPolicyTagIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicyTagIamPolicy(ctx *pulumi.Context,
	name string, args *PolicyTagIamPolicyArgs, opts ...pulumi.ResourceOption) (*PolicyTagIamPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.PolicyTag == nil {
		return nil, errors.New("missing required argument 'PolicyTag'")
	}
	if args == nil {
		args = &PolicyTagIamPolicyArgs{}
	}
	var resource PolicyTagIamPolicy
	err := ctx.RegisterResource("gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyTagIamPolicy gets an existing PolicyTagIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyTagIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyTagIamPolicyState, opts ...pulumi.ResourceOption) (*PolicyTagIamPolicy, error) {
	var resource PolicyTagIamPolicy
	err := ctx.ReadResource("gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyTagIamPolicy resources.
type policyTagIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag *string `pulumi:"policyTag"`
}

type PolicyTagIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringPtrInput
}

func (PolicyTagIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamPolicyState)(nil)).Elem()
}

type policyTagIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag string `pulumi:"policyTag"`
}

// The set of arguments for constructing a PolicyTagIamPolicy resource.
type PolicyTagIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	PolicyTag pulumi.StringInput
}

func (PolicyTagIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagIamPolicyArgs)(nil)).Elem()
}
