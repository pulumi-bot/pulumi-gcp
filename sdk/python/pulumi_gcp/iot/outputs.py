# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'DeviceConfig',
    'DeviceCredential',
    'DeviceCredentialPublicKey',
    'DeviceGatewayConfig',
    'DeviceLastErrorStatus',
    'DeviceState',
    'RegistryCredential',
    'RegistryCredentialPublicKeyCertificate',
    'RegistryEventNotificationConfigItem',
    'RegistryHttpConfig',
    'RegistryMqttConfig',
    'RegistryStateNotificationConfig',
]

@pulumi.output_type
class DeviceConfig(dict):
    binary_data: Optional[str] = pulumi.output_property("binaryData")
    cloud_update_time: Optional[str] = pulumi.output_property("cloudUpdateTime")
    device_ack_time: Optional[str] = pulumi.output_property("deviceAckTime")
    version: Optional[str] = pulumi.output_property("version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DeviceCredential(dict):
    expiration_time: Optional[str] = pulumi.output_property("expirationTime")
    """
    The time at which this credential becomes invalid.
    """
    public_key: 'outputs.DeviceCredentialPublicKey' = pulumi.output_property("publicKey")
    """
    A public key used to verify the signature of JSON Web Tokens (JWTs).  Structure is documented below.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DeviceCredentialPublicKey(dict):
    format: str = pulumi.output_property("format")
    """
    The format of the key.
    """
    key: str = pulumi.output_property("key")
    """
    The key data.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DeviceGatewayConfig(dict):
    gateway_auth_method: Optional[str] = pulumi.output_property("gatewayAuthMethod")
    """
    Indicates whether the device is a gateway.
    """
    gateway_type: Optional[str] = pulumi.output_property("gatewayType")
    """
    Indicates whether the device is a gateway.
    """
    last_accessed_gateway_id: Optional[str] = pulumi.output_property("lastAccessedGatewayId")
    """
    -
    The ID of the gateway the device accessed most recently.
    """
    last_accessed_gateway_time: Optional[str] = pulumi.output_property("lastAccessedGatewayTime")
    """
    -
    The most recent time at which the device accessed the gateway specified in last_accessed_gateway.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DeviceLastErrorStatus(dict):
    details: Optional[List[Dict[str, Any]]] = pulumi.output_property("details")
    message: Optional[str] = pulumi.output_property("message")
    number: Optional[float] = pulumi.output_property("number")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DeviceState(dict):
    binary_data: Optional[str] = pulumi.output_property("binaryData")
    update_time: Optional[str] = pulumi.output_property("updateTime")

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


