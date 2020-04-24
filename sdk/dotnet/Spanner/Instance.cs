// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Spanner
{
    /// <summary>
    /// An isolated set of Cloud Spanner resources on which databases can be
    /// hosted.
    /// 
    /// 
    /// To get more information about Instance, see:
    /// 
    /// * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/spanner/)
    /// </summary>
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Required)
        /// The name of the instance's configuration (similar but not
        /// quite the same as a region) which defines defines the geographic placement and
        /// replication of your databases in this instance. It determines where your data
        /// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
        /// In order to obtain a valid list please consult the
        /// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The descriptive name for this instance as it appears in UIs. Must be
        /// unique per project and between 4 and 30 characters in length.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// A unique identifier for the instance, which cannot be changed after
        /// the instance is created. The name must be between 6 and 30 characters
        /// in length.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The number of nodes allocated to this instance.
        /// </summary>
        [Output("numNodes")]
        public Output<int?> NumNodes { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Instance status: 'CREATING' or 'READY'.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("gcp:spanner/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:spanner/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The name of the instance's configuration (similar but not
        /// quite the same as a region) which defines defines the geographic placement and
        /// replication of your databases in this instance. It determines where your data
        /// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
        /// In order to obtain a valid list please consult the
        /// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The descriptive name for this instance as it appears in UIs. Must be
        /// unique per project and between 4 and 30 characters in length.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// -
        /// (Optional)
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// A unique identifier for the instance, which cannot be changed after
        /// the instance is created. The name must be between 6 and 30 characters
        /// in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The number of nodes allocated to this instance.
        /// </summary>
        [Input("numNodes")]
        public Input<int>? NumNodes { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The name of the instance's configuration (similar but not
        /// quite the same as a region) which defines defines the geographic placement and
        /// replication of your databases in this instance. It determines where your data
        /// is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
        /// In order to obtain a valid list please consult the
        /// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The descriptive name for this instance as it appears in UIs. Must be
        /// unique per project and between 4 and 30 characters in length.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// -
        /// (Optional)
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// A unique identifier for the instance, which cannot be changed after
        /// the instance is created. The name must be between 6 and 30 characters
        /// in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The number of nodes allocated to this instance.
        /// </summary>
        [Input("numNodes")]
        public Input<int>? NumNodes { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Instance status: 'CREATING' or 'READY'.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public InstanceState()
        {
        }
    }
}
