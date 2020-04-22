// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging.Outputs
{

    [OutputType]
    public sealed class MetricMetricDescriptorLabel
    {
        /// <summary>
        /// -
        /// (Optional)
        /// A description of this metric, which is used in documentation. The maximum length of the
        /// description is 8000 characters.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// -
        /// (Required)
        /// The label key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// -
        /// (Optional)
        /// The type of data that can be assigned to the label.
        /// </summary>
        public readonly string? ValueType;

        [OutputConstructor]
        private MetricMetricDescriptorLabel(
            string? description,

            string key,

            string? valueType)
        {
            Description = description;
            Key = key;
            ValueType = valueType;
        }
    }
}
