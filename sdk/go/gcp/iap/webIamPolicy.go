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
//  $ pulumi import gcp:iap/webIamPolicy:WebIamPolicy editor "projects/{{project}}/iap_web roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webIamPolicy:WebIamPolicy editor "projects/{{project}}/iap_web roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webIamPolicy:WebIamPolicy editor projects/{{project}}/iap_web
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewWebIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebIamPolicy(ctx *pulumi.Context,
	name string, args *WebIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource WebIamPolicy
	err := ctx.RegisterResource("gcp:iap/webIamPolicy:WebIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebIamPolicy gets an existing WebIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebIamPolicyState, opts ...pulumi.ResourceOption) (*WebIamPolicy, error) {
	var resource WebIamPolicy
	err := ctx.ReadResource("gcp:iap/webIamPolicy:WebIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebIamPolicy resources.
type webIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type WebIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webIamPolicyState)(nil)).Elem()
}

type webIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a WebIamPolicy resource.
type WebIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webIamPolicyArgs)(nil)).Elem()
}

type WebIamPolicyInput interface {
	pulumi.Input

	ToWebIamPolicyOutput() WebIamPolicyOutput
	ToWebIamPolicyOutputWithContext(ctx context.Context) WebIamPolicyOutput
}

func (WebIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*WebIamPolicy)(nil)).Elem()
}

func (i WebIamPolicy) ToWebIamPolicyOutput() WebIamPolicyOutput {
	return i.ToWebIamPolicyOutputWithContext(context.Background())
}

func (i WebIamPolicy) ToWebIamPolicyOutputWithContext(ctx context.Context) WebIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamPolicyOutput)
}

type WebIamPolicyOutput struct {
	*pulumi.OutputState
}

func (WebIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebIamPolicyOutput)(nil)).Elem()
}

func (o WebIamPolicyOutput) ToWebIamPolicyOutput() WebIamPolicyOutput {
	return o
}

func (o WebIamPolicyOutput) ToWebIamPolicyOutputWithContext(ctx context.Context) WebIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebIamPolicyOutput{})
}
