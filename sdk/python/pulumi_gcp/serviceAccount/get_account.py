# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAccountResult:
    """
    A collection of values returned by getAccount.
    """
    def __init__(__self__, account_id=None, display_name=None, email=None, id=None, name=None, project=None, unique_id=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        __self__.account_id = account_id
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        """
        The display name for the service account.
        """
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        __self__.email = email
        """
        The e-mail address of the service account. This value
        should be referenced from any `organizations.getIAMPolicy` data sources
        that would grant the service account privileges.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The fully-qualified name of the service account.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if unique_id and not isinstance(unique_id, str):
            raise TypeError("Expected argument 'unique_id' to be a str")
        __self__.unique_id = unique_id
        """
        The unique id of the service account.
        """
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

def get_account(account_id=None,project=None,opts=None):
    """
    Get the service account from a project. For more information see
    the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_service_account.html.markdown.


    :param str account_id: The Service account id.  (This is the part of the service account's email field that comes before the @ symbol.)
    :param str project: The ID of the project that the service account is present in.
           Defaults to the provider project configuration.
    """
    __args__ = dict()


    __args__['accountId'] = account_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:serviceAccount/getAccount:getAccount', __args__, opts=opts).value

    return AwaitableGetAccountResult(
        account_id=__ret__.get('accountId'),
        display_name=__ret__.get('displayName'),
        email=__ret__.get('email'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        unique_id=__ret__.get('uniqueId'))
