// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// A Security Policy defines an IP blacklist or whitelist that protects load balanced Google Cloud services by denying or permitting traffic from specified IP ranges. For more information
    /// see the [official documentation](https://cloud.google.com/armor/docs/configure-security-policies)
    /// and the [API](https://cloud.google.com/compute/docs/reference/rest/beta/securityPolicies).
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var policy = new Gcp.Compute.SecurityPolicy("policy", new Gcp.Compute.SecurityPolicyArgs
    ///         {
    ///             Rules = 
    ///             {
    ///                 new Gcp.Compute.Inputs.SecurityPolicyRuleArgs
    ///                 {
    ///                     Action = "deny(403)",
    ///                     Description = "Deny access to IPs in 9.9.9.0/24",
    ///                     Match = new Gcp.Compute.Inputs.SecurityPolicyRuleMatchArgs
    ///                     {
    ///                         Config = new Gcp.Compute.Inputs.SecurityPolicyRuleMatchConfigArgs
    ///                         {
    ///                             SrcIpRanges = 
    ///                             {
    ///                                 "9.9.9.0/24",
    ///                             },
    ///                         },
    ///                         VersionedExpr = "SRC_IPS_V1",
    ///                     },
    ///                     Priority = 1000,
    ///                 },
    ///                 new Gcp.Compute.Inputs.SecurityPolicyRuleArgs
    ///                 {
    ///                     Action = "allow",
    ///                     Description = "default rule",
    ///                     Match = new Gcp.Compute.Inputs.SecurityPolicyRuleMatchArgs
    ///                     {
    ///                         Config = new Gcp.Compute.Inputs.SecurityPolicyRuleMatchConfigArgs
    ///                         {
    ///                             SrcIpRanges = 
    ///                             {
    ///                                 "*",
    ///                             },
    ///                         },
    ///                         VersionedExpr = "SRC_IPS_V1",
    ///                     },
    ///                     Priority = 2147483647,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
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
        public Output<ImmutableArray<Outputs.SecurityPolicyRule>> Rules { get; private set; } = null!;

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
            : base("gcp:compute/securityPolicy:SecurityPolicy", name, args ?? new SecurityPolicyArgs(), MakeResourceOptions(options, ""))
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
        private InputList<Inputs.SecurityPolicyRuleArgs>? _rules;

        /// <summary>
        /// The set of rules that belong to this policy. There must always be a default
        /// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
        /// security policy, a default rule with action "allow" will be added. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecurityPolicyRuleArgs>());
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
        private InputList<Inputs.SecurityPolicyRuleGetArgs>? _rules;

        /// <summary>
        /// The set of rules that belong to this policy. There must always be a default
        /// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
        /// security policy, a default rule with action "allow" will be added. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecurityPolicyRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecurityPolicyRuleGetArgs>());
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
}
