# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['User']


class User(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new Google SQL User on a Google SQL User Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/users).

        ## Example Usage

        Example creating a SQL User.

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        db_name_suffix = random.RandomId("dbNameSuffix", byte_length=4)
        master = gcp.sql.DatabaseInstance("master", settings=gcp.sql.DatabaseInstanceSettingsArgs(
            tier="db-f1-micro",
        ))
        users = gcp.sql.User("users",
            instance=master.name,
            host="me.com",
            password="changeme")
        ```

        ## Import

        SQL users for MySQL databases can be imported using the `project`, `instance`, `host` and `name`, e.g.

        ```sh
         $ pulumi import gcp:sql/user:User users my-project/master-instance/my-domain.com/me
        ```

         SQL users for PostgreSQL databases can be imported using the `project`, `instance` and `name`, e.g.

        ```sh
         $ pulumi import gcp:sql/user:User users my-project/master-instance/me
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: The host the user can connect from. This is only supported
               for MySQL instances. Don't set this field for PostgreSQL instances.
               Can be an IP address. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the user. Changing this forces a new resource
               to be created.
        :param pulumi.Input[str] password: The password for the user. Can be updated. For Postgres
               instances this is a Required field.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
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

            __props__['host'] = host
            if instance is None and not opts.urn:
                raise TypeError("Missing required property 'instance'")
            __props__['instance'] = instance
            __props__['name'] = name
            __props__['password'] = password
            __props__['project'] = project
        super(User, __self__).__init__(
            'gcp:sql/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            host: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: The host the user can connect from. This is only supported
               for MySQL instances. Don't set this field for PostgreSQL instances.
               Can be an IP address. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the user. Changing this forces a new resource
               to be created.
        :param pulumi.Input[str] password: The password for the user. Can be updated. For Postgres
               instances this is a Required field.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["host"] = host
        __props__["instance"] = instance
        __props__["name"] = name
        __props__["password"] = password
        __props__["project"] = project
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        The host the user can connect from. This is only supported
        for MySQL instances. Don't set this field for PostgreSQL instances.
        Can be an IP address. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        The name of the Cloud SQL instance. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the user. Changing this forces a new resource
        to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password for the user. Can be updated. For Postgres
        instances this is a Required field.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

