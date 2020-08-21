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

__all__ = ['AccountIamBinding']


class AccountIamBinding(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account_id: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['AccountIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows creation and management of a single binding within IAM policy for
        an existing Google Cloud Platform Billing Account.

        > **Note:** This resource __must not__ be used in conjunction with
           `billing.AccountIamMember` for the __same role__ or they will fight over
           what your policy should be.

        > **Note:** On create, this resource will overwrite members of any existing roles.
            Use `pulumi import` and inspect the output to ensure
            your existing members are preserved.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account_id: The billing account id.
        :param pulumi.Input[List[pulumi.Input[str]]] members: A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        :param pulumi.Input[str] role: The role that should be applied.
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

            if billing_account_id is None:
                raise TypeError("Missing required property 'billing_account_id'")
            __props__['billing_account_id'] = billing_account_id
            __props__['condition'] = condition
            if members is None:
                raise TypeError("Missing required property 'members'")
            __props__['members'] = members
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(AccountIamBinding, __self__).__init__(
            'gcp:billing/accountIamBinding:AccountIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            billing_account_id: Optional[pulumi.Input[str]] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['AccountIamBindingConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'AccountIamBinding':
        """
        Get an existing AccountIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account_id: The billing account id.
        :param pulumi.Input[str] etag: (Computed) The etag of the billing account's IAM policy.
        :param pulumi.Input[List[pulumi.Input[str]]] members: A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        :param pulumi.Input[str] role: The role that should be applied.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["billing_account_id"] = billing_account_id
        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["members"] = members
        __props__["role"] = role
        return AccountIamBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingAccountId")
    def billing_account_id(self) -> str:
        """
        The billing account id.
        """
        return pulumi.get(self, "billing_account_id")

    @property
    @pulumi.getter
    def condition(self) -> Optional['outputs.AccountIamBindingCondition']:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        (Computed) The etag of the billing account's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def members(self) -> List[str]:
        """
        A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The role that should be applied.
        """
        return pulumi.get(self, "role")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

