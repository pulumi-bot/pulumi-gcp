// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows management of Organization policies for a Google Organization. For more information see
// [the official
// documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
// [API](https://cloud.google.com/resource-manager/reference/rest/v1/organizations/setOrgPolicy).
//
// ## Example Usage
//
// To set policy with a [boolean constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-boolean-constraints):
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewPolicy(ctx, "serialPortPolicy", &organizations.PolicyArgs{
// 			BooleanPolicy: &organizations.PolicyBooleanPolicyArgs{
// 				Enforced: pulumi.Bool(true),
// 			},
// 			Constraint: pulumi.String("compute.disableSerialPortAccess"),
// 			OrgId:      pulumi.String("123456789"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To set a policy with a [list constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-list-constraints):
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewPolicy(ctx, "servicesPolicy", &organizations.PolicyArgs{
// 			Constraint: pulumi.String("serviceuser.services"),
// 			ListPolicy: &organizations.PolicyListPolicyArgs{
// 				Allow: &organizations.PolicyListPolicyAllowArgs{
// 					All: pulumi.Bool(true),
// 				},
// 			},
// 			OrgId: pulumi.String("123456789"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Or to deny some services, use the following instead:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewPolicy(ctx, "servicesPolicy", &organizations.PolicyArgs{
// 			Constraint: pulumi.String("serviceuser.services"),
// 			ListPolicy: &organizations.PolicyListPolicyArgs{
// 				Deny: &organizations.PolicyListPolicyDenyArgs{
// 					Values: pulumi.StringArray{
// 						pulumi.String("cloudresourcemanager.googleapis.com"),
// 					},
// 				},
// 				SuggestedValue: pulumi.String("compute.googleapis.com"),
// 			},
// 			OrgId: pulumi.String("123456789"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To restore the default organization policy, use the following instead:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewPolicy(ctx, "servicesPolicy", &organizations.PolicyArgs{
// 			Constraint: pulumi.String("serviceuser.services"),
// 			OrgId:      pulumi.String("123456789"),
// 			RestorePolicy: &organizations.PolicyRestorePolicyArgs{
// 				Default: pulumi.Bool(true),
// 			},
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
// Organization Policies can be imported using the `org_id` and the `constraint`, e.g.
//
// ```sh
//  $ pulumi import gcp:organizations/policy:Policy services_policy 123456789/constraints/serviceuser.services
// ```
//
//  It is all right if the constraint contains a slash, as in the example above.
type Policy struct {
	pulumi.CustomResourceState

	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy PolicyBooleanPolicyPtrOutput `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringOutput `pulumi:"constraint"`
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy PolicyListPolicyPtrOutput `pulumi:"listPolicy"`
	// The numeric ID of the organization to set the policy for.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy PolicyRestorePolicyPtrOutput `pulumi:"restorePolicy"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Version of the Policy. Default version is 0.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Constraint == nil {
		return nil, errors.New("invalid value for required argument 'Constraint'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	var resource Policy
	err := ctx.RegisterResource("gcp:organizations/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("gcp:organizations/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy *PolicyBooleanPolicy `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint *string `pulumi:"constraint"`
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag *string `pulumi:"etag"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy *PolicyListPolicy `pulumi:"listPolicy"`
	// The numeric ID of the organization to set the policy for.
	OrgId *string `pulumi:"orgId"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy *PolicyRestorePolicy `pulumi:"restorePolicy"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime *string `pulumi:"updateTime"`
	// Version of the Policy. Default version is 0.
	Version *int `pulumi:"version"`
}

type PolicyState struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy PolicyBooleanPolicyPtrInput
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringPtrInput
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag pulumi.StringPtrInput
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy PolicyListPolicyPtrInput
	// The numeric ID of the organization to set the policy for.
	OrgId pulumi.StringPtrInput
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy PolicyRestorePolicyPtrInput
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringPtrInput
	// Version of the Policy. Default version is 0.
	Version pulumi.IntPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy *PolicyBooleanPolicy `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint string `pulumi:"constraint"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy *PolicyListPolicy `pulumi:"listPolicy"`
	// The numeric ID of the organization to set the policy for.
	OrgId string `pulumi:"orgId"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy *PolicyRestorePolicy `pulumi:"restorePolicy"`
	// Version of the Policy. Default version is 0.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy PolicyBooleanPolicyPtrInput
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringInput
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy PolicyListPolicyPtrInput
	// The numeric ID of the organization to set the policy for.
	OrgId pulumi.StringInput
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy PolicyRestorePolicyPtrInput
	// Version of the Policy. Default version is 0.
	Version pulumi.IntPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

func (i *Policy) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

type PolicyPtrInput interface {
	pulumi.Input

	ToPolicyPtrOutput() PolicyPtrOutput
	ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput
}

type policyPtrType PolicyArgs

func (*policyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (i *policyPtrType) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *policyPtrType) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput).ToPolicyPtrOutput()
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//          PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Policy)(nil))
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//          PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Policy)(nil))
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct {
	*pulumi.OutputState
}

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o.ToPolicyPtrOutputWithContext(context.Background())
}

func (o PolicyOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o.ApplyT(func(v Policy) *Policy {
		return &v
	}).(PolicyPtrOutput)
}

type PolicyPtrOutput struct {
	*pulumi.OutputState
}

func (PolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (o PolicyPtrOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Policy)(nil))
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Policy {
		return vs[0].([]Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Policy)(nil))
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Policy {
		return vs[0].(map[string]Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyPtrOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
