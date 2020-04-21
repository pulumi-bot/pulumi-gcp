// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DeploymentManager.Inputs
{

    public sealed class DeploymentTargetImportGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The full contents of the template that you want to import.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The name of the template to import, as declared in the YAML
        /// configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DeploymentTargetImportGetArgs()
        {
        }
    }
}
