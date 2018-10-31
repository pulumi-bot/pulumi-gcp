# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GlobalAddress(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, address_type=None, description=None, ip_version=None, labels=None, name=None, network=None, prefix_length=None, project=None, purpose=None):
        """Create a GlobalAddress resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['addressType'] = address_type

        __props__['description'] = description

        __props__['ipVersion'] = ip_version

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['network'] = network

        __props__['prefixLength'] = prefix_length

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

