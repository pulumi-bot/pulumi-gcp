// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class AutoscalerAutoscalingPolicyCpuUtilizationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fraction of backend capacity utilization (set in HTTP(s) load
        /// balancing configuration) that autoscaler should maintain. Must
        /// be a positive float value. If not defined, the default is 0.8.
        /// </summary>
        [Input("target", required: true)]
        public Input<double> Target { get; set; } = null!;

        public AutoscalerAutoscalingPolicyCpuUtilizationGetArgs()
        {
        }
    }
}
