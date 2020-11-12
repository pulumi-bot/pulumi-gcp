// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine
{
    /// <summary>
    /// A single firewall rule that is evaluated against incoming traffic
    /// and provides an action to take on matched requests.
    /// 
    /// To get more information about FirewallRule, see:
    /// 
    /// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/appengine/docs/standard/python/creating-firewalls#creating_firewall_rules)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// FirewallRule can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/firewallRule:FirewallRule default apps/{{project}}/firewall/ingressRules/{{priority}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/firewallRule:FirewallRule default {{project}}/{{priority}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/firewallRule:FirewallRule default {{priority}}
    /// ```
    /// </summary>
    public partial class FirewallRule : Pulumi.CustomResource
    {
        /// <summary>
        /// The action to take if this rule matches.
        /// Possible values are `UNSPECIFIED_ACTION`, `ALLOW`, and `DENY`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// An optional string description of this rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A positive integer that defines the order of rule evaluation.
        /// Rules with the lowest priority are evaluated first.
        /// A default rule at priority Int32.MaxValue matches all IPv4 and
        /// IPv6 traffic when no previous rule matches. Only the action of
        /// this rule can be modified by the user.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// IP address or range, defined using CIDR notation, of requests that this rule applies to.
        /// </summary>
        [Output("sourceRange")]
        public Output<string> SourceRange { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallRule(string name, FirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("gcp:appengine/firewallRule:FirewallRule", name, args ?? new FirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallRule(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
            : base("gcp:appengine/firewallRule:FirewallRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallRule Get(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallRule(name, id, state, options);
        }
    }

    public sealed class FirewallRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take if this rule matches.
        /// Possible values are `UNSPECIFIED_ACTION`, `ALLOW`, and `DENY`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// An optional string description of this rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A positive integer that defines the order of rule evaluation.
        /// Rules with the lowest priority are evaluated first.
        /// A default rule at priority Int32.MaxValue matches all IPv4 and
        /// IPv6 traffic when no previous rule matches. Only the action of
        /// this rule can be modified by the user.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// IP address or range, defined using CIDR notation, of requests that this rule applies to.
        /// </summary>
        [Input("sourceRange", required: true)]
        public Input<string> SourceRange { get; set; } = null!;

        public FirewallRuleArgs()
        {
        }
    }

    public sealed class FirewallRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take if this rule matches.
        /// Possible values are `UNSPECIFIED_ACTION`, `ALLOW`, and `DENY`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// An optional string description of this rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A positive integer that defines the order of rule evaluation.
        /// Rules with the lowest priority are evaluated first.
        /// A default rule at priority Int32.MaxValue matches all IPv4 and
        /// IPv6 traffic when no previous rule matches. Only the action of
        /// this rule can be modified by the user.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// IP address or range, defined using CIDR notation, of requests that this rule applies to.
        /// </summary>
        [Input("sourceRange")]
        public Input<string>? SourceRange { get; set; }

        public FirewallRuleState()
        {
        }
    }
}
