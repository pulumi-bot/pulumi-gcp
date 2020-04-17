// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Outputs
{

    [OutputType]
    public sealed class FlexibleAppVersionDeploymentCloudBuildOptions
    {
        /// <summary>
        /// -
        /// (Required)
        /// Path to the yaml file used in deployment, used to determine runtime configuration details.
        /// </summary>
        public readonly string AppYamlPath;
        /// <summary>
        /// -
        /// (Optional)
        /// The Cloud Build timeout used as part of any dependent builds performed by version creation. Defaults to 10 minutes.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        /// </summary>
        public readonly string? CloudBuildTimeout;

        [OutputConstructor]
        private FlexibleAppVersionDeploymentCloudBuildOptions(
            string appYamlPath,

            string? cloudBuildTimeout)
        {
            AppYamlPath = appYamlPath;
            CloudBuildTimeout = cloudBuildTimeout;
        }
    }
}
