// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve Storage Transfer service account for this project
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_transfer_project_service_account.html.markdown.
        /// </summary>
        public static Task<GetTransferProjectServieAccountResult> GetTransferProjectServieAccount(GetTransferProjectServieAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransferProjectServieAccountResult>("gcp:storage/getTransferProjectServieAccount:getTransferProjectServieAccount", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetTransferProjectServieAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project ID. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetTransferProjectServieAccountArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetTransferProjectServieAccountResult
    {
        /// <summary>
        /// Email address of the default service account used by Storage Transfer Jobs running in this project
        /// </summary>
        public readonly string Email;
        public readonly string Project;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetTransferProjectServieAccountResult(
            string email,
            string project,
            string id)
        {
            Email = email;
            Project = project;
            Id = id;
        }
    }
}
