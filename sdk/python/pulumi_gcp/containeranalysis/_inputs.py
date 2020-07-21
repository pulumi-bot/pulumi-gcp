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
class NoteAttestationAuthorityArgs:
    hint: pulumi.Input['NoteAttestationAuthorityHintArgs'] = pulumi.input_property("hint")
    """
    This submessage provides human-readable hints about the purpose of
    the AttestationAuthority. Because the name of a Note acts as its
    resource reference, it is important to disambiguate the canonical
    name of the Note (which might be a UUID for security purposes)
    from "readable" names more suitable for debug output. Note that
    these hints should NOT be used to look up AttestationAuthorities
    in security sensitive contexts, such as when looking up
    Attestations to verify.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, hint: pulumi.Input['NoteAttestationAuthorityHintArgs']) -> None:
        """
        :param pulumi.Input['NoteAttestationAuthorityHintArgs'] hint: This submessage provides human-readable hints about the purpose of
               the AttestationAuthority. Because the name of a Note acts as its
               resource reference, it is important to disambiguate the canonical
               name of the Note (which might be a UUID for security purposes)
               from "readable" names more suitable for debug output. Note that
               these hints should NOT be used to look up AttestationAuthorities
               in security sensitive contexts, such as when looking up
               Attestations to verify.  Structure is documented below.
        """
        __self__.hint = hint

@pulumi.input_type
class NoteAttestationAuthorityHintArgs:
    human_readable_name: pulumi.Input[str] = pulumi.input_property("humanReadableName")
    """
    The human readable name of this Attestation Authority, for
    example "qa".
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, human_readable_name: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] human_readable_name: The human readable name of this Attestation Authority, for
               example "qa".
        """
        __self__.human_readable_name = human_readable_name

@pulumi.input_type
class NoteRelatedUrlArgs:
    url: pulumi.Input[str] = pulumi.input_property("url")
    """
    Specific URL associated with the resource.
    """
    label: Optional[pulumi.Input[str]] = pulumi.input_property("label")
    """
    Label to describe usage of the URL
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, url: pulumi.Input[str], label: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] url: Specific URL associated with the resource.
        :param pulumi.Input[str] label: Label to describe usage of the URL
        """
        __self__.url = url
        __self__.label = label

@pulumi.input_type
class OccurenceAttestationArgs:
    serialized_payload: pulumi.Input[str] = pulumi.input_property("serializedPayload")
    """
    The serialized payload that is verified by one or
    more signatures. A base64-encoded string.
    """
    signatures: pulumi.Input[List[pulumi.Input['OccurenceAttestationSignatureArgs']]] = pulumi.input_property("signatures")
    """
    One or more signatures over serializedPayload.
    Verifier implementations should consider this attestation
    message verified if at least one signature verifies
    serializedPayload. See Signature in common.proto for more
    details on signature structure and verification.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, serialized_payload: pulumi.Input[str], signatures: pulumi.Input[List[pulumi.Input['OccurenceAttestationSignatureArgs']]]) -> None:
        """
        :param pulumi.Input[str] serialized_payload: The serialized payload that is verified by one or
               more signatures. A base64-encoded string.
        :param pulumi.Input[List[pulumi.Input['OccurenceAttestationSignatureArgs']]] signatures: One or more signatures over serializedPayload.
               Verifier implementations should consider this attestation
               message verified if at least one signature verifies
               serializedPayload. See Signature in common.proto for more
               details on signature structure and verification.  Structure is documented below.
        """
        __self__.serialized_payload = serialized_payload
        __self__.signatures = signatures

@pulumi.input_type
class OccurenceAttestationSignatureArgs:
    public_key_id: pulumi.Input[str] = pulumi.input_property("publicKeyId")
    """
    The identifier for the public key that verifies this
    signature. MUST be an RFC3986 conformant
    URI. * When possible, the key id should be an
    immutable reference, such as a cryptographic digest.
    Examples of valid values:
    * OpenPGP V4 public key fingerprint. See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr
    for more details on this scheme.
    * `openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA`
    * RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER serialization):
    * "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU"
    """
    signature: Optional[pulumi.Input[str]] = pulumi.input_property("signature")
    """
    The content of the signature, an opaque bytestring.
    The payload that this signature verifies MUST be
    unambiguously provided with the Signature during
    verification. A wrapper message might provide the
    payload explicitly. Alternatively, a message might
    have a canonical serialization that can always be
    unambiguously computed to derive the payload.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, public_key_id: pulumi.Input[str], signature: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] public_key_id: The identifier for the public key that verifies this
               signature. MUST be an RFC3986 conformant
               URI. * When possible, the key id should be an
               immutable reference, such as a cryptographic digest.
               Examples of valid values:
               * OpenPGP V4 public key fingerprint. See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr
               for more details on this scheme.
               * `openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA`
               * RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER serialization):
               * "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU"
        :param pulumi.Input[str] signature: The content of the signature, an opaque bytestring.
               The payload that this signature verifies MUST be
               unambiguously provided with the Signature during
               verification. A wrapper message might provide the
               payload explicitly. Alternatively, a message might
               have a canonical serialization that can always be
               unambiguously computed to derive the payload.
        """
        __self__.public_key_id = public_key_id
        __self__.signature = signature

