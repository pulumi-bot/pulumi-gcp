// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetResourcePolicy
    {
        /// <summary>
        /// Provide access to a Resource Policy's attributes. For more information see [the official documentation](https://cloud.google.com/compute/docs/disks/scheduled-snapshots) or the [API](https://cloud.google.com/compute/docs/reference/rest/beta/resourcePolicies).
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var daily = Output.Create(Gcp.Compute.GetResourcePolicy.InvokeAsync(new Gcp.Compute.GetResourcePolicyArgs
        ///         {
        ///             Name = "daily",
        ///             Region = "us-central1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        public static Task<GetResourcePolicyResult> InvokeAsync(GetResourcePolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcePolicyResult>("gcp:compute/getResourcePolicy:getResourcePolicy", args ?? new GetResourcePolicyArgs(), options.WithVersion());
    }


    public sealed class GetResourcePolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Resource Policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Project from which to list the Resource Policy. Defaults to project declared in the provider.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// Region where the Resource Policy resides.
        /// </summary>
        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetResourcePolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourcePolicyResult
    {
        /// <summary>
        /// Description of this Resource Policy.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? Project;
        public readonly string Region;
        /// <summary>
        /// The URI of the resource.
        /// </summary>
        public readonly string SelfLink;

        [OutputConstructor]
        private GetResourcePolicyResult(
            string description,

            string id,

            string name,

            string? project,

            string region,

            string selfLink)
        {
            Description = description;
            Id = id;
            Name = name;
            Project = project;
            Region = region;
            SelfLink = selfLink;
        }
    }
}
