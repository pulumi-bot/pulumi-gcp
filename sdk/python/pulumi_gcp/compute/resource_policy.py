# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ResourcePolicy(pulumi.CustomResource):
    group_placement_policy: pulumi.Output[dict]
    """
    Policy for creating snapshots of persistent disks.  Structure is documented below.

      * `availabilityDomainCount` (`float`) - The number of availability domains instances will be spread across. If two instances are in different
        availability domain, they will not be put in the same low latency network
      * `collocation` (`str`) - Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
        Specify `COLLOCATED` to enable collocation. Can only be specified with `vm_count`. If compute instances are created
        with a COLLOCATED policy, then exactly `vm_count` instances must be created at the same time with the resource policy
        attached.
      * `vmCount` (`float`) - Number of vms in this placement group.
    """
    name: pulumi.Output[str]
    """
    The name of the resource, provided by the client when initially creating
    the resource. The resource name must be 1-63 characters long, and comply
    with RFC1035. Specifically, the name must be 1-63 characters long and
    match the regular expression `a-z`? which means the
    first character must be a lowercase letter, and all following characters
    must be a dash, lowercase letter, or digit, except the last character,
    which cannot be a dash.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    Region where resource policy resides.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    snapshot_schedule_policy: pulumi.Output[dict]
    """
    Policy for creating snapshots of persistent disks.  Structure is documented below.

      * `retention_policy` (`dict`) - Retention policy applied to snapshots created by this resource policy.  Structure is documented below.
        * `maxRetentionDays` (`float`) - Maximum age of the snapshot that is allowed to be kept.
        * `onSourceDiskDelete` (`str`) - Specifies the behavior to apply to scheduled snapshots when
          the source disk is deleted.

      * `schedule` (`dict`) - Contains one of an `hourlySchedule`, `dailySchedule`, or `weeklySchedule`.  Structure is documented below.
        * `dailySchedule` (`dict`) - The policy will execute every nth day at the specified time.  Structure is documented below.
          * `daysInCycle` (`float`) - The number of days between snapshots.
          * `startTime` (`str`) - Time within the window to start the operations.
            It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

        * `hourlySchedule` (`dict`) - The policy will execute every nth hour starting at the specified time.  Structure is documented below.
          * `hoursInCycle` (`float`) - The number of hours between snapshots.
          * `startTime` (`str`) - Time within the window to start the operations.
            It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

        * `weeklySchedule` (`dict`) - Allows specifying a snapshot time for each day of the week.  Structure is documented below.
          * `dayOfWeeks` (`list`) - May contain up to seven (one for each day of the week) snapshot times.  Structure is documented below.
            * `day` (`str`) - The day of the week to create the snapshot. e.g. MONDAY
            * `startTime` (`str`) - Time within the window to start the operations.
              It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

      * `snapshotProperties` (`dict`) - Properties with which the snapshots are created, such as labels.  Structure is documented below.
        * `guestFlush` (`bool`) - Whether to perform a 'guest aware' snapshot.
        * `labels` (`dict`) - A set of key-value pairs.
        * `storageLocations` (`str`) - Cloud Storage bucket location to store the auto snapshot
          (regional or multi-regional)
    """
    def __init__(__self__, resource_name, opts=None, group_placement_policy=None, name=None, project=None, region=None, snapshot_schedule_policy=None, __props__=None, __name__=None, __opts__=None):
        """
        A policy that can be attached to a resource to specify or schedule actions on that resource.



        ## Example Usage - Resource Policy Basic


        ```python
        import pulumi
        import pulumi_gcp as gcp

        foo = gcp.compute.ResourcePolicy("foo",
            region="us-central1",
            snapshot_schedule_policy={
                "schedule": {
                    "dailySchedule": {
                        "daysInCycle": 1,
                        "startTime": "04:00",
                    },
                },
            })
        ```
        ## Example Usage - Resource Policy Full


        ```python
        import pulumi
        import pulumi_gcp as gcp

        bar = gcp.compute.ResourcePolicy("bar",
            region="us-central1",
            snapshot_schedule_policy={
                "retention_policy": {
                    "maxRetentionDays": 10,
                    "onSourceDiskDelete": "KEEP_AUTO_SNAPSHOTS",
                },
                "schedule": {
                    "hourlySchedule": {
                        "hoursInCycle": 20,
                        "startTime": "23:00",
                    },
                },
                "snapshotProperties": {
                    "guestFlush": True,
                    "labels": {
                        "myLabel": "value",
                    },
                    "storageLocations": "us",
                },
            })
        ```
        ## Example Usage - Resource Policy Placement Policy


        ```python
        import pulumi
        import pulumi_gcp as gcp

        baz = gcp.compute.ResourcePolicy("baz",
            group_placement_policy={
                "collocation": "COLLOCATED",
                "vmCount": 2,
            },
            region="us-central1")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] group_placement_policy: Policy for creating snapshots of persistent disks.  Structure is documented below.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating
               the resource. The resource name must be 1-63 characters long, and comply
               with RFC1035. Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z`? which means the
               first character must be a lowercase letter, and all following characters
               must be a dash, lowercase letter, or digit, except the last character,
               which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where resource policy resides.
        :param pulumi.Input[dict] snapshot_schedule_policy: Policy for creating snapshots of persistent disks.  Structure is documented below.

        The **group_placement_policy** object supports the following:

          * `availabilityDomainCount` (`pulumi.Input[float]`) - The number of availability domains instances will be spread across. If two instances are in different
            availability domain, they will not be put in the same low latency network
          * `collocation` (`pulumi.Input[str]`) - Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
            Specify `COLLOCATED` to enable collocation. Can only be specified with `vm_count`. If compute instances are created
            with a COLLOCATED policy, then exactly `vm_count` instances must be created at the same time with the resource policy
            attached.
          * `vmCount` (`pulumi.Input[float]`) - Number of vms in this placement group.

        The **snapshot_schedule_policy** object supports the following:

          * `retention_policy` (`pulumi.Input[dict]`) - Retention policy applied to snapshots created by this resource policy.  Structure is documented below.
            * `maxRetentionDays` (`pulumi.Input[float]`) - Maximum age of the snapshot that is allowed to be kept.
            * `onSourceDiskDelete` (`pulumi.Input[str]`) - Specifies the behavior to apply to scheduled snapshots when
              the source disk is deleted.

          * `schedule` (`pulumi.Input[dict]`) - Contains one of an `hourlySchedule`, `dailySchedule`, or `weeklySchedule`.  Structure is documented below.
            * `dailySchedule` (`pulumi.Input[dict]`) - The policy will execute every nth day at the specified time.  Structure is documented below.
              * `daysInCycle` (`pulumi.Input[float]`) - The number of days between snapshots.
              * `startTime` (`pulumi.Input[str]`) - Time within the window to start the operations.
                It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

            * `hourlySchedule` (`pulumi.Input[dict]`) - The policy will execute every nth hour starting at the specified time.  Structure is documented below.
              * `hoursInCycle` (`pulumi.Input[float]`) - The number of hours between snapshots.
              * `startTime` (`pulumi.Input[str]`) - Time within the window to start the operations.
                It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

            * `weeklySchedule` (`pulumi.Input[dict]`) - Allows specifying a snapshot time for each day of the week.  Structure is documented below.
              * `dayOfWeeks` (`pulumi.Input[list]`) - May contain up to seven (one for each day of the week) snapshot times.  Structure is documented below.
                * `day` (`pulumi.Input[str]`) - The day of the week to create the snapshot. e.g. MONDAY
                * `startTime` (`pulumi.Input[str]`) - Time within the window to start the operations.
                  It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

          * `snapshotProperties` (`pulumi.Input[dict]`) - Properties with which the snapshots are created, such as labels.  Structure is documented below.
            * `guestFlush` (`pulumi.Input[bool]`) - Whether to perform a 'guest aware' snapshot.
            * `labels` (`pulumi.Input[dict]`) - A set of key-value pairs.
            * `storageLocations` (`pulumi.Input[str]`) - Cloud Storage bucket location to store the auto snapshot
              (regional or multi-regional)
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['group_placement_policy'] = group_placement_policy
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
            __props__['snapshot_schedule_policy'] = snapshot_schedule_policy
            __props__['self_link'] = None
        super(ResourcePolicy, __self__).__init__(
            'gcp:compute/resourcePolicy:ResourcePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, group_placement_policy=None, name=None, project=None, region=None, self_link=None, snapshot_schedule_policy=None):
        """
        Get an existing ResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] group_placement_policy: Policy for creating snapshots of persistent disks.  Structure is documented below.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating
               the resource. The resource name must be 1-63 characters long, and comply
               with RFC1035. Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z`? which means the
               first character must be a lowercase letter, and all following characters
               must be a dash, lowercase letter, or digit, except the last character,
               which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where resource policy resides.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[dict] snapshot_schedule_policy: Policy for creating snapshots of persistent disks.  Structure is documented below.

        The **group_placement_policy** object supports the following:

          * `availabilityDomainCount` (`pulumi.Input[float]`) - The number of availability domains instances will be spread across. If two instances are in different
            availability domain, they will not be put in the same low latency network
          * `collocation` (`pulumi.Input[str]`) - Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
            Specify `COLLOCATED` to enable collocation. Can only be specified with `vm_count`. If compute instances are created
            with a COLLOCATED policy, then exactly `vm_count` instances must be created at the same time with the resource policy
            attached.
          * `vmCount` (`pulumi.Input[float]`) - Number of vms in this placement group.

        The **snapshot_schedule_policy** object supports the following:

          * `retention_policy` (`pulumi.Input[dict]`) - Retention policy applied to snapshots created by this resource policy.  Structure is documented below.
            * `maxRetentionDays` (`pulumi.Input[float]`) - Maximum age of the snapshot that is allowed to be kept.
            * `onSourceDiskDelete` (`pulumi.Input[str]`) - Specifies the behavior to apply to scheduled snapshots when
              the source disk is deleted.

          * `schedule` (`pulumi.Input[dict]`) - Contains one of an `hourlySchedule`, `dailySchedule`, or `weeklySchedule`.  Structure is documented below.
            * `dailySchedule` (`pulumi.Input[dict]`) - The policy will execute every nth day at the specified time.  Structure is documented below.
              * `daysInCycle` (`pulumi.Input[float]`) - The number of days between snapshots.
              * `startTime` (`pulumi.Input[str]`) - Time within the window to start the operations.
                It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

            * `hourlySchedule` (`pulumi.Input[dict]`) - The policy will execute every nth hour starting at the specified time.  Structure is documented below.
              * `hoursInCycle` (`pulumi.Input[float]`) - The number of hours between snapshots.
              * `startTime` (`pulumi.Input[str]`) - Time within the window to start the operations.
                It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

            * `weeklySchedule` (`pulumi.Input[dict]`) - Allows specifying a snapshot time for each day of the week.  Structure is documented below.
              * `dayOfWeeks` (`pulumi.Input[list]`) - May contain up to seven (one for each day of the week) snapshot times.  Structure is documented below.
                * `day` (`pulumi.Input[str]`) - The day of the week to create the snapshot. e.g. MONDAY
                * `startTime` (`pulumi.Input[str]`) - Time within the window to start the operations.
                  It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.

          * `snapshotProperties` (`pulumi.Input[dict]`) - Properties with which the snapshots are created, such as labels.  Structure is documented below.
            * `guestFlush` (`pulumi.Input[bool]`) - Whether to perform a 'guest aware' snapshot.
            * `labels` (`pulumi.Input[dict]`) - A set of key-value pairs.
            * `storageLocations` (`pulumi.Input[str]`) - Cloud Storage bucket location to store the auto snapshot
              (regional or multi-regional)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["group_placement_policy"] = group_placement_policy
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["snapshot_schedule_policy"] = snapshot_schedule_policy
        return ResourcePolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
