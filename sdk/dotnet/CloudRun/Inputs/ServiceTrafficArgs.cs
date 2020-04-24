// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceTrafficArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// LatestRevision may be optionally provided to indicate that the latest ready
        /// Revision of the Configuration should be used for this traffic target. When
        /// provided LatestRevision must be true if RevisionName is empty; it must be
        /// false when RevisionName is non-empty.
        /// </summary>
        [Input("latestRevision")]
        public Input<bool>? LatestRevision { get; set; }

        /// <summary>
        /// Percent specifies percent of the traffic to this Revision or Configuration.
        /// </summary>
        [Input("percent", required: true)]
        public Input<int> Percent { get; set; } = null!;

        /// <summary>
        /// RevisionName of a specific revision to which to send this portion of traffic.
        /// </summary>
        [Input("revisionName")]
        public Input<string>? RevisionName { get; set; }

        public ServiceTrafficArgs()
        {
        }
    }
}
