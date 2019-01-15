# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GlobalAddress(pulumi.CustomResource):
    address: pulumi.Output[str]
    address_type: pulumi.Output[str]
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    ip_version: pulumi.Output[str]
    label_fingerprint: pulumi.Output[str]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    network: pulumi.Output[str]
    prefix_length: pulumi.Output[int]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    purpose: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, __name__, __opts__=None, address_type=None, description=None, ip_version=None, labels=None, name=None, network=None, prefix_length=None, project=None, purpose=None):
        """
        Represents a Global Address resource. Global addresses are used for
        HTTP(S) load balancing.
        
        
        To get more information about GlobalAddress, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/latest/globalAddresses)
        * How-to Guides
            * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address)
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=global_address_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] address_type
        :param pulumi.Input[str] description
        :param pulumi.Input[str] ip_version
        :param pulumi.Input[dict] labels
        :param pulumi.Input[str] name
        :param pulumi.Input[str] network
        :param pulumi.Input[int] prefix_length
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['address_type'] = address_type

        __props__['description'] = description

        __props__['ip_version'] = ip_version

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['network'] = network

        __props__['prefix_length'] = prefix_length

        __props__['project'] = project

        __props__['purpose'] = purpose

        __props__['address'] = None
        __props__['creation_timestamp'] = None
        __props__['label_fingerprint'] = None
        __props__['self_link'] = None

        super(GlobalAddress, __self__).__init__(
            'gcp:compute/globalAddress:GlobalAddress',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

