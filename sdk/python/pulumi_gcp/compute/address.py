# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Address(pulumi.CustomResource):
    """
    Represents an Address resource.
    
    Each virtual machine instance has an ephemeral internal IP address and,
    optionally, an external IP address. To communicate between instances on
    the same network, you can use an instance's internal IP address. To
    communicate with the Internet and instances outside of the same network,
    you must specify the instance's external IP address.
    
    Internal IP addresses are ephemeral and only belong to an instance for
    the lifetime of the instance; if the instance is deleted and recreated,
    the instance is assigned a new internal IP address, either by Compute
    Engine or by you. External IP addresses can be either ephemeral or
    static.
    
    
    To get more information about Address, see:
    
    * [API documentation](https://cloud.google.com/compute/docs/reference/beta/addresses)
    * How-to Guides
        * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/instances-and-network)
        * [Reserving a Static Internal IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)
    """
    def __init__(__self__, __name__, __opts__=None, address=None, address_type=None, description=None, labels=None, name=None, network_tier=None, project=None, region=None, subnetwork=None):
        """Create a Address resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['address'] = address

        __props__['address_type'] = address_type

        __props__['description'] = description

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['network_tier'] = network_tier

        __props__['project'] = project

        __props__['region'] = region

        __props__['subnetwork'] = subnetwork

        __props__['creation_timestamp'] = None
        __props__['label_fingerprint'] = None
        __props__['self_link'] = None
        __props__['users'] = None

        super(Address, __self__).__init__(
            'gcp:compute/address:Address',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

