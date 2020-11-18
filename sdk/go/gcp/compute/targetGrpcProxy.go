// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a Target gRPC Proxy resource. A target gRPC proxy is a component
// of load balancers intended for load balancing gRPC traffic. Global forwarding
// rules reference a target gRPC proxy. The Target gRPC Proxy references
// a URL map which specifies how traffic routes to gRPC backend services.
//
// To get more information about TargetGrpcProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetGrpcProxies)
// * How-to Guides
//     * [Using Target gRPC Proxies](https://cloud.google.com/traffic-director/docs/proxyless-overview)
//
// ## Example Usage
type TargetGrpcProxy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to
	// patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest
	// fingerprint, make a get() request to retrieve the TargetGrpcProxy. A base64-encoded string.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL with id for the resource.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap pulumi.StringPtrOutput `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless pulumi.BoolPtrOutput `pulumi:"validateForProxyless"`
}

// NewTargetGrpcProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetGrpcProxy(ctx *pulumi.Context,
	name string, args *TargetGrpcProxyArgs, opts ...pulumi.ResourceOption) (*TargetGrpcProxy, error) {
	if args == nil {
		args = &TargetGrpcProxyArgs{}
	}

	var resource TargetGrpcProxy
	err := ctx.RegisterResource("gcp:compute/targetGrpcProxy:TargetGrpcProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGrpcProxy gets an existing TargetGrpcProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGrpcProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGrpcProxyState, opts ...pulumi.ResourceOption) (*TargetGrpcProxy, error) {
	var resource TargetGrpcProxy
	err := ctx.ReadResource("gcp:compute/targetGrpcProxy:TargetGrpcProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGrpcProxy resources.
type targetGrpcProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to
	// patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest
	// fingerprint, make a get() request to retrieve the TargetGrpcProxy. A base64-encoded string.
	Fingerprint *string `pulumi:"fingerprint"`
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Server-defined URL with id for the resource.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap *string `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless *bool `pulumi:"validateForProxyless"`
}

type TargetGrpcProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to
	// patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest
	// fingerprint, make a get() request to retrieve the TargetGrpcProxy. A base64-encoded string.
	Fingerprint pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Server-defined URL with id for the resource.
	SelfLinkWithId pulumi.StringPtrInput
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap pulumi.StringPtrInput
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless pulumi.BoolPtrInput
}

func (TargetGrpcProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGrpcProxyState)(nil)).Elem()
}

type targetGrpcProxyArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap *string `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless *bool `pulumi:"validateForProxyless"`
}

// The set of arguments for constructing a TargetGrpcProxy resource.
type TargetGrpcProxyArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap pulumi.StringPtrInput
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless pulumi.BoolPtrInput
}

func (TargetGrpcProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGrpcProxyArgs)(nil)).Elem()
}

type TargetGrpcProxyInput interface {
	pulumi.Input

	ToTargetGrpcProxyOutput() TargetGrpcProxyOutput
	ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput
}

func (TargetGrpcProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetGrpcProxy)(nil)).Elem()
}

func (i TargetGrpcProxy) ToTargetGrpcProxyOutput() TargetGrpcProxyOutput {
	return i.ToTargetGrpcProxyOutputWithContext(context.Background())
}

func (i TargetGrpcProxy) ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGrpcProxyOutput)
}

type TargetGrpcProxyOutput struct {
	*pulumi.OutputState
}

func (TargetGrpcProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetGrpcProxyOutput)(nil)).Elem()
}

func (o TargetGrpcProxyOutput) ToTargetGrpcProxyOutput() TargetGrpcProxyOutput {
	return o
}

func (o TargetGrpcProxyOutput) ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TargetGrpcProxyOutput{})
}
