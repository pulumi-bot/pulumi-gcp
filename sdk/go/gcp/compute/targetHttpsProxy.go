// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a TargetHttpsProxy resource, which is used by one or more
// global forwarding rule to route incoming HTTPS requests to a URL map.
//
//
// To get more information about TargetHttpsProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpsProxies)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
type TargetHttpsProxy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// -
	// (Required)
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
	// -
	// (Optional)
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE. Not specifying this field is equivalent to
	// specifying NONE.
	QuicOverride pulumi.StringPtrOutput `pulumi:"quicOverride"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// -
	// (Required)
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates pulumi.StringArrayOutput `pulumi:"sslCertificates"`
	// -
	// (Optional)
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrOutput `pulumi:"sslPolicy"`
	// -
	// (Required)
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewTargetHttpsProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpsProxy(ctx *pulumi.Context,
	name string, args *TargetHttpsProxyArgs, opts ...pulumi.ResourceOption) (*TargetHttpsProxy, error) {
	if args == nil || args.SslCertificates == nil {
		return nil, errors.New("missing required argument 'SslCertificates'")
	}
	if args == nil || args.UrlMap == nil {
		return nil, errors.New("missing required argument 'UrlMap'")
	}
	if args == nil {
		args = &TargetHttpsProxyArgs{}
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
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// -
	// (Required)
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
	// -
	// (Optional)
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE. Not specifying this field is equivalent to
	// specifying NONE.
	QuicOverride *string `pulumi:"quicOverride"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// -
	// (Required)
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// -
	// (Optional)
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
	// -
	// (Required)
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap *string `pulumi:"urlMap"`
}

type TargetHttpsProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// -
	// (Required)
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
	// -
	// (Optional)
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE. Not specifying this field is equivalent to
	// specifying NONE.
	QuicOverride pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// -
	// (Required)
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates pulumi.StringArrayInput
	// -
	// (Optional)
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
	// -
	// (Required)
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringPtrInput
}

func (TargetHttpsProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpsProxyState)(nil)).Elem()
}

type targetHttpsProxyArgs struct {
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// -
	// (Required)
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
	// -
	// (Optional)
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE. Not specifying this field is equivalent to
	// specifying NONE.
	QuicOverride *string `pulumi:"quicOverride"`
	// -
	// (Required)
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// -
	// (Optional)
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
	// -
	// (Required)
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap string `pulumi:"urlMap"`
}

// The set of arguments for constructing a TargetHttpsProxy resource.
type TargetHttpsProxyArgs struct {
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// -
	// (Required)
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
	// -
	// (Optional)
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, uses the QUIC policy with no user overrides, which is
	// equivalent to DISABLE. Not specifying this field is equivalent to
	// specifying NONE.
	QuicOverride pulumi.StringPtrInput
	// -
	// (Required)
	// A list of SslCertificate resources that are used to authenticate
	// connections between users and the load balancer. At least one SSL
	// certificate must be specified.
	SslCertificates pulumi.StringArrayInput
	// -
	// (Optional)
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
	// -
	// (Required)
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringInput
}

func (TargetHttpsProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpsProxyArgs)(nil)).Elem()
}
