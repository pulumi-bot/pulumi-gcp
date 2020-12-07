// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package firestore

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Cloud Firestore indexes enable simple and complex queries against documents in a database.
//  This resource manages composite indexes and not single
// field indexes.
//
// To get more information about Index, see:
//
// * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.collectionGroups.indexes)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/firestore/docs/query-data/indexing)
//
// > **Warning:** This resource creates a Firestore Index on a project that already has
// Firestore enabled. If you haven't already enabled it, you can create a
// `appengine.Application` resource with `databaseType` set to
// `"CLOUD_FIRESTORE"` to do so. Your Firestore location will be the same as
// the App Engine location specified.
//
// ## Example Usage
// ### Firestore Index Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/firestore"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := firestore.NewIndex(ctx, "my_index", &firestore.IndexArgs{
// 			Collection: pulumi.String("chatrooms"),
// 			Fields: firestore.IndexFieldArray{
// 				&firestore.IndexFieldArgs{
// 					FieldPath: pulumi.String("name"),
// 					Order:     pulumi.String("ASCENDING"),
// 				},
// 				&firestore.IndexFieldArgs{
// 					FieldPath: pulumi.String("description"),
// 					Order:     pulumi.String("DESCENDING"),
// 				},
// 				&firestore.IndexFieldArgs{
// 					FieldPath: pulumi.String("__name__"),
// 					Order:     pulumi.String("DESCENDING"),
// 				},
// 			},
// 			Project: pulumi.String("my-project-name"),
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
// Index can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:firestore/index:Index default {{name}}
// ```
type Index struct {
	pulumi.CustomResourceState

	// The collection being indexed.
	Collection pulumi.StringOutput `pulumi:"collection"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// The fields supported by this index. The last field entry is always for
	// the field path `__name__`. If, on creation, `__name__` was not
	// specified as the last field, it will be added automatically with the
	// same direction as that of the last field defined. If the final field
	// in a composite index is not directional, the `__name__` will be
	// ordered `"ASCENDING"` (unless explicitly specified otherwise).
	// Structure is documented below.
	Fields IndexFieldArrayOutput `pulumi:"fields"`
	// A server defined name for this index. Format:
	// 'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The scope at which a query is run.
	// Default value is `COLLECTION`.
	// Possible values are `COLLECTION` and `COLLECTION_GROUP`.
	QueryScope pulumi.StringPtrOutput `pulumi:"queryScope"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Collection == nil {
		return nil, errors.New("invalid value for required argument 'Collection'")
	}
	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	var resource Index
	err := ctx.RegisterResource("gcp:firestore/index:Index", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndex gets an existing Index resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexState, opts ...pulumi.ResourceOption) (*Index, error) {
	var resource Index
	err := ctx.ReadResource("gcp:firestore/index:Index", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Index resources.
type indexState struct {
	// The collection being indexed.
	Collection *string `pulumi:"collection"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database *string `pulumi:"database"`
	// The fields supported by this index. The last field entry is always for
	// the field path `__name__`. If, on creation, `__name__` was not
	// specified as the last field, it will be added automatically with the
	// same direction as that of the last field defined. If the final field
	// in a composite index is not directional, the `__name__` will be
	// ordered `"ASCENDING"` (unless explicitly specified otherwise).
	// Structure is documented below.
	Fields []IndexField `pulumi:"fields"`
	// A server defined name for this index. Format:
	// 'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The scope at which a query is run.
	// Default value is `COLLECTION`.
	// Possible values are `COLLECTION` and `COLLECTION_GROUP`.
	QueryScope *string `pulumi:"queryScope"`
}

type IndexState struct {
	// The collection being indexed.
	Collection pulumi.StringPtrInput
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrInput
	// The fields supported by this index. The last field entry is always for
	// the field path `__name__`. If, on creation, `__name__` was not
	// specified as the last field, it will be added automatically with the
	// same direction as that of the last field defined. If the final field
	// in a composite index is not directional, the `__name__` will be
	// ordered `"ASCENDING"` (unless explicitly specified otherwise).
	// Structure is documented below.
	Fields IndexFieldArrayInput
	// A server defined name for this index. Format:
	// 'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The scope at which a query is run.
	// Default value is `COLLECTION`.
	// Possible values are `COLLECTION` and `COLLECTION_GROUP`.
	QueryScope pulumi.StringPtrInput
}

func (IndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexState)(nil)).Elem()
}

type indexArgs struct {
	// The collection being indexed.
	Collection string `pulumi:"collection"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database *string `pulumi:"database"`
	// The fields supported by this index. The last field entry is always for
	// the field path `__name__`. If, on creation, `__name__` was not
	// specified as the last field, it will be added automatically with the
	// same direction as that of the last field defined. If the final field
	// in a composite index is not directional, the `__name__` will be
	// ordered `"ASCENDING"` (unless explicitly specified otherwise).
	// Structure is documented below.
	Fields []IndexField `pulumi:"fields"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The scope at which a query is run.
	// Default value is `COLLECTION`.
	// Possible values are `COLLECTION` and `COLLECTION_GROUP`.
	QueryScope *string `pulumi:"queryScope"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// The collection being indexed.
	Collection pulumi.StringInput
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrInput
	// The fields supported by this index. The last field entry is always for
	// the field path `__name__`. If, on creation, `__name__` was not
	// specified as the last field, it will be added automatically with the
	// same direction as that of the last field defined. If the final field
	// in a composite index is not directional, the `__name__` will be
	// ordered `"ASCENDING"` (unless explicitly specified otherwise).
	// Structure is documented below.
	Fields IndexFieldArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The scope at which a query is run.
	// Default value is `COLLECTION`.
	// Possible values are `COLLECTION` and `COLLECTION_GROUP`.
	QueryScope pulumi.StringPtrInput
}

func (IndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexArgs)(nil)).Elem()
}

type IndexInput interface {
	pulumi.Input

	ToIndexOutput() IndexOutput
	ToIndexOutputWithContext(ctx context.Context) IndexOutput
}

func (Index) ElementType() reflect.Type {
	return reflect.TypeOf((*Index)(nil)).Elem()
}

func (i Index) ToIndexOutput() IndexOutput {
	return i.ToIndexOutputWithContext(context.Background())
}

func (i Index) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexOutput)
}

type IndexOutput struct {
	*pulumi.OutputState
}

func (IndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexOutput)(nil)).Elem()
}

func (o IndexOutput) ToIndexOutput() IndexOutput {
	return o
}

func (o IndexOutput) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IndexOutput{})
}
