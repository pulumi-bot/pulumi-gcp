// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Authoritatively manages metadata common to all instances for a project in GCE. For more information see
    /// [the official documentation](https://cloud.google.com/compute/docs/storing-retrieving-metadata)
    /// and
    /// [API](https://cloud.google.com/compute/docs/reference/latest/projects/setCommonInstanceMetadata).
    /// 
    /// &gt; **Note:**  This resource manages all project-level metadata including project-level ssh keys.
    /// Keys unset in config but set on the server will be removed. If you want to manage only single
    /// key/value pairs within the project metadata rather than the entire set, then use
    /// google_compute_project_metadata_item.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_project_metadata.html.markdown.
    /// </summary>
    public partial class ProjectMetadata : Pulumi.CustomResource
    {
        /// <summary>
        /// A series of key value pairs.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectMetadata resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectMetadata(string name, ProjectMetadataArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/projectMetadata:ProjectMetadata", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ProjectMetadata(string name, Input<string> id, ProjectMetadataState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/projectMetadata:ProjectMetadata", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectMetadata resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectMetadata Get(string name, Input<string> id, ProjectMetadataState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectMetadata(name, id, state, options);
        }
    }

    public sealed class ProjectMetadataArgs : Pulumi.ResourceArgs
    {
        [Input("metadata", required: true)]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A series of key value pairs.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectMetadataArgs()
        {
        }
    }

    public sealed class ProjectMetadataState : Pulumi.ResourceArgs
    {
        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A series of key value pairs.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectMetadataState()
        {
        }
    }
}
