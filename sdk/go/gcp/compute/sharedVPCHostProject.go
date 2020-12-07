// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables the Google Compute Engine
// [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
// feature for a project, assigning it as a Shared VPC host project.
//
// For more information, see,
// [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
// where the Shared VPC feature is referred to by its former name "XPN".
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		host, err := compute.NewSharedVPCHostProject(ctx, "host", &compute.SharedVPCHostProjectArgs{
// 			Project: pulumi.String("host-project-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSharedVPCServiceProject(ctx, "service1", &compute.SharedVPCServiceProjectArgs{
// 			HostProject:    host.Project,
// 			ServiceProject: pulumi.String("service-project-id-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSharedVPCServiceProject(ctx, "service2", &compute.SharedVPCServiceProjectArgs{
// 			HostProject:    host.Project,
// 			ServiceProject: pulumi.String("service-project-id-2"),
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
// Google Compute Engine Shared VPC host project feature can be imported using the `project`, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/sharedVPCHostProject:SharedVPCHostProject host host-project-id
// ```
type SharedVPCHostProject struct {
	pulumi.CustomResourceState

	// The ID of the project that will serve as a Shared VPC host project
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewSharedVPCHostProject registers a new resource with the given unique name, arguments, and options.
func NewSharedVPCHostProject(ctx *pulumi.Context,
	name string, args *SharedVPCHostProjectArgs, opts ...pulumi.ResourceOption) (*SharedVPCHostProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource SharedVPCHostProject
	err := ctx.RegisterResource("gcp:compute/sharedVPCHostProject:SharedVPCHostProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSharedVPCHostProject gets an existing SharedVPCHostProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedVPCHostProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedVPCHostProjectState, opts ...pulumi.ResourceOption) (*SharedVPCHostProject, error) {
	var resource SharedVPCHostProject
	err := ctx.ReadResource("gcp:compute/sharedVPCHostProject:SharedVPCHostProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SharedVPCHostProject resources.
type sharedVPCHostProjectState struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project *string `pulumi:"project"`
}

type SharedVPCHostProjectState struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project pulumi.StringPtrInput
}

func (SharedVPCHostProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedVPCHostProjectState)(nil)).Elem()
}

type sharedVPCHostProjectArgs struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a SharedVPCHostProject resource.
type SharedVPCHostProjectArgs struct {
	// The ID of the project that will serve as a Shared VPC host project
	Project pulumi.StringInput
}

func (SharedVPCHostProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedVPCHostProjectArgs)(nil)).Elem()
}

type SharedVPCHostProjectInput interface {
	pulumi.Input

	ToSharedVPCHostProjectOutput() SharedVPCHostProjectOutput
	ToSharedVPCHostProjectOutputWithContext(ctx context.Context) SharedVPCHostProjectOutput
}

func (SharedVPCHostProject) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedVPCHostProject)(nil)).Elem()
}

func (i SharedVPCHostProject) ToSharedVPCHostProjectOutput() SharedVPCHostProjectOutput {
	return i.ToSharedVPCHostProjectOutputWithContext(context.Background())
}

func (i SharedVPCHostProject) ToSharedVPCHostProjectOutputWithContext(ctx context.Context) SharedVPCHostProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedVPCHostProjectOutput)
}

type SharedVPCHostProjectOutput struct {
	*pulumi.OutputState
}

func (SharedVPCHostProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedVPCHostProjectOutput)(nil)).Elem()
}

func (o SharedVPCHostProjectOutput) ToSharedVPCHostProjectOutput() SharedVPCHostProjectOutput {
	return o
}

func (o SharedVPCHostProjectOutput) ToSharedVPCHostProjectOutputWithContext(ctx context.Context) SharedVPCHostProjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SharedVPCHostProjectOutput{})
}
