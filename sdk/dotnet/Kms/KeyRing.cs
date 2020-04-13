// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    /// <summary>
    /// A `KeyRing` is a toplevel logical grouping of `CryptoKeys`.
    /// 
    /// 
    /// &gt; **Note:** KeyRings cannot be deleted from Google Cloud Platform.
    /// Destroying a provider-managed KeyRing will remove it from state but
    /// *will not delete the resource on the server.*
    /// 
    /// 
    /// To get more information about KeyRing, see:
    /// 
    /// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings)
    /// * How-to Guides
    ///     * [Creating a key ring](https://cloud.google.com/kms/docs/creating-keys#create_a_key_ring)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/kms_key_ring.html.markdown.
    /// </summary>
    public partial class KeyRing : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Required)
        /// The location for the KeyRing.
        /// A full list of valid locations can be found by running `gcloud kms locations list`.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The resource name for the KeyRing.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a KeyRing resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyRing(string name, KeyRingArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/keyRing:KeyRing", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private KeyRing(string name, Input<string> id, KeyRingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/keyRing:KeyRing", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyRing resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyRing Get(string name, Input<string> id, KeyRingState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyRing(name, id, state, options);
        }
    }

    public sealed class KeyRingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The location for the KeyRing.
        /// A full list of valid locations can be found by running `gcloud kms locations list`.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The resource name for the KeyRing.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public KeyRingArgs()
        {
        }
    }

    public sealed class KeyRingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The location for the KeyRing.
        /// A full list of valid locations can be found by running `gcloud kms locations list`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The resource name for the KeyRing.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public KeyRingState()
        {
        }
    }
}
