# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Snapshot']


class Snapshot(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str] = pulumi.output_property("creationTimestamp")
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    An optional description of this resource.
    """
    disk_size_gb: pulumi.Output[float] = pulumi.output_property("diskSizeGb")
    """
    Size of the snapshot, specified in GB.
    """
    label_fingerprint: pulumi.Output[str] = pulumi.output_property("labelFingerprint")
    """
    The fingerprint used for optimistic locking of this resource. Used internally during updates.
    """
    labels: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("labels")
    """
    Labels to apply to this Snapshot.
    """
    licenses: pulumi.Output[List[str]] = pulumi.output_property("licenses")
    """
    A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
    attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
    encryption key.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Name of the resource; provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035. Specifically, the name must be 1-63 characters long and match
    the regular expression `a-z?` which means the
    first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last
    character, which cannot be a dash.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str] = pulumi.output_property("selfLink")
    """
    The URI of the created resource.
    """
    snapshot_encryption_key: pulumi.Output[Optional['outputs.SnapshotSnapshotEncryptionKey']] = pulumi.output_property("snapshotEncryptionKey")
    """
    The customer-supplied encryption key of the snapshot. Required if the
    source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
    """
    snapshot_id: pulumi.Output[float] = pulumi.output_property("snapshotId")
    """
    The unique identifier for the resource.
    """
    source_disk: pulumi.Output[str] = pulumi.output_property("sourceDisk")
    """
    A reference to the disk used to create this snapshot.
    """
    source_disk_encryption_key: pulumi.Output[Optional['outputs.SnapshotSourceDiskEncryptionKey']] = pulumi.output_property("sourceDiskEncryptionKey")
    """
    The customer-supplied encryption key of the source snapshot. Required
    if the source snapshot is protected by a customer-supplied encryption
    key.  Structure is documented below.
    """
    source_disk_link: pulumi.Output[str] = pulumi.output_property("sourceDiskLink")
    storage_bytes: pulumi.Output[float] = pulumi.output_property("storageBytes")
    """
    A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot
    creation/deletion.
    """
    zone: pulumi.Output[str] = pulumi.output_property("zone")
    """
    A reference to the zone where the disk is hosted.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, description: Optional[pulumi.Input[str]] = None, labels: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, snapshot_encryption_key: Optional[pulumi.Input[pulumi.InputType['SnapshotSnapshotEncryptionKeyArgs']]] = None, source_disk: Optional[pulumi.Input[str]] = None, source_disk_encryption_key: Optional[pulumi.Input[pulumi.InputType['SnapshotSourceDiskEncryptionKeyArgs']]] = None, zone: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Represents a Persistent Disk Snapshot resource.

        Use snapshots to back up data from your persistent disks. Snapshots are
        different from public images and custom images, which are used primarily
        to create instances or configure instance templates. Snapshots are useful
        for periodic backup of the data on your persistent disks. You can create
        snapshots from persistent disks even while they are attached to running
        instances.

        Snapshots are incremental, so you can create regular snapshots on a
        persistent disk faster and at a much lower cost than if you regularly
        created a full image of the disk.

        To get more information about Snapshot, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)

        > **Warning:** All arguments including `snapshot_encryption_key.raw_key` and `source_disk_encryption_key.raw_key` will be stored in the raw
        state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).

        ## Example Usage
        ### Snapshot Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        debian = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        persistent = gcp.compute.Disk("persistent",
            image=debian.self_link,
            size=10,
            type="pd-ssd",
            zone="us-central1-a")
        snapshot = gcp.compute.Snapshot("snapshot",
            source_disk=persistent.name,
            zone="us-central1-a",
            labels={
                "my_label": "value",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] labels: Labels to apply to this Snapshot.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['SnapshotSnapshotEncryptionKeyArgs']] snapshot_encryption_key: The customer-supplied encryption key of the snapshot. Required if the
               source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
        :param pulumi.Input[str] source_disk: A reference to the disk used to create this snapshot.
        :param pulumi.Input[pulumi.InputType['SnapshotSourceDiskEncryptionKeyArgs']] source_disk_encryption_key: The customer-supplied encryption key of the source snapshot. Required
               if the source snapshot is protected by a customer-supplied encryption
               key.  Structure is documented below.
        :param pulumi.Input[str] zone: A reference to the zone where the disk is hosted.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['project'] = project
            __props__['snapshot_encryption_key'] = snapshot_encryption_key
            if source_disk is None:
                raise TypeError("Missing required property 'source_disk'")
            __props__['source_disk'] = source_disk
            __props__['source_disk_encryption_key'] = source_disk_encryption_key
            __props__['zone'] = zone
            __props__['creation_timestamp'] = None
            __props__['disk_size_gb'] = None
            __props__['label_fingerprint'] = None
            __props__['licenses'] = None
            __props__['self_link'] = None
            __props__['snapshot_id'] = None
            __props__['source_disk_link'] = None
            __props__['storage_bytes'] = None
        super(Snapshot, __self__).__init__(
            'gcp:compute/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, creation_timestamp: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, disk_size_gb: Optional[pulumi.Input[float]] = None, label_fingerprint: Optional[pulumi.Input[str]] = None, labels: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, licenses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, self_link: Optional[pulumi.Input[str]] = None, snapshot_encryption_key: Optional[pulumi.Input[pulumi.InputType['SnapshotSnapshotEncryptionKeyArgs']]] = None, snapshot_id: Optional[pulumi.Input[float]] = None, source_disk: Optional[pulumi.Input[str]] = None, source_disk_encryption_key: Optional[pulumi.Input[pulumi.InputType['SnapshotSourceDiskEncryptionKeyArgs']]] = None, source_disk_link: Optional[pulumi.Input[str]] = None, storage_bytes: Optional[pulumi.Input[float]] = None, zone: Optional[pulumi.Input[str]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[float] disk_size_gb: Size of the snapshot, specified in GB.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] labels: Labels to apply to this Snapshot.
        :param pulumi.Input[List[pulumi.Input[str]]] licenses: A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses
               attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied
               encryption key.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[pulumi.InputType['SnapshotSnapshotEncryptionKeyArgs']] snapshot_encryption_key: The customer-supplied encryption key of the snapshot. Required if the
               source snapshot is protected by a customer-supplied encryption key.  Structure is documented below.
        :param pulumi.Input[float] snapshot_id: The unique identifier for the resource.
        :param pulumi.Input[str] source_disk: A reference to the disk used to create this snapshot.
        :param pulumi.Input[pulumi.InputType['SnapshotSourceDiskEncryptionKeyArgs']] source_disk_encryption_key: The customer-supplied encryption key of the source snapshot. Required
               if the source snapshot is protected by a customer-supplied encryption
               key.  Structure is documented below.
        :param pulumi.Input[float] storage_bytes: A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot
               creation/deletion.
        :param pulumi.Input[str] zone: A reference to the zone where the disk is hosted.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["disk_size_gb"] = disk_size_gb
        __props__["label_fingerprint"] = label_fingerprint
        __props__["labels"] = labels
        __props__["licenses"] = licenses
        __props__["name"] = name
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["snapshot_encryption_key"] = snapshot_encryption_key
        __props__["snapshot_id"] = snapshot_id
        __props__["source_disk"] = source_disk
        __props__["source_disk_encryption_key"] = source_disk_encryption_key
        __props__["source_disk_link"] = source_disk_link
        __props__["storage_bytes"] = storage_bytes
        __props__["zone"] = zone
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

