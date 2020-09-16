# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetOrganizationResult',
    'AwaitableGetOrganizationResult',
    'get_organization',
]

@pulumi.output_type
class GetOrganizationResult:
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, create_time=None, directory_customer_id=None, domain=None, id=None, lifecycle_state=None, name=None, org_id=None, organization=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if directory_customer_id and not isinstance(directory_customer_id, str):
            raise TypeError("Expected argument 'directory_customer_id' to be a str")
        pulumi.set(__self__, "directory_customer_id", directory_customer_id)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lifecycle_state and not isinstance(lifecycle_state, str):
            raise TypeError("Expected argument 'lifecycle_state' to be a str")
        pulumi.set(__self__, "lifecycle_state", lifecycle_state)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="directoryCustomerId")
    def directory_customer_id(self) -> str:
        """
        The Google for Work customer ID of the Organization.
        """
        return pulumi.get(self, "directory_customer_id")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lifecycleState")
    def lifecycle_state(self) -> str:
        """
        The Organization's current lifecycle state.
        """
        return pulumi.get(self, "lifecycle_state")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the Organization in the form `organizations/{organization_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        """
        The Organization ID.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        return pulumi.get(self, "organization")


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


def get_organization(domain: Optional[str] = None,
                     organization: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationResult:
    """
    Use this data source to get information about a Google Cloud Organization.


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
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getOrganization:getOrganization', __args__, opts=opts, typ=GetOrganizationResult).value

    return AwaitableGetOrganizationResult(
        create_time=__ret__.create_time,
        directory_customer_id=__ret__.directory_customer_id,
        domain=__ret__.domain,
        id=__ret__.id,
        lifecycle_state=__ret__.lifecycle_state,
        name=__ret__.name,
        org_id=__ret__.org_id,
        organization=__ret__.organization)
