# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("gcp.kms.Registry has been deprecated in favor of gcp.iot.Registry", DeprecationWarning)
class Registry(pulumi.CustomResource):
    credentials: pulumi.Output[list]
    event_notification_configs: pulumi.Output[list]
    http_config: pulumi.Output[dict]
    log_level: pulumi.Output[str]
    mqtt_config: pulumi.Output[dict]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    region: pulumi.Output[str]
    state_notification_config: pulumi.Output[dict]
    warnings.warn("gcp.kms.Registry has been deprecated in favor of gcp.iot.Registry", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, credentials=None, event_notification_configs=None, http_config=None, log_level=None, mqtt_config=None, name=None, project=None, region=None, state_notification_config=None, __props__=None, __name__=None, __opts__=None):
        """

        Deprecated: gcp.kms.Registry has been deprecated in favor of gcp.iot.Registry

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **credentials** object supports the following:

          * `publicKeyCertificate` (`pulumi.Input[dict]`)
            * `certificate` (`pulumi.Input[str]`)
            * `format` (`pulumi.Input[str]`)

        The **event_notification_configs** object supports the following:

          * `pubsub_topic_name` (`pulumi.Input[str]`)
          * `subfolderMatches` (`pulumi.Input[str]`)

        The **http_config** object supports the following:

          * `http_enabled_state` (`pulumi.Input[str]`)

        The **mqtt_config** object supports the following:

          * `mqtt_enabled_state` (`pulumi.Input[str]`)

        The **state_notification_config** object supports the following:

          * `pubsub_topic_name` (`pulumi.Input[str]`)
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, credentials=None, event_notification_configs=None, http_config=None, log_level=None, mqtt_config=None, name=None, project=None, region=None, state_notification_config=None):
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **credentials** object supports the following:

          * `publicKeyCertificate` (`pulumi.Input[dict]`)
            * `certificate` (`pulumi.Input[str]`)
            * `format` (`pulumi.Input[str]`)

        The **event_notification_configs** object supports the following:

          * `pubsub_topic_name` (`pulumi.Input[str]`)
          * `subfolderMatches` (`pulumi.Input[str]`)

        The **http_config** object supports the following:

          * `http_enabled_state` (`pulumi.Input[str]`)

        The **mqtt_config** object supports the following:

          * `mqtt_enabled_state` (`pulumi.Input[str]`)

        The **state_notification_config** object supports the following:

          * `pubsub_topic_name` (`pulumi.Input[str]`)
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

