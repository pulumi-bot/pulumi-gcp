// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BinaryAuthorization
{
    /// <summary>
    /// A policy for container image binary authorization.
    /// 
    /// 
    /// To get more information about Policy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/binary-authorization/)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_policy.html.markdown.
    /// </summary>
    public partial class Policy : Pulumi.CustomResource
    {
        /// <summary>
        /// A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist
        /// pattern, the image's admission requests will always be permitted regardless of your admission rules.
        /// </summary>
        [Output("admissionWhitelistPatterns")]
        public Output<ImmutableArray<Outputs.PolicyAdmissionWhitelistPatterns>> AdmissionWhitelistPatterns { get; private set; } = null!;

        /// <summary>
        /// Per-cluster admission rules. An admission rule specifies either that all container images used in a pod
        /// creation request must be attested to by one or more attestors, that all pod creations will be allowed, or
        /// that all pod creations will be denied. There can be at most one admission rule per cluster spec. Identifier
        /// format: '{{location}}.{{clusterId}}'. A location is either a compute zone (e.g. 'us-central1-a') or a region
        /// (e.g. 'us-central1').
        /// </summary>
        [Output("clusterAdmissionRules")]
        public Output<ImmutableArray<Outputs.PolicyClusterAdmissionRules>> ClusterAdmissionRules { get; private set; } = null!;

        /// <summary>
        /// Default admission rule for a cluster without a per-cluster admission rule.
        /// </summary>
        [Output("defaultAdmissionRule")]
        public Output<Outputs.PolicyDefaultAdmissionRule> DefaultAdmissionRule { get; private set; } = null!;

        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Controls the evaluation of a Google-maintained global admission policy for common system-level images.
        /// Images not covered by the global policy will be subject to the project admission policy.
        /// </summary>
        [Output("globalPolicyEvaluationMode")]
        public Output<string> GlobalPolicyEvaluationMode { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:binaryauthorization/policy:Policy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:binaryauthorization/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : Pulumi.ResourceArgs
    {
        [Input("admissionWhitelistPatterns")]
        private InputList<Inputs.PolicyAdmissionWhitelistPatternsArgs>? _admissionWhitelistPatterns;

        /// <summary>
        /// A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist
        /// pattern, the image's admission requests will always be permitted regardless of your admission rules.
        /// </summary>
        public InputList<Inputs.PolicyAdmissionWhitelistPatternsArgs> AdmissionWhitelistPatterns
        {
            get => _admissionWhitelistPatterns ?? (_admissionWhitelistPatterns = new InputList<Inputs.PolicyAdmissionWhitelistPatternsArgs>());
            set => _admissionWhitelistPatterns = value;
        }

        [Input("clusterAdmissionRules")]
        private InputList<Inputs.PolicyClusterAdmissionRulesArgs>? _clusterAdmissionRules;

        /// <summary>
        /// Per-cluster admission rules. An admission rule specifies either that all container images used in a pod
        /// creation request must be attested to by one or more attestors, that all pod creations will be allowed, or
        /// that all pod creations will be denied. There can be at most one admission rule per cluster spec. Identifier
        /// format: '{{location}}.{{clusterId}}'. A location is either a compute zone (e.g. 'us-central1-a') or a region
        /// (e.g. 'us-central1').
        /// </summary>
        public InputList<Inputs.PolicyClusterAdmissionRulesArgs> ClusterAdmissionRules
        {
            get => _clusterAdmissionRules ?? (_clusterAdmissionRules = new InputList<Inputs.PolicyClusterAdmissionRulesArgs>());
            set => _clusterAdmissionRules = value;
        }

        /// <summary>
        /// Default admission rule for a cluster without a per-cluster admission rule.
        /// </summary>
        [Input("defaultAdmissionRule", required: true)]
        public Input<Inputs.PolicyDefaultAdmissionRuleArgs> DefaultAdmissionRule { get; set; } = null!;

        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Controls the evaluation of a Google-maintained global admission policy for common system-level images.
        /// Images not covered by the global policy will be subject to the project admission policy.
        /// </summary>
        [Input("globalPolicyEvaluationMode")]
        public Input<string>? GlobalPolicyEvaluationMode { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public PolicyArgs()
        {
        }
    }

    public sealed class PolicyState : Pulumi.ResourceArgs
    {
        [Input("admissionWhitelistPatterns")]
        private InputList<Inputs.PolicyAdmissionWhitelistPatternsGetArgs>? _admissionWhitelistPatterns;

        /// <summary>
        /// A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist
        /// pattern, the image's admission requests will always be permitted regardless of your admission rules.
        /// </summary>
        public InputList<Inputs.PolicyAdmissionWhitelistPatternsGetArgs> AdmissionWhitelistPatterns
        {
            get => _admissionWhitelistPatterns ?? (_admissionWhitelistPatterns = new InputList<Inputs.PolicyAdmissionWhitelistPatternsGetArgs>());
            set => _admissionWhitelistPatterns = value;
        }

        [Input("clusterAdmissionRules")]
        private InputList<Inputs.PolicyClusterAdmissionRulesGetArgs>? _clusterAdmissionRules;

        /// <summary>
        /// Per-cluster admission rules. An admission rule specifies either that all container images used in a pod
        /// creation request must be attested to by one or more attestors, that all pod creations will be allowed, or
        /// that all pod creations will be denied. There can be at most one admission rule per cluster spec. Identifier
        /// format: '{{location}}.{{clusterId}}'. A location is either a compute zone (e.g. 'us-central1-a') or a region
        /// (e.g. 'us-central1').
        /// </summary>
        public InputList<Inputs.PolicyClusterAdmissionRulesGetArgs> ClusterAdmissionRules
        {
            get => _clusterAdmissionRules ?? (_clusterAdmissionRules = new InputList<Inputs.PolicyClusterAdmissionRulesGetArgs>());
            set => _clusterAdmissionRules = value;
        }

        /// <summary>
        /// Default admission rule for a cluster without a per-cluster admission rule.
        /// </summary>
        [Input("defaultAdmissionRule")]
        public Input<Inputs.PolicyDefaultAdmissionRuleGetArgs>? DefaultAdmissionRule { get; set; }

        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Controls the evaluation of a Google-maintained global admission policy for common system-level images.
        /// Images not covered by the global policy will be subject to the project admission policy.
        /// </summary>
        [Input("globalPolicyEvaluationMode")]
        public Input<string>? GlobalPolicyEvaluationMode { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public PolicyState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class PolicyAdmissionWhitelistPatternsArgs : Pulumi.ResourceArgs
    {
        [Input("namePattern", required: true)]
        public Input<string> NamePattern { get; set; } = null!;

        public PolicyAdmissionWhitelistPatternsArgs()
        {
        }
    }

    public sealed class PolicyAdmissionWhitelistPatternsGetArgs : Pulumi.ResourceArgs
    {
        [Input("namePattern", required: true)]
        public Input<string> NamePattern { get; set; } = null!;

        public PolicyAdmissionWhitelistPatternsGetArgs()
        {
        }
    }

    public sealed class PolicyClusterAdmissionRulesArgs : Pulumi.ResourceArgs
    {
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        [Input("enforcementMode", required: true)]
        public Input<string> EnforcementMode { get; set; } = null!;

        [Input("evaluationMode", required: true)]
        public Input<string> EvaluationMode { get; set; } = null!;

        [Input("requireAttestationsBies")]
        private InputList<string>? _requireAttestationsBies;
        public InputList<string> RequireAttestationsBies
        {
            get => _requireAttestationsBies ?? (_requireAttestationsBies = new InputList<string>());
            set => _requireAttestationsBies = value;
        }

        public PolicyClusterAdmissionRulesArgs()
        {
        }
    }

    public sealed class PolicyClusterAdmissionRulesGetArgs : Pulumi.ResourceArgs
    {
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        [Input("enforcementMode", required: true)]
        public Input<string> EnforcementMode { get; set; } = null!;

        [Input("evaluationMode", required: true)]
        public Input<string> EvaluationMode { get; set; } = null!;

        [Input("requireAttestationsBies")]
        private InputList<string>? _requireAttestationsBies;
        public InputList<string> RequireAttestationsBies
        {
            get => _requireAttestationsBies ?? (_requireAttestationsBies = new InputList<string>());
            set => _requireAttestationsBies = value;
        }

        public PolicyClusterAdmissionRulesGetArgs()
        {
        }
    }

    public sealed class PolicyDefaultAdmissionRuleArgs : Pulumi.ResourceArgs
    {
        [Input("enforcementMode", required: true)]
        public Input<string> EnforcementMode { get; set; } = null!;

        [Input("evaluationMode", required: true)]
        public Input<string> EvaluationMode { get; set; } = null!;

        [Input("requireAttestationsBies")]
        private InputList<string>? _requireAttestationsBies;
        public InputList<string> RequireAttestationsBies
        {
            get => _requireAttestationsBies ?? (_requireAttestationsBies = new InputList<string>());
            set => _requireAttestationsBies = value;
        }

        public PolicyDefaultAdmissionRuleArgs()
        {
        }
    }

    public sealed class PolicyDefaultAdmissionRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("enforcementMode", required: true)]
        public Input<string> EnforcementMode { get; set; } = null!;

        [Input("evaluationMode", required: true)]
        public Input<string> EvaluationMode { get; set; } = null!;

        [Input("requireAttestationsBies")]
        private InputList<string>? _requireAttestationsBies;
        public InputList<string> RequireAttestationsBies
        {
            get => _requireAttestationsBies ?? (_requireAttestationsBies = new InputList<string>());
            set => _requireAttestationsBies = value;
        }

        public PolicyDefaultAdmissionRuleGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class PolicyAdmissionWhitelistPatterns
    {
        public readonly string NamePattern;

        [OutputConstructor]
        private PolicyAdmissionWhitelistPatterns(string namePattern)
        {
            NamePattern = namePattern;
        }
    }

    [OutputType]
    public sealed class PolicyClusterAdmissionRules
    {
        public readonly string Cluster;
        public readonly string EnforcementMode;
        public readonly string EvaluationMode;
        public readonly ImmutableArray<string> RequireAttestationsBies;

        [OutputConstructor]
        private PolicyClusterAdmissionRules(
            string cluster,
            string enforcementMode,
            string evaluationMode,
            ImmutableArray<string> requireAttestationsBies)
        {
            Cluster = cluster;
            EnforcementMode = enforcementMode;
            EvaluationMode = evaluationMode;
            RequireAttestationsBies = requireAttestationsBies;
        }
    }

    [OutputType]
    public sealed class PolicyDefaultAdmissionRule
    {
        public readonly string EnforcementMode;
        public readonly string EvaluationMode;
        public readonly ImmutableArray<string> RequireAttestationsBies;

        [OutputConstructor]
        private PolicyDefaultAdmissionRule(
            string enforcementMode,
            string evaluationMode,
            ImmutableArray<string> requireAttestationsBies)
        {
            EnforcementMode = enforcementMode;
            EvaluationMode = evaluationMode;
            RequireAttestationsBies = requireAttestationsBies;
        }
    }
    }
}
