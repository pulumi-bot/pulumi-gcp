# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class NetworkEndpointGroup(pulumi.CustomResource):
    default_port: pulumi.Output[float]
    """
    The default port used if the port number is not specified in the
    network endpoint.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource. Provide this property when
    you create the resource.
    """
    name: pulumi.Output[str]
    """
    Name of the resource; provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035. Specifically, the name must be 1-63 characters long and match
    the regular expression `a-z?` which means the
    first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last
    character, which cannot be a dash.
    """
    network: pulumi.Output[str]
    """
    The network to which all network endpoints in the NEG belong.
    Uses "default" project network if unspecified.
    """
    network_endpoint_type: pulumi.Output[str]
    """
    Type of network endpoints in this network endpoint group.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    size: pulumi.Output[float]
    """
    Number of network endpoints in the network endpoint group.
    """
    subnetwork: pulumi.Output[str]
    """
    Optional subnetwork to which all network endpoints in the NEG belong.
    """
    zone: pulumi.Output[str]
    """
    Zone where the network endpoint group is located.
    """
    def __init__(__self__, resource_name, opts=None, default_port=None, description=None, name=None, network=None, network_endpoint_type=None, project=None, subnetwork=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        Network endpoint groups (NEGs) are zonal resources that represent
        collections of IP address and port combinations for GCP resources within a
        single subnet. Each IP address and port combination is called a network
        endpoint.

        Network endpoint groups can be used as backends in backend services for
        HTTP(S), TCP proxy, and SSL proxy load balancers. You cannot use NEGs as a
        backend with internal load balancers. Because NEG backends allow you to
        specify IP addresses and ports, you can distribute traffic in a granular
        fashion among applications or containers running within VM instances.

        To get more information about NetworkEndpointGroup, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)

        ## Example Usage
        ### Network Endpoint Group

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.Network("defaultNetwork", auto_create_subnetworks=False)
        default_subnetwork = gcp.compute.Subnetwork("defaultSubnetwork",
            ip_cidr_range="10.0.0.0/16",
            region="us-central1",
            network=default_network.id)
        neg = gcp.compute.NetworkEndpointGroup("neg",
            network=default_network.id,
            subnetwork=default_subnetwork.id,
            default_port="90",
            zone="us-central1-a")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] default_port: The default port used if the port number is not specified in the
               network endpoint.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: The network to which all network endpoints in the NEG belong.
               Uses "default" project network if unspecified.
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] subnetwork: Optional subnetwork to which all network endpoints in the NEG belong.
        :param pulumi.Input[str] zone: Zone where the network endpoint group is located.
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

            __props__['default_port'] = default_port
            __props__['description'] = description
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['network_endpoint_type'] = network_endpoint_type
            __props__['project'] = project
            __props__['subnetwork'] = subnetwork
            __props__['zone'] = zone
            __props__['self_link'] = None
            __props__['size'] = None
        super(NetworkEndpointGroup, __self__).__init__(
            'gcp:compute/networkEndpointGroup:NetworkEndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, default_port=None, description=None, name=None, network=None, network_endpoint_type=None, project=None, self_link=None, size=None, subnetwork=None, zone=None):
        """
        Get an existing NetworkEndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] default_port: The default port used if the port number is not specified in the
               network endpoint.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: The network to which all network endpoints in the NEG belong.
               Uses "default" project network if unspecified.
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[float] size: Number of network endpoints in the network endpoint group.
        :param pulumi.Input[str] subnetwork: Optional subnetwork to which all network endpoints in the NEG belong.
        :param pulumi.Input[str] zone: Zone where the network endpoint group is located.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["default_port"] = default_port
        __props__["description"] = description
        __props__["name"] = name
        __props__["network"] = network
        __props__["network_endpoint_type"] = network_endpoint_type
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["size"] = size
        __props__["subnetwork"] = subnetwork
        __props__["zone"] = zone
        return NetworkEndpointGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
