// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
//
// * `spanner.DatabaseIAMPolicy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
//
// > **Warning:** It's entirely possibly to lock yourself out of your database using `spanner.DatabaseIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
//
// * `spanner.DatabaseIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
// * `spanner.DatabaseIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
//
// > **Note:** `spanner.DatabaseIAMPolicy` **cannot** be used in conjunction with `spanner.DatabaseIAMBinding` and `spanner.DatabaseIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `spanner.DatabaseIAMBinding` resources **can be** used in conjunction with `spanner.DatabaseIAMMember` resources **only if** they do not grant privilege to the same role.
type DatabaseIAMBinding struct {
	pulumi.CustomResourceState

	Condition DatabaseIAMBindingConditionPtrOutput `pulumi:"condition"`
	// The name of the Spanner database.
	Database pulumi.StringOutput `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringOutput      `pulumi:"instance"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewDatabaseIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewDatabaseIAMBinding(ctx *pulumi.Context,
	name string, args *DatabaseIAMBindingArgs, opts ...pulumi.ResourceOption) (*DatabaseIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource DatabaseIAMBinding
	err := ctx.RegisterResource("gcp:spanner/databaseIAMBinding:DatabaseIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseIAMBinding gets an existing DatabaseIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseIAMBindingState, opts ...pulumi.ResourceOption) (*DatabaseIAMBinding, error) {
	var resource DatabaseIAMBinding
	err := ctx.ReadResource("gcp:spanner/databaseIAMBinding:DatabaseIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseIAMBinding resources.
type databaseIAMBindingState struct {
	Condition *DatabaseIAMBindingCondition `pulumi:"condition"`
	// The name of the Spanner database.
	Database *string `pulumi:"database"`
	// (Computed) The etag of the database's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the Spanner instance the database belongs to.
	Instance *string  `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type DatabaseIAMBindingState struct {
	Condition DatabaseIAMBindingConditionPtrInput
	// The name of the Spanner database.
	Database pulumi.StringPtrInput
	// (Computed) The etag of the database's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (DatabaseIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMBindingState)(nil)).Elem()
}

type databaseIAMBindingArgs struct {
	Condition *DatabaseIAMBindingCondition `pulumi:"condition"`
	// The name of the Spanner database.
	Database string `pulumi:"database"`
	// The name of the Spanner instance the database belongs to.
	Instance string   `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a DatabaseIAMBinding resource.
type DatabaseIAMBindingArgs struct {
	Condition DatabaseIAMBindingConditionPtrInput
	// The name of the Spanner database.
	Database pulumi.StringInput
	// The name of the Spanner instance the database belongs to.
	Instance pulumi.StringInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (DatabaseIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseIAMBindingArgs)(nil)).Elem()
}

type DatabaseIAMBindingInput interface {
	pulumi.Input

	ToDatabaseIAMBindingOutput() DatabaseIAMBindingOutput
	ToDatabaseIAMBindingOutputWithContext(ctx context.Context) DatabaseIAMBindingOutput
}

func (DatabaseIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIAMBinding)(nil)).Elem()
}

func (i DatabaseIAMBinding) ToDatabaseIAMBindingOutput() DatabaseIAMBindingOutput {
	return i.ToDatabaseIAMBindingOutputWithContext(context.Background())
}

func (i DatabaseIAMBinding) ToDatabaseIAMBindingOutputWithContext(ctx context.Context) DatabaseIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIAMBindingOutput)
}

type DatabaseIAMBindingOutput struct {
	*pulumi.OutputState
}

func (DatabaseIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIAMBindingOutput)(nil)).Elem()
}

func (o DatabaseIAMBindingOutput) ToDatabaseIAMBindingOutput() DatabaseIAMBindingOutput {
	return o
}

func (o DatabaseIAMBindingOutput) ToDatabaseIAMBindingOutputWithContext(ctx context.Context) DatabaseIAMBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseIAMBindingOutput{})
}
