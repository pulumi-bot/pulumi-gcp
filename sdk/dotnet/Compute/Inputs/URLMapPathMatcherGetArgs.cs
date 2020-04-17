// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The backend service or backend bucket to use when none of the given rules match.
        /// </summary>
        [Input("defaultService")]
        public Input<string>? DefaultService { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource. Provide this property when you create
        /// the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies changes to request and response headers that need to take effect for
        /// the selected backendService. The headerAction specified here take effect after
        /// headerAction specified under pathMatcher.  Structure is documented below.
        /// </summary>
        [Input("headerAction")]
        public Input<Inputs.URLMapPathMatcherHeaderActionGetArgs>? HeaderAction { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. Provided by the client when the resource is created. The
        /// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
        /// name must be 1-63 characters long and match the regular expression
        /// `a-z?` which means the first character must be a lowercase
        /// letter, and all following characters must be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("pathRules")]
        private InputList<Inputs.URLMapPathMatcherPathRuleGetArgs>? _pathRules;

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
        public InputList<Inputs.URLMapPathMatcherPathRuleGetArgs> PathRules
        {
            get => _pathRules ?? (_pathRules = new InputList<Inputs.URLMapPathMatcherPathRuleGetArgs>());
            set => _pathRules = value;
        }

        [Input("routeRules")]
        private InputList<Inputs.URLMapPathMatcherRouteRuleGetArgs>? _routeRules;

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
        public InputList<Inputs.URLMapPathMatcherRouteRuleGetArgs> RouteRules
        {
            get => _routeRules ?? (_routeRules = new InputList<Inputs.URLMapPathMatcherRouteRuleGetArgs>());
            set => _routeRules = value;
        }

        public URLMapPathMatcherGetArgs()
        {
        }
    }
}
