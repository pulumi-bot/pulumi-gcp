// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class FlexibleAppVersionDeploymentFileGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// SHA1 checksum of the file
        /// </summary>
        [Input("sha1Sum")]
        public Input<string>? Sha1Sum { get; set; }

        /// <summary>
        /// Source URL
        /// </summary>
        [Input("sourceUrl", required: true)]
        public Input<string> SourceUrl { get; set; } = null!;

        public FlexibleAppVersionDeploymentFileGetArgs()
        {
        }
    }
}
