# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ResourcePolicyArgs', 'ResourcePolicy']

@pulumi.input_type
class ResourcePolicyArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 group_placement_policy: Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']] = None,
                 instance_schedule_policy: Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 snapshot_schedule_policy: Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']] = None):
        """
        The set of arguments for constructing a ResourcePolicy resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs'] group_placement_policy: Resource policy for instances used for placement configuration.
               Structure is documented below.
        :param pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs'] instance_schedule_policy: Resource policy for scheduling instance operations.
               Structure is documented below.
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
        :param pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs'] snapshot_schedule_policy: Policy for creating snapshots of persistent disks.
               Structure is documented below.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group_placement_policy is not None:
            pulumi.set(__self__, "group_placement_policy", group_placement_policy)
        if instance_schedule_policy is not None:
            pulumi.set(__self__, "instance_schedule_policy", instance_schedule_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if snapshot_schedule_policy is not None:
            pulumi.set(__self__, "snapshot_schedule_policy", snapshot_schedule_policy)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="groupPlacementPolicy")
    def group_placement_policy(self) -> Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']]:
        """
        Resource policy for instances used for placement configuration.
        Structure is documented below.
        """
        return pulumi.get(self, "group_placement_policy")

    @group_placement_policy.setter
    def group_placement_policy(self, value: Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']]):
        pulumi.set(self, "group_placement_policy", value)

    @property
    @pulumi.getter(name="instanceSchedulePolicy")
    def instance_schedule_policy(self) -> Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']]:
        """
        Resource policy for scheduling instance operations.
        Structure is documented below.
        """
        return pulumi.get(self, "instance_schedule_policy")

    @instance_schedule_policy.setter
    def instance_schedule_policy(self, value: Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']]):
        pulumi.set(self, "instance_schedule_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource, provided by the client when initially creating
        the resource. The resource name must be 1-63 characters long, and comply
        with RFC1035. Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z`? which means the
        first character must be a lowercase letter, and all following characters
        must be a dash, lowercase letter, or digit, except the last character,
        which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region where resource policy resides.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="snapshotSchedulePolicy")
    def snapshot_schedule_policy(self) -> Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']]:
        """
        Policy for creating snapshots of persistent disks.
        Structure is documented below.
        """
        return pulumi.get(self, "snapshot_schedule_policy")

    @snapshot_schedule_policy.setter
    def snapshot_schedule_policy(self, value: Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']]):
        pulumi.set(self, "snapshot_schedule_policy", value)


@pulumi.input_type
class _ResourcePolicyState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 group_placement_policy: Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']] = None,
                 instance_schedule_policy: Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 snapshot_schedule_policy: Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']] = None):
        """
        Input properties used for looking up and filtering ResourcePolicy resources.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs'] group_placement_policy: Resource policy for instances used for placement configuration.
               Structure is documented below.
        :param pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs'] instance_schedule_policy: Resource policy for scheduling instance operations.
               Structure is documented below.
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
        :param pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs'] snapshot_schedule_policy: Policy for creating snapshots of persistent disks.
               Structure is documented below.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group_placement_policy is not None:
            pulumi.set(__self__, "group_placement_policy", group_placement_policy)
        if instance_schedule_policy is not None:
            pulumi.set(__self__, "instance_schedule_policy", instance_schedule_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if snapshot_schedule_policy is not None:
            pulumi.set(__self__, "snapshot_schedule_policy", snapshot_schedule_policy)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="groupPlacementPolicy")
    def group_placement_policy(self) -> Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']]:
        """
        Resource policy for instances used for placement configuration.
        Structure is documented below.
        """
        return pulumi.get(self, "group_placement_policy")

    @group_placement_policy.setter
    def group_placement_policy(self, value: Optional[pulumi.Input['ResourcePolicyGroupPlacementPolicyArgs']]):
        pulumi.set(self, "group_placement_policy", value)

    @property
    @pulumi.getter(name="instanceSchedulePolicy")
    def instance_schedule_policy(self) -> Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']]:
        """
        Resource policy for scheduling instance operations.
        Structure is documented below.
        """
        return pulumi.get(self, "instance_schedule_policy")

    @instance_schedule_policy.setter
    def instance_schedule_policy(self, value: Optional[pulumi.Input['ResourcePolicyInstanceSchedulePolicyArgs']]):
        pulumi.set(self, "instance_schedule_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource, provided by the client when initially creating
        the resource. The resource name must be 1-63 characters long, and comply
        with RFC1035. Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z`? which means the
        first character must be a lowercase letter, and all following characters
        must be a dash, lowercase letter, or digit, except the last character,
        which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region where resource policy resides.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="snapshotSchedulePolicy")
    def snapshot_schedule_policy(self) -> Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']]:
        """
        Policy for creating snapshots of persistent disks.
        Structure is documented below.
        """
        return pulumi.get(self, "snapshot_schedule_policy")

    @snapshot_schedule_policy.setter
    def snapshot_schedule_policy(self, value: Optional[pulumi.Input['ResourcePolicySnapshotSchedulePolicyArgs']]):
        pulumi.set(self, "snapshot_schedule_policy", value)


class ResourcePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_placement_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyGroupPlacementPolicyArgs']]] = None,
                 instance_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyInstanceSchedulePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 snapshot_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicySnapshotSchedulePolicyArgs']]] = None,
                 __props__=None):
        """
        A policy that can be attached to a resource to specify or schedule actions on that resource.

        ## Example Usage
        ### Resource Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        foo = gcp.compute.ResourcePolicy("foo",
            region="us-central1",
            snapshot_schedule_policy=gcp.compute.ResourcePolicySnapshotSchedulePolicyArgs(
                schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleArgs(
                    daily_schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs(
                        days_in_cycle=1,
                        start_time="04:00",
                    ),
                ),
            ))
        ```
        ### Resource Policy Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        bar = gcp.compute.ResourcePolicy("bar",
            region="us-central1",
            snapshot_schedule_policy=gcp.compute.ResourcePolicySnapshotSchedulePolicyArgs(
                retention_policy={
                    "maxRetentionDays": 10,
                    "onSourceDiskDelete": "KEEP_AUTO_SNAPSHOTS",
                },
                schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleArgs(
                    hourly_schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleArgs(
                        hours_in_cycle=20,
                        start_time="23:00",
                    ),
                ),
                snapshot_properties=gcp.compute.ResourcePolicySnapshotSchedulePolicySnapshotPropertiesArgs(
                    guest_flush=True,
                    labels={
                        "myLabel": "value",
                    },
                    storage_locations="us",
                ),
            ))
        ```
        ### Resource Policy Placement Policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        baz = gcp.compute.ResourcePolicy("baz",
            group_placement_policy=gcp.compute.ResourcePolicyGroupPlacementPolicyArgs(
                collocation="COLLOCATED",
                vm_count=2,
            ),
            region="us-central1")
        ```
        ### Resource Policy Instance Schedule Policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        hourly = gcp.compute.ResourcePolicy("hourly",
            description="Start and stop instances",
            instance_schedule_policy=gcp.compute.ResourcePolicyInstanceSchedulePolicyArgs(
                time_zone="US/Central",
                vm_start_schedule=gcp.compute.ResourcePolicyInstanceSchedulePolicyVmStartScheduleArgs(
                    schedule="0 * * * *",
                ),
                vm_stop_schedule=gcp.compute.ResourcePolicyInstanceSchedulePolicyVmStopScheduleArgs(
                    schedule="15 * * * *",
                ),
            ),
            region="us-central1")
        ```

        ## Import

        ResourcePolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[pulumi.InputType['ResourcePolicyGroupPlacementPolicyArgs']] group_placement_policy: Resource policy for instances used for placement configuration.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['ResourcePolicyInstanceSchedulePolicyArgs']] instance_schedule_policy: Resource policy for scheduling instance operations.
               Structure is documented below.
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
        :param pulumi.Input[pulumi.InputType['ResourcePolicySnapshotSchedulePolicyArgs']] snapshot_schedule_policy: Policy for creating snapshots of persistent disks.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: Optional[ResourcePolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A policy that can be attached to a resource to specify or schedule actions on that resource.

        ## Example Usage
        ### Resource Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        foo = gcp.compute.ResourcePolicy("foo",
            region="us-central1",
            snapshot_schedule_policy=gcp.compute.ResourcePolicySnapshotSchedulePolicyArgs(
                schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleArgs(
                    daily_schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs(
                        days_in_cycle=1,
                        start_time="04:00",
                    ),
                ),
            ))
        ```
        ### Resource Policy Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        bar = gcp.compute.ResourcePolicy("bar",
            region="us-central1",
            snapshot_schedule_policy=gcp.compute.ResourcePolicySnapshotSchedulePolicyArgs(
                retention_policy={
                    "maxRetentionDays": 10,
                    "onSourceDiskDelete": "KEEP_AUTO_SNAPSHOTS",
                },
                schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleArgs(
                    hourly_schedule=gcp.compute.ResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleArgs(
                        hours_in_cycle=20,
                        start_time="23:00",
                    ),
                ),
                snapshot_properties=gcp.compute.ResourcePolicySnapshotSchedulePolicySnapshotPropertiesArgs(
                    guest_flush=True,
                    labels={
                        "myLabel": "value",
                    },
                    storage_locations="us",
                ),
            ))
        ```
        ### Resource Policy Placement Policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        baz = gcp.compute.ResourcePolicy("baz",
            group_placement_policy=gcp.compute.ResourcePolicyGroupPlacementPolicyArgs(
                collocation="COLLOCATED",
                vm_count=2,
            ),
            region="us-central1")
        ```
        ### Resource Policy Instance Schedule Policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        hourly = gcp.compute.ResourcePolicy("hourly",
            description="Start and stop instances",
            instance_schedule_policy=gcp.compute.ResourcePolicyInstanceSchedulePolicyArgs(
                time_zone="US/Central",
                vm_start_schedule=gcp.compute.ResourcePolicyInstanceSchedulePolicyVmStartScheduleArgs(
                    schedule="0 * * * *",
                ),
                vm_stop_schedule=gcp.compute.ResourcePolicyInstanceSchedulePolicyVmStopScheduleArgs(
                    schedule="15 * * * *",
                ),
            ),
            region="us-central1")
        ```

        ## Import

        ResourcePolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param ResourcePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourcePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_placement_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyGroupPlacementPolicyArgs']]] = None,
                 instance_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyInstanceSchedulePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 snapshot_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicySnapshotSchedulePolicyArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourcePolicyArgs.__new__(ResourcePolicyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["group_placement_policy"] = group_placement_policy
            __props__.__dict__["instance_schedule_policy"] = instance_schedule_policy
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["snapshot_schedule_policy"] = snapshot_schedule_policy
            __props__.__dict__["self_link"] = None
        super(ResourcePolicy, __self__).__init__(
            'gcp:compute/resourcePolicy:ResourcePolicy',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            group_placement_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyGroupPlacementPolicyArgs']]] = None,
            instance_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicyInstanceSchedulePolicyArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            snapshot_schedule_policy: Optional[pulumi.Input[pulumi.InputType['ResourcePolicySnapshotSchedulePolicyArgs']]] = None) -> 'ResourcePolicy':
        """
        Get an existing ResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[pulumi.InputType['ResourcePolicyGroupPlacementPolicyArgs']] group_placement_policy: Resource policy for instances used for placement configuration.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['ResourcePolicyInstanceSchedulePolicyArgs']] instance_schedule_policy: Resource policy for scheduling instance operations.
               Structure is documented below.
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
        :param pulumi.Input[pulumi.InputType['ResourcePolicySnapshotSchedulePolicyArgs']] snapshot_schedule_policy: Policy for creating snapshots of persistent disks.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourcePolicyState.__new__(_ResourcePolicyState)

        __props__.__dict__["description"] = description
        __props__.__dict__["group_placement_policy"] = group_placement_policy
        __props__.__dict__["instance_schedule_policy"] = instance_schedule_policy
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["snapshot_schedule_policy"] = snapshot_schedule_policy
        return ResourcePolicy(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupPlacementPolicy")
    def group_placement_policy(self) -> pulumi.Output[Optional['outputs.ResourcePolicyGroupPlacementPolicy']]:
        """
        Resource policy for instances used for placement configuration.
        Structure is documented below.
        """
        return pulumi.get(self, "group_placement_policy")

    @property
    @pulumi.getter(name="instanceSchedulePolicy")
    def instance_schedule_policy(self) -> pulumi.Output[Optional['outputs.ResourcePolicyInstanceSchedulePolicy']]:
        """
        Resource policy for scheduling instance operations.
        Structure is documented below.
        """
        return pulumi.get(self, "instance_schedule_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource, provided by the client when initially creating
        the resource. The resource name must be 1-63 characters long, and comply
        with RFC1035. Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z`? which means the
        first character must be a lowercase letter, and all following characters
        must be a dash, lowercase letter, or digit, except the last character,
        which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region where resource policy resides.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="snapshotSchedulePolicy")
    def snapshot_schedule_policy(self) -> pulumi.Output[Optional['outputs.ResourcePolicySnapshotSchedulePolicy']]:
        """
        Policy for creating snapshots of persistent disks.
        Structure is documented below.
        """
        return pulumi.get(self, "snapshot_schedule_policy")

