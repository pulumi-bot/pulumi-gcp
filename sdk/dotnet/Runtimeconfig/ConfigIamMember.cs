// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Runtimeconfig
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/runtimeconfig_config_iam_member.html.markdown.
    /// </summary>
    public partial class ConfigIamMember : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.ConfigIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigIamMember(string name, ConfigIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:runtimeconfig/configIamMember:ConfigIamMember", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ConfigIamMember(string name, Input<string> id, ConfigIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:runtimeconfig/configIamMember:ConfigIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigIamMember Get(string name, Input<string> id, ConfigIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigIamMember(name, id, state, options);
        }
    }

    public sealed class ConfigIamMemberArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.ConfigIamMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public ConfigIamMemberArgs()
        {
        }
    }

    public sealed class ConfigIamMemberState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.ConfigIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public ConfigIamMemberState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ConfigIamMemberConditionArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public ConfigIamMemberConditionArgs()
        {
        }
    }

    public sealed class ConfigIamMemberConditionGetArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public ConfigIamMemberConditionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ConfigIamMemberCondition
    {
        public readonly string? Description;
        public readonly string Expression;
        public readonly string Title;

        [OutputConstructor]
        private ConfigIamMemberCondition(
            string? description,
            string expression,
            string title)
        {
            Description = description;
            Expression = expression;
            Title = title;
        }
    }
    }
}
