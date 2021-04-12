# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'AttestorAttestationAuthorityNoteArgs',
    'AttestorAttestationAuthorityNotePublicKeyArgs',
    'AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs',
    'AttestorIamBindingConditionArgs',
    'AttestorIamMemberConditionArgs',
    'PolicyAdmissionWhitelistPatternArgs',
    'PolicyClusterAdmissionRuleArgs',
    'PolicyDefaultAdmissionRuleArgs',
]

@pulumi.input_type
class AttestorAttestationAuthorityNoteArgs:
    def __init__(__self__, *,
                 note_reference: pulumi.Input[str],
                 delegation_service_account_email: Optional[pulumi.Input[str]] = None,
                 public_keys: Optional[pulumi.Input[Sequence[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyArgs']]]] = None):
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
        :param pulumi.Input[Sequence[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyArgs']]] public_keys: Public keys that verify attestations signed by this attestor. This
               field may be updated.
               If this field is non-empty, one of the specified public keys must
               verify that an attestation was signed by this attestor for the
               image specified in the admission request.
               If this field is empty, this attestor always returns that no valid
               attestations exist.
               Structure is documented below.
        """
        pulumi.set(__self__, "note_reference", note_reference)
        if delegation_service_account_email is not None:
            pulumi.set(__self__, "delegation_service_account_email", delegation_service_account_email)
        if public_keys is not None:
            pulumi.set(__self__, "public_keys", public_keys)

    @property
    @pulumi.getter(name="noteReference")
    def note_reference(self) -> pulumi.Input[str]:
        """
        The resource name of a ATTESTATION_AUTHORITY Note, created by the
        user. If the Note is in a different project from the Attestor, it
        should be specified in the format `projects/*/notes/*` (or the legacy
        `providers/*/notes/*`). This field may not be updated.
        An attestation by this attestor is stored as a Container Analysis
        ATTESTATION_AUTHORITY Occurrence that names a container image
        and that links to this Note.
        """
        return pulumi.get(self, "note_reference")

    @note_reference.setter
    def note_reference(self, value: pulumi.Input[str]):
        pulumi.set(self, "note_reference", value)

    @property
    @pulumi.getter(name="delegationServiceAccountEmail")
    def delegation_service_account_email(self) -> Optional[pulumi.Input[str]]:
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
        return pulumi.get(self, "delegation_service_account_email")

    @delegation_service_account_email.setter
    def delegation_service_account_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delegation_service_account_email", value)

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyArgs']]]]:
        """
        Public keys that verify attestations signed by this attestor. This
        field may be updated.
        If this field is non-empty, one of the specified public keys must
        verify that an attestation was signed by this attestor for the
        image specified in the admission request.
        If this field is empty, this attestor always returns that no valid
        attestations exist.
        Structure is documented below.
        """
        return pulumi.get(self, "public_keys")

    @public_keys.setter
    def public_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyArgs']]]]):
        pulumi.set(self, "public_keys", value)


@pulumi.input_type
class AttestorAttestationAuthorityNotePublicKeyArgs:
    def __init__(__self__, *,
                 ascii_armored_pgp_public_key: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 pkix_public_key: Optional[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs']] = None):
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
               encoding of the public key.
               Structure is documented below.
        """
        if ascii_armored_pgp_public_key is not None:
            pulumi.set(__self__, "ascii_armored_pgp_public_key", ascii_armored_pgp_public_key)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if pkix_public_key is not None:
            pulumi.set(__self__, "pkix_public_key", pkix_public_key)

    @property
    @pulumi.getter(name="asciiArmoredPgpPublicKey")
    def ascii_armored_pgp_public_key(self) -> Optional[pulumi.Input[str]]:
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
        return pulumi.get(self, "ascii_armored_pgp_public_key")

    @ascii_armored_pgp_public_key.setter
    def ascii_armored_pgp_public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ascii_armored_pgp_public_key", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive comment. This field may be updated.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of this public key. Signatures verified by BinAuthz
        must include the ID of the public key that can be used to
        verify them, and that ID must match the contents of this
        field exactly. Additional restrictions on this field can
        be imposed based on which public key type is encapsulated.
        See the documentation on publicKey cases below for details.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="pkixPublicKey")
    def pkix_public_key(self) -> Optional[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs']]:
        """
        A raw PKIX SubjectPublicKeyInfo format public key.
        NOTE: id may be explicitly provided by the caller when using this
        type of public key, but it MUST be a valid RFC3986 URI. If id is left
        blank, a default one will be computed based on the digest of the DER
        encoding of the public key.
        Structure is documented below.
        """
        return pulumi.get(self, "pkix_public_key")

    @pkix_public_key.setter
    def pkix_public_key(self, value: Optional[pulumi.Input['AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs']]):
        pulumi.set(self, "pkix_public_key", value)


@pulumi.input_type
class AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs:
    def __init__(__self__, *,
                 public_key_pem: Optional[pulumi.Input[str]] = None,
                 signature_algorithm: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] public_key_pem: A PEM-encoded public key, as described in
               `https://tools.ietf.org/html/rfc7468#section-13`
        :param pulumi.Input[str] signature_algorithm: The signature algorithm used to verify a message against
               a signature using this key. These signature algorithm must
               match the structure and any object identifiers encoded in
               publicKeyPem (i.e. this algorithm must match that of the
               public key).
        """
        if public_key_pem is not None:
            pulumi.set(__self__, "public_key_pem", public_key_pem)
        if signature_algorithm is not None:
            pulumi.set(__self__, "signature_algorithm", signature_algorithm)

    @property
    @pulumi.getter(name="publicKeyPem")
    def public_key_pem(self) -> Optional[pulumi.Input[str]]:
        """
        A PEM-encoded public key, as described in
        `https://tools.ietf.org/html/rfc7468#section-13`
        """
        return pulumi.get(self, "public_key_pem")

    @public_key_pem.setter
    def public_key_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key_pem", value)

    @property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The signature algorithm used to verify a message against
        a signature using this key. These signature algorithm must
        match the structure and any object identifiers encoded in
        publicKeyPem (i.e. this algorithm must match that of the
        public key).
        """
        return pulumi.get(self, "signature_algorithm")

    @signature_algorithm.setter
    def signature_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signature_algorithm", value)


@pulumi.input_type
class AttestorIamBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class AttestorIamMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class PolicyAdmissionWhitelistPatternArgs:
    def __init__(__self__, *,
                 name_pattern: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name_pattern: An image name pattern to whitelist, in the form
               `registry/path/to/image`. This supports a trailing * as a
               wildcard, but this is allowed only in text after the registry/
               part.
        """
        pulumi.set(__self__, "name_pattern", name_pattern)

    @property
    @pulumi.getter(name="namePattern")
    def name_pattern(self) -> pulumi.Input[str]:
        """
        An image name pattern to whitelist, in the form
        `registry/path/to/image`. This supports a trailing * as a
        wildcard, but this is allowed only in text after the registry/
        part.
        """
        return pulumi.get(self, "name_pattern")

    @name_pattern.setter
    def name_pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "name_pattern", value)


@pulumi.input_type
class PolicyClusterAdmissionRuleArgs:
    def __init__(__self__, *,
                 cluster: pulumi.Input[str],
                 enforcement_mode: pulumi.Input[str],
                 evaluation_mode: pulumi.Input[str],
                 require_attestations_bies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] cluster: The identifier for this object. Format specified above.
        :param pulumi.Input[str] enforcement_mode: The action when a pod creation is denied by the admission rule.
               Possible values are `ENFORCED_BLOCK_AND_AUDIT_LOG` and `DRYRUN_AUDIT_LOG_ONLY`.
        :param pulumi.Input[str] evaluation_mode: How this admission rule will be evaluated.
               Possible values are `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, and `ALWAYS_DENY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] require_attestations_bies: The resource names of the attestors that must attest to a
               container image. If the attestor is in a different project from the
               policy, it should be specified in the format `projects/*/attestors/*`.
               Each attestor must exist before a policy can reference it. To add an
               attestor to a policy the principal issuing the policy change
               request must be able to read the attestor resource.
               Note: this field must be non-empty when the evaluation_mode field
               specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        pulumi.set(__self__, "cluster", cluster)
        pulumi.set(__self__, "enforcement_mode", enforcement_mode)
        pulumi.set(__self__, "evaluation_mode", evaluation_mode)
        if require_attestations_bies is not None:
            pulumi.set(__self__, "require_attestations_bies", require_attestations_bies)

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Input[str]:
        """
        The identifier for this object. Format specified above.
        """
        return pulumi.get(self, "cluster")

    @cluster.setter
    def cluster(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster", value)

    @property
    @pulumi.getter(name="enforcementMode")
    def enforcement_mode(self) -> pulumi.Input[str]:
        """
        The action when a pod creation is denied by the admission rule.
        Possible values are `ENFORCED_BLOCK_AND_AUDIT_LOG` and `DRYRUN_AUDIT_LOG_ONLY`.
        """
        return pulumi.get(self, "enforcement_mode")

    @enforcement_mode.setter
    def enforcement_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "enforcement_mode", value)

    @property
    @pulumi.getter(name="evaluationMode")
    def evaluation_mode(self) -> pulumi.Input[str]:
        """
        How this admission rule will be evaluated.
        Possible values are `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, and `ALWAYS_DENY`.
        """
        return pulumi.get(self, "evaluation_mode")

    @evaluation_mode.setter
    def evaluation_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "evaluation_mode", value)

    @property
    @pulumi.getter(name="requireAttestationsBies")
    def require_attestations_bies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
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
        return pulumi.get(self, "require_attestations_bies")

    @require_attestations_bies.setter
    def require_attestations_bies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "require_attestations_bies", value)


@pulumi.input_type
class PolicyDefaultAdmissionRuleArgs:
    def __init__(__self__, *,
                 enforcement_mode: pulumi.Input[str],
                 evaluation_mode: pulumi.Input[str],
                 require_attestations_bies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] enforcement_mode: The action when a pod creation is denied by the admission rule.
               Possible values are `ENFORCED_BLOCK_AND_AUDIT_LOG` and `DRYRUN_AUDIT_LOG_ONLY`.
        :param pulumi.Input[str] evaluation_mode: How this admission rule will be evaluated.
               Possible values are `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, and `ALWAYS_DENY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] require_attestations_bies: The resource names of the attestors that must attest to a
               container image. If the attestor is in a different project from the
               policy, it should be specified in the format `projects/*/attestors/*`.
               Each attestor must exist before a policy can reference it. To add an
               attestor to a policy the principal issuing the policy change
               request must be able to read the attestor resource.
               Note: this field must be non-empty when the evaluation_mode field
               specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        pulumi.set(__self__, "enforcement_mode", enforcement_mode)
        pulumi.set(__self__, "evaluation_mode", evaluation_mode)
        if require_attestations_bies is not None:
            pulumi.set(__self__, "require_attestations_bies", require_attestations_bies)

    @property
    @pulumi.getter(name="enforcementMode")
    def enforcement_mode(self) -> pulumi.Input[str]:
        """
        The action when a pod creation is denied by the admission rule.
        Possible values are `ENFORCED_BLOCK_AND_AUDIT_LOG` and `DRYRUN_AUDIT_LOG_ONLY`.
        """
        return pulumi.get(self, "enforcement_mode")

    @enforcement_mode.setter
    def enforcement_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "enforcement_mode", value)

    @property
    @pulumi.getter(name="evaluationMode")
    def evaluation_mode(self) -> pulumi.Input[str]:
        """
        How this admission rule will be evaluated.
        Possible values are `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, and `ALWAYS_DENY`.
        """
        return pulumi.get(self, "evaluation_mode")

    @evaluation_mode.setter
    def evaluation_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "evaluation_mode", value)

    @property
    @pulumi.getter(name="requireAttestationsBies")
    def require_attestations_bies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
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
        return pulumi.get(self, "require_attestations_bies")

    @require_attestations_bies.setter
    def require_attestations_bies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "require_attestations_bies", value)


