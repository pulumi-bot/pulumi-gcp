// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a TargetHttpProxy resource, which is used by one or more global
// forwarding rule to route incoming HTTP requests to a URL map.
//
// To get more information about TargetHttpProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpProxies)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
//
// ## Example Usage
type TargetHttpProxy struct {
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
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewTargetHttpProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpProxy(ctx *pulumi.Context,
	name string, args *TargetHttpProxyArgs, opts ...pulumi.ResourceOption) (*TargetHttpProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.UrlMap == nil {
		return nil, errors.New("invalid value for required argument 'UrlMap'")
	}
	var resource TargetHttpProxy
	err := ctx.RegisterResource("gcp:compute/targetHttpProxy:TargetHttpProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetHttpProxy gets an existing TargetHttpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetHttpProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetHttpProxyState, opts ...pulumi.ResourceOption) (*TargetHttpProxy, error) {
	var resource TargetHttpProxy
	err := ctx.ReadResource("gcp:compute/targetHttpProxy:TargetHttpProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetHttpProxy resources.
type targetHttpProxyState struct {
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
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap *string `pulumi:"urlMap"`
}

type TargetHttpProxyState struct {
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
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringPtrInput
}

func (TargetHttpProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpProxyState)(nil)).Elem()
}

type targetHttpProxyArgs struct {
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
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap string `pulumi:"urlMap"`
}

// The set of arguments for constructing a TargetHttpProxy resource.
type TargetHttpProxyArgs struct {
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
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringInput
}

func (TargetHttpProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpProxyArgs)(nil)).Elem()
}
