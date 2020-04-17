// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionEndpointsApiServiceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Endpoints service configuration ID as specified by the Service Management API. For example "2016-09-19r1".
        /// By default, the rollout strategy for Endpoints is "FIXED". This means that Endpoints starts up with a particular configuration ID.
        /// When a new configuration is rolled out, Endpoints must be given the new configuration ID. The configId field is used to give the configuration ID
        /// and is required in this case.
        /// Endpoints also has a rollout strategy called "MANAGED". When using this, Endpoints fetches the latest configuration and does not need
        /// the configuration ID. In this case, configId must be omitted.
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Enable or disable trace sampling. By default, this is set to false for enabled.
        /// </summary>
        [Input("disableTraceSampling")]
        public Input<bool>? DisableTraceSampling { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// Google Compute Engine network where the virtual machines are created. Specify the short name, not the resource path.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Endpoints rollout strategy. If FIXED, configId must be specified. If MANAGED, configId must be omitted. Default is "FIXED".
        /// </summary>
        [Input("rolloutStrategy")]
        public Input<string>? RolloutStrategy { get; set; }

        public FlexibleAppVersionEndpointsApiServiceGetArgs()
        {
        }
    }
}
