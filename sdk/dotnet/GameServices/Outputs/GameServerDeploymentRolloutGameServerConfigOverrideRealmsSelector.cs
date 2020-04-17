// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices.Outputs
{

    [OutputType]
    public sealed class GameServerDeploymentRolloutGameServerConfigOverrideRealmsSelector
    {
        /// <summary>
        /// -
        /// (Optional)
        /// List of realms to match against.
        /// </summary>
        public readonly ImmutableArray<string> Realms;

        [OutputConstructor]
        private GameServerDeploymentRolloutGameServerConfigOverrideRealmsSelector(ImmutableArray<string> realms)
        {
            Realms = realms;
        }
    }
}
