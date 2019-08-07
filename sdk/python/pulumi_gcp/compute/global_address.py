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
    def __init__(__self__, resource_name, opts=None, address=None, address_type=None, description=None, ip_version=None, labels=None, name=None, network=None, prefix_length=None, project=None, purpose=None, __props__=None, __name__=None, __opts__=None):
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
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
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
        super(GlobalAddress, __self__).__init__(
            'gcp:compute/globalAddress:GlobalAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address=None, address_type=None, creation_timestamp=None, description=None, ip_version=None, label_fingerprint=None, labels=None, name=None, network=None, prefix_length=None, project=None, purpose=None, self_link=None):
        """
        Get an existing GlobalAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_global_address.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["address"] = address
        __props__["address_type"] = address_type
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["ip_version"] = ip_version
        __props__["label_fingerprint"] = label_fingerprint
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["network"] = network
        __props__["prefix_length"] = prefix_length
        __props__["project"] = project
        __props__["purpose"] = purpose
        __props__["self_link"] = self_link
        return GlobalAddress(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

