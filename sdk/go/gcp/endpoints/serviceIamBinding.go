// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package endpoints

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for Service Management Service. Each of these resources serves a different use case:
// 
// * `endpoints.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `endpoints.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `endpoints.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
// 
// > **Note:** `endpoints.ServiceIamPolicy` **cannot** be used in conjunction with `endpoints.ServiceIamBinding` and `endpoints.ServiceIamMember` or they will fight over what your policy should be.
// 
// > **Note:** `endpoints.ServiceIamBinding` resources **can be** used in conjunction with `endpoints.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/endpoints_service_iam.html.markdown.
type ServiceIamBinding struct {
	pulumi.CustomResourceState

	Condition ServiceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
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
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ServiceIamBindingArgs{}
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
	Etag *string `pulumi:"etag"`
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
	Etag pulumi.StringPtrInput
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
	Members []string `pulumi:"members"`
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
	Members pulumi.StringArrayInput
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

