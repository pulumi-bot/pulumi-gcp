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

@pulumi.output_type
class CryptoKeyIAMBindingCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """
    expression: str = pulumi.output_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: str = pulumi.output_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CryptoKeyIAMMemberCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """
    expression: str = pulumi.output_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: str = pulumi.output_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CryptoKeyVersionTemplate(dict):
    algorithm: str = pulumi.output_property("algorithm")
    """
    The algorithm to use when creating a version based on this template.
    See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.
    """
    protection_level: Optional[str] = pulumi.output_property("protectionLevel")
    """
    The protection level to use when creating a version based on this template.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KeyRingIAMBindingCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """
    expression: str = pulumi.output_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: str = pulumi.output_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KeyRingIAMMemberCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """
    expression: str = pulumi.output_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: str = pulumi.output_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KeyRingImportJobAttestation(dict):
    content: Optional[str] = pulumi.output_property("content")
    format: Optional[str] = pulumi.output_property("format")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KeyRingImportJobPublicKey(dict):
    pem: Optional[str] = pulumi.output_property("pem")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RegistryCredential(dict):
    public_key_certificate: 'outputs.RegistryCredentialPublicKeyCertificate' = pulumi.output_property("publicKeyCertificate")
    """
    A public key certificate format and data.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RegistryCredentialPublicKeyCertificate(dict):
    certificate: str = pulumi.output_property("certificate")
    """
    The certificate data.
    """
    format: str = pulumi.output_property("format")
    """
    The field allows only `X509_CERTIFICATE_PEM`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RegistryEventNotificationConfigItem(dict):
    pubsub_topic_name: str = pulumi.output_property("pubsubTopicName")
    """
    PubSub topic name to publish device events.
    """
    subfolder_matches: Optional[str] = pulumi.output_property("subfolderMatches")
    """
    If the subfolder name matches this string exactly, this
    configuration will be used. The string must not include the
    leading '/' character. If empty, all strings are matched. Empty
    value can only be used for the last `event_notification_configs`
    item.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RegistryHttpConfig(dict):
    http_enabled_state: str = pulumi.output_property("httpEnabledState")
    """
    The field allows `HTTP_ENABLED` or `HTTP_DISABLED`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RegistryMqttConfig(dict):
    mqtt_enabled_state: str = pulumi.output_property("mqttEnabledState")
    """
    The field allows `MQTT_ENABLED` or `MQTT_DISABLED`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RegistryStateNotificationConfig(dict):
    pubsub_topic_name: str = pulumi.output_property("pubsubTopicName")
    """
    PubSub topic name to publish device events.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetKMSCryptoKeyVersionPublicKey(dict):
    algorithm: str = pulumi.output_property("algorithm")
    """
    The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
    """
    pem: str = pulumi.output_property("pem")
    """
    The public key, encoded in PEM format. For more information, see the RFC 7468 sections for General Considerations and Textual Encoding of Subject Public Key Info.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetKMSCryptoKeyVersionTemplate(dict):
    algorithm: str = pulumi.output_property("algorithm")
    protection_level: str = pulumi.output_property("protectionLevel")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


