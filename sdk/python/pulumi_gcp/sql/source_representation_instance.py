# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SourceRepresentationInstance(pulumi.CustomResource):
    database_version: pulumi.Output[str]
    """
    The MySQL version running on your source database server.
    """
    host: pulumi.Output[str]
    """
    The externally accessible IPv4 address for the source database server.
    """
    name: pulumi.Output[str]
    """
    The name of the source representation instance. Use any valid Cloud SQL instance name.
    """
    port: pulumi.Output[float]
    """
    The externally accessible port for the source database server.
    Defaults to 3306.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The Region in which the created instance should reside.
    If it is not provided, the provider region is used.
    """
    def __init__(__self__, resource_name, opts=None, database_version=None, host=None, name=None, port=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        A source representation instance is a Cloud SQL instance that represents
        the source database server to the Cloud SQL replica. It is visible in the
        Cloud Console and appears the same as a regular Cloud SQL instance, but it
        contains no data, requires no configuration or maintenance, and does not
        affect billing. You cannot update the source representation instance.



        ## Example Usage

        ### Sql Source Representation Instance Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.sql.SourceRepresentationInstance("instance",
            database_version="MYSQL_5_7",
            host="10.20.30.40",
            port=3306,
            region="us-central1")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_version: The MySQL version running on your source database server.
        :param pulumi.Input[str] host: The externally accessible IPv4 address for the source database server.
        :param pulumi.Input[str] name: The name of the source representation instance. Use any valid Cloud SQL instance name.
        :param pulumi.Input[float] port: The externally accessible port for the source database server.
               Defaults to 3306.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Region in which the created instance should reside.
               If it is not provided, the provider region is used.
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

            if database_version is None:
                raise TypeError("Missing required property 'database_version'")
            __props__['database_version'] = database_version
            if host is None:
                raise TypeError("Missing required property 'host'")
            __props__['host'] = host
            __props__['name'] = name
            __props__['port'] = port
            __props__['project'] = project
            __props__['region'] = region
        super(SourceRepresentationInstance, __self__).__init__(
            'gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, database_version=None, host=None, name=None, port=None, project=None, region=None):
        """
        Get an existing SourceRepresentationInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_version: The MySQL version running on your source database server.
        :param pulumi.Input[str] host: The externally accessible IPv4 address for the source database server.
        :param pulumi.Input[str] name: The name of the source representation instance. Use any valid Cloud SQL instance name.
        :param pulumi.Input[float] port: The externally accessible port for the source database server.
               Defaults to 3306.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Region in which the created instance should reside.
               If it is not provided, the provider region is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["database_version"] = database_version
        __props__["host"] = host
        __props__["name"] = name
        __props__["port"] = port
        __props__["project"] = project
        __props__["region"] = region
        return SourceRepresentationInstance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

