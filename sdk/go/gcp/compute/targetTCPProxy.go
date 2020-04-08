// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a TargetTcpProxy resource, which is used by one or more
// global forwarding rule to route incoming TCP requests to a Backend
// service.
//
//
// To get more information about TargetTcpProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetTcpProxies)
// * How-to Guides
//     * [Setting Up TCP proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/tcp-proxy)
type TargetTCPProxy struct {
	pulumi.CustomResourceState

	// A reference to the BackendService resource.
	BackendService pulumi.StringOutput `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is
	// NONE.
	ProxyHeader pulumi.StringPtrOutput `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId pulumi.IntOutput `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewTargetTCPProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetTCPProxy(ctx *pulumi.Context,
	name string, args *TargetTCPProxyArgs, opts ...pulumi.ResourceOption) (*TargetTCPProxy, error) {
	if args == nil || args.BackendService == nil {
		return nil, errors.New("missing required argument 'BackendService'")
	}
	if args == nil {
		args = &TargetTCPProxyArgs{}
	}
	var resource TargetTCPProxy
	err := ctx.RegisterResource("gcp:compute/targetTCPProxy:TargetTCPProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetTCPProxy gets an existing TargetTCPProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetTCPProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetTCPProxyState, opts ...pulumi.ResourceOption) (*TargetTCPProxy, error) {
	var resource TargetTCPProxy
	err := ctx.ReadResource("gcp:compute/targetTCPProxy:TargetTCPProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetTCPProxy resources.
type targetTCPProxyState struct {
	// A reference to the BackendService resource.
	BackendService *string `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is
	// NONE.
	ProxyHeader *string `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId *int `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type TargetTCPProxyState struct {
	// A reference to the BackendService resource.
	BackendService pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is
	// NONE.
	ProxyHeader pulumi.StringPtrInput
	// The unique identifier for the resource.
	ProxyId pulumi.IntPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (TargetTCPProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetTCPProxyState)(nil)).Elem()
}

type targetTCPProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService string `pulumi:"backendService"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is
	// NONE.
	ProxyHeader *string `pulumi:"proxyHeader"`
}

// The set of arguments for constructing a TargetTCPProxy resource.
type TargetTCPProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService pulumi.StringInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is
	// NONE.
	ProxyHeader pulumi.StringPtrInput
}

func (TargetTCPProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetTCPProxyArgs)(nil)).Elem()
}
