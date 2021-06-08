# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProjectExclusionArgs', 'ProjectExclusion']

@pulumi.input_type
class ProjectExclusionArgs:
    def __init__(__self__, *,
                 filter: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectExclusion resource.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] name: The name of the logging exclusion.
        :param pulumi.Input[str] project: The project to create the exclusion in. If omitted, the project associated with the provider is
               used.
        """
        pulumi.set(__self__, "filter", filter)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Input[str]:
        """
        The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: pulumi.Input[str]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this exclusion rule should be disabled or not. This defaults to
        false.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the logging exclusion.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project to create the exclusion in. If omitted, the project associated with the provider is
        used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _ProjectExclusionState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectExclusion resources.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging exclusion.
        :param pulumi.Input[str] project: The project to create the exclusion in. If omitted, the project associated with the provider is
               used.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this exclusion rule should be disabled or not. This defaults to
        false.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the logging exclusion.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project to create the exclusion in. If omitted, the project associated with the provider is
        used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class ProjectExclusion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a project-level logging exclusion. For more information see:

        * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.exclusions)
        * How-to Guides
            * [Excluding Logs](https://cloud.google.com/logging/docs/exclusions)

        > You can specify exclusions for log sinks created by the provider by using the exclusions field of `logging.ProjectSink`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_exclusion = gcp.logging.ProjectExclusion("my-exclusion",
            description="Exclude GCE instance debug logs",
            filter="resource.type = gce_instance AND severity <= DEBUG")
        ```

        ## Import

        Project-level logging exclusions can be imported using their URI, e.g.

        ```sh
         $ pulumi import gcp:logging/projectExclusion:ProjectExclusion my_exclusion projects/my-project/exclusions/my-exclusion
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging exclusion.
        :param pulumi.Input[str] project: The project to create the exclusion in. If omitted, the project associated with the provider is
               used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: ProjectExclusionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a project-level logging exclusion. For more information see:

        * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.exclusions)
        * How-to Guides
            * [Excluding Logs](https://cloud.google.com/logging/docs/exclusions)

        > You can specify exclusions for log sinks created by the provider by using the exclusions field of `logging.ProjectSink`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_exclusion = gcp.logging.ProjectExclusion("my-exclusion",
            description="Exclude GCE instance debug logs",
            filter="resource.type = gce_instance AND severity <= DEBUG")
        ```

        ## Import

        Project-level logging exclusions can be imported using their URI, e.g.

        ```sh
         $ pulumi import gcp:logging/projectExclusion:ProjectExclusion my_exclusion projects/my-project/exclusions/my-exclusion
        ```

        :param str resource_name_: The name of the resource.
        :param ProjectExclusionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectExclusionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            __props__ = ProjectExclusionArgs.__new__(ProjectExclusionArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["disabled"] = disabled
            if filter is None and not opts.urn:
                raise TypeError("Missing required property 'filter'")
            __props__.__dict__["filter"] = filter
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
        super(ProjectExclusion, __self__).__init__(
            'gcp:logging/projectExclusion:ProjectExclusion',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'ProjectExclusion':
        """
        Get an existing ProjectExclusion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging exclusion.
        :param pulumi.Input[str] project: The project to create the exclusion in. If omitted, the project associated with the provider is
               used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectExclusionState.__new__(_ProjectExclusionState)

        __props__.__dict__["description"] = description
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["filter"] = filter
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        return ProjectExclusion(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A human-readable description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether this exclusion rule should be disabled or not. This defaults to
        false.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[str]:
        """
        The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the logging exclusion.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project to create the exclusion in. If omitted, the project associated with the provider is
        used.
        """
        return pulumi.get(self, "project")

