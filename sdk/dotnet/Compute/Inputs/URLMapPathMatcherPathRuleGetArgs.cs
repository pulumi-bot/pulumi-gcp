// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherPathRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("paths", required: true)]
        private InputList<string>? _paths;

        /// <summary>
        /// -
        /// (Required)
        /// The list of path patterns to match. Each must start with / and the only place a
        /// * is allowed is at the end following a /. The string fed to the path matcher
        /// does not include any text after the first ? or #, and those chars are not
        /// allowed here.
        /// </summary>
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// In response to a matching path, the load balancer performs advanced routing
        /// actions like URL rewrites, header transformations, etc. prior to forwarding the
        /// request to the selected backend. If routeAction specifies any
        /// weightedBackendServices, service must not be set. Conversely if service is set,
        /// routeAction cannot contain any  weightedBackendServices. Only one of routeAction
        /// or urlRedirect must be set.  Structure is documented below.
        /// </summary>
        [Input("routeAction")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionGetArgs>? RouteAction { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The backend service or backend bucket to use if any of the given paths match.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// When a path pattern is matched, the request is redirected to a URL specified by
        /// urlRedirect. If urlRedirect is specified, service or routeAction must not be
        /// set.  Structure is documented below.
        /// </summary>
        [Input("urlRedirect")]
        public Input<Inputs.URLMapPathMatcherPathRuleUrlRedirectGetArgs>? UrlRedirect { get; set; }

        public URLMapPathMatcherPathRuleGetArgs()
        {
        }
    }
}
