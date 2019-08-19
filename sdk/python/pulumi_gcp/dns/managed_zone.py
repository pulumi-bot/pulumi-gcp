# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ManagedZone(pulumi.CustomResource):
    description: pulumi.Output[str]
    dns_name: pulumi.Output[str]
    dnssec_config: pulumi.Output[dict]
    forwarding_config: pulumi.Output[dict]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    name_servers: pulumi.Output[list]
    peering_config: pulumi.Output[dict]
    private_visibility_config: pulumi.Output[dict]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    visibility: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, description=None, dns_name=None, dnssec_config=None, forwarding_config=None, labels=None, name=None, peering_config=None, private_visibility_config=None, project=None, visibility=None, __props__=None, __name__=None, __opts__=None):
        """
        A zone is a subtree of the DNS namespace under one administrative
        responsibility. A ManagedZone is a resource that represents a DNS zone
        hosted by the Cloud DNS service.
        
        
        To get more information about ManagedZone, see:
        
        * [API documentation](https://cloud.google.com/dns/api/v1/managedZones)
        * How-to Guides
            * [Managing Zones](https://cloud.google.com/dns/zones/)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_managed_zone.html.markdown.
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
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            if dns_name is None:
                raise TypeError("Missing required property 'dns_name'")
            __props__['dns_name'] = dns_name
            __props__['dnssec_config'] = dnssec_config
            __props__['forwarding_config'] = forwarding_config
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['peering_config'] = peering_config
            __props__['private_visibility_config'] = private_visibility_config
            __props__['project'] = project
            __props__['visibility'] = visibility
            __props__['name_servers'] = None
        super(ManagedZone, __self__).__init__(
            'gcp:dns/managedZone:ManagedZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, dns_name=None, dnssec_config=None, forwarding_config=None, labels=None, name=None, name_servers=None, peering_config=None, private_visibility_config=None, project=None, visibility=None):
        """
        Get an existing ManagedZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_managed_zone.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["description"] = description
        __props__["dns_name"] = dns_name
        __props__["dnssec_config"] = dnssec_config
        __props__["forwarding_config"] = forwarding_config
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["name_servers"] = name_servers
        __props__["peering_config"] = peering_config
        __props__["private_visibility_config"] = private_visibility_config
        __props__["project"] = project
        __props__["visibility"] = visibility
        return ManagedZone(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

