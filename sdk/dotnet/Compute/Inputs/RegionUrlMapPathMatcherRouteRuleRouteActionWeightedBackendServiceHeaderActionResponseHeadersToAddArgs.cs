// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAddArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The name of the header.
        /// </summary>
        [Input("headerName", required: true)]
        public Input<string> HeaderName { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The value of the header to add.
        /// </summary>
        [Input("headerValue", required: true)]
        public Input<string> HeaderValue { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// If false, headerValue is appended to any values that already exist for the
        /// header. If true, headerValue is set for the header, discarding any values that
        /// were set for that header.
        /// </summary>
        [Input("replace", required: true)]
        public Input<bool> Replace { get; set; } = null!;

        public RegionUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAddArgs()
        {
        }
    }
}
