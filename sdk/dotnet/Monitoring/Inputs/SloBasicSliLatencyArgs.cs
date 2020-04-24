// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Inputs
{

    public sealed class SloBasicSliLatencyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A duration string, e.g. 10s.
        /// Good service is defined to be the count of requests made to
        /// this service that return in no more than threshold.
        /// </summary>
        [Input("threshold", required: true)]
        public Input<string> Threshold { get; set; } = null!;

        public SloBasicSliLatencyArgs()
        {
        }
    }
}
