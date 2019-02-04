# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Firewall(pulumi.CustomResource):
    allows: pulumi.Output[list]
    creation_timestamp: pulumi.Output[str]
    denies: pulumi.Output[list]
    description: pulumi.Output[str]
    destination_ranges: pulumi.Output[list]
    direction: pulumi.Output[str]
    disabled: pulumi.Output[bool]
    enable_logging: pulumi.Output[bool]
    name: pulumi.Output[str]
    network: pulumi.Output[str]
    priority: pulumi.Output[int]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    source_ranges: pulumi.Output[list]
    source_service_accounts: pulumi.Output[str]
    source_tags: pulumi.Output[list]
    target_service_accounts: pulumi.Output[str]
    target_tags: pulumi.Output[list]
    def __init__(__self__, __name__, __opts__=None, allows=None, denies=None, description=None, destination_ranges=None, direction=None, disabled=None, enable_logging=None, name=None, network=None, priority=None, project=None, source_ranges=None, source_service_accounts=None, source_tags=None, target_service_accounts=None, target_tags=None):
        """
        Each network has its own firewall controlling access to and from the
        instances.
        
        All traffic to instances, even from other instances, is blocked by the
        firewall unless firewall rules are created to allow it.
        
        The default network has automatically created firewall rules that are
        shown in default firewall rules. No manually created network has
        automatically created firewall rules except for a default "allow" rule for
        outgoing traffic and a default "deny" for incoming traffic. For all
        networks except the default network, you must create any firewall rules
        you need.
        
        
        To get more information about Firewall, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/latest/firewalls)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/vpc/docs/firewalls)
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=firewall_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] allows
        :param pulumi.Input[list] denies
        :param pulumi.Input[str] description
        :param pulumi.Input[list] destination_ranges
        :param pulumi.Input[str] direction
        :param pulumi.Input[bool] disabled
        :param pulumi.Input[bool] enable_logging
        :param pulumi.Input[str] name
        :param pulumi.Input[str] network
        :param pulumi.Input[int] priority
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] source_ranges
        :param pulumi.Input[str] source_service_accounts
        :param pulumi.Input[list] source_tags
        :param pulumi.Input[str] target_service_accounts
        :param pulumi.Input[list] target_tags
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allows'] = allows

        __props__['denies'] = denies

        __props__['description'] = description

        __props__['destination_ranges'] = destination_ranges

        __props__['direction'] = direction

        __props__['disabled'] = disabled

        __props__['enable_logging'] = enable_logging

        __props__['name'] = name

        if network is None:
            raise TypeError('Missing required property network')
        __props__['network'] = network

        __props__['priority'] = priority

        __props__['project'] = project

        __props__['source_ranges'] = source_ranges

        __props__['source_service_accounts'] = source_service_accounts

        __props__['source_tags'] = source_tags

        __props__['target_service_accounts'] = target_service_accounts

        __props__['target_tags'] = target_tags

        __props__['creation_timestamp'] = None
        __props__['self_link'] = None

        super(Firewall, __self__).__init__(
            'gcp:compute/firewall:Firewall',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

