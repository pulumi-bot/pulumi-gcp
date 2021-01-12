// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a TargetHttpsProxy resource, which is used by one or more
// global forwarding rule to route incoming HTTPS requests to a URL map.
//
// To get more information about TargetHttpsProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpsProxies)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
//
// ## Example Usage
//
// ## Import
//
// TargetHttpsProxy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/targetHttpsProxy:TargetHttpsProxy default projects/{{project}}/global/targetHttpsProxies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetHttpsProxy:TargetHttpsProxy default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetHttpsProxy:TargetHttpsProxy default {{name}}
// ```
type TargetHttpsProxy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The unique identifier for the resource.
	ProxyId pulumi.IntOutput `pulumi:"proxyId"`
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE.
	// Default value is `NONE`.
	// Possible values are `NONE`, `ENABLE`, and `DISABLE`.
	QuicOverride pulumi.StringPtrOutput `pulumi:"quicOverride"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates pulumi.StringArrayOutput `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrOutput `pulumi:"sslPolicy"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewTargetHttpsProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpsProxy(ctx *pulumi.Context,
	name string, args *TargetHttpsProxyArgs, opts ...pulumi.ResourceOption) (*TargetHttpsProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SslCertificates == nil {
		return nil, errors.New("invalid value for required argument 'SslCertificates'")
	}
	if args.UrlMap == nil {
		return nil, errors.New("invalid value for required argument 'UrlMap'")
	}
	var resource TargetHttpsProxy
	err := ctx.RegisterResource("gcp:compute/targetHttpsProxy:TargetHttpsProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetHttpsProxy gets an existing TargetHttpsProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetHttpsProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetHttpsProxyState, opts ...pulumi.ResourceOption) (*TargetHttpsProxy, error) {
	var resource TargetHttpsProxy
	err := ctx.ReadResource("gcp:compute/targetHttpsProxy:TargetHttpsProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetHttpsProxy resources.
type targetHttpsProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The unique identifier for the resource.
	ProxyId *int `pulumi:"proxyId"`
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE.
	// Default value is `NONE`.
	// Possible values are `NONE`, `ENABLE`, and `DISABLE`.
	QuicOverride *string `pulumi:"quicOverride"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap *string `pulumi:"urlMap"`
}

type TargetHttpsProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The unique identifier for the resource.
	ProxyId pulumi.IntPtrInput
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE.
	// Default value is `NONE`.
	// Possible values are `NONE`, `ENABLE`, and `DISABLE`.
	QuicOverride pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates pulumi.StringArrayInput
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringPtrInput
}

func (TargetHttpsProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpsProxyState)(nil)).Elem()
}

type targetHttpsProxyArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE.
	// Default value is `NONE`.
	// Possible values are `NONE`, `ENABLE`, and `DISABLE`.
	QuicOverride *string `pulumi:"quicOverride"`
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap string `pulumi:"urlMap"`
}

// The set of arguments for constructing a TargetHttpsProxy resource.
type TargetHttpsProxyArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE.
	// Default value is `NONE`.
	// Possible values are `NONE`, `ENABLE`, and `DISABLE`.
	QuicOverride pulumi.StringPtrInput
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates pulumi.StringArrayInput
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringInput
}

func (TargetHttpsProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpsProxyArgs)(nil)).Elem()
}

type TargetHttpsProxyInput interface {
	pulumi.Input

	ToTargetHttpsProxyOutput() TargetHttpsProxyOutput
	ToTargetHttpsProxyOutputWithContext(ctx context.Context) TargetHttpsProxyOutput
}

func (*TargetHttpsProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetHttpsProxy)(nil))
}

func (i *TargetHttpsProxy) ToTargetHttpsProxyOutput() TargetHttpsProxyOutput {
	return i.ToTargetHttpsProxyOutputWithContext(context.Background())
}

func (i *TargetHttpsProxy) ToTargetHttpsProxyOutputWithContext(ctx context.Context) TargetHttpsProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpsProxyOutput)
}

func (i *TargetHttpsProxy) ToTargetHttpsProxyPtrOutput() TargetHttpsProxyPtrOutput {
	return i.ToTargetHttpsProxyPtrOutputWithContext(context.Background())
}

func (i *TargetHttpsProxy) ToTargetHttpsProxyPtrOutputWithContext(ctx context.Context) TargetHttpsProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpsProxyPtrOutput)
}

type TargetHttpsProxyPtrInput interface {
	pulumi.Input

	ToTargetHttpsProxyPtrOutput() TargetHttpsProxyPtrOutput
	ToTargetHttpsProxyPtrOutputWithContext(ctx context.Context) TargetHttpsProxyPtrOutput
}

type targetHttpsProxyPtrType TargetHttpsProxyArgs

func (*targetHttpsProxyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetHttpsProxy)(nil))
}

func (i *targetHttpsProxyPtrType) ToTargetHttpsProxyPtrOutput() TargetHttpsProxyPtrOutput {
	return i.ToTargetHttpsProxyPtrOutputWithContext(context.Background())
}

func (i *targetHttpsProxyPtrType) ToTargetHttpsProxyPtrOutputWithContext(ctx context.Context) TargetHttpsProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpsProxyOutput).ToTargetHttpsProxyPtrOutput()
}

// TargetHttpsProxyArrayInput is an input type that accepts TargetHttpsProxyArray and TargetHttpsProxyArrayOutput values.
// You can construct a concrete instance of `TargetHttpsProxyArrayInput` via:
//
//          TargetHttpsProxyArray{ TargetHttpsProxyArgs{...} }
type TargetHttpsProxyArrayInput interface {
	pulumi.Input

	ToTargetHttpsProxyArrayOutput() TargetHttpsProxyArrayOutput
	ToTargetHttpsProxyArrayOutputWithContext(context.Context) TargetHttpsProxyArrayOutput
}

type TargetHttpsProxyArray []TargetHttpsProxyInput

func (TargetHttpsProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetHttpsProxy)(nil))
}

func (i TargetHttpsProxyArray) ToTargetHttpsProxyArrayOutput() TargetHttpsProxyArrayOutput {
	return i.ToTargetHttpsProxyArrayOutputWithContext(context.Background())
}

func (i TargetHttpsProxyArray) ToTargetHttpsProxyArrayOutputWithContext(ctx context.Context) TargetHttpsProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpsProxyArrayOutput)
}

// TargetHttpsProxyMapInput is an input type that accepts TargetHttpsProxyMap and TargetHttpsProxyMapOutput values.
// You can construct a concrete instance of `TargetHttpsProxyMapInput` via:
//
//          TargetHttpsProxyMap{ "key": TargetHttpsProxyArgs{...} }
type TargetHttpsProxyMapInput interface {
	pulumi.Input

	ToTargetHttpsProxyMapOutput() TargetHttpsProxyMapOutput
	ToTargetHttpsProxyMapOutputWithContext(context.Context) TargetHttpsProxyMapOutput
}

type TargetHttpsProxyMap map[string]TargetHttpsProxyInput

func (TargetHttpsProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TargetHttpsProxy)(nil))
}

func (i TargetHttpsProxyMap) ToTargetHttpsProxyMapOutput() TargetHttpsProxyMapOutput {
	return i.ToTargetHttpsProxyMapOutputWithContext(context.Background())
}

func (i TargetHttpsProxyMap) ToTargetHttpsProxyMapOutputWithContext(ctx context.Context) TargetHttpsProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpsProxyMapOutput)
}

type TargetHttpsProxyOutput struct {
	*pulumi.OutputState
}

func (TargetHttpsProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetHttpsProxy)(nil))
}

func (o TargetHttpsProxyOutput) ToTargetHttpsProxyOutput() TargetHttpsProxyOutput {
	return o
}

func (o TargetHttpsProxyOutput) ToTargetHttpsProxyOutputWithContext(ctx context.Context) TargetHttpsProxyOutput {
	return o
}

func (o TargetHttpsProxyOutput) ToTargetHttpsProxyPtrOutput() TargetHttpsProxyPtrOutput {
	return o.ToTargetHttpsProxyPtrOutputWithContext(context.Background())
}

func (o TargetHttpsProxyOutput) ToTargetHttpsProxyPtrOutputWithContext(ctx context.Context) TargetHttpsProxyPtrOutput {
	return o.ApplyT(func(v TargetHttpsProxy) *TargetHttpsProxy {
		return &v
	}).(TargetHttpsProxyPtrOutput)
}

type TargetHttpsProxyPtrOutput struct {
	*pulumi.OutputState
}

func (TargetHttpsProxyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetHttpsProxy)(nil))
}

func (o TargetHttpsProxyPtrOutput) ToTargetHttpsProxyPtrOutput() TargetHttpsProxyPtrOutput {
	return o
}

func (o TargetHttpsProxyPtrOutput) ToTargetHttpsProxyPtrOutputWithContext(ctx context.Context) TargetHttpsProxyPtrOutput {
	return o
}

type TargetHttpsProxyArrayOutput struct{ *pulumi.OutputState }

func (TargetHttpsProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetHttpsProxy)(nil))
}

func (o TargetHttpsProxyArrayOutput) ToTargetHttpsProxyArrayOutput() TargetHttpsProxyArrayOutput {
	return o
}

func (o TargetHttpsProxyArrayOutput) ToTargetHttpsProxyArrayOutputWithContext(ctx context.Context) TargetHttpsProxyArrayOutput {
	return o
}

func (o TargetHttpsProxyArrayOutput) Index(i pulumi.IntInput) TargetHttpsProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetHttpsProxy {
		return vs[0].([]TargetHttpsProxy)[vs[1].(int)]
	}).(TargetHttpsProxyOutput)
}

type TargetHttpsProxyMapOutput struct{ *pulumi.OutputState }

func (TargetHttpsProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TargetHttpsProxy)(nil))
}

func (o TargetHttpsProxyMapOutput) ToTargetHttpsProxyMapOutput() TargetHttpsProxyMapOutput {
	return o
}

func (o TargetHttpsProxyMapOutput) ToTargetHttpsProxyMapOutputWithContext(ctx context.Context) TargetHttpsProxyMapOutput {
	return o
}

func (o TargetHttpsProxyMapOutput) MapIndex(k pulumi.StringInput) TargetHttpsProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TargetHttpsProxy {
		return vs[0].(map[string]TargetHttpsProxy)[vs[1].(string)]
	}).(TargetHttpsProxyOutput)
}

func init() {
	pulumi.RegisterOutputType(TargetHttpsProxyOutput{})
	pulumi.RegisterOutputType(TargetHttpsProxyPtrOutput{})
	pulumi.RegisterOutputType(TargetHttpsProxyArrayOutput{})
	pulumi.RegisterOutputType(TargetHttpsProxyMapOutput{})
}
