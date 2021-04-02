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

__all__ = ['OrganizationSinkArgs', 'OrganizationSink']

@pulumi.input_type
class OrganizationSinkArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str],
                 org_id: pulumi.Input[str],
                 bigquery_options: Optional[pulumi.Input['OrganizationSinkBigqueryOptionsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclusions: Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationSinkExclusionArgs']]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 include_children: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationSink resource.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[str] org_id: The numeric ID of the organization to be exported to the sink.
        :param pulumi.Input['OrganizationSinkBigqueryOptionsArgs'] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this exclusion.
        :param pulumi.Input[bool] disabled: If set to True, then this exclusion is disabled and it does not exclude any log entries.
        :param pulumi.Input[Sequence[pulumi.Input['OrganizationSinkExclusionArgs']]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        :param pulumi.Input[str] filter: An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[bool] include_children: Whether or not to include children organizations in the sink export. If true, logs
               associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        :param pulumi.Input[str] name: A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        """
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "org_id", org_id)
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
        if include_children is not None:
            pulumi.set(__self__, "include_children", include_children)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        The destination of the sink (or, in other words, where logs are written to). Can be a
        Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
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
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        The numeric ID of the organization to be exported to the sink.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="bigqueryOptions")
    def bigquery_options(self) -> Optional[pulumi.Input['OrganizationSinkBigqueryOptionsArgs']]:
        """
        Options that affect sinks exporting data to BigQuery. Structure documented below.
        """
        return pulumi.get(self, "bigquery_options")

    @bigquery_options.setter
    def bigquery_options(self, value: Optional[pulumi.Input['OrganizationSinkBigqueryOptionsArgs']]):
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
    def exclusions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationSinkExclusionArgs']]]]:
        """
        Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        """
        return pulumi.get(self, "exclusions")

    @exclusions.setter
    def exclusions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationSinkExclusionArgs']]]]):
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
    @pulumi.getter(name="includeChildren")
    def include_children(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to include children organizations in the sink export. If true, logs
        associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        """
        return pulumi.get(self, "include_children")

    @include_children.setter
    def include_children(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_children", value)

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


@pulumi.input_type
class _OrganizationSinkState:
    def __init__(__self__, *,
                 bigquery_options: Optional[pulumi.Input['OrganizationSinkBigqueryOptionsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclusions: Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationSinkExclusionArgs']]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 include_children: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 writer_identity: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrganizationSink resources.
        :param pulumi.Input['OrganizationSinkBigqueryOptionsArgs'] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this exclusion.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[bool] disabled: If set to True, then this exclusion is disabled and it does not exclude any log entries.
        :param pulumi.Input[Sequence[pulumi.Input['OrganizationSinkExclusionArgs']]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        :param pulumi.Input[str] filter: An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[bool] include_children: Whether or not to include children organizations in the sink export. If true, logs
               associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        :param pulumi.Input[str] name: A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        :param pulumi.Input[str] org_id: The numeric ID of the organization to be exported to the sink.
        :param pulumi.Input[str] writer_identity: The identity associated with this sink. This identity must be granted write access to the
               configured `destination`.
        """
        if bigquery_options is not None:
            pulumi.set(__self__, "bigquery_options", bigquery_options)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if exclusions is not None:
            pulumi.set(__self__, "exclusions", exclusions)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if include_children is not None:
            pulumi.set(__self__, "include_children", include_children)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if writer_identity is not None:
            pulumi.set(__self__, "writer_identity", writer_identity)

    @property
    @pulumi.getter(name="bigqueryOptions")
    def bigquery_options(self) -> Optional[pulumi.Input['OrganizationSinkBigqueryOptionsArgs']]:
        """
        Options that affect sinks exporting data to BigQuery. Structure documented below.
        """
        return pulumi.get(self, "bigquery_options")

    @bigquery_options.setter
    def bigquery_options(self, value: Optional[pulumi.Input['OrganizationSinkBigqueryOptionsArgs']]):
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
    def destination(self) -> Optional[pulumi.Input[str]]:
        """
        The destination of the sink (or, in other words, where logs are written to). Can be a
        Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
        ```python
        import pulumi
        ```
        The writer associated with the sink must have access to write to the above resource.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination", value)

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
    def exclusions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationSinkExclusionArgs']]]]:
        """
        Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        """
        return pulumi.get(self, "exclusions")

    @exclusions.setter
    def exclusions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OrganizationSinkExclusionArgs']]]]):
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
    @pulumi.getter(name="includeChildren")
    def include_children(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to include children organizations in the sink export. If true, logs
        associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        """
        return pulumi.get(self, "include_children")

    @include_children.setter
    def include_children(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_children", value)

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
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The numeric ID of the organization to be exported to the sink.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="writerIdentity")
    def writer_identity(self) -> Optional[pulumi.Input[str]]:
        """
        The identity associated with this sink. This identity must be granted write access to the
        configured `destination`.
        """
        return pulumi.get(self, "writer_identity")

    @writer_identity.setter
    def writer_identity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "writer_identity", value)


class OrganizationSink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bigquery_options: Optional[pulumi.Input[pulumi.InputType['OrganizationSinkBigqueryOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclusions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationSinkExclusionArgs']]]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 include_children: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a organization-level logging sink. For more information see:
        * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/organizations.sinks)
        * How-to Guides
            * [Exporting Logs](https://cloud.google.com/logging/docs/export)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        log_bucket = gcp.storage.Bucket("log-bucket")
        my_sink = gcp.logging.OrganizationSink("my-sink",
            description="some explaination on what this is",
            org_id="123456789",
            destination=log_bucket.name.apply(lambda name: f"storage.googleapis.com/{name}"),
            filter="resource.type = gce_instance AND severity >= WARNING")
        log_writer = gcp.projects.IAMMember("log-writer",
            role="roles/storage.objectCreator",
            member=my_sink.writer_identity)
        ```

        ## Import

        Organization-level logging sinks can be imported using this format

        ```sh
         $ pulumi import gcp:logging/organizationSink:OrganizationSink my_sink organizations/{{organization_id}}/sinks/{{sink_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OrganizationSinkBigqueryOptionsArgs']] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this exclusion.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[bool] disabled: If set to True, then this exclusion is disabled and it does not exclude any log entries.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationSinkExclusionArgs']]]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        :param pulumi.Input[str] filter: An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[bool] include_children: Whether or not to include children organizations in the sink export. If true, logs
               associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        :param pulumi.Input[str] name: A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        :param pulumi.Input[str] org_id: The numeric ID of the organization to be exported to the sink.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationSinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a organization-level logging sink. For more information see:
        * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/organizations.sinks)
        * How-to Guides
            * [Exporting Logs](https://cloud.google.com/logging/docs/export)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        log_bucket = gcp.storage.Bucket("log-bucket")
        my_sink = gcp.logging.OrganizationSink("my-sink",
            description="some explaination on what this is",
            org_id="123456789",
            destination=log_bucket.name.apply(lambda name: f"storage.googleapis.com/{name}"),
            filter="resource.type = gce_instance AND severity >= WARNING")
        log_writer = gcp.projects.IAMMember("log-writer",
            role="roles/storage.objectCreator",
            member=my_sink.writer_identity)
        ```

        ## Import

        Organization-level logging sinks can be imported using this format

        ```sh
         $ pulumi import gcp:logging/organizationSink:OrganizationSink my_sink organizations/{{organization_id}}/sinks/{{sink_id}}
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationSinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationSinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bigquery_options: Optional[pulumi.Input[pulumi.InputType['OrganizationSinkBigqueryOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 exclusions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationSinkExclusionArgs']]]]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 include_children: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = OrganizationSinkArgs.__new__(OrganizationSinkArgs)

            __props__.__dict__['bigquery_options'] = bigquery_options
            __props__.__dict__['description'] = description
            if destination is None and not opts.urn:
                raise TypeError("Missing required property 'destination'")
            __props__.__dict__['destination'] = destination
            __props__.__dict__['disabled'] = disabled
            __props__.__dict__['exclusions'] = exclusions
            __props__.__dict__['filter'] = filter
            __props__.__dict__['include_children'] = include_children
            __props__.__dict__['name'] = name
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__['org_id'] = org_id
            __props__.__dict__['writer_identity'] = None
        super(OrganizationSink, __self__).__init__(
            'gcp:logging/organizationSink:OrganizationSink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bigquery_options: Optional[pulumi.Input[pulumi.InputType['OrganizationSinkBigqueryOptionsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            exclusions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationSinkExclusionArgs']]]]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            include_children: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            writer_identity: Optional[pulumi.Input[str]] = None) -> 'OrganizationSink':
        """
        Get an existing OrganizationSink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OrganizationSinkBigqueryOptionsArgs']] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] description: A description of this exclusion.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
               ```python
               import pulumi
               ```
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[bool] disabled: If set to True, then this exclusion is disabled and it does not exclude any log entries.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrganizationSinkExclusionArgs']]]] exclusions: Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        :param pulumi.Input[str] filter: An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[bool] include_children: Whether or not to include children organizations in the sink export. If true, logs
               associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        :param pulumi.Input[str] name: A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        :param pulumi.Input[str] org_id: The numeric ID of the organization to be exported to the sink.
        :param pulumi.Input[str] writer_identity: The identity associated with this sink. This identity must be granted write access to the
               configured `destination`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationSinkState.__new__(_OrganizationSinkState)

        __props__.__dict__['bigquery_options'] = bigquery_options
        __props__.__dict__['description'] = description
        __props__.__dict__['destination'] = destination
        __props__.__dict__['disabled'] = disabled
        __props__.__dict__['exclusions'] = exclusions
        __props__.__dict__['filter'] = filter
        __props__.__dict__['include_children'] = include_children
        __props__.__dict__['name'] = name
        __props__.__dict__['org_id'] = org_id
        __props__.__dict__['writer_identity'] = writer_identity
        return OrganizationSink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bigqueryOptions")
    def bigquery_options(self) -> pulumi.Output['outputs.OrganizationSinkBigqueryOptions']:
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
        Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
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
    def exclusions(self) -> pulumi.Output[Optional[Sequence['outputs.OrganizationSinkExclusion']]]:
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
    @pulumi.getter(name="includeChildren")
    def include_children(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to include children organizations in the sink export. If true, logs
        associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        """
        return pulumi.get(self, "include_children")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The numeric ID of the organization to be exported to the sink.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="writerIdentity")
    def writer_identity(self) -> pulumi.Output[str]:
        """
        The identity associated with this sink. This identity must be granted write access to the
        configured `destination`.
        """
        return pulumi.get(self, "writer_identity")

