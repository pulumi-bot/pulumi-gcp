# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Subscription(pulumi.CustomResource):
    ack_deadline_seconds: pulumi.Output[float]
    """
    The maximum number of seconds a
    subscriber has to acknowledge a received message, otherwise the message is
    redelivered. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    A unique name for the resource, required by pubsub.
    Changing this forces a new resource to be created.
    """
    path: pulumi.Output[str]
    """
    Path of the subscription in the format `projects/{project}/subscriptions/{sub}`
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    push_config: pulumi.Output[dict]
    """
    Block configuration for push options. More
    configuration options are detailed below.
    """
    topic: pulumi.Output[str]
    """
    The topic name or id to bind this subscription to, required by pubsub.
    Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, ack_deadline_seconds=None, name=None, project=None, push_config=None, topic=None, __name__=None, __opts__=None):
        """
        Creates a subscription in Google's pubsub queueing system. For more information see
        [the official documentation](https://cloud.google.com/pubsub/docs) and
        [API](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] ack_deadline_seconds: The maximum number of seconds a
               subscriber has to acknowledge a received message, otherwise the message is
               redelivered. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: A unique name for the resource, required by pubsub.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[dict] push_config: Block configuration for push options. More
               configuration options are detailed below.
        :param pulumi.Input[str] topic: The topic name or id to bind this subscription to, required by pubsub.
               Changing this forces a new resource to be created.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['ack_deadline_seconds'] = ack_deadline_seconds

        __props__['name'] = name

        __props__['project'] = project

        __props__['push_config'] = push_config

        if topic is None:
            raise TypeError('Missing required property topic')
        __props__['topic'] = topic

        __props__['path'] = None

        super(Subscription, __self__).__init__(
            'gcp:pubsub/subscription:Subscription',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

