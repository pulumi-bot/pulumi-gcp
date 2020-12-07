// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpcaccess

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Serverless VPC Access connector resource.
//
// To get more information about Connector, see:
//
// * [API documentation](https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors)
// * How-to Guides
//     * [Configuring Serverless VPC Access](https://cloud.google.com/vpc/docs/configure-serverless-vpc-access)
//
// ## Example Usage
// ### VPC Access Connector
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/vpcaccess"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vpcaccess.NewConnector(ctx, "connector", &vpcaccess.ConnectorArgs{
// 			IpCidrRange: pulumi.String("10.8.0.0/28"),
// 			Network:     pulumi.String("default"),
// 			Region:      pulumi.String("us-central1"),
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
// Connector can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:vpcaccess/connector:Connector default projects/{{project}}/locations/{{region}}/connectors/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:vpcaccess/connector:Connector default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:vpcaccess/connector:Connector default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:vpcaccess/connector:Connector default {{name}}
// ```
type Connector struct {
	pulumi.CustomResourceState

	// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// Maximum throughput of the connector in Mbps, must be greater than `minThroughput`. Default is 1000.
	MaxThroughput pulumi.IntPtrOutput `pulumi:"maxThroughput"`
	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput pulumi.IntPtrOutput `pulumi:"minThroughput"`
	// The name of the resource (Max 25 characters).
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of a VPC network.
	Network pulumi.StringOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where the VPC Access connector resides
	Region pulumi.StringOutput `pulumi:"region"`
	// The fully qualified name of this VPC connector
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// State of the VPC access connector.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewConnector registers a new resource with the given unique name, arguments, and options.
func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpCidrRange == nil {
		return nil, errors.New("invalid value for required argument 'IpCidrRange'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource Connector
	err := ctx.RegisterResource("gcp:vpcaccess/connector:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnector gets an existing Connector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("gcp:vpcaccess/connector:Connector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connector resources.
type connectorState struct {
	// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// Maximum throughput of the connector in Mbps, must be greater than `minThroughput`. Default is 1000.
	MaxThroughput *int `pulumi:"maxThroughput"`
	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput *int `pulumi:"minThroughput"`
	// The name of the resource (Max 25 characters).
	Name *string `pulumi:"name"`
	// Name of a VPC network.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the VPC Access connector resides
	Region *string `pulumi:"region"`
	// The fully qualified name of this VPC connector
	SelfLink *string `pulumi:"selfLink"`
	// State of the VPC access connector.
	State *string `pulumi:"state"`
}

type ConnectorState struct {
	// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
	IpCidrRange pulumi.StringPtrInput
	// Maximum throughput of the connector in Mbps, must be greater than `minThroughput`. Default is 1000.
	MaxThroughput pulumi.IntPtrInput
	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput pulumi.IntPtrInput
	// The name of the resource (Max 25 characters).
	Name pulumi.StringPtrInput
	// Name of a VPC network.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the VPC Access connector resides
	Region pulumi.StringPtrInput
	// The fully qualified name of this VPC connector
	SelfLink pulumi.StringPtrInput
	// State of the VPC access connector.
	State pulumi.StringPtrInput
}

func (ConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorState)(nil)).Elem()
}

type connectorArgs struct {
	// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
	IpCidrRange string `pulumi:"ipCidrRange"`
	// Maximum throughput of the connector in Mbps, must be greater than `minThroughput`. Default is 1000.
	MaxThroughput *int `pulumi:"maxThroughput"`
	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput *int `pulumi:"minThroughput"`
	// The name of the resource (Max 25 characters).
	Name *string `pulumi:"name"`
	// Name of a VPC network.
	Network string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the VPC Access connector resides
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a Connector resource.
type ConnectorArgs struct {
	// The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
	IpCidrRange pulumi.StringInput
	// Maximum throughput of the connector in Mbps, must be greater than `minThroughput`. Default is 1000.
	MaxThroughput pulumi.IntPtrInput
	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput pulumi.IntPtrInput
	// The name of the resource (Max 25 characters).
	Name pulumi.StringPtrInput
	// Name of a VPC network.
	Network pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the VPC Access connector resides
	Region pulumi.StringInput
}

func (ConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorArgs)(nil)).Elem()
}

type ConnectorInput interface {
	pulumi.Input

	ToConnectorOutput() ConnectorOutput
	ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput
}

func (Connector) ElementType() reflect.Type {
	return reflect.TypeOf((*Connector)(nil)).Elem()
}

func (i Connector) ToConnectorOutput() ConnectorOutput {
	return i.ToConnectorOutputWithContext(context.Background())
}

func (i Connector) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorOutput)
}

type ConnectorOutput struct {
	*pulumi.OutputState
}

func (ConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorOutput)(nil)).Elem()
}

func (o ConnectorOutput) ToConnectorOutput() ConnectorOutput {
	return o
}

func (o ConnectorOutput) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectorOutput{})
}
