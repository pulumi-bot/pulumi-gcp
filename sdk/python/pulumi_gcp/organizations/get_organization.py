# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetOrganizationResult:
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, create_time=None, directory_customer_id=None, domain=None, lifecycle_state=None, name=None, organization=None, id=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        __self__.create_time = create_time
        """
        Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        if directory_customer_id and not isinstance(directory_customer_id, str):
            raise TypeError("Expected argument 'directory_customer_id' to be a str")
        __self__.directory_customer_id = directory_customer_id
        """
        The Google for Work customer ID of the Organization.
        """
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        __self__.domain = domain
        if lifecycle_state and not isinstance(lifecycle_state, str):
            raise TypeError("Expected argument 'lifecycle_state' to be a str")
        __self__.lifecycle_state = lifecycle_state
        """
        The Organization's current lifecycle state.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The resource name of the Organization in the form `organizations/{organization_id}`.
        """
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        __self__.organization = organization
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        delattr(self, "__await__")
        delattr(self, "__iter__")
        return self

    __iter__ = __await__

def get_organization(domain=None,organization=None,opts=None):
    """
    Use this data source to get information about a Google Cloud Organization.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/organization.html.markdown.
    """
    __args__ = dict()

    __args__['domain'] = domain
    __args__['organization'] = organization
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getOrganization:getOrganization', __args__, opts=opts).value

    return GetOrganizationResult(
        create_time=__ret__.get('createTime'),
        directory_customer_id=__ret__.get('directoryCustomerId'),
        domain=__ret__.get('domain'),
        lifecycle_state=__ret__.get('lifecycleState'),
        name=__ret__.get('name'),
        organization=__ret__.get('organization'),
        id=__ret__.get('id'))
