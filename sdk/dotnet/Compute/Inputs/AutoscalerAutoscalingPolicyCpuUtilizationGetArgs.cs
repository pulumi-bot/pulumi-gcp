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
        /// -
        /// (Required)
        /// URL of the managed instance group that this autoscaler will scale.
        /// </summary>
        [Input("target", required: true)]
        public Input<double> Target { get; set; } = null!;

        public AutoscalerAutoscalingPolicyCpuUtilizationGetArgs()
        {
        }
    }
}
