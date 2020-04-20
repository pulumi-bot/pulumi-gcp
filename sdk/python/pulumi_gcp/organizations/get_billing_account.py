# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetBillingAccountResult:
    """
    A collection of values returned by getBillingAccount.
    """
    def __init__(__self__, billing_account=None, display_name=None, id=None, name=None, open=None, project_ids=None):
        if billing_account and not isinstance(billing_account, str):
            raise TypeError("Expected argument 'billing_account' to be a str")
        __self__.billing_account = billing_account
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
        """
        if open and not isinstance(open, bool):
            raise TypeError("Expected argument 'open' to be a bool")
        __self__.open = open
        if project_ids and not isinstance(project_ids, list):
            raise TypeError("Expected argument 'project_ids' to be a list")
        __self__.project_ids = project_ids
        """
        The IDs of any projects associated with the billing account.
        """
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

def get_billing_account(billing_account=None,display_name=None,open=None,opts=None):
    """
    Use this data source to get information about a Google Billing Account.


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
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getBillingAccount:getBillingAccount', __args__, opts=opts).value

    return AwaitableGetBillingAccountResult(
        billing_account=__ret__.get('billingAccount'),
        display_name=__ret__.get('displayName'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        open=__ret__.get('open'),
        project_ids=__ret__.get('projectIds'))
