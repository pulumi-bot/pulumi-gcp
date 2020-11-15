// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Traffic routing configuration for versions within a single service. Traffic splits define how traffic directed to the service is assigned to versions.
//
// To get more information about ServiceSplitTraffic, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services)
//
// ## Example Usage
//
// ## Import
//
// ServiceSplitTraffic can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:appengine/engineSplitTraffic:EngineSplitTraffic default apps/{{project}}/services/{{service}}
// ```
//
// ```sh
//  $ pulumi import gcp:appengine/engineSplitTraffic:EngineSplitTraffic default {{project}}/{{service}}
// ```
//
// ```sh
//  $ pulumi import gcp:appengine/engineSplitTraffic:EngineSplitTraffic default {{service}}
// ```
type EngineSplitTraffic struct {
	pulumi.CustomResourceState

	// If set to true traffic will be migrated to this version.
	MigrateTraffic pulumi.BoolPtrOutput `pulumi:"migrateTraffic"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the service these settings apply to.
	Service pulumi.StringOutput `pulumi:"service"`
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split EngineSplitTrafficSplitOutput `pulumi:"split"`
}

// NewEngineSplitTraffic registers a new resource with the given unique name, arguments, and options.
func NewEngineSplitTraffic(ctx *pulumi.Context,
	name string, args *EngineSplitTrafficArgs, opts ...pulumi.ResourceOption) (*EngineSplitTraffic, error) {
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil || args.Split == nil {
		return nil, errors.New("missing required argument 'Split'")
	}
	if args == nil {
		args = &EngineSplitTrafficArgs{}
	}
	var resource EngineSplitTraffic
	err := ctx.RegisterResource("gcp:appengine/engineSplitTraffic:EngineSplitTraffic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEngineSplitTraffic gets an existing EngineSplitTraffic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEngineSplitTraffic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EngineSplitTrafficState, opts ...pulumi.ResourceOption) (*EngineSplitTraffic, error) {
	var resource EngineSplitTraffic
	err := ctx.ReadResource("gcp:appengine/engineSplitTraffic:EngineSplitTraffic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EngineSplitTraffic resources.
type engineSplitTrafficState struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic *bool `pulumi:"migrateTraffic"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The name of the service these settings apply to.
	Service *string `pulumi:"service"`
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split *EngineSplitTrafficSplit `pulumi:"split"`
}

type EngineSplitTrafficState struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The name of the service these settings apply to.
	Service pulumi.StringPtrInput
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split EngineSplitTrafficSplitPtrInput
}

func (EngineSplitTrafficState) ElementType() reflect.Type {
	return reflect.TypeOf((*engineSplitTrafficState)(nil)).Elem()
}

type engineSplitTrafficArgs struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic *bool `pulumi:"migrateTraffic"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The name of the service these settings apply to.
	Service string `pulumi:"service"`
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split EngineSplitTrafficSplit `pulumi:"split"`
}

// The set of arguments for constructing a EngineSplitTraffic resource.
type EngineSplitTrafficArgs struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The name of the service these settings apply to.
	Service pulumi.StringInput
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split EngineSplitTrafficSplitInput
}

func (EngineSplitTrafficArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*engineSplitTrafficArgs)(nil)).Elem()
}

type EngineSplitTrafficInput interface {
	pulumi.Input

	ToEngineSplitTrafficOutput() EngineSplitTrafficOutput
	ToEngineSplitTrafficOutputWithContext(ctx context.Context) EngineSplitTrafficOutput
}

func (EngineSplitTraffic) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineSplitTraffic)(nil)).Elem()
}

func (i EngineSplitTraffic) ToEngineSplitTrafficOutput() EngineSplitTrafficOutput {
	return i.ToEngineSplitTrafficOutputWithContext(context.Background())
}

func (i EngineSplitTraffic) ToEngineSplitTrafficOutputWithContext(ctx context.Context) EngineSplitTrafficOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineSplitTrafficOutput)
}

type EngineSplitTrafficOutput struct {
	*pulumi.OutputState
}

func (EngineSplitTrafficOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineSplitTrafficOutput)(nil)).Elem()
}

func (o EngineSplitTrafficOutput) ToEngineSplitTrafficOutput() EngineSplitTrafficOutput {
	return o
}

func (o EngineSplitTrafficOutput) ToEngineSplitTrafficOutputWithContext(ctx context.Context) EngineSplitTrafficOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EngineSplitTrafficOutput{})
}
