// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionDeploymentCloudBuildOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// Path to the yaml file used in deployment, used to determine runtime configuration details.
        /// </summary>
        [Input("appYamlPath", required: true)]
        public Input<string> AppYamlPath { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The Cloud Build timeout used as part of any dependent builds performed by version creation. Defaults to 10 minutes.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        [Input("cloudBuildTimeout")]
        public Input<string>? CloudBuildTimeout { get; set; }

        public FlexibleAppVersionDeploymentCloudBuildOptionsArgs()
        {
        }
    }
}
