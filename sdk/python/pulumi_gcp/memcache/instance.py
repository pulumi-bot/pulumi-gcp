# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Instance']


class Instance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_network: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 memcache_parameters: Optional[pulumi.Input[pulumi.InputType['InstanceMemcacheParametersArgs']]] = None,
                 memcache_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_config: Optional[pulumi.Input[pulumi.InputType['InstanceNodeConfigArgs']]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A Google Cloud Memcache instance.

        To get more information about Instance, see:

        * [API documentation](https://cloud.google.com/memorystore/docs/memcached/reference/rest)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/memcache/docs/creating-instances)

        ## Example Usage

        ## Import

        Instance can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:memcache/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
        ```

        ```sh
         $ pulumi import gcp:memcache/instance:Instance default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:memcache/instance:Instance default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:memcache/instance:Instance default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorized_network: The full name of the GCE network to connect the instance to.  If not provided,
               'default' will be used.
        :param pulumi.Input[str] display_name: A user-visible name for the instance.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[pulumi.InputType['InstanceMemcacheParametersArgs']] memcache_parameters: User-specified parameters for this memcache instance.
               Structure is documented below.
        :param pulumi.Input[str] memcache_version: The major version of Memcached software. If not provided, latest supported version will be used.
               Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
               determined by our system based on the latest supported minor version.
               Default value is `MEMCACHE_1_5`.
               Possible values are `MEMCACHE_1_5`.
        :param pulumi.Input[str] name: The resource name of the instance.
        :param pulumi.Input[pulumi.InputType['InstanceNodeConfigArgs']] node_config: Configuration for memcache nodes.
               Structure is documented below.
        :param pulumi.Input[int] node_count: Number of nodes in the memcache instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the Memcache instance. If it is not provided, the provider region is used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: Zones where memcache nodes should be provisioned.  If not
               provided, all zones will be used.
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

            __props__['authorized_network'] = authorized_network
            __props__['display_name'] = display_name
            __props__['labels'] = labels
            __props__['memcache_parameters'] = memcache_parameters
            __props__['memcache_version'] = memcache_version
            __props__['name'] = name
            if node_config is None and not opts.urn:
                raise TypeError("Missing required property 'node_config'")
            __props__['node_config'] = node_config
            if node_count is None and not opts.urn:
                raise TypeError("Missing required property 'node_count'")
            __props__['node_count'] = node_count
            __props__['project'] = project
            __props__['region'] = region
            __props__['zones'] = zones
            __props__['create_time'] = None
            __props__['discovery_endpoint'] = None
            __props__['memcache_full_version'] = None
            __props__['memcache_nodes'] = None
        super(Instance, __self__).__init__(
            'gcp:memcache/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorized_network: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            discovery_endpoint: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            memcache_full_version: Optional[pulumi.Input[str]] = None,
            memcache_nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceMemcacheNodeArgs']]]]] = None,
            memcache_parameters: Optional[pulumi.Input[pulumi.InputType['InstanceMemcacheParametersArgs']]] = None,
            memcache_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_config: Optional[pulumi.Input[pulumi.InputType['InstanceNodeConfigArgs']]] = None,
            node_count: Optional[pulumi.Input[int]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorized_network: The full name of the GCE network to connect the instance to.  If not provided,
               'default' will be used.
        :param pulumi.Input[str] create_time: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] discovery_endpoint: Endpoint for Discovery API
        :param pulumi.Input[str] display_name: A user-visible name for the instance.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] memcache_full_version: The full version of memcached server running on this instance.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceMemcacheNodeArgs']]]] memcache_nodes: Additional information about the instance state, if available.
        :param pulumi.Input[pulumi.InputType['InstanceMemcacheParametersArgs']] memcache_parameters: User-specified parameters for this memcache instance.
               Structure is documented below.
        :param pulumi.Input[str] memcache_version: The major version of Memcached software. If not provided, latest supported version will be used.
               Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
               determined by our system based on the latest supported minor version.
               Default value is `MEMCACHE_1_5`.
               Possible values are `MEMCACHE_1_5`.
        :param pulumi.Input[str] name: The resource name of the instance.
        :param pulumi.Input[pulumi.InputType['InstanceNodeConfigArgs']] node_config: Configuration for memcache nodes.
               Structure is documented below.
        :param pulumi.Input[int] node_count: Number of nodes in the memcache instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the Memcache instance. If it is not provided, the provider region is used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: Zones where memcache nodes should be provisioned.  If not
               provided, all zones will be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorized_network"] = authorized_network
        __props__["create_time"] = create_time
        __props__["discovery_endpoint"] = discovery_endpoint
        __props__["display_name"] = display_name
        __props__["labels"] = labels
        __props__["memcache_full_version"] = memcache_full_version
        __props__["memcache_nodes"] = memcache_nodes
        __props__["memcache_parameters"] = memcache_parameters
        __props__["memcache_version"] = memcache_version
        __props__["name"] = name
        __props__["node_config"] = node_config
        __props__["node_count"] = node_count
        __props__["project"] = project
        __props__["region"] = region
        __props__["zones"] = zones
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizedNetwork")
    def authorized_network(self) -> pulumi.Output[Optional[str]]:
        """
        The full name of the GCE network to connect the instance to.  If not provided,
        'default' will be used.
        """
        return pulumi.get(self, "authorized_network")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="discoveryEndpoint")
    def discovery_endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint for Discovery API
        """
        return pulumi.get(self, "discovery_endpoint")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A user-visible name for the instance.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource labels to represent user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="memcacheFullVersion")
    def memcache_full_version(self) -> pulumi.Output[str]:
        """
        The full version of memcached server running on this instance.
        """
        return pulumi.get(self, "memcache_full_version")

    @property
    @pulumi.getter(name="memcacheNodes")
    def memcache_nodes(self) -> pulumi.Output[Sequence['outputs.InstanceMemcacheNode']]:
        """
        Additional information about the instance state, if available.
        """
        return pulumi.get(self, "memcache_nodes")

    @property
    @pulumi.getter(name="memcacheParameters")
    def memcache_parameters(self) -> pulumi.Output[Optional['outputs.InstanceMemcacheParameters']]:
        """
        User-specified parameters for this memcache instance.
        Structure is documented below.
        """
        return pulumi.get(self, "memcache_parameters")

    @property
    @pulumi.getter(name="memcacheVersion")
    def memcache_version(self) -> pulumi.Output[Optional[str]]:
        """
        The major version of Memcached software. If not provided, latest supported version will be used.
        Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
        determined by our system based on the latest supported minor version.
        Default value is `MEMCACHE_1_5`.
        Possible values are `MEMCACHE_1_5`.
        """
        return pulumi.get(self, "memcache_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> pulumi.Output['outputs.InstanceNodeConfig']:
        """
        Configuration for memcache nodes.
        Structure is documented below.
        """
        return pulumi.get(self, "node_config")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[int]:
        """
        Number of nodes in the memcache instance.
        """
        return pulumi.get(self, "node_count")

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
    def region(self) -> pulumi.Output[str]:
        """
        The region of the Memcache instance. If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Sequence[str]]:
        """
        Zones where memcache nodes should be provisioned.  If not
        provided, all zones will be used.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

