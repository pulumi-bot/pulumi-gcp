// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine Instance. Each of these resources serves a different use case:
//
// * `compute.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `compute.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `compute.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `compute.InstanceIAMPolicy` **cannot** be used in conjunction with `compute.InstanceIAMBinding` and `compute.InstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.InstanceIAMBinding` resources **can be** used in conjunction with `compute.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/zones/{{zone}}/instances/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine instance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/instanceIAMBinding:InstanceIAMBinding editor "projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/instanceIAMBinding:InstanceIAMBinding editor "projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/instanceIAMBinding:InstanceIAMBinding editor projects/{{project}}/zones/{{zone}}/instances/{{instance}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIAMBinding struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition InstanceIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringOutput      `pulumi:"instanceName"`
	Members      pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewInstanceIAMBinding(ctx *pulumi.Context,
	name string, args *InstanceIAMBindingArgs, opts ...pulumi.ResourceOption) (*InstanceIAMBinding, error) {
	if args == nil || args.InstanceName == nil {
		return nil, errors.New("missing required argument 'InstanceName'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &InstanceIAMBindingArgs{}
	}
	var resource InstanceIAMBinding
	err := ctx.RegisterResource("gcp:compute/instanceIAMBinding:InstanceIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIAMBinding gets an existing InstanceIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIAMBindingState, opts ...pulumi.ResourceOption) (*InstanceIAMBinding, error) {
	var resource InstanceIAMBinding
	err := ctx.ReadResource("gcp:compute/instanceIAMBinding:InstanceIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIAMBinding resources.
type instanceIAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *InstanceIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName *string  `pulumi:"instanceName"`
	Members      []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone *string `pulumi:"zone"`
}

type InstanceIAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition InstanceIAMBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringPtrInput
	Members      pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringPtrInput
}

func (InstanceIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMBindingState)(nil)).Elem()
}

type instanceIAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *InstanceIAMBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName string   `pulumi:"instanceName"`
	Members      []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceIAMBinding resource.
type InstanceIAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition InstanceIAMBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringInput
	Members      pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringPtrInput
}

func (InstanceIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMBindingArgs)(nil)).Elem()
}
