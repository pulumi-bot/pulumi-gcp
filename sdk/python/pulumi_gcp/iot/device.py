# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Device(pulumi.CustomResource):
    blocked: pulumi.Output[bool]
    """
    If a device is blocked, connections or requests from this device will fail.
    """
    config: pulumi.Output[dict]
    """
    The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.

      * `binaryData` (`str`)
      * `cloudUpdateTime` (`str`)
      * `deviceAckTime` (`str`)
      * `version` (`str`)
    """
    credentials: pulumi.Output[list]
    """
    The credentials used to authenticate this device.  Structure is documented below.

      * `expiration_time` (`str`) - The time at which this credential becomes invalid.
      * `public_key` (`dict`) - A public key used to verify the signature of JSON Web Tokens (JWTs).  Structure is documented below.
        * `format` (`str`) - The format of the key.
        * `key` (`str`) - The key data.
    """
    gateway_config: pulumi.Output[dict]
    """
    Gateway-related configuration and state.  Structure is documented below.

      * `gatewayAuthMethod` (`str`) - Indicates whether the device is a gateway.
      * `gatewayType` (`str`) - Indicates whether the device is a gateway.
      * `lastAccessedGatewayId` (`str`) - -
        The ID of the gateway the device accessed most recently.
      * `lastAccessedGatewayTime` (`str`) - -
        The most recent time at which the device accessed the gateway specified in last_accessed_gateway.
    """
    last_config_ack_time: pulumi.Output[str]
    """
    The last time a cloud-to-device config version acknowledgment was received from the device.
    """
    last_config_send_time: pulumi.Output[str]
    """
    The last time a cloud-to-device config version was sent to the device.
    """
    last_error_status: pulumi.Output[dict]
    """
    The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.

      * `details` (`list`)
      * `message` (`str`)
      * `number` (`float`)
    """
    last_error_time: pulumi.Output[str]
    """
    The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
    """
    last_event_time: pulumi.Output[str]
    """
    The last time a telemetry event was received.
    """
    last_heartbeat_time: pulumi.Output[str]
    """
    The last time an MQTT PINGREQ was received.
    """
    last_state_time: pulumi.Output[str]
    """
    The last time a state event was received.
    """
    log_level: pulumi.Output[str]
    """
    The logging verbosity for device activity.
    """
    metadata: pulumi.Output[dict]
    """
    The metadata key-value pairs assigned to the device.
    """
    name: pulumi.Output[str]
    """
    A unique name for the resource.
    """
    num_id: pulumi.Output[str]
    """
    A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
    unique.
    """
    registry: pulumi.Output[str]
    """
    The name of the device registry where this device should be created.
    """
    state: pulumi.Output[dict]
    """
    The state most recently received from the device.

      * `binaryData` (`str`)
      * `update_time` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, blocked=None, credentials=None, gateway_config=None, log_level=None, metadata=None, name=None, registry=None, __props__=None, __name__=None, __opts__=None):
        """
        A Google Cloud IoT Core device.

        To get more information about Device, see:

        * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/iot/docs/)

        ## Example Usage
        ### Cloudiot Device Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        registry = gcp.iot.Registry("registry")
        test_device = gcp.iot.Device("test-device", registry=registry.id)
        ```
        ### Cloudiot Device Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        registry = gcp.iot.Registry("registry")
        test_device = gcp.iot.Device("test-device",
            registry=registry.id,
            credentials=[gcp.iot.DeviceCredentialArgs(
                public_key={
                    "format": "RSA_PEM",
                    "key": (lambda path: open(path).read())("test-fixtures/rsa_public.pem"),
                },
            )],
            blocked=False,
            log_level="INFO",
            metadata={
                "test_key_1": "test_value_1",
            },
            gateway_config=gcp.iot.DeviceGatewayConfigArgs(
                gateway_type="NON_GATEWAY",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocked: If a device is blocked, connections or requests from this device will fail.
        :param pulumi.Input[list] credentials: The credentials used to authenticate this device.  Structure is documented below.
        :param pulumi.Input[dict] gateway_config: Gateway-related configuration and state.  Structure is documented below.
        :param pulumi.Input[str] log_level: The logging verbosity for device activity.
        :param pulumi.Input[dict] metadata: The metadata key-value pairs assigned to the device.
        :param pulumi.Input[str] name: A unique name for the resource.
        :param pulumi.Input[str] registry: The name of the device registry where this device should be created.

        The **credentials** object supports the following:

          * `expiration_time` (`pulumi.Input[str]`) - The time at which this credential becomes invalid.
          * `public_key` (`pulumi.Input[dict]`) - A public key used to verify the signature of JSON Web Tokens (JWTs).  Structure is documented below.
            * `format` (`pulumi.Input[str]`) - The format of the key.
            * `key` (`pulumi.Input[str]`) - The key data.

        The **gateway_config** object supports the following:

          * `gatewayAuthMethod` (`pulumi.Input[str]`) - Indicates whether the device is a gateway.
          * `gatewayType` (`pulumi.Input[str]`) - Indicates whether the device is a gateway.
          * `lastAccessedGatewayId` (`pulumi.Input[str]`) - -
            The ID of the gateway the device accessed most recently.
          * `lastAccessedGatewayTime` (`pulumi.Input[str]`) - -
            The most recent time at which the device accessed the gateway specified in last_accessed_gateway.
        """
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

            __props__['blocked'] = blocked
            __props__['credentials'] = credentials
            __props__['gateway_config'] = gateway_config
            __props__['log_level'] = log_level
            __props__['metadata'] = metadata
            __props__['name'] = name
            if registry is None:
                raise TypeError("Missing required property 'registry'")
            __props__['registry'] = registry
            __props__['config'] = None
            __props__['last_config_ack_time'] = None
            __props__['last_config_send_time'] = None
            __props__['last_error_status'] = None
            __props__['last_error_time'] = None
            __props__['last_event_time'] = None
            __props__['last_heartbeat_time'] = None
            __props__['last_state_time'] = None
            __props__['num_id'] = None
            __props__['state'] = None
        super(Device, __self__).__init__(
            'gcp:iot/device:Device',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, blocked=None, config=None, credentials=None, gateway_config=None, last_config_ack_time=None, last_config_send_time=None, last_error_status=None, last_error_time=None, last_event_time=None, last_heartbeat_time=None, last_state_time=None, log_level=None, metadata=None, name=None, num_id=None, registry=None, state=None):
        """
        Get an existing Device resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] blocked: If a device is blocked, connections or requests from this device will fail.
        :param pulumi.Input[dict] config: The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
        :param pulumi.Input[list] credentials: The credentials used to authenticate this device.  Structure is documented below.
        :param pulumi.Input[dict] gateway_config: Gateway-related configuration and state.  Structure is documented below.
        :param pulumi.Input[str] last_config_ack_time: The last time a cloud-to-device config version acknowledgment was received from the device.
        :param pulumi.Input[str] last_config_send_time: The last time a cloud-to-device config version was sent to the device.
        :param pulumi.Input[dict] last_error_status: The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
        :param pulumi.Input[str] last_error_time: The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
        :param pulumi.Input[str] last_event_time: The last time a telemetry event was received.
        :param pulumi.Input[str] last_heartbeat_time: The last time an MQTT PINGREQ was received.
        :param pulumi.Input[str] last_state_time: The last time a state event was received.
        :param pulumi.Input[str] log_level: The logging verbosity for device activity.
        :param pulumi.Input[dict] metadata: The metadata key-value pairs assigned to the device.
        :param pulumi.Input[str] name: A unique name for the resource.
        :param pulumi.Input[str] num_id: A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
               unique.
        :param pulumi.Input[str] registry: The name of the device registry where this device should be created.
        :param pulumi.Input[dict] state: The state most recently received from the device.

        The **config** object supports the following:

          * `binaryData` (`pulumi.Input[str]`)
          * `cloudUpdateTime` (`pulumi.Input[str]`)
          * `deviceAckTime` (`pulumi.Input[str]`)
          * `version` (`pulumi.Input[str]`)

        The **credentials** object supports the following:

          * `expiration_time` (`pulumi.Input[str]`) - The time at which this credential becomes invalid.
          * `public_key` (`pulumi.Input[dict]`) - A public key used to verify the signature of JSON Web Tokens (JWTs).  Structure is documented below.
            * `format` (`pulumi.Input[str]`) - The format of the key.
            * `key` (`pulumi.Input[str]`) - The key data.

        The **gateway_config** object supports the following:

          * `gatewayAuthMethod` (`pulumi.Input[str]`) - Indicates whether the device is a gateway.
          * `gatewayType` (`pulumi.Input[str]`) - Indicates whether the device is a gateway.
          * `lastAccessedGatewayId` (`pulumi.Input[str]`) - -
            The ID of the gateway the device accessed most recently.
          * `lastAccessedGatewayTime` (`pulumi.Input[str]`) - -
            The most recent time at which the device accessed the gateway specified in last_accessed_gateway.

        The **last_error_status** object supports the following:

          * `details` (`pulumi.Input[list]`)
          * `message` (`pulumi.Input[str]`)
          * `number` (`pulumi.Input[float]`)

        The **state** object supports the following:

          * `binaryData` (`pulumi.Input[str]`)
          * `update_time` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["blocked"] = blocked
        __props__["config"] = config
        __props__["credentials"] = credentials
        __props__["gateway_config"] = gateway_config
        __props__["last_config_ack_time"] = last_config_ack_time
        __props__["last_config_send_time"] = last_config_send_time
        __props__["last_error_status"] = last_error_status
        __props__["last_error_time"] = last_error_time
        __props__["last_event_time"] = last_event_time
        __props__["last_heartbeat_time"] = last_heartbeat_time
        __props__["last_state_time"] = last_state_time
        __props__["log_level"] = log_level
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["num_id"] = num_id
        __props__["registry"] = registry
        __props__["state"] = state
        return Device(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
