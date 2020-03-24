// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package datastore

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Describes a composite index for Cloud Datastore.
//
//
// To get more information about Index, see:
//
// * [API documentation](https://cloud.google.com/datastore/docs/reference/admin/rest/v1/projects.indexes)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/datastore/docs/concepts/indexes)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/datastore_index.html.markdown.
type DataStoreIndex struct {
	pulumi.CustomResourceState

	// Policy for including ancestors in the index. Either 'ALL_ANCESTORS' or 'NONE', the default is 'NONE'.
	Ancestor pulumi.StringPtrOutput `pulumi:"ancestor"`
	// The index id.
	IndexId pulumi.StringOutput `pulumi:"indexId"`
	// The entity kind which the index applies to.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// An ordered list of properties to index on.
	Properties DataStoreIndexPropertyArrayOutput `pulumi:"properties"`
}

// NewDataStoreIndex registers a new resource with the given unique name, arguments, and options.
func NewDataStoreIndex(ctx *pulumi.Context,
	name string, args *DataStoreIndexArgs, opts ...pulumi.ResourceOption) (*DataStoreIndex, error) {
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil {
		args = &DataStoreIndexArgs{}
	}
	var resource DataStoreIndex
	err := ctx.RegisterResource("gcp:datastore/dataStoreIndex:DataStoreIndex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataStoreIndex gets an existing DataStoreIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataStoreIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataStoreIndexState, opts ...pulumi.ResourceOption) (*DataStoreIndex, error) {
	var resource DataStoreIndex
	err := ctx.ReadResource("gcp:datastore/dataStoreIndex:DataStoreIndex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataStoreIndex resources.
type dataStoreIndexState struct {
	// Policy for including ancestors in the index. Either 'ALL_ANCESTORS' or 'NONE', the default is 'NONE'.
	Ancestor *string `pulumi:"ancestor"`
	// The index id.
	IndexId *string `pulumi:"indexId"`
	// The entity kind which the index applies to.
	Kind *string `pulumi:"kind"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An ordered list of properties to index on.
	Properties []DataStoreIndexProperty `pulumi:"properties"`
}

type DataStoreIndexState struct {
	// Policy for including ancestors in the index. Either 'ALL_ANCESTORS' or 'NONE', the default is 'NONE'.
	Ancestor pulumi.StringPtrInput
	// The index id.
	IndexId pulumi.StringPtrInput
	// The entity kind which the index applies to.
	Kind pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An ordered list of properties to index on.
	Properties DataStoreIndexPropertyArrayInput
}

func (DataStoreIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataStoreIndexState)(nil)).Elem()
}

type dataStoreIndexArgs struct {
	// Policy for including ancestors in the index. Either 'ALL_ANCESTORS' or 'NONE', the default is 'NONE'.
	Ancestor *string `pulumi:"ancestor"`
	// The entity kind which the index applies to.
	Kind string `pulumi:"kind"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An ordered list of properties to index on.
	Properties []DataStoreIndexProperty `pulumi:"properties"`
}

// The set of arguments for constructing a DataStoreIndex resource.
type DataStoreIndexArgs struct {
	// Policy for including ancestors in the index. Either 'ALL_ANCESTORS' or 'NONE', the default is 'NONE'.
	Ancestor pulumi.StringPtrInput
	// The entity kind which the index applies to.
	Kind pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An ordered list of properties to index on.
	Properties DataStoreIndexPropertyArrayInput
}

func (DataStoreIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataStoreIndexArgs)(nil)).Elem()
}

