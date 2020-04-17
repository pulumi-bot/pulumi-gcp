// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub
{
    /// <summary>
    /// A named resource to which messages are sent by publishers.
    /// 
    /// 
    /// To get more information about Topic, see:
    /// 
    /// * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
    /// * How-to Guides
    ///     * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)
    /// </summary>
    public partial class Topic : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The resource name of the Cloud KMS CryptoKey to be used to protect access
        /// to messages published on this topic. Your project's PubSub service account
        /// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
        /// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
        /// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        /// </summary>
        [Output("kmsKeyName")]
        public Output<string?> KmsKeyName { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// A set of key/value label pairs to assign to this Topic.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Policy constraining the set of Google Cloud Platform regions where
        /// messages published to the topic may be stored. If not present, then no
        /// constraints are in effect.  Structure is documented below.
        /// </summary>
        [Output("messageStoragePolicy")]
        public Output<Outputs.TopicMessageStoragePolicy> MessageStoragePolicy { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// Name of the topic.
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
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:pubsub/topic:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("gcp:pubsub/topic:Topic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The resource name of the Cloud KMS CryptoKey to be used to protect access
        /// to messages published on this topic. Your project's PubSub service account
        /// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
        /// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
        /// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// -
        /// (Optional)
        /// A set of key/value label pairs to assign to this Topic.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Policy constraining the set of Google Cloud Platform regions where
        /// messages published to the topic may be stored. If not present, then no
        /// constraints are in effect.  Structure is documented below.
        /// </summary>
        [Input("messageStoragePolicy")]
        public Input<Inputs.TopicMessageStoragePolicyArgs>? MessageStoragePolicy { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the topic.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public TopicArgs()
        {
        }
    }

    public sealed class TopicState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The resource name of the Cloud KMS CryptoKey to be used to protect access
        /// to messages published on this topic. Your project's PubSub service account
        /// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
        /// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
        /// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// -
        /// (Optional)
        /// A set of key/value label pairs to assign to this Topic.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Policy constraining the set of Google Cloud Platform regions where
        /// messages published to the topic may be stored. If not present, then no
        /// constraints are in effect.  Structure is documented below.
        /// </summary>
        [Input("messageStoragePolicy")]
        public Input<Inputs.TopicMessageStoragePolicyGetArgs>? MessageStoragePolicy { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the topic.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public TopicState()
        {
        }
    }
}
