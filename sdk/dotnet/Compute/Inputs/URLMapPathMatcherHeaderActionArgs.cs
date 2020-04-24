// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherHeaderActionArgs : Pulumi.ResourceArgs
    {
        [Input("requestHeadersToAdds")]
        private InputList<Inputs.URLMapPathMatcherHeaderActionRequestHeadersToAddArgs>? _requestHeadersToAdds;

        /// <summary>
        /// -
        /// (Optional)
        /// Headers to add to a matching request prior to forwarding the request to the
        /// backendService.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.URLMapPathMatcherHeaderActionRequestHeadersToAddArgs> RequestHeadersToAdds
        {
            get => _requestHeadersToAdds ?? (_requestHeadersToAdds = new InputList<Inputs.URLMapPathMatcherHeaderActionRequestHeadersToAddArgs>());
            set => _requestHeadersToAdds = value;
        }

        [Input("requestHeadersToRemoves")]
        private InputList<string>? _requestHeadersToRemoves;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of header names for headers that need to be removed from the request
        /// prior to forwarding the request to the backendService.
        /// </summary>
        public InputList<string> RequestHeadersToRemoves
        {
            get => _requestHeadersToRemoves ?? (_requestHeadersToRemoves = new InputList<string>());
            set => _requestHeadersToRemoves = value;
        }

        [Input("responseHeadersToAdds")]
        private InputList<Inputs.URLMapPathMatcherHeaderActionResponseHeadersToAddArgs>? _responseHeadersToAdds;

        /// <summary>
        /// -
        /// (Optional)
        /// Headers to add the response prior to sending the response back to the client.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.URLMapPathMatcherHeaderActionResponseHeadersToAddArgs> ResponseHeadersToAdds
        {
            get => _responseHeadersToAdds ?? (_responseHeadersToAdds = new InputList<Inputs.URLMapPathMatcherHeaderActionResponseHeadersToAddArgs>());
            set => _responseHeadersToAdds = value;
        }

        [Input("responseHeadersToRemoves")]
        private InputList<string>? _responseHeadersToRemoves;

        /// <summary>
        /// -
        /// (Optional)
        /// A list of header names for headers that need to be removed from the response
        /// prior to sending the response back to the client.
        /// </summary>
        public InputList<string> ResponseHeadersToRemoves
        {
            get => _responseHeadersToRemoves ?? (_responseHeadersToRemoves = new InputList<string>());
            set => _responseHeadersToRemoves = value;
        }

        public URLMapPathMatcherHeaderActionArgs()
        {
        }
    }
}
