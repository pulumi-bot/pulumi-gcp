// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get project details.
        /// For more information see
        /// [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects#Project)
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/project.html.markdown.
        /// </summary>
        public static Task<GetProjectResult> GetProject(GetProjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("gcp:organizations/getProject:getProject", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetProjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project ID. If it is not provided, the provider project is used.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetProjectArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetProjectResult
    {
        public readonly ImmutableArray<Outputs.GetProjectAppEnginesResult> AppEngines;
        public readonly bool AutoCreateNetwork;
        public readonly string BillingAccount;
        public readonly string FolderId;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly string Number;
        public readonly string OrgId;
        public readonly string PolicyData;
        public readonly string PolicyEtag;
        public readonly string? ProjectId;
        public readonly bool SkipDelete;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetProjectResult(
            ImmutableArray<Outputs.GetProjectAppEnginesResult> appEngines,
            bool autoCreateNetwork,
            string billingAccount,
            string folderId,
            ImmutableDictionary<string, string> labels,
            string name,
            string number,
            string orgId,
            string policyData,
            string policyEtag,
            string? projectId,
            bool skipDelete,
            string id)
        {
            AppEngines = appEngines;
            AutoCreateNetwork = autoCreateNetwork;
            BillingAccount = billingAccount;
            FolderId = folderId;
            Labels = labels;
            Name = name;
            Number = number;
            OrgId = orgId;
            PolicyData = policyData;
            PolicyEtag = policyEtag;
            ProjectId = projectId;
            SkipDelete = skipDelete;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetProjectAppEnginesFeatureSettingsResult
    {
        public readonly bool SplitHealthChecks;

        [OutputConstructor]
        private GetProjectAppEnginesFeatureSettingsResult(bool splitHealthChecks)
        {
            SplitHealthChecks = splitHealthChecks;
        }
    }

    [OutputType]
    public sealed class GetProjectAppEnginesResult
    {
        public readonly string AuthDomain;
        public readonly string CodeBucket;
        public readonly string DefaultBucket;
        public readonly string DefaultHostname;
        public readonly ImmutableArray<GetProjectAppEnginesFeatureSettingsResult> FeatureSettings;
        public readonly string GcrDomain;
        public readonly string LocationId;
        public readonly string Name;
        public readonly string ServingStatus;
        public readonly ImmutableArray<GetProjectAppEnginesUrlDispatchRulesResult> UrlDispatchRules;

        [OutputConstructor]
        private GetProjectAppEnginesResult(
            string authDomain,
            string codeBucket,
            string defaultBucket,
            string defaultHostname,
            ImmutableArray<GetProjectAppEnginesFeatureSettingsResult> featureSettings,
            string gcrDomain,
            string locationId,
            string name,
            string servingStatus,
            ImmutableArray<GetProjectAppEnginesUrlDispatchRulesResult> urlDispatchRules)
        {
            AuthDomain = authDomain;
            CodeBucket = codeBucket;
            DefaultBucket = defaultBucket;
            DefaultHostname = defaultHostname;
            FeatureSettings = featureSettings;
            GcrDomain = gcrDomain;
            LocationId = locationId;
            Name = name;
            ServingStatus = servingStatus;
            UrlDispatchRules = urlDispatchRules;
        }
    }

    [OutputType]
    public sealed class GetProjectAppEnginesUrlDispatchRulesResult
    {
        public readonly string Domain;
        public readonly string Path;
        public readonly string Service;

        [OutputConstructor]
        private GetProjectAppEnginesUrlDispatchRulesResult(
            string domain,
            string path,
            string service)
        {
            Domain = domain;
            Path = path;
            Service = service;
        }
    }
    }
}
