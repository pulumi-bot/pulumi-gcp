// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceTemplateSpecContainerEnvFromSecretRefArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The ConfigMap to select from.  Structure is documented below.
        /// </summary>
        [Input("localObjectReference")]
        public Input<Inputs.ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReferenceArgs>? LocalObjectReference { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specify whether the ConfigMap must be defined
        /// </summary>
        [Input("optional")]
        public Input<bool>? Optional { get; set; }

        public ServiceTemplateSpecContainerEnvFromSecretRefArgs()
        {
        }
    }
}
