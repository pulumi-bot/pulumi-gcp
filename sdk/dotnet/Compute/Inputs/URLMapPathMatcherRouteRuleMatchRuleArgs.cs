// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherRouteRuleMatchRuleArgs : Pulumi.ResourceArgs
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
        [Input("fullPathMatch")]
        public Input<string>? FullPathMatch { get; set; }

        [Input("headerMatches")]
        private InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs>? _headerMatches;

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a list of header match criteria, all of which must match corresponding
        /// headers in the request.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs> HeaderMatches
        {
            get => _headerMatches ?? (_headerMatches = new InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs>());
            set => _headerMatches = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies that prefixMatch and fullPathMatch matches are case sensitive.
        /// Defaults to false.
        /// </summary>
        [Input("ignoreCase")]
        public Input<bool>? IgnoreCase { get; set; }

        [Input("metadataFilters")]
        private InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs>? _metadataFilters;

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
        public InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs> MetadataFilters
        {
            get => _metadataFilters ?? (_metadataFilters = new InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs>());
            set => _metadataFilters = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// For satifying the matchRule condition, the request's path must begin with the
        /// specified prefixMatch. prefixMatch must begin with a /. The value must be
        /// between 1 and 1024 characters. Only one of prefixMatch, fullPathMatch or
        /// regexMatch must be specified.
        /// </summary>
        [Input("prefixMatch")]
        public Input<string>? PrefixMatch { get; set; }

        [Input("queryParameterMatches")]
        private InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs>? _queryParameterMatches;

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a list of query parameter match criteria, all of which must match
        /// corresponding query parameters in the request.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs> QueryParameterMatches
        {
            get => _queryParameterMatches ?? (_queryParameterMatches = new InputList<Inputs.URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs>());
            set => _queryParameterMatches = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// For satifying the matchRule condition, the path of the request must satisfy the
        /// regular expression specified in regexMatch after removing any query parameters
        /// and anchor supplied with the original URL. For regular expression grammar please
        /// see en.cppreference.com/w/cpp/regex/ecmascript  Only one of prefixMatch,
        /// fullPathMatch or regexMatch must be specified.
        /// </summary>
        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        public URLMapPathMatcherRouteRuleMatchRuleArgs()
        {
        }
    }
}
