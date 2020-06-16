// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BinaryAuthorization
{
    /// <summary>
    /// An attestor that attests to container image artifacts.
    /// 
    /// To get more information about Attestor, see:
    /// 
    /// * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/binary-authorization/)
    /// </summary>
    public partial class Attestor : Pulumi.CustomResource
    {
        /// <summary>
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.  Structure is documented below.
        /// </summary>
        [Output("attestationAuthorityNote")]
        public Output<Outputs.AttestorAttestationAuthorityNote> AttestationAuthorityNote { get; private set; } = null!;

        /// <summary>
        /// A descriptive comment. This field may be updated. The field may be
        /// displayed in chooser dialogs.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The resource name.
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
        /// Create a Attestor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Attestor(string name, AttestorArgs args, CustomResourceOptions? options = null)
            : base("gcp:binaryauthorization/attestor:Attestor", name, args ?? new AttestorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Attestor(string name, Input<string> id, AttestorState? state = null, CustomResourceOptions? options = null)
            : base("gcp:binaryauthorization/attestor:Attestor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Attestor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Attestor Get(string name, Input<string> id, AttestorState? state = null, CustomResourceOptions? options = null)
        {
            return new Attestor(name, id, state, options);
        }
    }

    public sealed class AttestorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.  Structure is documented below.
        /// </summary>
        [Input("attestationAuthorityNote", required: true)]
        public Input<Inputs.AttestorAttestationAuthorityNoteArgs> AttestationAuthorityNote { get; set; } = null!;

        /// <summary>
        /// A descriptive comment. This field may be updated. The field may be
        /// displayed in chooser dialogs.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public AttestorArgs()
        {
        }
    }

    public sealed class AttestorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.  Structure is documented below.
        /// </summary>
        [Input("attestationAuthorityNote")]
        public Input<Inputs.AttestorAttestationAuthorityNoteGetArgs>? AttestationAuthorityNote { get; set; }

        /// <summary>
        /// A descriptive comment. This field may be updated. The field may be
        /// displayed in chooser dialogs.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public AttestorState()
        {
        }
    }
}
