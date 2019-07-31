# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NetworkPeering(pulumi.CustomResource):
    auto_create_routes: pulumi.Output[bool]
    """
    If set to `true`, the routes between the two networks will
    be created and managed automatically. Defaults to `true`.
    """
    export_custom_routes: pulumi.Output[bool]
    import_custom_routes: pulumi.Output[bool]
    name: pulumi.Output[str]
    """
    Name of the peering.
    """
    network: pulumi.Output[str]
    """
    Resource link of the network to add a peering to.
    """
    peer_network: pulumi.Output[str]
    """
    Resource link of the peer network.
    """
    state: pulumi.Output[str]
    """
    State for the peering.
    """
    state_details: pulumi.Output[str]
    """
    Details about the current state of the peering.
    """
    def __init__(__self__, resource_name, opts=None, auto_create_routes=None, export_custom_routes=None, import_custom_routes=None, name=None, network=None, peer_network=None, __name__=None, __opts__=None):
        """
        Manages a network peering within GCE. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/vpc/vpc-peering)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/networks).
        
        > **Note:** Both network must create a peering with each other for the peering to be functional.
        
        > **Note:** Subnets IP ranges across peered VPC networks cannot overlap.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_routes: If set to `true`, the routes between the two networks will
               be created and managed automatically. Defaults to `true`.
        :param pulumi.Input[str] name: Name of the peering.
        :param pulumi.Input[str] network: Resource link of the network to add a peering to.
        :param pulumi.Input[str] peer_network: Resource link of the peer network.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_network_peering.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['auto_create_routes'] = auto_create_routes

        __props__['export_custom_routes'] = export_custom_routes

        __props__['import_custom_routes'] = import_custom_routes

        __props__['name'] = name

        if network is None:
            raise TypeError("Missing required property 'network'")
        __props__['network'] = network

        if peer_network is None:
            raise TypeError("Missing required property 'peer_network'")
        __props__['peer_network'] = peer_network

        __props__['state'] = None
        __props__['state_details'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(NetworkPeering, __self__).__init__(
            'gcp:compute/networkPeering:NetworkPeering',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

