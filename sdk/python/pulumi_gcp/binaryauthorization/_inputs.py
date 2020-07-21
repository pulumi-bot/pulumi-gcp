# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs

@pulumi.input_type
class AttestorAttestationAuthorityNoteArgs:
    note_reference: pulumi.Input[str] = pulumi.input_property("noteReference")
    """
    The resource name of a ATTESTATION_AUTHORITY Note, created by the
    user. If the Note is in a different project from the Attestor, it
    should be specified in the format `projects/*/notes/*` (or the legacy
    `providers/*/notes/*`). This field may not be updated.
    An attestation by this attestor is stored as a Container Analysis
    ATTESTATION_AUTHORITY Occurrence that names a container image
    and that links to this Note.
    """
    delegation_service_account_email: Optional[pulumi.Input[str]] = pulumi.input_property("delegationServiceAccountEmail")
    """
    -
    This field will contain the service account email address that
    this Attestor will use as the principal when querying Container
    Analysis. Attestor administrators must grant this service account
    the IAM role needed to read attestations from the noteReference in
    Container Analysis (containeranalysis.notes.occurrences.viewer).
    This email address is fixed for the lifetime of the Attestor, but
    callers should not make any other assumptions about the service
    account email; future versions may use an email based on a
    different naming pattern.
    """
    public_keys: Optional[pulumi.Input[List[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyArgs']]]] = pulumi.input_property("publicKeys")
    """
    Public keys that verify attestations signed by this attestor. This
    field may be updated.
    If this field is non-empty, one of the specified public keys must
    verify that an attestation was signed by this attestor for the
    image specified in the admission request.
    If this field is empty, this attestor always returns that no valid
    attestations exist.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, note_reference: pulumi.Input[str], delegation_service_account_email: Optional[pulumi.Input[str]] = None, public_keys: Optional[pulumi.Input[List[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyArgs']]]] = None) -> None:
        """
        :param pulumi.Input[str] note_reference: The resource name of a ATTESTATION_AUTHORITY Note, created by the
               user. If the Note is in a different project from the Attestor, it
               should be specified in the format `projects/*/notes/*` (or the legacy
               `providers/*/notes/*`). This field may not be updated.
               An attestation by this attestor is stored as a Container Analysis
               ATTESTATION_AUTHORITY Occurrence that names a container image
               and that links to this Note.
        :param pulumi.Input[str] delegation_service_account_email: -
               This field will contain the service account email address that
               this Attestor will use as the principal when querying Container
               Analysis. Attestor administrators must grant this service account
               the IAM role needed to read attestations from the noteReference in
               Container Analysis (containeranalysis.notes.occurrences.viewer).
               This email address is fixed for the lifetime of the Attestor, but
               callers should not make any other assumptions about the service
               account email; future versions may use an email based on a
               different naming pattern.
        :param pulumi.Input[List[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyArgs']]] public_keys: Public keys that verify attestations signed by this attestor. This
               field may be updated.
               If this field is non-empty, one of the specified public keys must
               verify that an attestation was signed by this attestor for the
               image specified in the admission request.
               If this field is empty, this attestor always returns that no valid
               attestations exist.  Structure is documented below.
        """
        __self__.note_reference = note_reference
        __self__.delegation_service_account_email = delegation_service_account_email
        __self__.public_keys = public_keys

@pulumi.input_type
class AttestorAttestationAuthorityNotePublicKeyArgs:
    ascii_armored_pgp_public_key: Optional[pulumi.Input[str]] = pulumi.input_property("asciiArmoredPgpPublicKey")
    """
    ASCII-armored representation of a PGP public key, as the
    entire output by the command
    `gpg --export --armor foo@example.com` (either LF or CRLF
    line endings). When using this field, id should be left
    blank. The BinAuthz API handlers will calculate the ID
    and fill it in automatically. BinAuthz computes this ID
    as the OpenPGP RFC4880 V4 fingerprint, represented as
    upper-case hex. If id is provided by the caller, it will
    be overwritten by the API-calculated ID.
    """
    comment: Optional[pulumi.Input[str]] = pulumi.input_property("comment")
    """
    A descriptive comment. This field may be updated.
    """
    id: Optional[pulumi.Input[str]] = pulumi.input_property("id")
    """
    The ID of this public key. Signatures verified by BinAuthz
    must include the ID of the public key that can be used to
    verify them, and that ID must match the contents of this
    field exactly. Additional restrictions on this field can
    be imposed based on which public key type is encapsulated.
    See the documentation on publicKey cases below for details.
    """
    pkix_public_key: Optional[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs']] = pulumi.input_property("pkixPublicKey")
    """
    A raw PKIX SubjectPublicKeyInfo format public key.
    NOTE: id may be explicitly provided by the caller when using this
    type of public key, but it MUST be a valid RFC3986 URI. If id is left
    blank, a default one will be computed based on the digest of the DER
    encoding of the public key.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, ascii_armored_pgp_public_key: Optional[pulumi.Input[str]] = None, comment: Optional[pulumi.Input[str]] = None, id: Optional[pulumi.Input[str]] = None, pkix_public_key: Optional[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs']] = None) -> None:
        """
        :param pulumi.Input[str] ascii_armored_pgp_public_key: ASCII-armored representation of a PGP public key, as the
               entire output by the command
               `gpg --export --armor foo@example.com` (either LF or CRLF
               line endings). When using this field, id should be left
               blank. The BinAuthz API handlers will calculate the ID
               and fill it in automatically. BinAuthz computes this ID
               as the OpenPGP RFC4880 V4 fingerprint, represented as
               upper-case hex. If id is provided by the caller, it will
               be overwritten by the API-calculated ID.
        :param pulumi.Input[str] comment: A descriptive comment. This field may be updated.
        :param pulumi.Input[str] id: The ID of this public key. Signatures verified by BinAuthz
               must include the ID of the public key that can be used to
               verify them, and that ID must match the contents of this
               field exactly. Additional restrictions on this field can
               be imposed based on which public key type is encapsulated.
               See the documentation on publicKey cases below for details.
        :param pulumi.Input['AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs'] pkix_public_key: A raw PKIX SubjectPublicKeyInfo format public key.
               NOTE: id may be explicitly provided by the caller when using this
               type of public key, but it MUST be a valid RFC3986 URI. If id is left
               blank, a default one will be computed based on the digest of the DER
               encoding of the public key.  Structure is documented below.
        """
        __self__.ascii_armored_pgp_public_key = ascii_armored_pgp_public_key
        __self__.comment = comment
        __self__.id = id
        __self__.pkix_public_key = pkix_public_key

@pulumi.input_type
class AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs:
    public_key_pem: Optional[pulumi.Input[str]] = pulumi.input_property("publicKeyPem")
    """
    A PEM-encoded public key, as described in
    `https://tools.ietf.org/html/rfc7468#section-13`
    """
    signature_algorithm: Optional[pulumi.Input[str]] = pulumi.input_property("signatureAlgorithm")
    """
    The signature algorithm used to verify a message against
    a signature using this key. These signature algorithm must
    match the structure and any object identifiers encoded in
    publicKeyPem (i.e. this algorithm must match that of the
    public key).
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, public_key_pem: Optional[pulumi.Input[str]] = None, signature_algorithm: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] public_key_pem: A PEM-encoded public key, as described in
               `https://tools.ietf.org/html/rfc7468#section-13`
        :param pulumi.Input[str] signature_algorithm: The signature algorithm used to verify a message against
               a signature using this key. These signature algorithm must
               match the structure and any object identifiers encoded in
               publicKeyPem (i.e. this algorithm must match that of the
               public key).
        """
        __self__.public_key_pem = public_key_pem
        __self__.signature_algorithm = signature_algorithm

@pulumi.input_type
class AttestorIamBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    title: pulumi.Input[str] = pulumi.input_property("title")
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class AttestorIamMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    title: pulumi.Input[str] = pulumi.input_property("title")
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class PolicyAdmissionWhitelistPatternArgs:
    name_pattern: pulumi.Input[str] = pulumi.input_property("namePattern")
    """
    An image name pattern to whitelist, in the form
    `registry/path/to/image`. This supports a trailing * as a
    wildcard, but this is allowed only in text after the registry/
    part.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name_pattern: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] name_pattern: An image name pattern to whitelist, in the form
               `registry/path/to/image`. This supports a trailing * as a
               wildcard, but this is allowed only in text after the registry/
               part.
        """
        __self__.name_pattern = name_pattern

@pulumi.input_type
class PolicyClusterAdmissionRuleArgs:
    cluster: pulumi.Input[str] = pulumi.input_property("cluster")
    """
    The identifier for this object. Format specified above.
    """
    enforcement_mode: pulumi.Input[str] = pulumi.input_property("enforcementMode")
    """
    The action when a pod creation is denied by the admission rule.
    """
    evaluation_mode: pulumi.Input[str] = pulumi.input_property("evaluationMode")
    """
    How this admission rule will be evaluated.
    """
    require_attestations_bies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("requireAttestationsBies")
    """
    The resource names of the attestors that must attest to a
    container image. If the attestor is in a different project from the
    policy, it should be specified in the format `projects/*/attestors/*`.
    Each attestor must exist before a policy can reference it. To add an
    attestor to a policy the principal issuing the policy change
    request must be able to read the attestor resource.
    Note: this field must be non-empty when the evaluation_mode field
    specifies REQUIRE_ATTESTATION, otherwise it must be empty.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, cluster: pulumi.Input[str], enforcement_mode: pulumi.Input[str], evaluation_mode: pulumi.Input[str], require_attestations_bies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input[str] cluster: The identifier for this object. Format specified above.
        :param pulumi.Input[str] enforcement_mode: The action when a pod creation is denied by the admission rule.
        :param pulumi.Input[str] evaluation_mode: How this admission rule will be evaluated.
        :param pulumi.Input[List[pulumi.Input[str]]] require_attestations_bies: The resource names of the attestors that must attest to a
               container image. If the attestor is in a different project from the
               policy, it should be specified in the format `projects/*/attestors/*`.
               Each attestor must exist before a policy can reference it. To add an
               attestor to a policy the principal issuing the policy change
               request must be able to read the attestor resource.
               Note: this field must be non-empty when the evaluation_mode field
               specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        __self__.cluster = cluster
        __self__.enforcement_mode = enforcement_mode
        __self__.evaluation_mode = evaluation_mode
        __self__.require_attestations_bies = require_attestations_bies

@pulumi.input_type
class PolicyDefaultAdmissionRuleArgs:
    enforcement_mode: pulumi.Input[str] = pulumi.input_property("enforcementMode")
    """
    The action when a pod creation is denied by the admission rule.
    """
    evaluation_mode: pulumi.Input[str] = pulumi.input_property("evaluationMode")
    """
    How this admission rule will be evaluated.
    """
    require_attestations_bies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("requireAttestationsBies")
    """
    The resource names of the attestors that must attest to a
    container image. If the attestor is in a different project from the
    policy, it should be specified in the format `projects/*/attestors/*`.
    Each attestor must exist before a policy can reference it. To add an
    attestor to a policy the principal issuing the policy change
    request must be able to read the attestor resource.
    Note: this field must be non-empty when the evaluation_mode field
    specifies REQUIRE_ATTESTATION, otherwise it must be empty.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, enforcement_mode: pulumi.Input[str], evaluation_mode: pulumi.Input[str], require_attestations_bies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input[str] enforcement_mode: The action when a pod creation is denied by the admission rule.
        :param pulumi.Input[str] evaluation_mode: How this admission rule will be evaluated.
        :param pulumi.Input[List[pulumi.Input[str]]] require_attestations_bies: The resource names of the attestors that must attest to a
               container image. If the attestor is in a different project from the
               policy, it should be specified in the format `projects/*/attestors/*`.
               Each attestor must exist before a policy can reference it. To add an
               attestor to a policy the principal issuing the policy change
               request must be able to read the attestor resource.
               Note: this field must be non-empty when the evaluation_mode field
               specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        __self__.enforcement_mode = enforcement_mode
        __self__.evaluation_mode = evaluation_mode
        __self__.require_attestations_bies = require_attestations_bies

