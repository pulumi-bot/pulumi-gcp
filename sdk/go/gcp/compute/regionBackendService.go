// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
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

	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec pulumi.IntPtrOutput `pulumi:"affinityCookieTtlSec"`
	// The set of backends that serve this RegionBackendService.
	Backends RegionBackendServiceBackendArrayOutput `pulumi:"backends"`
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	CircuitBreakers RegionBackendServiceCircuitBreakersPtrOutput `pulumi:"circuitBreakers"`
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec pulumi.IntPtrOutput `pulumi:"connectionDrainingTimeoutSec"`
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies when all of the following are true - *
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED * 'protocol' is set to HTTP, HTTPS, or HTTP2 * 'locality_lb_policy'
	// is set to MAGLEV or RING_HASH
	ConsistentHash RegionBackendServiceConsistentHashPtrOutput `pulumi:"consistentHash"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Policy for failovers.
	FailoverPolicy RegionBackendServiceFailoverPolicyPtrOutput `pulumi:"failoverPolicy"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks pulumi.StringOutput `pulumi:"healthChecks"`
	// Indicates what kind of load balancing this regional backend service will be used for. A backend service created for one
	// type of load balancing cannot be used with the other(s). Must be 'INTERNAL' or 'INTERNAL_MANAGED'. Defaults to
	// 'INTERNAL'.
	LoadBalancingScheme pulumi.StringPtrOutput `pulumi:"loadBalancingScheme"`
	// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
	// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
	// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
	// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
	// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
	// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
	// the same address as the destination address of the incoming connection before the connection was redirected to the load
	// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
	// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the 'load_balancing_scheme' is set to
	// INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	LocalityLbPolicy pulumi.StringPtrOutput `pulumi:"localityLbPolicy"`
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig RegionBackendServiceLogConfigPtrOutput `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing
	// scheme is set to INTERNAL.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	OutlierDetection RegionBackendServiceOutlierDetectionPtrOutput `pulumi:"outlierDetection"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The protocol this RegionBackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, SSL,
	// TCP, and UDP. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in
	// errors if used with the GA API.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity pulumi.StringOutput `pulumi:"sessionAffinity"`
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
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
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec *int `pulumi:"affinityCookieTtlSec"`
	// The set of backends that serve this RegionBackendService.
	Backends []RegionBackendServiceBackend `pulumi:"backends"`
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	CircuitBreakers *RegionBackendServiceCircuitBreakers `pulumi:"circuitBreakers"`
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec *int `pulumi:"connectionDrainingTimeoutSec"`
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies when all of the following are true - *
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED * 'protocol' is set to HTTP, HTTPS, or HTTP2 * 'locality_lb_policy'
	// is set to MAGLEV or RING_HASH
	ConsistentHash *RegionBackendServiceConsistentHash `pulumi:"consistentHash"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Policy for failovers.
	FailoverPolicy *RegionBackendServiceFailoverPolicy `pulumi:"failoverPolicy"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint *string `pulumi:"fingerprint"`
	// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks *string `pulumi:"healthChecks"`
	// Indicates what kind of load balancing this regional backend service will be used for. A backend service created for one
	// type of load balancing cannot be used with the other(s). Must be 'INTERNAL' or 'INTERNAL_MANAGED'. Defaults to
	// 'INTERNAL'.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
	// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
	// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
	// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
	// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
	// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
	// the same address as the destination address of the incoming connection before the connection was redirected to the load
	// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
	// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the 'load_balancing_scheme' is set to
	// INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	LocalityLbPolicy *string `pulumi:"localityLbPolicy"`
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig *RegionBackendServiceLogConfig `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing
	// scheme is set to INTERNAL.
	Network *string `pulumi:"network"`
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	OutlierDetection *RegionBackendServiceOutlierDetection `pulumi:"outlierDetection"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The protocol this RegionBackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, SSL,
	// TCP, and UDP. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in
	// errors if used with the GA API.
	Protocol *string `pulumi:"protocol"`
	// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity *string `pulumi:"sessionAffinity"`
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
	TimeoutSec *int `pulumi:"timeoutSec"`
}

type RegionBackendServiceState struct {
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec pulumi.IntPtrInput
	// The set of backends that serve this RegionBackendService.
	Backends RegionBackendServiceBackendArrayInput
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	CircuitBreakers RegionBackendServiceCircuitBreakersPtrInput
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec pulumi.IntPtrInput
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies when all of the following are true - *
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED * 'protocol' is set to HTTP, HTTPS, or HTTP2 * 'locality_lb_policy'
	// is set to MAGLEV or RING_HASH
	ConsistentHash RegionBackendServiceConsistentHashPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Policy for failovers.
	FailoverPolicy RegionBackendServiceFailoverPolicyPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringPtrInput
	// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks pulumi.StringPtrInput
	// Indicates what kind of load balancing this regional backend service will be used for. A backend service created for one
	// type of load balancing cannot be used with the other(s). Must be 'INTERNAL' or 'INTERNAL_MANAGED'. Defaults to
	// 'INTERNAL'.
	LoadBalancingScheme pulumi.StringPtrInput
	// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
	// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
	// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
	// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
	// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
	// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
	// the same address as the destination address of the incoming connection before the connection was redirected to the load
	// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
	// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the 'load_balancing_scheme' is set to
	// INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	LocalityLbPolicy pulumi.StringPtrInput
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig RegionBackendServiceLogConfigPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing
	// scheme is set to INTERNAL.
	Network pulumi.StringPtrInput
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	OutlierDetection RegionBackendServiceOutlierDetectionPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The protocol this RegionBackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, SSL,
	// TCP, and UDP. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in
	// errors if used with the GA API.
	Protocol pulumi.StringPtrInput
	// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity pulumi.StringPtrInput
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
	TimeoutSec pulumi.IntPtrInput
}

func (RegionBackendServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionBackendServiceState)(nil)).Elem()
}

type regionBackendServiceArgs struct {
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec *int `pulumi:"affinityCookieTtlSec"`
	// The set of backends that serve this RegionBackendService.
	Backends []RegionBackendServiceBackend `pulumi:"backends"`
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	CircuitBreakers *RegionBackendServiceCircuitBreakers `pulumi:"circuitBreakers"`
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec *int `pulumi:"connectionDrainingTimeoutSec"`
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies when all of the following are true - *
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED * 'protocol' is set to HTTP, HTTPS, or HTTP2 * 'locality_lb_policy'
	// is set to MAGLEV or RING_HASH
	ConsistentHash *RegionBackendServiceConsistentHash `pulumi:"consistentHash"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Policy for failovers.
	FailoverPolicy *RegionBackendServiceFailoverPolicy `pulumi:"failoverPolicy"`
	// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks string `pulumi:"healthChecks"`
	// Indicates what kind of load balancing this regional backend service will be used for. A backend service created for one
	// type of load balancing cannot be used with the other(s). Must be 'INTERNAL' or 'INTERNAL_MANAGED'. Defaults to
	// 'INTERNAL'.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
	// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
	// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
	// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
	// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
	// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
	// the same address as the destination address of the incoming connection before the connection was redirected to the load
	// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
	// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the 'load_balancing_scheme' is set to
	// INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	LocalityLbPolicy *string `pulumi:"localityLbPolicy"`
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig *RegionBackendServiceLogConfig `pulumi:"logConfig"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing
	// scheme is set to INTERNAL.
	Network *string `pulumi:"network"`
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	OutlierDetection *RegionBackendServiceOutlierDetection `pulumi:"outlierDetection"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The protocol this RegionBackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, SSL,
	// TCP, and UDP. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in
	// errors if used with the GA API.
	Protocol *string `pulumi:"protocol"`
	// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity *string `pulumi:"sessionAffinity"`
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
	TimeoutSec *int `pulumi:"timeoutSec"`
}

// The set of arguments for constructing a RegionBackendService resource.
type RegionBackendServiceArgs struct {
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and
	// lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the
	// load balancing scheme is INTERNAL, this field is not used.
	AffinityCookieTtlSec pulumi.IntPtrInput
	// The set of backends that serve this RegionBackendService.
	Backends RegionBackendServiceBackendArrayInput
	// Settings controlling the volume of connections to a backend service. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	CircuitBreakers RegionBackendServiceCircuitBreakersPtrInput
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec pulumi.IntPtrInput
	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
	// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
	// destination host will be lost when one or more hosts are added/removed from the destination service. This field
	// specifies parameters that control consistent hashing. This field only applies when all of the following are true - *
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED * 'protocol' is set to HTTP, HTTPS, or HTTP2 * 'locality_lb_policy'
	// is set to MAGLEV or RING_HASH
	ConsistentHash RegionBackendServiceConsistentHashPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Policy for failovers.
	FailoverPolicy RegionBackendServiceFailoverPolicyPtrInput
	// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks pulumi.StringInput
	// Indicates what kind of load balancing this regional backend service will be used for. A backend service created for one
	// type of load balancing cannot be used with the other(s). Must be 'INTERNAL' or 'INTERNAL_MANAGED'. Defaults to
	// 'INTERNAL'.
	LoadBalancingScheme pulumi.StringPtrInput
	// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
	// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
	// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
	// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
	// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
	// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
	// the same address as the destination address of the incoming connection before the connection was redirected to the load
	// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
	// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
	// https://ai.google/research/pubs/pub44824 This field is applicable only when the 'load_balancing_scheme' is set to
	// INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	LocalityLbPolicy pulumi.StringPtrInput
	// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
	// enabled, logs will be exported to Stackdriver.
	LogConfig RegionBackendServiceLogConfigPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which this backend service belongs. This field can only be specified when the load balancing
	// scheme is set to INTERNAL.
	Network pulumi.StringPtrInput
	// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
	// 'load_balancing_scheme' is set to INTERNAL_MANAGED and the 'protocol' is set to HTTP, HTTPS, or HTTP2.
	OutlierDetection RegionBackendServiceOutlierDetectionPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The protocol this RegionBackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, SSL,
	// TCP, and UDP. The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer types and may result in
	// errors if used with the GA API.
	Protocol pulumi.StringPtrInput
	// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// Type of session affinity to use. The default is NONE. Session affinity is not applicable if the protocol is UDP.
	SessionAffinity pulumi.StringPtrInput
	// How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is
	// [1, 86400].
	TimeoutSec pulumi.IntPtrInput
}

func (RegionBackendServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionBackendServiceArgs)(nil)).Elem()
}
