// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a RuntimeConfig variable in Google Cloud. For more information, see the
// [official documentation](https://cloud.google.com/deployment-manager/runtime-configurator/),
// or the
// [JSON API](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/).
//
// ## Import
//
// Runtime Config Variables can be imported using the `name` or full variable name, e.g.
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/variable:Variable myvariable myconfig/myvariable
// ```
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/variable:Variable myvariable projects/my-gcp-project/configs/myconfig/variables/myvariable
// ```
//
//  When importing using only the name, the provider project must be set.
type Variable struct {
	pulumi.CustomResourceState

	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text pulumi.StringPtrOutput `pulumi:"text"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds, representing when the variable was last updated.
	// Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringOutput    `pulumi:"updateTime"`
	Value      pulumi.StringPtrOutput `pulumi:"value"`
}

// NewVariable registers a new resource with the given unique name, arguments, and options.
func NewVariable(ctx *pulumi.Context,
	name string, args *VariableArgs, opts ...pulumi.ResourceOption) (*Variable, error) {
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil {
		args = &VariableArgs{}
	}
	var resource Variable
	err := ctx.RegisterResource("gcp:runtimeconfig/variable:Variable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVariable gets an existing Variable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableState, opts ...pulumi.ResourceOption) (*Variable, error) {
	var resource Variable
	err := ctx.ReadResource("gcp:runtimeconfig/variable:Variable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Variable resources.
type variableState struct {
	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name *string `pulumi:"name"`
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent *string `pulumi:"parent"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text *string `pulumi:"text"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds, representing when the variable was last updated.
	// Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime *string `pulumi:"updateTime"`
	Value      *string `pulumi:"value"`
}

type VariableState struct {
	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name pulumi.StringPtrInput
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text pulumi.StringPtrInput
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds, representing when the variable was last updated.
	// Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringPtrInput
	Value      pulumi.StringPtrInput
}

func (VariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableState)(nil)).Elem()
}

type variableArgs struct {
	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name *string `pulumi:"name"`
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent string `pulumi:"parent"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text  *string `pulumi:"text"`
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a Variable resource.
type VariableArgs struct {
	// The name of the variable to manage. Note that variable
	// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	Name pulumi.StringPtrInput
	// The name of the RuntimeConfig resource containing this
	// variable.
	Parent pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// or `value` - (Required) The content to associate with the variable.
	// Exactly one of `text` or `variable` must be specified. If `text` is specified,
	// it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	// is specified, it must be base64 encoded and less than 4096 bytes in length.
	Text  pulumi.StringPtrInput
	Value pulumi.StringPtrInput
}

func (VariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableArgs)(nil)).Elem()
}

type VariableInput interface {
	pulumi.Input

	ToVariableOutput() VariableOutput
	ToVariableOutputWithContext(ctx context.Context) VariableOutput
}

func (Variable) ElementType() reflect.Type {
	return reflect.TypeOf((*Variable)(nil)).Elem()
}

func (i Variable) ToVariableOutput() VariableOutput {
	return i.ToVariableOutputWithContext(context.Background())
}

func (i Variable) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableOutput)
}

type VariableOutput struct {
	*pulumi.OutputState
}

func (VariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableOutput)(nil)).Elem()
}

func (o VariableOutput) ToVariableOutput() VariableOutput {
	return o
}

func (o VariableOutput) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VariableOutput{})
}
