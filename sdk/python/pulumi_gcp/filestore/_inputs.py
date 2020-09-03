# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'InstanceFileSharesArgs',
    'InstanceFileSharesNfsExportOptionArgs',
    'InstanceNetworkArgs',
]

@pulumi.input_type
class InstanceFileSharesArgs:
    def __init__(__self__, *,
                 capacity_gb: pulumi.Input[float],
                 name: pulumi.Input[str],
                 nfs_export_options: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceFileSharesNfsExportOptionArgs']]]] = None):
        """
        :param pulumi.Input[float] capacity_gb: File share capacity in GiB. This must be at least 1024 GiB
               for the standard tier, or 2560 GiB for the premium tier.
        :param pulumi.Input[str] name: The name of the fileshare (16 characters or less)
        """
        pulumi.set(__self__, "capacity_gb", capacity_gb)
        pulumi.set(__self__, "name", name)
        if nfs_export_options is not None:
            pulumi.set(__self__, "nfs_export_options", nfs_export_options)

    @property
    @pulumi.getter(name="capacityGb")
    def capacity_gb(self) -> pulumi.Input[float]:
        """
        File share capacity in GiB. This must be at least 1024 GiB
        for the standard tier, or 2560 GiB for the premium tier.
        """
        return pulumi.get(self, "capacity_gb")

    @capacity_gb.setter
    def capacity_gb(self, value: pulumi.Input[float]):
        pulumi.set(self, "capacity_gb", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the fileshare (16 characters or less)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nfsExportOptions")
    def nfs_export_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceFileSharesNfsExportOptionArgs']]]]:
        return pulumi.get(self, "nfs_export_options")

    @nfs_export_options.setter
    def nfs_export_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceFileSharesNfsExportOptionArgs']]]]):
        pulumi.set(self, "nfs_export_options", value)


@pulumi.input_type
class InstanceFileSharesNfsExportOptionArgs:
    def __init__(__self__, *,
                 access_mode: Optional[pulumi.Input[str]] = None,
                 anon_gid: Optional[pulumi.Input[float]] = None,
                 anon_uid: Optional[pulumi.Input[float]] = None,
                 ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 squash_mode: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] access_mode: Either READ_ONLY, for allowing only read requests on the exported directory,
               or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
               Default value is `READ_WRITE`.
               Possible values are `READ_ONLY` and `READ_WRITE`.
        :param pulumi.Input[float] anon_gid: An integer representing the anonymous group id with a default value of 65534.
               Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
               if this field is specified for other squashMode settings.
        :param pulumi.Input[float] anon_uid: An integer representing the anonymous user id with a default value of 65534.
               Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
               if this field is specified for other squashMode settings.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_ranges: List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
               Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
               The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
        :param pulumi.Input[str] squash_mode: Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
               for not allowing root access. The default is NO_ROOT_SQUASH.
               Default value is `NO_ROOT_SQUASH`.
               Possible values are `NO_ROOT_SQUASH` and `ROOT_SQUASH`.
        """
        if access_mode is not None:
            pulumi.set(__self__, "access_mode", access_mode)
        if anon_gid is not None:
            pulumi.set(__self__, "anon_gid", anon_gid)
        if anon_uid is not None:
            pulumi.set(__self__, "anon_uid", anon_uid)
        if ip_ranges is not None:
            pulumi.set(__self__, "ip_ranges", ip_ranges)
        if squash_mode is not None:
            pulumi.set(__self__, "squash_mode", squash_mode)

    @property
    @pulumi.getter(name="accessMode")
    def access_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Either READ_ONLY, for allowing only read requests on the exported directory,
        or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
        Default value is `READ_WRITE`.
        Possible values are `READ_ONLY` and `READ_WRITE`.
        """
        return pulumi.get(self, "access_mode")

    @access_mode.setter
    def access_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_mode", value)

    @property
    @pulumi.getter(name="anonGid")
    def anon_gid(self) -> Optional[pulumi.Input[float]]:
        """
        An integer representing the anonymous group id with a default value of 65534.
        Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
        if this field is specified for other squashMode settings.
        """
        return pulumi.get(self, "anon_gid")

    @anon_gid.setter
    def anon_gid(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "anon_gid", value)

    @property
    @pulumi.getter(name="anonUid")
    def anon_uid(self) -> Optional[pulumi.Input[float]]:
        """
        An integer representing the anonymous user id with a default value of 65534.
        Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
        if this field is specified for other squashMode settings.
        """
        return pulumi.get(self, "anon_uid")

    @anon_uid.setter
    def anon_uid(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "anon_uid", value)

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
        Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
        The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
        """
        return pulumi.get(self, "ip_ranges")

    @ip_ranges.setter
    def ip_ranges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_ranges", value)

    @property
    @pulumi.getter(name="squashMode")
    def squash_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
        for not allowing root access. The default is NO_ROOT_SQUASH.
        Default value is `NO_ROOT_SQUASH`.
        Possible values are `NO_ROOT_SQUASH` and `ROOT_SQUASH`.
        """
        return pulumi.get(self, "squash_mode")

    @squash_mode.setter
    def squash_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "squash_mode", value)


@pulumi.input_type
class InstanceNetworkArgs:
    def __init__(__self__, *,
                 modes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 network: pulumi.Input[str],
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] modes: IP versions for which the instance has
               IP addresses assigned.
               Each value may be one of `ADDRESS_MODE_UNSPECIFIED`, `MODE_IPV4`, and `MODE_IPV6`.
        :param pulumi.Input[str] network: The name of the GCE VPC network to which the
               instance is connected.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: -
               A list of IPv4 or IPv6 addresses.
        :param pulumi.Input[str] reserved_ip_range: A /29 CIDR block that identifies the range of IP
               addresses reserved for this instance.
        """
        pulumi.set(__self__, "modes", modes)
        pulumi.set(__self__, "network", network)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if reserved_ip_range is not None:
            pulumi.set(__self__, "reserved_ip_range", reserved_ip_range)

    @property
    @pulumi.getter
    def modes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        IP versions for which the instance has
        IP addresses assigned.
        Each value may be one of `ADDRESS_MODE_UNSPECIFIED`, `MODE_IPV4`, and `MODE_IPV6`.
        """
        return pulumi.get(self, "modes")

    @modes.setter
    def modes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "modes", value)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input[str]:
        """
        The name of the GCE VPC network to which the
        instance is connected.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input[str]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        -
        A list of IPv4 or IPv6 addresses.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> Optional[pulumi.Input[str]]:
        """
        A /29 CIDR block that identifies the range of IP
        addresses reserved for this instance.
        """
        return pulumi.get(self, "reserved_ip_range")

    @reserved_ip_range.setter
    def reserved_ip_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reserved_ip_range", value)


