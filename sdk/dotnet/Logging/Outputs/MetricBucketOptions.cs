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
    public sealed class MetricBucketOptions
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a set of buckets with arbitrary widths.  Structure is documented below.
        /// </summary>
        public readonly Outputs.MetricBucketOptionsExplicitBuckets? ExplicitBuckets;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies an exponential sequence of buckets that have a width that is proportional to the value of
        /// the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.  Structure is documented below.
        /// </summary>
        public readonly Outputs.MetricBucketOptionsExponentialBuckets? ExponentialBuckets;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
        /// Each bucket represents a constant absolute uncertainty on the specific value in the bucket.  Structure is documented below.
        /// </summary>
        public readonly Outputs.MetricBucketOptionsLinearBuckets? LinearBuckets;

        [OutputConstructor]
        private MetricBucketOptions(
            Outputs.MetricBucketOptionsExplicitBuckets? explicitBuckets,

            Outputs.MetricBucketOptionsExponentialBuckets? exponentialBuckets,

            Outputs.MetricBucketOptionsLinearBuckets? linearBuckets)
        {
            ExplicitBuckets = explicitBuckets;
            ExponentialBuckets = exponentialBuckets;
            LinearBuckets = linearBuckets;
        }
    }
}
