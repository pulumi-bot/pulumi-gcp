// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ServiceAccount
{
    public static class GetAccount
    {
        /// <summary>
        /// Get the service account from a project. For more information see
        /// the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("gcp:serviceAccount/getAccount:getAccount", args ?? new GetAccountArgs(), options.WithVersion());
    }


    public sealed class GetAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Service account id.  (This is the part of the service account's email field that comes before the @ symbol.)
        /// </summary>
        [Input("accountId", required: true)]
        public string AccountId { get; set; } = null!;

        /// <summary>
        /// The ID of the project that the service account is present in.
        /// Defaults to the provider project configuration.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        public readonly string AccountId;
        /// <summary>
        /// The display name for the service account.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The e-mail address of the service account. This value
        /// should be referenced from any `gcp.organizationsgetIAMPolicy` data sources
        /// that would grant the service account privileges.
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
        public readonly string? Project;
        /// <summary>
        /// The unique id of the service account.
        /// </summary>
        public readonly string UniqueId;

        [OutputConstructor]
        private GetAccountResult(
            string accountId,

            string displayName,

            string email,

            string id,

            string name,

            string? project,

            string uniqueId)
        {
            AccountId = accountId;
            DisplayName = displayName;
            Email = email;
            Id = id;
            Name = name;
            Project = project;
            UniqueId = uniqueId;
        }
    }
}
