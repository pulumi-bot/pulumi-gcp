// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Inputs
{

    public sealed class UptimeCheckConfigHttpCheckGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The authentication information. Optional when creating an HTTP check; defaults to empty.  Structure is documented below.
        /// </summary>
        [Input("authInfo")]
        public Input<Inputs.UptimeCheckConfigHttpCheckAuthInfoGetArgs>? AuthInfo { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;

        /// <summary>
        /// -
        /// (Optional)
        /// The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with ******.
        /// </summary>
        [Input("maskHeaders")]
        public Input<bool>? MaskHeaders { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. Optional (defaults to "/").
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) and path to construct the full URL. Optional (defaults to 80 without SSL, or 443 with SSL).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// If true, use HTTPS instead of HTTP to run the check.
        /// </summary>
        [Input("useSsl")]
        public Input<bool>? UseSsl { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitoredResource is set to uptime_url. If useSsl is false, setting validateSsl to true has no effect.
        /// </summary>
        [Input("validateSsl")]
        public Input<bool>? ValidateSsl { get; set; }

        public UptimeCheckConfigHttpCheckGetArgs()
        {
        }
    }
}
