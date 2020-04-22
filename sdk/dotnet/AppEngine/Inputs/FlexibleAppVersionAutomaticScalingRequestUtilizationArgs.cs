// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionAutomaticScalingRequestUtilizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Target number of concurrent requests.
        /// </summary>
        [Input("targetConcurrentRequests")]
        public Input<double>? TargetConcurrentRequests { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Target requests per second.
        /// </summary>
        [Input("targetRequestCountPerSecond")]
        public Input<string>? TargetRequestCountPerSecond { get; set; }

        public FlexibleAppVersionAutomaticScalingRequestUtilizationArgs()
        {
        }
    }
}
