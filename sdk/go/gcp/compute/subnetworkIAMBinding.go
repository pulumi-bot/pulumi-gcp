// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine Subnetwork. Each of these resources serves a different use case:
//
// * `compute.SubnetworkIAMPolicy`: Authoritative. Sets the IAM policy for the subnetwork and replaces any existing policy already attached.
// * `compute.SubnetworkIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subnetwork are preserved.
// * `compute.SubnetworkIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subnetwork are preserved.
//
// > **Note:** `compute.SubnetworkIAMPolicy` **cannot** be used in conjunction with `compute.SubnetworkIAMBinding` and `compute.SubnetworkIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.SubnetworkIAMBinding` resources **can be** used in conjunction with `compute.SubnetworkIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_compute\_subnetwork\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/compute.networkUser",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetworkIAMPolicy(ctx, "policy", &compute.SubnetworkIAMPolicyArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/compute.networkUser",
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
// 		_, err = compute.NewSubnetworkIAMPolicy(ctx, "policy", &compute.SubnetworkIAMPolicyArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_compute\_subnetwork\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewSubnetworkIAMBinding(ctx, "binding", &compute.SubnetworkIAMBindingArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			Role:       pulumi.String("roles/compute.networkUser"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewSubnetworkIAMBinding(ctx, "binding", &compute.SubnetworkIAMBindingArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			Role:       pulumi.String("roles/compute.networkUser"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &compute.SubnetworkIAMBindingConditionArgs{
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
// ## google\_compute\_subnetwork\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewSubnetworkIAMMember(ctx, "member", &compute.SubnetworkIAMMemberArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			Role:       pulumi.String("roles/compute.networkUser"),
// 			Member:     pulumi.String("user:jane@example.com"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewSubnetworkIAMMember(ctx, "member", &compute.SubnetworkIAMMemberArgs{
// 			Project:    pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Project),
// 			Region:     pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Region),
// 			Subnetwork: pulumi.Any(google_compute_subnetwork.Network - with - private - secondary - ip - ranges.Name),
// 			Role:       pulumi.String("roles/compute.networkUser"),
// 			Member:     pulumi.String("user:jane@example.com"),
// 			Condition: &compute.SubnetworkIAMMemberConditionArgs{
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/regions/{{region}}/subnetworks/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine subnetwork IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMBinding:SubnetworkIAMBinding editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMBinding:SubnetworkIAMBinding editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMBinding:SubnetworkIAMBinding editor projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type SubnetworkIAMBinding struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition SubnetworkIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
}

// NewSubnetworkIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewSubnetworkIAMBinding(ctx *pulumi.Context,
	name string, args *SubnetworkIAMBindingArgs, opts ...pulumi.ResourceOption) (*SubnetworkIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Subnetwork == nil {
		return nil, errors.New("invalid value for required argument 'Subnetwork'")
	}
	var resource SubnetworkIAMBinding
	err := ctx.RegisterResource("gcp:compute/subnetworkIAMBinding:SubnetworkIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetworkIAMBinding gets an existing SubnetworkIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetworkIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetworkIAMBindingState, opts ...pulumi.ResourceOption) (*SubnetworkIAMBinding, error) {
	var resource SubnetworkIAMBinding
	err := ctx.ReadResource("gcp:compute/subnetworkIAMBinding:SubnetworkIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetworkIAMBinding resources.
type subnetworkIAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *SubnetworkIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork *string `pulumi:"subnetwork"`
}

type SubnetworkIAMBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition SubnetworkIAMBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringPtrInput
}

func (SubnetworkIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkIAMBindingState)(nil)).Elem()
}

type subnetworkIAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *SubnetworkIAMBindingCondition `pulumi:"condition"`
	Members   []string                       `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a SubnetworkIAMBinding resource.
type SubnetworkIAMBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition SubnetworkIAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringInput
}

func (SubnetworkIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkIAMBindingArgs)(nil)).Elem()
}

type SubnetworkIAMBindingInput interface {
	pulumi.Input

	ToSubnetworkIAMBindingOutput() SubnetworkIAMBindingOutput
	ToSubnetworkIAMBindingOutputWithContext(ctx context.Context) SubnetworkIAMBindingOutput
}

func (*SubnetworkIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetworkIAMBinding)(nil))
}

func (i *SubnetworkIAMBinding) ToSubnetworkIAMBindingOutput() SubnetworkIAMBindingOutput {
	return i.ToSubnetworkIAMBindingOutputWithContext(context.Background())
}

func (i *SubnetworkIAMBinding) ToSubnetworkIAMBindingOutputWithContext(ctx context.Context) SubnetworkIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkIAMBindingOutput)
}

type SubnetworkIAMBindingOutput struct {
	*pulumi.OutputState
}

func (SubnetworkIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetworkIAMBinding)(nil))
}

func (o SubnetworkIAMBindingOutput) ToSubnetworkIAMBindingOutput() SubnetworkIAMBindingOutput {
	return o
}

func (o SubnetworkIAMBindingOutput) ToSubnetworkIAMBindingOutputWithContext(ctx context.Context) SubnetworkIAMBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetworkIAMBindingOutput{})
}
