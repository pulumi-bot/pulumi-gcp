# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Autoscalar(pulumi.CustomResource):
    """
    #     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
    #
    # ----------------------------------------------------------------------------
    #
    #     This file is automatically generated by Magic Modules and manual
    #     changes will be clobbered when the file is regenerated.
    #
    #     Please read more about how to change this file in
    #     .github/CONTRIBUTING.md.
    #
    # ----------------------------------------------------------------------------
    layout: "google"
    page_title: "Google: google_compute_autoscaler"
    sidebar_current: "docs-google-compute-autoscaler"
    description: |-
      Represents an Autoscaler resource.
    ---
    
    # google\_compute\_autoscaler
    
    Represents an Autoscaler resource.
    
    Autoscalers allow you to automatically scale virtual machine instances in
    managed instance groups according to an autoscaling policy that you
    define.
    
    
    To get more information about Autoscaler, see:
    
    * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers)
    * How-to Guides
        * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
    """
    def __init__(__self__, __name__, __opts__=None, autoscaling_policy=None, description=None, name=None, project=None, target=None, zone=None):
        """Create a Autoscalar resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
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
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

