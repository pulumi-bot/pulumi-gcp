// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging.Inputs
{

    public sealed class MetricBucketOptionsExplicitBucketsArgs : Pulumi.ResourceArgs
    {
        [Input("bounds", required: true)]
        private InputList<double>? _bounds;

        /// <summary>
        /// -
        /// (Required)
        /// The values must be monotonically increasing.
        /// </summary>
        public InputList<double> Bounds
        {
            get => _bounds ?? (_bounds = new InputList<double>());
            set => _bounds = value;
        }

        public MetricBucketOptionsExplicitBucketsArgs()
        {
        }
    }
}
