// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_health_check.html.markdown.
type HealthCheck struct {
	pulumi.CustomResourceState

	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec pulumi.IntPtrOutput `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrOutput `pulumi:"healthyThreshold"`
	// A nested object resource
	Http2HealthCheck HealthCheckHttp2HealthCheckPtrOutput `pulumi:"http2HealthCheck"`
	// A nested object resource
	HttpHealthCheck HealthCheckHttpHealthCheckPtrOutput `pulumi:"httpHealthCheck"`
	// A nested object resource
	HttpsHealthCheck HealthCheckHttpsHealthCheckPtrOutput `pulumi:"httpsHealthCheck"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A nested object resource
	SslHealthCheck HealthCheckSslHealthCheckPtrOutput `pulumi:"sslHealthCheck"`
	// A nested object resource
	TcpHealthCheck HealthCheckTcpHealthCheckPtrOutput `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to
	// have greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrOutput `pulumi:"timeoutSec"`
	// The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
	Type pulumi.StringOutput `pulumi:"type"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
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
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec *int `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// A nested object resource
	Http2HealthCheck *HealthCheckHttp2HealthCheck `pulumi:"http2HealthCheck"`
	// A nested object resource
	HttpHealthCheck *HealthCheckHttpHealthCheck `pulumi:"httpHealthCheck"`
	// A nested object resource
	HttpsHealthCheck *HealthCheckHttpsHealthCheck `pulumi:"httpsHealthCheck"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A nested object resource
	SslHealthCheck *HealthCheckSslHealthCheck `pulumi:"sslHealthCheck"`
	// A nested object resource
	TcpHealthCheck *HealthCheckTcpHealthCheck `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to
	// have greater value than checkIntervalSec.
	TimeoutSec *int `pulumi:"timeoutSec"`
	// The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
	Type *string `pulumi:"type"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

type HealthCheckState struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrInput
	// A nested object resource
	Http2HealthCheck HealthCheckHttp2HealthCheckPtrInput
	// A nested object resource
	HttpHealthCheck HealthCheckHttpHealthCheckPtrInput
	// A nested object resource
	HttpsHealthCheck HealthCheckHttpsHealthCheckPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A nested object resource
	SslHealthCheck HealthCheckSslHealthCheckPtrInput
	// A nested object resource
	TcpHealthCheck HealthCheckTcpHealthCheckPtrInput
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to
	// have greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrInput
	// The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
	Type pulumi.StringPtrInput
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (HealthCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckState)(nil)).Elem()
}

type healthCheckArgs struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec *int `pulumi:"checkIntervalSec"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// A nested object resource
	Http2HealthCheck *HealthCheckHttp2HealthCheck `pulumi:"http2HealthCheck"`
	// A nested object resource
	HttpHealthCheck *HealthCheckHttpHealthCheck `pulumi:"httpHealthCheck"`
	// A nested object resource
	HttpsHealthCheck *HealthCheckHttpsHealthCheck `pulumi:"httpsHealthCheck"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A nested object resource
	SslHealthCheck *HealthCheckSslHealthCheck `pulumi:"sslHealthCheck"`
	// A nested object resource
	TcpHealthCheck *HealthCheckTcpHealthCheck `pulumi:"tcpHealthCheck"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to
	// have greater value than checkIntervalSec.
	TimeoutSec *int `pulumi:"timeoutSec"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}

// The set of arguments for constructing a HealthCheck resource.
type HealthCheckArgs struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold pulumi.IntPtrInput
	// A nested object resource
	Http2HealthCheck HealthCheckHttp2HealthCheckPtrInput
	// A nested object resource
	HttpHealthCheck HealthCheckHttpHealthCheckPtrInput
	// A nested object resource
	HttpsHealthCheck HealthCheckHttpsHealthCheckPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A nested object resource
	SslHealthCheck HealthCheckSslHealthCheckPtrInput
	// A nested object resource
	TcpHealthCheck HealthCheckTcpHealthCheckPtrInput
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to
	// have greater value than checkIntervalSec.
	TimeoutSec pulumi.IntPtrInput
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold pulumi.IntPtrInput
}

func (HealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckArgs)(nil)).Elem()
}

