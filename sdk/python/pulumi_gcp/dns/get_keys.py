# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetKeysResult',
    'AwaitableGetKeysResult',
    'get_keys',
]

@pulumi.output_type
class GetKeysResult:
    """
    A collection of values returned by getKeys.
    """
    def __init__(__self__, id=None, key_signing_keys=None, managed_zone=None, project=None, zone_signing_keys=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_signing_keys and not isinstance(key_signing_keys, list):
            raise TypeError("Expected argument 'key_signing_keys' to be a list")
        pulumi.set(__self__, "key_signing_keys", key_signing_keys)
        if managed_zone and not isinstance(managed_zone, str):
            raise TypeError("Expected argument 'managed_zone' to be a str")
        pulumi.set(__self__, "managed_zone", managed_zone)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if zone_signing_keys and not isinstance(zone_signing_keys, list):
            raise TypeError("Expected argument 'zone_signing_keys' to be a list")
        pulumi.set(__self__, "zone_signing_keys", zone_signing_keys)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keySigningKeys")
    def key_signing_keys(self) -> Sequence['outputs.GetKeysKeySigningKeyResult']:
        """
        A list of Key-signing key (KSK) records. Structure is documented below. Additionally, the DS record is provided:
        """
        return pulumi.get(self, "key_signing_keys")

    @property
    @pulumi.getter(name="managedZone")
    def managed_zone(self) -> str:
        return pulumi.get(self, "managed_zone")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="zoneSigningKeys")
    def zone_signing_keys(self) -> Sequence['outputs.GetKeysZoneSigningKeyResult']:
        """
        A list of Zone-signing key (ZSK) records. Structure is documented below.
        """
        return pulumi.get(self, "zone_signing_keys")


class AwaitableGetKeysResult(GetKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeysResult(
            id=self.id,
            key_signing_keys=self.key_signing_keys,
            managed_zone=self.managed_zone,
            project=self.project,
            zone_signing_keys=self.zone_signing_keys)


def get_keys(managed_zone: Optional[str] = None,
             project: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeysResult:
    """
    Get the DNSKEY and DS records of DNSSEC-signed managed zones. For more information see the
    [official documentation](https://cloud.google.com/dns/docs/dnskeys/)
    and [API](https://cloud.google.com/dns/docs/reference/v1/dnsKeys).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    foo = gcp.dns.ManagedZone("foo",
        dns_name="foo.bar.",
        dnssec_config=gcp.dns.ManagedZoneDnssecConfigArgs(
            state="on",
            non_existence="nsec3",
        ))
    foo_dns_keys = foo.id.apply(lambda id: gcp.dns.get_keys(managed_zone=id))
    pulumi.export("fooDnsDsRecord", foo_dns_keys.key_signing_keys[0].ds_record)
    ```


    :param str managed_zone: The name or id of the Cloud DNS managed zone.
    :param str project: The ID of the project in which the resource belongs. If `project` is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['managedZone'] = managed_zone
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:dns/getKeys:getKeys', __args__, opts=opts, typ=GetKeysResult).value

    return AwaitableGetKeysResult(
        id=__ret__.id,
        key_signing_keys=__ret__.key_signing_keys,
        managed_zone=__ret__.managed_zone,
        project=__ret__.project,
        zone_signing_keys=__ret__.zone_signing_keys)
