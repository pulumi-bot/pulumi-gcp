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
    public sealed class FlexibleAppVersionDeploymentFile
    {
        /// <summary>
        /// -
        /// (Required)
        /// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// -
        /// (Optional)
        /// SHA1 checksum of the file
        /// </summary>
        public readonly string? Sha1Sum;
        /// <summary>
        /// -
        /// (Required)
        /// Source URL
        /// </summary>
        public readonly string SourceUrl;

        [OutputConstructor]
        private FlexibleAppVersionDeploymentFile(
            string name,

            string? sha1Sum,

            string sourceUrl)
        {
            Name = name;
            Sha1Sum = sha1Sum;
            SourceUrl = sourceUrl;
        }
    }
}
