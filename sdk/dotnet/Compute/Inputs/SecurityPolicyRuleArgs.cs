// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SecurityPolicyRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take when `match` matches the request. Valid values:
        /// * "allow" : allow access to target
        /// * "deny(status)" : deny access to target, returns the  HTTP response code specified (valid values are 403, 404 and 502)
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// An optional description of this security policy. Max size is 2048.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A match condition that incoming traffic is evaluated against.
        /// If it evaluates to true, the corresponding `action` is enforced. Structure is documented below.
        /// </summary>
        [Input("match", required: true)]
        public Input<Inputs.SecurityPolicyRuleMatchArgs> Match { get; set; } = null!;

        /// <summary>
        /// When set to true, the `action` specified above is not enforced.
        /// Stackdriver logs for requests that trigger a preview action are annotated as such.
        /// </summary>
        [Input("preview")]
        public Input<bool>? Preview { get; set; }

        /// <summary>
        /// An unique positive integer indicating the priority of evaluation for a rule.
        /// Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        public SecurityPolicyRuleArgs()
        {
        }
    }
}
