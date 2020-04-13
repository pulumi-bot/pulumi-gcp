// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecurityCenter
{
    /// <summary>
    /// A Cloud Security Command Center's (Cloud SCC) finding source. A finding
    /// source is an entity or a mechanism that can produce a finding. A source is
    /// like a container of findings that come from the same scanner, logger,
    /// monitor, etc.
    /// 
    /// 
    /// To get more information about Source, see:
    /// 
    /// * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1beta1/organizations.sources)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/binary-authorization/)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/scc_source.html.markdown.
    /// </summary>
    public partial class Source : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The description of the source (max of 1024 characters).
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The source’s display name. A source’s display name must be unique
        /// amongst its siblings, for example, two sources with the same parent
        /// can't share the same display name. The display name must start and end
        /// with a letter or digit, may contain letters, digits, spaces, hyphens,
        /// and underscores, and can be no longer than 32 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The organization whose Cloud Security Command Center the Source
        /// lives in.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;


        /// <summary>
        /// Create a Source resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Source(string name, SourceArgs args, CustomResourceOptions? options = null)
            : base("gcp:securitycenter/source:Source", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Source(string name, Input<string> id, SourceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:securitycenter/source:Source", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Source resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Source Get(string name, Input<string> id, SourceState? state = null, CustomResourceOptions? options = null)
        {
            return new Source(name, id, state, options);
        }
    }

    public sealed class SourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The description of the source (max of 1024 characters).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The source’s display name. A source’s display name must be unique
        /// amongst its siblings, for example, two sources with the same parent
        /// can't share the same display name. The display name must start and end
        /// with a letter or digit, may contain letters, digits, spaces, hyphens,
        /// and underscores, and can be no longer than 32 characters.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The organization whose Cloud Security Command Center the Source
        /// lives in.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        public SourceArgs()
        {
        }
    }

    public sealed class SourceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The description of the source (max of 1024 characters).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The source’s display name. A source’s display name must be unique
        /// amongst its siblings, for example, two sources with the same parent
        /// can't share the same display name. The display name must start and end
        /// with a letter or digit, may contain letters, digits, spaces, hyphens,
        /// and underscores, and can be no longer than 32 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The resource name of this source, in the format 'organizations/{{organization}}/sources/{{source}}'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The organization whose Cloud Security Command Center the Source
        /// lives in.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        public SourceState()
        {
        }
    }
}
