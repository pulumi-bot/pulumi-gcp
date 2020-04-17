// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class ClusterNetworkPolicy
    {
        /// <summary>
        /// Whether node auto-provisioning is enabled. Resource
        /// limits for `cpu` and `memory` must be defined to enable node auto-provisioning.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED.
        /// </summary>
        public readonly string? Provider;

        [OutputConstructor]
        private ClusterNetworkPolicy(
            bool enabled,

            string? provider)
        {
            Enabled = enabled;
            Provider = provider;
        }
    }
}
