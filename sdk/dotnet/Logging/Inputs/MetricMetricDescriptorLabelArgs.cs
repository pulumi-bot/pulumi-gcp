// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging.Inputs
{

    public sealed class MetricMetricDescriptorLabelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// A human-readable description for the label.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The label key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// Whether the measurement is an integer, a floating-point number, etc.
        /// Some combinations of metricKind and valueType might not be supported.
        /// For counter metrics, set this to INT64.
        /// </summary>
        [Input("valueType")]
        public Input<string>? ValueType { get; set; }

        public MetricMetricDescriptorLabelArgs()
        {
        }
    }
}
