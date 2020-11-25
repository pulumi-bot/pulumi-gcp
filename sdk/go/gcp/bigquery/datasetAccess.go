// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// This resource does not support import.
type DatasetAccess struct {
	pulumi.CustomResourceState

	// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
	// stored in state as a different member type
	ApiUpdatedMember pulumi.BoolOutput `pulumi:"apiUpdatedMember"`
	// The ID of the dataset containing this table.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// An email address of a Google Group to grant access to.
	GroupByEmail pulumi.StringPtrOutput `pulumi:"groupByEmail"`
	// Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: `allUsers`
	IamMember pulumi.StringPtrOutput `pulumi:"iamMember"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles are
	// supported. Predefined roles that have equivalent basic roles are
	// swapped by the API to their basic counterparts, and will show a diff
	// post-create. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// A special group to grant access to. Possible values include:
	SpecialGroup pulumi.StringPtrOutput `pulumi:"specialGroup"`
	// An email address of a user to grant access to. For example:
	// fred@example.com
	UserByEmail pulumi.StringPtrOutput `pulumi:"userByEmail"`
	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	View DatasetAccessViewPtrOutput `pulumi:"view"`
}

// NewDatasetAccess registers a new resource with the given unique name, arguments, and options.
func NewDatasetAccess(ctx *pulumi.Context,
	name string, args *DatasetAccessArgs, opts ...pulumi.ResourceOption) (*DatasetAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	var resource DatasetAccess
	err := ctx.RegisterResource("gcp:bigquery/datasetAccess:DatasetAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetAccess gets an existing DatasetAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetAccessState, opts ...pulumi.ResourceOption) (*DatasetAccess, error) {
	var resource DatasetAccess
	err := ctx.ReadResource("gcp:bigquery/datasetAccess:DatasetAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetAccess resources.
type datasetAccessState struct {
	// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
	// stored in state as a different member type
	ApiUpdatedMember *bool `pulumi:"apiUpdatedMember"`
	// The ID of the dataset containing this table.
	DatasetId *string `pulumi:"datasetId"`
	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	Domain *string `pulumi:"domain"`
	// An email address of a Google Group to grant access to.
	GroupByEmail *string `pulumi:"groupByEmail"`
	// Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: `allUsers`
	IamMember *string `pulumi:"iamMember"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles are
	// supported. Predefined roles that have equivalent basic roles are
	// swapped by the API to their basic counterparts, and will show a diff
	// post-create. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	Role *string `pulumi:"role"`
	// A special group to grant access to. Possible values include:
	SpecialGroup *string `pulumi:"specialGroup"`
	// An email address of a user to grant access to. For example:
	// fred@example.com
	UserByEmail *string `pulumi:"userByEmail"`
	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	View *DatasetAccessView `pulumi:"view"`
}

type DatasetAccessState struct {
	// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
	// stored in state as a different member type
	ApiUpdatedMember pulumi.BoolPtrInput
	// The ID of the dataset containing this table.
	DatasetId pulumi.StringPtrInput
	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	Domain pulumi.StringPtrInput
	// An email address of a Google Group to grant access to.
	GroupByEmail pulumi.StringPtrInput
	// Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: `allUsers`
	IamMember pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles are
	// supported. Predefined roles that have equivalent basic roles are
	// swapped by the API to their basic counterparts, and will show a diff
	// post-create. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	Role pulumi.StringPtrInput
	// A special group to grant access to. Possible values include:
	SpecialGroup pulumi.StringPtrInput
	// An email address of a user to grant access to. For example:
	// fred@example.com
	UserByEmail pulumi.StringPtrInput
	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	View DatasetAccessViewPtrInput
}

func (DatasetAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetAccessState)(nil)).Elem()
}

type datasetAccessArgs struct {
	// The ID of the dataset containing this table.
	DatasetId string `pulumi:"datasetId"`
	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	Domain *string `pulumi:"domain"`
	// An email address of a Google Group to grant access to.
	GroupByEmail *string `pulumi:"groupByEmail"`
	// Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: `allUsers`
	IamMember *string `pulumi:"iamMember"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles are
	// supported. Predefined roles that have equivalent basic roles are
	// swapped by the API to their basic counterparts, and will show a diff
	// post-create. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	Role *string `pulumi:"role"`
	// A special group to grant access to. Possible values include:
	SpecialGroup *string `pulumi:"specialGroup"`
	// An email address of a user to grant access to. For example:
	// fred@example.com
	UserByEmail *string `pulumi:"userByEmail"`
	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	View *DatasetAccessView `pulumi:"view"`
}

// The set of arguments for constructing a DatasetAccess resource.
type DatasetAccessArgs struct {
	// The ID of the dataset containing this table.
	DatasetId pulumi.StringInput
	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	Domain pulumi.StringPtrInput
	// An email address of a Google Group to grant access to.
	GroupByEmail pulumi.StringPtrInput
	// Some other type of member that appears in the IAM Policy but isn't a user,
	// group, domain, or special group. For example: `allUsers`
	IamMember pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles are
	// supported. Predefined roles that have equivalent basic roles are
	// swapped by the API to their basic counterparts, and will show a diff
	// post-create. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	Role pulumi.StringPtrInput
	// A special group to grant access to. Possible values include:
	SpecialGroup pulumi.StringPtrInput
	// An email address of a user to grant access to. For example:
	// fred@example.com
	UserByEmail pulumi.StringPtrInput
	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	View DatasetAccessViewPtrInput
}

func (DatasetAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetAccessArgs)(nil)).Elem()
}

type DatasetAccessInput interface {
	pulumi.Input

	ToDatasetAccessOutput() DatasetAccessOutput
	ToDatasetAccessOutputWithContext(ctx context.Context) DatasetAccessOutput
}

func (DatasetAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetAccess)(nil)).Elem()
}

func (i DatasetAccess) ToDatasetAccessOutput() DatasetAccessOutput {
	return i.ToDatasetAccessOutputWithContext(context.Background())
}

func (i DatasetAccess) ToDatasetAccessOutputWithContext(ctx context.Context) DatasetAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetAccessOutput)
}

type DatasetAccessOutput struct {
	*pulumi.OutputState
}

func (DatasetAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetAccessOutput)(nil)).Elem()
}

func (o DatasetAccessOutput) ToDatasetAccessOutput() DatasetAccessOutput {
	return o
}

func (o DatasetAccessOutput) ToDatasetAccessOutputWithContext(ctx context.Context) DatasetAccessOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetAccessOutput{})
}
