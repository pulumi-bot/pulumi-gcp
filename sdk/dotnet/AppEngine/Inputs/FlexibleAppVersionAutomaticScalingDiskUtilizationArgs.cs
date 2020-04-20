// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionAutomaticScalingDiskUtilizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Target bytes read per second.
        /// </summary>
        [Input("targetReadBytesPerSecond")]
        public Input<int>? TargetReadBytesPerSecond { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Target ops read per seconds.
        /// </summary>
        [Input("targetReadOpsPerSecond")]
        public Input<int>? TargetReadOpsPerSecond { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Target bytes written per second.
        /// </summary>
        [Input("targetWriteBytesPerSecond")]
        public Input<int>? TargetWriteBytesPerSecond { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Target ops written per second.
        /// </summary>
        [Input("targetWriteOpsPerSecond")]
        public Input<int>? TargetWriteOpsPerSecond { get; set; }

        public FlexibleAppVersionAutomaticScalingDiskUtilizationArgs()
        {
        }
    }
}
