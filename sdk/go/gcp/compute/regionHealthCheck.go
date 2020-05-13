// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Health Checks determine whether instances are responsive and able to do work.
// They are an important part of a comprehensive load balancing configuration,
// as they enable monitoring instances behind load balancers.
//
// Health Checks poll instances at a specified interval. Instances that
// do not respond successfully to some number of probes in a row are marked
// as unhealthy. No new connections are sent to unhealthy instances,
// though existing connections will continue. The health check will
// continue to poll unhealthy instances. If an instance later responds
// successfully to some number of consecutive probes, it is marked
// healthy again and can receive new connections.
//
//
// To get more information about RegionHealthCheck, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionHealthChecks)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/health-checks)
type RegionHealthCheck struct {
	pulumi.CustomResourceState

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec pulumi.IntPtrOutput `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrOutput `pulumi:"healthyThreshold"`
	// A nested object resource  Structure is documented below.
	Http2HealthCheck RegionHealthCheckHttp2HealthCheckPtrOutput `pulumi:"http2HealthCheck"`
	// A nested object resource  Structure is documented below.
	HttpHealthCheck RegionHealthCheckHttpHealthCheckPtrOutput `pulumi:"httpHealthCheck"`
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck RegionHealthCheckHttpsHealthCheckPtrOutput `pulumi:"httpsHealthCheck"`
	// Configure logging on this health check.  Structure is documented below.
	LogConfig RegionHealthCheckLogConfigPtrOutput `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the created health check should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A nested object resource  Structure is documented below.
	SslHealthCheck RegionHealthCheckSslHealthCheckPtrOutput `pulumi:"sslHealthCheck"`
	// A nested object resource  Structure is documented below.
	TcpHealthCheck RegionHealthCheckTcpHealthCheckPtrOutput `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrOutput `pulumi:"timeoutSec"`
	// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
	Type pulumi.StringOutput `pulumi:"type"`
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrOutput `pulumi:"unhealthyThreshold"`
}

// NewRegionHealthCheck registers a new resource with the given unique name, arguments, and options.
func NewRegionHealthCheck(ctx *pulumi.Context,
	name string, args *RegionHealthCheckArgs, opts ...pulumi.ResourceOption) (*RegionHealthCheck, error) {
	if args == nil {
		args = &RegionHealthCheckArgs{}
	}
	var resource RegionHealthCheck
	err := ctx.RegisterResource("gcp:compute/regionHealthCheck:RegionHealthCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionHealthCheck gets an existing RegionHealthCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionHealthCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionHealthCheckState, opts ...pulumi.ResourceOption) (*RegionHealthCheck, error) {
	var resource RegionHealthCheck
	err := ctx.ReadResource("gcp:compute/regionHealthCheck:RegionHealthCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionHealthCheck resources.
type regionHealthCheckState struct {
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *int `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// A nested object resource  Structure is documented below.
	Http2HealthCheck *RegionHealthCheckHttp2HealthCheck `pulumi:"http2HealthCheck"`
	// A nested object resource  Structure is documented below.
	HttpHealthCheck *RegionHealthCheckHttpHealthCheck `pulumi:"httpHealthCheck"`
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck *RegionHealthCheckHttpsHealthCheck `pulumi:"httpsHealthCheck"`
	// Configure logging on this health check.  Structure is documented below.
	LogConfig *RegionHealthCheckLogConfig `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created health check should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A nested object resource  Structure is documented below.
	SslHealthCheck *RegionHealthCheckSslHealthCheck `pulumi:"sslHealthCheck"`
	// A nested object resource  Structure is documented below.
	TcpHealthCheck *RegionHealthCheckTcpHealthCheck `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *int `pulumi:"timeoutSec"`
	// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
	Type *string `pulumi:"type"`
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

type RegionHealthCheckState struct {
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrInput
	// A nested object resource  Structure is documented below.
	Http2HealthCheck RegionHealthCheckHttp2HealthCheckPtrInput
	// A nested object resource  Structure is documented below.
	HttpHealthCheck RegionHealthCheckHttpHealthCheckPtrInput
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck RegionHealthCheckHttpsHealthCheckPtrInput
	// Configure logging on this health check.  Structure is documented below.
	LogConfig RegionHealthCheckLogConfigPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created health check should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A nested object resource  Structure is documented below.
	SslHealthCheck RegionHealthCheckSslHealthCheckPtrInput
	// A nested object resource  Structure is documented below.
	TcpHealthCheck RegionHealthCheckTcpHealthCheckPtrInput
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrInput
	// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
	Type pulumi.StringPtrInput
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (RegionHealthCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionHealthCheckState)(nil)).Elem()
}

type regionHealthCheckArgs struct {
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *int `pulumi:"checkIntervalSec"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// A nested object resource  Structure is documented below.
	Http2HealthCheck *RegionHealthCheckHttp2HealthCheck `pulumi:"http2HealthCheck"`
	// A nested object resource  Structure is documented below.
	HttpHealthCheck *RegionHealthCheckHttpHealthCheck `pulumi:"httpHealthCheck"`
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck *RegionHealthCheckHttpsHealthCheck `pulumi:"httpsHealthCheck"`
	// Configure logging on this health check.  Structure is documented below.
	LogConfig *RegionHealthCheckLogConfig `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created health check should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// A nested object resource  Structure is documented below.
	SslHealthCheck *RegionHealthCheckSslHealthCheck `pulumi:"sslHealthCheck"`
	// A nested object resource  Structure is documented below.
	TcpHealthCheck *RegionHealthCheckTcpHealthCheck `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *int `pulumi:"timeoutSec"`
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

// The set of arguments for constructing a RegionHealthCheck resource.
type RegionHealthCheckArgs struct {
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrInput
	// A nested object resource  Structure is documented below.
	Http2HealthCheck RegionHealthCheckHttp2HealthCheckPtrInput
	// A nested object resource  Structure is documented below.
	HttpHealthCheck RegionHealthCheckHttpHealthCheckPtrInput
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck RegionHealthCheckHttpsHealthCheckPtrInput
	// Configure logging on this health check.  Structure is documented below.
	LogConfig RegionHealthCheckLogConfigPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created health check should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// A nested object resource  Structure is documented below.
	SslHealthCheck RegionHealthCheckSslHealthCheckPtrInput
	// A nested object resource  Structure is documented below.
	TcpHealthCheck RegionHealthCheckTcpHealthCheckPtrInput
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrInput
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (RegionHealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionHealthCheckArgs)(nil)).Elem()
}
