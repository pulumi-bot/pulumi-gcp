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
    public sealed class AutoscalerAutoscalingPolicyMetric
    {
        public readonly string? Filter;
        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. The name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        public readonly double? SingleInstanceAssignment;
        /// <summary>
        /// -
        /// (Required)
        /// URL of the managed instance group that this autoscaler will scale.
        /// </summary>
        public readonly double? Target;
        /// <summary>
        /// -
        /// (Optional)
        /// Defines how target utilization value is expressed for a
        /// Stackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND,
        /// or DELTA_PER_MINUTE.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private AutoscalerAutoscalingPolicyMetric(
            string? filter,

            string name,

            double? singleInstanceAssignment,

            double? target,

            string? type)
        {
            Filter = filter;
            Name = name;
            SingleInstanceAssignment = singleInstanceAssignment;
            Target = target;
            Type = type;
        }
    }
}
