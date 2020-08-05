// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iam
{
    public static class GetRule
    {
        /// <summary>
        /// Use this data source to get information about a Google IAM Role.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var roleinfo = Output.Create(Gcp.Iam.GetRule.InvokeAsync(new Gcp.Iam.GetRuleArgs
        ///         {
        ///             Name = "roles/compute.viewer",
        ///         }));
        ///         this.TheRolePermissions = roleinfo.Apply(roleinfo =&gt; roleinfo.IncludedPermissions);
        ///     }
        /// 
        ///     [Output("theRolePermissions")]
        ///     public Output&lt;string&gt; TheRolePermissions { get; set; }
        /// }
        /// ```
        /// </summary>
        public static Task<GetRuleResult> InvokeAsync(GetRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRuleResult>("gcp:iam/getRule:getRule", args ?? new GetRuleArgs(), options.WithVersion());
    }


    public sealed class GetRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Role to lookup in the form `roles/{ROLE_NAME}`, `organizations/{ORGANIZATION_ID}/roles/{ROLE_NAME}` or `projects/{PROJECT_ID}/roles/{ROLE_NAME}`
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// specifies the list of one or more permissions to include in the custom role, such as - `iam.roles.get`
        /// </summary>
        public readonly ImmutableArray<string> IncludedPermissions;
        public readonly string Name;
        /// <summary>
        /// indicates the stage of a role in the launch lifecycle, such as `GA`, `BETA` or `ALPHA`.
        /// </summary>
        public readonly string Stage;
        /// <summary>
        /// is a friendly title for the role, such as "Role Viewer"
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private GetRuleResult(
            string id,

            ImmutableArray<string> includedPermissions,

            string name,

            string stage,

            string title)
        {
            Id = id;
            IncludedPermissions = includedPermissions;
            Name = name;
            Stage = stage;
            Title = title;
        }
    }
}
