# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ManagedZone(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A textual description field. Defaults to 'Managed by Terraform'.
    """
    dns_name: pulumi.Output[str]
    """
    The fully qualified DNS name of this zone, e.g. `terraform.io.`.
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to the instance.
    """
    name: pulumi.Output[str]
    """
    A unique name for the resource, required by GCE.
    Changing this forces a new resource to be created.
    """
    name_servers: pulumi.Output[list]
    """
    The list of nameservers that will be authoritative for this
    domain. Use NS records to redirect from your DNS provider to these names,
    thus making Google Cloud DNS authoritative for this zone.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, dns_name=None, labels=None, name=None, project=None):
        """
        Manages a zone within Google Cloud DNS. For more information see [the official documentation](https://cloud.google.com/dns/zones/) and
        [API](https://cloud.google.com/dns/api/v1/managedZones).
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] description: A textual description field. Defaults to 'Managed by Terraform'.
        :param pulumi.Input[str] dns_name: The fully qualified DNS name of this zone, e.g. `terraform.io.`.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the instance.
        :param pulumi.Input[str] name: A unique name for the resource, required by GCE.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
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

