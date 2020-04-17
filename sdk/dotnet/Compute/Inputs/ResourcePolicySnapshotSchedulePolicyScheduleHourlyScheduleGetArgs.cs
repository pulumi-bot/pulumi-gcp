// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class ResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The number of hours between snapshots.
        /// </summary>
        [Input("hoursInCycle", required: true)]
        public Input<int> HoursInCycle { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// Time within the window to start the operations.
        /// It must be in an hourly format "HH:MM",
        /// where HH : [00-23] and MM : [00] GMT.
        /// eg: 21:00
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public ResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleGetArgs()
        {
        }
    }
}
