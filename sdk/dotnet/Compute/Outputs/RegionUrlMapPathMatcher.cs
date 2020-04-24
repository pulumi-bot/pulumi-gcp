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
    public sealed class RegionUrlMapPathMatcher
    {
        /// <summary>
        /// -
        /// (Required)
        /// A reference to a RegionBackendService resource. This will be used if
        /// none of the pathRules defined by this PathMatcher is matched by
        /// the URL's path portion.
        /// </summary>
        public readonly string DefaultService;
        /// <summary>
        /// -
        /// (Optional)
        /// Description of this test case.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// -
        /// (Required)
        /// The name of the query parameter to match. The query parameter must exist in the
        /// request, in the absence of which the request match fails.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// -
        /// (Optional)
        /// The list of path rules. Use this list instead of routeRules when routing based
        /// on simple path matching is all that's required. The order by which path rules
        /// are specified does not matter. Matches are always done on the longest-path-first
        /// basis. For example: a pathRule with a path /a/b/c/* will match before /a/b/*
        /// irrespective of the order in which those paths appear in this list. Within a
        /// given pathMatcher, only one of pathRules or routeRules must be set.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RegionUrlMapPathMatcherPathRule> PathRules;
        /// <summary>
        /// -
        /// (Optional)
        /// The list of ordered HTTP route rules. Use this list instead of pathRules when
        /// advanced route matching and routing actions are desired. The order of specifying
        /// routeRules matters: the first rule that matches will cause its specified routing
        /// action to take effect. Within a given pathMatcher, only one of pathRules or
        /// routeRules must be set. routeRules are not supported in UrlMaps intended for
        /// External load balancers.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RegionUrlMapPathMatcherRouteRule> RouteRules;

        [OutputConstructor]
        private RegionUrlMapPathMatcher(
            string defaultService,

            string? description,

            string name,

            ImmutableArray<Outputs.RegionUrlMapPathMatcherPathRule> pathRules,

            ImmutableArray<Outputs.RegionUrlMapPathMatcherRouteRule> routeRules)
        {
            DefaultService = defaultService;
            Description = description;
            Name = name;
            PathRules = pathRules;
            RouteRules = routeRules;
        }
    }
}
