// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A source representation instance is a Cloud SQL instance that represents
// the source database server to the Cloud SQL replica. It is visible in the
// Cloud Console and appears the same as a regular Cloud SQL instance, but it
// contains no data, requires no configuration or maintenance, and does not
// affect billing. You cannot update the source representation instance.
//
//
//
// ## Example Usage
//
// ### Sql Source Representation Instance Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/sql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = sql.NewSourceRepresentationInstance(ctx, "instance", &sql.SourceRepresentationInstanceArgs{
// 			DatabaseVersion: pulumi.String("MYSQL_5_7"),
// 			Host:            pulumi.String("10.20.30.40"),
// 			Port:            pulumi.Int(3306),
// 			Region:          pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SourceRepresentationInstance struct {
	pulumi.CustomResourceState

	// The MySQL version running on your source database server.
	DatabaseVersion pulumi.StringOutput `pulumi:"databaseVersion"`
	// The externally accessible IPv4 address for the source database server.
	Host pulumi.StringOutput `pulumi:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewSourceRepresentationInstance registers a new resource with the given unique name, arguments, and options.
func NewSourceRepresentationInstance(ctx *pulumi.Context,
	name string, args *SourceRepresentationInstanceArgs, opts ...pulumi.ResourceOption) (*SourceRepresentationInstance, error) {
	if args == nil || args.DatabaseVersion == nil {
		return nil, errors.New("missing required argument 'DatabaseVersion'")
	}
	if args == nil || args.Host == nil {
		return nil, errors.New("missing required argument 'Host'")
	}
	if args == nil {
		args = &SourceRepresentationInstanceArgs{}
	}
	var resource SourceRepresentationInstance
	err := ctx.RegisterResource("gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceRepresentationInstance gets an existing SourceRepresentationInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceRepresentationInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceRepresentationInstanceState, opts ...pulumi.ResourceOption) (*SourceRepresentationInstance, error) {
	var resource SourceRepresentationInstance
	err := ctx.ReadResource("gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceRepresentationInstance resources.
type sourceRepresentationInstanceState struct {
	// The MySQL version running on your source database server.
	DatabaseVersion *string `pulumi:"databaseVersion"`
	// The externally accessible IPv4 address for the source database server.
	Host *string `pulumi:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name *string `pulumi:"name"`
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

type SourceRepresentationInstanceState struct {
	// The MySQL version running on your source database server.
	DatabaseVersion pulumi.StringPtrInput
	// The externally accessible IPv4 address for the source database server.
	Host pulumi.StringPtrInput
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name pulumi.StringPtrInput
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
}

func (SourceRepresentationInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRepresentationInstanceState)(nil)).Elem()
}

type sourceRepresentationInstanceArgs struct {
	// The MySQL version running on your source database server.
	DatabaseVersion string `pulumi:"databaseVersion"`
	// The externally accessible IPv4 address for the source database server.
	Host string `pulumi:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name *string `pulumi:"name"`
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a SourceRepresentationInstance resource.
type SourceRepresentationInstanceArgs struct {
	// The MySQL version running on your source database server.
	DatabaseVersion pulumi.StringInput
	// The externally accessible IPv4 address for the source database server.
	Host pulumi.StringInput
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name pulumi.StringPtrInput
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
}

func (SourceRepresentationInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceRepresentationInstanceArgs)(nil)).Elem()
}
