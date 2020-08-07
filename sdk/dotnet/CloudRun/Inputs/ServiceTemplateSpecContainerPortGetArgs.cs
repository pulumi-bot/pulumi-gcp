// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceTemplateSpecContainerPortGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Port number.
        /// </summary>
        [Input("containerPort", required: true)]
        public Input<int> ContainerPort { get; set; } = null!;

        /// <summary>
        /// Name of the port.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Protocol used on port. Defaults to TCP.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public ServiceTemplateSpecContainerPortGetArgs()
        {
        }
    }
}