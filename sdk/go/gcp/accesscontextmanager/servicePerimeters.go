// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Replace all existing Service Perimeters in an Access Policy with the Service Perimeters provided. This is done atomically.
// This is a bulk edit of all Service Perimeters and may override existing Service Perimeters created by `accesscontextmanager.ServicePerimeter`,
// thus causing a permadiff if used alongside `accesscontextmanager.ServicePerimeter` on the same parent.
//
// To get more information about ServicePerimeters, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
// * How-to Guides
//     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
//
// ## Example Usage
type ServicePerimeters struct {
	pulumi.CustomResourceState

	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters ServicePerimetersServicePerimeterArrayOutput `pulumi:"servicePerimeters"`
}

// NewServicePerimeters registers a new resource with the given unique name, arguments, and options.
func NewServicePerimeters(ctx *pulumi.Context,
	name string, args *ServicePerimetersArgs, opts ...pulumi.ResourceOption) (*ServicePerimeters, error) {
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil {
		args = &ServicePerimetersArgs{}
	}
	var resource ServicePerimeters
	err := ctx.RegisterResource("gcp:accesscontextmanager/servicePerimeters:ServicePerimeters", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePerimeters gets an existing ServicePerimeters resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePerimeters(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePerimetersState, opts ...pulumi.ResourceOption) (*ServicePerimeters, error) {
	var resource ServicePerimeters
	err := ctx.ReadResource("gcp:accesscontextmanager/servicePerimeters:ServicePerimeters", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePerimeters resources.
type servicePerimetersState struct {
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `pulumi:"parent"`
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters []ServicePerimetersServicePerimeter `pulumi:"servicePerimeters"`
}

type ServicePerimetersState struct {
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringPtrInput
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters ServicePerimetersServicePerimeterArrayInput
}

func (ServicePerimetersState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimetersState)(nil)).Elem()
}

type servicePerimetersArgs struct {
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent string `pulumi:"parent"`
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters []ServicePerimetersServicePerimeter `pulumi:"servicePerimeters"`
}

// The set of arguments for constructing a ServicePerimeters resource.
type ServicePerimetersArgs struct {
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent pulumi.StringInput
	// The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
	// Structure is documented below.
	ServicePerimeters ServicePerimetersServicePerimeterArrayInput
}

func (ServicePerimetersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimetersArgs)(nil)).Elem()
}

type ServicePerimetersInput interface {
	pulumi.Input

	ToServicePerimetersOutput() ServicePerimetersOutput
	ToServicePerimetersOutputWithContext(ctx context.Context) ServicePerimetersOutput
}

func (ServicePerimeters) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePerimeters)(nil)).Elem()
}

func (i ServicePerimeters) ToServicePerimetersOutput() ServicePerimetersOutput {
	return i.ToServicePerimetersOutputWithContext(context.Background())
}

func (i ServicePerimeters) ToServicePerimetersOutputWithContext(ctx context.Context) ServicePerimetersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimetersOutput)
}

type ServicePerimetersOutput struct {
	*pulumi.OutputState
}

func (ServicePerimetersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePerimetersOutput)(nil)).Elem()
}

func (o ServicePerimetersOutput) ToServicePerimetersOutput() ServicePerimetersOutput {
	return o
}

func (o ServicePerimetersOutput) ToServicePerimetersOutputWithContext(ctx context.Context) ServicePerimetersOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServicePerimetersOutput{})
}
