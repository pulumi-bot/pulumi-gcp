// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_network_endpoint.html.markdown.
type NetworkEndpoint struct {
	pulumi.CustomResourceState

	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
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
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil || args.IpAddress == nil {
		return nil, errors.New("missing required argument 'IpAddress'")
	}
	if args == nil || args.NetworkEndpointGroup == nil {
		return nil, errors.New("missing required argument 'NetworkEndpointGroup'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	if args == nil {
		args = &NetworkEndpointArgs{}
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
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance *string `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
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
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringPtrInput
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
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
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance string `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
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
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringInput
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
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

