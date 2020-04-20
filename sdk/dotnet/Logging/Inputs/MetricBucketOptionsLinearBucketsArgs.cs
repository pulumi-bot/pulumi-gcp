// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging.Inputs
{

    public sealed class MetricBucketOptionsLinearBucketsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Must be greater than 0.
        /// </summary>
        [Input("numFiniteBuckets")]
        public Input<int>? NumFiniteBuckets { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Lower bound of the first bucket.
        /// </summary>
        [Input("offset")]
        public Input<double>? Offset { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Must be greater than 0.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public MetricBucketOptionsLinearBucketsArgs()
        {
        }
    }
}
