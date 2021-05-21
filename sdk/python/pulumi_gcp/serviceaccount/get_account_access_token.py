# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAccountAccessTokenResult',
    'AwaitableGetAccountAccessTokenResult',
    'get_account_access_token',
]

@pulumi.output_type
class GetAccountAccessTokenResult:
    """
    A collection of values returned by getAccountAccessToken.
    """
    def __init__(__self__, access_token=None, delegates=None, id=None, lifetime=None, scopes=None, target_service_account=None):
        if access_token and not isinstance(access_token, str):
            raise TypeError("Expected argument 'access_token' to be a str")
        pulumi.set(__self__, "access_token", access_token)
        if delegates and not isinstance(delegates, list):
            raise TypeError("Expected argument 'delegates' to be a list")
        pulumi.set(__self__, "delegates", delegates)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lifetime and not isinstance(lifetime, str):
            raise TypeError("Expected argument 'lifetime' to be a str")
        pulumi.set(__self__, "lifetime", lifetime)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if target_service_account and not isinstance(target_service_account, str):
            raise TypeError("Expected argument 'target_service_account' to be a str")
        pulumi.set(__self__, "target_service_account", target_service_account)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> str:
        """
        The `access_token` representing the new generated identity.
        """
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter
    def delegates(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "delegates")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lifetime(self) -> Optional[str]:
        return pulumi.get(self, "lifetime")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence[str]:
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter(name="targetServiceAccount")
    def target_service_account(self) -> str:
        return pulumi.get(self, "target_service_account")


class AwaitableGetAccountAccessTokenResult(GetAccountAccessTokenResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountAccessTokenResult(
            access_token=self.access_token,
            delegates=self.delegates,
            id=self.id,
            lifetime=self.lifetime,
            scopes=self.scopes,
            target_service_account=self.target_service_account)


def get_account_access_token(delegates: Optional[Sequence[str]] = None,
                             lifetime: Optional[str] = None,
                             scopes: Optional[Sequence[str]] = None,
                             target_service_account: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountAccessTokenResult:
    """
    This data source provides a google `oauth2` `access_token` for a different service account than the one initially running the script.

    For more information see
    [the official documentation](https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials) as well as [iamcredentials.generateAccessToken()](https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken)


    :param Sequence[str] delegates: Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
    :param str lifetime: Lifetime of the impersonated token (defaults to its max: `3600s`).
    :param Sequence[str] scopes: The scopes the new credential should have (e.g. `["cloud-platform"]`)
    :param str target_service_account: The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
    """
    __args__ = dict()
    __args__['delegates'] = delegates
    __args__['lifetime'] = lifetime
    __args__['scopes'] = scopes
    __args__['targetServiceAccount'] = target_service_account
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:serviceAccount/getAccountAccessToken:getAccountAccessToken', __args__, opts=opts, typ=GetAccountAccessTokenResult).value

    return AwaitableGetAccountAccessTokenResult(
        access_token=__ret__.access_token,
        delegates=__ret__.delegates,
        id=__ret__.id,
        lifetime=__ret__.lifetime,
        scopes=__ret__.scopes,
        target_service_account=__ret__.target_service_account)


@_utilities.lift_output_func(get_account_access_token)
def get_account_access_token_output(delegates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                                    lifetime: Optional[pulumi.Input[str]] = None,
                                    scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                                    target_service_account: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountAccessTokenResult]:
    ...
