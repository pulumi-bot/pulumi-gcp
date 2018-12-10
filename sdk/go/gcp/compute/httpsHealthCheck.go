// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
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
// page_title: "Google: google_compute_https_health_check"
// sidebar_current: "docs-google-compute-https-health-check"
// description: |-
//   An HttpsHealthCheck resource.
// ---
// 
// # google\_compute\_https\_health\_check
// 
// An HttpsHealthCheck resource. This resource defines a template for how
// individual VMs should be checked for health, via HTTPS.
// 
// 
// ~> **Note:** google_compute_https_health_check is a legacy health check.
// The newer [google_compute_health_check](https://www.terraform.io/docs/providers/google/r/compute_health_check.html)
// should be preferred for all uses except
// [Network Load Balancers](https://cloud.google.com/compute/docs/load-balancing/network/)
// which still require the legacy version.
// 
// 
// To get more information about HttpsHealthCheck, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/latest/httpsHealthChecks)
// * How-to Guides
//     * [Adding Health Checks](https://cloud.google.com/compute/docs/load-balancing/health-checks#legacy_health_checks)
type HttpsHealthCheck struct {
	s *pulumi.ResourceState
}

// NewHttpsHealthCheck registers a new resource with the given unique name, arguments, and options.
func NewHttpsHealthCheck(ctx *pulumi.Context,
	name string, args *HttpsHealthCheckArgs, opts ...pulumi.ResourceOpt) (*HttpsHealthCheck, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["checkIntervalSec"] = nil
		inputs["description"] = nil
		inputs["healthyThreshold"] = nil
		inputs["host"] = nil
		inputs["name"] = nil
		inputs["port"] = nil
		inputs["project"] = nil
		inputs["requestPath"] = nil
		inputs["timeoutSec"] = nil
		inputs["unhealthyThreshold"] = nil
	} else {
		inputs["checkIntervalSec"] = args.CheckIntervalSec
		inputs["description"] = args.Description
		inputs["healthyThreshold"] = args.HealthyThreshold
		inputs["host"] = args.Host
		inputs["name"] = args.Name
		inputs["port"] = args.Port
		inputs["project"] = args.Project
		inputs["requestPath"] = args.RequestPath
		inputs["timeoutSec"] = args.TimeoutSec
		inputs["unhealthyThreshold"] = args.UnhealthyThreshold
	}
	inputs["creationTimestamp"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/httpsHealthCheck:HttpsHealthCheck", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HttpsHealthCheck{s: s}, nil
}

// GetHttpsHealthCheck gets an existing HttpsHealthCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpsHealthCheck(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HttpsHealthCheckState, opts ...pulumi.ResourceOpt) (*HttpsHealthCheck, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["checkIntervalSec"] = state.CheckIntervalSec
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["healthyThreshold"] = state.HealthyThreshold
		inputs["host"] = state.Host
		inputs["name"] = state.Name
		inputs["port"] = state.Port
		inputs["project"] = state.Project
		inputs["requestPath"] = state.RequestPath
		inputs["selfLink"] = state.SelfLink
		inputs["timeoutSec"] = state.TimeoutSec
		inputs["unhealthyThreshold"] = state.UnhealthyThreshold
	}
	s, err := ctx.ReadResource("gcp:compute/httpsHealthCheck:HttpsHealthCheck", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HttpsHealthCheck{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *HttpsHealthCheck) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *HttpsHealthCheck) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *HttpsHealthCheck) CheckIntervalSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["checkIntervalSec"])
}

func (r *HttpsHealthCheck) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *HttpsHealthCheck) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *HttpsHealthCheck) HealthyThreshold() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["healthyThreshold"])
}

func (r *HttpsHealthCheck) Host() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["host"])
}

func (r *HttpsHealthCheck) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *HttpsHealthCheck) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *HttpsHealthCheck) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *HttpsHealthCheck) RequestPath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["requestPath"])
}

// The URI of the created resource.
func (r *HttpsHealthCheck) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *HttpsHealthCheck) TimeoutSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["timeoutSec"])
}

func (r *HttpsHealthCheck) UnhealthyThreshold() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["unhealthyThreshold"])
}

// Input properties used for looking up and filtering HttpsHealthCheck resources.
type HttpsHealthCheckState struct {
	CheckIntervalSec interface{}
	CreationTimestamp interface{}
	Description interface{}
	HealthyThreshold interface{}
	Host interface{}
	Name interface{}
	Port interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	RequestPath interface{}
	// The URI of the created resource.
	SelfLink interface{}
	TimeoutSec interface{}
	UnhealthyThreshold interface{}
}

// The set of arguments for constructing a HttpsHealthCheck resource.
type HttpsHealthCheckArgs struct {
	CheckIntervalSec interface{}
	Description interface{}
	HealthyThreshold interface{}
	Host interface{}
	Name interface{}
	Port interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	RequestPath interface{}
	TimeoutSec interface{}
	UnhealthyThreshold interface{}
}
