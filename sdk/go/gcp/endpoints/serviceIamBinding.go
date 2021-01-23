// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package endpoints

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Endpoints Service. Each of these resources serves a different use case:
//
// * `endpoints.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `endpoints.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `endpoints.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
//
// > **Note:** `endpoints.ServiceIamPolicy` **cannot** be used in conjunction with `endpoints.ServiceIamBinding` and `endpoints.ServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `endpoints.ServiceIamBinding` resources **can be** used in conjunction with `endpoints.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_endpoints\_service\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/endpoints"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = endpoints.NewServiceIamPolicy(ctx, "policy", &endpoints.ServiceIamPolicyArgs{
// 			ServiceName: pulumi.Any(google_endpoints_service.Endpoints_service.Service_name),
// 			PolicyData:  pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_endpoints\_service\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/endpoints"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := endpoints.NewServiceIamBinding(ctx, "binding", &endpoints.ServiceIamBindingArgs{
// 			ServiceName: pulumi.Any(google_endpoints_service.Endpoints_service.Service_name),
// 			Role:        pulumi.String("roles/viewer"),
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
// ## google\_endpoints\_service\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/endpoints"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := endpoints.NewServiceIamMember(ctx, "member", &endpoints.ServiceIamMemberArgs{
// 			ServiceName: pulumi.Any(google_endpoints_service.Endpoints_service.Service_name),
// 			Role:        pulumi.String("roles/viewer"),
// 			Member:      pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* services/{{serviceName}} * {{serviceName}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Endpoints service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:endpoints/serviceIamBinding:ServiceIamBinding editor "services/{{serviceName}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:endpoints/serviceIamBinding:ServiceIamBinding editor "services/{{serviceName}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:endpoints/serviceIamBinding:ServiceIamBinding editor services/{{serviceName}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ServiceIamBinding struct {
	pulumi.CustomResourceState

	Condition ServiceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewServiceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewServiceIamBinding(ctx *pulumi.Context,
	name string, args *ServiceIamBindingArgs, opts ...pulumi.ResourceOption) (*ServiceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource ServiceIamBinding
	err := ctx.RegisterResource("gcp:endpoints/serviceIamBinding:ServiceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceIamBinding gets an existing ServiceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceIamBindingState, opts ...pulumi.ResourceOption) (*ServiceIamBinding, error) {
	var resource ServiceIamBinding
	err := ctx.ReadResource("gcp:endpoints/serviceIamBinding:ServiceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceIamBinding resources.
type serviceIamBindingState struct {
	Condition *ServiceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName *string `pulumi:"serviceName"`
}

type ServiceIamBindingState struct {
	Condition ServiceIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringPtrInput
}

func (ServiceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamBindingState)(nil)).Elem()
}

type serviceIamBindingArgs struct {
	Condition *ServiceIamBindingCondition `pulumi:"condition"`
	Members   []string                    `pulumi:"members"`
	// The role that should be applied. Only one
	// `endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ServiceIamBinding resource.
type ServiceIamBindingArgs struct {
	Condition ServiceIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringInput
}

func (ServiceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamBindingArgs)(nil)).Elem()
}

type ServiceIamBindingInput interface {
	pulumi.Input

	ToServiceIamBindingOutput() ServiceIamBindingOutput
	ToServiceIamBindingOutputWithContext(ctx context.Context) ServiceIamBindingOutput
}

func (*ServiceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIamBinding)(nil))
}

func (i *ServiceIamBinding) ToServiceIamBindingOutput() ServiceIamBindingOutput {
	return i.ToServiceIamBindingOutputWithContext(context.Background())
}

func (i *ServiceIamBinding) ToServiceIamBindingOutputWithContext(ctx context.Context) ServiceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamBindingOutput)
}

type ServiceIamBindingOutput struct {
	*pulumi.OutputState
}

func (ServiceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIamBinding)(nil))
}

func (o ServiceIamBindingOutput) ToServiceIamBindingOutput() ServiceIamBindingOutput {
	return o
}

func (o ServiceIamBindingOutput) ToServiceIamBindingOutputWithContext(ctx context.Context) ServiceIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceIamBindingOutput{})
}
