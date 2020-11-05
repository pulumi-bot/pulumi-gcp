// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"fmt"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// UrlMaps are used to route requests to a backend service based on rules
// that you define for the host and path of an incoming URL.
//
// To get more information about UrlMap, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps)
//
// ## Example Usage
type URLMap struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction URLMapDefaultRouteActionPtrOutput `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService pulumi.StringPtrOutput `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect URLMapDefaultUrlRedirectPtrOutput `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction URLMapHeaderActionPtrOutput `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules URLMapHostRuleArrayOutput `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId pulumi.IntOutput `pulumi:"mapId"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers URLMapPathMatcherArrayOutput `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests URLMapTestArrayOutput `pulumi:"tests"`
}

// NewURLMap registers a new resource with the given unique name, arguments, and options.
func NewURLMap(ctx *pulumi.Context,
	name string, args *URLMapArgs, opts ...pulumi.ResourceOption) (*URLMap, error) {
	if args == nil {
		args = &URLMapArgs{}
	}
	var resource URLMap
	err := ctx.RegisterResource("gcp:compute/uRLMap:URLMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURLMap gets an existing URLMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetURLMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *URLMapState, opts ...pulumi.ResourceOption) (*URLMap, error) {
	var resource URLMap
	err := ctx.ReadResource("gcp:compute/uRLMap:URLMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering URLMap resources.
type urlmapState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction *URLMapDefaultRouteAction `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect *URLMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint *string `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction *URLMapHeaderAction `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules []URLMapHostRule `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId *int `pulumi:"mapId"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name *string `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers []URLMapPathMatcher `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests []URLMapTest `pulumi:"tests"`
}

type URLMapState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction URLMapDefaultRouteActionPtrInput
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect URLMapDefaultUrlRedirectPtrInput
	// Description of this test case.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction URLMapHeaderActionPtrInput
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules URLMapHostRuleArrayInput
	// The unique identifier for the resource.
	MapId pulumi.IntPtrInput
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringPtrInput
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers URLMapPathMatcherArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests URLMapTestArrayInput
}

func (URLMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*urlmapState)(nil)).Elem()
}

type urlmapArgs struct {
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction *URLMapDefaultRouteAction `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect *URLMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description *string `pulumi:"description"`
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction *URLMapHeaderAction `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules []URLMapHostRule `pulumi:"hostRules"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name *string `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers []URLMapPathMatcher `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests []URLMapTest `pulumi:"tests"`
}

// The set of arguments for constructing a URLMap resource.
type URLMapArgs struct {
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction URLMapDefaultRouteActionPtrInput
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect URLMapDefaultUrlRedirectPtrInput
	// Description of this test case.
	Description pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction URLMapHeaderActionPtrInput
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules URLMapHostRuleArrayInput
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringPtrInput
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers URLMapPathMatcherArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests URLMapTestArrayInput
}

func (URLMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*urlmapArgs)(nil)).Elem()
}
