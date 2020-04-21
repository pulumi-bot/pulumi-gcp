// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionBackendServiceConsistentHashArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Hash is based on HTTP Cookie. This field describes a HTTP cookie
        /// that will be used as the hash key for the consistent hash load
        /// balancer. If the cookie is not present, it will be generated.
        /// This field is applicable if the sessionAffinity is set to HTTP_COOKIE.  Structure is documented below.
        /// </summary>
        [Input("httpCookie")]
        public Input<Inputs.RegionBackendServiceConsistentHashHttpCookieArgs>? HttpCookie { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The hash based on the value of the specified header field.
        /// This field is applicable if the sessionAffinity is set to HEADER_FIELD.
        /// </summary>
        [Input("httpHeaderName")]
        public Input<string>? HttpHeaderName { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The minimum number of virtual nodes to use for the hash ring.
        /// Larger ring sizes result in more granular load
        /// distributions. If the number of hosts in the load balancing pool
        /// is larger than the ring size, each host will be assigned a single
        /// virtual node.
        /// Defaults to 1024.
        /// </summary>
        [Input("minimumRingSize")]
        public Input<int>? MinimumRingSize { get; set; }

        public RegionBackendServiceConsistentHashArgs()
        {
        }
    }
}
