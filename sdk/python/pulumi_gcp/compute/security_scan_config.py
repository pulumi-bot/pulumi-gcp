# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecurityScanConfig(pulumi.CustomResource):
    authentication: pulumi.Output[dict]
    blacklist_patterns: pulumi.Output[list]
    display_name: pulumi.Output[str]
    export_to_security_command_center: pulumi.Output[str]
    max_qps: pulumi.Output[float]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    schedule: pulumi.Output[dict]
    starting_urls: pulumi.Output[list]
    target_platforms: pulumi.Output[list]
    user_agent: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, authentication=None, blacklist_patterns=None, display_name=None, export_to_security_command_center=None, max_qps=None, project=None, schedule=None, starting_urls=None, target_platforms=None, user_agent=None, __props__=None, __name__=None, __opts__=None):
        """
        A ScanConfig resource contains the configurations to launch a scan.
        
        To get more information about ScanConfig, see:
        
        * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
        * How-to Guides
            * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **authentication** object supports the following:
        
          * `custom_account` (`pulumi.Input[dict]`)
        
            * `login_url` (`pulumi.Input[str]`)
            * `password` (`pulumi.Input[str]`)
            * `username` (`pulumi.Input[str]`)
        
          * `google_account` (`pulumi.Input[dict]`)
        
            * `password` (`pulumi.Input[str]`)
            * `username` (`pulumi.Input[str]`)
        
        The **schedule** object supports the following:
        
          * `interval_duration_days` (`pulumi.Input[float]`)
          * `schedule_time` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/security_scanner_scan_config.html.markdown.
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

            __props__['authentication'] = authentication
            __props__['blacklist_patterns'] = blacklist_patterns
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['export_to_security_command_center'] = export_to_security_command_center
            __props__['max_qps'] = max_qps
            __props__['project'] = project
            __props__['schedule'] = schedule
            if starting_urls is None:
                raise TypeError("Missing required property 'starting_urls'")
            __props__['starting_urls'] = starting_urls
            __props__['target_platforms'] = target_platforms
            __props__['user_agent'] = user_agent
            __props__['name'] = None
        super(SecurityScanConfig, __self__).__init__(
            'gcp:compute/securityScanConfig:SecurityScanConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, authentication=None, blacklist_patterns=None, display_name=None, export_to_security_command_center=None, max_qps=None, name=None, project=None, schedule=None, starting_urls=None, target_platforms=None, user_agent=None):
        """
        Get an existing SecurityScanConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **authentication** object supports the following:
        
          * `custom_account` (`pulumi.Input[dict]`)
        
            * `login_url` (`pulumi.Input[str]`)
            * `password` (`pulumi.Input[str]`)
            * `username` (`pulumi.Input[str]`)
        
          * `google_account` (`pulumi.Input[dict]`)
        
            * `password` (`pulumi.Input[str]`)
            * `username` (`pulumi.Input[str]`)
        
        The **schedule** object supports the following:
        
          * `interval_duration_days` (`pulumi.Input[float]`)
          * `schedule_time` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/security_scanner_scan_config.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["authentication"] = authentication
        __props__["blacklist_patterns"] = blacklist_patterns
        __props__["display_name"] = display_name
        __props__["export_to_security_command_center"] = export_to_security_command_center
        __props__["max_qps"] = max_qps
        __props__["name"] = name
        __props__["project"] = project
        __props__["schedule"] = schedule
        __props__["starting_urls"] = starting_urls
        __props__["target_platforms"] = target_platforms
        __props__["user_agent"] = user_agent
        return SecurityScanConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

