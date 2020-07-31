# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['NetworkPeeringRoutesConfig']


class NetworkPeeringRoutesConfig(pulumi.CustomResource):
    export_custom_routes: pulumi.Output[bool] = pulumi.output_property("exportCustomRoutes")
    """
    Whether to export the custom routes to the peer network.
    """
    import_custom_routes: pulumi.Output[bool] = pulumi.output_property("importCustomRoutes")
    """
    Whether to import the custom routes to the peer network.
    """
    network: pulumi.Output[str] = pulumi.output_property("network")
    """
    The name of the primary network for the peering.
    """
    peering: pulumi.Output[str] = pulumi.output_property("peering")
    """
    Name of the peering.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, export_custom_routes: Optional[pulumi.Input[bool]] = None, import_custom_routes: Optional[pulumi.Input[bool]] = None, network: Optional[pulumi.Input[str]] = None, peering: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Manage a network peering's route settings without managing the peering as
        a whole. This resource is primarily intended for use with GCP-generated
        peerings that shouldn't otherwise be managed by other tools. Deleting this
        resource is a no-op and the peering will not be modified.

        To get more information about NetworkPeeringRoutesConfig, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/vpc/docs/vpc-peering)

        ## Example Usage
        ### Network Peering Routes Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        network_primary = gcp.compute.Network("networkPrimary", auto_create_subnetworks="false")
        network_secondary = gcp.compute.Network("networkSecondary", auto_create_subnetworks="false")
        peering_primary = gcp.compute.NetworkPeering("peeringPrimary",
            network=network_primary.id,
            peer_network=network_secondary.id,
            import_custom_routes=True,
            export_custom_routes=True)
        peering_primary_routes = gcp.compute.NetworkPeeringRoutesConfig("peeringPrimaryRoutes",
            peering=peering_primary.name,
            network=network_primary.name,
            import_custom_routes=True,
            export_custom_routes=True)
        peering_secondary = gcp.compute.NetworkPeering("peeringSecondary",
            network=network_secondary.id,
            peer_network=network_primary.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] export_custom_routes: Whether to export the custom routes to the peer network.
        :param pulumi.Input[bool] import_custom_routes: Whether to import the custom routes to the peer network.
        :param pulumi.Input[str] network: The name of the primary network for the peering.
        :param pulumi.Input[str] peering: Name of the peering.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if export_custom_routes is None:
                raise TypeError("Missing required property 'export_custom_routes'")
            __props__['export_custom_routes'] = export_custom_routes
            if import_custom_routes is None:
                raise TypeError("Missing required property 'import_custom_routes'")
            __props__['import_custom_routes'] = import_custom_routes
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            if peering is None:
                raise TypeError("Missing required property 'peering'")
            __props__['peering'] = peering
            __props__['project'] = project
        super(NetworkPeeringRoutesConfig, __self__).__init__(
            'gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, export_custom_routes: Optional[pulumi.Input[bool]] = None, import_custom_routes: Optional[pulumi.Input[bool]] = None, network: Optional[pulumi.Input[str]] = None, peering: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None) -> 'NetworkPeeringRoutesConfig':
        """
        Get an existing NetworkPeeringRoutesConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] export_custom_routes: Whether to export the custom routes to the peer network.
        :param pulumi.Input[bool] import_custom_routes: Whether to import the custom routes to the peer network.
        :param pulumi.Input[str] network: The name of the primary network for the peering.
        :param pulumi.Input[str] peering: Name of the peering.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["export_custom_routes"] = export_custom_routes
        __props__["import_custom_routes"] = import_custom_routes
        __props__["network"] = network
        __props__["peering"] = peering
        __props__["project"] = project
        return NetworkPeeringRoutesConfig(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

