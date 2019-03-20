# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    alternative_location_id: pulumi.Output[str]
    authorized_network: pulumi.Output[str]
    create_time: pulumi.Output[str]
    current_location_id: pulumi.Output[str]
    display_name: pulumi.Output[str]
    host: pulumi.Output[str]
    labels: pulumi.Output[dict]
    location_id: pulumi.Output[str]
    memory_size_gb: pulumi.Output[float]
    name: pulumi.Output[str]
    port: pulumi.Output[float]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    redis_configs: pulumi.Output[dict]
    redis_version: pulumi.Output[str]
    region: pulumi.Output[str]
    reserved_ip_range: pulumi.Output[str]
    tier: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, alternative_location_id=None, authorized_network=None, display_name=None, labels=None, location_id=None, memory_size_gb=None, name=None, project=None, redis_configs=None, redis_version=None, region=None, reserved_ip_range=None, tier=None, __name__=None, __opts__=None):
        """
        A Google Cloud Redis instance.
        
        
        To get more information about Instance, see:
        
        * [API documentation](https://cloud.google.com/memorystore/docs/redis/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/memorystore/docs/redis/)
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=redis_instance_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
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

        __props__['alternative_location_id'] = alternative_location_id

        __props__['authorized_network'] = authorized_network

        __props__['display_name'] = display_name

        __props__['labels'] = labels

        __props__['location_id'] = location_id

        if memory_size_gb is None:
            raise TypeError('Missing required property memory_size_gb')
        __props__['memory_size_gb'] = memory_size_gb

        __props__['name'] = name

        __props__['project'] = project

        __props__['redis_configs'] = redis_configs

        __props__['redis_version'] = redis_version

        __props__['region'] = region

        __props__['reserved_ip_range'] = reserved_ip_range

        __props__['tier'] = tier

        __props__['create_time'] = None
        __props__['current_location_id'] = None
        __props__['host'] = None
        __props__['port'] = None

        super(Instance, __self__).__init__(
            'gcp:redis/instance:Instance',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

