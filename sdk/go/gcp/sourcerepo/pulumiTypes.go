// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sourcerepo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type RepositoryIamBindingCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// RepositoryIamBindingConditionInput is an input type that accepts RepositoryIamBindingConditionArgs and RepositoryIamBindingConditionOutput values.
// You can construct a concrete instance of `RepositoryIamBindingConditionInput` via:
//
//          RepositoryIamBindingConditionArgs{...}
type RepositoryIamBindingConditionInput interface {
	pulumi.Input

	ToRepositoryIamBindingConditionOutput() RepositoryIamBindingConditionOutput
	ToRepositoryIamBindingConditionOutputWithContext(context.Context) RepositoryIamBindingConditionOutput
}

type RepositoryIamBindingConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (RepositoryIamBindingConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamBindingCondition)(nil)).Elem()
}

func (i RepositoryIamBindingConditionArgs) ToRepositoryIamBindingConditionOutput() RepositoryIamBindingConditionOutput {
	return i.ToRepositoryIamBindingConditionOutputWithContext(context.Background())
}

func (i RepositoryIamBindingConditionArgs) ToRepositoryIamBindingConditionOutputWithContext(ctx context.Context) RepositoryIamBindingConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamBindingConditionOutput)
}

func (i RepositoryIamBindingConditionArgs) ToRepositoryIamBindingConditionPtrOutput() RepositoryIamBindingConditionPtrOutput {
	return i.ToRepositoryIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i RepositoryIamBindingConditionArgs) ToRepositoryIamBindingConditionPtrOutputWithContext(ctx context.Context) RepositoryIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamBindingConditionOutput).ToRepositoryIamBindingConditionPtrOutputWithContext(ctx)
}

// RepositoryIamBindingConditionPtrInput is an input type that accepts RepositoryIamBindingConditionArgs, RepositoryIamBindingConditionPtr and RepositoryIamBindingConditionPtrOutput values.
// You can construct a concrete instance of `RepositoryIamBindingConditionPtrInput` via:
//
//          RepositoryIamBindingConditionArgs{...}
//
//  or:
//
//          nil
type RepositoryIamBindingConditionPtrInput interface {
	pulumi.Input

	ToRepositoryIamBindingConditionPtrOutput() RepositoryIamBindingConditionPtrOutput
	ToRepositoryIamBindingConditionPtrOutputWithContext(context.Context) RepositoryIamBindingConditionPtrOutput
}

type repositoryIamBindingConditionPtrType RepositoryIamBindingConditionArgs

func RepositoryIamBindingConditionPtr(v *RepositoryIamBindingConditionArgs) RepositoryIamBindingConditionPtrInput {
	return (*repositoryIamBindingConditionPtrType)(v)
}

func (*repositoryIamBindingConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamBindingCondition)(nil)).Elem()
}

func (i *repositoryIamBindingConditionPtrType) ToRepositoryIamBindingConditionPtrOutput() RepositoryIamBindingConditionPtrOutput {
	return i.ToRepositoryIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i *repositoryIamBindingConditionPtrType) ToRepositoryIamBindingConditionPtrOutputWithContext(ctx context.Context) RepositoryIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamBindingConditionPtrOutput)
}

type RepositoryIamBindingConditionOutput struct{ *pulumi.OutputState }

func (RepositoryIamBindingConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamBindingCondition)(nil)).Elem()
}

func (o RepositoryIamBindingConditionOutput) ToRepositoryIamBindingConditionOutput() RepositoryIamBindingConditionOutput {
	return o
}

func (o RepositoryIamBindingConditionOutput) ToRepositoryIamBindingConditionOutputWithContext(ctx context.Context) RepositoryIamBindingConditionOutput {
	return o
}

func (o RepositoryIamBindingConditionOutput) ToRepositoryIamBindingConditionPtrOutput() RepositoryIamBindingConditionPtrOutput {
	return o.ToRepositoryIamBindingConditionPtrOutputWithContext(context.Background())
}

func (o RepositoryIamBindingConditionOutput) ToRepositoryIamBindingConditionPtrOutputWithContext(ctx context.Context) RepositoryIamBindingConditionPtrOutput {
	return o.ApplyT(func(v RepositoryIamBindingCondition) *RepositoryIamBindingCondition {
		return &v
	}).(RepositoryIamBindingConditionPtrOutput)
}
func (o RepositoryIamBindingConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RepositoryIamBindingConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v RepositoryIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o RepositoryIamBindingConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v RepositoryIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type RepositoryIamBindingConditionPtrOutput struct{ *pulumi.OutputState }

func (RepositoryIamBindingConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamBindingCondition)(nil)).Elem()
}

func (o RepositoryIamBindingConditionPtrOutput) ToRepositoryIamBindingConditionPtrOutput() RepositoryIamBindingConditionPtrOutput {
	return o
}

func (o RepositoryIamBindingConditionPtrOutput) ToRepositoryIamBindingConditionPtrOutputWithContext(ctx context.Context) RepositoryIamBindingConditionPtrOutput {
	return o
}

func (o RepositoryIamBindingConditionPtrOutput) Elem() RepositoryIamBindingConditionOutput {
	return o.ApplyT(func(v *RepositoryIamBindingCondition) RepositoryIamBindingCondition { return *v }).(RepositoryIamBindingConditionOutput)
}

func (o RepositoryIamBindingConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryIamBindingConditionPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Expression
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryIamBindingConditionPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryIamBindingCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

type RepositoryIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// RepositoryIamMemberConditionInput is an input type that accepts RepositoryIamMemberConditionArgs and RepositoryIamMemberConditionOutput values.
// You can construct a concrete instance of `RepositoryIamMemberConditionInput` via:
//
//          RepositoryIamMemberConditionArgs{...}
type RepositoryIamMemberConditionInput interface {
	pulumi.Input

	ToRepositoryIamMemberConditionOutput() RepositoryIamMemberConditionOutput
	ToRepositoryIamMemberConditionOutputWithContext(context.Context) RepositoryIamMemberConditionOutput
}

type RepositoryIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (RepositoryIamMemberConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamMemberCondition)(nil)).Elem()
}

func (i RepositoryIamMemberConditionArgs) ToRepositoryIamMemberConditionOutput() RepositoryIamMemberConditionOutput {
	return i.ToRepositoryIamMemberConditionOutputWithContext(context.Background())
}

func (i RepositoryIamMemberConditionArgs) ToRepositoryIamMemberConditionOutputWithContext(ctx context.Context) RepositoryIamMemberConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberConditionOutput)
}

func (i RepositoryIamMemberConditionArgs) ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput {
	return i.ToRepositoryIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i RepositoryIamMemberConditionArgs) ToRepositoryIamMemberConditionPtrOutputWithContext(ctx context.Context) RepositoryIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberConditionOutput).ToRepositoryIamMemberConditionPtrOutputWithContext(ctx)
}

// RepositoryIamMemberConditionPtrInput is an input type that accepts RepositoryIamMemberConditionArgs, RepositoryIamMemberConditionPtr and RepositoryIamMemberConditionPtrOutput values.
// You can construct a concrete instance of `RepositoryIamMemberConditionPtrInput` via:
//
//          RepositoryIamMemberConditionArgs{...}
//
//  or:
//
//          nil
type RepositoryIamMemberConditionPtrInput interface {
	pulumi.Input

	ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput
	ToRepositoryIamMemberConditionPtrOutputWithContext(context.Context) RepositoryIamMemberConditionPtrOutput
}

type repositoryIamMemberConditionPtrType RepositoryIamMemberConditionArgs

func RepositoryIamMemberConditionPtr(v *RepositoryIamMemberConditionArgs) RepositoryIamMemberConditionPtrInput {
	return (*repositoryIamMemberConditionPtrType)(v)
}

func (*repositoryIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamMemberCondition)(nil)).Elem()
}

func (i *repositoryIamMemberConditionPtrType) ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput {
	return i.ToRepositoryIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *repositoryIamMemberConditionPtrType) ToRepositoryIamMemberConditionPtrOutputWithContext(ctx context.Context) RepositoryIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamMemberConditionPtrOutput)
}

type RepositoryIamMemberConditionOutput struct{ *pulumi.OutputState }

func (RepositoryIamMemberConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryIamMemberCondition)(nil)).Elem()
}

func (o RepositoryIamMemberConditionOutput) ToRepositoryIamMemberConditionOutput() RepositoryIamMemberConditionOutput {
	return o
}

func (o RepositoryIamMemberConditionOutput) ToRepositoryIamMemberConditionOutputWithContext(ctx context.Context) RepositoryIamMemberConditionOutput {
	return o
}

func (o RepositoryIamMemberConditionOutput) ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput {
	return o.ToRepositoryIamMemberConditionPtrOutputWithContext(context.Background())
}

func (o RepositoryIamMemberConditionOutput) ToRepositoryIamMemberConditionPtrOutputWithContext(ctx context.Context) RepositoryIamMemberConditionPtrOutput {
	return o.ApplyT(func(v RepositoryIamMemberCondition) *RepositoryIamMemberCondition {
		return &v
	}).(RepositoryIamMemberConditionPtrOutput)
}
func (o RepositoryIamMemberConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RepositoryIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v RepositoryIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o RepositoryIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v RepositoryIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type RepositoryIamMemberConditionPtrOutput struct{ *pulumi.OutputState }

func (RepositoryIamMemberConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamMemberCondition)(nil)).Elem()
}

func (o RepositoryIamMemberConditionPtrOutput) ToRepositoryIamMemberConditionPtrOutput() RepositoryIamMemberConditionPtrOutput {
	return o
}

func (o RepositoryIamMemberConditionPtrOutput) ToRepositoryIamMemberConditionPtrOutputWithContext(ctx context.Context) RepositoryIamMemberConditionPtrOutput {
	return o
}

func (o RepositoryIamMemberConditionPtrOutput) Elem() RepositoryIamMemberConditionOutput {
	return o.ApplyT(func(v *RepositoryIamMemberCondition) RepositoryIamMemberCondition { return *v }).(RepositoryIamMemberConditionOutput)
}

func (o RepositoryIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryIamMemberConditionPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Expression
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryIamMemberConditionPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryIamMemberCondition) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

type RepositoryPubsubConfig struct {
	// The format of the Cloud Pub/Sub messages.
	// - PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
	// - JSON: The message payload is a JSON string of SourceRepoEvent.
	MessageFormat string `pulumi:"messageFormat"`
	// Email address of the service account used for publishing Cloud Pub/Sub messages.
	// This service account needs to be in the same project as the PubsubConfig. When added,
	// the caller needs to have iam.serviceAccounts.actAs permission on this service account.
	// If unspecified, it defaults to the compute engine default service account.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The identifier for this object. Format specified above.
	Topic string `pulumi:"topic"`
}

// RepositoryPubsubConfigInput is an input type that accepts RepositoryPubsubConfigArgs and RepositoryPubsubConfigOutput values.
// You can construct a concrete instance of `RepositoryPubsubConfigInput` via:
//
//          RepositoryPubsubConfigArgs{...}
type RepositoryPubsubConfigInput interface {
	pulumi.Input

	ToRepositoryPubsubConfigOutput() RepositoryPubsubConfigOutput
	ToRepositoryPubsubConfigOutputWithContext(context.Context) RepositoryPubsubConfigOutput
}

type RepositoryPubsubConfigArgs struct {
	// The format of the Cloud Pub/Sub messages.
	// - PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
	// - JSON: The message payload is a JSON string of SourceRepoEvent.
	MessageFormat pulumi.StringInput `pulumi:"messageFormat"`
	// Email address of the service account used for publishing Cloud Pub/Sub messages.
	// This service account needs to be in the same project as the PubsubConfig. When added,
	// the caller needs to have iam.serviceAccounts.actAs permission on this service account.
	// If unspecified, it defaults to the compute engine default service account.
	ServiceAccountEmail pulumi.StringPtrInput `pulumi:"serviceAccountEmail"`
	// The identifier for this object. Format specified above.
	Topic pulumi.StringInput `pulumi:"topic"`
}

func (RepositoryPubsubConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPubsubConfig)(nil)).Elem()
}

func (i RepositoryPubsubConfigArgs) ToRepositoryPubsubConfigOutput() RepositoryPubsubConfigOutput {
	return i.ToRepositoryPubsubConfigOutputWithContext(context.Background())
}

func (i RepositoryPubsubConfigArgs) ToRepositoryPubsubConfigOutputWithContext(ctx context.Context) RepositoryPubsubConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPubsubConfigOutput)
}

// RepositoryPubsubConfigArrayInput is an input type that accepts RepositoryPubsubConfigArray and RepositoryPubsubConfigArrayOutput values.
// You can construct a concrete instance of `RepositoryPubsubConfigArrayInput` via:
//
//          RepositoryPubsubConfigArray{ RepositoryPubsubConfigArgs{...} }
type RepositoryPubsubConfigArrayInput interface {
	pulumi.Input

	ToRepositoryPubsubConfigArrayOutput() RepositoryPubsubConfigArrayOutput
	ToRepositoryPubsubConfigArrayOutputWithContext(context.Context) RepositoryPubsubConfigArrayOutput
}

type RepositoryPubsubConfigArray []RepositoryPubsubConfigInput

func (RepositoryPubsubConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryPubsubConfig)(nil)).Elem()
}

func (i RepositoryPubsubConfigArray) ToRepositoryPubsubConfigArrayOutput() RepositoryPubsubConfigArrayOutput {
	return i.ToRepositoryPubsubConfigArrayOutputWithContext(context.Background())
}

func (i RepositoryPubsubConfigArray) ToRepositoryPubsubConfigArrayOutputWithContext(ctx context.Context) RepositoryPubsubConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPubsubConfigArrayOutput)
}

type RepositoryPubsubConfigOutput struct{ *pulumi.OutputState }

func (RepositoryPubsubConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPubsubConfig)(nil)).Elem()
}

func (o RepositoryPubsubConfigOutput) ToRepositoryPubsubConfigOutput() RepositoryPubsubConfigOutput {
	return o
}

func (o RepositoryPubsubConfigOutput) ToRepositoryPubsubConfigOutputWithContext(ctx context.Context) RepositoryPubsubConfigOutput {
	return o
}

// The format of the Cloud Pub/Sub messages.
// - PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
// - JSON: The message payload is a JSON string of SourceRepoEvent.
func (o RepositoryPubsubConfigOutput) MessageFormat() pulumi.StringOutput {
	return o.ApplyT(func(v RepositoryPubsubConfig) string { return v.MessageFormat }).(pulumi.StringOutput)
}

// Email address of the service account used for publishing Cloud Pub/Sub messages.
// This service account needs to be in the same project as the PubsubConfig. When added,
// the caller needs to have iam.serviceAccounts.actAs permission on this service account.
// If unspecified, it defaults to the compute engine default service account.
func (o RepositoryPubsubConfigOutput) ServiceAccountEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryPubsubConfig) *string { return v.ServiceAccountEmail }).(pulumi.StringPtrOutput)
}

// The identifier for this object. Format specified above.
func (o RepositoryPubsubConfigOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v RepositoryPubsubConfig) string { return v.Topic }).(pulumi.StringOutput)
}

type RepositoryPubsubConfigArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPubsubConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryPubsubConfig)(nil)).Elem()
}

func (o RepositoryPubsubConfigArrayOutput) ToRepositoryPubsubConfigArrayOutput() RepositoryPubsubConfigArrayOutput {
	return o
}

func (o RepositoryPubsubConfigArrayOutput) ToRepositoryPubsubConfigArrayOutputWithContext(ctx context.Context) RepositoryPubsubConfigArrayOutput {
	return o
}

func (o RepositoryPubsubConfigArrayOutput) Index(i pulumi.IntInput) RepositoryPubsubConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepositoryPubsubConfig {
		return vs[0].([]RepositoryPubsubConfig)[vs[1].(int)]
	}).(RepositoryPubsubConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryIamBindingConditionOutput{})
	pulumi.RegisterOutputType(RepositoryIamBindingConditionPtrOutput{})
	pulumi.RegisterOutputType(RepositoryIamMemberConditionOutput{})
	pulumi.RegisterOutputType(RepositoryIamMemberConditionPtrOutput{})
	pulumi.RegisterOutputType(RepositoryPubsubConfigOutput{})
	pulumi.RegisterOutputType(RepositoryPubsubConfigArrayOutput{})
}
