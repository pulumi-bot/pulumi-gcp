# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ProjectSinkArgs', 'ProjectSink']

@pulumi.input_type
class ProjectSinkArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str],
                 bigquery_options: Optional[pulumi.Input['ProjectSinkBigqueryOptionsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclusions: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectSinkExclusionArgs']]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 unique_writer_identity: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ProjectSink resource.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input['ProjectSinkBigqueryOptionsArgs'] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this exclusion.
        :param pulumi.Input[bool] disabled: If set to True, then this exclusion is disabled and it does not exclude any log entries.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectSinkExclusionArgs']]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        :param pulumi.Input[str] filter: An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        :param pulumi.Input[str] project: The ID of the project to create the sink in. If omitted, the project associated with the provider is
               used.
        :param pulumi.Input[bool] unique_writer_identity: Whether or not to create a unique identity associated with this sink. If `false`
               (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
               then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
               `bigquery_options`, you must set `unique_writer_identity` to true.
        """
        pulumi.set(__self__, "destination", destination)
        if bigquery_options is not None:
            pulumi.set(__self__, "bigquery_options", bigquery_options)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if exclusions is not None:
            pulumi.set(__self__, "exclusions", exclusions)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if unique_writer_identity is not None:
            pulumi.set(__self__, "unique_writer_identity", unique_writer_identity)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        The destination of the sink (or, in other words, where logs are written to). Can be a
        Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
        ```python
        import pulumi
        ```
        The writer associated with the sink must have access to write to the above resource.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter(name="bigqueryOptions")
    def bigquery_options(self) -> Optional[pulumi.Input['ProjectSinkBigqueryOptionsArgs']]:
        """
        Options that affect sinks exporting data to BigQuery. Structure documented below.
        """
        return pulumi.get(self, "bigquery_options")

    @bigquery_options.setter
    def bigquery_options(self, value: Optional[pulumi.Input['ProjectSinkBigqueryOptionsArgs']]):
        pulumi.set(self, "bigquery_options", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of this exclusion.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to True, then this exclusion is disabled and it does not exclude any log entries.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def exclusions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectSinkExclusionArgs']]]]:
        """
        Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        """
        return pulumi.get(self, "exclusions")

    @exclusions.setter
    def exclusions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectSinkExclusionArgs']]]]):
        pulumi.set(self, "exclusions", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
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
        A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project to create the sink in. If omitted, the project associated with the provider is
        used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="uniqueWriterIdentity")
    def unique_writer_identity(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to create a unique identity associated with this sink. If `false`
        (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
        then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
        `bigquery_options`, you must set `unique_writer_identity` to true.
        """
        return pulumi.get(self, "unique_writer_identity")

    @unique_writer_identity.setter
    def unique_writer_identity(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unique_writer_identity", value)


class ProjectSink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bigquery_options: Optional[pulumi.Input[pulumi.InputType['ProjectSinkBigqueryOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclusions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectSinkExclusionArgs']]]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 unique_writer_identity: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Project-level logging sinks can be imported using their URI, e.g.

        ```sh
         $ pulumi import gcp:logging/projectSink:ProjectSink my_sink projects/my-project/sinks/my-sink
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ProjectSinkBigqueryOptionsArgs']] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this exclusion.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[bool] disabled: If set to True, then this exclusion is disabled and it does not exclude any log entries.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectSinkExclusionArgs']]]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        :param pulumi.Input[str] filter: An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        :param pulumi.Input[str] project: The ID of the project to create the sink in. If omitted, the project associated with the provider is
               used.
        :param pulumi.Input[bool] unique_writer_identity: Whether or not to create a unique identity associated with this sink. If `false`
               (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
               then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
               `bigquery_options`, you must set `unique_writer_identity` to true.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectSinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Project-level logging sinks can be imported using their URI, e.g.

        ```sh
         $ pulumi import gcp:logging/projectSink:ProjectSink my_sink projects/my-project/sinks/my-sink
        ```

        :param str resource_name: The name of the resource.
        :param ProjectSinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectSinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bigquery_options: Optional[pulumi.Input[pulumi.InputType['ProjectSinkBigqueryOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclusions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectSinkExclusionArgs']]]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 unique_writer_identity: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['bigquery_options'] = bigquery_options
            __props__['description'] = description
            if destination is None and not opts.urn:
                raise TypeError("Missing required property 'destination'")
            __props__['destination'] = destination
            __props__['disabled'] = disabled
            __props__['exclusions'] = exclusions
            __props__['filter'] = filter
            __props__['name'] = name
            __props__['project'] = project
            __props__['unique_writer_identity'] = unique_writer_identity
            __props__['writer_identity'] = None
        super(ProjectSink, __self__).__init__(
            'gcp:logging/projectSink:ProjectSink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bigquery_options: Optional[pulumi.Input[pulumi.InputType['ProjectSinkBigqueryOptionsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            exclusions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectSinkExclusionArgs']]]]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            unique_writer_identity: Optional[pulumi.Input[bool]] = None,
            writer_identity: Optional[pulumi.Input[str]] = None) -> 'ProjectSink':
        """
        Get an existing ProjectSink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ProjectSinkBigqueryOptionsArgs']] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this exclusion.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[bool] disabled: If set to True, then this exclusion is disabled and it does not exclude any log entries.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProjectSinkExclusionArgs']]]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        :param pulumi.Input[str] filter: An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        :param pulumi.Input[str] project: The ID of the project to create the sink in. If omitted, the project associated with the provider is
               used.
        :param pulumi.Input[bool] unique_writer_identity: Whether or not to create a unique identity associated with this sink. If `false`
               (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
               then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
               `bigquery_options`, you must set `unique_writer_identity` to true.
        :param pulumi.Input[str] writer_identity: The identity associated with this sink. This identity must be granted write access to the
               configured `destination`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bigquery_options"] = bigquery_options
        __props__["description"] = description
        __props__["destination"] = destination
        __props__["disabled"] = disabled
        __props__["exclusions"] = exclusions
        __props__["filter"] = filter
        __props__["name"] = name
        __props__["project"] = project
        __props__["unique_writer_identity"] = unique_writer_identity
        __props__["writer_identity"] = writer_identity
        return ProjectSink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bigqueryOptions")
    def bigquery_options(self) -> pulumi.Output['outputs.ProjectSinkBigqueryOptions']:
        """
        Options that affect sinks exporting data to BigQuery. Structure documented below.
        """
        return pulumi.get(self, "bigquery_options")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of this exclusion.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[str]:
        """
        The destination of the sink (or, in other words, where logs are written to). Can be a
        Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
        ```python
        import pulumi
        ```
        The writer associated with the sink must have access to write to the above resource.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to True, then this exclusion is disabled and it does not exclude any log entries.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def exclusions(self) -> pulumi.Output[Optional[Sequence['outputs.ProjectSinkExclusion']]]:
        """
        Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        """
        return pulumi.get(self, "exclusions")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional[str]]:
        """
        An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project to create the sink in. If omitted, the project associated with the provider is
        used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="uniqueWriterIdentity")
    def unique_writer_identity(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to create a unique identity associated with this sink. If `false`
        (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
        then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
        `bigquery_options`, you must set `unique_writer_identity` to true.
        """
        return pulumi.get(self, "unique_writer_identity")

    @property
    @pulumi.getter(name="writerIdentity")
    def writer_identity(self) -> pulumi.Output[str]:
        """
        The identity associated with this sink. This identity must be granted write access to the
        configured `destination`.
        """
        return pulumi.get(self, "writer_identity")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

