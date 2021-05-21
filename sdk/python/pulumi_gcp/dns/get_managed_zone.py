# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetManagedZoneResult',
    'AwaitableGetManagedZoneResult',
    'get_managed_zone',
]

@pulumi.output_type
class GetManagedZoneResult:
    """
    A collection of values returned by getManagedZone.
    """
    def __init__(__self__, description=None, dns_name=None, id=None, name=None, name_servers=None, project=None, visibility=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_servers and not isinstance(name_servers, list):
            raise TypeError("Expected argument 'name_servers' to be a list")
        pulumi.set(__self__, "name_servers", name_servers)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A textual description field.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        """
        The fully qualified DNS name of this zone, e.g. `example.io.`.
        """
        return pulumi.get(self, "dns_name")

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
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> Sequence[str]:
        """
        The list of nameservers that will be authoritative for this
        domain. Use NS records to redirect from your DNS provider to these names,
        thus making Google Cloud DNS authoritative for this zone.
        """
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        The zone's visibility: public zones are exposed to the Internet,
        while private zones are visible only to Virtual Private Cloud resources.
        """
        return pulumi.get(self, "visibility")


class AwaitableGetManagedZoneResult(GetManagedZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedZoneResult(
            description=self.description,
            dns_name=self.dns_name,
            id=self.id,
            name=self.name,
            name_servers=self.name_servers,
            project=self.project,
            visibility=self.visibility)


def get_managed_zone(name: Optional[str] = None,
                     project: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagedZoneResult:
    """
    Provides access to a zone's attributes within Google Cloud DNS.
    For more information see
    [the official documentation](https://cloud.google.com/dns/zones/)
    and
    [API](https://cloud.google.com/dns/api/v1/managedZones).

    ```python
    import pulumi
    import pulumi_gcp as gcp

    env_dns_zone = gcp.dns.get_managed_zone(name="qa-zone")
    dns = gcp.dns.RecordSet("dns",
        name=f"my-address.{env_dns_zone.dns_name}",
        type="TXT",
        ttl=300,
        managed_zone=env_dns_zone.name,
        rrdatas=["test"])
    ```


    :param str name: A unique name for the resource.
    :param str project: The ID of the project for the Google Cloud DNS zone.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:dns/getManagedZone:getManagedZone', __args__, opts=opts, typ=GetManagedZoneResult).value

    return AwaitableGetManagedZoneResult(
        description=__ret__.description,
        dns_name=__ret__.dns_name,
        id=__ret__.id,
        name=__ret__.name,
        name_servers=__ret__.name_servers,
        project=__ret__.project,
        visibility=__ret__.visibility)


@_utilities.lift_output_func(get_managed_zone)
def get_managed_zone_output(name: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetManagedZoneResult]:
    ...
