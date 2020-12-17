// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy Web. Each of these resources serves a different use case:
//
// * `iap.WebIamPolicy`: Authoritative. Sets the IAM policy for the web and replaces any existing policy already attached.
// * `iap.WebIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the web are preserved.
// * `iap.WebIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the web are preserved.
//
// > **Note:** `iap.WebIamPolicy` **cannot** be used in conjunction with `iap.WebIamBinding` and `iap.WebIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebIamBinding` resources **can be** used in conjunction with `iap.WebIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_web\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.httpsResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewWebIamPolicy(ctx, "policy", &iap.WebIamPolicyArgs{
// 			Project:    pulumi.Any(google_project_service.Project_service.Project),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.httpsResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewWebIamPolicy(ctx, "policy", &iap.WebIamPolicyArgs{
// 			Project:    pulumi.Any(google_project_service.Project_service.Project),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_web\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebIamBinding(ctx, "binding", &iap.WebIamBindingArgs{
// 			Project: pulumi.Any(google_project_service.Project_service.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebIamBinding(ctx, "binding", &iap.WebIamBindingArgs{
// 			Project: pulumi.Any(google_project_service.Project_service.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &iap.WebIamBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_web\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebIamMember(ctx, "member", &iap.WebIamMemberArgs{
// 			Project: pulumi.Any(google_project_service.Project_service.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebIamMember(ctx, "member", &iap.WebIamMemberArgs{
// 			Project: pulumi.Any(google_project_service.Project_service.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Condition: &iap.WebIamMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web * {{project}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy web IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webIamBinding:WebIamBinding editor "projects/{{project}}/iap_web roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webIamBinding:WebIamBinding editor "projects/{{project}}/iap_web roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webIamBinding:WebIamBinding editor projects/{{project}}/iap_web
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebIamBinding struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewWebIamBinding registers a new resource with the given unique name, arguments, and options.
func NewWebIamBinding(ctx *pulumi.Context,
	name string, args *WebIamBindingArgs, opts ...pulumi.ResourceOption) (*WebIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource WebIamBinding
	err := ctx.RegisterResource("gcp:iap/webIamBinding:WebIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebIamBinding gets an existing WebIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebIamBindingState, opts ...pulumi.ResourceOption) (*WebIamBinding, error) {
	var resource WebIamBinding
	err := ctx.ReadResource("gcp:iap/webIamBinding:WebIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebIamBinding resources.
type webIamBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type WebIamBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (WebIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webIamBindingState)(nil)).Elem()
}

type webIamBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebIamBindingCondition `pulumi:"condition"`
	Members   []string                `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a WebIamBinding resource.
type WebIamBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (WebIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webIamBindingArgs)(nil)).Elem()
}

type WebIamBindingInput interface {
	pulumi.Input

	ToWebIamBindingOutput() WebIamBindingOutput
	ToWebIamBindingOutputWithContext(ctx context.Context) WebIamBindingOutput
}

func (*WebIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*WebIamBinding)(nil))
}

func (i *WebIamBinding) ToWebIamBindingOutput() WebIamBindingOutput {
	return i.ToWebIamBindingOutputWithContext(context.Background())
}

func (i *WebIamBinding) ToWebIamBindingOutputWithContext(ctx context.Context) WebIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamBindingOutput)
}

func (i *WebIamBinding) ToWebIamBindingPtrOutput() WebIamBindingPtrOutput {
	return i.ToWebIamBindingPtrOutputWithContext(context.Background())
}

func (i *WebIamBinding) ToWebIamBindingPtrOutputWithContext(ctx context.Context) WebIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamBindingPtrOutput)
}

type WebIamBindingPtrInput interface {
	pulumi.Input

	ToWebIamBindingPtrOutput() WebIamBindingPtrOutput
	ToWebIamBindingPtrOutputWithContext(ctx context.Context) WebIamBindingPtrOutput
}

type webIamBindingPtrType WebIamBindingArgs

func (*webIamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebIamBinding)(nil))
}

func (i *webIamBindingPtrType) ToWebIamBindingPtrOutput() WebIamBindingPtrOutput {
	return i.ToWebIamBindingPtrOutputWithContext(context.Background())
}

func (i *webIamBindingPtrType) ToWebIamBindingPtrOutputWithContext(ctx context.Context) WebIamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamBindingOutput).ToWebIamBindingPtrOutput()
}

// WebIamBindingArrayInput is an input type that accepts WebIamBindingArray and WebIamBindingArrayOutput values.
// You can construct a concrete instance of `WebIamBindingArrayInput` via:
//
//          WebIamBindingArray{ WebIamBindingArgs{...} }
type WebIamBindingArrayInput interface {
	pulumi.Input

	ToWebIamBindingArrayOutput() WebIamBindingArrayOutput
	ToWebIamBindingArrayOutputWithContext(context.Context) WebIamBindingArrayOutput
}

type WebIamBindingArray []WebIamBindingInput

func (WebIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebIamBinding)(nil))
}

func (i WebIamBindingArray) ToWebIamBindingArrayOutput() WebIamBindingArrayOutput {
	return i.ToWebIamBindingArrayOutputWithContext(context.Background())
}

func (i WebIamBindingArray) ToWebIamBindingArrayOutputWithContext(ctx context.Context) WebIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamBindingArrayOutput)
}

// WebIamBindingMapInput is an input type that accepts WebIamBindingMap and WebIamBindingMapOutput values.
// You can construct a concrete instance of `WebIamBindingMapInput` via:
//
//          WebIamBindingMap{ "key": WebIamBindingArgs{...} }
type WebIamBindingMapInput interface {
	pulumi.Input

	ToWebIamBindingMapOutput() WebIamBindingMapOutput
	ToWebIamBindingMapOutputWithContext(context.Context) WebIamBindingMapOutput
}

type WebIamBindingMap map[string]WebIamBindingInput

func (WebIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebIamBinding)(nil))
}

func (i WebIamBindingMap) ToWebIamBindingMapOutput() WebIamBindingMapOutput {
	return i.ToWebIamBindingMapOutputWithContext(context.Background())
}

func (i WebIamBindingMap) ToWebIamBindingMapOutputWithContext(ctx context.Context) WebIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamBindingMapOutput)
}

type WebIamBindingOutput struct {
	*pulumi.OutputState
}

func (WebIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebIamBinding)(nil))
}

func (o WebIamBindingOutput) ToWebIamBindingOutput() WebIamBindingOutput {
	return o
}

func (o WebIamBindingOutput) ToWebIamBindingOutputWithContext(ctx context.Context) WebIamBindingOutput {
	return o
}

func (o WebIamBindingOutput) ToWebIamBindingPtrOutput() WebIamBindingPtrOutput {
	return o.ToWebIamBindingPtrOutputWithContext(context.Background())
}

func (o WebIamBindingOutput) ToWebIamBindingPtrOutputWithContext(ctx context.Context) WebIamBindingPtrOutput {
	return o.ApplyT(func(v WebIamBinding) *WebIamBinding {
		return &v
	}).(WebIamBindingPtrOutput)
}

type WebIamBindingPtrOutput struct {
	*pulumi.OutputState
}

func (WebIamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebIamBinding)(nil))
}

func (o WebIamBindingPtrOutput) ToWebIamBindingPtrOutput() WebIamBindingPtrOutput {
	return o
}

func (o WebIamBindingPtrOutput) ToWebIamBindingPtrOutputWithContext(ctx context.Context) WebIamBindingPtrOutput {
	return o
}

type WebIamBindingArrayOutput struct{ *pulumi.OutputState }

func (WebIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebIamBinding)(nil))
}

func (o WebIamBindingArrayOutput) ToWebIamBindingArrayOutput() WebIamBindingArrayOutput {
	return o
}

func (o WebIamBindingArrayOutput) ToWebIamBindingArrayOutputWithContext(ctx context.Context) WebIamBindingArrayOutput {
	return o
}

func (o WebIamBindingArrayOutput) Index(i pulumi.IntInput) WebIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebIamBinding {
		return vs[0].([]WebIamBinding)[vs[1].(int)]
	}).(WebIamBindingOutput)
}

type WebIamBindingMapOutput struct{ *pulumi.OutputState }

func (WebIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebIamBinding)(nil))
}

func (o WebIamBindingMapOutput) ToWebIamBindingMapOutput() WebIamBindingMapOutput {
	return o
}

func (o WebIamBindingMapOutput) ToWebIamBindingMapOutputWithContext(ctx context.Context) WebIamBindingMapOutput {
	return o
}

func (o WebIamBindingMapOutput) MapIndex(k pulumi.StringInput) WebIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebIamBinding {
		return vs[0].(map[string]WebIamBinding)[vs[1].(string)]
	}).(WebIamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(WebIamBindingOutput{})
	pulumi.RegisterOutputType(WebIamBindingPtrOutput{})
	pulumi.RegisterOutputType(WebIamBindingArrayOutput{})
	pulumi.RegisterOutputType(WebIamBindingMapOutput{})
}
