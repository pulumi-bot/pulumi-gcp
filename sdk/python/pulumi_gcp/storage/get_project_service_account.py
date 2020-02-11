# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetProjectServiceAccountResult:
    """
    A collection of values returned by getProjectServiceAccount.
    """
    def __init__(__self__, email_address=None, id=None, project=None, user_project=None):
        if email_address and not isinstance(email_address, str):
            raise TypeError("Expected argument 'email_address' to be a str")
        __self__.email_address = email_address
        """
        The email address of the service account. This value is often used to refer to the service account
        in order to grant IAM permissions.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if user_project and not isinstance(user_project, str):
            raise TypeError("Expected argument 'user_project' to be a str")
        __self__.user_project = user_project
class AwaitableGetProjectServiceAccountResult(GetProjectServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectServiceAccountResult(
            email_address=self.email_address,
            id=self.id,
            project=self.project,
            user_project=self.user_project)

def get_project_service_account(project=None,user_project=None,opts=None):
    """
    Get the email address of a project's unique Google Cloud Storage service account.

    Each Google Cloud project has a unique service account for use with Google Cloud Storage. Only this
    special service account can be used to set up `storage.Notification` resources.

    For more information see
    [the API reference](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_project_service_account.html.markdown.


    :param str project: The project the unique service account was created for. If it is not provided, the provider project is used.
    :param str user_project: The project the lookup originates from. This field is used if you are making the request
           from a different account than the one you are finding the service account for.
    """
    __args__ = dict()


    __args__['project'] = project
    __args__['userProject'] = user_project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:storage/getProjectServiceAccount:getProjectServiceAccount', __args__, opts=opts).value

    return AwaitableGetProjectServiceAccountResult(
        email_address=__ret__.get('emailAddress'),
        id=__ret__.get('id'),
        project=__ret__.get('project'),
        user_project=__ret__.get('userProject'))
