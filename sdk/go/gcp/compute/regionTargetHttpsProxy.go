// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a RegionTargetHttpsProxy resource, which is used by one or more
// forwarding rules to route incoming HTTPS requests to a URL map.
//
// To get more information about RegionTargetHttpsProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/regionTargetHttpsProxies)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
type RegionTargetHttpsProxy struct {
	pulumi.CustomResourceState

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
	// The unique identifier for the resource.
	ProxyId pulumi.IntOutput `pulumi:"proxyId"`
	// The Region in which the created target https proxy should reside. If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of RegionSslCertificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, exactly one SSL certificate must be specified.
	SslCertificates pulumi.StringArrayOutput `pulumi:"sslCertificates"`
	// A reference to the RegionUrlMap resource that defines the mapping from URL to the RegionBackendService.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewRegionTargetHttpsProxy registers a new resource with the given unique name, arguments, and options.
func NewRegionTargetHttpsProxy(ctx *pulumi.Context,
	name string, args *RegionTargetHttpsProxyArgs, opts ...pulumi.ResourceOption) (*RegionTargetHttpsProxy, error) {
	if args == nil || args.SslCertificates == nil {
		return nil, errors.New("missing required argument 'SslCertificates'")
	}
	if args == nil || args.UrlMap == nil {
		return nil, errors.New("missing required argument 'UrlMap'")
	}
	if args == nil {
		args = &RegionTargetHttpsProxyArgs{}
	}
	var resource RegionTargetHttpsProxy
	err := ctx.RegisterResource("gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionTargetHttpsProxy gets an existing RegionTargetHttpsProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionTargetHttpsProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionTargetHttpsProxyState, opts ...pulumi.ResourceOption) (*RegionTargetHttpsProxy, error) {
	var resource RegionTargetHttpsProxy
	err := ctx.ReadResource("gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionTargetHttpsProxy resources.
type regionTargetHttpsProxyState struct {
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
	// The unique identifier for the resource.
	ProxyId *int `pulumi:"proxyId"`
	// The Region in which the created target https proxy should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A list of RegionSslCertificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, exactly one SSL certificate must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// A reference to the RegionUrlMap resource that defines the mapping from URL to the RegionBackendService.
	UrlMap *string `pulumi:"urlMap"`
}

type RegionTargetHttpsProxyState struct {
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
	// The unique identifier for the resource.
	ProxyId pulumi.IntPtrInput
	// The Region in which the created target https proxy should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A list of RegionSslCertificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, exactly one SSL certificate must be specified.
	SslCertificates pulumi.StringArrayInput
	// A reference to the RegionUrlMap resource that defines the mapping from URL to the RegionBackendService.
	UrlMap pulumi.StringPtrInput
}

func (RegionTargetHttpsProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionTargetHttpsProxyState)(nil)).Elem()
}

type regionTargetHttpsProxyArgs struct {
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
	// The Region in which the created target https proxy should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// A list of RegionSslCertificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, exactly one SSL certificate must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// A reference to the RegionUrlMap resource that defines the mapping from URL to the RegionBackendService.
	UrlMap string `pulumi:"urlMap"`
}

// The set of arguments for constructing a RegionTargetHttpsProxy resource.
type RegionTargetHttpsProxyArgs struct {
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
	// The Region in which the created target https proxy should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// A list of RegionSslCertificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, exactly one SSL certificate must be specified.
	SslCertificates pulumi.StringArrayInput
	// A reference to the RegionUrlMap resource that defines the mapping from URL to the RegionBackendService.
	UrlMap pulumi.StringInput
}

func (RegionTargetHttpsProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionTargetHttpsProxyArgs)(nil)).Elem()
}
