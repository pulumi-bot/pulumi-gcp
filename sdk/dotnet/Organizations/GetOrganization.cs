// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static class GetOrganization
    {
        /// <summary>
        /// Get information about a Google Cloud Organization. Note that you must have the `roles/resourcemanager.organizationViewer` role (or equivalent permissions) at the organization level to use this datasource.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var org = Output.Create(Gcp.Organizations.GetOrganization.InvokeAsync(new Gcp.Organizations.GetOrganizationArgs
        ///         {
        ///             Domain = "example.com",
        ///         }));
        ///         var sales = new Gcp.Organizations.Folder("sales", new Gcp.Organizations.FolderArgs
        ///         {
        ///             DisplayName = "Sales",
        ///             Parent = org.Apply(org =&gt; org.Name),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        public static Task<GetOrganizationResult> InvokeAsync(GetOrganizationArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationResult>("gcp:organizations/getOrganization:getOrganization", args ?? new GetOrganizationArgs(), options.WithVersion());

        public static Output<GetOrganizationResult> Apply(GetOrganizationApplyArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetOrganizationApplyArgs();
            return Pulumi.Output.All(
                args.Domain.Box(),
                args.Organization.Box()
            ).Apply(a => {
                    var args = new GetOrganizationArgs();
                    a[0].Set(args, nameof(args.Domain));
                    a[1].Set(args, nameof(args.Organization));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetOrganizationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain name of the Organization.
        /// </summary>
        [Input("domain")]
        public string? Domain { get; set; }

        /// <summary>
        /// The Organization's numeric ID, including an optional `organizations/` prefix.
        /// </summary>
        [Input("organization")]
        public string? Organization { get; set; }

        public GetOrganizationArgs()
        {
        }
    }

    public sealed class GetOrganizationApplyArgs
    {
        /// <summary>
        /// The domain name of the Organization.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The Organization's numeric ID, including an optional `organizations/` prefix.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        public GetOrganizationApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrganizationResult
    {
        /// <summary>
        /// Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The Google for Work customer ID of the Organization.
        /// </summary>
        public readonly string DirectoryCustomerId;
        public readonly string Domain;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Organization's current lifecycle state.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// The resource name of the Organization in the form `organizations/{organization_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Organization ID.
        /// </summary>
        public readonly string OrgId;
        public readonly string? Organization;

        [OutputConstructor]
        private GetOrganizationResult(
            string createTime,

            string directoryCustomerId,

            string domain,

            string id,

            string lifecycleState,

            string name,

            string orgId,

            string? organization)
        {
            CreateTime = createTime;
            DirectoryCustomerId = directoryCustomerId;
            Domain = domain;
            Id = id;
            LifecycleState = lifecycleState;
            Name = name;
            OrgId = orgId;
            Organization = organization;
        }
    }
}
