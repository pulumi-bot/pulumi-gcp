# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['DatabaseArgs', 'Database']

@pulumi.input_type
class DatabaseArgs:
    def __init__(__self__, *,
                 instance: pulumi.Input[str],
                 ddls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Database resource.
        :param pulumi.Input[str] instance: The instance to create the database on.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ddls: An optional list of DDL statements to run inside the newly created
               database. Statements can create tables, indexes, etc. These statements
               execute atomically with the creation of the database: if there is an
               error in any statement, the database is not created.
        :param pulumi.Input[bool] deletion_protection: Whether or not to allow the provider to destroy the instance. Unless this field is set to false
               in state, a `destroy` or `update` that would delete the instance will fail.
        :param pulumi.Input[str] name: A unique identifier for the database, which cannot be changed after
               the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "instance", instance)
        if ddls is not None:
            pulumi.set(__self__, "ddls", ddls)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Input[str]:
        """
        The instance to create the database on.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter
    def ddls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An optional list of DDL statements to run inside the newly created
        database. Statements can create tables, indexes, etc. These statements
        execute atomically with the creation of the database: if there is an
        error in any statement, the database is not created.
        """
        return pulumi.get(self, "ddls")

    @ddls.setter
    def ddls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ddls", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to allow the provider to destroy the instance. Unless this field is set to false
        in state, a `destroy` or `update` that would delete the instance will fail.
        """
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for the database, which cannot be changed after
        the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
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


class Database(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ddls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A Cloud Spanner Database which is hosted on a Spanner instance.

        To get more information about Database, see:

        * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.databases)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/spanner/)

        > **Warning:** It is strongly recommended to set `lifecycle { prevent_destroy = true }` on databases in order to prevent accidental data loss.

        ## Example Usage
        ### Spanner Database Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        main = gcp.spanner.Instance("main",
            config="regional-europe-west1",
            display_name="main-instance")
        database = gcp.spanner.Database("database",
            instance=main.name,
            ddls=[
                "CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
                "CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
            ],
            deletion_protection=False)
        ```

        ## Import

        Database can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:spanner/database:Database default projects/{{project}}/instances/{{instance}}/databases/{{name}}
        ```

        ```sh
         $ pulumi import gcp:spanner/database:Database default instances/{{instance}}/databases/{{name}}
        ```

        ```sh
         $ pulumi import gcp:spanner/database:Database default {{project}}/{{instance}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:spanner/database:Database default {{instance}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ddls: An optional list of DDL statements to run inside the newly created
               database. Statements can create tables, indexes, etc. These statements
               execute atomically with the creation of the database: if there is an
               error in any statement, the database is not created.
        :param pulumi.Input[bool] deletion_protection: Whether or not to allow the provider to destroy the instance. Unless this field is set to false
               in state, a `destroy` or `update` that would delete the instance will fail.
        :param pulumi.Input[str] instance: The instance to create the database on.
        :param pulumi.Input[str] name: A unique identifier for the database, which cannot be changed after
               the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Cloud Spanner Database which is hosted on a Spanner instance.

        To get more information about Database, see:

        * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.databases)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/spanner/)

        > **Warning:** It is strongly recommended to set `lifecycle { prevent_destroy = true }` on databases in order to prevent accidental data loss.

        ## Example Usage
        ### Spanner Database Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        main = gcp.spanner.Instance("main",
            config="regional-europe-west1",
            display_name="main-instance")
        database = gcp.spanner.Database("database",
            instance=main.name,
            ddls=[
                "CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
                "CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
            ],
            deletion_protection=False)
        ```

        ## Import

        Database can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:spanner/database:Database default projects/{{project}}/instances/{{instance}}/databases/{{name}}
        ```

        ```sh
         $ pulumi import gcp:spanner/database:Database default instances/{{instance}}/databases/{{name}}
        ```

        ```sh
         $ pulumi import gcp:spanner/database:Database default {{project}}/{{instance}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:spanner/database:Database default {{instance}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param DatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ddls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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

            __props__['ddls'] = ddls
            __props__['deletion_protection'] = deletion_protection
            if instance is None and not opts.urn:
                raise TypeError("Missing required property 'instance'")
            __props__['instance'] = instance
            __props__['name'] = name
            __props__['project'] = project
            __props__['state'] = None
        super(Database, __self__).__init__(
            'gcp:spanner/database:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ddls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            deletion_protection: Optional[pulumi.Input[bool]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'Database':
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ddls: An optional list of DDL statements to run inside the newly created
               database. Statements can create tables, indexes, etc. These statements
               execute atomically with the creation of the database: if there is an
               error in any statement, the database is not created.
        :param pulumi.Input[bool] deletion_protection: Whether or not to allow the provider to destroy the instance. Unless this field is set to false
               in state, a `destroy` or `update` that would delete the instance will fail.
        :param pulumi.Input[str] instance: The instance to create the database on.
        :param pulumi.Input[str] name: A unique identifier for the database, which cannot be changed after
               the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: An explanation of the status of the database.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ddls"] = ddls
        __props__["deletion_protection"] = deletion_protection
        __props__["instance"] = instance
        __props__["name"] = name
        __props__["project"] = project
        __props__["state"] = state
        return Database(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ddls(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An optional list of DDL statements to run inside the newly created
        database. Statements can create tables, indexes, etc. These statements
        execute atomically with the creation of the database: if there is an
        error in any statement, the database is not created.
        """
        return pulumi.get(self, "ddls")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to allow the provider to destroy the instance. Unless this field is set to false
        in state, a `destroy` or `update` that would delete the instance will fail.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        The instance to create the database on.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique identifier for the database, which cannot be changed after
        the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
        """
        return pulumi.get(self, "name")

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
    def state(self) -> pulumi.Output[str]:
        """
        An explanation of the status of the database.
        """
        return pulumi.get(self, "state")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

