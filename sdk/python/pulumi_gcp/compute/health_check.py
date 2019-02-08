# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class HealthCheck(pulumi.CustomResource):
    check_interval_sec: pulumi.Output[int]
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    healthy_threshold: pulumi.Output[int]
    http_health_check: pulumi.Output[dict]
    https_health_check: pulumi.Output[dict]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    ssl_health_check: pulumi.Output[dict]
    tcp_health_check: pulumi.Output[dict]
    timeout_sec: pulumi.Output[int]
    type: pulumi.Output[str]
    unhealthy_threshold: pulumi.Output[int]
    def __init__(__self__, resource_name, opts=None, check_interval_sec=None, description=None, healthy_threshold=None, http_health_check=None, https_health_check=None, name=None, project=None, ssl_health_check=None, tcp_health_check=None, timeout_sec=None, unhealthy_threshold=None, __name__=None, __opts__=None):
        """
        Health Checks determine whether instances are responsive and able to do work.
        They are an important part of a comprehensive load balancing configuration,
        as they enable monitoring instances behind load balancers.
        
        Health Checks poll instances at a specified interval. Instances that
        do not respond successfully to some number of probes in a row are marked
        as unhealthy. No new connections are sent to unhealthy instances,
        though existing connections will continue. The health check will
        continue to poll unhealthy instances. If an instance later responds
        successfully to some number of consecutive probes, it is marked
        healthy again and can receive new connections.
        
        
        To get more information about HealthCheck, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/latest/healthChecks)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/load-balancing/docs/health-checks)
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=health_check_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] check_interval_sec
        :param pulumi.Input[str] description
        :param pulumi.Input[int] healthy_threshold
        :param pulumi.Input[dict] http_health_check
        :param pulumi.Input[dict] https_health_check
        :param pulumi.Input[str] name
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] ssl_health_check
        :param pulumi.Input[dict] tcp_health_check
        :param pulumi.Input[int] timeout_sec
        :param pulumi.Input[int] unhealthy_threshold
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

        __props__['check_interval_sec'] = check_interval_sec

        __props__['description'] = description

        __props__['healthy_threshold'] = healthy_threshold

        __props__['http_health_check'] = http_health_check

        __props__['https_health_check'] = https_health_check

        __props__['name'] = name

        __props__['project'] = project

        __props__['ssl_health_check'] = ssl_health_check

        __props__['tcp_health_check'] = tcp_health_check

        __props__['timeout_sec'] = timeout_sec

        __props__['unhealthy_threshold'] = unhealthy_threshold

        __props__['creation_timestamp'] = None
        __props__['self_link'] = None
        __props__['type'] = None

        super(HealthCheck, __self__).__init__(
            'gcp:compute/healthCheck:HealthCheck',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

