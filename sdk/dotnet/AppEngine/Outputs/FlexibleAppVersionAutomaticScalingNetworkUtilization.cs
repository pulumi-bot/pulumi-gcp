// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Outputs
{

    [OutputType]
    public sealed class FlexibleAppVersionAutomaticScalingNetworkUtilization
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Target bytes received per second.
        /// </summary>
        public readonly int? TargetReceivedBytesPerSecond;
        /// <summary>
        /// -
        /// (Optional)
        /// Target packets received per second.
        /// </summary>
        public readonly int? TargetReceivedPacketsPerSecond;
        /// <summary>
        /// -
        /// (Optional)
        /// Target bytes sent per second.
        /// </summary>
        public readonly int? TargetSentBytesPerSecond;
        /// <summary>
        /// -
        /// (Optional)
        /// Target packets sent per second.
        /// </summary>
        public readonly int? TargetSentPacketsPerSecond;

        [OutputConstructor]
        private FlexibleAppVersionAutomaticScalingNetworkUtilization(
            int? targetReceivedBytesPerSecond,

            int? targetReceivedPacketsPerSecond,

            int? targetSentBytesPerSecond,

            int? targetSentPacketsPerSecond)
        {
            TargetReceivedBytesPerSecond = targetReceivedBytesPerSecond;
            TargetReceivedPacketsPerSecond = targetReceivedPacketsPerSecond;
            TargetSentBytesPerSecond = targetSentBytesPerSecond;
            TargetSentPacketsPerSecond = targetSentPacketsPerSecond;
        }
    }
}
