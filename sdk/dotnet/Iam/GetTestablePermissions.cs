// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iam
{
    public static class GetTestablePermissions
    {
        /// <summary>
        /// Retrieve a list of testable permissions for a resource. Testable permissions mean the permissions that user can add or remove in a role at a given resource. The resource can be referenced either via the full resource name or via a URI.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Retrieve all the supported permissions able to be set on `my-project` that are in either GA or BETA. This is useful for dynamically constructing custom roles.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var perms = Output.Create(Gcp.Iam.GetTestablePermissions.InvokeAsync(new Gcp.Iam.GetTestablePermissionsArgs
        ///         {
        ///             FullResourceName = "//cloudresourcemanager.googleapis.com/projects/my-project",
        ///             Stages = 
        ///             {
        ///                 "GA",
        ///                 "BETA",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTestablePermissionsResult> InvokeAsync(GetTestablePermissionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTestablePermissionsResult>("gcp:iam/getTestablePermissions:getTestablePermissions", args ?? new GetTestablePermissionsArgs(), options.WithVersion());
    }


    public sealed class GetTestablePermissionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The level of support for custom roles. Can be one of `"NOT_SUPPORTED"`, `"SUPPORTED"`, `"TESTING"`. Default is `"SUPPORTED"`
        /// </summary>
        [Input("customSupportLevel")]
        public string? CustomSupportLevel { get; set; }

        /// <summary>
        /// See [full resource name documentation](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more detail.
        /// </summary>
        [Input("fullResourceName", required: true)]
        public string FullResourceName { get; set; } = null!;

        [Input("stages")]
        private List<string>? _stages;

        /// <summary>
        /// The acceptable release stages of the permission in the output. Note that `BETA` does not include permissions in `GA`, but you can specify both with `["GA", "BETA"]` for example. Can be a list of `"ALPHA"`, `"BETA"`, `"GA"`, `"DEPRECATED"`. Default is `["GA"]`.
        /// </summary>
        public List<string> Stages
        {
            get => _stages ?? (_stages = new List<string>());
            set => _stages = value;
        }

        public GetTestablePermissionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTestablePermissionsResult
    {
        /// <summary>
        /// The the support level of this permission for custom roles.
        /// </summary>
        public readonly string? CustomSupportLevel;
        public readonly string FullResourceName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of permissions matching the provided input. Structure is defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTestablePermissionsPermissionResult> Permissions;
        public readonly ImmutableArray<string> Stages;

        [OutputConstructor]
        private GetTestablePermissionsResult(
            string? customSupportLevel,

            string fullResourceName,

            string id,

            ImmutableArray<Outputs.GetTestablePermissionsPermissionResult> permissions,

            ImmutableArray<string> stages)
        {
            CustomSupportLevel = customSupportLevel;
            FullResourceName = fullResourceName;
            Id = id;
            Permissions = permissions;
            Stages = stages;
        }
    }
}
