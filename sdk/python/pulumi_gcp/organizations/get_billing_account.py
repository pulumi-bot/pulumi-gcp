# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetBillingAccountResult',
    'AwaitableGetBillingAccountResult',
    'get_billing_account',
]

@pulumi.output_type
class GetBillingAccountResult:
    """
    A collection of values returned by getBillingAccount.
    """
    def __init__(__self__, billing_account=None, display_name=None, id=None, name=None, open=None, project_ids=None):
        if billing_account and not isinstance(billing_account, str):
            raise TypeError("Expected argument 'billing_account' to be a str")
        pulumi.set(__self__, "billing_account", billing_account)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if open and not isinstance(open, bool):
            raise TypeError("Expected argument 'open' to be a bool")
        pulumi.set(__self__, "open", open)
        if project_ids and not isinstance(project_ids, list):
            raise TypeError("Expected argument 'project_ids' to be a list")
        pulumi.set(__self__, "project_ids", project_ids)

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> Optional[str]:
        return pulumi.get(self, "billing_account")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def open(self) -> bool:
        return pulumi.get(self, "open")

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> Sequence[str]:
        """
        The IDs of any projects associated with the billing account.
        """
        return pulumi.get(self, "project_ids")


class AwaitableGetBillingAccountResult(GetBillingAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBillingAccountResult(
            billing_account=self.billing_account,
            display_name=self.display_name,
            id=self.id,
            name=self.name,
            open=self.open,
            project_ids=self.project_ids)


def get_billing_account(billing_account: Optional[str] = None,
                        display_name: Optional[str] = None,
                        open: Optional[bool] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBillingAccountResult:
    """
    Use this data source to get information about a Google Billing Account.

    ```python
    import pulumi
    import pulumi_gcp as gcp

    acct = gcp.organizations.get_billing_account(display_name="My Billing Account",
        open=True)
    my_project = gcp.organizations.Project("myProject",
        project_id="your-project-id",
        org_id="1234567",
        billing_account=acct.id)
    ```


    :param str billing_account: The name of the billing account in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
    :param str display_name: The display name of the billing account.
    :param bool open: `true` if the billing account is open, `false` if the billing account is closed.
    """
    __args__ = dict()
    __args__['billingAccount'] = billing_account
    __args__['displayName'] = display_name
    __args__['open'] = open
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getBillingAccount:getBillingAccount', __args__, opts=opts, typ=GetBillingAccountResult).value

    return AwaitableGetBillingAccountResult(
        billing_account=__ret__.billing_account,
        display_name=__ret__.display_name,
        id=__ret__.id,
        name=__ret__.name,
        open=__ret__.open,
        project_ids=__ret__.project_ids)


@_utilities.lift_output_func(get_billing_account)
def get_billing_account_apply(billing_account: Optional[pulumi.Input[str]] = None,
                              display_name: Optional[pulumi.Input[str]] = None,
                              open: Optional[pulumi.Input[bool]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBillingAccountResult]:
    ...
