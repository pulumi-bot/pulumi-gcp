# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GCPolicy(pulumi.CustomResource):
    column_family: pulumi.Output[str]
    """
    The name of the column family.
    """
    instance_name: pulumi.Output[str]
    """
    The name of the Bigtable instance.
    """
    max_ages: pulumi.Output[list]
    """
    GC policy that applies to all cells older than the given age.

      * `days` (`float`) - Number of days before applying GC policy.
    """
    max_versions: pulumi.Output[list]
    """
    GC policy that applies to all versions of a cell except for the most recent.

      * `number` (`float`) - Number of version before applying the GC policy.
    """
    mode: pulumi.Output[str]
    """
    If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
    """
    table: pulumi.Output[str]
    """
    The name of the table.
    """
    def __init__(__self__, resource_name, opts=None, column_family=None, instance_name=None, max_ages=None, max_versions=None, mode=None, project=None, table=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a Google Cloud Bigtable GC Policy inside a family. For more information see
        [the official documentation](https://cloud.google.com/bigtable/) and
        [API](https://cloud.google.com/bigtable/docs/go/reference).


        ## Example Usage



        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.bigtable.Instance("instance", cluster=[{
            "cluster_id": "tf-instance-cluster",
            "zone": "us-central1-b",
            "num_nodes": 3,
            "storageType": "HDD",
        }])
        table = gcp.bigtable.Table("table",
            instance_name=instance.name,
            column_family=[{
                "family": "name",
            }])
        policy = gcp.bigtable.GCPolicy("policy",
            instance_name=instance.name,
            table=table.name,
            column_family="name",
            max_age=[{
                "days": 7,
            }])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] column_family: The name of the column family.
        :param pulumi.Input[str] instance_name: The name of the Bigtable instance.
        :param pulumi.Input[list] max_ages: GC policy that applies to all cells older than the given age.
        :param pulumi.Input[list] max_versions: GC policy that applies to all versions of a cell except for the most recent.
        :param pulumi.Input[str] mode: If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] table: The name of the table.

        The **max_ages** object supports the following:

          * `days` (`pulumi.Input[float]`) - Number of days before applying GC policy.

        The **max_versions** object supports the following:

          * `number` (`pulumi.Input[float]`) - Number of version before applying the GC policy.
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

            if column_family is None:
                raise TypeError("Missing required property 'column_family'")
            __props__['column_family'] = column_family
            if instance_name is None:
                raise TypeError("Missing required property 'instance_name'")
            __props__['instance_name'] = instance_name
            __props__['max_ages'] = max_ages
            __props__['max_versions'] = max_versions
            __props__['mode'] = mode
            __props__['project'] = project
            if table is None:
                raise TypeError("Missing required property 'table'")
            __props__['table'] = table
        super(GCPolicy, __self__).__init__(
            'gcp:bigtable/gCPolicy:GCPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, column_family=None, instance_name=None, max_ages=None, max_versions=None, mode=None, project=None, table=None):
        """
        Get an existing GCPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] column_family: The name of the column family.
        :param pulumi.Input[str] instance_name: The name of the Bigtable instance.
        :param pulumi.Input[list] max_ages: GC policy that applies to all cells older than the given age.
        :param pulumi.Input[list] max_versions: GC policy that applies to all versions of a cell except for the most recent.
        :param pulumi.Input[str] mode: If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] table: The name of the table.

        The **max_ages** object supports the following:

          * `days` (`pulumi.Input[float]`) - Number of days before applying GC policy.

        The **max_versions** object supports the following:

          * `number` (`pulumi.Input[float]`) - Number of version before applying the GC policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["column_family"] = column_family
        __props__["instance_name"] = instance_name
        __props__["max_ages"] = max_ages
        __props__["max_versions"] = max_versions
        __props__["mode"] = mode
        __props__["project"] = project
        __props__["table"] = table
        return GCPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

