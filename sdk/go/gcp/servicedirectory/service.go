// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An individual service. A service contains a name and optional metadata.
//
// To get more information about Service, see:
//
// * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services)
// * How-to Guides
//     * [Configuring a service](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_a_service)
//
// ## Example Usage
// ### Service Directory Service Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleNamespace, err := servicedirectory.NewNamespace(ctx, "exampleNamespace", &servicedirectory.NamespaceArgs{
// 			NamespaceId: pulumi.String("example-namespace"),
// 			Location:    pulumi.String("us-central1"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicedirectory.NewService(ctx, "exampleService", &servicedirectory.ServiceArgs{
// 			ServiceId: pulumi.String("example-service"),
// 			Namespace: exampleNamespace.ID(),
// 			Metadata: pulumi.StringMap{
// 				"stage":  pulumi.String("prod"),
// 				"region": pulumi.String("us-central1"),
// 			},
// 		}, pulumi.Provider(google_beta))
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
// Service can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:servicedirectory/service:Service default projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:servicedirectory/service:Service default {{project}}/{{location}}/{{namespace_id}}/{{service_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:servicedirectory/service:Service default {{location}}/{{namespace_id}}/{{service_id}}
// ```
type Service struct {
	pulumi.CustomResourceState

	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource name of the namespace this service will belong to.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	var resource Service
	err := ctx.RegisterResource("gcp:servicedirectory/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("gcp:servicedirectory/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata map[string]string `pulumi:"metadata"`
	// The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
	Name *string `pulumi:"name"`
	// The resource name of the namespace this service will belong to.
	Namespace *string `pulumi:"namespace"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId *string `pulumi:"serviceId"`
}

type ServiceState struct {
	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapInput
	// The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
	Name pulumi.StringPtrInput
	// The resource name of the namespace this service will belong to.
	Namespace pulumi.StringPtrInput
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata map[string]string `pulumi:"metadata"`
	// The resource name of the namespace this service will belong to.
	Namespace string `pulumi:"namespace"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Metadata for the service. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapInput
	// The resource name of the namespace this service will belong to.
	Namespace pulumi.StringInput
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	ServiceId pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct {
	*pulumi.OutputState
}

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Service)(nil))
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
