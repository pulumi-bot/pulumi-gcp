// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// A Backend Service defines a group of virtual machines that will serve
    /// traffic for load balancing. This resource is a global backend service,
    /// appropriate for external load balancing or self-managed internal load balancing.
    /// For managed internal load balancing, use a regional backend service instead.
    /// 
    /// Currently self-managed internal load balancing is only available in beta.
    /// 
    /// 
    /// To get more information about BackendService, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendServices)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/backend-service)
    /// </summary>
    public partial class BackendService : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Lifetime of cookies in seconds if session_affinity is
        /// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
        /// only until the end of the browser session (or equivalent). The
        /// maximum allowed value for TTL is one day.
        /// When the load balancing scheme is INTERNAL, this field is not used.
        /// </summary>
        [Output("affinityCookieTtlSec")]
        public Output<int?> AffinityCookieTtlSec { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The set of backends that serve this BackendService.  Structure is documented below.
        /// </summary>
        [Output("backends")]
        public Output<ImmutableArray<Outputs.BackendServiceBackend>> Backends { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Cloud CDN configuration for this BackendService.  Structure is documented below.
        /// </summary>
        [Output("cdnPolicy")]
        public Output<Outputs.BackendServiceCdnPolicy> CdnPolicy { get; private set; } = null!;

        /// <summary>
        /// Settings controlling the volume of connections to a backend service. This field is applicable only when the
        /// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Output("circuitBreakers")]
        public Output<Outputs.BackendServiceCircuitBreakers?> CircuitBreakers { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Time for which instance will be drained (not accept new
        /// connections, but still work to finish started).
        /// </summary>
        [Output("connectionDrainingTimeoutSec")]
        public Output<int?> ConnectionDrainingTimeoutSec { get; private set; } = null!;

        /// <summary>
        /// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
        /// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
        /// destination host will be lost when one or more hosts are added/removed from the destination service. This field
        /// specifies parameters that control consistent hashing. This field only applies if the load_balancing_scheme is set to
        /// INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is set to MAGLEV or RING_HASH.
        /// </summary>
        [Output("consistentHash")]
        public Output<Outputs.BackendServiceConsistentHash?> ConsistentHash { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied requests.
        /// </summary>
        [Output("customRequestHeaders")]
        public Output<ImmutableArray<string>> CustomRequestHeaders { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// If true, enable Cloud CDN for this BackendService.
        /// </summary>
        [Output("enableCdn")]
        public Output<bool?> EnableCdn { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource
        /// for health checking this BackendService. Currently at most one health
        /// check can be specified, and a health check is required.
        /// For internal load balancing, a URL to a HealthCheck resource must be specified instead.
        /// </summary>
        [Output("healthChecks")]
        public Output<string> HealthChecks { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Settings for enabling Cloud Identity Aware Proxy  Structure is documented below.
        /// </summary>
        [Output("iap")]
        public Output<Outputs.BackendServiceIap?> Iap { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Indicates whether the backend service will be used with internal or
        /// external load balancing. A backend service created for one type of
        /// load balancing cannot be used with the other. Must be `EXTERNAL` or
        /// `INTERNAL_SELF_MANAGED` for a global backend service. Defaults to `EXTERNAL`.
        /// </summary>
        [Output("loadBalancingScheme")]
        public Output<string?> LoadBalancingScheme { get; private set; } = null!;

        /// <summary>
        /// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
        /// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
        /// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
        /// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
        /// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
        /// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
        /// the same address as the destination address of the incoming connection before the connection was redirected to the load
        /// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
        /// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
        /// https://ai.google/research/pubs/pub44824 This field is applicable only when the load_balancing_scheme is set to
        /// INTERNAL_SELF_MANAGED.
        /// </summary>
        [Output("localityLbPolicy")]
        public Output<string?> LocalityLbPolicy { get; private set; } = null!;

        /// <summary>
        /// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
        /// enabled, logs will be exported to Stackdriver.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.BackendServiceLogConfig> LogConfig { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Name of the cookie.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
        /// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Output("outlierDetection")]
        public Output<Outputs.BackendServiceOutlierDetection?> OutlierDetection { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Name of backend port. The same name should appear in the instance
        /// groups referenced by this service. Required when the load balancing
        /// scheme is EXTERNAL.
        /// </summary>
        [Output("portName")]
        public Output<string> PortName { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The protocol this BackendService uses to communicate with backends.
        /// Possible values are HTTP, HTTPS, HTTP2, TCP, and SSL. The default is
        /// HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
        /// types and may result in errors if used with the GA API.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The security policy associated with this backend service.
        /// </summary>
        [Output("securityPolicy")]
        public Output<string?> SecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Type of session affinity to use. The default is NONE. Session affinity is
        /// not applicable if the protocol is UDP.
        /// </summary>
        [Output("sessionAffinity")]
        public Output<string> SessionAffinity { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// How many seconds to wait for the backend before considering it a
        /// failed request. Default is 30 seconds. Valid range is [1, 86400].
        /// </summary>
        [Output("timeoutSec")]
        public Output<int> TimeoutSec { get; private set; } = null!;


        /// <summary>
        /// Create a BackendService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendService(string name, BackendServiceArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/backendService:BackendService", name, args ?? new BackendServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendService(string name, Input<string> id, BackendServiceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/backendService:BackendService", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BackendService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendService Get(string name, Input<string> id, BackendServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new BackendService(name, id, state, options);
        }
    }

    public sealed class BackendServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Lifetime of cookies in seconds if session_affinity is
        /// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
        /// only until the end of the browser session (or equivalent). The
        /// maximum allowed value for TTL is one day.
        /// When the load balancing scheme is INTERNAL, this field is not used.
        /// </summary>
        [Input("affinityCookieTtlSec")]
        public Input<int>? AffinityCookieTtlSec { get; set; }

        [Input("backends")]
        private InputList<Inputs.BackendServiceBackendArgs>? _backends;

        /// <summary>
        /// -
        /// (Optional)
        /// The set of backends that serve this BackendService.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.BackendServiceBackendArgs> Backends
        {
            get => _backends ?? (_backends = new InputList<Inputs.BackendServiceBackendArgs>());
            set => _backends = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Cloud CDN configuration for this BackendService.  Structure is documented below.
        /// </summary>
        [Input("cdnPolicy")]
        public Input<Inputs.BackendServiceCdnPolicyArgs>? CdnPolicy { get; set; }

        /// <summary>
        /// Settings controlling the volume of connections to a backend service. This field is applicable only when the
        /// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("circuitBreakers")]
        public Input<Inputs.BackendServiceCircuitBreakersArgs>? CircuitBreakers { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Time for which instance will be drained (not accept new
        /// connections, but still work to finish started).
        /// </summary>
        [Input("connectionDrainingTimeoutSec")]
        public Input<int>? ConnectionDrainingTimeoutSec { get; set; }

        /// <summary>
        /// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
        /// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
        /// destination host will be lost when one or more hosts are added/removed from the destination service. This field
        /// specifies parameters that control consistent hashing. This field only applies if the load_balancing_scheme is set to
        /// INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is set to MAGLEV or RING_HASH.
        /// </summary>
        [Input("consistentHash")]
        public Input<Inputs.BackendServiceConsistentHashArgs>? ConsistentHash { get; set; }

        [Input("customRequestHeaders")]
        private InputList<string>? _customRequestHeaders;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied requests.
        /// </summary>
        public InputList<string> CustomRequestHeaders
        {
            get => _customRequestHeaders ?? (_customRequestHeaders = new InputList<string>());
            set => _customRequestHeaders = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// If true, enable Cloud CDN for this BackendService.
        /// </summary>
        [Input("enableCdn")]
        public Input<bool>? EnableCdn { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource
        /// for health checking this BackendService. Currently at most one health
        /// check can be specified, and a health check is required.
        /// For internal load balancing, a URL to a HealthCheck resource must be specified instead.
        /// </summary>
        [Input("healthChecks", required: true)]
        public Input<string> HealthChecks { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Settings for enabling Cloud Identity Aware Proxy  Structure is documented below.
        /// </summary>
        [Input("iap")]
        public Input<Inputs.BackendServiceIapArgs>? Iap { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Indicates whether the backend service will be used with internal or
        /// external load balancing. A backend service created for one type of
        /// load balancing cannot be used with the other. Must be `EXTERNAL` or
        /// `INTERNAL_SELF_MANAGED` for a global backend service. Defaults to `EXTERNAL`.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        /// <summary>
        /// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
        /// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
        /// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
        /// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
        /// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
        /// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
        /// the same address as the destination address of the incoming connection before the connection was redirected to the load
        /// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
        /// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
        /// https://ai.google/research/pubs/pub44824 This field is applicable only when the load_balancing_scheme is set to
        /// INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("localityLbPolicy")]
        public Input<string>? LocalityLbPolicy { get; set; }

        /// <summary>
        /// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
        /// enabled, logs will be exported to Stackdriver.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.BackendServiceLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Name of the cookie.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
        /// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("outlierDetection")]
        public Input<Inputs.BackendServiceOutlierDetectionArgs>? OutlierDetection { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Name of backend port. The same name should appear in the instance
        /// groups referenced by this service. Required when the load balancing
        /// scheme is EXTERNAL.
        /// </summary>
        [Input("portName")]
        public Input<string>? PortName { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The protocol this BackendService uses to communicate with backends.
        /// Possible values are HTTP, HTTPS, HTTP2, TCP, and SSL. The default is
        /// HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
        /// types and may result in errors if used with the GA API.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The security policy associated with this backend service.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Type of session affinity to use. The default is NONE. Session affinity is
        /// not applicable if the protocol is UDP.
        /// </summary>
        [Input("sessionAffinity")]
        public Input<string>? SessionAffinity { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// How many seconds to wait for the backend before considering it a
        /// failed request. Default is 30 seconds. Valid range is [1, 86400].
        /// </summary>
        [Input("timeoutSec")]
        public Input<int>? TimeoutSec { get; set; }

        public BackendServiceArgs()
        {
        }
    }

    public sealed class BackendServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Lifetime of cookies in seconds if session_affinity is
        /// GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
        /// only until the end of the browser session (or equivalent). The
        /// maximum allowed value for TTL is one day.
        /// When the load balancing scheme is INTERNAL, this field is not used.
        /// </summary>
        [Input("affinityCookieTtlSec")]
        public Input<int>? AffinityCookieTtlSec { get; set; }

        [Input("backends")]
        private InputList<Inputs.BackendServiceBackendGetArgs>? _backends;

        /// <summary>
        /// -
        /// (Optional)
        /// The set of backends that serve this BackendService.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.BackendServiceBackendGetArgs> Backends
        {
            get => _backends ?? (_backends = new InputList<Inputs.BackendServiceBackendGetArgs>());
            set => _backends = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Cloud CDN configuration for this BackendService.  Structure is documented below.
        /// </summary>
        [Input("cdnPolicy")]
        public Input<Inputs.BackendServiceCdnPolicyGetArgs>? CdnPolicy { get; set; }

        /// <summary>
        /// Settings controlling the volume of connections to a backend service. This field is applicable only when the
        /// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("circuitBreakers")]
        public Input<Inputs.BackendServiceCircuitBreakersGetArgs>? CircuitBreakers { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Time for which instance will be drained (not accept new
        /// connections, but still work to finish started).
        /// </summary>
        [Input("connectionDrainingTimeoutSec")]
        public Input<int>? ConnectionDrainingTimeoutSec { get; set; }

        /// <summary>
        /// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or
        /// other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular
        /// destination host will be lost when one or more hosts are added/removed from the destination service. This field
        /// specifies parameters that control consistent hashing. This field only applies if the load_balancing_scheme is set to
        /// INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is set to MAGLEV or RING_HASH.
        /// </summary>
        [Input("consistentHash")]
        public Input<Inputs.BackendServiceConsistentHashGetArgs>? ConsistentHash { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("customRequestHeaders")]
        private InputList<string>? _customRequestHeaders;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied requests.
        /// </summary>
        public InputList<string> CustomRequestHeaders
        {
            get => _customRequestHeaders ?? (_customRequestHeaders = new InputList<string>());
            set => _customRequestHeaders = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// If true, enable Cloud CDN for this BackendService.
        /// </summary>
        [Input("enableCdn")]
        public Input<bool>? EnableCdn { get; set; }

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource
        /// for health checking this BackendService. Currently at most one health
        /// check can be specified, and a health check is required.
        /// For internal load balancing, a URL to a HealthCheck resource must be specified instead.
        /// </summary>
        [Input("healthChecks")]
        public Input<string>? HealthChecks { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Settings for enabling Cloud Identity Aware Proxy  Structure is documented below.
        /// </summary>
        [Input("iap")]
        public Input<Inputs.BackendServiceIapGetArgs>? Iap { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Indicates whether the backend service will be used with internal or
        /// external load balancing. A backend service created for one type of
        /// load balancing cannot be used with the other. Must be `EXTERNAL` or
        /// `INTERNAL_SELF_MANAGED` for a global backend service. Defaults to `EXTERNAL`.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        /// <summary>
        /// The load balancing algorithm used within the scope of the locality. The possible values are - ROUND_ROBIN - This is a
        /// simple policy in which each healthy backend is selected in round robin order. LEAST_REQUEST - An O(1) algorithm which
        /// selects two random healthy hosts and picks the host which has fewer active requests. RING_HASH - The ring/modulo hash
        /// load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a
        /// host from a set of N hosts only affects 1/N of the requests. RANDOM - The load balancer selects a random healthy host.
        /// ORIGINAL_DESTINATION - Backend host is selected based on the client connection metadata, i.e., connections are opened to
        /// the same address as the destination address of the incoming connection before the connection was redirected to the load
        /// balancer. MAGLEV - used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash
        /// but has faster table lookup build times and host selection times. For more information about Maglev, refer to
        /// https://ai.google/research/pubs/pub44824 This field is applicable only when the load_balancing_scheme is set to
        /// INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("localityLbPolicy")]
        public Input<string>? LocalityLbPolicy { get; set; }

        /// <summary>
        /// This field denotes the logging options for the load balancer traffic served by this backend service. If logging is
        /// enabled, logs will be exported to Stackdriver.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.BackendServiceLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Name of the cookie.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Settings controlling eviction of unhealthy hosts from the load balancing pool. This field is applicable only when the
        /// load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Input("outlierDetection")]
        public Input<Inputs.BackendServiceOutlierDetectionGetArgs>? OutlierDetection { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Name of backend port. The same name should appear in the instance
        /// groups referenced by this service. Required when the load balancing
        /// scheme is EXTERNAL.
        /// </summary>
        [Input("portName")]
        public Input<string>? PortName { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The protocol this BackendService uses to communicate with backends.
        /// Possible values are HTTP, HTTPS, HTTP2, TCP, and SSL. The default is
        /// HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
        /// types and may result in errors if used with the GA API.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The security policy associated with this backend service.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Type of session affinity to use. The default is NONE. Session affinity is
        /// not applicable if the protocol is UDP.
        /// </summary>
        [Input("sessionAffinity")]
        public Input<string>? SessionAffinity { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// How many seconds to wait for the backend before considering it a
        /// failed request. Default is 30 seconds. Valid range is [1, 86400].
        /// </summary>
        [Input("timeoutSec")]
        public Input<int>? TimeoutSec { get; set; }

        public BackendServiceState()
        {
        }
    }
}
