# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceGroupNamedPortInitArgs', 'InstanceGroupNamedPort']

@pulumi.input_type
class InstanceGroupNamedPortInitArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 port: pulumi.Input[int],
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceGroupNamedPort resource.
        :param pulumi.Input[str] group: The name of the instance group.
        :param pulumi.Input[int] port: The port number, which can be a value between 1 and 65535.
        :param pulumi.Input[str] name: The name for this named port. The name must be 1-63 characters
               long, and comply with RFC1035.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] zone: The zone of the instance group.
        """
        pulumi.set(__self__, "group", group)
        pulumi.set(__self__, "port", port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The name of the instance group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The port number, which can be a value between 1 and 65535.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for this named port. The name must be 1-63 characters
        long, and comply with RFC1035.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone of the instance group.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _InstanceGroupNamedPortState:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceGroupNamedPort resources.
        :param pulumi.Input[str] group: The name of the instance group.
        :param pulumi.Input[str] name: The name for this named port. The name must be 1-63 characters
               long, and comply with RFC1035.
        :param pulumi.Input[int] port: The port number, which can be a value between 1 and 65535.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] zone: The zone of the instance group.
        """
        if group is not None:
            pulumi.set(__self__, "group", group)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the instance group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for this named port. The name must be 1-63 characters
        long, and comply with RFC1035.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port number, which can be a value between 1 and 65535.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

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
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone of the instance group.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class InstanceGroupNamedPort(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Mange the named ports setting for a managed instance group without
        managing the group as whole. This resource is primarily intended for use
        with GKE-generated groups that shouldn't otherwise be managed by other
        tools.

        To get more information about InstanceGroupNamedPort, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroup)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/)

        ## Example Usage
        ### Instance Group Named Port Gke

        ```python
        import pulumi
        import pulumi_gcp as gcp

        container_network = gcp.compute.Network("containerNetwork", auto_create_subnetworks=False)
        container_subnetwork = gcp.compute.Subnetwork("containerSubnetwork",
            region="us-central1",
            network=container_network.name,
            ip_cidr_range="10.0.36.0/24")
        my_cluster = gcp.container.Cluster("myCluster",
            location="us-central1-a",
            initial_node_count=1,
            network=container_network.name,
            subnetwork=container_subnetwork.name,
            ip_allocation_policy=gcp.container.ClusterIpAllocationPolicyArgs(
                cluster_ipv4_cidr_block="/19",
                services_ipv4_cidr_block="/22",
            ))
        my_port = gcp.compute.InstanceGroupNamedPort("myPort",
            group=my_cluster.instance_group_urls[0],
            zone="us-central1-a",
            port=8080)
        my_ports = gcp.compute.InstanceGroupNamedPort("myPorts",
            group=my_cluster.instance_group_urls[0],
            zone="us-central1-a",
            port=4443)
        ```

        ## Import

        InstanceGroupNamedPort can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{project}}/{{zone}}/{{group}}/{{port}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{zone}}/{{group}}/{{port}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{group}}/{{port}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The name of the instance group.
        :param pulumi.Input[str] name: The name for this named port. The name must be 1-63 characters
               long, and comply with RFC1035.
        :param pulumi.Input[int] port: The port number, which can be a value between 1 and 65535.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] zone: The zone of the instance group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceGroupNamedPortInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Mange the named ports setting for a managed instance group without
        managing the group as whole. This resource is primarily intended for use
        with GKE-generated groups that shouldn't otherwise be managed by other
        tools.

        To get more information about InstanceGroupNamedPort, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroup)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/)

        ## Example Usage
        ### Instance Group Named Port Gke

        ```python
        import pulumi
        import pulumi_gcp as gcp

        container_network = gcp.compute.Network("containerNetwork", auto_create_subnetworks=False)
        container_subnetwork = gcp.compute.Subnetwork("containerSubnetwork",
            region="us-central1",
            network=container_network.name,
            ip_cidr_range="10.0.36.0/24")
        my_cluster = gcp.container.Cluster("myCluster",
            location="us-central1-a",
            initial_node_count=1,
            network=container_network.name,
            subnetwork=container_subnetwork.name,
            ip_allocation_policy=gcp.container.ClusterIpAllocationPolicyArgs(
                cluster_ipv4_cidr_block="/19",
                services_ipv4_cidr_block="/22",
            ))
        my_port = gcp.compute.InstanceGroupNamedPort("myPort",
            group=my_cluster.instance_group_urls[0],
            zone="us-central1-a",
            port=8080)
        my_ports = gcp.compute.InstanceGroupNamedPort("myPorts",
            group=my_cluster.instance_group_urls[0],
            zone="us-central1-a",
            port=4443)
        ```

        ## Import

        InstanceGroupNamedPort can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{project}}/{{zone}}/{{group}}/{{port}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{zone}}/{{group}}/{{port}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{group}}/{{port}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param InstanceGroupNamedPortInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceGroupNamedPortInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceGroupNamedPortInitArgs.__new__(InstanceGroupNamedPortInitArgs)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["name"] = name
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            __props__.__dict__["project"] = project
            __props__.__dict__["zone"] = zone
        super(InstanceGroupNamedPort, __self__).__init__(
            'gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            project: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceGroupNamedPort':
        """
        Get an existing InstanceGroupNamedPort resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The name of the instance group.
        :param pulumi.Input[str] name: The name for this named port. The name must be 1-63 characters
               long, and comply with RFC1035.
        :param pulumi.Input[int] port: The port number, which can be a value between 1 and 65535.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] zone: The zone of the instance group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceGroupNamedPortState.__new__(_InstanceGroupNamedPortState)

        __props__.__dict__["group"] = group
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["project"] = project
        __props__.__dict__["zone"] = zone
        return InstanceGroupNamedPort(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The name of the instance group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for this named port. The name must be 1-63 characters
        long, and comply with RFC1035.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The port number, which can be a value between 1 and 65535.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone of the instance group.
        """
        return pulumi.get(self, "zone")

