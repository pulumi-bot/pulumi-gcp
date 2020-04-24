// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default RegionBackendService resource. Before
        /// forwarding the request to backendService, the loadbalancer applies any relevant
        /// headerActions specified as part of this backendServiceWeight.
        /// </summary>
        [Input("backendService", required: true)]
        public Input<string> BackendService { get; set; } = null!;

        public RegionUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyGetArgs()
        {
        }
    }
}
