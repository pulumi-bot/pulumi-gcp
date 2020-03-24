// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public partial class SecurityPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// An optional description of this rule. Max size is 64.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The name of the security policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The set of rules that belong to this policy. There must always be a default
        /// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
        /// security policy, a default rule with action "allow" will be added. Structure is documented below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.SecurityPolicyRules>> Rules { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityPolicy(string name, SecurityPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/securityPolicy:SecurityPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecurityPolicy(string name, Input<string> id, SecurityPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/securityPolicy:SecurityPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityPolicy Get(string name, Input<string> id, SecurityPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityPolicy(name, id, state, options);
        }
    }

    public sealed class SecurityPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this rule. Max size is 64.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the security policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("rules")]
        private InputList<Inputs.SecurityPolicyRulesArgs>? _rules;

        /// <summary>
        /// The set of rules that belong to this policy. There must always be a default
        /// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
        /// security policy, a default rule with action "allow" will be added. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRulesArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecurityPolicyRulesArgs>());
            set => _rules = value;
        }

        public SecurityPolicyArgs()
        {
        }
    }

    public sealed class SecurityPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this rule. Max size is 64.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Fingerprint of this resource.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The name of the security policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("rules")]
        private InputList<Inputs.SecurityPolicyRulesGetArgs>? _rules;

        /// <summary>
        /// The set of rules that belong to this policy. There must always be a default
        /// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
        /// security policy, a default rule with action "allow" will be added. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRulesGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecurityPolicyRulesGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public SecurityPolicyState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SecurityPolicyRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take when `match` matches the request. Valid values:
        /// * "allow" : allow access to target
        /// * "deny(status)" : deny access to target, returns the  HTTP response code specified (valid values are 403, 404 and 502)
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// An optional description of this rule. Max size is 64.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A match condition that incoming traffic is evaluated against.
        /// If it evaluates to true, the corresponding `action` is enforced. Structure is documented below.
        /// </summary>
        [Input("match", required: true)]
        public Input<SecurityPolicyRulesMatchArgs> Match { get; set; } = null!;

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

        public SecurityPolicyRulesArgs()
        {
        }
    }

    public sealed class SecurityPolicyRulesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take when `match` matches the request. Valid values:
        /// * "allow" : allow access to target
        /// * "deny(status)" : deny access to target, returns the  HTTP response code specified (valid values are 403, 404 and 502)
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// An optional description of this rule. Max size is 64.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A match condition that incoming traffic is evaluated against.
        /// If it evaluates to true, the corresponding `action` is enforced. Structure is documented below.
        /// </summary>
        [Input("match", required: true)]
        public Input<SecurityPolicyRulesMatchGetArgs> Match { get; set; } = null!;

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

        public SecurityPolicyRulesGetArgs()
        {
        }
    }

    public sealed class SecurityPolicyRulesMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration options available when specifying `versioned_expr`.
        /// This field must be specified if `versioned_expr` is specified and cannot be specified if `versioned_expr` is not specified.
        /// Structure is documented below.
        /// </summary>
        [Input("config")]
        public Input<SecurityPolicyRulesMatchConfigArgs>? Config { get; set; }

        /// <summary>
        /// User defined CEVAL expression. A CEVAL expression is used to specify match criteria
        /// such as origin.ip, source.region_code and contents in the request header.
        /// Structure is documented below.
        /// </summary>
        [Input("expr")]
        public Input<SecurityPolicyRulesMatchExprArgs>? Expr { get; set; }

        /// <summary>
        /// Predefined rule expression. If this field is specified, `config` must also be specified.
        /// Available options:
        /// * SRC_IPS_V1: Must specify the corresponding `src_ip_ranges` field in `config`.
        /// </summary>
        [Input("versionedExpr")]
        public Input<string>? VersionedExpr { get; set; }

        public SecurityPolicyRulesMatchArgs()
        {
        }
    }

    public sealed class SecurityPolicyRulesMatchConfigArgs : Pulumi.ResourceArgs
    {
        [Input("srcIpRanges", required: true)]
        private InputList<string>? _srcIpRanges;

        /// <summary>
        /// Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation
        /// to match against inbound traffic. There is a limit of 5 IP ranges per rule. A value of '\*' matches all IPs
        /// (can be used to override the default behavior).
        /// </summary>
        public InputList<string> SrcIpRanges
        {
            get => _srcIpRanges ?? (_srcIpRanges = new InputList<string>());
            set => _srcIpRanges = value;
        }

        public SecurityPolicyRulesMatchConfigArgs()
        {
        }
    }

    public sealed class SecurityPolicyRulesMatchConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("srcIpRanges", required: true)]
        private InputList<string>? _srcIpRanges;

        /// <summary>
        /// Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation
        /// to match against inbound traffic. There is a limit of 5 IP ranges per rule. A value of '\*' matches all IPs
        /// (can be used to override the default behavior).
        /// </summary>
        public InputList<string> SrcIpRanges
        {
            get => _srcIpRanges ?? (_srcIpRanges = new InputList<string>());
            set => _srcIpRanges = value;
        }

        public SecurityPolicyRulesMatchConfigGetArgs()
        {
        }
    }

    public sealed class SecurityPolicyRulesMatchExprArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Textual representation of an expression in Common Expression Language syntax.
        /// The application context of the containing message determines which well-known feature set of CEL is supported.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        public SecurityPolicyRulesMatchExprArgs()
        {
        }
    }

    public sealed class SecurityPolicyRulesMatchExprGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Textual representation of an expression in Common Expression Language syntax.
        /// The application context of the containing message determines which well-known feature set of CEL is supported.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        public SecurityPolicyRulesMatchExprGetArgs()
        {
        }
    }

    public sealed class SecurityPolicyRulesMatchGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration options available when specifying `versioned_expr`.
        /// This field must be specified if `versioned_expr` is specified and cannot be specified if `versioned_expr` is not specified.
        /// Structure is documented below.
        /// </summary>
        [Input("config")]
        public Input<SecurityPolicyRulesMatchConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// User defined CEVAL expression. A CEVAL expression is used to specify match criteria
        /// such as origin.ip, source.region_code and contents in the request header.
        /// Structure is documented below.
        /// </summary>
        [Input("expr")]
        public Input<SecurityPolicyRulesMatchExprGetArgs>? Expr { get; set; }

        /// <summary>
        /// Predefined rule expression. If this field is specified, `config` must also be specified.
        /// Available options:
        /// * SRC_IPS_V1: Must specify the corresponding `src_ip_ranges` field in `config`.
        /// </summary>
        [Input("versionedExpr")]
        public Input<string>? VersionedExpr { get; set; }

        public SecurityPolicyRulesMatchGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SecurityPolicyRules
    {
        /// <summary>
        /// Action to take when `match` matches the request. Valid values:
        /// * "allow" : allow access to target
        /// * "deny(status)" : deny access to target, returns the  HTTP response code specified (valid values are 403, 404 and 502)
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// An optional description of this rule. Max size is 64.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A match condition that incoming traffic is evaluated against.
        /// If it evaluates to true, the corresponding `action` is enforced. Structure is documented below.
        /// </summary>
        public readonly SecurityPolicyRulesMatch Match;
        /// <summary>
        /// When set to true, the `action` specified above is not enforced.
        /// Stackdriver logs for requests that trigger a preview action are annotated as such.
        /// </summary>
        public readonly bool? Preview;
        /// <summary>
        /// An unique positive integer indicating the priority of evaluation for a rule.
        /// Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
        /// </summary>
        public readonly int Priority;

        [OutputConstructor]
        private SecurityPolicyRules(
            string action,
            string? description,
            SecurityPolicyRulesMatch match,
            bool? preview,
            int priority)
        {
            Action = action;
            Description = description;
            Match = match;
            Preview = preview;
            Priority = priority;
        }
    }

    [OutputType]
    public sealed class SecurityPolicyRulesMatch
    {
        /// <summary>
        /// The configuration options available when specifying `versioned_expr`.
        /// This field must be specified if `versioned_expr` is specified and cannot be specified if `versioned_expr` is not specified.
        /// Structure is documented below.
        /// </summary>
        public readonly SecurityPolicyRulesMatchConfig? Config;
        /// <summary>
        /// User defined CEVAL expression. A CEVAL expression is used to specify match criteria
        /// such as origin.ip, source.region_code and contents in the request header.
        /// Structure is documented below.
        /// </summary>
        public readonly SecurityPolicyRulesMatchExpr? Expr;
        /// <summary>
        /// Predefined rule expression. If this field is specified, `config` must also be specified.
        /// Available options:
        /// * SRC_IPS_V1: Must specify the corresponding `src_ip_ranges` field in `config`.
        /// </summary>
        public readonly string? VersionedExpr;

        [OutputConstructor]
        private SecurityPolicyRulesMatch(
            SecurityPolicyRulesMatchConfig? config,
            SecurityPolicyRulesMatchExpr? expr,
            string? versionedExpr)
        {
            Config = config;
            Expr = expr;
            VersionedExpr = versionedExpr;
        }
    }

    [OutputType]
    public sealed class SecurityPolicyRulesMatchConfig
    {
        /// <summary>
        /// Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation
        /// to match against inbound traffic. There is a limit of 5 IP ranges per rule. A value of '\*' matches all IPs
        /// (can be used to override the default behavior).
        /// </summary>
        public readonly ImmutableArray<string> SrcIpRanges;

        [OutputConstructor]
        private SecurityPolicyRulesMatchConfig(ImmutableArray<string> srcIpRanges)
        {
            SrcIpRanges = srcIpRanges;
        }
    }

    [OutputType]
    public sealed class SecurityPolicyRulesMatchExpr
    {
        /// <summary>
        /// Textual representation of an expression in Common Expression Language syntax.
        /// The application context of the containing message determines which well-known feature set of CEL is supported.
        /// </summary>
        public readonly string Expression;

        [OutputConstructor]
        private SecurityPolicyRulesMatchExpr(string expression)
        {
            Expression = expression;
        }
    }
    }
}
