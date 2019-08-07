# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Dataset(pulumi.CustomResource):
    accesses: pulumi.Output[list]
    """
    An array of objects that define dataset access for
    one or more entities. Structure is documented below.
    """
    creation_time: pulumi.Output[float]
    """
    The time when this dataset was created, in milliseconds since the epoch.
    """
    dataset_id: pulumi.Output[str]
    """
    The ID of the dataset containing this table.
    """
    default_partition_expiration_ms: pulumi.Output[float]
    """
    The default partition expiration
    for all partitioned tables in the dataset, in milliseconds.
    """
    default_table_expiration_ms: pulumi.Output[float]
    """
    The default lifetime of all
    tables in the dataset, in milliseconds. The minimum value is 3600000
    milliseconds (one hour).
    """
    delete_contents_on_destroy: pulumi.Output[bool]
    """
    If set to `true`, delete all the
    tables in the dataset when destroying the resource; otherwise, destroying
    the resource will fail if tables are present.
    """
    description: pulumi.Output[str]
    """
    A user-friendly description of the dataset.
    """
    etag: pulumi.Output[str]
    """
    A hash of the resource.
    """
    friendly_name: pulumi.Output[str]
    """
    A descriptive name for the dataset.
    """
    labels: pulumi.Output[dict]
    """
    A mapping of labels to assign to the resource.
    """
    last_modified_time: pulumi.Output[float]
    """
    The date when this dataset or any of its tables was last modified,
    in milliseconds since the epoch.
    """
    location: pulumi.Output[str]
    """
    The geographic location where the dataset should reside.
    See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, resource_name, opts=None, accesses=None, dataset_id=None, default_partition_expiration_ms=None, default_table_expiration_ms=None, delete_contents_on_destroy=None, description=None, friendly_name=None, labels=None, location=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a dataset resource for Google BigQuery. For more information see
        [the official documentation](https://cloud.google.com/bigquery/docs/) and
        [API](https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] accesses: An array of objects that define dataset access for
               one or more entities. Structure is documented below.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this table.
        :param pulumi.Input[float] default_partition_expiration_ms: The default partition expiration
               for all partitioned tables in the dataset, in milliseconds.
        :param pulumi.Input[float] default_table_expiration_ms: The default lifetime of all
               tables in the dataset, in milliseconds. The minimum value is 3600000
               milliseconds (one hour).
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the
               tables in the dataset when destroying the resource; otherwise, destroying
               the resource will fail if tables are present.
        :param pulumi.Input[str] description: A user-friendly description of the dataset.
        :param pulumi.Input[str] friendly_name: A descriptive name for the dataset.
        :param pulumi.Input[dict] labels: A mapping of labels to assign to the resource.
        :param pulumi.Input[str] location: The geographic location where the dataset should reside.
               See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_dataset.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            __props__['accesses'] = accesses
            if dataset_id is None:
                raise TypeError("Missing required property 'dataset_id'")
            __props__['dataset_id'] = dataset_id
            __props__['default_partition_expiration_ms'] = default_partition_expiration_ms
            __props__['default_table_expiration_ms'] = default_table_expiration_ms
            __props__['delete_contents_on_destroy'] = delete_contents_on_destroy
            __props__['description'] = description
            __props__['friendly_name'] = friendly_name
            __props__['labels'] = labels
            __props__['location'] = location
            __props__['project'] = project
            __props__['creation_time'] = None
            __props__['etag'] = None
            __props__['last_modified_time'] = None
            __props__['self_link'] = None
        super(Dataset, __self__).__init__(
            'gcp:bigquery/dataset:Dataset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accesses=None, creation_time=None, dataset_id=None, default_partition_expiration_ms=None, default_table_expiration_ms=None, delete_contents_on_destroy=None, description=None, etag=None, friendly_name=None, labels=None, last_modified_time=None, location=None, project=None, self_link=None):
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] accesses: An array of objects that define dataset access for
               one or more entities. Structure is documented below.
        :param pulumi.Input[float] creation_time: The time when this dataset was created, in milliseconds since the epoch.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this table.
        :param pulumi.Input[float] default_partition_expiration_ms: The default partition expiration
               for all partitioned tables in the dataset, in milliseconds.
        :param pulumi.Input[float] default_table_expiration_ms: The default lifetime of all
               tables in the dataset, in milliseconds. The minimum value is 3600000
               milliseconds (one hour).
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the
               tables in the dataset when destroying the resource; otherwise, destroying
               the resource will fail if tables are present.
        :param pulumi.Input[str] description: A user-friendly description of the dataset.
        :param pulumi.Input[str] etag: A hash of the resource.
        :param pulumi.Input[str] friendly_name: A descriptive name for the dataset.
        :param pulumi.Input[dict] labels: A mapping of labels to assign to the resource.
        :param pulumi.Input[float] last_modified_time: The date when this dataset or any of its tables was last modified,
               in milliseconds since the epoch.
        :param pulumi.Input[str] location: The geographic location where the dataset should reside.
               See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_dataset.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["accesses"] = accesses
        __props__["creation_time"] = creation_time
        __props__["dataset_id"] = dataset_id
        __props__["default_partition_expiration_ms"] = default_partition_expiration_ms
        __props__["default_table_expiration_ms"] = default_table_expiration_ms
        __props__["delete_contents_on_destroy"] = delete_contents_on_destroy
        __props__["description"] = description
        __props__["etag"] = etag
        __props__["friendly_name"] = friendly_name
        __props__["labels"] = labels
        __props__["last_modified_time"] = last_modified_time
        __props__["location"] = location
        __props__["project"] = project
        __props__["self_link"] = self_link
        return Dataset(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

