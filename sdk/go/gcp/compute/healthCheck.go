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
// To get more information about HealthCheck, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/healthChecks)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/health-checks)
type HealthCheck struct {
	pulumi.CustomResourceState

	// -
	// (Optional)
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec pulumi.IntPtrOutput `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// -
	// (Optional)
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrOutput `pulumi:"healthyThreshold"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	Http2HealthCheck HealthCheckHttp2HealthCheckPtrOutput `pulumi:"http2HealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpHealthCheck HealthCheckHttpHealthCheckPtrOutput `pulumi:"httpHealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck HealthCheckHttpsHealthCheckPtrOutput `pulumi:"httpsHealthCheck"`
	// -
	// (Required)
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
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	SslHealthCheck HealthCheckSslHealthCheckPtrOutput `pulumi:"sslHealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	TcpHealthCheck HealthCheckTcpHealthCheckPtrOutput `pulumi:"tcpHealthCheck"`
	// -
	// (Optional)
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrOutput `pulumi:"timeoutSec"`
	// The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
	Type pulumi.StringOutput `pulumi:"type"`
	// -
	// (Optional)
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrOutput `pulumi:"unhealthyThreshold"`
}

// NewHealthCheck registers a new resource with the given unique name, arguments, and options.
func NewHealthCheck(ctx *pulumi.Context,
	name string, args *HealthCheckArgs, opts ...pulumi.ResourceOption) (*HealthCheck, error) {
	if args == nil {
		args = &HealthCheckArgs{}
	}
	var resource HealthCheck
	err := ctx.RegisterResource("gcp:compute/healthCheck:HealthCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthCheck gets an existing HealthCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthCheckState, opts ...pulumi.ResourceOption) (*HealthCheck, error) {
	var resource HealthCheck
	err := ctx.ReadResource("gcp:compute/healthCheck:HealthCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthCheck resources.
type healthCheckState struct {
	// -
	// (Optional)
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *int `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	Http2HealthCheck *HealthCheckHttp2HealthCheck `pulumi:"http2HealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpHealthCheck *HealthCheckHttpHealthCheck `pulumi:"httpHealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck *HealthCheckHttpsHealthCheck `pulumi:"httpsHealthCheck"`
	// -
	// (Required)
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
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	SslHealthCheck *HealthCheckSslHealthCheck `pulumi:"sslHealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	TcpHealthCheck *HealthCheckTcpHealthCheck `pulumi:"tcpHealthCheck"`
	// -
	// (Optional)
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *int `pulumi:"timeoutSec"`
	// The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
	Type *string `pulumi:"type"`
	// -
	// (Optional)
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

type HealthCheckState struct {
	// -
	// (Optional)
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	Http2HealthCheck HealthCheckHttp2HealthCheckPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpHealthCheck HealthCheckHttpHealthCheckPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck HealthCheckHttpsHealthCheckPtrInput
	// -
	// (Required)
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
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	SslHealthCheck HealthCheckSslHealthCheckPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	TcpHealthCheck HealthCheckTcpHealthCheckPtrInput
	// -
	// (Optional)
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrInput
	// The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
	Type pulumi.StringPtrInput
	// -
	// (Optional)
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (HealthCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckState)(nil)).Elem()
}

type healthCheckArgs struct {
	// -
	// (Optional)
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *int `pulumi:"checkIntervalSec"`
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	Http2HealthCheck *HealthCheckHttp2HealthCheck `pulumi:"http2HealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpHealthCheck *HealthCheckHttpHealthCheck `pulumi:"httpHealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck *HealthCheckHttpsHealthCheck `pulumi:"httpsHealthCheck"`
	// -
	// (Required)
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
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	SslHealthCheck *HealthCheckSslHealthCheck `pulumi:"sslHealthCheck"`
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	TcpHealthCheck *HealthCheckTcpHealthCheck `pulumi:"tcpHealthCheck"`
	// -
	// (Optional)
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *int `pulumi:"timeoutSec"`
	// -
	// (Optional)
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

// The set of arguments for constructing a HealthCheck resource.
type HealthCheckArgs struct {
	// -
	// (Optional)
	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec pulumi.IntPtrInput
	// -
	// (Optional)
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	Http2HealthCheck HealthCheckHttp2HealthCheckPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpHealthCheck HealthCheckHttpHealthCheckPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	HttpsHealthCheck HealthCheckHttpsHealthCheckPtrInput
	// -
	// (Required)
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
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	SslHealthCheck HealthCheckSslHealthCheckPtrInput
	// -
	// (Optional)
	// A nested object resource  Structure is documented below.
	TcpHealthCheck HealthCheckTcpHealthCheckPtrInput
	// -
	// (Optional)
	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (HealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckArgs)(nil)).Elem()
}
