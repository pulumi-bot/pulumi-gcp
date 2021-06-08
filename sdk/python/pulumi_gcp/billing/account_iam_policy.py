# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountIamPolicyArgs', 'AccountIamPolicy']

@pulumi.input_type
class AccountIamPolicyArgs:
    def __init__(__self__, *,
                 billing_account_id: pulumi.Input[str],
                 policy_data: pulumi.Input[str]):
        """
        The set of arguments for constructing a AccountIamPolicy resource.
        """
        pulumi.set(__self__, "billing_account_id", billing_account_id)
        pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter(name="billingAccountId")
    def billing_account_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "billing_account_id")

    @billing_account_id.setter
    def billing_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "billing_account_id", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)


@pulumi.input_type
class _AccountIamPolicyState:
    def __init__(__self__, *,
                 billing_account_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccountIamPolicy resources.
        """
        if billing_account_id is not None:
            pulumi.set(__self__, "billing_account_id", billing_account_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter(name="billingAccountId")
    def billing_account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "billing_account_id")

    @billing_account_id.setter
    def billing_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_account_id", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)


class AccountIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AccountIamPolicy resource with the given unique name, props, and options.
        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: AccountIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AccountIamPolicy resource with the given unique name, props, and options.
        :param str resource_name_: The name of the resource.
        :param AccountIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountIamPolicyArgs.__new__(AccountIamPolicyArgs)

            if billing_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'billing_account_id'")
            __props__.__dict__["billing_account_id"] = billing_account_id
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["etag"] = None
        super(AccountIamPolicy, __self__).__init__(
            'gcp:billing/accountIamPolicy:AccountIamPolicy',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            billing_account_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None) -> 'AccountIamPolicy':
        """
        Get an existing AccountIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountIamPolicyState.__new__(_AccountIamPolicyState)

        __props__.__dict__["billing_account_id"] = billing_account_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        return AccountIamPolicy(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingAccountId")
    def billing_account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "billing_account_id")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        return pulumi.get(self, "policy_data")

