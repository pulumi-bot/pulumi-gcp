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
    public sealed class URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderAction
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Headers to add to a matching request prior to forwarding the request to the
        /// backendService.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionRequestHeadersToAdd> RequestHeadersToAdds;
        /// <summary>
        /// -
        /// (Optional)
        /// A list of header names for headers that need to be removed from the request
        /// prior to forwarding the request to the backendService.
        /// </summary>
        public readonly ImmutableArray<string> RequestHeadersToRemoves;
        /// <summary>
        /// -
        /// (Optional)
        /// Headers to add the response prior to sending the response back to the client.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAdd> ResponseHeadersToAdds;
        /// <summary>
        /// -
        /// (Optional)
        /// A list of header names for headers that need to be removed from the response
        /// prior to sending the response back to the client.
        /// </summary>
        public readonly ImmutableArray<string> ResponseHeadersToRemoves;

        [OutputConstructor]
        private URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderAction(
            ImmutableArray<Outputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionRequestHeadersToAdd> requestHeadersToAdds,

            ImmutableArray<string> requestHeadersToRemoves,

            ImmutableArray<Outputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAdd> responseHeadersToAdds,

            ImmutableArray<string> responseHeadersToRemoves)
        {
            RequestHeadersToAdds = requestHeadersToAdds;
            RequestHeadersToRemoves = requestHeadersToRemoves;
            ResponseHeadersToAdds = responseHeadersToAdds;
            ResponseHeadersToRemoves = responseHeadersToRemoves;
        }
    }
}
