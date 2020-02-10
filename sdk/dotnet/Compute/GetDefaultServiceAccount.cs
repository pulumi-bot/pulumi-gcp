// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve default service account for this project
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_default_service_account.html.markdown.
        /// </summary>
        public static Task<GetDefaultServiceAccountResult> GetDefaultServiceAccount(GetDefaultServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefaultServiceAccountResult>("gcp:compute/getDefaultServiceAccount:getDefaultServiceAccount", args ?? InvokeArgs.Empty, options.WithVersion());
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
        /// The fully-qualified name of the service account.
        /// </summary>
        public readonly string Name;
        public readonly string Project;
        /// <summary>
        /// The unique id of the service account.
        /// </summary>
        public readonly string UniqueId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDefaultServiceAccountResult(
            string displayName,
            string email,
            string name,
            string project,
            string uniqueId,
            string id)
        {
            DisplayName = displayName;
            Email = email;
            Name = name;
            Project = project;
            UniqueId = uniqueId;
            Id = id;
        }
    }
}
