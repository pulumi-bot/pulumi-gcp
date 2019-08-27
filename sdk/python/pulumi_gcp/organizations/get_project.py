# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, app_engines=None, auto_create_network=None, billing_account=None, folder_id=None, labels=None, name=None, number=None, org_id=None, policy_data=None, policy_etag=None, project_id=None, skip_delete=None, id=None):
        if app_engines and not isinstance(app_engines, list):
            raise TypeError("Expected argument 'app_engines' to be a list")
        __self__.app_engines = app_engines
        if auto_create_network and not isinstance(auto_create_network, bool):
            raise TypeError("Expected argument 'auto_create_network' to be a bool")
        __self__.auto_create_network = auto_create_network
        if billing_account and not isinstance(billing_account, str):
            raise TypeError("Expected argument 'billing_account' to be a str")
        __self__.billing_account = billing_account
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        __self__.folder_id = folder_id
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if number and not isinstance(number, str):
            raise TypeError("Expected argument 'number' to be a str")
        __self__.number = number
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        __self__.org_id = org_id
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        __self__.policy_data = policy_data
        if policy_etag and not isinstance(policy_etag, str):
            raise TypeError("Expected argument 'policy_etag' to be a str")
        __self__.policy_etag = policy_etag
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        if skip_delete and not isinstance(skip_delete, bool):
            raise TypeError("Expected argument 'skip_delete' to be a bool")
        __self__.skip_delete = skip_delete
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            app_engines=self.app_engines,
            auto_create_network=self.auto_create_network,
            billing_account=self.billing_account,
            folder_id=self.folder_id,
            labels=self.labels,
            name=self.name,
            number=self.number,
            org_id=self.org_id,
            policy_data=self.policy_data,
            policy_etag=self.policy_etag,
            project_id=self.project_id,
            skip_delete=self.skip_delete,
            id=self.id)

def get_project(project_id=None,opts=None):
    """
    Use this data source to get project details.
    For more information see 
    [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects#Project)
    
    :param str project_id: The project ID. If it is not provided, the provider project is used.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/project.html.markdown.
    """
    __args__ = dict()

    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getProject:getProject', __args__, opts=opts).value

    return AwaitableGetProjectResult(
        app_engines=__ret__.get('appEngines'),
        auto_create_network=__ret__.get('autoCreateNetwork'),
        billing_account=__ret__.get('billingAccount'),
        folder_id=__ret__.get('folderId'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        number=__ret__.get('number'),
        org_id=__ret__.get('orgId'),
        policy_data=__ret__.get('policyData'),
        policy_etag=__ret__.get('policyEtag'),
        project_id=__ret__.get('projectId'),
        skip_delete=__ret__.get('skipDelete'),
        id=__ret__.get('id'))
