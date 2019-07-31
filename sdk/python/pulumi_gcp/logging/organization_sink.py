# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class OrganizationSink(pulumi.CustomResource):
    destination: pulumi.Output[str]
    """
    The destination of the sink (or, in other words, where logs are written to). Can be a
    Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
    The writer associated with the sink must have access to write to the above resource.
    """
    filter: pulumi.Output[str]
    """
    The filter to apply when exporting logs. Only log entries that match the filter are exported.
    See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
    write a filter.
    """
    include_children: pulumi.Output[bool]
    """
    Whether or not to include children organizations in the sink export. If true, logs
    associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
    """
    name: pulumi.Output[str]
    """
    The name of the logging sink.
    """
    org_id: pulumi.Output[str]
    """
    The numeric ID of the organization to be exported to the sink.
    """
    writer_identity: pulumi.Output[str]
    """
    The identity associated with this sink. This identity must be granted write access to the
    configured `destination`.
    """
    def __init__(__self__, resource_name, opts=None, destination=None, filter=None, include_children=None, name=None, org_id=None, __name__=None, __opts__=None):
        """
        Create a OrganizationSink resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[str] filter: The filter to apply when exporting logs. Only log entries that match the filter are exported.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[bool] include_children: Whether or not to include children organizations in the sink export. If true, logs
               associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        :param pulumi.Input[str] name: The name of the logging sink.
        :param pulumi.Input[str] org_id: The numeric ID of the organization to be exported to the sink.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_organization_sink.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if destination is None:
            raise TypeError("Missing required property 'destination'")
        __props__['destination'] = destination

        __props__['filter'] = filter

        __props__['include_children'] = include_children

        __props__['name'] = name

        if org_id is None:
            raise TypeError("Missing required property 'org_id'")
        __props__['org_id'] = org_id

        __props__['writer_identity'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(OrganizationSink, __self__).__init__(
            'gcp:logging/organizationSink:OrganizationSink',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

