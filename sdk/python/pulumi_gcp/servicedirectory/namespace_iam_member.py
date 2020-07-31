# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['NamespaceIamMember']


class NamespaceIamMember(pulumi.CustomResource):
    condition: pulumi.Output[Optional['outputs.NamespaceIamMemberCondition']] = pulumi.output_property("condition")
    etag: pulumi.Output[str] = pulumi.output_property("etag")
    """
    (Computed) The etag of the IAM policy.
    """
    member: pulumi.Output[str] = pulumi.output_property("member")
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Used to find the parent resource to bind the IAM policy to
    """
    role: pulumi.Output[str] = pulumi.output_property("role")
    """
    The role that should be applied. Only one
    `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, condition: Optional[pulumi.Input[pulumi.InputType['NamespaceIamMemberConditionArgs']]] = None, member: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, role: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Three different resources help you manage your IAM policy for Service Directory Namespace. Each of these resources serves a different use case:

        * `servicedirectory.NamespaceIamPolicy`: Authoritative. Sets the IAM policy for the namespace and replaces any existing policy already attached.
        * `servicedirectory.NamespaceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the namespace are preserved.
        * `servicedirectory.NamespaceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the namespace are preserved.

        > **Note:** `servicedirectory.NamespaceIamPolicy` **cannot** be used in conjunction with `servicedirectory.NamespaceIamBinding` and `servicedirectory.NamespaceIamMember` or they will fight over what your policy should be.

        > **Note:** `servicedirectory.NamespaceIamBinding` resources **can be** used in conjunction with `servicedirectory.NamespaceIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_service\_directory\_namespace\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.servicedirectory.NamespaceIamPolicy("policy", policy_data=admin.policy_data)
        ```

        ## google\_service\_directory\_namespace\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.servicedirectory.NamespaceIamBinding("binding",
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_service\_directory\_namespace\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.servicedirectory.NamespaceIamMember("member",
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
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
            if member is None:
                raise TypeError("Missing required property 'member'")
            __props__['member'] = member
            __props__['name'] = name
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(NamespaceIamMember, __self__).__init__(
            'gcp:servicedirectory/namespaceIamMember:NamespaceIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, condition: Optional[pulumi.Input[pulumi.InputType['NamespaceIamMemberConditionArgs']]] = None, etag: Optional[pulumi.Input[str]] = None, member: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, role: Optional[pulumi.Input[str]] = None) -> 'NamespaceIamMember':
        """
        Get an existing NamespaceIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["member"] = member
        __props__["name"] = name
        __props__["role"] = role
        return NamespaceIamMember(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

