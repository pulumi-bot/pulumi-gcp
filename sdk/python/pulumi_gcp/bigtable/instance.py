# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    cluster: pulumi.Output[dict]
    """
    A block of cluster configuration options. Either `cluster` or `cluster_id` must be used. Only one cluster may be specified. See structure below.
    """
    cluster_id: pulumi.Output[str]
    """
    The ID of the Cloud Bigtable cluster.
    """
    display_name: pulumi.Output[str]
    """
    The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
    """
    instance_type: pulumi.Output[str]
    """
    The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
    """
    name: pulumi.Output[str]
    """
    The name of the Cloud Bigtable instance.
    """
    num_nodes: pulumi.Output[int]
    """
    The number of nodes in your Cloud Bigtable cluster. Minimum of `3` for a `PRODUCTION` instance. Cannot be set for a `DEVELOPMENT` instance.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    storage_type: pulumi.Output[str]
    """
    The storage type to use. One of `"SSD"` or `"HDD"`. Defaults to `"SSD"`.
    """
    zone: pulumi.Output[str]
    """
    The zone to create the Cloud Bigtable cluster in. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
    """
    def __init__(__self__, resource_name, opts=None, cluster=None, cluster_id=None, display_name=None, instance_type=None, name=None, num_nodes=None, project=None, storage_type=None, zone=None, __name__=None, __opts__=None):
        """
        Creates a Google Bigtable instance. For more information see
        [the official documentation](https://cloud.google.com/bigtable/) and
        [API](https://cloud.google.com/bigtable/docs/go/reference).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] cluster: A block of cluster configuration options. Either `cluster` or `cluster_id` must be used. Only one cluster may be specified. See structure below.
        :param pulumi.Input[str] cluster_id: The ID of the Cloud Bigtable cluster.
        :param pulumi.Input[str] display_name: The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        :param pulumi.Input[str] instance_type: The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        :param pulumi.Input[str] name: The name of the Cloud Bigtable instance.
        :param pulumi.Input[int] num_nodes: The number of nodes in your Cloud Bigtable cluster. Minimum of `3` for a `PRODUCTION` instance. Cannot be set for a `DEVELOPMENT` instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] storage_type: The storage type to use. One of `"SSD"` or `"HDD"`. Defaults to `"SSD"`.
        :param pulumi.Input[str] zone: The zone to create the Cloud Bigtable cluster in. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
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

        __props__['cluster'] = cluster

        __props__['cluster_id'] = cluster_id

        __props__['display_name'] = display_name

        __props__['instance_type'] = instance_type

        __props__['name'] = name

        __props__['num_nodes'] = num_nodes

        __props__['project'] = project

        __props__['storage_type'] = storage_type

        __props__['zone'] = zone

        super(Instance, __self__).__init__(
            'gcp:bigtable/instance:Instance',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

