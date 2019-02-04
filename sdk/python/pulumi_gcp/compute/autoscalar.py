# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Autoscalar(pulumi.CustomResource):
    autoscaling_policy: pulumi.Output[dict]
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    target: pulumi.Output[str]
    zone: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, autoscaling_policy=None, description=None, name=None, project=None, target=None, zone=None, __name__=None, __opts__=None):
        """
        Represents an Autoscaler resource.
        
        Autoscalers allow you to automatically scale virtual machine instances in
        managed instance groups according to an autoscaling policy that you
        define.
        
        
        To get more information about Autoscaler, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers)
        * How-to Guides
            * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=autoscaler_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] autoscaling_policy
        :param pulumi.Input[str] description
        :param pulumi.Input[str] name
        :param pulumi.Input[str] project
        :param pulumi.Input[str] target
        :param pulumi.Input[str] zone
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

        if not autoscaling_policy:
            raise TypeError('Missing required property autoscaling_policy')
        __props__['autoscaling_policy'] = autoscaling_policy

        __props__['description'] = description

        __props__['name'] = name

        __props__['project'] = project

        if not target:
            raise TypeError('Missing required property target')
        __props__['target'] = target

        __props__['zone'] = zone

        __props__['creation_timestamp'] = None
        __props__['self_link'] = None

        super(Autoscalar, __self__).__init__(
            'gcp:compute/autoscalar:Autoscalar',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

