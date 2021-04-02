# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PolicyTagIamMemberArgs', 'PolicyTagIamMember']

@pulumi.input_type
class PolicyTagIamMemberArgs:
    def __init__(__self__, *,
                 member: pulumi.Input[str],
                 policy_tag: pulumi.Input[str],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['PolicyTagIamMemberConditionArgs']] = None):
        """
        The set of arguments for constructing a PolicyTagIamMember resource.
        :param pulumi.Input[str] policy_tag: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        pulumi.set(__self__, "member", member)
        pulumi.set(__self__, "policy_tag", policy_tag)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter
    def member(self) -> pulumi.Input[str]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: pulumi.Input[str]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter(name="policyTag")
    def policy_tag(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "policy_tag")

    @policy_tag.setter
    def policy_tag(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_tag", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['PolicyTagIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['PolicyTagIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)


class PolicyTagIamMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['PolicyTagIamMemberConditionArgs']]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 policy_tag: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Three different resources help you manage your IAM policy for Data catalog PolicyTag. Each of these resources serves a different use case:

        * `datacatalog.PolicyTagIamPolicy`: Authoritative. Sets the IAM policy for the policytag and replaces any existing policy already attached.
        * `datacatalog.PolicyTagIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the policytag are preserved.
        * `datacatalog.PolicyTagIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the policytag are preserved.

        > **Note:** `datacatalog.PolicyTagIamPolicy` **cannot** be used in conjunction with `datacatalog.PolicyTagIamBinding` and `datacatalog.PolicyTagIamMember` or they will fight over what your policy should be.

        > **Note:** `datacatalog.PolicyTagIamBinding` resources **can be** used in conjunction with `datacatalog.PolicyTagIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_data\_catalog\_policy\_tag\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.datacatalog.PolicyTagIamPolicy("policy",
            policy_tag=google_data_catalog_policy_tag["basic_policy_tag"]["name"],
            policy_data=admin.policy_data)
        ```

        ## google\_data\_catalog\_policy\_tag\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.datacatalog.PolicyTagIamBinding("binding",
            policy_tag=google_data_catalog_policy_tag["basic_policy_tag"]["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_data\_catalog\_policy\_tag\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.datacatalog.PolicyTagIamMember("member",
            policy_tag=google_data_catalog_policy_tag["basic_policy_tag"]["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* {{policy_tag}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor "{{policy_tag}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor "{{policy_tag}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor {{policy_tag}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_tag: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyTagIamMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Data catalog PolicyTag. Each of these resources serves a different use case:

        * `datacatalog.PolicyTagIamPolicy`: Authoritative. Sets the IAM policy for the policytag and replaces any existing policy already attached.
        * `datacatalog.PolicyTagIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the policytag are preserved.
        * `datacatalog.PolicyTagIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the policytag are preserved.

        > **Note:** `datacatalog.PolicyTagIamPolicy` **cannot** be used in conjunction with `datacatalog.PolicyTagIamBinding` and `datacatalog.PolicyTagIamMember` or they will fight over what your policy should be.

        > **Note:** `datacatalog.PolicyTagIamBinding` resources **can be** used in conjunction with `datacatalog.PolicyTagIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_data\_catalog\_policy\_tag\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.datacatalog.PolicyTagIamPolicy("policy",
            policy_tag=google_data_catalog_policy_tag["basic_policy_tag"]["name"],
            policy_data=admin.policy_data)
        ```

        ## google\_data\_catalog\_policy\_tag\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.datacatalog.PolicyTagIamBinding("binding",
            policy_tag=google_data_catalog_policy_tag["basic_policy_tag"]["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_data\_catalog\_policy\_tag\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.datacatalog.PolicyTagIamMember("member",
            policy_tag=google_data_catalog_policy_tag["basic_policy_tag"]["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* {{policy_tag}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor "{{policy_tag}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor "{{policy_tag}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:datacatalog/policyTagIamMember:PolicyTagIamMember editor {{policy_tag}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param PolicyTagIamMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyTagIamMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['PolicyTagIamMemberConditionArgs']]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 policy_tag: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
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

            __props__['condition'] = condition
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__['member'] = member
            if policy_tag is None and not opts.urn:
                raise TypeError("Missing required property 'policy_tag'")
            __props__['policy_tag'] = policy_tag
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(PolicyTagIamMember, __self__).__init__(
            'gcp:datacatalog/policyTagIamMember:PolicyTagIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['PolicyTagIamMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            policy_tag: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'PolicyTagIamMember':
        """
        Get an existing PolicyTagIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_tag: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["member"] = member
        __props__["policy_tag"] = policy_tag
        __props__["role"] = role
        return PolicyTagIamMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.PolicyTagIamMemberCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter(name="policyTag")
    def policy_tag(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "policy_tag")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

