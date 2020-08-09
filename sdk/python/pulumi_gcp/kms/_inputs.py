# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'CryptoKeyIAMBindingConditionArgs',
    'CryptoKeyIAMMemberConditionArgs',
    'CryptoKeyVersionTemplateArgs',
    'KeyRingIAMBindingConditionArgs',
    'KeyRingIAMMemberConditionArgs',
    'KeyRingImportJobAttestationArgs',
    'KeyRingImportJobPublicKeyArgs',
    'RegistryCredentialArgs',
    'RegistryCredentialPublicKeyCertificateArgs',
    'RegistryEventNotificationConfigItemArgs',
    'RegistryHttpConfigArgs',
    'RegistryMqttConfigArgs',
    'RegistryStateNotificationConfigArgs',
]

@pulumi.input_type
class CryptoKeyIAMBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        ...

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        ...

    @title.setter
    def title(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class CryptoKeyIAMMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        ...

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        ...

    @title.setter
    def title(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class CryptoKeyVersionTemplateArgs:
    def __init__(__self__, *,
                 algorithm: pulumi.Input[str],
                 protection_level: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] algorithm: The algorithm to use when creating a version based on this template.
               See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.
        :param pulumi.Input[str] protection_level: The protection level to use when creating a version based on this template.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "protectionLevel", protection_level)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Input[str]:
        """
        The algorithm to use when creating a version based on this template.
        See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.
        """
        ...

    @algorithm.setter
    def algorithm(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> Optional[pulumi.Input[str]]:
        """
        The protection level to use when creating a version based on this template.
        """
        ...

    @protection_level.setter
    def protection_level(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class KeyRingIAMBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        ...

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        ...

    @title.setter
    def title(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class KeyRingIAMMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        ...

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        ...

    @title.setter
    def title(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class KeyRingImportJobAttestationArgs:
    def __init__(__self__, *,
                 content: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        ...

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        ...

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class KeyRingImportJobPublicKeyArgs:
    def __init__(__self__, *,
                 pem: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "pem", pem)

    @property
    @pulumi.getter
    def pem(self) -> Optional[pulumi.Input[str]]:
        ...

    @pem.setter
    def pem(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class RegistryCredentialArgs:
    def __init__(__self__, *,
                 public_key_certificate: pulumi.Input['RegistryCredentialPublicKeyCertificateArgs']):
        """
        :param pulumi.Input['RegistryCredentialPublicKeyCertificateArgs'] public_key_certificate: A public key certificate format and data.
        """
        pulumi.set(__self__, "publicKeyCertificate", public_key_certificate)

    @property
    @pulumi.getter(name="publicKeyCertificate")
    def public_key_certificate(self) -> pulumi.Input['RegistryCredentialPublicKeyCertificateArgs']:
        """
        A public key certificate format and data.
        """
        ...

    @public_key_certificate.setter
    def public_key_certificate(self, value: pulumi.Input['RegistryCredentialPublicKeyCertificateArgs']):
        ...


@pulumi.input_type
class RegistryCredentialPublicKeyCertificateArgs:
    def __init__(__self__, *,
                 certificate: pulumi.Input[str],
                 format: pulumi.Input[str]):
        """
        :param pulumi.Input[str] certificate: The certificate data.
        :param pulumi.Input[str] format: The field allows only `X509_CERTIFICATE_PEM`.
        """
        pulumi.set(__self__, "certificate", certificate)
        pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Input[str]:
        """
        The certificate data.
        """
        ...

    @certificate.setter
    def certificate(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def format(self) -> pulumi.Input[str]:
        """
        The field allows only `X509_CERTIFICATE_PEM`.
        """
        ...

    @format.setter
    def format(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class RegistryEventNotificationConfigItemArgs:
    def __init__(__self__, *,
                 pubsub_topic_name: pulumi.Input[str],
                 subfolder_matches: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] pubsub_topic_name: PubSub topic name to publish device events.
        :param pulumi.Input[str] subfolder_matches: If the subfolder name matches this string exactly, this
               configuration will be used. The string must not include the
               leading '/' character. If empty, all strings are matched. Empty
               value can only be used for the last `event_notification_configs`
               item.
        """
        pulumi.set(__self__, "pubsubTopicName", pubsub_topic_name)
        pulumi.set(__self__, "subfolderMatches", subfolder_matches)

    @property
    @pulumi.getter(name="pubsubTopicName")
    def pubsub_topic_name(self) -> pulumi.Input[str]:
        """
        PubSub topic name to publish device events.
        """
        ...

    @pubsub_topic_name.setter
    def pubsub_topic_name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="subfolderMatches")
    def subfolder_matches(self) -> Optional[pulumi.Input[str]]:
        """
        If the subfolder name matches this string exactly, this
        configuration will be used. The string must not include the
        leading '/' character. If empty, all strings are matched. Empty
        value can only be used for the last `event_notification_configs`
        item.
        """
        ...

    @subfolder_matches.setter
    def subfolder_matches(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class RegistryHttpConfigArgs:
    def __init__(__self__, *,
                 http_enabled_state: pulumi.Input[str]):
        """
        :param pulumi.Input[str] http_enabled_state: The field allows `HTTP_ENABLED` or `HTTP_DISABLED`.
        """
        pulumi.set(__self__, "httpEnabledState", http_enabled_state)

    @property
    @pulumi.getter(name="httpEnabledState")
    def http_enabled_state(self) -> pulumi.Input[str]:
        """
        The field allows `HTTP_ENABLED` or `HTTP_DISABLED`.
        """
        ...

    @http_enabled_state.setter
    def http_enabled_state(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class RegistryMqttConfigArgs:
    def __init__(__self__, *,
                 mqtt_enabled_state: pulumi.Input[str]):
        """
        :param pulumi.Input[str] mqtt_enabled_state: The field allows `MQTT_ENABLED` or `MQTT_DISABLED`.
        """
        pulumi.set(__self__, "mqttEnabledState", mqtt_enabled_state)

    @property
    @pulumi.getter(name="mqttEnabledState")
    def mqtt_enabled_state(self) -> pulumi.Input[str]:
        """
        The field allows `MQTT_ENABLED` or `MQTT_DISABLED`.
        """
        ...

    @mqtt_enabled_state.setter
    def mqtt_enabled_state(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class RegistryStateNotificationConfigArgs:
    def __init__(__self__, *,
                 pubsub_topic_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] pubsub_topic_name: PubSub topic name to publish device events.
        """
        pulumi.set(__self__, "pubsubTopicName", pubsub_topic_name)

    @property
    @pulumi.getter(name="pubsubTopicName")
    def pubsub_topic_name(self) -> pulumi.Input[str]:
        """
        PubSub topic name to publish device events.
        """
        ...

    @pubsub_topic_name.setter
    def pubsub_topic_name(self, value: pulumi.Input[str]):
        ...


