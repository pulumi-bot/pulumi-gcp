# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Table(pulumi.CustomResource):
    creation_time: pulumi.Output[float]
    """
    The time when this table was created, in milliseconds since the epoch.
    """
    dataset_id: pulumi.Output[str]
    """
    The dataset ID to create the table in.
    Changing this forces a new resource to be created.
    """
    description: pulumi.Output[str]
    """
    The field description.
    """
    etag: pulumi.Output[str]
    """
    A hash of the resource.
    """
    expiration_time: pulumi.Output[float]
    """
    The time when this table expires, in
    milliseconds since the epoch. If not present, the table will persist
    indefinitely. Expired tables will be deleted and their storage
    reclaimed.
    """
    external_data_configuration: pulumi.Output[dict]
    """
    Describes the data format,
    location, and other properties of a table stored outside of BigQuery.
    By defining these properties, the data source can then be queried as
    if it were a standard BigQuery table. Structure is documented below.
    """
    friendly_name: pulumi.Output[str]
    """
    A descriptive name for the table.
    """
    labels: pulumi.Output[dict]
    """
    A mapping of labels to assign to the resource.
    """
    last_modified_time: pulumi.Output[float]
    """
    The time when this table was last modified, in milliseconds since the epoch.
    """
    location: pulumi.Output[str]
    """
    The geographic location where the table resides. This value is inherited from the dataset.
    """
    num_bytes: pulumi.Output[float]
    """
    The size of this table in bytes, excluding any data in the streaming buffer.
    """
    num_long_term_bytes: pulumi.Output[float]
    """
    The number of bytes in the table that are considered "long-term storage".
    """
    num_rows: pulumi.Output[float]
    """
    The number of rows of data in this table, excluding any data in the streaming buffer.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    schema: pulumi.Output[str]
    """
    A JSON schema for the table. Schema is required
    for CSV and JSON formats and is disallowed for Google Cloud
    Bigtable, Cloud Datastore backups, and Avro formats when using
    external tables. For more information see the
    [BigQuery API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#resource).
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    table_id: pulumi.Output[str]
    """
    A unique ID for the resource.
    Changing this forces a new resource to be created.
    """
    time_partitioning: pulumi.Output[dict]
    """
    If specified, configures time-based
    partitioning for this table. Structure is documented below.
    """
    type: pulumi.Output[str]
    """
    Describes the table type.
    """
    view: pulumi.Output[dict]
    """
    If specified, configures this table as a view.
    Structure is documented below.
    """
    def __init__(__self__, resource_name, opts=None, dataset_id=None, description=None, expiration_time=None, external_data_configuration=None, friendly_name=None, labels=None, project=None, schema=None, table_id=None, time_partitioning=None, view=None, __name__=None, __opts__=None):
        """
        Creates a table resource in a dataset for Google BigQuery. For more information see
        [the official documentation](https://cloud.google.com/bigquery/docs/) and
        [API](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset_id: The dataset ID to create the table in.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: The field description.
        :param pulumi.Input[float] expiration_time: The time when this table expires, in
               milliseconds since the epoch. If not present, the table will persist
               indefinitely. Expired tables will be deleted and their storage
               reclaimed.
        :param pulumi.Input[dict] external_data_configuration: Describes the data format,
               location, and other properties of a table stored outside of BigQuery.
               By defining these properties, the data source can then be queried as
               if it were a standard BigQuery table. Structure is documented below.
        :param pulumi.Input[str] friendly_name: A descriptive name for the table.
        :param pulumi.Input[dict] labels: A mapping of labels to assign to the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] schema: A JSON schema for the table. Schema is required
               for CSV and JSON formats and is disallowed for Google Cloud
               Bigtable, Cloud Datastore backups, and Avro formats when using
               external tables. For more information see the
               [BigQuery API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#resource).
        :param pulumi.Input[str] table_id: A unique ID for the resource.
               Changing this forces a new resource to be created.
        :param pulumi.Input[dict] time_partitioning: If specified, configures time-based
               partitioning for this table. Structure is documented below.
        :param pulumi.Input[dict] view: If specified, configures this table as a view.
               Structure is documented below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_table.html.markdown.
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

        if dataset_id is None:
            raise TypeError("Missing required property 'dataset_id'")
        __props__['dataset_id'] = dataset_id

        __props__['description'] = description

        __props__['expiration_time'] = expiration_time

        __props__['external_data_configuration'] = external_data_configuration

        __props__['friendly_name'] = friendly_name

        __props__['labels'] = labels

        __props__['project'] = project

        __props__['schema'] = schema

        if table_id is None:
            raise TypeError("Missing required property 'table_id'")
        __props__['table_id'] = table_id

        __props__['time_partitioning'] = time_partitioning

        __props__['view'] = view

        __props__['creation_time'] = None
        __props__['etag'] = None
        __props__['last_modified_time'] = None
        __props__['location'] = None
        __props__['num_bytes'] = None
        __props__['num_long_term_bytes'] = None
        __props__['num_rows'] = None
        __props__['self_link'] = None
        __props__['type'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Table, __self__).__init__(
            'gcp:bigquery/table:Table',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

