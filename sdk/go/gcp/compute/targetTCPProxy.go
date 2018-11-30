// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// #     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
// #
// # ----------------------------------------------------------------------------
// #
// #     This file is automatically generated by Magic Modules and manual
// #     changes will be clobbered when the file is regenerated.
// #
// #     Please read more about how to change this file in
// #     .github/CONTRIBUTING.md.
// #
// # ----------------------------------------------------------------------------
// layout: "google"
// page_title: "Google: google_compute_target_tcp_proxy"
// sidebar_current: "docs-google-compute-target-tcp-proxy"
// description: |-
//   Represents a TargetTcpProxy resource, which is used by one or more
//   global forwarding rule to route incoming TCP requests to a Backend
//   service.
// ---
// 
// # google\_compute\_target\_tcp\_proxy
// 
// Represents a TargetTcpProxy resource, which is used by one or more
// global forwarding rule to route incoming TCP requests to a Backend
// service.
// 
// 
// To get more information about TargetTcpProxy, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/latest/targetTcpProxies)
// * How-to Guides
//     * [Setting Up TCP proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/tcp-proxy)
type TargetTCPProxy struct {
	s *pulumi.ResourceState
}

// NewTargetTCPProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetTCPProxy(ctx *pulumi.Context,
	name string, args *TargetTCPProxyArgs, opts ...pulumi.ResourceOpt) (*TargetTCPProxy, error) {
	if args == nil || args.BackendService == nil {
		return nil, errors.New("missing required argument 'BackendService'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backendService"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["proxyHeader"] = nil
	} else {
		inputs["backendService"] = args.BackendService
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["proxyHeader"] = args.ProxyHeader
	}
	inputs["creationTimestamp"] = nil
	inputs["proxyId"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/targetTCPProxy:TargetTCPProxy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetTCPProxy{s: s}, nil
}

// GetTargetTCPProxy gets an existing TargetTCPProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetTCPProxy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TargetTCPProxyState, opts ...pulumi.ResourceOpt) (*TargetTCPProxy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backendService"] = state.BackendService
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["proxyHeader"] = state.ProxyHeader
		inputs["proxyId"] = state.ProxyId
		inputs["selfLink"] = state.SelfLink
	}
	s, err := ctx.ReadResource("gcp:compute/targetTCPProxy:TargetTCPProxy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetTCPProxy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TargetTCPProxy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TargetTCPProxy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *TargetTCPProxy) BackendService() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backendService"])
}

func (r *TargetTCPProxy) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *TargetTCPProxy) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *TargetTCPProxy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *TargetTCPProxy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *TargetTCPProxy) ProxyHeader() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["proxyHeader"])
}

func (r *TargetTCPProxy) ProxyId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["proxyId"])
}

// The URI of the created resource.
func (r *TargetTCPProxy) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// Input properties used for looking up and filtering TargetTCPProxy resources.
type TargetTCPProxyState struct {
	BackendService interface{}
	CreationTimestamp interface{}
	Description interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	ProxyHeader interface{}
	ProxyId interface{}
	// The URI of the created resource.
	SelfLink interface{}
}

// The set of arguments for constructing a TargetTCPProxy resource.
type TargetTCPProxyArgs struct {
	BackendService interface{}
	Description interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	ProxyHeader interface{}
}
