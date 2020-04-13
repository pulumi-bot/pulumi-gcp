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
    /// 
    /// To get more information about Attestor, see:
    /// 
    /// * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/binary-authorization/)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_attestor.html.markdown.
    /// </summary>
    public partial class Attestor : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Required)
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.  Structure is documented below.
        /// </summary>
        [Output("attestationAuthorityNote")]
        public Output<Outputs.AttestorAttestationAuthorityNote> AttestationAuthorityNote { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// A descriptive comment. This field may be updated. The field may be
        /// displayed in chooser dialogs.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
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
            : base("gcp:binaryauthorization/attestor:Attestor", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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
        /// -
        /// (Required)
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.  Structure is documented below.
        /// </summary>
        [Input("attestationAuthorityNote", required: true)]
        public Input<Inputs.AttestorAttestationAuthorityNoteArgs> AttestationAuthorityNote { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// A descriptive comment. This field may be updated. The field may be
        /// displayed in chooser dialogs.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Required)
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
        /// -
        /// (Required)
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.  Structure is documented below.
        /// </summary>
        [Input("attestationAuthorityNote")]
        public Input<Inputs.AttestorAttestationAuthorityNoteGetArgs>? AttestationAuthorityNote { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// A descriptive comment. This field may be updated. The field may be
        /// displayed in chooser dialogs.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// -
        /// (Required)
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

    namespace Inputs
    {

    public sealed class AttestorAttestationAuthorityNoteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// This field will contain the service account email address that
        /// this Attestor will use as the principal when querying Container
        /// Analysis. Attestor administrators must grant this service account
        /// the IAM role needed to read attestations from the noteReference in
        /// Container Analysis (containeranalysis.notes.occurrences.viewer).
        /// This email address is fixed for the lifetime of the Attestor, but
        /// callers should not make any other assumptions about the service
        /// account email; future versions may use an email based on a
        /// different naming pattern.
        /// </summary>
        [Input("delegationServiceAccountEmail")]
        public Input<string>? DelegationServiceAccountEmail { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The resource name of a ATTESTATION_AUTHORITY Note, created by the
        /// user. If the Note is in a different project from the Attestor, it
        /// should be specified in the format `projects/*/notes/*` (or the legacy
        /// `providers/*/notes/*`). This field may not be updated.
        /// An attestation by this attestor is stored as a Container Analysis
        /// ATTESTATION_AUTHORITY Occurrence that names a container image
        /// and that links to this Note.
        /// </summary>
        [Input("noteReference", required: true)]
        public Input<string> NoteReference { get; set; } = null!;

        [Input("publicKeys")]
        private InputList<AttestorAttestationAuthorityNotePublicKeysArgs>? _publicKeys;

        /// <summary>
        /// -
        /// (Optional)
        /// Public keys that verify attestations signed by this attestor. This
        /// field may be updated.
        /// If this field is non-empty, one of the specified public keys must
        /// verify that an attestation was signed by this attestor for the
        /// image specified in the admission request.
        /// If this field is empty, this attestor always returns that no valid
        /// attestations exist.  Structure is documented below.
        /// </summary>
        public InputList<AttestorAttestationAuthorityNotePublicKeysArgs> PublicKeys
        {
            get => _publicKeys ?? (_publicKeys = new InputList<AttestorAttestationAuthorityNotePublicKeysArgs>());
            set => _publicKeys = value;
        }

        public AttestorAttestationAuthorityNoteArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNoteGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// This field will contain the service account email address that
        /// this Attestor will use as the principal when querying Container
        /// Analysis. Attestor administrators must grant this service account
        /// the IAM role needed to read attestations from the noteReference in
        /// Container Analysis (containeranalysis.notes.occurrences.viewer).
        /// This email address is fixed for the lifetime of the Attestor, but
        /// callers should not make any other assumptions about the service
        /// account email; future versions may use an email based on a
        /// different naming pattern.
        /// </summary>
        [Input("delegationServiceAccountEmail")]
        public Input<string>? DelegationServiceAccountEmail { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The resource name of a ATTESTATION_AUTHORITY Note, created by the
        /// user. If the Note is in a different project from the Attestor, it
        /// should be specified in the format `projects/*/notes/*` (or the legacy
        /// `providers/*/notes/*`). This field may not be updated.
        /// An attestation by this attestor is stored as a Container Analysis
        /// ATTESTATION_AUTHORITY Occurrence that names a container image
        /// and that links to this Note.
        /// </summary>
        [Input("noteReference", required: true)]
        public Input<string> NoteReference { get; set; } = null!;

        [Input("publicKeys")]
        private InputList<AttestorAttestationAuthorityNotePublicKeysGetArgs>? _publicKeys;

        /// <summary>
        /// -
        /// (Optional)
        /// Public keys that verify attestations signed by this attestor. This
        /// field may be updated.
        /// If this field is non-empty, one of the specified public keys must
        /// verify that an attestation was signed by this attestor for the
        /// image specified in the admission request.
        /// If this field is empty, this attestor always returns that no valid
        /// attestations exist.  Structure is documented below.
        /// </summary>
        public InputList<AttestorAttestationAuthorityNotePublicKeysGetArgs> PublicKeys
        {
            get => _publicKeys ?? (_publicKeys = new InputList<AttestorAttestationAuthorityNotePublicKeysGetArgs>());
            set => _publicKeys = value;
        }

        public AttestorAttestationAuthorityNoteGetArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNotePublicKeysArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// ASCII-armored representation of a PGP public key, as the
        /// entire output by the command
        /// `gpg --export --armor foo@example.com` (either LF or CRLF
        /// line endings). When using this field, id should be left
        /// blank. The BinAuthz API handlers will calculate the ID
        /// and fill it in automatically. BinAuthz computes this ID
        /// as the OpenPGP RFC4880 V4 fingerprint, represented as
        /// upper-case hex. If id is provided by the caller, it will
        /// be overwritten by the API-calculated ID.
        /// </summary>
        [Input("asciiArmoredPgpPublicKey")]
        public Input<string>? AsciiArmoredPgpPublicKey { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// A descriptive comment. This field may be updated.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The ID of this public key. Signatures verified by BinAuthz
        /// must include the ID of the public key that can be used to
        /// verify them, and that ID must match the contents of this
        /// field exactly. Additional restrictions on this field can
        /// be imposed based on which public key type is encapsulated.
        /// See the documentation on publicKey cases below for details.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// A raw PKIX SubjectPublicKeyInfo format public key.
        /// NOTE: id may be explicitly provided by the caller when using this
        /// type of public key, but it MUST be a valid RFC3986 URI. If id is left
        /// blank, a default one will be computed based on the digest of the DER
        /// encoding of the public key.  Structure is documented below.
        /// </summary>
        [Input("pkixPublicKey")]
        public Input<AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyArgs>? PkixPublicKey { get; set; }

        public AttestorAttestationAuthorityNotePublicKeysArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNotePublicKeysGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// ASCII-armored representation of a PGP public key, as the
        /// entire output by the command
        /// `gpg --export --armor foo@example.com` (either LF or CRLF
        /// line endings). When using this field, id should be left
        /// blank. The BinAuthz API handlers will calculate the ID
        /// and fill it in automatically. BinAuthz computes this ID
        /// as the OpenPGP RFC4880 V4 fingerprint, represented as
        /// upper-case hex. If id is provided by the caller, it will
        /// be overwritten by the API-calculated ID.
        /// </summary>
        [Input("asciiArmoredPgpPublicKey")]
        public Input<string>? AsciiArmoredPgpPublicKey { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// A descriptive comment. This field may be updated.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The ID of this public key. Signatures verified by BinAuthz
        /// must include the ID of the public key that can be used to
        /// verify them, and that ID must match the contents of this
        /// field exactly. Additional restrictions on this field can
        /// be imposed based on which public key type is encapsulated.
        /// See the documentation on publicKey cases below for details.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// A raw PKIX SubjectPublicKeyInfo format public key.
        /// NOTE: id may be explicitly provided by the caller when using this
        /// type of public key, but it MUST be a valid RFC3986 URI. If id is left
        /// blank, a default one will be computed based on the digest of the DER
        /// encoding of the public key.  Structure is documented below.
        /// </summary>
        [Input("pkixPublicKey")]
        public Input<AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyGetArgs>? PkixPublicKey { get; set; }

        public AttestorAttestationAuthorityNotePublicKeysGetArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// A PEM-encoded public key, as described in
        /// `https://tools.ietf.org/html/rfc7468#section-13`
        /// </summary>
        [Input("publicKeyPem")]
        public Input<string>? PublicKeyPem { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The signature algorithm used to verify a message against
        /// a signature using this key. These signature algorithm must
        /// match the structure and any object identifiers encoded in
        /// publicKeyPem (i.e. this algorithm must match that of the
        /// public key).
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        public AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// A PEM-encoded public key, as described in
        /// `https://tools.ietf.org/html/rfc7468#section-13`
        /// </summary>
        [Input("publicKeyPem")]
        public Input<string>? PublicKeyPem { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The signature algorithm used to verify a message against
        /// a signature using this key. These signature algorithm must
        /// match the structure and any object identifiers encoded in
        /// publicKeyPem (i.e. this algorithm must match that of the
        /// public key).
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        public AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AttestorAttestationAuthorityNote
    {
        /// <summary>
        /// -
        /// This field will contain the service account email address that
        /// this Attestor will use as the principal when querying Container
        /// Analysis. Attestor administrators must grant this service account
        /// the IAM role needed to read attestations from the noteReference in
        /// Container Analysis (containeranalysis.notes.occurrences.viewer).
        /// This email address is fixed for the lifetime of the Attestor, but
        /// callers should not make any other assumptions about the service
        /// account email; future versions may use an email based on a
        /// different naming pattern.
        /// </summary>
        public readonly string DelegationServiceAccountEmail;
        /// <summary>
        /// -
        /// (Required)
        /// The resource name of a ATTESTATION_AUTHORITY Note, created by the
        /// user. If the Note is in a different project from the Attestor, it
        /// should be specified in the format `projects/*/notes/*` (or the legacy
        /// `providers/*/notes/*`). This field may not be updated.
        /// An attestation by this attestor is stored as a Container Analysis
        /// ATTESTATION_AUTHORITY Occurrence that names a container image
        /// and that links to this Note.
        /// </summary>
        public readonly string NoteReference;
        /// <summary>
        /// -
        /// (Optional)
        /// Public keys that verify attestations signed by this attestor. This
        /// field may be updated.
        /// If this field is non-empty, one of the specified public keys must
        /// verify that an attestation was signed by this attestor for the
        /// image specified in the admission request.
        /// If this field is empty, this attestor always returns that no valid
        /// attestations exist.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<AttestorAttestationAuthorityNotePublicKeys> PublicKeys;

        [OutputConstructor]
        private AttestorAttestationAuthorityNote(
            string delegationServiceAccountEmail,
            string noteReference,
            ImmutableArray<AttestorAttestationAuthorityNotePublicKeys> publicKeys)
        {
            DelegationServiceAccountEmail = delegationServiceAccountEmail;
            NoteReference = noteReference;
            PublicKeys = publicKeys;
        }
    }

    [OutputType]
    public sealed class AttestorAttestationAuthorityNotePublicKeys
    {
        /// <summary>
        /// -
        /// (Optional)
        /// ASCII-armored representation of a PGP public key, as the
        /// entire output by the command
        /// `gpg --export --armor foo@example.com` (either LF or CRLF
        /// line endings). When using this field, id should be left
        /// blank. The BinAuthz API handlers will calculate the ID
        /// and fill it in automatically. BinAuthz computes this ID
        /// as the OpenPGP RFC4880 V4 fingerprint, represented as
        /// upper-case hex. If id is provided by the caller, it will
        /// be overwritten by the API-calculated ID.
        /// </summary>
        public readonly string? AsciiArmoredPgpPublicKey;
        /// <summary>
        /// -
        /// (Optional)
        /// A descriptive comment. This field may be updated.
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// -
        /// (Optional)
        /// The ID of this public key. Signatures verified by BinAuthz
        /// must include the ID of the public key that can be used to
        /// verify them, and that ID must match the contents of this
        /// field exactly. Additional restrictions on this field can
        /// be imposed based on which public key type is encapsulated.
        /// See the documentation on publicKey cases below for details.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// -
        /// (Optional)
        /// A raw PKIX SubjectPublicKeyInfo format public key.
        /// NOTE: id may be explicitly provided by the caller when using this
        /// type of public key, but it MUST be a valid RFC3986 URI. If id is left
        /// blank, a default one will be computed based on the digest of the DER
        /// encoding of the public key.  Structure is documented below.
        /// </summary>
        public readonly AttestorAttestationAuthorityNotePublicKeysPkixPublicKey? PkixPublicKey;

        [OutputConstructor]
        private AttestorAttestationAuthorityNotePublicKeys(
            string? asciiArmoredPgpPublicKey,
            string? comment,
            string id,
            AttestorAttestationAuthorityNotePublicKeysPkixPublicKey? pkixPublicKey)
        {
            AsciiArmoredPgpPublicKey = asciiArmoredPgpPublicKey;
            Comment = comment;
            Id = id;
            PkixPublicKey = pkixPublicKey;
        }
    }

    [OutputType]
    public sealed class AttestorAttestationAuthorityNotePublicKeysPkixPublicKey
    {
        /// <summary>
        /// -
        /// (Optional)
        /// A PEM-encoded public key, as described in
        /// `https://tools.ietf.org/html/rfc7468#section-13`
        /// </summary>
        public readonly string? PublicKeyPem;
        /// <summary>
        /// -
        /// (Optional)
        /// The signature algorithm used to verify a message against
        /// a signature using this key. These signature algorithm must
        /// match the structure and any object identifiers encoded in
        /// publicKeyPem (i.e. this algorithm must match that of the
        /// public key).
        /// </summary>
        public readonly string? SignatureAlgorithm;

        [OutputConstructor]
        private AttestorAttestationAuthorityNotePublicKeysPkixPublicKey(
            string? publicKeyPem,
            string? signatureAlgorithm)
        {
            PublicKeyPem = publicKeyPem;
            SignatureAlgorithm = signatureAlgorithm;
        }
    }
    }
}
