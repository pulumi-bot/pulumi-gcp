# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['SubAccountArgs', 'SubAccount']

@pulumi.input_type
class SubAccountArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 master_billing_account: pulumi.Input[str],
                 deletion_policy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SubAccount resource.
        :param pulumi.Input[str] display_name: The display name of the billing account.
        :param pulumi.Input[str] master_billing_account: The name of the master billing account that the subaccount
               will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        :param pulumi.Input[str] deletion_policy: If set to "RENAME_ON_DESTROY" the billing account display_name
               will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
               Default is "".
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "master_billing_account", master_billing_account)
        if deletion_policy is not None:
            pulumi.set(__self__, "deletion_policy", deletion_policy)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the billing account.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="masterBillingAccount")
    def master_billing_account(self) -> pulumi.Input[str]:
        """
        The name of the master billing account that the subaccount
        will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        """
        return pulumi.get(self, "master_billing_account")

    @master_billing_account.setter
    def master_billing_account(self, value: pulumi.Input[str]):
        pulumi.set(self, "master_billing_account", value)

    @property
    @pulumi.getter(name="deletionPolicy")
    def deletion_policy(self) -> Optional[pulumi.Input[str]]:
        """
        If set to "RENAME_ON_DESTROY" the billing account display_name
        will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
        Default is "".
        """
        return pulumi.get(self, "deletion_policy")

    @deletion_policy.setter
    def deletion_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deletion_policy", value)


class SubAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deletion_policy: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 master_billing_account: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows creation and management of a Google Cloud Billing Subaccount.

        !> **WARNING:** Deleting this resource will not delete or close the billing subaccount.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        subaccount = gcp.billing.SubAccount("subaccount",
            display_name="My Billing Account",
            master_billing_account="012345-567890-ABCDEF")
        ```

        ## Import

        Billing Subaccounts can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:billing/subAccount:SubAccount default billingAccounts/{billing_account_id}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deletion_policy: If set to "RENAME_ON_DESTROY" the billing account display_name
               will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
               Default is "".
        :param pulumi.Input[str] display_name: The display name of the billing account.
        :param pulumi.Input[str] master_billing_account: The name of the master billing account that the subaccount
               will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows creation and management of a Google Cloud Billing Subaccount.

        !> **WARNING:** Deleting this resource will not delete or close the billing subaccount.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        subaccount = gcp.billing.SubAccount("subaccount",
            display_name="My Billing Account",
            master_billing_account="012345-567890-ABCDEF")
        ```

        ## Import

        Billing Subaccounts can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:billing/subAccount:SubAccount default billingAccounts/{billing_account_id}
        ```

        :param str resource_name: The name of the resource.
        :param SubAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deletion_policy: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 master_billing_account: Optional[pulumi.Input[str]] = None,
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

            __props__['deletion_policy'] = deletion_policy
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if master_billing_account is None and not opts.urn:
                raise TypeError("Missing required property 'master_billing_account'")
            __props__['master_billing_account'] = master_billing_account
            __props__['billing_account_id'] = None
            __props__['name'] = None
            __props__['open'] = None
        super(SubAccount, __self__).__init__(
            'gcp:billing/subAccount:SubAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            billing_account_id: Optional[pulumi.Input[str]] = None,
            deletion_policy: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            master_billing_account: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            open: Optional[pulumi.Input[bool]] = None) -> 'SubAccount':
        """
        Get an existing SubAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account_id: The billing account id.
        :param pulumi.Input[str] deletion_policy: If set to "RENAME_ON_DESTROY" the billing account display_name
               will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
               Default is "".
        :param pulumi.Input[str] display_name: The display name of the billing account.
        :param pulumi.Input[str] master_billing_account: The name of the master billing account that the subaccount
               will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        :param pulumi.Input[str] name: The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
        :param pulumi.Input[bool] open: `true` if the billing account is open, `false` if the billing account is closed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["billing_account_id"] = billing_account_id
        __props__["deletion_policy"] = deletion_policy
        __props__["display_name"] = display_name
        __props__["master_billing_account"] = master_billing_account
        __props__["name"] = name
        __props__["open"] = open
        return SubAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingAccountId")
    def billing_account_id(self) -> pulumi.Output[str]:
        """
        The billing account id.
        """
        return pulumi.get(self, "billing_account_id")

    @property
    @pulumi.getter(name="deletionPolicy")
    def deletion_policy(self) -> pulumi.Output[Optional[str]]:
        """
        If set to "RENAME_ON_DESTROY" the billing account display_name
        will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
        Default is "".
        """
        return pulumi.get(self, "deletion_policy")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the billing account.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="masterBillingAccount")
    def master_billing_account(self) -> pulumi.Output[str]:
        """
        The name of the master billing account that the subaccount
        will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        """
        return pulumi.get(self, "master_billing_account")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def open(self) -> pulumi.Output[bool]:
        """
        `true` if the billing account is open, `false` if the billing account is closed.
        """
        return pulumi.get(self, "open")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

