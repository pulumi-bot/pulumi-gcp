// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class StandardAppVersionEntrypointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The format should be a shell command that can be fed to bash -c.
        /// </summary>
        [Input("shell", required: true)]
        public Input<string> Shell { get; set; } = null!;

        public StandardAppVersionEntrypointArgs()
        {
        }
    }
}
