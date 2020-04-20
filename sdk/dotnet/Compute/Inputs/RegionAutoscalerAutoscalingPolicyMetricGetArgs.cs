// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionAutoscalerAutoscalingPolicyMetricGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// A filter string to be used as the filter string for
        /// a Stackdriver Monitoring TimeSeries.list API call.
        /// This filter is used to select a specific TimeSeries for
        /// the purpose of autoscaling and to determine whether the metric
        /// is exporting per-instance or per-group data.
        /// You can only use the AND operator for joining selectors.
        /// You can only use direct equality comparison operator (=) without
        /// any functions for each selector.
        /// You can specify the metric in both the filter string and in the
        /// metric field. However, if specified in both places, the metric must
        /// be identical.
        /// The monitored resource type determines what kind of values are
        /// expected for the metric. If it is a gce_instance, the autoscaler
        /// expects the metric to include a separate TimeSeries for each
        /// instance in a group. In such a case, you cannot filter on resource
        /// labels.
        /// If the resource type is any other value, the autoscaler expects
        /// this metric to contain values that apply to the entire autoscaled
        /// instance group and resource label filtering can be performed to
        /// point autoscaler at the correct TimeSeries to scale upon.
        /// This is called a per-group metric for the purpose of autoscaling.
        /// If not specified, the type defaults to gce_instance.
        /// You should provide a filter that is selective enough to pick just
        /// one TimeSeries for the autoscaled group or for each of the instances
        /// (if you are using gce_instance resource type). If multiple
        /// TimeSeries are returned upon the query execution, the autoscaler
        /// will sum their respective values to obtain its scaling value.
        /// </summary>
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

        /// <summary>
        /// -
        /// (Optional)
        /// If scaling is based on a per-group metric value that represents the
        /// total amount of work to be done or resource usage, set this value to
        /// an amount assigned for a single instance of the scaled group.
        /// The autoscaler will keep the number of instances proportional to the
        /// value of this metric, the metric itself should not change value due
        /// to group resizing.
        /// For example, a good metric to use with the target is
        /// `pubsub.googleapis.com/subscription/num_undelivered_messages`
        /// or a custom metric exporting the total number of requests coming to
        /// your instances.
        /// A bad example would be a metric exporting an average or median
        /// latency, since this value can't include a chunk assignable to a
        /// single instance, it could be better used with utilization_target
        /// instead.
        /// </summary>
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

        public RegionAutoscalerAutoscalingPolicyMetricGetArgs()
        {
        }
    }
}
