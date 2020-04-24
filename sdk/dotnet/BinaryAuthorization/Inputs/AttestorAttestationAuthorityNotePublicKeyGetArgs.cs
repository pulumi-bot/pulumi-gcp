// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BinaryAuthorization.Inputs
{

    public sealed class AttestorAttestationAuthorityNotePublicKeyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
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
        /// A descriptive comment. This field may be updated.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
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
        /// A raw PKIX SubjectPublicKeyInfo format public key.
        /// NOTE: id may be explicitly provided by the caller when using this
        /// type of public key, but it MUST be a valid RFC3986 URI. If id is left
        /// blank, a default one will be computed based on the digest of the DER
        /// encoding of the public key.  Structure is documented below.
        /// </summary>
        [Input("pkixPublicKey")]
        public Input<Inputs.AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyGetArgs>? PkixPublicKey { get; set; }

        public AttestorAttestationAuthorityNotePublicKeyGetArgs()
        {
        }
    }
}
