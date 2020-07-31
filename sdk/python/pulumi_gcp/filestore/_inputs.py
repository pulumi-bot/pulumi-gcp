# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'InstanceFileSharesArgs',
    'InstanceFileSharesNfsExportOptionArgs',
    'InstanceNetworkArgs',
]

@pulumi.input_type
class InstanceFileSharesArgs:
    capacity_gb: pulumi.Input[float] = pulumi.input_property("capacityGb")
    """
    File share capacity in GiB. This must be at least 1024 GiB
    for the standard tier, or 2560 GiB for the premium tier.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    The name of the fileshare (16 characters or less)
    """
    nfs_export_options: Optional[pulumi.Input[List[pulumi.Input['InstanceFileSharesNfsExportOptionArgs']]]] = pulumi.input_property("nfsExportOptions")

    # pylint: disable=no-self-argument
    def __init__(__self__, *, capacity_gb: pulumi.Input[float], name: pulumi.Input[str], nfs_export_options: Optional[pulumi.Input[List[pulumi.Input['InstanceFileSharesNfsExportOptionArgs']]]] = None) -> None:
        """
        :param pulumi.Input[float] capacity_gb: File share capacity in GiB. This must be at least 1024 GiB
               for the standard tier, or 2560 GiB for the premium tier.
        :param pulumi.Input[str] name: The name of the fileshare (16 characters or less)
        """
        __self__.capacity_gb = capacity_gb
        __self__.name = name
        __self__.nfs_export_options = nfs_export_options

@pulumi.input_type
class InstanceFileSharesNfsExportOptionArgs:
    access_mode: Optional[pulumi.Input[str]] = pulumi.input_property("accessMode")
    """
    Either READ_ONLY, for allowing only read requests on the exported directory,
    or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
    """
    anon_gid: Optional[pulumi.Input[float]] = pulumi.input_property("anonGid")
    """
    An integer representing the anonymous group id with a default value of 65534.
    Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
    if this field is specified for other squashMode settings.
    """
    anon_uid: Optional[pulumi.Input[float]] = pulumi.input_property("anonUid")
    """
    An integer representing the anonymous user id with a default value of 65534.
    Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
    if this field is specified for other squashMode settings.
    """
    ip_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("ipRanges")
    """
    List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
    Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
    The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
    """
    squash_mode: Optional[pulumi.Input[str]] = pulumi.input_property("squashMode")
    """
    Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
    for not allowing root access. The default is NO_ROOT_SQUASH.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, access_mode: Optional[pulumi.Input[str]] = None, anon_gid: Optional[pulumi.Input[float]] = None, anon_uid: Optional[pulumi.Input[float]] = None, ip_ranges: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, squash_mode: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] access_mode: Either READ_ONLY, for allowing only read requests on the exported directory,
               or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
        :param pulumi.Input[float] anon_gid: An integer representing the anonymous group id with a default value of 65534.
               Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
               if this field is specified for other squashMode settings.
        :param pulumi.Input[float] anon_uid: An integer representing the anonymous user id with a default value of 65534.
               Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
               if this field is specified for other squashMode settings.
        :param pulumi.Input[List[pulumi.Input[str]]] ip_ranges: List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
               Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
               The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
        :param pulumi.Input[str] squash_mode: Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
               for not allowing root access. The default is NO_ROOT_SQUASH.
        """
        __self__.access_mode = access_mode
        __self__.anon_gid = anon_gid
        __self__.anon_uid = anon_uid
        __self__.ip_ranges = ip_ranges
        __self__.squash_mode = squash_mode

@pulumi.input_type
class InstanceNetworkArgs:
    modes: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("modes")
    """
    IP versions for which the instance has
    IP addresses assigned.
    """
    network: pulumi.Input[str] = pulumi.input_property("network")
    """
    The name of the GCE VPC network to which the
    instance is connected.
    """
    ip_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("ipAddresses")
    """
    -
    A list of IPv4 or IPv6 addresses.
    """
    reserved_ip_range: Optional[pulumi.Input[str]] = pulumi.input_property("reservedIpRange")
    """
    A /29 CIDR block that identifies the range of IP
    addresses reserved for this instance.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, modes: pulumi.Input[List[pulumi.Input[str]]], network: pulumi.Input[str], ip_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, reserved_ip_range: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] modes: IP versions for which the instance has
               IP addresses assigned.
        :param pulumi.Input[str] network: The name of the GCE VPC network to which the
               instance is connected.
        :param pulumi.Input[List[pulumi.Input[str]]] ip_addresses: -
               A list of IPv4 or IPv6 addresses.
        :param pulumi.Input[str] reserved_ip_range: A /29 CIDR block that identifies the range of IP
               addresses reserved for this instance.
        """
        __self__.modes = modes
        __self__.network = network
        __self__.ip_addresses = ip_addresses
        __self__.reserved_ip_range = reserved_ip_range

