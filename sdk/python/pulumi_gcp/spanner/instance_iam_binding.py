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

__all__ = ['InstanceIAMBindingArgs', 'InstanceIAMBinding']

@pulumi.input_type
class InstanceIAMBindingArgs:
    def __init__(__self__, *,
                 instance: pulumi.Input[str],
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['InstanceIAMBindingConditionArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceIAMBinding resource.
        :param pulumi.Input[str] instance: The name of the instance.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        pulumi.set(__self__, "instance", instance)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Input[str]:
        """
        The name of the instance.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['InstanceIAMBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['InstanceIAMBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class InstanceIAMBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['InstanceIAMBindingConditionArgs']]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:

        * `spanner.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.

        > **Warning:** It's entirely possibly to lock yourself out of your instance using `spanner.InstanceIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.

        * `spanner.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `spanner.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.

        > **Note:** `spanner.InstanceIAMPolicy` **cannot** be used in conjunction with `spanner.InstanceIAMBinding` and `spanner.InstanceIAMMember` or they will fight over what your policy should be.

        > **Note:** `spanner.InstanceIAMBinding` resources **can be** used in conjunction with `spanner.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\_spanner\_instance\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        instance = gcp.spanner.InstanceIAMPolicy("instance",
            instance="your-instance-name",
            policy_data=admin.policy_data)
        ```

        ## google\_spanner\_instance\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.spanner.InstanceIAMBinding("instance",
            instance="your-instance-name",
            members=["user:jane@example.com"],
            role="roles/compute.networkUser")
        ```

        ## google\_spanner\_instance\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.spanner.InstanceIAMMember("instance",
            instance="your-instance-name",
            member="user:jane@example.com",
            role="roles/compute.networkUser")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* {{project}}/{{name}} * {{name}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the account, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance "project-name/instance-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance "project-name/instance-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance project-name/instance-name
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance: The name of the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceIAMBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:

        * `spanner.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.

        > **Warning:** It's entirely possibly to lock yourself out of your instance using `spanner.InstanceIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.

        * `spanner.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `spanner.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.

        > **Note:** `spanner.InstanceIAMPolicy` **cannot** be used in conjunction with `spanner.InstanceIAMBinding` and `spanner.InstanceIAMMember` or they will fight over what your policy should be.

        > **Note:** `spanner.InstanceIAMBinding` resources **can be** used in conjunction with `spanner.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\_spanner\_instance\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        instance = gcp.spanner.InstanceIAMPolicy("instance",
            instance="your-instance-name",
            policy_data=admin.policy_data)
        ```

        ## google\_spanner\_instance\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.spanner.InstanceIAMBinding("instance",
            instance="your-instance-name",
            members=["user:jane@example.com"],
            role="roles/compute.networkUser")
        ```

        ## google\_spanner\_instance\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.spanner.InstanceIAMMember("instance",
            instance="your-instance-name",
            member="user:jane@example.com",
            role="roles/compute.networkUser")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* {{project}}/{{name}} * {{name}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the account, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance "project-name/instance-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance "project-name/instance-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance project-name/instance-name
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param InstanceIAMBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceIAMBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['InstanceIAMBindingConditionArgs']]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            if instance is None and not opts.urn:
                raise TypeError("Missing required property 'instance'")
            __props__['instance'] = instance
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__['members'] = members
            __props__['project'] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(InstanceIAMBinding, __self__).__init__(
            'gcp:spanner/instanceIAMBinding:InstanceIAMBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['InstanceIAMBindingConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'InstanceIAMBinding':
        """
        Get an existing InstanceIAMBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the instance's IAM policy.
        :param pulumi.Input[str] instance: The name of the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["instance"] = instance
        __props__["members"] = members
        __props__["project"] = project
        __props__["role"] = role
        return InstanceIAMBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.InstanceIAMBindingCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the instance's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        The name of the instance.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

