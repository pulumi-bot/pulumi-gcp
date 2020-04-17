// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecretIamBindingCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// SecretIamBindingConditionInput is an input type that accepts SecretIamBindingConditionArgs and SecretIamBindingConditionOutput values.
// You can construct a concrete instance of `SecretIamBindingConditionInput` via:
//
// 		 SecretIamBindingConditionArgs{...}
//
type SecretIamBindingConditionInput interface {
	pulumi.Input

	ToSecretIamBindingConditionOutput() SecretIamBindingConditionOutput
	ToSecretIamBindingConditionOutputWithContext(context.Context) SecretIamBindingConditionOutput
}

type SecretIamBindingConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (SecretIamBindingConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretIamBindingCondition)(nil)).Elem()
}

func (i SecretIamBindingConditionArgs) ToSecretIamBindingConditionOutput() SecretIamBindingConditionOutput {
	return i.ToSecretIamBindingConditionOutputWithContext(context.Background())
}

func (i SecretIamBindingConditionArgs) ToSecretIamBindingConditionOutputWithContext(ctx context.Context) SecretIamBindingConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretIamBindingConditionOutput)
}

func (i SecretIamBindingConditionArgs) ToSecretIamBindingConditionPtrOutput() SecretIamBindingConditionPtrOutput {
	return i.ToSecretIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i SecretIamBindingConditionArgs) ToSecretIamBindingConditionPtrOutputWithContext(ctx context.Context) SecretIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretIamBindingConditionOutput).ToSecretIamBindingConditionPtrOutputWithContext(ctx)
}

// SecretIamBindingConditionPtrInput is an input type that accepts SecretIamBindingConditionArgs, SecretIamBindingConditionPtr and SecretIamBindingConditionPtrOutput values.
// You can construct a concrete instance of `SecretIamBindingConditionPtrInput` via:
//
// 		 SecretIamBindingConditionArgs{...}
//
//  or:
//
// 		 nil
//
type SecretIamBindingConditionPtrInput interface {
	pulumi.Input

	ToSecretIamBindingConditionPtrOutput() SecretIamBindingConditionPtrOutput
	ToSecretIamBindingConditionPtrOutputWithContext(context.Context) SecretIamBindingConditionPtrOutput
}

type secretIamBindingConditionPtrType SecretIamBindingConditionArgs

func SecretIamBindingConditionPtr(v *SecretIamBindingConditionArgs) SecretIamBindingConditionPtrInput {
	return (*secretIamBindingConditionPtrType)(v)
}

func (*secretIamBindingConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretIamBindingCondition)(nil)).Elem()
}

func (i *secretIamBindingConditionPtrType) ToSecretIamBindingConditionPtrOutput() SecretIamBindingConditionPtrOutput {
	return i.ToSecretIamBindingConditionPtrOutputWithContext(context.Background())
}

func (i *secretIamBindingConditionPtrType) ToSecretIamBindingConditionPtrOutputWithContext(ctx context.Context) SecretIamBindingConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretIamBindingConditionPtrOutput)
}

type SecretIamBindingConditionOutput struct{ *pulumi.OutputState }

func (SecretIamBindingConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretIamBindingCondition)(nil)).Elem()
}

func (o SecretIamBindingConditionOutput) ToSecretIamBindingConditionOutput() SecretIamBindingConditionOutput {
	return o
}

func (o SecretIamBindingConditionOutput) ToSecretIamBindingConditionOutputWithContext(ctx context.Context) SecretIamBindingConditionOutput {
	return o
}

func (o SecretIamBindingConditionOutput) ToSecretIamBindingConditionPtrOutput() SecretIamBindingConditionPtrOutput {
	return o.ToSecretIamBindingConditionPtrOutputWithContext(context.Background())
}

func (o SecretIamBindingConditionOutput) ToSecretIamBindingConditionPtrOutputWithContext(ctx context.Context) SecretIamBindingConditionPtrOutput {
	return o.ApplyT(func(v SecretIamBindingCondition) *SecretIamBindingCondition {
		return &v
	}).(SecretIamBindingConditionPtrOutput)
}
func (o SecretIamBindingConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecretIamBindingConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v SecretIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o SecretIamBindingConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v SecretIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type SecretIamBindingConditionPtrOutput struct{ *pulumi.OutputState }

func (SecretIamBindingConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretIamBindingCondition)(nil)).Elem()
}

func (o SecretIamBindingConditionPtrOutput) ToSecretIamBindingConditionPtrOutput() SecretIamBindingConditionPtrOutput {
	return o
}

func (o SecretIamBindingConditionPtrOutput) ToSecretIamBindingConditionPtrOutputWithContext(ctx context.Context) SecretIamBindingConditionPtrOutput {
	return o
}

func (o SecretIamBindingConditionPtrOutput) Elem() SecretIamBindingConditionOutput {
	return o.ApplyT(func(v *SecretIamBindingCondition) SecretIamBindingCondition { return *v }).(SecretIamBindingConditionOutput)
}

func (o SecretIamBindingConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretIamBindingCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecretIamBindingConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v SecretIamBindingCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o SecretIamBindingConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v SecretIamBindingCondition) string { return v.Title }).(pulumi.StringOutput)
}

type SecretIamMemberCondition struct {
	Description *string `pulumi:"description"`
	Expression  string  `pulumi:"expression"`
	Title       string  `pulumi:"title"`
}

// SecretIamMemberConditionInput is an input type that accepts SecretIamMemberConditionArgs and SecretIamMemberConditionOutput values.
// You can construct a concrete instance of `SecretIamMemberConditionInput` via:
//
// 		 SecretIamMemberConditionArgs{...}
//
type SecretIamMemberConditionInput interface {
	pulumi.Input

	ToSecretIamMemberConditionOutput() SecretIamMemberConditionOutput
	ToSecretIamMemberConditionOutputWithContext(context.Context) SecretIamMemberConditionOutput
}

type SecretIamMemberConditionArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Expression  pulumi.StringInput    `pulumi:"expression"`
	Title       pulumi.StringInput    `pulumi:"title"`
}

func (SecretIamMemberConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretIamMemberCondition)(nil)).Elem()
}

func (i SecretIamMemberConditionArgs) ToSecretIamMemberConditionOutput() SecretIamMemberConditionOutput {
	return i.ToSecretIamMemberConditionOutputWithContext(context.Background())
}

func (i SecretIamMemberConditionArgs) ToSecretIamMemberConditionOutputWithContext(ctx context.Context) SecretIamMemberConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretIamMemberConditionOutput)
}

func (i SecretIamMemberConditionArgs) ToSecretIamMemberConditionPtrOutput() SecretIamMemberConditionPtrOutput {
	return i.ToSecretIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i SecretIamMemberConditionArgs) ToSecretIamMemberConditionPtrOutputWithContext(ctx context.Context) SecretIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretIamMemberConditionOutput).ToSecretIamMemberConditionPtrOutputWithContext(ctx)
}

// SecretIamMemberConditionPtrInput is an input type that accepts SecretIamMemberConditionArgs, SecretIamMemberConditionPtr and SecretIamMemberConditionPtrOutput values.
// You can construct a concrete instance of `SecretIamMemberConditionPtrInput` via:
//
// 		 SecretIamMemberConditionArgs{...}
//
//  or:
//
// 		 nil
//
type SecretIamMemberConditionPtrInput interface {
	pulumi.Input

	ToSecretIamMemberConditionPtrOutput() SecretIamMemberConditionPtrOutput
	ToSecretIamMemberConditionPtrOutputWithContext(context.Context) SecretIamMemberConditionPtrOutput
}

type secretIamMemberConditionPtrType SecretIamMemberConditionArgs

func SecretIamMemberConditionPtr(v *SecretIamMemberConditionArgs) SecretIamMemberConditionPtrInput {
	return (*secretIamMemberConditionPtrType)(v)
}

func (*secretIamMemberConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretIamMemberCondition)(nil)).Elem()
}

func (i *secretIamMemberConditionPtrType) ToSecretIamMemberConditionPtrOutput() SecretIamMemberConditionPtrOutput {
	return i.ToSecretIamMemberConditionPtrOutputWithContext(context.Background())
}

func (i *secretIamMemberConditionPtrType) ToSecretIamMemberConditionPtrOutputWithContext(ctx context.Context) SecretIamMemberConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretIamMemberConditionPtrOutput)
}

type SecretIamMemberConditionOutput struct{ *pulumi.OutputState }

func (SecretIamMemberConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretIamMemberCondition)(nil)).Elem()
}

func (o SecretIamMemberConditionOutput) ToSecretIamMemberConditionOutput() SecretIamMemberConditionOutput {
	return o
}

func (o SecretIamMemberConditionOutput) ToSecretIamMemberConditionOutputWithContext(ctx context.Context) SecretIamMemberConditionOutput {
	return o
}

func (o SecretIamMemberConditionOutput) ToSecretIamMemberConditionPtrOutput() SecretIamMemberConditionPtrOutput {
	return o.ToSecretIamMemberConditionPtrOutputWithContext(context.Background())
}

func (o SecretIamMemberConditionOutput) ToSecretIamMemberConditionPtrOutputWithContext(ctx context.Context) SecretIamMemberConditionPtrOutput {
	return o.ApplyT(func(v SecretIamMemberCondition) *SecretIamMemberCondition {
		return &v
	}).(SecretIamMemberConditionPtrOutput)
}
func (o SecretIamMemberConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecretIamMemberConditionOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v SecretIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o SecretIamMemberConditionOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v SecretIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type SecretIamMemberConditionPtrOutput struct{ *pulumi.OutputState }

func (SecretIamMemberConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretIamMemberCondition)(nil)).Elem()
}

func (o SecretIamMemberConditionPtrOutput) ToSecretIamMemberConditionPtrOutput() SecretIamMemberConditionPtrOutput {
	return o
}

func (o SecretIamMemberConditionPtrOutput) ToSecretIamMemberConditionPtrOutputWithContext(ctx context.Context) SecretIamMemberConditionPtrOutput {
	return o
}

func (o SecretIamMemberConditionPtrOutput) Elem() SecretIamMemberConditionOutput {
	return o.ApplyT(func(v *SecretIamMemberCondition) SecretIamMemberCondition { return *v }).(SecretIamMemberConditionOutput)
}

func (o SecretIamMemberConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretIamMemberCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecretIamMemberConditionPtrOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v SecretIamMemberCondition) string { return v.Expression }).(pulumi.StringOutput)
}

func (o SecretIamMemberConditionPtrOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v SecretIamMemberCondition) string { return v.Title }).(pulumi.StringOutput)
}

type SecretReplication struct {
	// -
	// (Optional)
	// The Secret will automatically be replicated without any restrictions.
	Automatic *bool `pulumi:"automatic"`
	// -
	// (Optional)
	// The Secret will automatically be replicated without any restrictions.  Structure is documented below.
	UserManaged *SecretReplicationUserManaged `pulumi:"userManaged"`
}

// SecretReplicationInput is an input type that accepts SecretReplicationArgs and SecretReplicationOutput values.
// You can construct a concrete instance of `SecretReplicationInput` via:
//
// 		 SecretReplicationArgs{...}
//
type SecretReplicationInput interface {
	pulumi.Input

	ToSecretReplicationOutput() SecretReplicationOutput
	ToSecretReplicationOutputWithContext(context.Context) SecretReplicationOutput
}

type SecretReplicationArgs struct {
	// -
	// (Optional)
	// The Secret will automatically be replicated without any restrictions.
	Automatic pulumi.BoolPtrInput `pulumi:"automatic"`
	// -
	// (Optional)
	// The Secret will automatically be replicated without any restrictions.  Structure is documented below.
	UserManaged SecretReplicationUserManagedPtrInput `pulumi:"userManaged"`
}

func (SecretReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretReplication)(nil)).Elem()
}

func (i SecretReplicationArgs) ToSecretReplicationOutput() SecretReplicationOutput {
	return i.ToSecretReplicationOutputWithContext(context.Background())
}

func (i SecretReplicationArgs) ToSecretReplicationOutputWithContext(ctx context.Context) SecretReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicationOutput)
}

func (i SecretReplicationArgs) ToSecretReplicationPtrOutput() SecretReplicationPtrOutput {
	return i.ToSecretReplicationPtrOutputWithContext(context.Background())
}

func (i SecretReplicationArgs) ToSecretReplicationPtrOutputWithContext(ctx context.Context) SecretReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicationOutput).ToSecretReplicationPtrOutputWithContext(ctx)
}

// SecretReplicationPtrInput is an input type that accepts SecretReplicationArgs, SecretReplicationPtr and SecretReplicationPtrOutput values.
// You can construct a concrete instance of `SecretReplicationPtrInput` via:
//
// 		 SecretReplicationArgs{...}
//
//  or:
//
// 		 nil
//
type SecretReplicationPtrInput interface {
	pulumi.Input

	ToSecretReplicationPtrOutput() SecretReplicationPtrOutput
	ToSecretReplicationPtrOutputWithContext(context.Context) SecretReplicationPtrOutput
}

type secretReplicationPtrType SecretReplicationArgs

func SecretReplicationPtr(v *SecretReplicationArgs) SecretReplicationPtrInput {
	return (*secretReplicationPtrType)(v)
}

func (*secretReplicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretReplication)(nil)).Elem()
}

func (i *secretReplicationPtrType) ToSecretReplicationPtrOutput() SecretReplicationPtrOutput {
	return i.ToSecretReplicationPtrOutputWithContext(context.Background())
}

func (i *secretReplicationPtrType) ToSecretReplicationPtrOutputWithContext(ctx context.Context) SecretReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicationPtrOutput)
}

type SecretReplicationOutput struct{ *pulumi.OutputState }

func (SecretReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretReplication)(nil)).Elem()
}

func (o SecretReplicationOutput) ToSecretReplicationOutput() SecretReplicationOutput {
	return o
}

func (o SecretReplicationOutput) ToSecretReplicationOutputWithContext(ctx context.Context) SecretReplicationOutput {
	return o
}

func (o SecretReplicationOutput) ToSecretReplicationPtrOutput() SecretReplicationPtrOutput {
	return o.ToSecretReplicationPtrOutputWithContext(context.Background())
}

func (o SecretReplicationOutput) ToSecretReplicationPtrOutputWithContext(ctx context.Context) SecretReplicationPtrOutput {
	return o.ApplyT(func(v SecretReplication) *SecretReplication {
		return &v
	}).(SecretReplicationPtrOutput)
}

// -
// (Optional)
// The Secret will automatically be replicated without any restrictions.
func (o SecretReplicationOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecretReplication) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

// -
// (Optional)
// The Secret will automatically be replicated without any restrictions.  Structure is documented below.
func (o SecretReplicationOutput) UserManaged() SecretReplicationUserManagedPtrOutput {
	return o.ApplyT(func(v SecretReplication) *SecretReplicationUserManaged { return v.UserManaged }).(SecretReplicationUserManagedPtrOutput)
}

type SecretReplicationPtrOutput struct{ *pulumi.OutputState }

func (SecretReplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretReplication)(nil)).Elem()
}

func (o SecretReplicationPtrOutput) ToSecretReplicationPtrOutput() SecretReplicationPtrOutput {
	return o
}

func (o SecretReplicationPtrOutput) ToSecretReplicationPtrOutputWithContext(ctx context.Context) SecretReplicationPtrOutput {
	return o
}

func (o SecretReplicationPtrOutput) Elem() SecretReplicationOutput {
	return o.ApplyT(func(v *SecretReplication) SecretReplication { return *v }).(SecretReplicationOutput)
}

// -
// (Optional)
// The Secret will automatically be replicated without any restrictions.
func (o SecretReplicationPtrOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecretReplication) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

// -
// (Optional)
// The Secret will automatically be replicated without any restrictions.  Structure is documented below.
func (o SecretReplicationPtrOutput) UserManaged() SecretReplicationUserManagedPtrOutput {
	return o.ApplyT(func(v SecretReplication) *SecretReplicationUserManaged { return v.UserManaged }).(SecretReplicationUserManagedPtrOutput)
}

type SecretReplicationUserManaged struct {
	// -
	// (Required)
	// The list of Replicas for this Secret. Cannot be empty.  Structure is documented below.
	Replicas []SecretReplicationUserManagedReplica `pulumi:"replicas"`
}

// SecretReplicationUserManagedInput is an input type that accepts SecretReplicationUserManagedArgs and SecretReplicationUserManagedOutput values.
// You can construct a concrete instance of `SecretReplicationUserManagedInput` via:
//
// 		 SecretReplicationUserManagedArgs{...}
//
type SecretReplicationUserManagedInput interface {
	pulumi.Input

	ToSecretReplicationUserManagedOutput() SecretReplicationUserManagedOutput
	ToSecretReplicationUserManagedOutputWithContext(context.Context) SecretReplicationUserManagedOutput
}

type SecretReplicationUserManagedArgs struct {
	// -
	// (Required)
	// The list of Replicas for this Secret. Cannot be empty.  Structure is documented below.
	Replicas SecretReplicationUserManagedReplicaArrayInput `pulumi:"replicas"`
}

func (SecretReplicationUserManagedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretReplicationUserManaged)(nil)).Elem()
}

func (i SecretReplicationUserManagedArgs) ToSecretReplicationUserManagedOutput() SecretReplicationUserManagedOutput {
	return i.ToSecretReplicationUserManagedOutputWithContext(context.Background())
}

func (i SecretReplicationUserManagedArgs) ToSecretReplicationUserManagedOutputWithContext(ctx context.Context) SecretReplicationUserManagedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicationUserManagedOutput)
}

func (i SecretReplicationUserManagedArgs) ToSecretReplicationUserManagedPtrOutput() SecretReplicationUserManagedPtrOutput {
	return i.ToSecretReplicationUserManagedPtrOutputWithContext(context.Background())
}

func (i SecretReplicationUserManagedArgs) ToSecretReplicationUserManagedPtrOutputWithContext(ctx context.Context) SecretReplicationUserManagedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicationUserManagedOutput).ToSecretReplicationUserManagedPtrOutputWithContext(ctx)
}

// SecretReplicationUserManagedPtrInput is an input type that accepts SecretReplicationUserManagedArgs, SecretReplicationUserManagedPtr and SecretReplicationUserManagedPtrOutput values.
// You can construct a concrete instance of `SecretReplicationUserManagedPtrInput` via:
//
// 		 SecretReplicationUserManagedArgs{...}
//
//  or:
//
// 		 nil
//
type SecretReplicationUserManagedPtrInput interface {
	pulumi.Input

	ToSecretReplicationUserManagedPtrOutput() SecretReplicationUserManagedPtrOutput
	ToSecretReplicationUserManagedPtrOutputWithContext(context.Context) SecretReplicationUserManagedPtrOutput
}

type secretReplicationUserManagedPtrType SecretReplicationUserManagedArgs

func SecretReplicationUserManagedPtr(v *SecretReplicationUserManagedArgs) SecretReplicationUserManagedPtrInput {
	return (*secretReplicationUserManagedPtrType)(v)
}

func (*secretReplicationUserManagedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretReplicationUserManaged)(nil)).Elem()
}

func (i *secretReplicationUserManagedPtrType) ToSecretReplicationUserManagedPtrOutput() SecretReplicationUserManagedPtrOutput {
	return i.ToSecretReplicationUserManagedPtrOutputWithContext(context.Background())
}

func (i *secretReplicationUserManagedPtrType) ToSecretReplicationUserManagedPtrOutputWithContext(ctx context.Context) SecretReplicationUserManagedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicationUserManagedPtrOutput)
}

type SecretReplicationUserManagedOutput struct{ *pulumi.OutputState }

func (SecretReplicationUserManagedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretReplicationUserManaged)(nil)).Elem()
}

func (o SecretReplicationUserManagedOutput) ToSecretReplicationUserManagedOutput() SecretReplicationUserManagedOutput {
	return o
}

func (o SecretReplicationUserManagedOutput) ToSecretReplicationUserManagedOutputWithContext(ctx context.Context) SecretReplicationUserManagedOutput {
	return o
}

func (o SecretReplicationUserManagedOutput) ToSecretReplicationUserManagedPtrOutput() SecretReplicationUserManagedPtrOutput {
	return o.ToSecretReplicationUserManagedPtrOutputWithContext(context.Background())
}

func (o SecretReplicationUserManagedOutput) ToSecretReplicationUserManagedPtrOutputWithContext(ctx context.Context) SecretReplicationUserManagedPtrOutput {
	return o.ApplyT(func(v SecretReplicationUserManaged) *SecretReplicationUserManaged {
		return &v
	}).(SecretReplicationUserManagedPtrOutput)
}

// -
// (Required)
// The list of Replicas for this Secret. Cannot be empty.  Structure is documented below.
func (o SecretReplicationUserManagedOutput) Replicas() SecretReplicationUserManagedReplicaArrayOutput {
	return o.ApplyT(func(v SecretReplicationUserManaged) []SecretReplicationUserManagedReplica { return v.Replicas }).(SecretReplicationUserManagedReplicaArrayOutput)
}

type SecretReplicationUserManagedPtrOutput struct{ *pulumi.OutputState }

func (SecretReplicationUserManagedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretReplicationUserManaged)(nil)).Elem()
}

func (o SecretReplicationUserManagedPtrOutput) ToSecretReplicationUserManagedPtrOutput() SecretReplicationUserManagedPtrOutput {
	return o
}

func (o SecretReplicationUserManagedPtrOutput) ToSecretReplicationUserManagedPtrOutputWithContext(ctx context.Context) SecretReplicationUserManagedPtrOutput {
	return o
}

func (o SecretReplicationUserManagedPtrOutput) Elem() SecretReplicationUserManagedOutput {
	return o.ApplyT(func(v *SecretReplicationUserManaged) SecretReplicationUserManaged { return *v }).(SecretReplicationUserManagedOutput)
}

// -
// (Required)
// The list of Replicas for this Secret. Cannot be empty.  Structure is documented below.
func (o SecretReplicationUserManagedPtrOutput) Replicas() SecretReplicationUserManagedReplicaArrayOutput {
	return o.ApplyT(func(v SecretReplicationUserManaged) []SecretReplicationUserManagedReplica { return v.Replicas }).(SecretReplicationUserManagedReplicaArrayOutput)
}

type SecretReplicationUserManagedReplica struct {
	// -
	// (Required)
	// The canonical IDs of the location to replicate data. For example: "us-east1".
	Location string `pulumi:"location"`
}

// SecretReplicationUserManagedReplicaInput is an input type that accepts SecretReplicationUserManagedReplicaArgs and SecretReplicationUserManagedReplicaOutput values.
// You can construct a concrete instance of `SecretReplicationUserManagedReplicaInput` via:
//
// 		 SecretReplicationUserManagedReplicaArgs{...}
//
type SecretReplicationUserManagedReplicaInput interface {
	pulumi.Input

	ToSecretReplicationUserManagedReplicaOutput() SecretReplicationUserManagedReplicaOutput
	ToSecretReplicationUserManagedReplicaOutputWithContext(context.Context) SecretReplicationUserManagedReplicaOutput
}

type SecretReplicationUserManagedReplicaArgs struct {
	// -
	// (Required)
	// The canonical IDs of the location to replicate data. For example: "us-east1".
	Location pulumi.StringInput `pulumi:"location"`
}

func (SecretReplicationUserManagedReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretReplicationUserManagedReplica)(nil)).Elem()
}

func (i SecretReplicationUserManagedReplicaArgs) ToSecretReplicationUserManagedReplicaOutput() SecretReplicationUserManagedReplicaOutput {
	return i.ToSecretReplicationUserManagedReplicaOutputWithContext(context.Background())
}

func (i SecretReplicationUserManagedReplicaArgs) ToSecretReplicationUserManagedReplicaOutputWithContext(ctx context.Context) SecretReplicationUserManagedReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicationUserManagedReplicaOutput)
}

// SecretReplicationUserManagedReplicaArrayInput is an input type that accepts SecretReplicationUserManagedReplicaArray and SecretReplicationUserManagedReplicaArrayOutput values.
// You can construct a concrete instance of `SecretReplicationUserManagedReplicaArrayInput` via:
//
// 		 SecretReplicationUserManagedReplicaArray{ SecretReplicationUserManagedReplicaArgs{...} }
//
type SecretReplicationUserManagedReplicaArrayInput interface {
	pulumi.Input

	ToSecretReplicationUserManagedReplicaArrayOutput() SecretReplicationUserManagedReplicaArrayOutput
	ToSecretReplicationUserManagedReplicaArrayOutputWithContext(context.Context) SecretReplicationUserManagedReplicaArrayOutput
}

type SecretReplicationUserManagedReplicaArray []SecretReplicationUserManagedReplicaInput

func (SecretReplicationUserManagedReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretReplicationUserManagedReplica)(nil)).Elem()
}

func (i SecretReplicationUserManagedReplicaArray) ToSecretReplicationUserManagedReplicaArrayOutput() SecretReplicationUserManagedReplicaArrayOutput {
	return i.ToSecretReplicationUserManagedReplicaArrayOutputWithContext(context.Background())
}

func (i SecretReplicationUserManagedReplicaArray) ToSecretReplicationUserManagedReplicaArrayOutputWithContext(ctx context.Context) SecretReplicationUserManagedReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicationUserManagedReplicaArrayOutput)
}

type SecretReplicationUserManagedReplicaOutput struct{ *pulumi.OutputState }

func (SecretReplicationUserManagedReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretReplicationUserManagedReplica)(nil)).Elem()
}

func (o SecretReplicationUserManagedReplicaOutput) ToSecretReplicationUserManagedReplicaOutput() SecretReplicationUserManagedReplicaOutput {
	return o
}

func (o SecretReplicationUserManagedReplicaOutput) ToSecretReplicationUserManagedReplicaOutputWithContext(ctx context.Context) SecretReplicationUserManagedReplicaOutput {
	return o
}

// -
// (Required)
// The canonical IDs of the location to replicate data. For example: "us-east1".
func (o SecretReplicationUserManagedReplicaOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v SecretReplicationUserManagedReplica) string { return v.Location }).(pulumi.StringOutput)
}

type SecretReplicationUserManagedReplicaArrayOutput struct{ *pulumi.OutputState }

func (SecretReplicationUserManagedReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretReplicationUserManagedReplica)(nil)).Elem()
}

func (o SecretReplicationUserManagedReplicaArrayOutput) ToSecretReplicationUserManagedReplicaArrayOutput() SecretReplicationUserManagedReplicaArrayOutput {
	return o
}

func (o SecretReplicationUserManagedReplicaArrayOutput) ToSecretReplicationUserManagedReplicaArrayOutputWithContext(ctx context.Context) SecretReplicationUserManagedReplicaArrayOutput {
	return o
}

func (o SecretReplicationUserManagedReplicaArrayOutput) Index(i pulumi.IntInput) SecretReplicationUserManagedReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretReplicationUserManagedReplica {
		return vs[0].([]SecretReplicationUserManagedReplica)[vs[1].(int)]
	}).(SecretReplicationUserManagedReplicaOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretIamBindingConditionOutput{})
	pulumi.RegisterOutputType(SecretIamBindingConditionPtrOutput{})
	pulumi.RegisterOutputType(SecretIamMemberConditionOutput{})
	pulumi.RegisterOutputType(SecretIamMemberConditionPtrOutput{})
	pulumi.RegisterOutputType(SecretReplicationOutput{})
	pulumi.RegisterOutputType(SecretReplicationPtrOutput{})
	pulumi.RegisterOutputType(SecretReplicationUserManagedOutput{})
	pulumi.RegisterOutputType(SecretReplicationUserManagedPtrOutput{})
	pulumi.RegisterOutputType(SecretReplicationUserManagedReplicaOutput{})
	pulumi.RegisterOutputType(SecretReplicationUserManagedReplicaArrayOutput{})
}
