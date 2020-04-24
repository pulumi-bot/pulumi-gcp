// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherPathRuleRouteActionTimeoutArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Span of time that's a fraction of a second at nanosecond resolution. Durations
        /// less than one second are represented with a 0 `seconds` field and a positive
        /// `nanos` field. Must be from 0 to 999,999,999 inclusive.
        /// </summary>
        [Input("nanos")]
        public Input<int>? Nanos { get; set; }

        /// <summary>
        /// Span of time at a resolution of a second. Must be from 0 to 315,576,000,000
        /// inclusive.
        /// </summary>
        [Input("seconds", required: true)]
        public Input<string> Seconds { get; set; } = null!;

        public URLMapPathMatcherPathRuleRouteActionTimeoutArgs()
        {
        }
    }
}
