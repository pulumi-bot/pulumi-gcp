# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClientOpenIdUserInfoResult:
    """
    A collection of values returned by getClientOpenIdUserInfo.
    """
    def __init__(__self__, email=None, id=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        __self__.email = email
        """
        The email of the account used by the provider to authenticate with GCP.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
class AwaitableGetClientOpenIdUserInfoResult(GetClientOpenIdUserInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientOpenIdUserInfoResult(
            email=self.email,
            id=self.id)

def get_client_open_id_user_info(opts=None):
    """
    Get OpenID userinfo about the credentials used with the Google provider,
    specifically the email.

    This datasource enables you to export the email of the account you've
    authenticated the provider with; this can be used alongside
    `data.google_client_config`'s `access_token` to perform OpenID Connect
    authentication with GKE and configure an RBAC role for the email used.

    > This resource will only work as expected if the provider is configured to
    use the `https://www.googleapis.com/auth/userinfo.email` scope! You will
    receive an error otherwise.
    """
    __args__ = dict()


    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getClientOpenIdUserInfo:getClientOpenIdUserInfo', __args__, opts=opts).value

    return AwaitableGetClientOpenIdUserInfoResult(
        email=__ret__.get('email'),
        id=__ret__.get('id'))
