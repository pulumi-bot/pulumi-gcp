// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class URLMapPathMatcherRouteRuleMatchRule
    {
        /// <summary>
        /// -
        /// (Optional)
        /// For satifying the matchRule condition, the path of the request must exactly
        /// match the value specified in fullPathMatch after removing any query parameters
        /// and anchor that may be part of the original URL. FullPathMatch must be between 1
        /// and 1024 characters. Only one of prefixMatch, fullPathMatch or regexMatch must
        /// be specified.
        /// </summary>
        public readonly string? FullPathMatch;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a list of header match criteria, all of which must match corresponding
        /// headers in the request.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.URLMapPathMatcherRouteRuleMatchRuleHeaderMatch> HeaderMatches;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies that prefixMatch and fullPathMatch matches are case sensitive.
        /// Defaults to false.
        /// </summary>
        public readonly bool? IgnoreCase;
        /// <summary>
        /// -
        /// (Optional)
        /// Opaque filter criteria used by Loadbalancer to restrict routing configuration to
        /// a limited set xDS compliant clients. In their xDS requests to Loadbalancer, xDS
        /// clients present node metadata. If a match takes place, the relevant routing
        /// configuration is made available to those proxies. For each metadataFilter in
        /// this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the
        /// filterLabels must match the corresponding label provided in the metadata. If its
        /// filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match
        /// with corresponding labels in the provided metadata. metadataFilters specified
        /// here can be overrides those specified in ForwardingRule that refers to this
        /// UrlMap. metadataFilters only applies to Loadbalancers that have their
        /// loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilter> MetadataFilters;
        /// <summary>
        /// -
        /// (Optional)
        /// For satifying the matchRule condition, the request's path must begin with the
        /// specified prefixMatch. prefixMatch must begin with a /. The value must be
        /// between 1 and 1024 characters. Only one of prefixMatch, fullPathMatch or
        /// regexMatch must be specified.
        /// </summary>
        public readonly string? PrefixMatch;
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a list of query parameter match criteria, all of which must match
        /// corresponding query parameters in the request.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatch> QueryParameterMatches;
        /// <summary>
        /// -
        /// (Optional)
        /// For satifying the matchRule condition, the path of the request must satisfy the
        /// regular expression specified in regexMatch after removing any query parameters
        /// and anchor supplied with the original URL. For regular expression grammar please
        /// see en.cppreference.com/w/cpp/regex/ecmascript  Only one of prefixMatch,
        /// fullPathMatch or regexMatch must be specified.
        /// </summary>
        public readonly string? RegexMatch;

        [OutputConstructor]
        private URLMapPathMatcherRouteRuleMatchRule(
            string? fullPathMatch,

            ImmutableArray<Outputs.URLMapPathMatcherRouteRuleMatchRuleHeaderMatch> headerMatches,

            bool? ignoreCase,

            ImmutableArray<Outputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilter> metadataFilters,

            string? prefixMatch,

            ImmutableArray<Outputs.URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatch> queryParameterMatches,

            string? regexMatch)
        {
            FullPathMatch = fullPathMatch;
            HeaderMatches = headerMatches;
            IgnoreCase = ignoreCase;
            MetadataFilters = metadataFilters;
            PrefixMatch = prefixMatch;
            QueryParameterMatches = queryParameterMatches;
            RegexMatch = regexMatch;
        }
    }
}
