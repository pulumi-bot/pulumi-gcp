# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetOrganizationPolicyResult:
    """
    A collection of values returned by getOrganizationPolicy.
    """
    def __init__(__self__, boolean_policies=None, constraint=None, etag=None, list_policies=None, project=None, restore_policies=None, update_time=None, version=None, id=None):
        if boolean_policies and not isinstance(boolean_policies, list):
            raise TypeError("Expected argument 'boolean_policies' to be a list")
        __self__.boolean_policies = boolean_policies
        if constraint and not isinstance(constraint, str):
            raise TypeError("Expected argument 'constraint' to be a str")
        __self__.constraint = constraint
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        __self__.etag = etag
        if list_policies and not isinstance(list_policies, list):
            raise TypeError("Expected argument 'list_policies' to be a list")
        __self__.list_policies = list_policies
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if restore_policies and not isinstance(restore_policies, list):
            raise TypeError("Expected argument 'restore_policies' to be a list")
        __self__.restore_policies = restore_policies
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        __self__.update_time = update_time
        if version and not isinstance(version, float):
            raise TypeError("Expected argument 'version' to be a float")
        __self__.version = version
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_organization_policy(constraint=None,project=None,opts=None):
    """
    Allows management of Organization policies for a Google Project. For more information see
    [the official
    documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/project_organization_policy.html.markdown.
    """
    __args__ = dict()

    __args__['constraint'] = constraint
    __args__['project'] = project
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('gcp:projects/getOrganizationPolicy:getOrganizationPolicy', __args__, opts=opts)

    return GetOrganizationPolicyResult(
        boolean_policies=__ret__.get('booleanPolicies'),
        constraint=__ret__.get('constraint'),
        etag=__ret__.get('etag'),
        list_policies=__ret__.get('listPolicies'),
        project=__ret__.get('project'),
        restore_policies=__ret__.get('restorePolicies'),
        update_time=__ret__.get('updateTime'),
        version=__ret__.get('version'),
        id=__ret__.get('id'))
