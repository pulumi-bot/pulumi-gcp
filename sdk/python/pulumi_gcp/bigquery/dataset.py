# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Dataset(pulumi.CustomResource):
    accesses: pulumi.Output[list]
    """
    An array of objects that define dataset access for one or more entities.  Structure is documented below.

      * `domain` (`str`) - A domain to grant access to. Any users signed in with the
        domain specified will be granted the specified access
      * `group_by_email` (`str`) - An email address of a Google Group to grant access to.
      * `role` (`str`) - Describes the rights granted to the user specified by the other
        member of the access object. Primitive, Predefined and custom
        roles are supported. Predefined roles that have equivalent
        primitive roles are swapped by the API to their Primitive
        counterparts. See
        [official docs](https://cloud.google.com/bigquery/docs/access-control).
      * `special_group` (`str`) - A special group to grant access to. Possible values include:
      * `user_by_email` (`str`) - An email address of a user to grant access to. For example:
        fred@example.com
      * `view` (`dict`) - A view from a different dataset to grant access to. Queries
        executed against that view will have read access to tables in
        this dataset. The role field is not required when this field is
        set. If that view is updated by any user, access to the view
        needs to be granted again via an update operation.  Structure is documented below.
        * `dataset_id` (`str`) - The ID of the dataset containing this table.
        * `project_id` (`str`) - The ID of the project containing this table.
        * `table_id` (`str`) - The ID of the table. The ID must contain only letters (a-z,
          A-Z), numbers (0-9), or underscores (_). The maximum length
          is 1,024 characters.
    """
    creation_time: pulumi.Output[float]
    """
    The time when this dataset was created, in milliseconds since the epoch.
    """
    dataset_id: pulumi.Output[str]
    """
    The ID of the dataset containing this table.
    """
    default_encryption_configuration: pulumi.Output[dict]
    """
    The default encryption key for all tables in the dataset. Once this property is set,
    all newly-created partitioned tables in the dataset will have encryption key set to
    this value, unless table creation request (or query) overrides the key.  Structure is documented below.

      * `kms_key_name` (`str`) - Describes the Cloud KMS encryption key that will be used to protect destination
        BigQuery table. The BigQuery Service Account associated with your project requires
        access to this encryption key.
    """
    default_partition_expiration_ms: pulumi.Output[float]
    """
    The default partition expiration for all partitioned tables in
    the dataset, in milliseconds.
    """
    default_table_expiration_ms: pulumi.Output[float]
    """
    The default lifetime of all tables in the dataset, in milliseconds.
    The minimum value is 3600000 milliseconds (one hour).
    """
    delete_contents_on_destroy: pulumi.Output[bool]
    """
    If set to `true`, delete all the tables in the
    dataset when destroying the resource; otherwise,
    destroying the resource will fail if tables are present.
    """
    description: pulumi.Output[str]
    """
    A user-friendly description of the dataset
    """
    etag: pulumi.Output[str]
    """
    A hash of the resource.
    """
    friendly_name: pulumi.Output[str]
    """
    A descriptive name for the dataset
    """
    labels: pulumi.Output[dict]
    """
    The labels associated with this dataset. You can use these to
    organize and group your datasets
    """
    last_modified_time: pulumi.Output[float]
    """
    The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
    """
    location: pulumi.Output[str]
    """
    The geographic location where the dataset should reside.
    See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, resource_name, opts=None, accesses=None, dataset_id=None, default_encryption_configuration=None, default_partition_expiration_ms=None, default_table_expiration_ms=None, delete_contents_on_destroy=None, description=None, friendly_name=None, labels=None, location=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Datasets allow you to organize and control access to your tables.


        To get more information about Dataset, see:

        * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets)
        * How-to Guides
            * [Datasets Intro](https://cloud.google.com/bigquery/docs/datasets-intro)

        ## Example Usage - Bigquery Dataset Basic


        ```python
        import pulumi
        import pulumi_gcp as gcp

        bqowner = gcp.service_account.Account("bqowner", account_id="bqowner")
        dataset = gcp.bigquery.Dataset("dataset",
            dataset_id="example_dataset",
            friendly_name="test",
            description="This is a test description",
            location="EU",
            default_table_expiration_ms=3600000,
            labels={
                "env": "default",
            },
            access=[
                {
                    "role": "OWNER",
                    "user_by_email": bqowner.email,
                },
                {
                    "role": "READER",
                    "domain": "hashicorp.com",
                },
            ])
        ```
        ## Example Usage - Bigquery Dataset Cmek


        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRing("keyRing", location="us")
        crypto_key = gcp.kms.CryptoKey("cryptoKey", key_ring=key_ring.id)
        dataset = gcp.bigquery.Dataset("dataset",
            dataset_id="example_dataset",
            friendly_name="test",
            description="This is a test description",
            location="US",
            default_table_expiration_ms=3600000,
            default_encryption_configuration={
                "kms_key_name": crypto_key.id,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] accesses: An array of objects that define dataset access for one or more entities.  Structure is documented below.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this table.
        :param pulumi.Input[dict] default_encryption_configuration: The default encryption key for all tables in the dataset. Once this property is set,
               all newly-created partitioned tables in the dataset will have encryption key set to
               this value, unless table creation request (or query) overrides the key.  Structure is documented below.
        :param pulumi.Input[float] default_partition_expiration_ms: The default partition expiration for all partitioned tables in
               the dataset, in milliseconds.
        :param pulumi.Input[float] default_table_expiration_ms: The default lifetime of all tables in the dataset, in milliseconds.
               The minimum value is 3600000 milliseconds (one hour).
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the tables in the
               dataset when destroying the resource; otherwise,
               destroying the resource will fail if tables are present.
        :param pulumi.Input[str] description: A user-friendly description of the dataset
        :param pulumi.Input[str] friendly_name: A descriptive name for the dataset
        :param pulumi.Input[dict] labels: The labels associated with this dataset. You can use these to
               organize and group your datasets
        :param pulumi.Input[str] location: The geographic location where the dataset should reside.
               See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **accesses** object supports the following:

          * `domain` (`pulumi.Input[str]`) - A domain to grant access to. Any users signed in with the
            domain specified will be granted the specified access
          * `group_by_email` (`pulumi.Input[str]`) - An email address of a Google Group to grant access to.
          * `role` (`pulumi.Input[str]`) - Describes the rights granted to the user specified by the other
            member of the access object. Primitive, Predefined and custom
            roles are supported. Predefined roles that have equivalent
            primitive roles are swapped by the API to their Primitive
            counterparts. See
            [official docs](https://cloud.google.com/bigquery/docs/access-control).
          * `special_group` (`pulumi.Input[str]`) - A special group to grant access to. Possible values include:
          * `user_by_email` (`pulumi.Input[str]`) - An email address of a user to grant access to. For example:
            fred@example.com
          * `view` (`pulumi.Input[dict]`) - A view from a different dataset to grant access to. Queries
            executed against that view will have read access to tables in
            this dataset. The role field is not required when this field is
            set. If that view is updated by any user, access to the view
            needs to be granted again via an update operation.  Structure is documented below.
            * `dataset_id` (`pulumi.Input[str]`) - The ID of the dataset containing this table.
            * `project_id` (`pulumi.Input[str]`) - The ID of the project containing this table.
            * `table_id` (`pulumi.Input[str]`) - The ID of the table. The ID must contain only letters (a-z,
              A-Z), numbers (0-9), or underscores (_). The maximum length
              is 1,024 characters.

        The **default_encryption_configuration** object supports the following:

          * `kms_key_name` (`pulumi.Input[str]`) - Describes the Cloud KMS encryption key that will be used to protect destination
            BigQuery table. The BigQuery Service Account associated with your project requires
            access to this encryption key.
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

            __props__['accesses'] = accesses
            if dataset_id is None:
                raise TypeError("Missing required property 'dataset_id'")
            __props__['dataset_id'] = dataset_id
            __props__['default_encryption_configuration'] = default_encryption_configuration
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
    def get(resource_name, id, opts=None, accesses=None, creation_time=None, dataset_id=None, default_encryption_configuration=None, default_partition_expiration_ms=None, default_table_expiration_ms=None, delete_contents_on_destroy=None, description=None, etag=None, friendly_name=None, labels=None, last_modified_time=None, location=None, project=None, self_link=None):
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] accesses: An array of objects that define dataset access for one or more entities.  Structure is documented below.
        :param pulumi.Input[float] creation_time: The time when this dataset was created, in milliseconds since the epoch.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this table.
        :param pulumi.Input[dict] default_encryption_configuration: The default encryption key for all tables in the dataset. Once this property is set,
               all newly-created partitioned tables in the dataset will have encryption key set to
               this value, unless table creation request (or query) overrides the key.  Structure is documented below.
        :param pulumi.Input[float] default_partition_expiration_ms: The default partition expiration for all partitioned tables in
               the dataset, in milliseconds.
        :param pulumi.Input[float] default_table_expiration_ms: The default lifetime of all tables in the dataset, in milliseconds.
               The minimum value is 3600000 milliseconds (one hour).
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the tables in the
               dataset when destroying the resource; otherwise,
               destroying the resource will fail if tables are present.
        :param pulumi.Input[str] description: A user-friendly description of the dataset
        :param pulumi.Input[str] etag: A hash of the resource.
        :param pulumi.Input[str] friendly_name: A descriptive name for the dataset
        :param pulumi.Input[dict] labels: The labels associated with this dataset. You can use these to
               organize and group your datasets
        :param pulumi.Input[float] last_modified_time: The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
        :param pulumi.Input[str] location: The geographic location where the dataset should reside.
               See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.

        The **accesses** object supports the following:

          * `domain` (`pulumi.Input[str]`) - A domain to grant access to. Any users signed in with the
            domain specified will be granted the specified access
          * `group_by_email` (`pulumi.Input[str]`) - An email address of a Google Group to grant access to.
          * `role` (`pulumi.Input[str]`) - Describes the rights granted to the user specified by the other
            member of the access object. Primitive, Predefined and custom
            roles are supported. Predefined roles that have equivalent
            primitive roles are swapped by the API to their Primitive
            counterparts. See
            [official docs](https://cloud.google.com/bigquery/docs/access-control).
          * `special_group` (`pulumi.Input[str]`) - A special group to grant access to. Possible values include:
          * `user_by_email` (`pulumi.Input[str]`) - An email address of a user to grant access to. For example:
            fred@example.com
          * `view` (`pulumi.Input[dict]`) - A view from a different dataset to grant access to. Queries
            executed against that view will have read access to tables in
            this dataset. The role field is not required when this field is
            set. If that view is updated by any user, access to the view
            needs to be granted again via an update operation.  Structure is documented below.
            * `dataset_id` (`pulumi.Input[str]`) - The ID of the dataset containing this table.
            * `project_id` (`pulumi.Input[str]`) - The ID of the project containing this table.
            * `table_id` (`pulumi.Input[str]`) - The ID of the table. The ID must contain only letters (a-z,
              A-Z), numbers (0-9), or underscores (_). The maximum length
              is 1,024 characters.

        The **default_encryption_configuration** object supports the following:

          * `kms_key_name` (`pulumi.Input[str]`) - Describes the Cloud KMS encryption key that will be used to protect destination
            BigQuery table. The BigQuery Service Account associated with your project requires
            access to this encryption key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accesses"] = accesses
        __props__["creation_time"] = creation_time
        __props__["dataset_id"] = dataset_id
        __props__["default_encryption_configuration"] = default_encryption_configuration
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
