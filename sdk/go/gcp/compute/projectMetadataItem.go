// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a single key/value pair on metadata common to all instances for
// a project in GCE. Using `compute.ProjectMetadataItem` lets you
// manage a single key/value setting in the provider rather than the entire
// project metadata map.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewProjectMetadataItem(ctx, "_default", &compute.ProjectMetadataItemArgs{
// 			Key:   pulumi.String("my_metadata"),
// 			Value: pulumi.String("my_value"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProjectMetadataItem struct {
	pulumi.CustomResourceState

	// The metadata key to set.
	Key pulumi.StringOutput `pulumi:"key"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The value to set for the given metadata key.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewProjectMetadataItem registers a new resource with the given unique name, arguments, and options.
func NewProjectMetadataItem(ctx *pulumi.Context,
	name string, args *ProjectMetadataItemArgs, opts ...pulumi.ResourceOption) (*ProjectMetadataItem, error) {
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &ProjectMetadataItemArgs{}
	}
	var resource ProjectMetadataItem
	err := ctx.RegisterResource("gcp:compute/projectMetadataItem:ProjectMetadataItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMetadataItem gets an existing ProjectMetadataItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMetadataItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMetadataItemState, opts ...pulumi.ResourceOption) (*ProjectMetadataItem, error) {
	var resource ProjectMetadataItem
	err := ctx.ReadResource("gcp:compute/projectMetadataItem:ProjectMetadataItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMetadataItem resources.
type projectMetadataItemState struct {
	// The metadata key to set.
	Key *string `pulumi:"key"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The value to set for the given metadata key.
	Value *string `pulumi:"value"`
}

type ProjectMetadataItemState struct {
	// The metadata key to set.
	Key pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The value to set for the given metadata key.
	Value pulumi.StringPtrInput
}

func (ProjectMetadataItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMetadataItemState)(nil)).Elem()
}

type projectMetadataItemArgs struct {
	// The metadata key to set.
	Key string `pulumi:"key"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The value to set for the given metadata key.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ProjectMetadataItem resource.
type ProjectMetadataItemArgs struct {
	// The metadata key to set.
	Key pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The value to set for the given metadata key.
	Value pulumi.StringInput
}

func (ProjectMetadataItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMetadataItemArgs)(nil)).Elem()
}
