// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The description of a dynamic collection of monitored resources. Each group
// has a filter that is matched against monitored resources and their
// associated metadata. If a group's filter matches an available monitored
// resource, then that resource is a member of that group.
//
// To get more information about Group, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.groups)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/monitoring/groups/)
//
// ## Example Usage
//
// ## Import
//
// Group can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:monitoring/group:Group default {{name}}
// ```
type Group struct {
	pulumi.CustomResourceState

	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The filter used to determine which monitored resources
	// belong to this group.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster pulumi.BoolPtrOutput `pulumi:"isCluster"`
	// A unique identifier for this group. The format is "projects/{project_id_or_number}/groups/{group_id}".
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName pulumi.StringPtrOutput `pulumi:"parentName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Filter == nil {
		return nil, errors.New("missing required argument 'Filter'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("gcp:monitoring/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("gcp:monitoring/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName *string `pulumi:"displayName"`
	// The filter used to determine which monitored resources
	// belong to this group.
	Filter *string `pulumi:"filter"`
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster *bool `pulumi:"isCluster"`
	// A unique identifier for this group. The format is "projects/{project_id_or_number}/groups/{group_id}".
	Name *string `pulumi:"name"`
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName *string `pulumi:"parentName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type GroupState struct {
	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName pulumi.StringPtrInput
	// The filter used to determine which monitored resources
	// belong to this group.
	Filter pulumi.StringPtrInput
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster pulumi.BoolPtrInput
	// A unique identifier for this group. The format is "projects/{project_id_or_number}/groups/{group_id}".
	Name pulumi.StringPtrInput
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName string `pulumi:"displayName"`
	// The filter used to determine which monitored resources
	// belong to this group.
	Filter string `pulumi:"filter"`
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster *bool `pulumi:"isCluster"`
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName *string `pulumi:"parentName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName pulumi.StringInput
	// The filter used to determine which monitored resources
	// belong to this group.
	Filter pulumi.StringInput
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster pulumi.BoolPtrInput
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil)).Elem()
}

func (i Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupOutput)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
