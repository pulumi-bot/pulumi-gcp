// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
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
type SubnetworkIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
}

// NewSubnetworkIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewSubnetworkIAMPolicy(ctx *pulumi.Context,
	name string, args *SubnetworkIAMPolicyArgs, opts ...pulumi.ResourceOption) (*SubnetworkIAMPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.Subnetwork == nil {
		return nil, errors.New("missing required argument 'Subnetwork'")
	}
	if args == nil {
		args = &SubnetworkIAMPolicyArgs{}
	}
	var resource SubnetworkIAMPolicy
	err := ctx.RegisterResource("gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetworkIAMPolicy gets an existing SubnetworkIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetworkIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetworkIAMPolicyState, opts ...pulumi.ResourceOption) (*SubnetworkIAMPolicy, error) {
	var resource SubnetworkIAMPolicy
	err := ctx.ReadResource("gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetworkIAMPolicy resources.
type subnetworkIAMPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork *string `pulumi:"subnetwork"`
}

type SubnetworkIAMPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringPtrInput
}

func (SubnetworkIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkIAMPolicyState)(nil)).Elem()
}

type subnetworkIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a SubnetworkIAMPolicy resource.
type SubnetworkIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringInput
}

func (SubnetworkIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkIAMPolicyArgs)(nil)).Elem()
}
