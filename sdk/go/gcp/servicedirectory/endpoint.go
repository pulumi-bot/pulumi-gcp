// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An individual endpoint that provides a service.
//
// To get more information about Endpoint, see:
//
// * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services.endpoints)
// * How-to Guides
//     * [Configuring an endpoint](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_an_endpoint)
//
// ## Example Usage
// ### Service Directory Endpoint Basic
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
// 		exampleService, err := servicedirectory.NewService(ctx, "exampleService", &servicedirectory.ServiceArgs{
// 			ServiceId: pulumi.String("example-service"),
// 			Namespace: exampleNamespace.ID(),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicedirectory.NewEndpoint(ctx, "exampleEndpoint", &servicedirectory.EndpointArgs{
// 			EndpointId: pulumi.String("example-endpoint"),
// 			Service:    exampleService.ID(),
// 			Metadata: pulumi.StringMap{
// 				"stage":  pulumi.String("prod"),
// 				"region": pulumi.String("us-central1"),
// 			},
// 			Address: pulumi.String("1.2.3.4"),
// 			Port:    pulumi.Int(5353),
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
// Endpoint can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:servicedirectory/endpoint:Endpoint default projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}/endpoints/{{endpoint_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:servicedirectory/endpoint:Endpoint default {{project}}/{{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:servicedirectory/endpoint:Endpoint default {{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	// IPv4 or IPv6 address of the endpoint.
	Address pulumi.StringPtrOutput `pulumi:"address"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// Metadata for the endpoint. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 512 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port that the endpoint is running on, must be in the
	// range of [0, 65535]. If unspecified, the default is 0.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The resource name of the service that this endpoint provides.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("gcp:servicedirectory/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("gcp:servicedirectory/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// IPv4 or IPv6 address of the endpoint.
	Address *string `pulumi:"address"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	EndpointId *string `pulumi:"endpointId"`
	// Metadata for the endpoint. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 512 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata map[string]string `pulumi:"metadata"`
	// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
	Name *string `pulumi:"name"`
	// Port that the endpoint is running on, must be in the
	// range of [0, 65535]. If unspecified, the default is 0.
	Port *int `pulumi:"port"`
	// The resource name of the service that this endpoint provides.
	Service *string `pulumi:"service"`
}

type EndpointState struct {
	// IPv4 or IPv6 address of the endpoint.
	Address pulumi.StringPtrInput
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	EndpointId pulumi.StringPtrInput
	// Metadata for the endpoint. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 512 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapInput
	// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
	Name pulumi.StringPtrInput
	// Port that the endpoint is running on, must be in the
	// range of [0, 65535]. If unspecified, the default is 0.
	Port pulumi.IntPtrInput
	// The resource name of the service that this endpoint provides.
	Service pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// IPv4 or IPv6 address of the endpoint.
	Address *string `pulumi:"address"`
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	EndpointId string `pulumi:"endpointId"`
	// Metadata for the endpoint. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 512 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata map[string]string `pulumi:"metadata"`
	// Port that the endpoint is running on, must be in the
	// range of [0, 65535]. If unspecified, the default is 0.
	Port *int `pulumi:"port"`
	// The resource name of the service that this endpoint provides.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// IPv4 or IPv6 address of the endpoint.
	Address pulumi.StringPtrInput
	// The Resource ID must be 1-63 characters long, including digits,
	// lowercase letters or the hyphen character.
	EndpointId pulumi.StringInput
	// Metadata for the endpoint. This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 512 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapInput
	// Port that the endpoint is running on, must be in the
	// range of [0, 65535]. If unspecified, the default is 0.
	Port pulumi.IntPtrInput
	// The resource name of the service that this endpoint provides.
	Service pulumi.StringInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil))
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

type EndpointOutput struct {
	*pulumi.OutputState
}

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil))
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
}
