// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SourceRepo
{
    /// <summary>
    /// A repository (or repo) is a Git repository storing versioned source content.
    /// 
    /// To get more information about Repository, see:
    /// 
    /// * [API documentation](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/source-repositories/)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Repository can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:sourcerepo/repository:Repository default projects/{{project}}/repos/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:sourcerepo/repository:Repository default {{name}}
    /// ```
    /// </summary>
    public partial class Repository : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name of the repository, of the form `{{repo}}`.
        /// The repo name may contain slashes. eg, `name/with/slash`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// How this repository publishes a change in the repository through Cloud Pub/Sub.
        /// Keyed by the topic names.
        /// Structure is documented below.
        /// </summary>
        [Output("pubsubConfigs")]
        public Output<ImmutableArray<Outputs.RepositoryPubsubConfig>> PubsubConfigs { get; private set; } = null!;

        /// <summary>
        /// The disk usage of the repo, in bytes.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// URL to clone the repository from Google Cloud Source Repositories.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:sourcerepo/repository:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
            : base("gcp:sourcerepo/repository:Repository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, state, options);
        }
    }

    public sealed class RepositoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource name of the repository, of the form `{{repo}}`.
        /// The repo name may contain slashes. eg, `name/with/slash`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pubsubConfigs")]
        private InputList<Inputs.RepositoryPubsubConfigArgs>? _pubsubConfigs;

        /// <summary>
        /// How this repository publishes a change in the repository through Cloud Pub/Sub.
        /// Keyed by the topic names.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.RepositoryPubsubConfigArgs> PubsubConfigs
        {
            get => _pubsubConfigs ?? (_pubsubConfigs = new InputList<Inputs.RepositoryPubsubConfigArgs>());
            set => _pubsubConfigs = value;
        }

        public RepositoryArgs()
        {
        }
    }

    public sealed class RepositoryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource name of the repository, of the form `{{repo}}`.
        /// The repo name may contain slashes. eg, `name/with/slash`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pubsubConfigs")]
        private InputList<Inputs.RepositoryPubsubConfigGetArgs>? _pubsubConfigs;

        /// <summary>
        /// How this repository publishes a change in the repository through Cloud Pub/Sub.
        /// Keyed by the topic names.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.RepositoryPubsubConfigGetArgs> PubsubConfigs
        {
            get => _pubsubConfigs ?? (_pubsubConfigs = new InputList<Inputs.RepositoryPubsubConfigGetArgs>());
            set => _pubsubConfigs = value;
        }

        /// <summary>
        /// The disk usage of the repo, in bytes.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// URL to clone the repository from Google Cloud Source Repositories.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public RepositoryState()
        {
        }
    }
}
