// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a RuntimeConfig resource in Google Cloud. For more information, see the
// [official documentation](https://cloud.google.com/deployment-manager/runtime-configurator/),
// or the
// [JSON API](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/).
//
// ## Example Usage
//
// Example creating a RuntimeConfig resource.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/runtimeconfig"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := runtimeconfig.NewConfig(ctx, "my_runtime_config", &runtimeconfig.ConfigArgs{
// 			Description: pulumi.String("Runtime configuration values for my service"),
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
// Runtime Configs can be imported using the `name` or full config name, e.g.
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/config:Config myconfig myconfig
// ```
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/config:Config myconfig projects/my-gcp-project/configs/myconfig
// ```
//
//  When importing using only the name, the provider project must be set.
type Config struct {
	pulumi.CustomResourceState

	// The description to associate with the runtime
	// config.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the runtime config.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConfig registers a new resource with the given unique name, arguments, and options.
func NewConfig(ctx *pulumi.Context,
	name string, args *ConfigArgs, opts ...pulumi.ResourceOption) (*Config, error) {
	if args == nil {
		args = &ConfigArgs{}
	}

	var resource Config
	err := ctx.RegisterResource("gcp:runtimeconfig/config:Config", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfig gets an existing Config resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigState, opts ...pulumi.ResourceOption) (*Config, error) {
	var resource Config
	err := ctx.ReadResource("gcp:runtimeconfig/config:Config", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Config resources.
type configState struct {
	// The description to associate with the runtime
	// config.
	Description *string `pulumi:"description"`
	// The name of the runtime config.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ConfigState struct {
	// The description to associate with the runtime
	// config.
	Description pulumi.StringPtrInput
	// The name of the runtime config.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*configState)(nil)).Elem()
}

type configArgs struct {
	// The description to associate with the runtime
	// config.
	Description *string `pulumi:"description"`
	// The name of the runtime config.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	// The description to associate with the runtime
	// config.
	Description pulumi.StringPtrInput
	// The name of the runtime config.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configArgs)(nil)).Elem()
}

type ConfigInput interface {
	pulumi.Input

	ToConfigOutput() ConfigOutput
	ToConfigOutputWithContext(ctx context.Context) ConfigOutput
}

func (Config) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil)).Elem()
}

func (i Config) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i Config) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

type ConfigOutput struct {
	*pulumi.OutputState
}

func (ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigOutput)(nil)).Elem()
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigOutput{})
}
