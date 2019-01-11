# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Network(pulumi.CustomResource):
    """
    Manages a network within GCE. For more information see
    [the official documentation](https://cloud.google.com/compute/docs/vpc)
    and
    [API](https://cloud.google.com/compute/docs/reference/latest/networks).
    """
    def __init__(__self__, __name__, __opts__=None, auto_create_subnetworks=None, description=None, ipv4_range=None, name=None, project=None, routing_mode=None):
        """Create a Network resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['auto_create_subnetworks'] = auto_create_subnetworks

        __props__['description'] = description

        __props__['ipv4_range'] = ipv4_range

        __props__['name'] = name

        __props__['project'] = project

        __props__['routing_mode'] = routing_mode

        __props__['gateway_ipv4'] = None
        __props__['self_link'] = None

        super(Network, __self__).__init__(
            'gcp:compute/network:Network',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

