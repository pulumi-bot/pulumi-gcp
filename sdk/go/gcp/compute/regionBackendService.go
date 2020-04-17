// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Region Backend Service defines a regionally-scoped group of virtual
// machines that will serve traffic for load balancing.
//
//
// To get more information about RegionBackendService, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/latest/regionBackendServices)
// * How-to Guides
//     * [Internal TCP/UDP Load Balancing](https://cloud.google.com/compute/docs/load-balancing/internal/)
type RegionBackendService struct {
	pulumi.CustomResourceState

	// -
	// (Optional)
	// Lifetime of cookies in seconds if sessionAffinity is
	// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	// only until the end of the browser session (or equivalent). The
	// maximum allowed value for TTL is one day.
	// When the load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec pulumi.IntPtrOutput `pulumi:"affinityCookieTtlSec"`
	// -
	// (Optional)
	// The set of backends that serve this RegionBackendService.  Structure is documented below.
	Backends RegionBackendServiceBackendArrayOutput `pulumi:"backends"`
	// -
	// (Optional)
	// Settings controlling the volume of connections to a backend service. This field
	// is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
	// and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	CircuitBreakers RegionBackendServiceCircuitBreakersPtrOutput `pulumi:"circuitBreakers"`
	// -
	// (Optional)
	// Time for which instance will be drained (not accept new
	// connections, but still work to finish started).
	ConnectionDrainingTimeoutSec pulumi.IntPtrOutput `pulumi:"connectionDrainingTimeoutSec"`
	// -
	// (Optional)
	// Consistent Hash-based load balancing can be used to provide soft session
	// affinity based on HTTP headers, cookies or other properties. This load balancing
	// policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the
	// destination service. This field specifies parameters that control consistent
	// hashing.
	// This field only applies when all of the following are true -
	ConsistentHash RegionBackendServiceConsistentHashPtrOutput `pulumi:"consistentHash"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// -
	// (Optional)
	// Policy for failovers.  Structure is documented below.
	FailoverPolicy RegionBackendServiceFailoverPolicyPtrOutput `pulumi:"failoverPolicy"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// -
	// (Required)
	// The set of URLs to HealthCheck resources for health checking
	// this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks pulumi.StringOutput `pulumi:"healthChecks"`
	// is set to INTERNAL_MANAGED
	LoadBalancingScheme pulumi.StringPtrOutput `pulumi:"loadBalancingScheme"`
	// is set to MAGLEV or RING_HASH  Structure is documented below.
	LocalityLbPolicy pulumi.StringPtrOutput `pulumi:"localityLbPolicy"`
	// -
	// (Optional)
	// This field denotes the logging options for the load balancer traffic served by this backend service.
	// If logging is enabled, logs will be exported to Stackdriver.  Structure is documented below.
	LogConfig RegionBackendServiceLogConfigPtrOutput `pulumi:"logConfig"`
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
	// -
	// (Optional)
	// The URL of the network to which this backend service belongs.
	// This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// -
	// (Optional)
	// Settings controlling eviction of unhealthy hosts from the load balancing pool.
	// This field is applicable only when the `loadBalancingScheme` is set
	// to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	OutlierDetection RegionBackendServiceOutlierDetectionPtrOutput `pulumi:"outlierDetection"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// is set to HTTP, HTTPS, or HTTP2
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// -
	// (Optional)
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// -
	// (Optional)
	// Type of session affinity to use. The default is NONE. Session affinity is
	// not applicable if the protocol is UDP.
	SessionAffinity pulumi.StringOutput `pulumi:"sessionAffinity"`
	// -
	// (Optional)
	// How many seconds to wait for the backend before considering it a
	// failed request. Default is 30 seconds. Valid range is [1, 86400].
	TimeoutSec pulumi.IntOutput `pulumi:"timeoutSec"`
}

// NewRegionBackendService registers a new resource with the given unique name, arguments, and options.
func NewRegionBackendService(ctx *pulumi.Context,
	name string, args *RegionBackendServiceArgs, opts ...pulumi.ResourceOption) (*RegionBackendService, error) {
	if args == nil || args.HealthChecks == nil {
		return nil, errors.New("missing required argument 'HealthChecks'")
	}
	if args == nil {
		args = &RegionBackendServiceArgs{}
	}
	var resource RegionBackendService
	err := ctx.RegisterResource("gcp:compute/regionBackendService:RegionBackendService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionBackendService gets an existing RegionBackendService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionBackendService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionBackendServiceState, opts ...pulumi.ResourceOption) (*RegionBackendService, error) {
	var resource RegionBackendService
	err := ctx.ReadResource("gcp:compute/regionBackendService:RegionBackendService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionBackendService resources.
type regionBackendServiceState struct {
	// -
	// (Optional)
	// Lifetime of cookies in seconds if sessionAffinity is
	// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	// only until the end of the browser session (or equivalent). The
	// maximum allowed value for TTL is one day.
	// When the load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec *int `pulumi:"affinityCookieTtlSec"`
	// -
	// (Optional)
	// The set of backends that serve this RegionBackendService.  Structure is documented below.
	Backends []RegionBackendServiceBackend `pulumi:"backends"`
	// -
	// (Optional)
	// Settings controlling the volume of connections to a backend service. This field
	// is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
	// and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	CircuitBreakers *RegionBackendServiceCircuitBreakers `pulumi:"circuitBreakers"`
	// -
	// (Optional)
	// Time for which instance will be drained (not accept new
	// connections, but still work to finish started).
	ConnectionDrainingTimeoutSec *int `pulumi:"connectionDrainingTimeoutSec"`
	// -
	// (Optional)
	// Consistent Hash-based load balancing can be used to provide soft session
	// affinity based on HTTP headers, cookies or other properties. This load balancing
	// policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the
	// destination service. This field specifies parameters that control consistent
	// hashing.
	// This field only applies when all of the following are true -
	ConsistentHash *RegionBackendServiceConsistentHash `pulumi:"consistentHash"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// Policy for failovers.  Structure is documented below.
	FailoverPolicy *RegionBackendServiceFailoverPolicy `pulumi:"failoverPolicy"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint *string `pulumi:"fingerprint"`
	// -
	// (Required)
	// The set of URLs to HealthCheck resources for health checking
	// this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks *string `pulumi:"healthChecks"`
	// is set to INTERNAL_MANAGED
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// is set to MAGLEV or RING_HASH  Structure is documented below.
	LocalityLbPolicy *string `pulumi:"localityLbPolicy"`
	// -
	// (Optional)
	// This field denotes the logging options for the load balancer traffic served by this backend service.
	// If logging is enabled, logs will be exported to Stackdriver.  Structure is documented below.
	LogConfig *RegionBackendServiceLogConfig `pulumi:"logConfig"`
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
	// -
	// (Optional)
	// The URL of the network to which this backend service belongs.
	// This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network *string `pulumi:"network"`
	// -
	// (Optional)
	// Settings controlling eviction of unhealthy hosts from the load balancing pool.
	// This field is applicable only when the `loadBalancingScheme` is set
	// to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	OutlierDetection *RegionBackendServiceOutlierDetection `pulumi:"outlierDetection"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// is set to HTTP, HTTPS, or HTTP2
	Protocol *string `pulumi:"protocol"`
	// -
	// (Optional)
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// -
	// (Optional)
	// Type of session affinity to use. The default is NONE. Session affinity is
	// not applicable if the protocol is UDP.
	SessionAffinity *string `pulumi:"sessionAffinity"`
	// -
	// (Optional)
	// How many seconds to wait for the backend before considering it a
	// failed request. Default is 30 seconds. Valid range is [1, 86400].
	TimeoutSec *int `pulumi:"timeoutSec"`
}

type RegionBackendServiceState struct {
	// -
	// (Optional)
	// Lifetime of cookies in seconds if sessionAffinity is
	// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	// only until the end of the browser session (or equivalent). The
	// maximum allowed value for TTL is one day.
	// When the load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec pulumi.IntPtrInput
	// -
	// (Optional)
	// The set of backends that serve this RegionBackendService.  Structure is documented below.
	Backends RegionBackendServiceBackendArrayInput
	// -
	// (Optional)
	// Settings controlling the volume of connections to a backend service. This field
	// is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
	// and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	CircuitBreakers RegionBackendServiceCircuitBreakersPtrInput
	// -
	// (Optional)
	// Time for which instance will be drained (not accept new
	// connections, but still work to finish started).
	ConnectionDrainingTimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// Consistent Hash-based load balancing can be used to provide soft session
	// affinity based on HTTP headers, cookies or other properties. This load balancing
	// policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the
	// destination service. This field specifies parameters that control consistent
	// hashing.
	// This field only applies when all of the following are true -
	ConsistentHash RegionBackendServiceConsistentHashPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// Policy for failovers.  Structure is documented below.
	FailoverPolicy RegionBackendServiceFailoverPolicyPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringPtrInput
	// -
	// (Required)
	// The set of URLs to HealthCheck resources for health checking
	// this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks pulumi.StringPtrInput
	// is set to INTERNAL_MANAGED
	LoadBalancingScheme pulumi.StringPtrInput
	// is set to MAGLEV or RING_HASH  Structure is documented below.
	LocalityLbPolicy pulumi.StringPtrInput
	// -
	// (Optional)
	// This field denotes the logging options for the load balancer traffic served by this backend service.
	// If logging is enabled, logs will be exported to Stackdriver.  Structure is documented below.
	LogConfig RegionBackendServiceLogConfigPtrInput
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
	// -
	// (Optional)
	// The URL of the network to which this backend service belongs.
	// This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network pulumi.StringPtrInput
	// -
	// (Optional)
	// Settings controlling eviction of unhealthy hosts from the load balancing pool.
	// This field is applicable only when the `loadBalancingScheme` is set
	// to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	OutlierDetection RegionBackendServiceOutlierDetectionPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// is set to HTTP, HTTPS, or HTTP2
	Protocol pulumi.StringPtrInput
	// -
	// (Optional)
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// -
	// (Optional)
	// Type of session affinity to use. The default is NONE. Session affinity is
	// not applicable if the protocol is UDP.
	SessionAffinity pulumi.StringPtrInput
	// -
	// (Optional)
	// How many seconds to wait for the backend before considering it a
	// failed request. Default is 30 seconds. Valid range is [1, 86400].
	TimeoutSec pulumi.IntPtrInput
}

func (RegionBackendServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionBackendServiceState)(nil)).Elem()
}

type regionBackendServiceArgs struct {
	// -
	// (Optional)
	// Lifetime of cookies in seconds if sessionAffinity is
	// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	// only until the end of the browser session (or equivalent). The
	// maximum allowed value for TTL is one day.
	// When the load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec *int `pulumi:"affinityCookieTtlSec"`
	// -
	// (Optional)
	// The set of backends that serve this RegionBackendService.  Structure is documented below.
	Backends []RegionBackendServiceBackend `pulumi:"backends"`
	// -
	// (Optional)
	// Settings controlling the volume of connections to a backend service. This field
	// is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
	// and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	CircuitBreakers *RegionBackendServiceCircuitBreakers `pulumi:"circuitBreakers"`
	// -
	// (Optional)
	// Time for which instance will be drained (not accept new
	// connections, but still work to finish started).
	ConnectionDrainingTimeoutSec *int `pulumi:"connectionDrainingTimeoutSec"`
	// -
	// (Optional)
	// Consistent Hash-based load balancing can be used to provide soft session
	// affinity based on HTTP headers, cookies or other properties. This load balancing
	// policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the
	// destination service. This field specifies parameters that control consistent
	// hashing.
	// This field only applies when all of the following are true -
	ConsistentHash *RegionBackendServiceConsistentHash `pulumi:"consistentHash"`
	// -
	// (Optional)
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// Policy for failovers.  Structure is documented below.
	FailoverPolicy *RegionBackendServiceFailoverPolicy `pulumi:"failoverPolicy"`
	// -
	// (Required)
	// The set of URLs to HealthCheck resources for health checking
	// this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks string `pulumi:"healthChecks"`
	// is set to INTERNAL_MANAGED
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// is set to MAGLEV or RING_HASH  Structure is documented below.
	LocalityLbPolicy *string `pulumi:"localityLbPolicy"`
	// -
	// (Optional)
	// This field denotes the logging options for the load balancer traffic served by this backend service.
	// If logging is enabled, logs will be exported to Stackdriver.  Structure is documented below.
	LogConfig *RegionBackendServiceLogConfig `pulumi:"logConfig"`
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
	// -
	// (Optional)
	// The URL of the network to which this backend service belongs.
	// This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network *string `pulumi:"network"`
	// -
	// (Optional)
	// Settings controlling eviction of unhealthy hosts from the load balancing pool.
	// This field is applicable only when the `loadBalancingScheme` is set
	// to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	OutlierDetection *RegionBackendServiceOutlierDetection `pulumi:"outlierDetection"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// is set to HTTP, HTTPS, or HTTP2
	Protocol *string `pulumi:"protocol"`
	// -
	// (Optional)
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// -
	// (Optional)
	// Type of session affinity to use. The default is NONE. Session affinity is
	// not applicable if the protocol is UDP.
	SessionAffinity *string `pulumi:"sessionAffinity"`
	// -
	// (Optional)
	// How many seconds to wait for the backend before considering it a
	// failed request. Default is 30 seconds. Valid range is [1, 86400].
	TimeoutSec *int `pulumi:"timeoutSec"`
}

// The set of arguments for constructing a RegionBackendService resource.
type RegionBackendServiceArgs struct {
	// -
	// (Optional)
	// Lifetime of cookies in seconds if sessionAffinity is
	// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	// only until the end of the browser session (or equivalent). The
	// maximum allowed value for TTL is one day.
	// When the load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec pulumi.IntPtrInput
	// -
	// (Optional)
	// The set of backends that serve this RegionBackendService.  Structure is documented below.
	Backends RegionBackendServiceBackendArrayInput
	// -
	// (Optional)
	// Settings controlling the volume of connections to a backend service. This field
	// is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
	// and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	CircuitBreakers RegionBackendServiceCircuitBreakersPtrInput
	// -
	// (Optional)
	// Time for which instance will be drained (not accept new
	// connections, but still work to finish started).
	ConnectionDrainingTimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// Consistent Hash-based load balancing can be used to provide soft session
	// affinity based on HTTP headers, cookies or other properties. This load balancing
	// policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the
	// destination service. This field specifies parameters that control consistent
	// hashing.
	// This field only applies when all of the following are true -
	ConsistentHash RegionBackendServiceConsistentHashPtrInput
	// -
	// (Optional)
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// Policy for failovers.  Structure is documented below.
	FailoverPolicy RegionBackendServiceFailoverPolicyPtrInput
	// -
	// (Required)
	// The set of URLs to HealthCheck resources for health checking
	// this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks pulumi.StringInput
	// is set to INTERNAL_MANAGED
	LoadBalancingScheme pulumi.StringPtrInput
	// is set to MAGLEV or RING_HASH  Structure is documented below.
	LocalityLbPolicy pulumi.StringPtrInput
	// -
	// (Optional)
	// This field denotes the logging options for the load balancer traffic served by this backend service.
	// If logging is enabled, logs will be exported to Stackdriver.  Structure is documented below.
	LogConfig RegionBackendServiceLogConfigPtrInput
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
	// -
	// (Optional)
	// The URL of the network to which this backend service belongs.
	// This field can only be specified when the load balancing scheme is set to INTERNAL.
	Network pulumi.StringPtrInput
	// -
	// (Optional)
	// Settings controlling eviction of unhealthy hosts from the load balancing pool.
	// This field is applicable only when the `loadBalancingScheme` is set
	// to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.  Structure is documented below.
	OutlierDetection RegionBackendServiceOutlierDetectionPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// is set to HTTP, HTTPS, or HTTP2
	Protocol pulumi.StringPtrInput
	// -
	// (Optional)
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// -
	// (Optional)
	// Type of session affinity to use. The default is NONE. Session affinity is
	// not applicable if the protocol is UDP.
	SessionAffinity pulumi.StringPtrInput
	// -
	// (Optional)
	// How many seconds to wait for the backend before considering it a
	// failed request. Default is 30 seconds. Valid range is [1, 86400].
	TimeoutSec pulumi.IntPtrInput
}

func (RegionBackendServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionBackendServiceArgs)(nil)).Elem()
}
