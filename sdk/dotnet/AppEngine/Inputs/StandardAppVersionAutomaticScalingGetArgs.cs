// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class StandardAppVersionAutomaticScalingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.
        /// Defaults to a runtime-specific value.
        /// </summary>
        [Input("maxConcurrentRequests")]
        public Input<int>? MaxConcurrentRequests { get; set; }

        /// <summary>
        /// Maximum number of idle instances that should be maintained for this version.
        /// </summary>
        [Input("maxIdleInstances")]
        public Input<int>? MaxIdleInstances { get; set; }

        /// <summary>
        /// Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        [Input("maxPendingLatency")]
        public Input<string>? MaxPendingLatency { get; set; }

        /// <summary>
        /// Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.
        /// </summary>
        [Input("minIdleInstances")]
        public Input<int>? MinIdleInstances { get; set; }

        /// <summary>
        /// Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        [Input("minPendingLatency")]
        public Input<string>? MinPendingLatency { get; set; }

        /// <summary>
        /// Scheduler settings for standard environment.  Structure is documented below.
        /// </summary>
        [Input("standardSchedulerSettings")]
        public Input<Inputs.StandardAppVersionAutomaticScalingStandardSchedulerSettingsGetArgs>? StandardSchedulerSettings { get; set; }

        public StandardAppVersionAutomaticScalingGetArgs()
        {
        }
    }
}
