# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class InstanceIamBinding(pulumi.CustomResource):
    condition: pulumi.Output[dict]
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the instances's IAM policy.
    """
    instance: pulumi.Output[str]
    """
    The name or relative resource id of the instance to manage IAM policies for.
    """
    members: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The project in which the instance belongs. If it
    is not provided, a default will be supplied.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, condition=None, instance=None, members=None, project=None, role=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:

        * `bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
        * `bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.

        > **Note:** `bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `bigtable.InstanceIamBinding` and `bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `bigtable.InstanceIamPolicy` replaces the entire policy.

        > **Note:** `bigtable.InstanceIamBinding` resources **can be** used in conjunction with `bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_bigtable\_instance\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/editor",
            "members": ["user:jane@example.com"],
        }])
        editor = gcp.bigtable.InstanceIamPolicy("editor",
            project="your-project",
            instance="your-bigtable-instance",
            policy_data=admin.policy_data)
        ```

        ## google\_bigtable\_instance\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.InstanceIamBinding("editor",
            instance="your-bigtable-instance",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\_bigtable\_instance\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.InstanceIamMember("editor",
            instance="your-bigtable-instance",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance to manage IAM policies for.
        :param pulumi.Input[str] project: The project in which the instance belongs. If it
               is not provided, a default will be supplied.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`)
          * `expression` (`pulumi.Input[str]`)
          * `title` (`pulumi.Input[str]`)
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

            __props__['condition'] = condition
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
        super(InstanceIamBinding, __self__).__init__(
            'gcp:bigtable/instanceIamBinding:InstanceIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, condition=None, etag=None, instance=None, members=None, project=None, role=None):
        """
        Get an existing InstanceIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the instances's IAM policy.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance to manage IAM policies for.
        :param pulumi.Input[str] project: The project in which the instance belongs. If it
               is not provided, a default will be supplied.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`)
          * `expression` (`pulumi.Input[str]`)
          * `title` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["instance"] = instance
        __props__["members"] = members
        __props__["project"] = project
        __props__["role"] = role
        return InstanceIamBinding(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
