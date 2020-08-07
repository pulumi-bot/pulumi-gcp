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

__all__ = ['KeyRingIAMMember']


class KeyRingIAMMember(pulumi.CustomResource):
    condition: pulumi.Output[Optional['outputs.KeyRingIAMMemberCondition']] = pulumi.property("condition")
    """
    An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
    Structure is documented below.
    """

    etag: pulumi.Output[str] = pulumi.property("etag")
    """
    (Computed) The etag of the key ring's IAM policy.
    """

    key_ring_id: pulumi.Output[str] = pulumi.property("keyRingId")
    """
    The key ring ID, in the form
    `{project_id}/{location_name}/{key_ring_name}` or
    `{location_name}/{key_ring_name}`. In the second form, the provider's
    project setting will be used as a fallback.
    """

    member: pulumi.Output[str] = pulumi.property("member")

    role: pulumi.Output[str] = pulumi.property("role")
    """
    The role that should be applied. Only one
    `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['KeyRingIAMMemberConditionArgs']]] = None,
                 key_ring_id: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:

        * `kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
        * `kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
        * `kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.

        > **Note:** `kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `kms.KeyRingIAMBinding` and `kms.KeyRingIAMMember` or they will fight over what your policy should be.

        > **Note:** `kms.KeyRingIAMBinding` resources **can be** used in conjunction with `kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\_kms\_key\_ring\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/editor",
            "members": ["user:jane@example.com"],
        }])
        key_ring = gcp.kms.KeyRingIAMPolicy("keyRing",
            key_ring_id=keyring.id,
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/editor",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            },
        }])
        key_ring = gcp.kms.KeyRingIAMPolicy("keyRing",
            key_ring_id=keyring.id,
            policy_data=admin.policy_data)
        ```

        ## google\_kms\_key\_ring\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRingIAMBinding("keyRing",
            key_ring_id="your-key-ring-id",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRingIAMBinding("keyRing",
            condition={
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            key_ring_id="your-key-ring-id",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\_kms\_key\_ring\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRingIAMMember("keyRing",
            key_ring_id="your-key-ring-id",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRingIAMMember("keyRing",
            condition={
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            key_ring_id="your-key-ring-id",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['KeyRingIAMMemberConditionArgs']] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] key_ring_id: The key ring ID, in the form
               `{project_id}/{location_name}/{key_ring_name}` or
               `{location_name}/{key_ring_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
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
            if key_ring_id is None:
                raise TypeError("Missing required property 'key_ring_id'")
            __props__['key_ring_id'] = key_ring_id
            if member is None:
                raise TypeError("Missing required property 'member'")
            __props__['member'] = member
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(KeyRingIAMMember, __self__).__init__(
            'gcp:kms/keyRingIAMMember:KeyRingIAMMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['KeyRingIAMMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            key_ring_id: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'KeyRingIAMMember':
        """
        Get an existing KeyRingIAMMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['KeyRingIAMMemberConditionArgs']] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the key ring's IAM policy.
        :param pulumi.Input[str] key_ring_id: The key ring ID, in the form
               `{project_id}/{location_name}/{key_ring_name}` or
               `{location_name}/{key_ring_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["key_ring_id"] = key_ring_id
        __props__["member"] = member
        __props__["role"] = role
        return KeyRingIAMMember(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

