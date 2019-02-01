# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Registry(pulumi.CustomResource):
    credentials: pulumi.Output[list]
    """
    List of public key certificates to authenticate devices. Structure is documented below. 
    """
    event_notification_config: pulumi.Output[dict]
    """
    A PubSub topics to publish device events. Structure is documented below.
    """
    http_config: pulumi.Output[dict]
    """
    Activate or deactivate HTTP. Structure is documented below.
    """
    mqtt_config: pulumi.Output[dict]
    """
    Activate or deactivate MQTT. Structure is documented below.
    """
    name: pulumi.Output[str]
    """
    A unique name for the resource, required by device registry.
    Changing this forces a new resource to be created.
    """
    project: pulumi.Output[str]
    """
    The project in which the resource belongs. If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The Region in which the created address should reside. If it is not provided, the provider region is used.
    """
    state_notification_config: pulumi.Output[dict]
    """
    A PubSub topic to publish device state updates. Structure is documented below.
    """
    def __init__(__self__, __name__, __opts__=None, credentials=None, event_notification_config=None, http_config=None, mqtt_config=None, name=None, project=None, region=None, state_notification_config=None):
        """
         Creates a device registry in Google's Cloud IoT Core platform. For more information see
        [the official documentation](https://cloud.google.com/iot/docs/) and
        [API](https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries).
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] credentials: List of public key certificates to authenticate devices. Structure is documented below. 
        :param pulumi.Input[dict] event_notification_config: A PubSub topics to publish device events. Structure is documented below.
        :param pulumi.Input[dict] http_config: Activate or deactivate HTTP. Structure is documented below.
        :param pulumi.Input[dict] mqtt_config: Activate or deactivate MQTT. Structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Region in which the created address should reside. If it is not provided, the provider region is used.
        :param pulumi.Input[dict] state_notification_config: A PubSub topic to publish device state updates. Structure is documented below.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['credentials'] = credentials

        __props__['event_notification_config'] = event_notification_config

        __props__['http_config'] = http_config

        __props__['mqtt_config'] = mqtt_config

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        __props__['state_notification_config'] = state_notification_config

        super(Registry, __self__).__init__(
            'gcp:kms/registry:Registry',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

