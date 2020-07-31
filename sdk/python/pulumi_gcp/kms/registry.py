# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Registry']

warnings.warn("gcp.kms.Registry has been deprecated in favor of gcp.iot.Registry", DeprecationWarning)


class Registry(pulumi.CustomResource):
    credentials: pulumi.Output[Optional[List['outputs.RegistryCredential']]] = pulumi.output_property("credentials")
    """
    List of public key certificates to authenticate devices.
    The structure is documented below.
    """
    event_notification_configs: pulumi.Output[List['outputs.RegistryEventNotificationConfigItem']] = pulumi.output_property("eventNotificationConfigs")
    """
    List of configurations for event notifications, such as PubSub topics
    to publish device events to.  Structure is documented below.
    """
    http_config: pulumi.Output['outputs.RegistryHttpConfig'] = pulumi.output_property("httpConfig")
    """
    Activate or deactivate HTTP.
    The structure is documented below.
    """
    log_level: pulumi.Output[Optional[str]] = pulumi.output_property("logLevel")
    """
    The default logging verbosity for activity from devices in this
    registry. Specifies which events should be written to logs. For
    example, if the LogLevel is ERROR, only events that terminate in
    errors will be logged. LogLevel is inclusive; enabling INFO logging
    will also enable ERROR logging.
    """
    mqtt_config: pulumi.Output['outputs.RegistryMqttConfig'] = pulumi.output_property("mqttConfig")
    """
    Activate or deactivate MQTT.
    The structure is documented below.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    A unique name for the resource, required by device registry.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str] = pulumi.output_property("region")
    """
    The region in which the created registry should reside.
    If it is not provided, the provider region is used.
    """
    state_notification_config: pulumi.Output[Optional['outputs.RegistryStateNotificationConfig']] = pulumi.output_property("stateNotificationConfig")
    """
    A PubSub topic to publish device state updates.
    The structure is documented below.
    """
    warnings.warn("gcp.kms.Registry has been deprecated in favor of gcp.iot.Registry", DeprecationWarning)

    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, credentials: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]]] = None, event_notification_configs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]]] = None, http_config: Optional[pulumi.Input[pulumi.InputType['RegistryHttpConfigArgs']]] = None, log_level: Optional[pulumi.Input[str]] = None, mqtt_config: Optional[pulumi.Input[pulumi.InputType['RegistryMqttConfigArgs']]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, region: Optional[pulumi.Input[str]] = None, state_notification_config: Optional[pulumi.Input[pulumi.InputType['RegistryStateNotificationConfigArgs']]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A Google Cloud IoT Core device registry.

        To get more information about DeviceRegistry, see:

        * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/iot/docs/)

        ## Example Usage
        ### Cloudiot Device Registry Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test_registry = gcp.iot.Registry("test-registry")
        ```
        ### Cloudiot Device Registry Single Event Notification Configs

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_telemetry = gcp.pubsub.Topic("default-telemetry")
        test_registry = gcp.iot.Registry("test-registry", event_notification_configs=[{
            "pubsub_topic_name": default_telemetry.id,
            "subfolderMatches": "",
        }])
        ```
        ### Cloudiot Device Registry Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_devicestatus = gcp.pubsub.Topic("default-devicestatus")
        default_telemetry = gcp.pubsub.Topic("default-telemetry")
        additional_telemetry = gcp.pubsub.Topic("additional-telemetry")
        test_registry = gcp.iot.Registry("test-registry",
            event_notification_configs=[
                {
                    "pubsub_topic_name": additional_telemetry.id,
                    "subfolderMatches": "test/path",
                },
                {
                    "pubsub_topic_name": default_telemetry.id,
                    "subfolderMatches": "",
                },
            ],
            state_notification_config={
                "pubsub_topic_name": default_devicestatus.id,
            },
            mqtt_config={
                "mqtt_enabled_state": "MQTT_ENABLED",
            },
            http_config={
                "http_enabled_state": "HTTP_ENABLED",
            },
            log_level="INFO",
            credentials=[{
                "publicKeyCertificate": {
                    "format": "X509_CERTIFICATE_PEM",
                    "certificate": (lambda path: open(path).read())("test-fixtures/rsa_cert.pem"),
                },
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]] credentials: List of public key certificates to authenticate devices.
               The structure is documented below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]] event_notification_configs: List of configurations for event notifications, such as PubSub topics
               to publish device events to.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['RegistryHttpConfigArgs']] http_config: Activate or deactivate HTTP.
               The structure is documented below.
        :param pulumi.Input[str] log_level: The default logging verbosity for activity from devices in this
               registry. Specifies which events should be written to logs. For
               example, if the LogLevel is ERROR, only events that terminate in
               errors will be logged. LogLevel is inclusive; enabling INFO logging
               will also enable ERROR logging.
        :param pulumi.Input[pulumi.InputType['RegistryMqttConfigArgs']] mqtt_config: Activate or deactivate MQTT.
               The structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[pulumi.InputType['RegistryStateNotificationConfigArgs']] state_notification_config: A PubSub topic to publish device state updates.
               The structure is documented below.
        """
        pulumi.log.warn("Registry is deprecated: gcp.kms.Registry has been deprecated in favor of gcp.iot.Registry")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['credentials'] = credentials
            __props__['event_notification_configs'] = event_notification_configs
            __props__['http_config'] = http_config
            __props__['log_level'] = log_level
            __props__['mqtt_config'] = mqtt_config
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
            __props__['state_notification_config'] = state_notification_config
        super(Registry, __self__).__init__(
            'gcp:kms/registry:Registry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, credentials: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]]] = None, event_notification_configs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]]] = None, http_config: Optional[pulumi.Input[pulumi.InputType['RegistryHttpConfigArgs']]] = None, log_level: Optional[pulumi.Input[str]] = None, mqtt_config: Optional[pulumi.Input[pulumi.InputType['RegistryMqttConfigArgs']]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, region: Optional[pulumi.Input[str]] = None, state_notification_config: Optional[pulumi.Input[pulumi.InputType['RegistryStateNotificationConfigArgs']]] = None) -> 'Registry':
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]] credentials: List of public key certificates to authenticate devices.
               The structure is documented below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]] event_notification_configs: List of configurations for event notifications, such as PubSub topics
               to publish device events to.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['RegistryHttpConfigArgs']] http_config: Activate or deactivate HTTP.
               The structure is documented below.
        :param pulumi.Input[str] log_level: The default logging verbosity for activity from devices in this
               registry. Specifies which events should be written to logs. For
               example, if the LogLevel is ERROR, only events that terminate in
               errors will be logged. LogLevel is inclusive; enabling INFO logging
               will also enable ERROR logging.
        :param pulumi.Input[pulumi.InputType['RegistryMqttConfigArgs']] mqtt_config: Activate or deactivate MQTT.
               The structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[pulumi.InputType['RegistryStateNotificationConfigArgs']] state_notification_config: A PubSub topic to publish device state updates.
               The structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["credentials"] = credentials
        __props__["event_notification_configs"] = event_notification_configs
        __props__["http_config"] = http_config
        __props__["log_level"] = log_level
        __props__["mqtt_config"] = mqtt_config
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        __props__["state_notification_config"] = state_notification_config
        return Registry(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

