// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ArtifactRegistry
{
    /// <summary>
    /// A repository for storing artifacts
    /// 
    /// To get more information about Repository, see:
    /// 
    /// * [API documentation](https://cloud.google.com/artifact-registry/docs/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/artifact-registry/docs/overview)
    /// 
    /// ## Example Usage
    /// ### Artifact Registry Repository Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var my_repo = new Gcp.ArtifactRegistry.Repository("my-repo", new Gcp.ArtifactRegistry.RepositoryArgs
    ///         {
    ///             Location = "us-central1",
    ///             RepositoryId = "my-repository",
    ///             Description = "example docker repository",
    ///             Format = "DOCKER",
    ///         }, new CustomResourceOptions {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Artifact Registry Repository Iam
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var my_repo = new Gcp.ArtifactRegistry.Repository("my-repo", new Gcp.ArtifactRegistry.RepositoryArgs
    ///         {
    ///             Location = "us-central1",
    ///             RepositoryId = "my-repository",
    ///             Description = "example docker repository with iam",
    ///             Format = "DOCKER",
    ///         }, new CustomResourceOptions {
    ///             Provider = google_beta,
    ///         });
    ///         var test_account = new Gcp.ServiceAccount.Account("test-account", new Gcp.ServiceAccount.AccountArgs
    ///         {
    ///             AccountId = "my-account",
    ///             DisplayName = "Test Service Account",
    ///         }, new CustomResourceOptions {
    ///             Provider = google_beta,
    ///         });
    ///         var test_iam = new Gcp.ArtifactRegistry.RepositoryIamMember("test-iam", new Gcp.ArtifactRegistry.RepositoryIamMemberArgs
    ///         {
    ///             Location = my_repo.Location,
    ///             Repository = my_repo.Name,
    ///             Role = "roles/artifactregistry.reader",
    ///             Member = test_account.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///         }, new CustomResourceOptions {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Repository : Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the repository was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The user-provided description of the repository.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The format of packages that are stored in the repoitory.
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// Labels with user-defined metadata.
        /// This field may contain up to 64 entries. Label keys and values may be no
        /// longer than 63 characters. Label keys must begin with a lowercase letter
        /// and may only contain lowercase letters, numeric characters, underscores,
        /// and dashes.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the location this repository is located in.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1"
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
        /// The last part of the repository name, for example:
        /// "repo1"
        /// </summary>
        [Output("repositoryId")]
        public Output<string> RepositoryId { get; private set; } = null!;

        /// <summary>
        /// The time when the repository was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs args, CustomResourceOptions? options = null)
            : base("gcp:artifactregistry/repository:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
            : base("gcp:artifactregistry/repository:Repository", name, state, MakeResourceOptions(options, id))
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
        /// The user-provided description of the repository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The format of packages that are stored in the repoitory.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels with user-defined metadata.
        /// This field may contain up to 64 entries. Label keys and values may be no
        /// longer than 63 characters. Label keys must begin with a lowercase letter
        /// and may only contain lowercase letters, numeric characters, underscores,
        /// and dashes.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the location this repository is located in.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The last part of the repository name, for example:
        /// "repo1"
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        public RepositoryArgs()
        {
        }
    }

    public sealed class RepositoryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when the repository was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The user-provided description of the repository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The format of packages that are stored in the repoitory.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels with user-defined metadata.
        /// This field may contain up to 64 entries. Label keys and values may be no
        /// longer than 63 characters. Label keys must begin with a lowercase letter
        /// and may only contain lowercase letters, numeric characters, underscores,
        /// and dashes.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the location this repository is located in.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The last part of the repository name, for example:
        /// "repo1"
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        /// <summary>
        /// The time when the repository was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public RepositoryState()
        {
        }
    }
}
