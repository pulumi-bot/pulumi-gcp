// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class BackendServiceCircuitBreakersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timeout for new network connections to hosts.  Structure is documented below.
        /// </summary>
        [Input("connectTimeout")]
        public Input<Inputs.BackendServiceCircuitBreakersConnectTimeoutArgs>? ConnectTimeout { get; set; }

        /// <summary>
        /// The maximum number of connections to the backend cluster.
        /// Defaults to 1024.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// The maximum number of pending requests to the backend cluster.
        /// Defaults to 1024.
        /// </summary>
        [Input("maxPendingRequests")]
        public Input<int>? MaxPendingRequests { get; set; }

        /// <summary>
        /// The maximum number of parallel requests to the backend cluster.
        /// Defaults to 1024.
        /// </summary>
        [Input("maxRequests")]
        public Input<int>? MaxRequests { get; set; }

        /// <summary>
        /// Maximum requests for a single backend connection. This parameter
        /// is respected by both the HTTP/1.1 and HTTP/2 implementations. If
        /// not specified, there is no limit. Setting this parameter to 1
        /// will effectively disable keep alive.
        /// </summary>
        [Input("maxRequestsPerConnection")]
        public Input<int>? MaxRequestsPerConnection { get; set; }

        /// <summary>
        /// The maximum number of parallel retries to the backend cluster.
        /// Defaults to 3.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        public BackendServiceCircuitBreakersArgs()
        {
        }
    }
}
