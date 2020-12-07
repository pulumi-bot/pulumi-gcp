// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudidentity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Membership defines a relationship between a Group and an entity belonging to that Group, referred to as a "member".
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
// ### Cloud Identity Group Membership User
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudidentity"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		group, err := cloudidentity.NewGroup(ctx, "group", &cloudidentity.GroupArgs{
// 			DisplayName: pulumi.String("my-identity-group"),
// 			Parent:      pulumi.String("customers/A01b123xz"),
// 			GroupKey: &cloudidentity.GroupGroupKeyArgs{
// 				Id: pulumi.String("my-identity-group@example.com"),
// 			},
// 			Labels: pulumi.StringMap{
// 				"cloudidentity.googleapis.com/groups.discussion_forum": pulumi.String(""),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudidentity.NewGroupMembership(ctx, "cloudIdentityGroupMembershipBasic", &cloudidentity.GroupMembershipArgs{
// 			Group: group.ID(),
// 			PreferredMemberKey: &cloudidentity.GroupMembershipPreferredMemberKeyArgs{
// 				Id: pulumi.String("cloud_identity_user@example.com"),
// 			},
// 			Roles: cloudidentity.GroupMembershipRoleArray{
// 				&cloudidentity.GroupMembershipRoleArgs{
// 					Name: pulumi.String("MEMBER"),
// 				},
// 				&cloudidentity.GroupMembershipRoleArgs{
// 					Name: pulumi.String("MANAGER"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// GroupMembership can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudidentity/groupMembership:GroupMembership default {{name}}
// ```
type GroupMembership struct {
	pulumi.CustomResourceState

	// The time when the Membership was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The name of the Group to create this membership in.
	Group pulumi.StringOutput `pulumi:"group"`
	// EntityKey of the member.
	MemberKey GroupMembershipMemberKeyOutput `pulumi:"memberKey"`
	// The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
	// Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
	Name pulumi.StringOutput `pulumi:"name"`
	// EntityKey of the member.
	// Structure is documented below.
	PreferredMemberKey GroupMembershipPreferredMemberKeyOutput `pulumi:"preferredMemberKey"`
	// The MembershipRoles that apply to the Membership.
	// Must not contain duplicate MembershipRoles with the same name.
	// Structure is documented below.
	Roles GroupMembershipRoleArrayOutput `pulumi:"roles"`
	// The type of the membership.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time when the Membership was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewGroupMembership(ctx *pulumi.Context,
	name string, args *GroupMembershipArgs, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	var resource GroupMembership
	err := ctx.RegisterResource("gcp:cloudidentity/groupMembership:GroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembership gets an existing GroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipState, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	var resource GroupMembership
	err := ctx.ReadResource("gcp:cloudidentity/groupMembership:GroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembership resources.
type groupMembershipState struct {
	// The time when the Membership was created.
	CreateTime *string `pulumi:"createTime"`
	// The name of the Group to create this membership in.
	Group *string `pulumi:"group"`
	// EntityKey of the member.
	MemberKey *GroupMembershipMemberKey `pulumi:"memberKey"`
	// The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
	// Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
	Name *string `pulumi:"name"`
	// EntityKey of the member.
	// Structure is documented below.
	PreferredMemberKey *GroupMembershipPreferredMemberKey `pulumi:"preferredMemberKey"`
	// The MembershipRoles that apply to the Membership.
	// Must not contain duplicate MembershipRoles with the same name.
	// Structure is documented below.
	Roles []GroupMembershipRole `pulumi:"roles"`
	// The type of the membership.
	Type *string `pulumi:"type"`
	// The time when the Membership was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type GroupMembershipState struct {
	// The time when the Membership was created.
	CreateTime pulumi.StringPtrInput
	// The name of the Group to create this membership in.
	Group pulumi.StringPtrInput
	// EntityKey of the member.
	MemberKey GroupMembershipMemberKeyPtrInput
	// The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
	// Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
	Name pulumi.StringPtrInput
	// EntityKey of the member.
	// Structure is documented below.
	PreferredMemberKey GroupMembershipPreferredMemberKeyPtrInput
	// The MembershipRoles that apply to the Membership.
	// Must not contain duplicate MembershipRoles with the same name.
	// Structure is documented below.
	Roles GroupMembershipRoleArrayInput
	// The type of the membership.
	Type pulumi.StringPtrInput
	// The time when the Membership was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (GroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipState)(nil)).Elem()
}

type groupMembershipArgs struct {
	// The name of the Group to create this membership in.
	Group string `pulumi:"group"`
	// EntityKey of the member.
	MemberKey *GroupMembershipMemberKey `pulumi:"memberKey"`
	// EntityKey of the member.
	// Structure is documented below.
	PreferredMemberKey *GroupMembershipPreferredMemberKey `pulumi:"preferredMemberKey"`
	// The MembershipRoles that apply to the Membership.
	// Must not contain duplicate MembershipRoles with the same name.
	// Structure is documented below.
	Roles []GroupMembershipRole `pulumi:"roles"`
}

// The set of arguments for constructing a GroupMembership resource.
type GroupMembershipArgs struct {
	// The name of the Group to create this membership in.
	Group pulumi.StringInput
	// EntityKey of the member.
	MemberKey GroupMembershipMemberKeyPtrInput
	// EntityKey of the member.
	// Structure is documented below.
	PreferredMemberKey GroupMembershipPreferredMemberKeyPtrInput
	// The MembershipRoles that apply to the Membership.
	// Must not contain duplicate MembershipRoles with the same name.
	// Structure is documented below.
	Roles GroupMembershipRoleArrayInput
}

func (GroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipArgs)(nil)).Elem()
}

type GroupMembershipInput interface {
	pulumi.Input

	ToGroupMembershipOutput() GroupMembershipOutput
	ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput
}

func (GroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembership)(nil)).Elem()
}

func (i GroupMembership) ToGroupMembershipOutput() GroupMembershipOutput {
	return i.ToGroupMembershipOutputWithContext(context.Background())
}

func (i GroupMembership) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipOutput)
}

type GroupMembershipOutput struct {
	*pulumi.OutputState
}

func (GroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembershipOutput)(nil)).Elem()
}

func (o GroupMembershipOutput) ToGroupMembershipOutput() GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupMembershipOutput{})
}
