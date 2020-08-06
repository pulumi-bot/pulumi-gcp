// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    /// <summary>
    /// A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
    /// (https://www.dicomstandard.org/about/) standard for Healthcare information exchange
    /// 
    /// To get more information about DicomStore, see:
    /// 
    /// * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.dicomStores)
    /// * How-to Guides
    ///     * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)
    /// 
    /// ## Example Usage
    /// ### Healthcare Dicom Store Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var topic = new Gcp.PubSub.Topic("topic", new Gcp.PubSub.TopicArgs
    ///         {
    ///         });
    ///         var dataset = new Gcp.Healthcare.Dataset("dataset", new Gcp.Healthcare.DatasetArgs
    ///         {
    ///             Location = "us-central1",
    ///         });
    ///         var @default = new Gcp.Healthcare.DicomStore("default", new Gcp.Healthcare.DicomStoreArgs
    ///         {
    ///             Dataset = dataset.Id,
    ///             NotificationConfig = new Gcp.Healthcare.Inputs.DicomStoreNotificationConfigArgs
    ///             {
    ///                 PubsubTopic = topic.Id,
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "label1", "labelvalue1" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DicomStore : Pulumi.CustomResource
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Output("dataset")]
        public Output<string> Dataset { get; private set; } = null!;

        /// <summary>
        /// User-supplied key-value pairs used to organize DICOM stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name for the DicomStore.
        /// ** Changing this property may recreate the Dicom store (removing all data) **
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Output("notificationConfig")]
        public Output<Outputs.DicomStoreNotificationConfig?> NotificationConfig { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of this dataset
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a DicomStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DicomStore(string name, DicomStoreArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/dicomStore:DicomStore", name, args ?? new DicomStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DicomStore(string name, Input<string> id, DicomStoreState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/dicomStore:DicomStore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DicomStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DicomStore Get(string name, Input<string> id, DicomStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new DicomStore(name, id, state, options);
        }
    }

    public sealed class DicomStoreArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Input("dataset", required: true)]
        public Input<string> Dataset { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize DICOM stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the DicomStore.
        /// ** Changing this property may recreate the Dicom store (removing all data) **
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.DicomStoreNotificationConfigArgs>? NotificationConfig { get; set; }

        public DicomStoreArgs()
        {
        }
    }

    public sealed class DicomStoreState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Input("dataset")]
        public Input<string>? Dataset { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize DICOM stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the DicomStore.
        /// ** Changing this property may recreate the Dicom store (removing all data) **
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A nested object resource  Structure is documented below.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.DicomStoreNotificationConfigGetArgs>? NotificationConfig { get; set; }

        /// <summary>
        /// The fully qualified name of this dataset
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public DicomStoreState()
        {
        }
    }
}
