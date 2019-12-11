// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/monitoring_uptime_check_config.html.markdown.
type UptimeCheckConfig struct {
	s *pulumi.ResourceState
}

// NewUptimeCheckConfig registers a new resource with the given unique name, arguments, and options.
func NewUptimeCheckConfig(ctx *pulumi.Context,
	name string, args *UptimeCheckConfigArgs, opts ...pulumi.ResourceOpt) (*UptimeCheckConfig, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Timeout == nil {
		return nil, errors.New("missing required argument 'Timeout'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["contentMatchers"] = nil
		inputs["displayName"] = nil
		inputs["httpCheck"] = nil
		inputs["monitoredResource"] = nil
		inputs["period"] = nil
		inputs["project"] = nil
		inputs["resourceGroup"] = nil
		inputs["selectedRegions"] = nil
		inputs["tcpCheck"] = nil
		inputs["timeout"] = nil
	} else {
		inputs["contentMatchers"] = args.ContentMatchers
		inputs["displayName"] = args.DisplayName
		inputs["httpCheck"] = args.HttpCheck
		inputs["monitoredResource"] = args.MonitoredResource
		inputs["period"] = args.Period
		inputs["project"] = args.Project
		inputs["resourceGroup"] = args.ResourceGroup
		inputs["selectedRegions"] = args.SelectedRegions
		inputs["tcpCheck"] = args.TcpCheck
		inputs["timeout"] = args.Timeout
	}
	inputs["name"] = nil
	inputs["uptimeCheckId"] = nil
	s, err := ctx.RegisterResource("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UptimeCheckConfig{s: s}, nil
}

// GetUptimeCheckConfig gets an existing UptimeCheckConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUptimeCheckConfig(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UptimeCheckConfigState, opts ...pulumi.ResourceOpt) (*UptimeCheckConfig, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["contentMatchers"] = state.ContentMatchers
		inputs["displayName"] = state.DisplayName
		inputs["httpCheck"] = state.HttpCheck
		inputs["monitoredResource"] = state.MonitoredResource
		inputs["name"] = state.Name
		inputs["period"] = state.Period
		inputs["project"] = state.Project
		inputs["resourceGroup"] = state.ResourceGroup
		inputs["selectedRegions"] = state.SelectedRegions
		inputs["tcpCheck"] = state.TcpCheck
		inputs["timeout"] = state.Timeout
		inputs["uptimeCheckId"] = state.UptimeCheckId
	}
	s, err := ctx.ReadResource("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UptimeCheckConfig{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *UptimeCheckConfig) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *UptimeCheckConfig) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and
// other entries will be ignored. The server will look for an exact match of the string in the page response's content.
// This field is optional and should only be specified if a content match is required.
func (r *UptimeCheckConfig) ContentMatchers() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["contentMatchers"])
}

// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver
// Workspace in order to make it easier to identify; however, uniqueness is not enforced.
func (r *UptimeCheckConfig) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// Contains information needed to make an HTTP or HTTPS check.
func (r *UptimeCheckConfig) HttpCheck() pulumi.Output {
	return r.s.State["httpCheck"]
}

// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The
// following monitored resource types are supported for uptime checks: uptime_url gce_instance gae_app aws_ec2_instance
// aws_elb_load_balancer
func (r *UptimeCheckConfig) MonitoredResource() pulumi.Output {
	return r.s.State["monitoredResource"]
}

// A unique resource name for this UptimeCheckConfig. The format is
// projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
func (r *UptimeCheckConfig) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5
// minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
func (r *UptimeCheckConfig) Period() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["period"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *UptimeCheckConfig) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The group resource associated with the configuration.
func (r *UptimeCheckConfig) ResourceGroup() pulumi.Output {
	return r.s.State["resourceGroup"]
}

// The list of regions from which the check will be run. Some regions contain one location, and others contain more than
// one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error
// message is returned. Not specifying this field will result in uptime checks running from all regions.
func (r *UptimeCheckConfig) SelectedRegions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["selectedRegions"])
}

// Contains information needed to make a TCP check.
func (r *UptimeCheckConfig) TcpCheck() pulumi.Output {
	return r.s.State["tcpCheck"]
}

// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats
// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
func (r *UptimeCheckConfig) Timeout() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["timeout"])
}

// The id of the uptime check
func (r *UptimeCheckConfig) UptimeCheckId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["uptimeCheckId"])
}

// Input properties used for looking up and filtering UptimeCheckConfig resources.
type UptimeCheckConfigState struct {
	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported,
	// and other entries will be ignored. The server will look for an exact match of the string in the page response's
	// content. This field is optional and should only be specified if a content match is required.
	ContentMatchers interface{}
	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver
	// Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName interface{}
	// Contains information needed to make an HTTP or HTTPS check.
	HttpCheck interface{}
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The
	// following monitored resource types are supported for uptime checks: uptime_url gce_instance gae_app aws_ec2_instance
	// aws_elb_load_balancer
	MonitoredResource interface{}
	// A unique resource name for this UptimeCheckConfig. The format is
	// projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
	Name interface{}
	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5
	// minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The group resource associated with the configuration.
	ResourceGroup interface{}
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than
	// one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error
	// message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions interface{}
	// Contains information needed to make a TCP check.
	TcpCheck interface{}
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout interface{}
	// The id of the uptime check
	UptimeCheckId interface{}
}

// The set of arguments for constructing a UptimeCheckConfig resource.
type UptimeCheckConfigArgs struct {
	// The expected content on the page the check is run against. Currently, only the first entry in the list is supported,
	// and other entries will be ignored. The server will look for an exact match of the string in the page response's
	// content. This field is optional and should only be specified if a content match is required.
	ContentMatchers interface{}
	// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver
	// Workspace in order to make it easier to identify; however, uniqueness is not enforced.
	DisplayName interface{}
	// Contains information needed to make an HTTP or HTTPS check.
	HttpCheck interface{}
	// The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The
	// following monitored resource types are supported for uptime checks: uptime_url gce_instance gae_app aws_ec2_instance
	// aws_elb_load_balancer
	MonitoredResource interface{}
	// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5
	// minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
	Period interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The group resource associated with the configuration.
	ResourceGroup interface{}
	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than
	// one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error
	// message is returned. Not specifying this field will result in uptime checks running from all regions.
	SelectedRegions interface{}
	// Contains information needed to make a TCP check.
	TcpCheck interface{}
	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
	Timeout interface{}
}
