// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Data catalog EntryGroup. Each of these resources serves a different use case:
//
// * `datacatalog.EntryGroupIamPolicy`: Authoritative. Sets the IAM policy for the entrygroup and replaces any existing policy already attached.
// * `datacatalog.EntryGroupIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the entrygroup are preserved.
// * `datacatalog.EntryGroupIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the entrygroup are preserved.
//
// > **Note:** `datacatalog.EntryGroupIamPolicy` **cannot** be used in conjunction with `datacatalog.EntryGroupIamBinding` and `datacatalog.EntryGroupIamMember` or they will fight over what your policy should be.
//
// > **Note:** `datacatalog.EntryGroupIamBinding` resources **can be** used in conjunction with `datacatalog.EntryGroupIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} * {{project}}/{{region}}/{{entry_group}} * {{region}}/{{entry_group}} * {{entry_group}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog entrygroup IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/entryGroupIamPolicy:EntryGroupIamPolicy editor "projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/entryGroupIamPolicy:EntryGroupIamPolicy editor "projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/entryGroupIamPolicy:EntryGroupIamPolicy editor projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type EntryGroupIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	EntryGroup pulumi.StringOutput `pulumi:"entryGroup"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
}

// NewEntryGroupIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewEntryGroupIamPolicy(ctx *pulumi.Context,
	name string, args *EntryGroupIamPolicyArgs, opts ...pulumi.ResourceOption) (*EntryGroupIamPolicy, error) {
	if args == nil || args.EntryGroup == nil {
		return nil, errors.New("missing required argument 'EntryGroup'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &EntryGroupIamPolicyArgs{}
	}
	var resource EntryGroupIamPolicy
	err := ctx.RegisterResource("gcp:datacatalog/entryGroupIamPolicy:EntryGroupIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryGroupIamPolicy gets an existing EntryGroupIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryGroupIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryGroupIamPolicyState, opts ...pulumi.ResourceOption) (*EntryGroupIamPolicy, error) {
	var resource EntryGroupIamPolicy
	err := ctx.ReadResource("gcp:datacatalog/entryGroupIamPolicy:EntryGroupIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryGroupIamPolicy resources.
type entryGroupIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup *string `pulumi:"entryGroup"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

type EntryGroupIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
}

func (EntryGroupIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamPolicyState)(nil)).Elem()
}

type entryGroupIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup string `pulumi:"entryGroup"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// The set of arguments for constructing a EntryGroupIamPolicy resource.
type EntryGroupIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
}

func (EntryGroupIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamPolicyArgs)(nil)).Elem()
}
