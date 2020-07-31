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

__all__ = ['Queue']


class Queue(pulumi.CustomResource):
    app_engine_routing_override: pulumi.Output[Optional['outputs.QueueAppEngineRoutingOverride']] = pulumi.output_property("appEngineRoutingOverride")
    """
    Overrides for task-level appEngineRouting. These settings apply only
    to App Engine tasks in this queue  Structure is documented below.
    """
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    The location of the queue
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The queue name.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    rate_limits: pulumi.Output['outputs.QueueRateLimits'] = pulumi.output_property("rateLimits")
    """
    Rate limits for task dispatches.
    The queue's actual dispatch rate is the result of:
    * Number of tasks in the queue
    * User-specified throttling: rateLimits, retryConfig, and the queue's state.
    * System throttling due to 429 (Too Many Requests) or 503 (Service
    Unavailable) responses from the worker, high error rates, or to
    smooth sudden large traffic spikes.  Structure is documented below.
    """
    retry_config: pulumi.Output['outputs.QueueRetryConfig'] = pulumi.output_property("retryConfig")
    """
    Settings that determine the retry behavior.  Structure is documented below.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, app_engine_routing_override: Optional[pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, rate_limits: Optional[pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']]] = None, retry_config: Optional[pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A named resource to which messages are sent by publishers.

        ## Example Usage
        ### Queue Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.cloudtasks.Queue("default", location="us-central1")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only
               to App Engine tasks in this queue  Structure is documented below.
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']] rate_limits: Rate limits for task dispatches.
               The queue's actual dispatch rate is the result of:
               * Number of tasks in the queue
               * User-specified throttling: rateLimits, retryConfig, and the queue's state.
               * System throttling due to 429 (Too Many Requests) or 503 (Service
               Unavailable) responses from the worker, high error rates, or to
               smooth sudden large traffic spikes.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']] retry_config: Settings that determine the retry behavior.  Structure is documented below.
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

            __props__['app_engine_routing_override'] = app_engine_routing_override
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['name'] = name
            __props__['project'] = project
            __props__['rate_limits'] = rate_limits
            __props__['retry_config'] = retry_config
        super(Queue, __self__).__init__(
            'gcp:cloudtasks/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, app_engine_routing_override: Optional[pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']]] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, rate_limits: Optional[pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']]] = None, retry_config: Optional[pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']]] = None) -> 'Queue':
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only
               to App Engine tasks in this queue  Structure is documented below.
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']] rate_limits: Rate limits for task dispatches.
               The queue's actual dispatch rate is the result of:
               * Number of tasks in the queue
               * User-specified throttling: rateLimits, retryConfig, and the queue's state.
               * System throttling due to 429 (Too Many Requests) or 503 (Service
               Unavailable) responses from the worker, high error rates, or to
               smooth sudden large traffic spikes.  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']] retry_config: Settings that determine the retry behavior.  Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_engine_routing_override"] = app_engine_routing_override
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["rate_limits"] = rate_limits
        __props__["retry_config"] = retry_config
        return Queue(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

