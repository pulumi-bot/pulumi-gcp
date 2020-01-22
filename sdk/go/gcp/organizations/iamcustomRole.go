// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package organizations

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows management of a customized Cloud IAM organization role. For more information see
// [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
// and
// [API](https://cloud.google.com/iam/reference/rest/v1/organizations.roles).
// 
// > **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
//  from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
//  same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
//  after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
//  made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
//  by this provider, and new roles cannot share that name.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/organization_iam_custom_role.html.markdown.
type IAMCustomRole struct {
	pulumi.CustomResourceState

	// (Optional) The current deleted state of the role.
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// A human-readable description for the role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	// The role id to use for this role.
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
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
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
	err := ctx.RegisterResource("gcp:organizations/iAMCustomRole:IAMCustomRole", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:organizations/iAMCustomRole:IAMCustomRole", name, id, state, &resource, opts...)
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
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId *string `pulumi:"orgId"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `pulumi:"permissions"`
	// The role id to use for this role.
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
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringPtrInput
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayInput
	// The role id to use for this role.
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
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId string `pulumi:"orgId"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `pulumi:"permissions"`
	// The role id to use for this role.
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
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringInput
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayInput
	// The role id to use for this role.
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

