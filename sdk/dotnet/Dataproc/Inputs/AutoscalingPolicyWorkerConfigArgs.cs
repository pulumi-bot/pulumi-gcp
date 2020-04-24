// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class AutoscalingPolicyWorkerConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of instances for this group. Note that by default, clusters will not use
        /// secondary workers. Required for secondary workers if the minimum secondary instances is set.
        /// Bounds: [minInstances, ). Defaults to 0.
        /// </summary>
        [Input("maxInstances", required: true)]
        public Input<int> MaxInstances { get; set; } = null!;

        /// <summary>
        /// Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
        /// </summary>
        [Input("minInstances")]
        public Input<int>? MinInstances { get; set; }

        /// <summary>
        /// Weight for the instance group, which is used to determine the fraction of total workers
        /// in the cluster from this instance group. For example, if primary workers have weight 2,
        /// and secondary workers have weight 1, the cluster will have approximately 2 primary workers
        /// for each secondary worker.
        /// The cluster may not reach the specified balance if constrained by min/max bounds or other
        /// autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
        /// primary workers will be added. The cluster can also be out of balance when created.
        /// If weight is not set on any instance group, the cluster will default to equal weight for
        /// all groups: the cluster will attempt to maintain an equal number of workers in each group
        /// within the configured size bounds for each group. If weight is set for one group only,
        /// the cluster will default to zero weight on the unset group. For example if weight is set
        /// only on primary workers, the cluster will use primary workers only and no secondary workers.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public AutoscalingPolicyWorkerConfigArgs()
        {
        }
    }
}
