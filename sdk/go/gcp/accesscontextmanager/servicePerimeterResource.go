// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows configuring a single GCP resource that should be inside of a service perimeter.
// This resource is intended to be used in cases where it is not possible to compile a full list
// of projects to include in a `accesscontextmanager.ServicePerimeter` resource,
// to enable them to be added separately.
//
// > **Note:** If this resource is used alongside a `accesscontextmanager.ServicePerimeter` resource,
// the service perimeter resource must have a `lifecycle` block with `ignoreChanges = [status[0].resources]` so
// they don't fight over which resources should be in the policy.
//
// To get more information about ServicePerimeterResource, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
// * How-to Guides
//     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the ACM API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
//
// ## Import
//
// ServicePerimeterResource can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource default {{perimeter_name}}/{{resource}}
// ```
type ServicePerimeterResource struct {
	pulumi.CustomResourceState

	// The name of the Service Perimeter to add this resource to.
	PerimeterName pulumi.StringOutput `pulumi:"perimeterName"`
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource pulumi.StringOutput `pulumi:"resource"`
}

// NewServicePerimeterResource registers a new resource with the given unique name, arguments, and options.
func NewServicePerimeterResource(ctx *pulumi.Context,
	name string, args *ServicePerimeterResourceArgs, opts ...pulumi.ResourceOption) (*ServicePerimeterResource, error) {
	if args == nil || args.PerimeterName == nil {
		return nil, errors.New("missing required argument 'PerimeterName'")
	}
	if args == nil || args.Resource == nil {
		return nil, errors.New("missing required argument 'Resource'")
	}
	if args == nil {
		args = &ServicePerimeterResourceArgs{}
	}
	var resource ServicePerimeterResource
	err := ctx.RegisterResource("gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePerimeterResource gets an existing ServicePerimeterResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePerimeterResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePerimeterResourceState, opts ...pulumi.ResourceOption) (*ServicePerimeterResource, error) {
	var resource ServicePerimeterResource
	err := ctx.ReadResource("gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePerimeterResource resources.
type servicePerimeterResourceState struct {
	// The name of the Service Perimeter to add this resource to.
	PerimeterName *string `pulumi:"perimeterName"`
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource *string `pulumi:"resource"`
}

type ServicePerimeterResourceState struct {
	// The name of the Service Perimeter to add this resource to.
	PerimeterName pulumi.StringPtrInput
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource pulumi.StringPtrInput
}

func (ServicePerimeterResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimeterResourceState)(nil)).Elem()
}

type servicePerimeterResourceArgs struct {
	// The name of the Service Perimeter to add this resource to.
	PerimeterName string `pulumi:"perimeterName"`
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource string `pulumi:"resource"`
}

// The set of arguments for constructing a ServicePerimeterResource resource.
type ServicePerimeterResourceArgs struct {
	// The name of the Service Perimeter to add this resource to.
	PerimeterName pulumi.StringInput
	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	Resource pulumi.StringInput
}

func (ServicePerimeterResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimeterResourceArgs)(nil)).Elem()
}
