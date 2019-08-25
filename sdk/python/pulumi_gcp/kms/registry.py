# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
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
    def __init__(__self__, resource_name, opts=None, credentials=None, event_notification_config=None, http_config=None, mqtt_config=None, name=None, project=None, region=None, state_notification_config=None, __props__=None, __name__=None, __opts__=None):
        """
         Creates a device registry in Google's Cloud IoT Core platform. For more information see
        [the official documentation](https://cloud.google.com/iot/docs/) and
        [API](https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] credentials: List of public key certificates to authenticate devices. Structure is documented below. 
        :param pulumi.Input[dict] event_notification_config: A PubSub topics to publish device events. Structure is documented below.
        :param pulumi.Input[dict] http_config: Activate or deactivate HTTP. Structure is documented below.
        :param pulumi.Input[dict] mqtt_config: Activate or deactivate MQTT. Structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Region in which the created address should reside. If it is not provided, the provider region is used.
        :param pulumi.Input[dict] state_notification_config: A PubSub topic to publish device state updates. Structure is documented below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudiot_registry.html.markdown.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
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
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, credentials=None, event_notification_config=None, http_config=None, mqtt_config=None, name=None, project=None, region=None, state_notification_config=None):
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] credentials: List of public key certificates to authenticate devices. Structure is documented below. 
        :param pulumi.Input[dict] event_notification_config: A PubSub topics to publish device events. Structure is documented below.
        :param pulumi.Input[dict] http_config: Activate or deactivate HTTP. Structure is documented below.
        :param pulumi.Input[dict] mqtt_config: Activate or deactivate MQTT. Structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Region in which the created address should reside. If it is not provided, the provider region is used.
        :param pulumi.Input[dict] state_notification_config: A PubSub topic to publish device state updates. Structure is documented below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudiot_registry.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["credentials"] = credentials
        __props__["event_notification_config"] = event_notification_config
        __props__["http_config"] = http_config
        __props__["mqtt_config"] = mqtt_config
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        __props__["state_notification_config"] = state_notification_config
        return Registry(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

