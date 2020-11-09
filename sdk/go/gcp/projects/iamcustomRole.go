// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows management of a customized Cloud IAM project role. For more information see
// [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
// and
// [API](https://cloud.google.com/iam/reference/rest/v1/projects.roles).
//
// > **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
//  from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
//  same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
//  after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
//  made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
//  by the provider, and new roles cannot share that name.
//
// ## Import
//
// Customized IAM project role can be imported using their URI, e.g.
//
// ```sh
//  $ pulumi import gcp:projects/iAMCustomRole:IAMCustomRole my-custom-role projects/my-project/roles/myCustomRole
// ```
type IAMCustomRole struct {
	pulumi.CustomResourceState

	// (Optional) The current deleted state of the role.
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// A human-readable description for the role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
	Name pulumi.StringOutput `pulumi:"name"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	// The project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringOutput `pulumi:"project"`
	// The camel case role id to use for this role. Cannot contain `-` characters.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage pulumi.StringPtrOutput `pulumi:"stage"`
	// A human-readable title for the role.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewIAMCustomRole registers a new resource with the given unique name, arguments, and options.
func NewIAMCustomRole(ctx *pulumi.Context,
	name string, args *IAMCustomRoleArgs, opts ...pulumi.ResourceOption) (*IAMCustomRole, error) {
	if args == nil || args.Permissions == nil {
		return nil, errors.New("missing required argument 'Permissions'")
	}
	if args == nil || args.RoleId == nil {
		return nil, errors.New("missing required argument 'RoleId'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	if args == nil {
		args = &IAMCustomRoleArgs{}
	}
	var resource IAMCustomRole
	err := ctx.RegisterResource("gcp:projects/iAMCustomRole:IAMCustomRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMCustomRole gets an existing IAMCustomRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMCustomRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMCustomRoleState, opts ...pulumi.ResourceOption) (*IAMCustomRole, error) {
	var resource IAMCustomRole
	err := ctx.ReadResource("gcp:projects/iAMCustomRole:IAMCustomRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMCustomRole resources.
type iamcustomRoleState struct {
	// (Optional) The current deleted state of the role.
	Deleted *bool `pulumi:"deleted"`
	// A human-readable description for the role.
	Description *string `pulumi:"description"`
	// The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
	Name *string `pulumi:"name"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `pulumi:"permissions"`
	// The project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `pulumi:"project"`
	// The camel case role id to use for this role. Cannot contain `-` characters.
	RoleId *string `pulumi:"roleId"`
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage *string `pulumi:"stage"`
	// A human-readable title for the role.
	Title *string `pulumi:"title"`
}

type IAMCustomRoleState struct {
	// (Optional) The current deleted state of the role.
	Deleted pulumi.BoolPtrInput
	// A human-readable description for the role.
	Description pulumi.StringPtrInput
	// The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
	Name pulumi.StringPtrInput
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayInput
	// The project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringPtrInput
	// The camel case role id to use for this role. Cannot contain `-` characters.
	RoleId pulumi.StringPtrInput
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage pulumi.StringPtrInput
	// A human-readable title for the role.
	Title pulumi.StringPtrInput
}

func (IAMCustomRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamcustomRoleState)(nil)).Elem()
}

type iamcustomRoleArgs struct {
	// A human-readable description for the role.
	Description *string `pulumi:"description"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `pulumi:"permissions"`
	// The project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `pulumi:"project"`
	// The camel case role id to use for this role. Cannot contain `-` characters.
	RoleId string `pulumi:"roleId"`
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage *string `pulumi:"stage"`
	// A human-readable title for the role.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a IAMCustomRole resource.
type IAMCustomRoleArgs struct {
	// A human-readable description for the role.
	Description pulumi.StringPtrInput
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayInput
	// The project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringPtrInput
	// The camel case role id to use for this role. Cannot contain `-` characters.
	RoleId pulumi.StringInput
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage pulumi.StringPtrInput
	// A human-readable title for the role.
	Title pulumi.StringInput
}

func (IAMCustomRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamcustomRoleArgs)(nil)).Elem()
}
