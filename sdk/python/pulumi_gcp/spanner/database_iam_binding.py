# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DatabaseIAMBinding']


class DatabaseIAMBinding(pulumi.CustomResource):
    condition: pulumi.Output[Optional['outputs.DatabaseIAMBindingCondition']] = pulumi.property("condition")

    database: pulumi.Output[str] = pulumi.property("database")
    """
    The name of the Spanner database.
    """

    etag: pulumi.Output[str] = pulumi.property("etag")
    """
    (Computed) The etag of the database's IAM policy.
    """

    instance: pulumi.Output[str] = pulumi.property("instance")
    """
    The name of the Spanner instance the database belongs to.
    """

    members: pulumi.Output[List[str]] = pulumi.property("members")

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """

    role: pulumi.Output[str] = pulumi.property("role")
    """
    The role that should be applied. Only one
    `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['DatabaseIAMBindingConditionArgs']]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:

        * `spanner.DatabaseIAMPolicy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.

        > **Warning:** It's entirely possibly to lock yourself out of your database using `spanner.DatabaseIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.

        * `spanner.DatabaseIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
        * `spanner.DatabaseIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.

        > **Note:** `spanner.DatabaseIAMPolicy` **cannot** be used in conjunction with `spanner.DatabaseIAMBinding` and `spanner.DatabaseIAMMember` or they will fight over what your policy should be.

        > **Note:** `spanner.DatabaseIAMBinding` resources **can be** used in conjunction with `spanner.DatabaseIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\_spanner\_database\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/editor",
            "members": ["user:jane@example.com"],
        }])
        database = gcp.spanner.DatabaseIAMPolicy("database",
            instance="your-instance-name",
            database="your-database-name",
            policy_data=admin.policy_data)
        ```

        ## google\_spanner\_database\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        database = gcp.spanner.DatabaseIAMBinding("database",
            database="your-database-name",
            instance="your-instance-name",
            members=["user:jane@example.com"],
            role="roles/compute.networkUser")
        ```

        ## google\_spanner\_database\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        database = gcp.spanner.DatabaseIAMMember("database",
            database="your-database-name",
            instance="your-instance-name",
            member="user:jane@example.com",
            role="roles/compute.networkUser")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The name of the Spanner database.
        :param pulumi.Input[str] instance: The name of the Spanner instance the database belongs to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
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

            __props__['condition'] = condition
            if database is None:
                raise TypeError("Missing required property 'database'")
            __props__['database'] = database
            if instance is None:
                raise TypeError("Missing required property 'instance'")
            __props__['instance'] = instance
            if members is None:
                raise TypeError("Missing required property 'members'")
            __props__['members'] = members
            __props__['project'] = project
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(DatabaseIAMBinding, __self__).__init__(
            'gcp:spanner/databaseIAMBinding:DatabaseIAMBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['DatabaseIAMBindingConditionArgs']]] = None,
            database: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'DatabaseIAMBinding':
        """
        Get an existing DatabaseIAMBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The name of the Spanner database.
        :param pulumi.Input[str] etag: (Computed) The etag of the database's IAM policy.
        :param pulumi.Input[str] instance: The name of the Spanner instance the database belongs to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["database"] = database
        __props__["etag"] = etag
        __props__["instance"] = instance
        __props__["members"] = members
        __props__["project"] = project
        __props__["role"] = role
        return DatabaseIAMBinding(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

