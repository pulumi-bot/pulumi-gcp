// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of Organization policies for a Google Project. For more information see
// [the official
// documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		policy, err := projects.LookupOrganizationPolicy(ctx, &projects.LookupOrganizationPolicyArgs{
// 			Project:    "project-id",
// 			Constraint: "constraints/serviceuser.services",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("version", policy.Version)
// 		return nil
// 	})
// }
// ```
func LookupOrganizationPolicy(ctx *pulumi.Context, args *LookupOrganizationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationPolicyResult, error) {
	var rv LookupOrganizationPolicyResult
	err := ctx.Invoke("gcp:projects/getOrganizationPolicy:getOrganizationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationPolicy.
type LookupOrganizationPolicyArgs struct {
	// (Required) The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint string `pulumi:"constraint"`
	// The project ID.
	Project string `pulumi:"project"`
}

// A collection of values returned by getOrganizationPolicy.
type LookupOrganizationPolicyResult struct {
	BooleanPolicies []GetOrganizationPolicyBooleanPolicy `pulumi:"booleanPolicies"`
	Constraint      string                               `pulumi:"constraint"`
	Etag            string                               `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                               `pulumi:"id"`
	ListPolicies    []GetOrganizationPolicyListPolicy    `pulumi:"listPolicies"`
	Project         string                               `pulumi:"project"`
	RestorePolicies []GetOrganizationPolicyRestorePolicy `pulumi:"restorePolicies"`
	UpdateTime      string                               `pulumi:"updateTime"`
	Version         int                                  `pulumi:"version"`
}

func LookupOrganizationPolicyOutput(ctx *pulumi.Context, args LookupOrganizationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationPolicyResult, error) {
			args := v.(LookupOrganizationPolicyArgs)
			r, err := LookupOrganizationPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupOrganizationPolicyResultOutput)
}

// A collection of arguments for invoking getOrganizationPolicy.
type LookupOrganizationPolicyOutputArgs struct {
	// (Required) The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringInput `pulumi:"constraint"`
	// The project ID.
	Project pulumi.StringInput `pulumi:"project"`
}

func (LookupOrganizationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getOrganizationPolicy.
type LookupOrganizationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationPolicyResult)(nil)).Elem()
}

func (o LookupOrganizationPolicyResultOutput) ToLookupOrganizationPolicyResultOutput() LookupOrganizationPolicyResultOutput {
	return o
}

func (o LookupOrganizationPolicyResultOutput) ToLookupOrganizationPolicyResultOutputWithContext(ctx context.Context) LookupOrganizationPolicyResultOutput {
	return o
}

func (o LookupOrganizationPolicyResultOutput) BooleanPolicies() GetOrganizationPolicyBooleanPolicyArrayOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) []GetOrganizationPolicyBooleanPolicy { return v.BooleanPolicies }).(GetOrganizationPolicyBooleanPolicyArrayOutput)
}

func (o LookupOrganizationPolicyResultOutput) Constraint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) string { return v.Constraint }).(pulumi.StringOutput)
}

func (o LookupOrganizationPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrganizationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrganizationPolicyResultOutput) ListPolicies() GetOrganizationPolicyListPolicyArrayOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) []GetOrganizationPolicyListPolicy { return v.ListPolicies }).(GetOrganizationPolicyListPolicyArrayOutput)
}

func (o LookupOrganizationPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupOrganizationPolicyResultOutput) RestorePolicies() GetOrganizationPolicyRestorePolicyArrayOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) []GetOrganizationPolicyRestorePolicy { return v.RestorePolicies }).(GetOrganizationPolicyRestorePolicyArrayOutput)
}

func (o LookupOrganizationPolicyResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func (o LookupOrganizationPolicyResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrganizationPolicyResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationPolicyResultOutput{})
}
