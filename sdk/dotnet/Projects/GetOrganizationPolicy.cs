// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Projects
{
    public static partial class Invokes
    {
        /// <summary>
        /// Allows management of Organization policies for a Google Project. For more information see
        /// [the official
        /// documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_project_organization_policy.html.markdown.
        /// </summary>
        [Obsolete("Use GetOrganizationPolicy.InvokeAsync() instead")]
        public static Task<GetOrganizationPolicyResult> GetOrganizationPolicy(GetOrganizationPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationPolicyResult>("gcp:projects/getOrganizationPolicy:getOrganizationPolicy", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetOrganizationPolicy
    {
        /// <summary>
        /// Allows management of Organization policies for a Google Project. For more information see
        /// [the official
        /// documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_project_organization_policy.html.markdown.
        /// </summary>
        public static Task<GetOrganizationPolicyResult> InvokeAsync(GetOrganizationPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationPolicyResult>("gcp:projects/getOrganizationPolicy:getOrganizationPolicy", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetOrganizationPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        /// </summary>
        [Input("constraint", required: true)]
        public string Constraint { get; set; } = null!;

        /// <summary>
        /// The project ID.
        /// </summary>
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetOrganizationPolicyArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetOrganizationPolicyResult
    {
        public readonly ImmutableArray<Outputs.GetOrganizationPolicyBooleanPoliciesResult> BooleanPolicies;
        public readonly string Constraint;
        public readonly string Etag;
        public readonly ImmutableArray<Outputs.GetOrganizationPolicyListPoliciesResult> ListPolicies;
        public readonly string Project;
        public readonly ImmutableArray<Outputs.GetOrganizationPolicyRestorePoliciesResult> RestorePolicies;
        public readonly string UpdateTime;
        public readonly int Version;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetOrganizationPolicyResult(
            ImmutableArray<Outputs.GetOrganizationPolicyBooleanPoliciesResult> booleanPolicies,
            string constraint,
            string etag,
            ImmutableArray<Outputs.GetOrganizationPolicyListPoliciesResult> listPolicies,
            string project,
            ImmutableArray<Outputs.GetOrganizationPolicyRestorePoliciesResult> restorePolicies,
            string updateTime,
            int version,
            string id)
        {
            BooleanPolicies = booleanPolicies;
            Constraint = constraint;
            Etag = etag;
            ListPolicies = listPolicies;
            Project = project;
            RestorePolicies = restorePolicies;
            UpdateTime = updateTime;
            Version = version;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetOrganizationPolicyBooleanPoliciesResult
    {
        public readonly bool Enforced;

        [OutputConstructor]
        private GetOrganizationPolicyBooleanPoliciesResult(bool enforced)
        {
            Enforced = enforced;
        }
    }

    [OutputType]
    public sealed class GetOrganizationPolicyListPoliciesAllowsResult
    {
        public readonly bool All;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetOrganizationPolicyListPoliciesAllowsResult(
            bool all,
            ImmutableArray<string> values)
        {
            All = all;
            Values = values;
        }
    }

    [OutputType]
    public sealed class GetOrganizationPolicyListPoliciesDeniesResult
    {
        public readonly bool All;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetOrganizationPolicyListPoliciesDeniesResult(
            bool all,
            ImmutableArray<string> values)
        {
            All = all;
            Values = values;
        }
    }

    [OutputType]
    public sealed class GetOrganizationPolicyListPoliciesResult
    {
        public readonly ImmutableArray<GetOrganizationPolicyListPoliciesAllowsResult> Allows;
        public readonly ImmutableArray<GetOrganizationPolicyListPoliciesDeniesResult> Denies;
        public readonly bool InheritFromParent;
        public readonly string SuggestedValue;

        [OutputConstructor]
        private GetOrganizationPolicyListPoliciesResult(
            ImmutableArray<GetOrganizationPolicyListPoliciesAllowsResult> allows,
            ImmutableArray<GetOrganizationPolicyListPoliciesDeniesResult> denies,
            bool inheritFromParent,
            string suggestedValue)
        {
            Allows = allows;
            Denies = denies;
            InheritFromParent = inheritFromParent;
            SuggestedValue = suggestedValue;
        }
    }

    [OutputType]
    public sealed class GetOrganizationPolicyRestorePoliciesResult
    {
        public readonly bool Default;

        [OutputConstructor]
        private GetOrganizationPolicyRestorePoliciesResult(bool @default)
        {
            Default = @default;
        }
    }
    }
}
