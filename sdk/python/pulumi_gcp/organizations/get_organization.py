# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetOrganizationResult:
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, create_time=None, directory_customer_id=None, domain=None, id=None, lifecycle_state=None, name=None, org_id=None, organization=None):
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
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
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
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        __self__.org_id = org_id
        """
        The Organization ID.
        """
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        __self__.organization = organization


class AwaitableGetOrganizationResult(GetOrganizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationResult(
            create_time=self.create_time,
            directory_customer_id=self.directory_customer_id,
            domain=self.domain,
            id=self.id,
            lifecycle_state=self.lifecycle_state,
            name=self.name,
            org_id=self.org_id,
            organization=self.organization)


def get_organization(domain=None, organization=None, opts=None):
    """
    Use this data source to get information about a Google Cloud Organization.

    ```python
    import pulumi
    import pulumi_gcp as gcp

    org = gcp.organizations.get_organization(domain="example.com")
    sales = gcp.organizations.Folder("sales",
        display_name="Sales",
        parent=org.name)
    ```


    :param str domain: The domain name of the Organization.
    :param str organization: The name of the Organization in the form `{organization_id}` or `organizations/{organization_id}`.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['organization'] = organization
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getOrganization:getOrganization', __args__, opts=opts).value

    return AwaitableGetOrganizationResult(
        create_time=__ret__.get('createTime'),
        directory_customer_id=__ret__.get('directoryCustomerId'),
        domain=__ret__.get('domain'),
        id=__ret__.get('id'),
        lifecycle_state=__ret__.get('lifecycleState'),
        name=__ret__.get('name'),
        org_id=__ret__.get('orgId'),
        organization=__ret__.get('organization'))
