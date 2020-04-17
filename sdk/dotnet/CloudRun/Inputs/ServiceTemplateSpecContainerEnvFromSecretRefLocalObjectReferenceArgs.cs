// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// Name must be unique within a namespace, within a Cloud Run region.
        /// Is required when creating resources. Name is primarily intended
        /// for creation idempotence and configuration definition. Cannot be updated.
        /// More info: http://kubernetes.io/docs/user-guide/identifiers#names
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReferenceArgs()
        {
        }
    }
}
