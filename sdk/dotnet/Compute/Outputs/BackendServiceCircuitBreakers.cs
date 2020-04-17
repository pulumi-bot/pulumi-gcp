// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class BackendServiceCircuitBreakers
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The timeout for new network connections to hosts.  Structure is documented below.
        /// </summary>
        public readonly Outputs.BackendServiceCircuitBreakersConnectTimeout? ConnectTimeout;
        /// <summary>
        /// -
        /// (Optional)
        /// The max number of simultaneous connections for the group. Can
        /// be used with either CONNECTION or UTILIZATION balancing modes.
        /// For CONNECTION mode, either maxConnections or one
        /// of maxConnectionsPerInstance or maxConnectionsPerEndpoint,
        /// as appropriate for group type, must be set.
        /// </summary>
        public readonly int? MaxConnections;
        /// <summary>
        /// -
        /// (Optional)
        /// The maximum number of pending requests to the backend cluster.
        /// Defaults to 1024.
        /// </summary>
        public readonly int? MaxPendingRequests;
        /// <summary>
        /// -
        /// (Optional)
        /// The maximum number of parallel requests to the backend cluster.
        /// Defaults to 1024.
        /// </summary>
        public readonly int? MaxRequests;
        /// <summary>
        /// -
        /// (Optional)
        /// Maximum requests for a single backend connection. This parameter
        /// is respected by both the HTTP/1.1 and HTTP/2 implementations. If
        /// not specified, there is no limit. Setting this parameter to 1
        /// will effectively disable keep alive.
        /// </summary>
        public readonly int? MaxRequestsPerConnection;
        /// <summary>
        /// -
        /// (Optional)
        /// The maximum number of parallel retries to the backend cluster.
        /// Defaults to 3.
        /// </summary>
        public readonly int? MaxRetries;

        [OutputConstructor]
        private BackendServiceCircuitBreakers(
            Outputs.BackendServiceCircuitBreakersConnectTimeout? connectTimeout,

            int? maxConnections,

            int? maxPendingRequests,

            int? maxRequests,

            int? maxRequestsPerConnection,

            int? maxRetries)
        {
            ConnectTimeout = connectTimeout;
            MaxConnections = maxConnections;
            MaxPendingRequests = maxPendingRequests;
            MaxRequests = maxRequests;
            MaxRequestsPerConnection = maxRequestsPerConnection;
            MaxRetries = maxRetries;
        }
    }
}
