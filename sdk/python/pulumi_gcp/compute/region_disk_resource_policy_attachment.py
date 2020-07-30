# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class RegionDiskResourcePolicyAttachment(pulumi.CustomResource):
    disk: pulumi.Output[str]
    """
    The name of the regional disk in which the resource policies are attached to.
    """
    name: pulumi.Output[str]
    """
    The resource policy to be attached to the disk for scheduling snapshot
    creation. Do not specify the self link.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    A reference to the region where the disk resides.
    """
    def __init__(__self__, resource_name, opts=None, disk=None, name=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Adds existing resource policies to a disk. You can only add one policy
        which will be applied to this disk for scheduling snapshot creation.

        > **Note:** This resource does not support zonal disks (`compute.Disk`). For zonal disks, please refer to the `compute.DiskResourcePolicyAttachment` resource.

        ## Example Usage
        ### Region Disk Resource Policy Attachment Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        disk = gcp.compute.Disk("disk",
            image="debian-cloud/debian-9",
            size=50,
            type="pd-ssd",
            zone="us-central1-a")
        snapdisk = gcp.compute.Snapshot("snapdisk",
            source_disk=disk.name,
            zone="us-central1-a")
        ssd = gcp.compute.RegionDisk("ssd",
            replica_zones=[
                "us-central1-a",
                "us-central1-f",
            ],
            snapshot=snapdisk.id,
            size=50,
            type="pd-ssd",
            region="us-central1")
        attachment = gcp.compute.RegionDiskResourcePolicyAttachment("attachment",
            disk=ssd.name,
            region="us-central1")
        policy = gcp.compute.ResourcePolicy("policy",
            region="us-central1",
            snapshot_schedule_policy=gcp.compute.ResourcePolicySnapshotSchedulePolicyArgs(
                schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleArgs(
                    daily_schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs(
                        days_in_cycle=1,
                        start_time="04:00",
                    ),
                ),
            ))
        my_image = gcp.compute.get_image(gcp.compute.GetImageArgsArgs(
            family="debian-9",
            project="debian-cloud",
        ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk: The name of the regional disk in which the resource policies are attached to.
        :param pulumi.Input[str] name: The resource policy to be attached to the disk for scheduling snapshot
               creation. Do not specify the self link.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: A reference to the region where the disk resides.
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

            if disk is None:
                raise TypeError("Missing required property 'disk'")
            __props__['disk'] = disk
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
        super(RegionDiskResourcePolicyAttachment, __self__).__init__(
            'gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, disk=None, name=None, project=None, region=None):
        """
        Get an existing RegionDiskResourcePolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk: The name of the regional disk in which the resource policies are attached to.
        :param pulumi.Input[str] name: The resource policy to be attached to the disk for scheduling snapshot
               creation. Do not specify the self link.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: A reference to the region where the disk resides.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["disk"] = disk
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        return RegionDiskResourcePolicyAttachment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
