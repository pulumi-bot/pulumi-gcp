# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
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
    def __init__(__self__, __name__, __opts__=None, destination=None, filter=None, include_children=None, name=None, org_id=None):
        """
        Manages a organization-level logging sink. For more information see
        [the official documentation](https://cloud.google.com/logging/docs/) and
        [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
        
        Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
        granted to the credentials used with terraform.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
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
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if destination is None:
            raise TypeError('Missing required property destination')
        __props__['destination'] = destination

        __props__['filter'] = filter

        __props__['include_children'] = include_children

        __props__['name'] = name

        if org_id is None:
            raise TypeError('Missing required property org_id')
        __props__['org_id'] = org_id

        __props__['writer_identity'] = None

        super(OrganizationSink, __self__).__init__(
            'gcp:logging/organizationSink:OrganizationSink',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

