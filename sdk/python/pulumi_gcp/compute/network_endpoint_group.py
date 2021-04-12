# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['NetworkEndpointGroupArgs', 'NetworkEndpointGroup']

@pulumi.input_type
class NetworkEndpointGroupArgs:
    def __init__(__self__, *,
                 network: pulumi.Input[str],
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NetworkEndpointGroup resource.
        :param pulumi.Input[str] network: The network to which all network endpoints in the NEG belong.
               Uses "default" project network if unspecified.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the
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
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group.
               Default value is `GCE_VM_IP_PORT`.
               Possible values are `GCE_VM_IP_PORT`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] subnetwork: Optional subnetwork to which all network endpoints in the NEG belong.
        :param pulumi.Input[str] zone: Zone where the network endpoint group is located.
        """
        pulumi.set(__self__, "network", network)
        if default_port is not None:
            pulumi.set(__self__, "default_port", default_port)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_endpoint_type is not None:
            pulumi.set(__self__, "network_endpoint_type", network_endpoint_type)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if subnetwork is not None:
            pulumi.set(__self__, "subnetwork", subnetwork)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input[str]:
        """
        The network to which all network endpoints in the NEG belong.
        Uses "default" project network if unspecified.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input[str]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="defaultPort")
    def default_port(self) -> Optional[pulumi.Input[int]]:
        """
        The default port used if the port number is not specified in the
        network endpoint.
        """
        return pulumi.get(self, "default_port")

    @default_port.setter
    def default_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_port", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource; provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of network endpoints in this network endpoint group.
        Default value is `GCE_VM_IP_PORT`.
        Possible values are `GCE_VM_IP_PORT`.
        """
        return pulumi.get(self, "network_endpoint_type")

    @network_endpoint_type.setter
    def network_endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_endpoint_type", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def subnetwork(self) -> Optional[pulumi.Input[str]]:
        """
        Optional subnetwork to which all network endpoints in the NEG belong.
        """
        return pulumi.get(self, "subnetwork")

    @subnetwork.setter
    def subnetwork(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnetwork", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Zone where the network endpoint group is located.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class NetworkEndpointGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

        Recreating a network endpoint group that's in use by another resource will give a
        `resourceInUseByAnotherResource` error. Use `lifecycle.create_before_destroy`
        to avoid this type of error.

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
            default_port=90,
            zone="us-central1-a")
        ```

        ## Import

        NetworkEndpointGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the
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
               Default value is `GCE_VM_IP_PORT`.
               Possible values are `GCE_VM_IP_PORT`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] subnetwork: Optional subnetwork to which all network endpoints in the NEG belong.
        :param pulumi.Input[str] zone: Zone where the network endpoint group is located.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkEndpointGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        Recreating a network endpoint group that's in use by another resource will give a
        `resourceInUseByAnotherResource` error. Use `lifecycle.create_before_destroy`
        to avoid this type of error.

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
            default_port=90,
            zone="us-central1-a")
        ```

        ## Import

        NetworkEndpointGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param NetworkEndpointGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkEndpointGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['default_port'] = default_port
            __props__['description'] = description
            __props__['name'] = name
            if network is None and not opts.urn:
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_port: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            network_endpoint_type: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            subnetwork: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'NetworkEndpointGroup':
        """
        Get an existing NetworkEndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the
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
               Default value is `GCE_VM_IP_PORT`.
               Possible values are `GCE_VM_IP_PORT`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[int] size: Number of network endpoints in the network endpoint group.
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

    @property
    @pulumi.getter(name="defaultPort")
    def default_port(self) -> pulumi.Output[Optional[int]]:
        """
        The default port used if the port number is not specified in the
        network endpoint.
        """
        return pulumi.get(self, "default_port")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource; provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        The network to which all network endpoints in the NEG belong.
        Uses "default" project network if unspecified.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of network endpoints in this network endpoint group.
        Default value is `GCE_VM_IP_PORT`.
        Possible values are `GCE_VM_IP_PORT`.
        """
        return pulumi.get(self, "network_endpoint_type")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Number of network endpoints in the network endpoint group.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def subnetwork(self) -> pulumi.Output[Optional[str]]:
        """
        Optional subnetwork to which all network endpoints in the NEG belong.
        """
        return pulumi.get(self, "subnetwork")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Zone where the network endpoint group is located.
        """
        return pulumi.get(self, "zone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

