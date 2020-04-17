// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class StandardAppVersionDeploymentFileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Name of the library. Example "django".
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// SHA1 checksum of the file
        /// </summary>
        [Input("sha1Sum")]
        public Input<string>? Sha1Sum { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// Source URL
        /// </summary>
        [Input("sourceUrl", required: true)]
        public Input<string> SourceUrl { get; set; } = null!;

        public StandardAppVersionDeploymentFileArgs()
        {
        }
    }
}
