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
//  $ pulumi import gcp:endpoints/serviceIamPolicy:ServiceIamPolicy editor "services/{{serviceName}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:endpoints/serviceIamPolicy:ServiceIamPolicy editor "services/{{serviceName}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:endpoints/serviceIamPolicy:ServiceIamPolicy editor services/{{serviceName}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ServiceIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceIamPolicy(ctx *pulumi.Context,
	name string, args *ServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource ServiceIamPolicy
	err := ctx.RegisterResource("gcp:endpoints/serviceIamPolicy:ServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceIamPolicy gets an existing ServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceIamPolicyState, opts ...pulumi.ResourceOption) (*ServiceIamPolicy, error) {
	var resource ServiceIamPolicy
	err := ctx.ReadResource("gcp:endpoints/serviceIamPolicy:ServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceIamPolicy resources.
type serviceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName *string `pulumi:"serviceName"`
}

type ServiceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringPtrInput
}

func (ServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamPolicyState)(nil)).Elem()
}

type serviceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ServiceIamPolicy resource.
type ServiceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringInput
}

func (ServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamPolicyArgs)(nil)).Elem()
}

type ServiceIamPolicyInput interface {
	pulumi.Input

	ToServiceIamPolicyOutput() ServiceIamPolicyOutput
	ToServiceIamPolicyOutputWithContext(ctx context.Context) ServiceIamPolicyOutput
}

func (*ServiceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIamPolicy)(nil))
}

func (i *ServiceIamPolicy) ToServiceIamPolicyOutput() ServiceIamPolicyOutput {
	return i.ToServiceIamPolicyOutputWithContext(context.Background())
}

func (i *ServiceIamPolicy) ToServiceIamPolicyOutputWithContext(ctx context.Context) ServiceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamPolicyOutput)
}

func (i *ServiceIamPolicy) ToServiceIamPolicyPtrOutput() ServiceIamPolicyPtrOutput {
	return i.ToServiceIamPolicyPtrOutputWithContext(context.Background())
}

func (i *ServiceIamPolicy) ToServiceIamPolicyPtrOutputWithContext(ctx context.Context) ServiceIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamPolicyPtrOutput)
}

type ServiceIamPolicyPtrInput interface {
	pulumi.Input

	ToServiceIamPolicyPtrOutput() ServiceIamPolicyPtrOutput
	ToServiceIamPolicyPtrOutputWithContext(ctx context.Context) ServiceIamPolicyPtrOutput
}

type serviceIamPolicyPtrType ServiceIamPolicyArgs

func (*serviceIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceIamPolicy)(nil))
}

func (i *serviceIamPolicyPtrType) ToServiceIamPolicyPtrOutput() ServiceIamPolicyPtrOutput {
	return i.ToServiceIamPolicyPtrOutputWithContext(context.Background())
}

func (i *serviceIamPolicyPtrType) ToServiceIamPolicyPtrOutputWithContext(ctx context.Context) ServiceIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamPolicyPtrOutput)
}

// ServiceIamPolicyArrayInput is an input type that accepts ServiceIamPolicyArray and ServiceIamPolicyArrayOutput values.
// You can construct a concrete instance of `ServiceIamPolicyArrayInput` via:
//
//          ServiceIamPolicyArray{ ServiceIamPolicyArgs{...} }
type ServiceIamPolicyArrayInput interface {
	pulumi.Input

	ToServiceIamPolicyArrayOutput() ServiceIamPolicyArrayOutput
	ToServiceIamPolicyArrayOutputWithContext(context.Context) ServiceIamPolicyArrayOutput
}

type ServiceIamPolicyArray []ServiceIamPolicyInput

func (ServiceIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServiceIamPolicy)(nil))
}

func (i ServiceIamPolicyArray) ToServiceIamPolicyArrayOutput() ServiceIamPolicyArrayOutput {
	return i.ToServiceIamPolicyArrayOutputWithContext(context.Background())
}

func (i ServiceIamPolicyArray) ToServiceIamPolicyArrayOutputWithContext(ctx context.Context) ServiceIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamPolicyArrayOutput)
}

// ServiceIamPolicyMapInput is an input type that accepts ServiceIamPolicyMap and ServiceIamPolicyMapOutput values.
// You can construct a concrete instance of `ServiceIamPolicyMapInput` via:
//
//          ServiceIamPolicyMap{ "key": ServiceIamPolicyArgs{...} }
type ServiceIamPolicyMapInput interface {
	pulumi.Input

	ToServiceIamPolicyMapOutput() ServiceIamPolicyMapOutput
	ToServiceIamPolicyMapOutputWithContext(context.Context) ServiceIamPolicyMapOutput
}

type ServiceIamPolicyMap map[string]ServiceIamPolicyInput

func (ServiceIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServiceIamPolicy)(nil))
}

func (i ServiceIamPolicyMap) ToServiceIamPolicyMapOutput() ServiceIamPolicyMapOutput {
	return i.ToServiceIamPolicyMapOutputWithContext(context.Background())
}

func (i ServiceIamPolicyMap) ToServiceIamPolicyMapOutputWithContext(ctx context.Context) ServiceIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIamPolicyMapOutput)
}

type ServiceIamPolicyOutput struct {
	*pulumi.OutputState
}

func (ServiceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIamPolicy)(nil))
}

func (o ServiceIamPolicyOutput) ToServiceIamPolicyOutput() ServiceIamPolicyOutput {
	return o
}

func (o ServiceIamPolicyOutput) ToServiceIamPolicyOutputWithContext(ctx context.Context) ServiceIamPolicyOutput {
	return o
}

func (o ServiceIamPolicyOutput) ToServiceIamPolicyPtrOutput() ServiceIamPolicyPtrOutput {
	return o.ToServiceIamPolicyPtrOutputWithContext(context.Background())
}

func (o ServiceIamPolicyOutput) ToServiceIamPolicyPtrOutputWithContext(ctx context.Context) ServiceIamPolicyPtrOutput {
	return o.ApplyT(func(v ServiceIamPolicy) *ServiceIamPolicy {
		return &v
	}).(ServiceIamPolicyPtrOutput)
}

type ServiceIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (ServiceIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceIamPolicy)(nil))
}

func (o ServiceIamPolicyPtrOutput) ToServiceIamPolicyPtrOutput() ServiceIamPolicyPtrOutput {
	return o
}

func (o ServiceIamPolicyPtrOutput) ToServiceIamPolicyPtrOutputWithContext(ctx context.Context) ServiceIamPolicyPtrOutput {
	return o
}

type ServiceIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (ServiceIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceIamPolicy)(nil))
}

func (o ServiceIamPolicyArrayOutput) ToServiceIamPolicyArrayOutput() ServiceIamPolicyArrayOutput {
	return o
}

func (o ServiceIamPolicyArrayOutput) ToServiceIamPolicyArrayOutputWithContext(ctx context.Context) ServiceIamPolicyArrayOutput {
	return o
}

func (o ServiceIamPolicyArrayOutput) Index(i pulumi.IntInput) ServiceIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceIamPolicy {
		return vs[0].([]ServiceIamPolicy)[vs[1].(int)]
	}).(ServiceIamPolicyOutput)
}

type ServiceIamPolicyMapOutput struct{ *pulumi.OutputState }

func (ServiceIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceIamPolicy)(nil))
}

func (o ServiceIamPolicyMapOutput) ToServiceIamPolicyMapOutput() ServiceIamPolicyMapOutput {
	return o
}

func (o ServiceIamPolicyMapOutput) ToServiceIamPolicyMapOutputWithContext(ctx context.Context) ServiceIamPolicyMapOutput {
	return o
}

func (o ServiceIamPolicyMapOutput) MapIndex(k pulumi.StringInput) ServiceIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceIamPolicy {
		return vs[0].(map[string]ServiceIamPolicy)[vs[1].(string)]
	}).(ServiceIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceIamPolicyOutput{})
	pulumi.RegisterOutputType(ServiceIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(ServiceIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(ServiceIamPolicyMapOutput{})
}
