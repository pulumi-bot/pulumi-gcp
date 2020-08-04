# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class InstanceGroupNamedPort(pulumi.CustomResource):
    group: pulumi.Output[str]
    """
    The name of the instance group.
    """
    name: pulumi.Output[str]
    """
    The name for this named port. The name must be 1-63 characters
    long, and comply with RFC1035.
    """
    port: pulumi.Output[float]
    """
    The port number, which can be a value between 1 and 65535.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    zone: pulumi.Output[str]
    """
    The zone of the instance group.
    """
    def __init__(__self__, resource_name, opts=None, group=None, name=None, port=None, project=None, zone=None, __props__=None, __name__=None, __opts__=None):
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

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The name of the instance group.
        :param pulumi.Input[str] name: The name for this named port. The name must be 1-63 characters
               long, and comply with RFC1035.
        :param pulumi.Input[float] port: The port number, which can be a value between 1 and 65535.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] zone: The zone of the instance group.
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

            if group is None:
                raise TypeError("Missing required property 'group'")
            __props__['group'] = group
            __props__['name'] = name
            if port is None:
                raise TypeError("Missing required property 'port'")
            __props__['port'] = port
            __props__['project'] = project
            __props__['zone'] = zone
        super(InstanceGroupNamedPort, __self__).__init__(
            'gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, group=None, name=None, port=None, project=None, zone=None):
        """
        Get an existing InstanceGroupNamedPort resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The name of the instance group.
        :param pulumi.Input[str] name: The name for this named port. The name must be 1-63 characters
               long, and comply with RFC1035.
        :param pulumi.Input[float] port: The port number, which can be a value between 1 and 65535.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] zone: The zone of the instance group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["group"] = group
        __props__["name"] = name
        __props__["port"] = port
        __props__["project"] = project
        __props__["zone"] = zone
        return InstanceGroupNamedPort(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
