// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package spanner

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A Cloud Spanner Database which is hosted on a Spanner instance.
//
//
// To get more information about Database, see:
//
// * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.databases)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/spanner/)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database.html.markdown.
type Database struct {
	pulumi.CustomResourceState

	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls pulumi.StringArrayOutput `pulumi:"ddls"`
	// The instance to create the database on.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// [a-z][-a-z0-9]*[a-z0-9].
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// An explanation of the status of the database.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil {
		args = &DatabaseArgs{}
	}
	var resource Database
	err := ctx.RegisterResource("gcp:spanner/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("gcp:spanner/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls []string `pulumi:"ddls"`
	// The instance to create the database on.
	Instance *string `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// [a-z][-a-z0-9]*[a-z0-9].
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An explanation of the status of the database.
	State *string `pulumi:"state"`
}

type DatabaseState struct {
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls pulumi.StringArrayInput
	// The instance to create the database on.
	Instance pulumi.StringPtrInput
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// [a-z][-a-z0-9]*[a-z0-9].
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An explanation of the status of the database.
	State pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls []string `pulumi:"ddls"`
	// The instance to create the database on.
	Instance string `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// [a-z][-a-z0-9]*[a-z0-9].
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls pulumi.StringArrayInput
	// The instance to create the database on.
	Instance pulumi.StringInput
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// [a-z][-a-z0-9]*[a-z0-9].
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

