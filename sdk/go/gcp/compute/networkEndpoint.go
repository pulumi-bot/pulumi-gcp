// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Network endpoint represents a IP address and port combination that is
// part of a specific network endpoint group (NEG). NEGs are zonal
// collections of these endpoints for GCP resources within a
// single subnet. **NOTE**: Network endpoints cannot be created outside of a
// network endpoint group.
//
// To get more information about NetworkEndpoint, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)
//
// ## Example Usage
// ### Network Endpoint
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
// 		opt0 := "debian-9"
// 		opt1 := "debian-cloud"
// 		myImage, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Family:  &opt0,
// 			Project: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := compute.NewNetwork(ctx, "defaultNetwork", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSubnetwork, err := compute.NewSubnetwork(ctx, "defaultSubnetwork", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.0.0.1/16"),
// 			Region:      pulumi.String("us-central1"),
// 			Network:     defaultNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewInstance(ctx, "endpoint_instance", &compute.InstanceArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String(myImage.SelfLink),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Subnetwork: defaultSubnetwork.ID(),
// 					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
// 						nil,
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetworkEndpoint(ctx, "default_endpoint", &compute.NetworkEndpointArgs{
// 			NetworkEndpointGroup: pulumi.Any(google_compute_network_endpoint_group.Neg.Name),
// 			Instance:             endpoint_instance.Name,
// 			Port:                 pulumi.Any(google_compute_network_endpoint_group.Neg.Default_port),
// 			IpAddress: pulumi.String(endpoint_instance.NetworkInterfaces.ApplyT(func(networkInterfaces []compute.InstanceNetworkInterface) (string, error) {
// 				return networkInterfaces[0].NetworkIp, nil
// 			}).(pulumi.StringOutput)),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetworkEndpointGroup(ctx, "group", &compute.NetworkEndpointGroupArgs{
// 			Network:     defaultNetwork.ID(),
// 			Subnetwork:  defaultSubnetwork.ID(),
// 			DefaultPort: pulumi.Int(90),
// 			Zone:        pulumi.String("us-central1-a"),
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
// NetworkEndpoint can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/networkEndpoint:NetworkEndpoint default projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/networkEndpoint:NetworkEndpoint default {{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/networkEndpoint:NetworkEndpoint default {{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/networkEndpoint:NetworkEndpoint default {{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
// ```
type NetworkEndpoint struct {
	pulumi.CustomResourceState

	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup pulumi.StringOutput `pulumi:"networkEndpointGroup"`
	// Port number of network endpoint.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNetworkEndpoint registers a new resource with the given unique name, arguments, and options.
func NewNetworkEndpoint(ctx *pulumi.Context,
	name string, args *NetworkEndpointArgs, opts ...pulumi.ResourceOption) (*NetworkEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	if args.NetworkEndpointGroup == nil {
		return nil, errors.New("invalid value for required argument 'NetworkEndpointGroup'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	var resource NetworkEndpoint
	err := ctx.RegisterResource("gcp:compute/networkEndpoint:NetworkEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkEndpoint gets an existing NetworkEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkEndpointState, opts ...pulumi.ResourceOption) (*NetworkEndpoint, error) {
	var resource NetworkEndpoint
	err := ctx.ReadResource("gcp:compute/networkEndpoint:NetworkEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkEndpoint resources.
type networkEndpointState struct {
	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	Instance *string `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IpAddress *string `pulumi:"ipAddress"`
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup *string `pulumi:"networkEndpointGroup"`
	// Port number of network endpoint.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Zone where the containing network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

type NetworkEndpointState struct {
	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringPtrInput
	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IpAddress pulumi.StringPtrInput
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup pulumi.StringPtrInput
	// Port number of network endpoint.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (NetworkEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEndpointState)(nil)).Elem()
}

type networkEndpointArgs struct {
	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	Instance string `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IpAddress string `pulumi:"ipAddress"`
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup string `pulumi:"networkEndpointGroup"`
	// Port number of network endpoint.
	Port int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Zone where the containing network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a NetworkEndpoint resource.
type NetworkEndpointArgs struct {
	// The name for a specific VM instance that the IP address belongs to.
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringInput
	// IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IpAddress pulumi.StringInput
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup pulumi.StringInput
	// Port number of network endpoint.
	Port pulumi.IntInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (NetworkEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEndpointArgs)(nil)).Elem()
}

type NetworkEndpointInput interface {
	pulumi.Input

	ToNetworkEndpointOutput() NetworkEndpointOutput
	ToNetworkEndpointOutputWithContext(ctx context.Context) NetworkEndpointOutput
}

func (*NetworkEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkEndpoint)(nil))
}

func (i *NetworkEndpoint) ToNetworkEndpointOutput() NetworkEndpointOutput {
	return i.ToNetworkEndpointOutputWithContext(context.Background())
}

func (i *NetworkEndpoint) ToNetworkEndpointOutputWithContext(ctx context.Context) NetworkEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEndpointOutput)
}

func (i *NetworkEndpoint) ToNetworkEndpointPtrOutput() NetworkEndpointPtrOutput {
	return i.ToNetworkEndpointPtrOutputWithContext(context.Background())
}

func (i *NetworkEndpoint) ToNetworkEndpointPtrOutputWithContext(ctx context.Context) NetworkEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEndpointPtrOutput)
}

type NetworkEndpointPtrInput interface {
	pulumi.Input

	ToNetworkEndpointPtrOutput() NetworkEndpointPtrOutput
	ToNetworkEndpointPtrOutputWithContext(ctx context.Context) NetworkEndpointPtrOutput
}

type networkEndpointPtrType NetworkEndpointArgs

func (*networkEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkEndpoint)(nil))
}

func (i *networkEndpointPtrType) ToNetworkEndpointPtrOutput() NetworkEndpointPtrOutput {
	return i.ToNetworkEndpointPtrOutputWithContext(context.Background())
}

func (i *networkEndpointPtrType) ToNetworkEndpointPtrOutputWithContext(ctx context.Context) NetworkEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEndpointOutput).ToNetworkEndpointPtrOutput()
}

// NetworkEndpointArrayInput is an input type that accepts NetworkEndpointArray and NetworkEndpointArrayOutput values.
// You can construct a concrete instance of `NetworkEndpointArrayInput` via:
//
//          NetworkEndpointArray{ NetworkEndpointArgs{...} }
type NetworkEndpointArrayInput interface {
	pulumi.Input

	ToNetworkEndpointArrayOutput() NetworkEndpointArrayOutput
	ToNetworkEndpointArrayOutputWithContext(context.Context) NetworkEndpointArrayOutput
}

type NetworkEndpointArray []NetworkEndpointInput

func (NetworkEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkEndpoint)(nil))
}

func (i NetworkEndpointArray) ToNetworkEndpointArrayOutput() NetworkEndpointArrayOutput {
	return i.ToNetworkEndpointArrayOutputWithContext(context.Background())
}

func (i NetworkEndpointArray) ToNetworkEndpointArrayOutputWithContext(ctx context.Context) NetworkEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEndpointArrayOutput)
}

// NetworkEndpointMapInput is an input type that accepts NetworkEndpointMap and NetworkEndpointMapOutput values.
// You can construct a concrete instance of `NetworkEndpointMapInput` via:
//
//          NetworkEndpointMap{ "key": NetworkEndpointArgs{...} }
type NetworkEndpointMapInput interface {
	pulumi.Input

	ToNetworkEndpointMapOutput() NetworkEndpointMapOutput
	ToNetworkEndpointMapOutputWithContext(context.Context) NetworkEndpointMapOutput
}

type NetworkEndpointMap map[string]NetworkEndpointInput

func (NetworkEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkEndpoint)(nil))
}

func (i NetworkEndpointMap) ToNetworkEndpointMapOutput() NetworkEndpointMapOutput {
	return i.ToNetworkEndpointMapOutputWithContext(context.Background())
}

func (i NetworkEndpointMap) ToNetworkEndpointMapOutputWithContext(ctx context.Context) NetworkEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEndpointMapOutput)
}

type NetworkEndpointOutput struct {
	*pulumi.OutputState
}

func (NetworkEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkEndpoint)(nil))
}

func (o NetworkEndpointOutput) ToNetworkEndpointOutput() NetworkEndpointOutput {
	return o
}

func (o NetworkEndpointOutput) ToNetworkEndpointOutputWithContext(ctx context.Context) NetworkEndpointOutput {
	return o
}

func (o NetworkEndpointOutput) ToNetworkEndpointPtrOutput() NetworkEndpointPtrOutput {
	return o.ToNetworkEndpointPtrOutputWithContext(context.Background())
}

func (o NetworkEndpointOutput) ToNetworkEndpointPtrOutputWithContext(ctx context.Context) NetworkEndpointPtrOutput {
	return o.ApplyT(func(v NetworkEndpoint) *NetworkEndpoint {
		return &v
	}).(NetworkEndpointPtrOutput)
}

type NetworkEndpointPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkEndpoint)(nil))
}

func (o NetworkEndpointPtrOutput) ToNetworkEndpointPtrOutput() NetworkEndpointPtrOutput {
	return o
}

func (o NetworkEndpointPtrOutput) ToNetworkEndpointPtrOutputWithContext(ctx context.Context) NetworkEndpointPtrOutput {
	return o
}

type NetworkEndpointArrayOutput struct{ *pulumi.OutputState }

func (NetworkEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkEndpoint)(nil))
}

func (o NetworkEndpointArrayOutput) ToNetworkEndpointArrayOutput() NetworkEndpointArrayOutput {
	return o
}

func (o NetworkEndpointArrayOutput) ToNetworkEndpointArrayOutputWithContext(ctx context.Context) NetworkEndpointArrayOutput {
	return o
}

func (o NetworkEndpointArrayOutput) Index(i pulumi.IntInput) NetworkEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkEndpoint {
		return vs[0].([]NetworkEndpoint)[vs[1].(int)]
	}).(NetworkEndpointOutput)
}

type NetworkEndpointMapOutput struct{ *pulumi.OutputState }

func (NetworkEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkEndpoint)(nil))
}

func (o NetworkEndpointMapOutput) ToNetworkEndpointMapOutput() NetworkEndpointMapOutput {
	return o
}

func (o NetworkEndpointMapOutput) ToNetworkEndpointMapOutputWithContext(ctx context.Context) NetworkEndpointMapOutput {
	return o
}

func (o NetworkEndpointMapOutput) MapIndex(k pulumi.StringInput) NetworkEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkEndpoint {
		return vs[0].(map[string]NetworkEndpoint)[vs[1].(string)]
	}).(NetworkEndpointOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkEndpointOutput{})
	pulumi.RegisterOutputType(NetworkEndpointPtrOutput{})
	pulumi.RegisterOutputType(NetworkEndpointArrayOutput{})
	pulumi.RegisterOutputType(NetworkEndpointMapOutput{})
}
