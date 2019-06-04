# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class UptimeCheckConfig(pulumi.CustomResource):
    content_matchers: pulumi.Output[list]
    display_name: pulumi.Output[str]
    http_check: pulumi.Output[dict]
    internal_checkers: pulumi.Output[list]
    is_internal: pulumi.Output[bool]
    monitored_resource: pulumi.Output[dict]
    name: pulumi.Output[str]
    period: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    resource_group: pulumi.Output[dict]
    selected_regions: pulumi.Output[list]
    tcp_check: pulumi.Output[dict]
    timeout: pulumi.Output[str]
    uptime_check_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, content_matchers=None, display_name=None, http_check=None, internal_checkers=None, is_internal=None, monitored_resource=None, period=None, project=None, resource_group=None, selected_regions=None, tcp_check=None, timeout=None, __name__=None, __opts__=None):
        """
        This message configures which resources and services to monitor for availability.
        
        
        To get more information about UptimeCheckConfig, see:
        
        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.uptimeCheckConfigs)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/monitoring/uptime-checks/)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

        __props__['content_matchers'] = content_matchers

        if display_name is None:
            raise TypeError("Missing required property 'display_name'")
        __props__['display_name'] = display_name

        __props__['http_check'] = http_check

        __props__['internal_checkers'] = internal_checkers

        __props__['is_internal'] = is_internal

        __props__['monitored_resource'] = monitored_resource

        __props__['period'] = period

        __props__['project'] = project

        __props__['resource_group'] = resource_group

        __props__['selected_regions'] = selected_regions

        __props__['tcp_check'] = tcp_check

        if timeout is None:
            raise TypeError("Missing required property 'timeout'")
        __props__['timeout'] = timeout

        __props__['name'] = None
        __props__['uptime_check_id'] = None

        super(UptimeCheckConfig, __self__).__init__(
            'gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

