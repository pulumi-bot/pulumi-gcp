# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
]

@pulumi.output_type
class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, account_id=None, display_name=None, email=None, id=None, name=None, project=None, unique_id=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if unique_id and not isinstance(unique_id, str):
            raise TypeError("Expected argument 'unique_id' to be a str")
        pulumi.set(__self__, "unique_id", unique_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name for the service account.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The e-mail address of the service account. This value
        should be referenced from any `organizations.getIAMPolicy` data sources
        that would grant the service account privileges.
        """
        return pulumi.get(self, "email")

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
        The fully-qualified name of the service account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> str:
        """
        The unique id of the service account.
        """
        return pulumi.get(self, "unique_id")


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            account_id=self.account_id,
            display_name=self.display_name,
            email=self.email,
            id=self.id,
            name=self.name,
            project=self.project,
            unique_id=self.unique_id)


def get_account(account_id: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    Get the service account from a project. For more information see
    the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    object_viewer = gcp.serviceAccount.get_account(account_id="object-viewer")
    ```


    :param str account_id: The Google service account ID. This be one of:
    :param str project: The ID of the project that the service account is present in.
           Defaults to the provider project configuration.
    """
    __args__ = dict()
    __args__['accountId'] = account_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:serviceAccount/getAccount:getAccount', __args__, opts=opts, typ=GetAccountResult).value

    return AwaitableGetAccountResult(
        account_id=__ret__.account_id,
        display_name=__ret__.display_name,
        email=__ret__.email,
        id=__ret__.id,
        name=__ret__.name,
        project=__ret__.project,
        unique_id=__ret__.unique_id)


@_utilities.lift_output_func(get_account)
def get_account_output(account_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountResult]:
    ...
