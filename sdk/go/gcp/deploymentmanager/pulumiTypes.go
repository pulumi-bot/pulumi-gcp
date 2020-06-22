// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deploymentmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type DeploymentLabel struct {
	// Key for label.
	Key *string `pulumi:"key"`
	// Value of label.
	Value *string `pulumi:"value"`
}

// DeploymentLabelInput is an input type that accepts DeploymentLabelArgs and DeploymentLabelOutput values.
// You can construct a concrete instance of `DeploymentLabelInput` via:
//
//          DeploymentLabelArgs{...}
type DeploymentLabelInput interface {
	pulumi.Input

	ToDeploymentLabelOutput() DeploymentLabelOutput
	ToDeploymentLabelOutputWithContext(context.Context) DeploymentLabelOutput
}

type DeploymentLabelArgs struct {
	// Key for label.
	Key pulumi.StringPtrInput `pulumi:"key"`
	// Value of label.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (DeploymentLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentLabel)(nil)).Elem()
}

func (i DeploymentLabelArgs) ToDeploymentLabelOutput() DeploymentLabelOutput {
	return i.ToDeploymentLabelOutputWithContext(context.Background())
}

func (i DeploymentLabelArgs) ToDeploymentLabelOutputWithContext(ctx context.Context) DeploymentLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentLabelOutput)
}

// DeploymentLabelArrayInput is an input type that accepts DeploymentLabelArray and DeploymentLabelArrayOutput values.
// You can construct a concrete instance of `DeploymentLabelArrayInput` via:
//
//          DeploymentLabelArray{ DeploymentLabelArgs{...} }
type DeploymentLabelArrayInput interface {
	pulumi.Input

	ToDeploymentLabelArrayOutput() DeploymentLabelArrayOutput
	ToDeploymentLabelArrayOutputWithContext(context.Context) DeploymentLabelArrayOutput
}

type DeploymentLabelArray []DeploymentLabelInput

func (DeploymentLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentLabel)(nil)).Elem()
}

func (i DeploymentLabelArray) ToDeploymentLabelArrayOutput() DeploymentLabelArrayOutput {
	return i.ToDeploymentLabelArrayOutputWithContext(context.Background())
}

func (i DeploymentLabelArray) ToDeploymentLabelArrayOutputWithContext(ctx context.Context) DeploymentLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentLabelArrayOutput)
}

type DeploymentLabelOutput struct{ *pulumi.OutputState }

func (DeploymentLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentLabel)(nil)).Elem()
}

func (o DeploymentLabelOutput) ToDeploymentLabelOutput() DeploymentLabelOutput {
	return o
}

func (o DeploymentLabelOutput) ToDeploymentLabelOutputWithContext(ctx context.Context) DeploymentLabelOutput {
	return o
}

// Key for label.
func (o DeploymentLabelOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentLabel) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// Value of label.
func (o DeploymentLabelOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentLabel) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DeploymentLabelArrayOutput struct{ *pulumi.OutputState }

func (DeploymentLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentLabel)(nil)).Elem()
}

func (o DeploymentLabelArrayOutput) ToDeploymentLabelArrayOutput() DeploymentLabelArrayOutput {
	return o
}

func (o DeploymentLabelArrayOutput) ToDeploymentLabelArrayOutputWithContext(ctx context.Context) DeploymentLabelArrayOutput {
	return o
}

func (o DeploymentLabelArrayOutput) Index(i pulumi.IntInput) DeploymentLabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeploymentLabel {
		return vs[0].([]DeploymentLabel)[vs[1].(int)]
	}).(DeploymentLabelOutput)
}

type DeploymentTarget struct {
	// The root configuration file to use for this deployment.  Structure is documented below.
	Config DeploymentTargetConfig `pulumi:"config"`
	// Specifies import files for this configuration. This can be
	// used to import templates or other files. For example, you might
	// import a text file in order to use the file in a template.  Structure is documented below.
	Imports []DeploymentTargetImport `pulumi:"imports"`
}

// DeploymentTargetInput is an input type that accepts DeploymentTargetArgs and DeploymentTargetOutput values.
// You can construct a concrete instance of `DeploymentTargetInput` via:
//
//          DeploymentTargetArgs{...}
type DeploymentTargetInput interface {
	pulumi.Input

	ToDeploymentTargetOutput() DeploymentTargetOutput
	ToDeploymentTargetOutputWithContext(context.Context) DeploymentTargetOutput
}

type DeploymentTargetArgs struct {
	// The root configuration file to use for this deployment.  Structure is documented below.
	Config DeploymentTargetConfigInput `pulumi:"config"`
	// Specifies import files for this configuration. This can be
	// used to import templates or other files. For example, you might
	// import a text file in order to use the file in a template.  Structure is documented below.
	Imports DeploymentTargetImportArrayInput `pulumi:"imports"`
}

func (DeploymentTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTarget)(nil)).Elem()
}

func (i DeploymentTargetArgs) ToDeploymentTargetOutput() DeploymentTargetOutput {
	return i.ToDeploymentTargetOutputWithContext(context.Background())
}

func (i DeploymentTargetArgs) ToDeploymentTargetOutputWithContext(ctx context.Context) DeploymentTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTargetOutput)
}

func (i DeploymentTargetArgs) ToDeploymentTargetPtrOutput() DeploymentTargetPtrOutput {
	return i.ToDeploymentTargetPtrOutputWithContext(context.Background())
}

func (i DeploymentTargetArgs) ToDeploymentTargetPtrOutputWithContext(ctx context.Context) DeploymentTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTargetOutput).ToDeploymentTargetPtrOutputWithContext(ctx)
}

// DeploymentTargetPtrInput is an input type that accepts DeploymentTargetArgs, DeploymentTargetPtr and DeploymentTargetPtrOutput values.
// You can construct a concrete instance of `DeploymentTargetPtrInput` via:
//
//          DeploymentTargetArgs{...}
//
//  or:
//
//          nil
type DeploymentTargetPtrInput interface {
	pulumi.Input

	ToDeploymentTargetPtrOutput() DeploymentTargetPtrOutput
	ToDeploymentTargetPtrOutputWithContext(context.Context) DeploymentTargetPtrOutput
}

type deploymentTargetPtrType DeploymentTargetArgs

func DeploymentTargetPtr(v *DeploymentTargetArgs) DeploymentTargetPtrInput {
	return (*deploymentTargetPtrType)(v)
}

func (*deploymentTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentTarget)(nil)).Elem()
}

func (i *deploymentTargetPtrType) ToDeploymentTargetPtrOutput() DeploymentTargetPtrOutput {
	return i.ToDeploymentTargetPtrOutputWithContext(context.Background())
}

func (i *deploymentTargetPtrType) ToDeploymentTargetPtrOutputWithContext(ctx context.Context) DeploymentTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTargetPtrOutput)
}

type DeploymentTargetOutput struct{ *pulumi.OutputState }

func (DeploymentTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTarget)(nil)).Elem()
}

func (o DeploymentTargetOutput) ToDeploymentTargetOutput() DeploymentTargetOutput {
	return o
}

func (o DeploymentTargetOutput) ToDeploymentTargetOutputWithContext(ctx context.Context) DeploymentTargetOutput {
	return o
}

func (o DeploymentTargetOutput) ToDeploymentTargetPtrOutput() DeploymentTargetPtrOutput {
	return o.ToDeploymentTargetPtrOutputWithContext(context.Background())
}

func (o DeploymentTargetOutput) ToDeploymentTargetPtrOutputWithContext(ctx context.Context) DeploymentTargetPtrOutput {
	return o.ApplyT(func(v DeploymentTarget) *DeploymentTarget {
		return &v
	}).(DeploymentTargetPtrOutput)
}

// The root configuration file to use for this deployment.  Structure is documented below.
func (o DeploymentTargetOutput) Config() DeploymentTargetConfigOutput {
	return o.ApplyT(func(v DeploymentTarget) DeploymentTargetConfig { return v.Config }).(DeploymentTargetConfigOutput)
}

// Specifies import files for this configuration. This can be
// used to import templates or other files. For example, you might
// import a text file in order to use the file in a template.  Structure is documented below.
func (o DeploymentTargetOutput) Imports() DeploymentTargetImportArrayOutput {
	return o.ApplyT(func(v DeploymentTarget) []DeploymentTargetImport { return v.Imports }).(DeploymentTargetImportArrayOutput)
}

type DeploymentTargetPtrOutput struct{ *pulumi.OutputState }

func (DeploymentTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentTarget)(nil)).Elem()
}

func (o DeploymentTargetPtrOutput) ToDeploymentTargetPtrOutput() DeploymentTargetPtrOutput {
	return o
}

func (o DeploymentTargetPtrOutput) ToDeploymentTargetPtrOutputWithContext(ctx context.Context) DeploymentTargetPtrOutput {
	return o
}

func (o DeploymentTargetPtrOutput) Elem() DeploymentTargetOutput {
	return o.ApplyT(func(v *DeploymentTarget) DeploymentTarget { return *v }).(DeploymentTargetOutput)
}

// The root configuration file to use for this deployment.  Structure is documented below.
func (o DeploymentTargetPtrOutput) Config() DeploymentTargetConfigPtrOutput {
	return o.ApplyT(func(v *DeploymentTarget) *DeploymentTargetConfig {
		if v == nil {
			return nil
		}
		return &v.Config
	}).(DeploymentTargetConfigPtrOutput)
}

// Specifies import files for this configuration. This can be
// used to import templates or other files. For example, you might
// import a text file in order to use the file in a template.  Structure is documented below.
func (o DeploymentTargetPtrOutput) Imports() DeploymentTargetImportArrayOutput {
	return o.ApplyT(func(v *DeploymentTarget) []DeploymentTargetImport {
		if v == nil {
			return nil
		}
		return v.Imports
	}).(DeploymentTargetImportArrayOutput)
}

type DeploymentTargetConfig struct {
	// The full contents of the template that you want to import.
	Content string `pulumi:"content"`
}

// DeploymentTargetConfigInput is an input type that accepts DeploymentTargetConfigArgs and DeploymentTargetConfigOutput values.
// You can construct a concrete instance of `DeploymentTargetConfigInput` via:
//
//          DeploymentTargetConfigArgs{...}
type DeploymentTargetConfigInput interface {
	pulumi.Input

	ToDeploymentTargetConfigOutput() DeploymentTargetConfigOutput
	ToDeploymentTargetConfigOutputWithContext(context.Context) DeploymentTargetConfigOutput
}

type DeploymentTargetConfigArgs struct {
	// The full contents of the template that you want to import.
	Content pulumi.StringInput `pulumi:"content"`
}

func (DeploymentTargetConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTargetConfig)(nil)).Elem()
}

func (i DeploymentTargetConfigArgs) ToDeploymentTargetConfigOutput() DeploymentTargetConfigOutput {
	return i.ToDeploymentTargetConfigOutputWithContext(context.Background())
}

func (i DeploymentTargetConfigArgs) ToDeploymentTargetConfigOutputWithContext(ctx context.Context) DeploymentTargetConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTargetConfigOutput)
}

func (i DeploymentTargetConfigArgs) ToDeploymentTargetConfigPtrOutput() DeploymentTargetConfigPtrOutput {
	return i.ToDeploymentTargetConfigPtrOutputWithContext(context.Background())
}

func (i DeploymentTargetConfigArgs) ToDeploymentTargetConfigPtrOutputWithContext(ctx context.Context) DeploymentTargetConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTargetConfigOutput).ToDeploymentTargetConfigPtrOutputWithContext(ctx)
}

// DeploymentTargetConfigPtrInput is an input type that accepts DeploymentTargetConfigArgs, DeploymentTargetConfigPtr and DeploymentTargetConfigPtrOutput values.
// You can construct a concrete instance of `DeploymentTargetConfigPtrInput` via:
//
//          DeploymentTargetConfigArgs{...}
//
//  or:
//
//          nil
type DeploymentTargetConfigPtrInput interface {
	pulumi.Input

	ToDeploymentTargetConfigPtrOutput() DeploymentTargetConfigPtrOutput
	ToDeploymentTargetConfigPtrOutputWithContext(context.Context) DeploymentTargetConfigPtrOutput
}

type deploymentTargetConfigPtrType DeploymentTargetConfigArgs

func DeploymentTargetConfigPtr(v *DeploymentTargetConfigArgs) DeploymentTargetConfigPtrInput {
	return (*deploymentTargetConfigPtrType)(v)
}

func (*deploymentTargetConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentTargetConfig)(nil)).Elem()
}

func (i *deploymentTargetConfigPtrType) ToDeploymentTargetConfigPtrOutput() DeploymentTargetConfigPtrOutput {
	return i.ToDeploymentTargetConfigPtrOutputWithContext(context.Background())
}

func (i *deploymentTargetConfigPtrType) ToDeploymentTargetConfigPtrOutputWithContext(ctx context.Context) DeploymentTargetConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTargetConfigPtrOutput)
}

type DeploymentTargetConfigOutput struct{ *pulumi.OutputState }

func (DeploymentTargetConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTargetConfig)(nil)).Elem()
}

func (o DeploymentTargetConfigOutput) ToDeploymentTargetConfigOutput() DeploymentTargetConfigOutput {
	return o
}

func (o DeploymentTargetConfigOutput) ToDeploymentTargetConfigOutputWithContext(ctx context.Context) DeploymentTargetConfigOutput {
	return o
}

func (o DeploymentTargetConfigOutput) ToDeploymentTargetConfigPtrOutput() DeploymentTargetConfigPtrOutput {
	return o.ToDeploymentTargetConfigPtrOutputWithContext(context.Background())
}

func (o DeploymentTargetConfigOutput) ToDeploymentTargetConfigPtrOutputWithContext(ctx context.Context) DeploymentTargetConfigPtrOutput {
	return o.ApplyT(func(v DeploymentTargetConfig) *DeploymentTargetConfig {
		return &v
	}).(DeploymentTargetConfigPtrOutput)
}

// The full contents of the template that you want to import.
func (o DeploymentTargetConfigOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentTargetConfig) string { return v.Content }).(pulumi.StringOutput)
}

type DeploymentTargetConfigPtrOutput struct{ *pulumi.OutputState }

func (DeploymentTargetConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentTargetConfig)(nil)).Elem()
}

func (o DeploymentTargetConfigPtrOutput) ToDeploymentTargetConfigPtrOutput() DeploymentTargetConfigPtrOutput {
	return o
}

func (o DeploymentTargetConfigPtrOutput) ToDeploymentTargetConfigPtrOutputWithContext(ctx context.Context) DeploymentTargetConfigPtrOutput {
	return o
}

func (o DeploymentTargetConfigPtrOutput) Elem() DeploymentTargetConfigOutput {
	return o.ApplyT(func(v *DeploymentTargetConfig) DeploymentTargetConfig { return *v }).(DeploymentTargetConfigOutput)
}

// The full contents of the template that you want to import.
func (o DeploymentTargetConfigPtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentTargetConfig) *string {
		if v == nil {
			return nil
		}
		return &v.Content
	}).(pulumi.StringPtrOutput)
}

type DeploymentTargetImport struct {
	// The full contents of the template that you want to import.
	Content *string `pulumi:"content"`
	// The name of the template to import, as declared in the YAML
	// configuration.
	Name *string `pulumi:"name"`
}

// DeploymentTargetImportInput is an input type that accepts DeploymentTargetImportArgs and DeploymentTargetImportOutput values.
// You can construct a concrete instance of `DeploymentTargetImportInput` via:
//
//          DeploymentTargetImportArgs{...}
type DeploymentTargetImportInput interface {
	pulumi.Input

	ToDeploymentTargetImportOutput() DeploymentTargetImportOutput
	ToDeploymentTargetImportOutputWithContext(context.Context) DeploymentTargetImportOutput
}

type DeploymentTargetImportArgs struct {
	// The full contents of the template that you want to import.
	Content pulumi.StringPtrInput `pulumi:"content"`
	// The name of the template to import, as declared in the YAML
	// configuration.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DeploymentTargetImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTargetImport)(nil)).Elem()
}

func (i DeploymentTargetImportArgs) ToDeploymentTargetImportOutput() DeploymentTargetImportOutput {
	return i.ToDeploymentTargetImportOutputWithContext(context.Background())
}

func (i DeploymentTargetImportArgs) ToDeploymentTargetImportOutputWithContext(ctx context.Context) DeploymentTargetImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTargetImportOutput)
}

// DeploymentTargetImportArrayInput is an input type that accepts DeploymentTargetImportArray and DeploymentTargetImportArrayOutput values.
// You can construct a concrete instance of `DeploymentTargetImportArrayInput` via:
//
//          DeploymentTargetImportArray{ DeploymentTargetImportArgs{...} }
type DeploymentTargetImportArrayInput interface {
	pulumi.Input

	ToDeploymentTargetImportArrayOutput() DeploymentTargetImportArrayOutput
	ToDeploymentTargetImportArrayOutputWithContext(context.Context) DeploymentTargetImportArrayOutput
}

type DeploymentTargetImportArray []DeploymentTargetImportInput

func (DeploymentTargetImportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentTargetImport)(nil)).Elem()
}

func (i DeploymentTargetImportArray) ToDeploymentTargetImportArrayOutput() DeploymentTargetImportArrayOutput {
	return i.ToDeploymentTargetImportArrayOutputWithContext(context.Background())
}

func (i DeploymentTargetImportArray) ToDeploymentTargetImportArrayOutputWithContext(ctx context.Context) DeploymentTargetImportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTargetImportArrayOutput)
}

type DeploymentTargetImportOutput struct{ *pulumi.OutputState }

func (DeploymentTargetImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTargetImport)(nil)).Elem()
}

func (o DeploymentTargetImportOutput) ToDeploymentTargetImportOutput() DeploymentTargetImportOutput {
	return o
}

func (o DeploymentTargetImportOutput) ToDeploymentTargetImportOutputWithContext(ctx context.Context) DeploymentTargetImportOutput {
	return o
}

// The full contents of the template that you want to import.
func (o DeploymentTargetImportOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentTargetImport) *string { return v.Content }).(pulumi.StringPtrOutput)
}

// The name of the template to import, as declared in the YAML
// configuration.
func (o DeploymentTargetImportOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentTargetImport) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DeploymentTargetImportArrayOutput struct{ *pulumi.OutputState }

func (DeploymentTargetImportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentTargetImport)(nil)).Elem()
}

func (o DeploymentTargetImportArrayOutput) ToDeploymentTargetImportArrayOutput() DeploymentTargetImportArrayOutput {
	return o
}

func (o DeploymentTargetImportArrayOutput) ToDeploymentTargetImportArrayOutputWithContext(ctx context.Context) DeploymentTargetImportArrayOutput {
	return o
}

func (o DeploymentTargetImportArrayOutput) Index(i pulumi.IntInput) DeploymentTargetImportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeploymentTargetImport {
		return vs[0].([]DeploymentTargetImport)[vs[1].(int)]
	}).(DeploymentTargetImportOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentLabelOutput{})
	pulumi.RegisterOutputType(DeploymentLabelArrayOutput{})
	pulumi.RegisterOutputType(DeploymentTargetOutput{})
	pulumi.RegisterOutputType(DeploymentTargetPtrOutput{})
	pulumi.RegisterOutputType(DeploymentTargetConfigOutput{})
	pulumi.RegisterOutputType(DeploymentTargetConfigPtrOutput{})
	pulumi.RegisterOutputType(DeploymentTargetImportOutput{})
	pulumi.RegisterOutputType(DeploymentTargetImportArrayOutput{})
}
