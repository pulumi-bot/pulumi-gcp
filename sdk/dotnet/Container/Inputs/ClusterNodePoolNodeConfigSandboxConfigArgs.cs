// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class ClusterNodePoolNodeConfigSandboxConfigArgs : Pulumi.ResourceArgs
    {
        [Input("sandboxType", required: true)]
        public Input<string> SandboxType { get; set; } = null!;

        public ClusterNodePoolNodeConfigSandboxConfigArgs()
        {
        }
    }
}
