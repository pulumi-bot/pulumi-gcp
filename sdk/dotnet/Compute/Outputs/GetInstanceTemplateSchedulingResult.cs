// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class GetInstanceTemplateSchedulingResult
    {
        /// <summary>
        /// Specifies whether the instance should be
        /// automatically restarted if it is terminated by Compute Engine (not
        /// terminated by a user). This defaults to true.
        /// </summary>
        public readonly bool AutomaticRestart;
        public readonly int MinNodeCpus;
        /// <summary>
        /// Specifies node affinities or anti-affinities
        /// to determine which sole-tenant nodes your instances and managed instance
        /// groups will use as host systems. Read more on sole-tenant node creation
        /// [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
        /// Structure documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTemplateSchedulingNodeAffinityResult> NodeAffinities;
        /// <summary>
        /// Defines the maintenance behavior for this
        /// instance.
        /// </summary>
        public readonly string OnHostMaintenance;
        /// <summary>
        /// Allows instance to be preempted. This defaults to
        /// false. Read more on this
        /// [here](https://cloud.google.com/compute/docs/instances/preemptible).
        /// </summary>
        public readonly bool Preemptible;

        [OutputConstructor]
        private GetInstanceTemplateSchedulingResult(
            bool automaticRestart,

            int minNodeCpus,

            ImmutableArray<Outputs.GetInstanceTemplateSchedulingNodeAffinityResult> nodeAffinities,

            string onHostMaintenance,

            bool preemptible)
        {
            AutomaticRestart = automaticRestart;
            MinNodeCpus = minNodeCpus;
            NodeAffinities = nodeAffinities;
            OnHostMaintenance = onHostMaintenance;
            Preemptible = preemptible;
        }
    }
}
