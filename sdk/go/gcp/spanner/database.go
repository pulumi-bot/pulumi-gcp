// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Cloud Spanner Database which is hosted on a Spanner instance.
//
// To get more information about Database, see:
//
// * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.databases)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/spanner/)
//
// > **Warning:** It is strongly recommended to set `lifecycle { preventDestroy = true }` on databases in order to prevent accidental data loss.
//
// ## Example Usage
// ### Spanner Database Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := spanner.NewInstance(ctx, "main", &spanner.InstanceArgs{
// 			Config:      pulumi.String("regional-europe-west1"),
// 			DisplayName: pulumi.String("main-instance"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = spanner.NewDatabase(ctx, "database", &spanner.DatabaseArgs{
// 			Instance: main.Name,
// 			Ddls: pulumi.StringArray{
// 				pulumi.String("CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)"),
// 				pulumi.String("CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)"),
// 			},
// 			DeletionProtection: pulumi.Bool(false),
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
// Database can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:spanner/database:Database default projects/{{project}}/instances/{{instance}}/databases/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:spanner/database:Database default instances/{{instance}}/databases/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:spanner/database:Database default {{project}}/{{instance}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:spanner/database:Database default {{instance}}/{{name}}
// ```
type Database struct {
	pulumi.CustomResourceState

	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	Ddls               pulumi.StringArrayOutput `pulumi:"ddls"`
	DeletionProtection pulumi.BoolPtrOutput     `pulumi:"deletionProtection"`
	// The instance to create the database on.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after
	// the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
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
	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	Ddls               []string `pulumi:"ddls"`
	DeletionProtection *bool    `pulumi:"deletionProtection"`
	// The instance to create the database on.
	Instance *string `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after
	// the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An explanation of the status of the database.
	State *string `pulumi:"state"`
}

type DatabaseState struct {
	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	Ddls               pulumi.StringArrayInput
	DeletionProtection pulumi.BoolPtrInput
	// The instance to create the database on.
	Instance pulumi.StringPtrInput
	// A unique identifier for the database, which cannot be changed after
	// the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
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
	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	Ddls               []string `pulumi:"ddls"`
	DeletionProtection *bool    `pulumi:"deletionProtection"`
	// The instance to create the database on.
	Instance string `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after
	// the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	Ddls               pulumi.StringArrayInput
	DeletionProtection pulumi.BoolPtrInput
	// The instance to create the database on.
	Instance pulumi.StringInput
	// A unique identifier for the database, which cannot be changed after
	// the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (Database) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil)).Elem()
}

func (i Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct {
	*pulumi.OutputState
}

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseOutput)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
