// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// VPN tunnel resource.
// 
// 
// To get more information about VpnTunnel, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels)
// * How-to Guides
//     * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
//     * [Networks and Tunnel Routing](https://cloud.google.com/vpn/docs/concepts/choosing-networks-routing)
// 
// > **Warning:** All arguments including the shared secret will be stored in the raw
// state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
type VPNTunnel struct {
	s *pulumi.ResourceState
}

// NewVPNTunnel registers a new resource with the given unique name, arguments, and options.
func NewVPNTunnel(ctx *pulumi.Context,
	name string, args *VPNTunnelArgs, opts ...pulumi.ResourceOpt) (*VPNTunnel, error) {
	if args == nil || args.PeerIp == nil {
		return nil, errors.New("missing required argument 'PeerIp'")
	}
	if args == nil || args.SharedSecret == nil {
		return nil, errors.New("missing required argument 'SharedSecret'")
	}
	if args == nil || args.TargetVpnGateway == nil {
		return nil, errors.New("missing required argument 'TargetVpnGateway'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["ikeVersion"] = nil
		inputs["labels"] = nil
		inputs["localTrafficSelectors"] = nil
		inputs["name"] = nil
		inputs["peerIp"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["remoteTrafficSelectors"] = nil
		inputs["router"] = nil
		inputs["sharedSecret"] = nil
		inputs["targetVpnGateway"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["ikeVersion"] = args.IkeVersion
		inputs["labels"] = args.Labels
		inputs["localTrafficSelectors"] = args.LocalTrafficSelectors
		inputs["name"] = args.Name
		inputs["peerIp"] = args.PeerIp
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["remoteTrafficSelectors"] = args.RemoteTrafficSelectors
		inputs["router"] = args.Router
		inputs["sharedSecret"] = args.SharedSecret
		inputs["targetVpnGateway"] = args.TargetVpnGateway
	}
	inputs["creationTimestamp"] = nil
	inputs["detailedStatus"] = nil
	inputs["labelFingerprint"] = nil
	inputs["selfLink"] = nil
	inputs["sharedSecretHash"] = nil
	s, err := ctx.RegisterResource("gcp:compute/vPNTunnel:VPNTunnel", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VPNTunnel{s: s}, nil
}

// GetVPNTunnel gets an existing VPNTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPNTunnel(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VPNTunnelState, opts ...pulumi.ResourceOpt) (*VPNTunnel, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["detailedStatus"] = state.DetailedStatus
		inputs["ikeVersion"] = state.IkeVersion
		inputs["labelFingerprint"] = state.LabelFingerprint
		inputs["labels"] = state.Labels
		inputs["localTrafficSelectors"] = state.LocalTrafficSelectors
		inputs["name"] = state.Name
		inputs["peerIp"] = state.PeerIp
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["remoteTrafficSelectors"] = state.RemoteTrafficSelectors
		inputs["router"] = state.Router
		inputs["selfLink"] = state.SelfLink
		inputs["sharedSecret"] = state.SharedSecret
		inputs["sharedSecretHash"] = state.SharedSecretHash
		inputs["targetVpnGateway"] = state.TargetVpnGateway
	}
	s, err := ctx.ReadResource("gcp:compute/vPNTunnel:VPNTunnel", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VPNTunnel{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VPNTunnel) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VPNTunnel) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *VPNTunnel) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *VPNTunnel) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *VPNTunnel) DetailedStatus() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["detailedStatus"])
}

func (r *VPNTunnel) IkeVersion() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ikeVersion"])
}

func (r *VPNTunnel) LabelFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labelFingerprint"])
}

func (r *VPNTunnel) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *VPNTunnel) LocalTrafficSelectors() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["localTrafficSelectors"])
}

func (r *VPNTunnel) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *VPNTunnel) PeerIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peerIp"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *VPNTunnel) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *VPNTunnel) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

func (r *VPNTunnel) RemoteTrafficSelectors() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["remoteTrafficSelectors"])
}

func (r *VPNTunnel) Router() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["router"])
}

// The URI of the created resource.
func (r *VPNTunnel) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *VPNTunnel) SharedSecret() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedSecret"])
}

func (r *VPNTunnel) SharedSecretHash() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedSecretHash"])
}

func (r *VPNTunnel) TargetVpnGateway() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetVpnGateway"])
}

// Input properties used for looking up and filtering VPNTunnel resources.
type VPNTunnelState struct {
	CreationTimestamp interface{}
	Description interface{}
	DetailedStatus interface{}
	IkeVersion interface{}
	LabelFingerprint interface{}
	Labels interface{}
	LocalTrafficSelectors interface{}
	Name interface{}
	PeerIp interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	RemoteTrafficSelectors interface{}
	Router interface{}
	// The URI of the created resource.
	SelfLink interface{}
	SharedSecret interface{}
	SharedSecretHash interface{}
	TargetVpnGateway interface{}
}

// The set of arguments for constructing a VPNTunnel resource.
type VPNTunnelArgs struct {
	Description interface{}
	IkeVersion interface{}
	Labels interface{}
	LocalTrafficSelectors interface{}
	Name interface{}
	PeerIp interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	RemoteTrafficSelectors interface{}
	Router interface{}
	SharedSecret interface{}
	TargetVpnGateway interface{}
}
