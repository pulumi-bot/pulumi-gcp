# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class HaVpnGateway(pulumi.CustomResource):
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    network: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    vpn_interfaces: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, description=None, name=None, network=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a VPN gateway running in GCP. This virtual device is managed
        by Google, but used only by you. This type of VPN Gateway allows for the creation
        of VPN solutions with higher availability than classic Target VPN Gateways.
        
        To get more information about HaVpnGateway, see:
        
        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/vpnGateways)
        * How-to Guides
            * [Choosing a VPN](https://cloud.google.com/vpn/docs/how-to/choosing-a-vpn)
            * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_ha_vpn_gateway.html.markdown.
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

            __props__['description'] = description
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['project'] = project
            __props__['region'] = region
            __props__['self_link'] = None
            __props__['vpn_interfaces'] = None
        super(HaVpnGateway, __self__).__init__(
            'gcp:compute/haVpnGateway:HaVpnGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, name=None, network=None, project=None, region=None, self_link=None, vpn_interfaces=None):
        """
        Get an existing HaVpnGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_ha_vpn_gateway.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["description"] = description
        __props__["name"] = name
        __props__["network"] = network
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["vpn_interfaces"] = vpn_interfaces
        return HaVpnGateway(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

