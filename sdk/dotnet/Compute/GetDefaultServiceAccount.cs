// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetDefaultServiceAccount
    {
        /// <summary>
        /// Use this data source to retrieve default service account for this project
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
        ///         var @default = Output.Create(Gcp.Compute.GetDefaultServiceAccount.InvokeAsync());
        ///         this.DefaultAccount = @default.Apply(@default =&gt; _default.Email);
        ///     }
        /// 
        ///     [Output("defaultAccount")]
        ///     public Output&lt;string&gt; DefaultAccount { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDefaultServiceAccountResult> InvokeAsync(GetDefaultServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefaultServiceAccountResult>("gcp:compute/getDefaultServiceAccount:getDefaultServiceAccount", args ?? new GetDefaultServiceAccountArgs(), options.WithVersion());
    }


    public sealed class GetDefaultServiceAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project ID. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetDefaultServiceAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDefaultServiceAccountResult
    {
        /// <summary>
        /// The display name for the service account.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Email address of the default service account used by VMs running in this project
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The fully-qualified name of the service account.
        /// </summary>
        public readonly string Name;
        public readonly string Project;
        /// <summary>
        /// The unique id of the service account.
        /// </summary>
        public readonly string UniqueId;

        [OutputConstructor]
        private GetDefaultServiceAccountResult(
            string displayName,

            string email,

            string id,

            string name,

            string project,

            string uniqueId)
        {
            DisplayName = displayName;
            Email = email;
            Id = id;
            Name = name;
            Project = project;
            UniqueId = uniqueId;
        }
    }
}
