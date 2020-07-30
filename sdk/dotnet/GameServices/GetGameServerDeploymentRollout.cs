// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices
{
    public static class GetGameServerDeploymentRollout
    {
        /// <summary>
        /// Use this data source to get the rollout state. 
        /// 
        /// https://cloud.google.com/game-servers/docs/reference/rest/v1beta/GameServerDeploymentRollout
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var qa = Output.Create(Gcp.GameServices.GetGameServerDeploymentRollout.InvokeAsync(new Gcp.GameServices.GetGameServerDeploymentRolloutArgs
        ///         {
        ///             DeploymentId = "tf-test-deployment-s8sn12jt2c",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGameServerDeploymentRolloutResult> InvokeAsync(GetGameServerDeploymentRolloutArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGameServerDeploymentRolloutResult>("gcp:gameservices/getGameServerDeploymentRollout:getGameServerDeploymentRollout", args ?? new GetGameServerDeploymentRolloutArgs(), options.WithVersion());
    }


    public sealed class GetGameServerDeploymentRolloutArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The deployment to get the rollout state from. Only 1 rollout must be associated with each deployment.
        /// </summary>
        [Input("deploymentId", required: true)]
        public string DeploymentId { get; set; } = null!;

        public GetGameServerDeploymentRolloutArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGameServerDeploymentRolloutResult
    {
        public readonly string DefaultGameServerConfig;
        public readonly string DeploymentId;
        public readonly ImmutableArray<Outputs.GetGameServerDeploymentRolloutGameServerConfigOverrideResult> GameServerConfigOverrides;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private GetGameServerDeploymentRolloutResult(
            string defaultGameServerConfig,

            string deploymentId,

            ImmutableArray<Outputs.GetGameServerDeploymentRolloutGameServerConfigOverrideResult> gameServerConfigOverrides,

            string id,

            string name,

            string project)
        {
            DefaultGameServerConfig = defaultGameServerConfig;
            DeploymentId = deploymentId;
            GameServerConfigOverrides = gameServerConfigOverrides;
            Id = id;
            Name = name;
            Project = project;
        }
    }
}
