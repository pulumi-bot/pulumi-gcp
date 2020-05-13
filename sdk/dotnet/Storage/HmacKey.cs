// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    /// <summary>
    /// The hmacKeys resource represents an HMAC key within Cloud Storage. The resource
    /// consists of a secret and HMAC key metadata. HMAC keys can be used as credentials
    /// for service accounts.
    /// 
    /// 
    /// To get more information about HmacKey, see:
    /// 
    /// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/storage/docs/authentication/managing-hmackeys)
    /// 
    /// &gt; **Warning:** All arguments including the `secret` value will be stored in the raw
    /// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
    /// On import, the `secret` value will not be retrieved.
    /// 
    /// &gt; **Warning:** All arguments including `secret` will be stored in the raw
    /// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
    /// </summary>
    public partial class HmacKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The access ID of the HMAC Key.
        /// </summary>
        [Output("accessId")]
        public Output<string> AccessId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// HMAC secret key material.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        /// <summary>
        /// The email address of the key's associated service account.
        /// </summary>
        [Output("serviceAccountEmail")]
        public Output<string> ServiceAccountEmail { get; private set; } = null!;

        /// <summary>
        /// The state of the key. Can be set to one of ACTIVE, INACTIVE.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// 'The creation time of the HMAC key in RFC 3339 format. '
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a HmacKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HmacKey(string name, HmacKeyArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/hmacKey:HmacKey", name, args ?? new HmacKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HmacKey(string name, Input<string> id, HmacKeyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/hmacKey:HmacKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HmacKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HmacKey Get(string name, Input<string> id, HmacKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new HmacKey(name, id, state, options);
        }
    }

    public sealed class HmacKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The email address of the key's associated service account.
        /// </summary>
        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        /// <summary>
        /// The state of the key. Can be set to one of ACTIVE, INACTIVE.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public HmacKeyArgs()
        {
        }
    }

    public sealed class HmacKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access ID of the HMAC Key.
        /// </summary>
        [Input("accessId")]
        public Input<string>? AccessId { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// HMAC secret key material.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// The email address of the key's associated service account.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// The state of the key. Can be set to one of ACTIVE, INACTIVE.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// 'The creation time of the HMAC key in RFC 3339 format. '
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        public HmacKeyState()
        {
        }
    }
}
