// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class AutoscalerAutoscalingPolicyMetricGetArgs : Pulumi.ResourceArgs
    {
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The identifier (type) of the Stackdriver Monitoring metric.
        /// The metric cannot have negative values.
        /// The metric must have a value type of INT64 or DOUBLE.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("singleInstanceAssignment")]
        public Input<double>? SingleInstanceAssignment { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// Fraction of backend capacity utilization (set in HTTP(s) load
        /// balancing configuration) that autoscaler should maintain. Must
        /// be a positive float value. If not defined, the default is 0.8.
        /// </summary>
        [Input("target")]
        public Input<double>? Target { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Defines how target utilization value is expressed for a
        /// Stackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND,
        /// or DELTA_PER_MINUTE.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AutoscalerAutoscalingPolicyMetricGetArgs()
        {
        }
    }
}
