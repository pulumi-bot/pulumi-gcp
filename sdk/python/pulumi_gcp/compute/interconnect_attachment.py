# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class InterconnectAttachment(pulumi.CustomResource):
    cloud_router_ip_address: pulumi.Output[str]
    creation_timestamp: pulumi.Output[str]
    customer_router_ip_address: pulumi.Output[str]
    description: pulumi.Output[str]
    google_reference_id: pulumi.Output[str]
    interconnect: pulumi.Output[str]
    name: pulumi.Output[str]
    private_interconnect_info: pulumi.Output[dict]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    router: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, interconnect=None, name=None, project=None, region=None, router=None):
        """
        Represents an InterconnectAttachment (VLAN attachment) resource. For more
        information, see Creating VLAN Attachments.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] description
        :param pulumi.Input[str] interconnect
        :param pulumi.Input[str] name
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region
        :param pulumi.Input[str] router
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        if interconnect is None:
            raise TypeError('Missing required property interconnect')
        __props__['interconnect'] = interconnect

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        if router is None:
            raise TypeError('Missing required property router')
        __props__['router'] = router

        __props__['cloud_router_ip_address'] = None
        __props__['creation_timestamp'] = None
        __props__['customer_router_ip_address'] = None
        __props__['google_reference_id'] = None
        __props__['private_interconnect_info'] = None
        __props__['self_link'] = None

        super(InterconnectAttachment, __self__).__init__(
            'gcp:compute/interconnectAttachment:InterconnectAttachment',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

