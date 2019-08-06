# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
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
    prefix_length: pulumi.Output[float]
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
    def __init__(__self__, resource_name, opts=None, address=None, address_type=None, description=None, ip_version=None, labels=None, name=None, network=None, prefix_length=None, project=None, purpose=None, __name__=None, __opts__=None):
        """
        Represents a Global Address resource. Global addresses are used for
        HTTP(S) load balancing.
        
        
        To get more information about GlobalAddress, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/globalAddresses)
        * How-to Guides
            * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_global_address.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__

        __props__ = dict()

        __props__['address'] = address
        __props__['address_type'] = address_type
        __props__['description'] = description
        __props__['ip_version'] = ip_version
        __props__['labels'] = labels
        __props__['name'] = name
        __props__['network'] = network
        __props__['prefix_length'] = prefix_length
        __props__['project'] = project
        __props__['purpose'] = purpose
        __props__['creation_timestamp'] = None
        __props__['label_fingerprint'] = None
        __props__['self_link'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(GlobalAddress, __self__).__init__(
            'gcp:compute/globalAddress:GlobalAddress',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

