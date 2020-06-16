// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild
{
    /// <summary>
    /// Configuration for an automated build in response to source repository changes.
    /// 
    /// To get more information about Trigger, see:
    /// 
    /// * [API documentation](https://cloud.google.com/cloud-build/docs/api/reference/rest/)
    /// * How-to Guides
    ///     * [Automating builds using build triggers](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds)
    /// 
    /// {{% examples %}}
    /// ## Example Usage
    /// {{% example %}}
    /// ### Cloudbuild Trigger Filename
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var filename_trigger = new Gcp.CloudBuild.Trigger("filename-trigger", new Gcp.CloudBuild.TriggerArgs
    ///         {
    ///             Filename = "cloudbuild.yaml",
    ///             Substitutions = 
    ///             {
    ///                 { "_BAZ", "qux" },
    ///                 { "_FOO", "bar" },
    ///             },
    ///             TriggerTemplate = new Gcp.CloudBuild.Inputs.TriggerTriggerTemplateArgs
    ///             {
    ///                 BranchName = "master",
    ///                 RepoName = "my-repo",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// {{% /example %}}
    /// {{% /examples %}}
    /// </summary>
    public partial class Trigger : Pulumi.CustomResource
    {
        /// <summary>
        /// Contents of the build template. Either a filename or build template must be provided.  Structure is documented below.
        /// </summary>
        [Output("build")]
        public Output<Outputs.TriggerBuild?> Build { get; private set; } = null!;

        /// <summary>
        /// Time when the trigger was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Human-readable description of the trigger.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
        /// </summary>
        [Output("filename")]
        public Output<string?> Filename { get; private set; } = null!;

        /// <summary>
        /// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
        /// One of `trigger_template` or `github` must be provided.  Structure is documented below.
        /// </summary>
        [Output("github")]
        public Output<Outputs.TriggerGithub?> Github { get; private set; } = null!;

        /// <summary>
        /// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
        /// extended with support for `**`.
        /// If ignoredFiles and changed files are both empty, then they are not
        /// used to determine whether or not to trigger a build.
        /// If ignoredFiles is not empty, then we ignore any files that match any
        /// of the ignored_file globs. If the change has no files that are outside
        /// of the ignoredFiles globs, then we do not trigger a build.
        /// </summary>
        [Output("ignoredFiles")]
        public Output<ImmutableArray<string>> IgnoredFiles { get; private set; } = null!;

        /// <summary>
        /// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
        /// extended with support for `**`.
        /// If any of the files altered in the commit pass the ignoredFiles filter
        /// and includedFiles is empty, then as far as this filter is concerned, we
        /// should trigger the build.
        /// If any of the files altered in the commit pass the ignoredFiles filter
        /// and includedFiles is not empty, then we make sure that at least one of
        /// those files matches a includedFiles glob. If not, then we do not trigger
        /// a build.
        /// </summary>
        [Output("includedFiles")]
        public Output<ImmutableArray<string>> IncludedFiles { get; private set; } = null!;

        /// <summary>
        /// Name of the volume to mount.
        /// Volume names must be unique per build step and must be valid names for
        /// Docker volumes. Each named volume must be used by at least two build steps.
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
        /// Substitutions data for Build resource.
        /// </summary>
        [Output("substitutions")]
        public Output<ImmutableDictionary<string, string>?> Substitutions { get; private set; } = null!;

        /// <summary>
        /// The unique identifier for the trigger.
        /// </summary>
        [Output("triggerId")]
        public Output<string> TriggerId { get; private set; } = null!;

        /// <summary>
        /// Template describing the types of source changes to trigger a build.
        /// Branch and tag names in trigger templates are interpreted as regular
        /// expressions. Any branch or tag change that matches that regular
        /// expression will trigger a build.
        /// One of `trigger_template` or `github` must be provided.  Structure is documented below.
        /// </summary>
        [Output("triggerTemplate")]
        public Output<Outputs.TriggerTriggerTemplate?> TriggerTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a Trigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trigger(string name, TriggerArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:cloudbuild/trigger:Trigger", name, args ?? new TriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trigger(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
            : base("gcp:cloudbuild/trigger:Trigger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trigger Get(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new Trigger(name, id, state, options);
        }
    }

    public sealed class TriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contents of the build template. Either a filename or build template must be provided.  Structure is documented below.
        /// </summary>
        [Input("build")]
        public Input<Inputs.TriggerBuildArgs>? Build { get; set; }

        /// <summary>
        /// Human-readable description of the trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
        /// </summary>
        [Input("filename")]
        public Input<string>? Filename { get; set; }

        /// <summary>
        /// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
        /// One of `trigger_template` or `github` must be provided.  Structure is documented below.
        /// </summary>
        [Input("github")]
        public Input<Inputs.TriggerGithubArgs>? Github { get; set; }

        [Input("ignoredFiles")]
        private InputList<string>? _ignoredFiles;

        /// <summary>
        /// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
        /// extended with support for `**`.
        /// If ignoredFiles and changed files are both empty, then they are not
        /// used to determine whether or not to trigger a build.
        /// If ignoredFiles is not empty, then we ignore any files that match any
        /// of the ignored_file globs. If the change has no files that are outside
        /// of the ignoredFiles globs, then we do not trigger a build.
        /// </summary>
        public InputList<string> IgnoredFiles
        {
            get => _ignoredFiles ?? (_ignoredFiles = new InputList<string>());
            set => _ignoredFiles = value;
        }

        [Input("includedFiles")]
        private InputList<string>? _includedFiles;

        /// <summary>
        /// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
        /// extended with support for `**`.
        /// If any of the files altered in the commit pass the ignoredFiles filter
        /// and includedFiles is empty, then as far as this filter is concerned, we
        /// should trigger the build.
        /// If any of the files altered in the commit pass the ignoredFiles filter
        /// and includedFiles is not empty, then we make sure that at least one of
        /// those files matches a includedFiles glob. If not, then we do not trigger
        /// a build.
        /// </summary>
        public InputList<string> IncludedFiles
        {
            get => _includedFiles ?? (_includedFiles = new InputList<string>());
            set => _includedFiles = value;
        }

        /// <summary>
        /// Name of the volume to mount.
        /// Volume names must be unique per build step and must be valid names for
        /// Docker volumes. Each named volume must be used by at least two build steps.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("substitutions")]
        private InputMap<string>? _substitutions;

        /// <summary>
        /// Substitutions data for Build resource.
        /// </summary>
        public InputMap<string> Substitutions
        {
            get => _substitutions ?? (_substitutions = new InputMap<string>());
            set => _substitutions = value;
        }

        /// <summary>
        /// Template describing the types of source changes to trigger a build.
        /// Branch and tag names in trigger templates are interpreted as regular
        /// expressions. Any branch or tag change that matches that regular
        /// expression will trigger a build.
        /// One of `trigger_template` or `github` must be provided.  Structure is documented below.
        /// </summary>
        [Input("triggerTemplate")]
        public Input<Inputs.TriggerTriggerTemplateArgs>? TriggerTemplate { get; set; }

        public TriggerArgs()
        {
        }
    }

    public sealed class TriggerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contents of the build template. Either a filename or build template must be provided.  Structure is documented below.
        /// </summary>
        [Input("build")]
        public Input<Inputs.TriggerBuildGetArgs>? Build { get; set; }

        /// <summary>
        /// Time when the trigger was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Human-readable description of the trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
        /// </summary>
        [Input("filename")]
        public Input<string>? Filename { get; set; }

        /// <summary>
        /// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
        /// One of `trigger_template` or `github` must be provided.  Structure is documented below.
        /// </summary>
        [Input("github")]
        public Input<Inputs.TriggerGithubGetArgs>? Github { get; set; }

        [Input("ignoredFiles")]
        private InputList<string>? _ignoredFiles;

        /// <summary>
        /// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
        /// extended with support for `**`.
        /// If ignoredFiles and changed files are both empty, then they are not
        /// used to determine whether or not to trigger a build.
        /// If ignoredFiles is not empty, then we ignore any files that match any
        /// of the ignored_file globs. If the change has no files that are outside
        /// of the ignoredFiles globs, then we do not trigger a build.
        /// </summary>
        public InputList<string> IgnoredFiles
        {
            get => _ignoredFiles ?? (_ignoredFiles = new InputList<string>());
            set => _ignoredFiles = value;
        }

        [Input("includedFiles")]
        private InputList<string>? _includedFiles;

        /// <summary>
        /// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
        /// extended with support for `**`.
        /// If any of the files altered in the commit pass the ignoredFiles filter
        /// and includedFiles is empty, then as far as this filter is concerned, we
        /// should trigger the build.
        /// If any of the files altered in the commit pass the ignoredFiles filter
        /// and includedFiles is not empty, then we make sure that at least one of
        /// those files matches a includedFiles glob. If not, then we do not trigger
        /// a build.
        /// </summary>
        public InputList<string> IncludedFiles
        {
            get => _includedFiles ?? (_includedFiles = new InputList<string>());
            set => _includedFiles = value;
        }

        /// <summary>
        /// Name of the volume to mount.
        /// Volume names must be unique per build step and must be valid names for
        /// Docker volumes. Each named volume must be used by at least two build steps.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("substitutions")]
        private InputMap<string>? _substitutions;

        /// <summary>
        /// Substitutions data for Build resource.
        /// </summary>
        public InputMap<string> Substitutions
        {
            get => _substitutions ?? (_substitutions = new InputMap<string>());
            set => _substitutions = value;
        }

        /// <summary>
        /// The unique identifier for the trigger.
        /// </summary>
        [Input("triggerId")]
        public Input<string>? TriggerId { get; set; }

        /// <summary>
        /// Template describing the types of source changes to trigger a build.
        /// Branch and tag names in trigger templates are interpreted as regular
        /// expressions. Any branch or tag change that matches that regular
        /// expression will trigger a build.
        /// One of `trigger_template` or `github` must be provided.  Structure is documented below.
        /// </summary>
        [Input("triggerTemplate")]
        public Input<Inputs.TriggerTriggerTemplateGetArgs>? TriggerTemplate { get; set; }

        public TriggerState()
        {
        }
    }
}
