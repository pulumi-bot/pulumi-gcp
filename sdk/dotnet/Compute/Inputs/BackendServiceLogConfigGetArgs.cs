// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class BackendServiceLogConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Whether to enable logging for the load balancer traffic served by this backend service.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// This field can only be specified if logging is enabled for this backend service. The value of
        /// the field must be in [0, 1]. This configures the sampling rate of requests to the load balancer
        /// where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported.
        /// The default value is 1.0.
        /// </summary>
        [Input("sampleRate")]
        public Input<double>? SampleRate { get; set; }

        public BackendServiceLogConfigGetArgs()
        {
        }
    }
}
