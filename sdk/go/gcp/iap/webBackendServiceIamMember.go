// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:
//
// * `iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
// * `iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
// * `iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.
//
// > **Note:** `iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `iap.WebBackendServiceIamBinding` and `iap.WebBackendServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_web\_backend\_service\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		_, err = iap.NewWebBackendServiceIamPolicy(ctx, "policy", &iap.WebBackendServiceIamPolicyArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			PolicyData:        pulumi.String(admin.PolicyData),
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		_, err = iap.NewWebBackendServiceIamPolicy(ctx, "policy", &iap.WebBackendServiceIamPolicyArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			PolicyData:        pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_web\_backend\_service\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &iap.WebBackendServiceIamBindingConditionArgs{
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
// ## google\_iap\_web\_backend\_service\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:            pulumi.String("user:jane@example.com"),
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:            pulumi.String("user:jane@example.com"),
// 			Condition: &iap.WebBackendServiceIamMemberConditionArgs{
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute/services/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webbackendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember editor projects/{{project}}/iap_web/compute/services/{{web_backend_service}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebBackendServiceIamMember struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringOutput `pulumi:"webBackendService"`
}

// NewWebBackendServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewWebBackendServiceIamMember(ctx *pulumi.Context,
	name string, args *WebBackendServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*WebBackendServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WebBackendService == nil {
		return nil, errors.New("invalid value for required argument 'WebBackendService'")
	}
	var resource WebBackendServiceIamMember
	err := ctx.RegisterResource("gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebBackendServiceIamMember gets an existing WebBackendServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebBackendServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebBackendServiceIamMemberState, opts ...pulumi.ResourceOption) (*WebBackendServiceIamMember, error) {
	var resource WebBackendServiceIamMember
	err := ctx.ReadResource("gcp:iap/webBackendServiceIamMember:WebBackendServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebBackendServiceIamMember resources.
type webBackendServiceIamMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebBackendServiceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService *string `pulumi:"webBackendService"`
}

type WebBackendServiceIamMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringPtrInput
}

func (WebBackendServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamMemberState)(nil)).Elem()
}

type webBackendServiceIamMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebBackendServiceIamMemberCondition `pulumi:"condition"`
	Member    string                               `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService string `pulumi:"webBackendService"`
}

// The set of arguments for constructing a WebBackendServiceIamMember resource.
type WebBackendServiceIamMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebBackendServiceIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringInput
}

func (WebBackendServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamMemberArgs)(nil)).Elem()
}

type WebBackendServiceIamMemberInput interface {
	pulumi.Input

	ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput
	ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput
}

func (WebBackendServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*WebBackendServiceIamMember)(nil)).Elem()
}

func (i WebBackendServiceIamMember) ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput {
	return i.ToWebBackendServiceIamMemberOutputWithContext(context.Background())
}

func (i WebBackendServiceIamMember) ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebBackendServiceIamMemberOutput)
}

type WebBackendServiceIamMemberOutput struct {
	*pulumi.OutputState
}

func (WebBackendServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebBackendServiceIamMemberOutput)(nil)).Elem()
}

func (o WebBackendServiceIamMemberOutput) ToWebBackendServiceIamMemberOutput() WebBackendServiceIamMemberOutput {
	return o
}

func (o WebBackendServiceIamMemberOutput) ToWebBackendServiceIamMemberOutputWithContext(ctx context.Context) WebBackendServiceIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebBackendServiceIamMemberOutput{})
}
