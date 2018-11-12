# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class ManagedZone(pulumi.CustomResource):
    """
    Manages a zone within Google Cloud DNS. For more information see [the official documentation](https://cloud.google.com/dns/zones/) and
    [API](https://cloud.google.com/dns/api/v1/managedZones).
    """
    def __init__(__self__, __name__, __opts__=None, description=None, dns_name=None, labels=None, name=None, project=None):
        """Create a ManagedZone resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        if not dns_name:
            raise TypeError('Missing required property dns_name')
        __props__['dns_name'] = dns_name

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['project'] = project

        __props__['name_servers'] = None

        super(ManagedZone, __self__).__init__(
            'gcp:dns/managedZone:ManagedZone',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

