// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging.Inputs
{

    public sealed class MetricMetricDescriptorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// A concise name for the metric, which can be displayed in user interfaces. Use sentence case
        /// without an ending period, for example "Request count". This field is optional but it is
        /// recommended to be set for any metrics associated with user-visible concepts, such as Quota.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputList<Inputs.MetricMetricDescriptorLabelArgs>? _labels;

        /// <summary>
        /// -
        /// (Optional)
        /// The set of labels that can be used to describe a specific instance of this metric type. For
        /// example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
        /// for the HTTP response code, response_code, so you can look at latencies for successful responses
        /// or just for responses that failed.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.MetricMetricDescriptorLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.MetricMetricDescriptorLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Whether the metric records instantaneous values, changes to a value, etc.
        /// Some combinations of metricKind and valueType might not be supported.
        /// For counter metrics, set this to DELTA.
        /// </summary>
        [Input("metricKind", required: true)]
        public Input<string> MetricKind { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The unit in which the metric value is reported. It is only applicable if the valueType is
        /// `INT64`, `DOUBLE`, or `DISTRIBUTION`. The supported units are a subset of
        /// [The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The type of data that can be assigned to the label.
        /// </summary>
        [Input("valueType", required: true)]
        public Input<string> ValueType { get; set; } = null!;

        public MetricMetricDescriptorArgs()
        {
        }
    }
}
