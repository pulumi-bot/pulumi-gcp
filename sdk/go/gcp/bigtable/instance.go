// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a Google Bigtable instance. For more information see
// [the official documentation](https://cloud.google.com/bigtable/) and
// [API](https://cloud.google.com/bigtable/docs/go/reference).
//
// ## Example Usage
type Instance struct {
	pulumi.CustomResourceState

	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters InstanceClusterArrayOutput `pulumi:"clusters"`
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	var resource Instance
	err := ctx.RegisterResource("gcp:bigtable/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:bigtable/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters []InstanceCluster `pulumi:"clusters"`
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName *string `pulumi:"displayName"`
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType *string `pulumi:"instanceType"`
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]string `pulumi:"labels"`
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type InstanceState struct {
	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters InstanceClusterArrayInput
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection pulumi.BoolPtrInput
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName pulumi.StringPtrInput
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.StringMapInput
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters []InstanceCluster `pulumi:"clusters"`
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName *string `pulumi:"displayName"`
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType *string `pulumi:"instanceType"`
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]string `pulumi:"labels"`
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// A block of cluster configuration options. This can be specified at least once, and up to 4 times.
	// See structure below.
	Clusters InstanceClusterArrayInput
	// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
	// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
	DeletionProtection pulumi.BoolPtrInput
	// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
	DisplayName pulumi.StringPtrInput
	// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
	// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
	// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
	// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
	// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
	//
	// Deprecated: It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels pulumi.StringMapInput
	// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil)).Elem()
}

func (i Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceOutput)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}
