// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// UrlMaps are used to route requests to a backend service based on rules
// that you define for the host and path of an incoming URL.
//
// ## Example Usage
//
// ## Import
//
// RegionUrlMap can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default projects/{{project}}/regions/{{region}}/urlMaps/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{name}}
// ```
type RegionUrlMap struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService pulumi.StringPtrOutput `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect RegionUrlMapDefaultUrlRedirectPtrOutput `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules RegionUrlMapHostRuleArrayOutput `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId pulumi.IntOutput `pulumi:"mapId"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers RegionUrlMapPathMatcherArrayOutput `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests RegionUrlMapTestArrayOutput `pulumi:"tests"`
}

// NewRegionUrlMap registers a new resource with the given unique name, arguments, and options.
func NewRegionUrlMap(ctx *pulumi.Context,
	name string, args *RegionUrlMapArgs, opts ...pulumi.ResourceOption) (*RegionUrlMap, error) {
	if args == nil {
		args = &RegionUrlMapArgs{}
	}

	var resource RegionUrlMap
	err := ctx.RegisterResource("gcp:compute/regionUrlMap:RegionUrlMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionUrlMap gets an existing RegionUrlMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionUrlMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionUrlMapState, opts ...pulumi.ResourceOption) (*RegionUrlMap, error) {
	var resource RegionUrlMap
	err := ctx.ReadResource("gcp:compute/regionUrlMap:RegionUrlMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionUrlMap resources.
type regionUrlMapState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect *RegionUrlMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint *string `pulumi:"fingerprint"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules []RegionUrlMapHostRule `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId *int `pulumi:"mapId"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name *string `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers []RegionUrlMapPathMatcher `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests []RegionUrlMapTest `pulumi:"tests"`
}

type RegionUrlMapState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect RegionUrlMapDefaultUrlRedirectPtrInput
	// Description of this test case.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringPtrInput
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules RegionUrlMapHostRuleArrayInput
	// The unique identifier for the resource.
	MapId pulumi.IntPtrInput
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringPtrInput
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers RegionUrlMapPathMatcherArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests RegionUrlMapTestArrayInput
}

func (RegionUrlMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionUrlMapState)(nil)).Elem()
}

type regionUrlMapArgs struct {
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect *RegionUrlMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description *string `pulumi:"description"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules []RegionUrlMapHostRule `pulumi:"hostRules"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name *string `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers []RegionUrlMapPathMatcher `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests []RegionUrlMapTest `pulumi:"tests"`
}

// The set of arguments for constructing a RegionUrlMap resource.
type RegionUrlMapArgs struct {
	// A reference to a RegionBackendService resource. This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect RegionUrlMapDefaultUrlRedirectPtrInput
	// Description of this test case.
	Description pulumi.StringPtrInput
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules RegionUrlMapHostRuleArrayInput
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringPtrInput
	// The name of the PathMatcher to use to match the path portion of
	// the URL if the hostRule matches the URL's host portion.
	PathMatchers RegionUrlMapPathMatcherArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the url map should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The list of expected URL mappings. Requests to update this UrlMap will
	// succeed only if all of the test cases pass.
	// Structure is documented below.
	Tests RegionUrlMapTestArrayInput
}

func (RegionUrlMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionUrlMapArgs)(nil)).Elem()
}

type RegionUrlMapInput interface {
	pulumi.Input

	ToRegionUrlMapOutput() RegionUrlMapOutput
	ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput
}

func (*RegionUrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMap)(nil))
}

func (i *RegionUrlMap) ToRegionUrlMapOutput() RegionUrlMapOutput {
	return i.ToRegionUrlMapOutputWithContext(context.Background())
}

func (i *RegionUrlMap) ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapOutput)
}

func (i *RegionUrlMap) ToRegionUrlMapPtrOutput() RegionUrlMapPtrOutput {
	return i.ToRegionUrlMapPtrOutputWithContext(context.Background())
}

func (i *RegionUrlMap) ToRegionUrlMapPtrOutputWithContext(ctx context.Context) RegionUrlMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapPtrOutput)
}

type RegionUrlMapPtrInput interface {
	pulumi.Input

	ToRegionUrlMapPtrOutput() RegionUrlMapPtrOutput
	ToRegionUrlMapPtrOutputWithContext(ctx context.Context) RegionUrlMapPtrOutput
}

type regionUrlMapPtrType RegionUrlMapArgs

func (*regionUrlMapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionUrlMap)(nil))
}

func (i *regionUrlMapPtrType) ToRegionUrlMapPtrOutput() RegionUrlMapPtrOutput {
	return i.ToRegionUrlMapPtrOutputWithContext(context.Background())
}

func (i *regionUrlMapPtrType) ToRegionUrlMapPtrOutputWithContext(ctx context.Context) RegionUrlMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapPtrOutput)
}

// RegionUrlMapArrayInput is an input type that accepts RegionUrlMapArray and RegionUrlMapArrayOutput values.
// You can construct a concrete instance of `RegionUrlMapArrayInput` via:
//
//          RegionUrlMapArray{ RegionUrlMapArgs{...} }
type RegionUrlMapArrayInput interface {
	pulumi.Input

	ToRegionUrlMapArrayOutput() RegionUrlMapArrayOutput
	ToRegionUrlMapArrayOutputWithContext(context.Context) RegionUrlMapArrayOutput
}

type RegionUrlMapArray []RegionUrlMapInput

func (RegionUrlMapArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RegionUrlMap)(nil))
}

func (i RegionUrlMapArray) ToRegionUrlMapArrayOutput() RegionUrlMapArrayOutput {
	return i.ToRegionUrlMapArrayOutputWithContext(context.Background())
}

func (i RegionUrlMapArray) ToRegionUrlMapArrayOutputWithContext(ctx context.Context) RegionUrlMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapArrayOutput)
}

// RegionUrlMapMapInput is an input type that accepts RegionUrlMapMap and RegionUrlMapMapOutput values.
// You can construct a concrete instance of `RegionUrlMapMapInput` via:
//
//          RegionUrlMapMap{ "key": RegionUrlMapArgs{...} }
type RegionUrlMapMapInput interface {
	pulumi.Input

	ToRegionUrlMapMapOutput() RegionUrlMapMapOutput
	ToRegionUrlMapMapOutputWithContext(context.Context) RegionUrlMapMapOutput
}

type RegionUrlMapMap map[string]RegionUrlMapInput

func (RegionUrlMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RegionUrlMap)(nil))
}

func (i RegionUrlMapMap) ToRegionUrlMapMapOutput() RegionUrlMapMapOutput {
	return i.ToRegionUrlMapMapOutputWithContext(context.Background())
}

func (i RegionUrlMapMap) ToRegionUrlMapMapOutputWithContext(ctx context.Context) RegionUrlMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionUrlMapMapOutput)
}

type RegionUrlMapOutput struct {
	*pulumi.OutputState
}

func (RegionUrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionUrlMap)(nil))
}

func (o RegionUrlMapOutput) ToRegionUrlMapOutput() RegionUrlMapOutput {
	return o
}

func (o RegionUrlMapOutput) ToRegionUrlMapOutputWithContext(ctx context.Context) RegionUrlMapOutput {
	return o
}

func (o RegionUrlMapOutput) ToRegionUrlMapPtrOutput() RegionUrlMapPtrOutput {
	return o.ToRegionUrlMapPtrOutputWithContext(context.Background())
}

func (o RegionUrlMapOutput) ToRegionUrlMapPtrOutputWithContext(ctx context.Context) RegionUrlMapPtrOutput {
	return o.ApplyT(func(v RegionUrlMap) *RegionUrlMap {
		return &v
	}).(RegionUrlMapPtrOutput)
}

type RegionUrlMapPtrOutput struct {
	*pulumi.OutputState
}

func (RegionUrlMapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionUrlMap)(nil))
}

func (o RegionUrlMapPtrOutput) ToRegionUrlMapPtrOutput() RegionUrlMapPtrOutput {
	return o
}

func (o RegionUrlMapPtrOutput) ToRegionUrlMapPtrOutputWithContext(ctx context.Context) RegionUrlMapPtrOutput {
	return o
}

type RegionUrlMapArrayOutput struct{ *pulumi.OutputState }

func (RegionUrlMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegionUrlMap)(nil))
}

func (o RegionUrlMapArrayOutput) ToRegionUrlMapArrayOutput() RegionUrlMapArrayOutput {
	return o
}

func (o RegionUrlMapArrayOutput) ToRegionUrlMapArrayOutputWithContext(ctx context.Context) RegionUrlMapArrayOutput {
	return o
}

func (o RegionUrlMapArrayOutput) Index(i pulumi.IntInput) RegionUrlMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegionUrlMap {
		return vs[0].([]RegionUrlMap)[vs[1].(int)]
	}).(RegionUrlMapOutput)
}

type RegionUrlMapMapOutput struct{ *pulumi.OutputState }

func (RegionUrlMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RegionUrlMap)(nil))
}

func (o RegionUrlMapMapOutput) ToRegionUrlMapMapOutput() RegionUrlMapMapOutput {
	return o
}

func (o RegionUrlMapMapOutput) ToRegionUrlMapMapOutputWithContext(ctx context.Context) RegionUrlMapMapOutput {
	return o
}

func (o RegionUrlMapMapOutput) MapIndex(k pulumi.StringInput) RegionUrlMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RegionUrlMap {
		return vs[0].(map[string]RegionUrlMap)[vs[1].(string)]
	}).(RegionUrlMapOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionUrlMapOutput{})
	pulumi.RegisterOutputType(RegionUrlMapPtrOutput{})
	pulumi.RegisterOutputType(RegionUrlMapArrayOutput{})
	pulumi.RegisterOutputType(RegionUrlMapMapOutput{})
}
