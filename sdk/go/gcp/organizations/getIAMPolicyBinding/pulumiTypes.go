// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getIAMPolicyBinding

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/organizations/getIAMPolicyBindingCondition"
)

type GetIAMPolicyBinding struct {
	Condition *organizationsgetIAMPolicyBindingCondition.GetIAMPolicyBindingCondition `pulumi:"condition"`
	// An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. It **can't** be used with the `organizations.Project` resource.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. It **can't** be used with the `organizations.Project` resource.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members []string `pulumi:"members"`
	// The role/permission that will be granted to the members.
	// See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
	// Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

type GetIAMPolicyBindingInput interface {
	pulumi.Input

	ToGetIAMPolicyBindingOutput() GetIAMPolicyBindingOutput
	ToGetIAMPolicyBindingOutputWithContext(context.Context) GetIAMPolicyBindingOutput
}

type GetIAMPolicyBindingArgs struct {
	Condition organizationsgetIAMPolicyBindingCondition.GetIAMPolicyBindingConditionPtrInput `pulumi:"condition"`
	// An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. It **can't** be used with the `organizations.Project` resource.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. It **can't** be used with the `organizations.Project` resource.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayInput `pulumi:"members"`
	// The role/permission that will be granted to the members.
	// See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
	// Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput `pulumi:"role"`
}

func (GetIAMPolicyBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIAMPolicyBinding)(nil)).Elem()
}

func (i GetIAMPolicyBindingArgs) ToGetIAMPolicyBindingOutput() GetIAMPolicyBindingOutput {
	return i.ToGetIAMPolicyBindingOutputWithContext(context.Background())
}

func (i GetIAMPolicyBindingArgs) ToGetIAMPolicyBindingOutputWithContext(ctx context.Context) GetIAMPolicyBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetIAMPolicyBindingOutput)
}

type GetIAMPolicyBindingArrayInput interface {
	pulumi.Input

	ToGetIAMPolicyBindingArrayOutput() GetIAMPolicyBindingArrayOutput
	ToGetIAMPolicyBindingArrayOutputWithContext(context.Context) GetIAMPolicyBindingArrayOutput
}

type GetIAMPolicyBindingArray []GetIAMPolicyBindingInput

func (GetIAMPolicyBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetIAMPolicyBinding)(nil)).Elem()
}

func (i GetIAMPolicyBindingArray) ToGetIAMPolicyBindingArrayOutput() GetIAMPolicyBindingArrayOutput {
	return i.ToGetIAMPolicyBindingArrayOutputWithContext(context.Background())
}

func (i GetIAMPolicyBindingArray) ToGetIAMPolicyBindingArrayOutputWithContext(ctx context.Context) GetIAMPolicyBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetIAMPolicyBindingArrayOutput)
}

type GetIAMPolicyBindingOutput struct { *pulumi.OutputState }

func (GetIAMPolicyBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIAMPolicyBinding)(nil)).Elem()
}

func (o GetIAMPolicyBindingOutput) ToGetIAMPolicyBindingOutput() GetIAMPolicyBindingOutput {
	return o
}

func (o GetIAMPolicyBindingOutput) ToGetIAMPolicyBindingOutputWithContext(ctx context.Context) GetIAMPolicyBindingOutput {
	return o
}

func (o GetIAMPolicyBindingOutput) Condition() organizationsgetIAMPolicyBindingCondition.GetIAMPolicyBindingConditionPtrOutput {
	return o.ApplyT(func (v GetIAMPolicyBinding) *organizationsgetIAMPolicyBindingCondition.GetIAMPolicyBindingCondition { return v.Condition }).(organizationsgetIAMPolicyBindingCondition.GetIAMPolicyBindingConditionPtrOutput)
}

// An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
// Each entry can have one of the following values:
// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. It **can't** be used with the `organizations.Project` resource.
// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. It **can't** be used with the `organizations.Project` resource.
// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o GetIAMPolicyBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetIAMPolicyBinding) []string { return v.Members }).(pulumi.StringArrayOutput)
}

// The role/permission that will be granted to the members.
// See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
// Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o GetIAMPolicyBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func (v GetIAMPolicyBinding) string { return v.Role }).(pulumi.StringOutput)
}

type GetIAMPolicyBindingArrayOutput struct { *pulumi.OutputState}

func (GetIAMPolicyBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetIAMPolicyBinding)(nil)).Elem()
}

func (o GetIAMPolicyBindingArrayOutput) ToGetIAMPolicyBindingArrayOutput() GetIAMPolicyBindingArrayOutput {
	return o
}

func (o GetIAMPolicyBindingArrayOutput) ToGetIAMPolicyBindingArrayOutputWithContext(ctx context.Context) GetIAMPolicyBindingArrayOutput {
	return o
}

func (o GetIAMPolicyBindingArrayOutput) Index(i pulumi.IntInput) GetIAMPolicyBindingOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetIAMPolicyBinding {
		return vs[0].([]GetIAMPolicyBinding)[vs[1].(int)]
	}).(GetIAMPolicyBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIAMPolicyBindingOutput{})
	pulumi.RegisterOutputType(GetIAMPolicyBindingArrayOutput{})
}
