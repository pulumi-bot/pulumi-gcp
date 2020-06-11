# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class TransferJob(pulumi.CustomResource):
    creation_time: pulumi.Output[str]
    """
    When the Transfer Job was created.
    """
    deletion_time: pulumi.Output[str]
    """
    When the Transfer Job was deleted.
    """
    description: pulumi.Output[str]
    """
    Unique description to identify the Transfer Job.
    """
    last_modification_time: pulumi.Output[str]
    """
    When the Transfer Job was last modified.
    """
    name: pulumi.Output[str]
    """
    The name of the Transfer Job.
    """
    project: pulumi.Output[str]
    """
    The project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    schedule: pulumi.Output[dict]
    """
    Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.

      * `scheduleEndDate` (`dict`) - The last day the recurring transfer will be run. If `schedule_end_date` is the same as `schedule_start_date`, the transfer will be executed only once. Structure documented below.
        * `day` (`float`) - Day of month. Must be from 1 to 31 and valid for the year and month.
        * `month` (`float`) - Month of year. Must be from 1 to 12.
        * `year` (`float`) - Year of date. Must be from 1 to 9999.

      * `scheduleStartDate` (`dict`) - The first day the recurring transfer is scheduled to run. If `schedule_start_date` is in the past, the transfer will run for the first time on the following day. Structure documented below.
        * `day` (`float`) - Day of month. Must be from 1 to 31 and valid for the year and month.
        * `month` (`float`) - Month of year. Must be from 1 to 12.
        * `year` (`float`) - Year of date. Must be from 1 to 9999.

      * `startTimeOfDay` (`dict`) - The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer's start time in a day is specified in your local timezone. Structure documented below.
        * `hours` (`float`) - Hours of day in 24 hour format. Should be from 0 to 23
        * `minutes` (`float`) - Minutes of hour of day. Must be from 0 to 59.
        * `nanos` (`float`) - Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
        * `seconds` (`float`) - Seconds of minutes of the time. Must normally be from 0 to 59.
    """
    status: pulumi.Output[str]
    """
    Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
    """
    transfer_spec: pulumi.Output[dict]
    """
    Transfer specification. Structure documented below.

      * `awsS3DataSource` (`dict`) - An AWS S3 data source. Structure documented below.
        * `awsAccessKey` (`dict`) - AWS credentials block.
          * `accessKeyId` (`str`) - AWS Key ID.
          * `secretAccessKey` (`str`) - AWS Secret Access Key.

        * `bucket_name` (`str`) - S3 Bucket name.

      * `gcsDataSink` (`dict`) - A Google Cloud Storage data sink. Structure documented below.
        * `bucket_name` (`str`) - S3 Bucket name.

      * `gcsDataSource` (`dict`) - A Google Cloud Storage data source. Structure documented below.
        * `bucket_name` (`str`) - S3 Bucket name.

      * `httpDataSource` (`dict`) - An HTTP URL data source. Structure documented below.
        * `listUrl` (`str`) - The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.

      * `objectConditions` (`dict`) - Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' `last_modification_time` do not exclude objects in a data sink. Structure documented below.
        * `excludePrefixes` (`list`) - `exclude_prefixes` must follow the requirements described for `include_prefixes`. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).
        * `includePrefixes` (`list`) - If `include_refixes` is specified, objects that satisfy the object conditions must have names that start with one of the `include_prefixes` and that do not start with any of the `exclude_prefixes`. If `include_prefixes` is not specified, all objects except those that have names starting with one of the `exclude_prefixes` must satisfy the object conditions. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).
        * `maxTimeElapsedSinceLastModification` (`str`) - A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        * `minTimeElapsedSinceLastModification` (`str`) - 
          A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".

      * `transferOptions` (`dict`) - Characteristics of how to treat files from datasource and sink during job. If the option `delete_objects_unique_in_sink` is true, object conditions based on objects' `last_modification_time` are ignored and do not exclude objects in a data source or a data sink. Structure documented below.
        * `deleteObjectsFromSourceAfterTransfer` (`bool`) - Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and `delete_objects_unique_in_sink` are mutually exclusive.
        * `deleteObjectsUniqueInSink` (`bool`) - Whether objects that exist only in the sink should be deleted. Note that this option and
          `delete_objects_from_source_after_transfer` are mutually exclusive.
        * `overwriteObjectsAlreadyExistingInSink` (`bool`) - Whether overwriting objects that already exist in the sink is allowed.
    """
    def __init__(__self__, resource_name, opts=None, description=None, project=None, schedule=None, status=None, transfer_spec=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a new Transfer Job in Google Cloud Storage Transfer.

        To get more information about Google Cloud Storage Transfer, see:

        * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
        * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs#TransferJob)
        * How-to Guides
            * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)

        ## Example Usage



        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.storage.get_transfer_project_servie_account(project=var["project"])
        s3_backup_bucket_bucket = gcp.storage.Bucket("s3-backup-bucketBucket",
            storage_class="NEARLINE",
            project=var["project"])
        s3_backup_bucket_bucket_iam_member = gcp.storage.BucketIAMMember("s3-backup-bucketBucketIAMMember",
            bucket=s3_backup_bucket_bucket.name,
            role="roles/storage.admin",
            member=f"serviceAccount:{default.email}")
        s3_bucket_nightly_backup = gcp.storage.TransferJob("s3-bucket-nightly-backup",
            description="Nightly backup of S3 bucket",
            project=var["project"],
            transfer_spec={
                "object_conditions": {
                    "maxTimeElapsedSinceLastModification": "600s",
                    "excludePrefixes": ["requests.gz"],
                },
                "transfer_options": {
                    "deleteObjectsUniqueInSink": False,
                },
                "aws_s3_data_source": {
                    "bucket_name": var["aws_s3_bucket"],
                    "aws_access_key": {
                        "accessKeyId": var["aws_access_key"],
                        "secretAccessKey": var["aws_secret_key"],
                    },
                },
                "gcs_data_sink": {
                    "bucket_name": s3_backup_bucket_bucket.name,
                },
            },
            schedule={
                "schedule_start_date": {
                    "year": 2018,
                    "month": 10,
                    "day": 1,
                },
                "schedule_end_date": {
                    "year": 2019,
                    "month": 1,
                    "day": 15,
                },
                "start_time_of_day": {
                    "hours": 23,
                    "minutes": 30,
                    "seconds": 0,
                    "nanos": 0,
                },
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Unique description to identify the Transfer Job.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[dict] schedule: Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
        :param pulumi.Input[str] status: Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        :param pulumi.Input[dict] transfer_spec: Transfer specification. Structure documented below.

        The **schedule** object supports the following:

          * `scheduleEndDate` (`pulumi.Input[dict]`) - The last day the recurring transfer will be run. If `schedule_end_date` is the same as `schedule_start_date`, the transfer will be executed only once. Structure documented below.
            * `day` (`pulumi.Input[float]`) - Day of month. Must be from 1 to 31 and valid for the year and month.
            * `month` (`pulumi.Input[float]`) - Month of year. Must be from 1 to 12.
            * `year` (`pulumi.Input[float]`) - Year of date. Must be from 1 to 9999.

          * `scheduleStartDate` (`pulumi.Input[dict]`) - The first day the recurring transfer is scheduled to run. If `schedule_start_date` is in the past, the transfer will run for the first time on the following day. Structure documented below.
            * `day` (`pulumi.Input[float]`) - Day of month. Must be from 1 to 31 and valid for the year and month.
            * `month` (`pulumi.Input[float]`) - Month of year. Must be from 1 to 12.
            * `year` (`pulumi.Input[float]`) - Year of date. Must be from 1 to 9999.

          * `startTimeOfDay` (`pulumi.Input[dict]`) - The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer's start time in a day is specified in your local timezone. Structure documented below.
            * `hours` (`pulumi.Input[float]`) - Hours of day in 24 hour format. Should be from 0 to 23
            * `minutes` (`pulumi.Input[float]`) - Minutes of hour of day. Must be from 0 to 59.
            * `nanos` (`pulumi.Input[float]`) - Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
            * `seconds` (`pulumi.Input[float]`) - Seconds of minutes of the time. Must normally be from 0 to 59.

        The **transfer_spec** object supports the following:

          * `awsS3DataSource` (`pulumi.Input[dict]`) - An AWS S3 data source. Structure documented below.
            * `awsAccessKey` (`pulumi.Input[dict]`) - AWS credentials block.
              * `accessKeyId` (`pulumi.Input[str]`) - AWS Key ID.
              * `secretAccessKey` (`pulumi.Input[str]`) - AWS Secret Access Key.

            * `bucket_name` (`pulumi.Input[str]`) - S3 Bucket name.

          * `gcsDataSink` (`pulumi.Input[dict]`) - A Google Cloud Storage data sink. Structure documented below.
            * `bucket_name` (`pulumi.Input[str]`) - S3 Bucket name.

          * `gcsDataSource` (`pulumi.Input[dict]`) - A Google Cloud Storage data source. Structure documented below.
            * `bucket_name` (`pulumi.Input[str]`) - S3 Bucket name.

          * `httpDataSource` (`pulumi.Input[dict]`) - An HTTP URL data source. Structure documented below.
            * `listUrl` (`pulumi.Input[str]`) - The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.

          * `objectConditions` (`pulumi.Input[dict]`) - Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' `last_modification_time` do not exclude objects in a data sink. Structure documented below.
            * `excludePrefixes` (`pulumi.Input[list]`) - `exclude_prefixes` must follow the requirements described for `include_prefixes`. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).
            * `includePrefixes` (`pulumi.Input[list]`) - If `include_refixes` is specified, objects that satisfy the object conditions must have names that start with one of the `include_prefixes` and that do not start with any of the `exclude_prefixes`. If `include_prefixes` is not specified, all objects except those that have names starting with one of the `exclude_prefixes` must satisfy the object conditions. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).
            * `maxTimeElapsedSinceLastModification` (`pulumi.Input[str]`) - A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
            * `minTimeElapsedSinceLastModification` (`pulumi.Input[str]`) - 
              A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".

          * `transferOptions` (`pulumi.Input[dict]`) - Characteristics of how to treat files from datasource and sink during job. If the option `delete_objects_unique_in_sink` is true, object conditions based on objects' `last_modification_time` are ignored and do not exclude objects in a data source or a data sink. Structure documented below.
            * `deleteObjectsFromSourceAfterTransfer` (`pulumi.Input[bool]`) - Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and `delete_objects_unique_in_sink` are mutually exclusive.
            * `deleteObjectsUniqueInSink` (`pulumi.Input[bool]`) - Whether objects that exist only in the sink should be deleted. Note that this option and
              `delete_objects_from_source_after_transfer` are mutually exclusive.
            * `overwriteObjectsAlreadyExistingInSink` (`pulumi.Input[bool]`) - Whether overwriting objects that already exist in the sink is allowed.
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

            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            __props__['project'] = project
            if schedule is None:
                raise TypeError("Missing required property 'schedule'")
            __props__['schedule'] = schedule
            __props__['status'] = status
            if transfer_spec is None:
                raise TypeError("Missing required property 'transfer_spec'")
            __props__['transfer_spec'] = transfer_spec
            __props__['creation_time'] = None
            __props__['deletion_time'] = None
            __props__['last_modification_time'] = None
            __props__['name'] = None
        super(TransferJob, __self__).__init__(
            'gcp:storage/transferJob:TransferJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, creation_time=None, deletion_time=None, description=None, last_modification_time=None, name=None, project=None, schedule=None, status=None, transfer_spec=None):
        """
        Get an existing TransferJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_time: When the Transfer Job was created.
        :param pulumi.Input[str] deletion_time: When the Transfer Job was deleted.
        :param pulumi.Input[str] description: Unique description to identify the Transfer Job.
        :param pulumi.Input[str] last_modification_time: When the Transfer Job was last modified.
        :param pulumi.Input[str] name: The name of the Transfer Job.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[dict] schedule: Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
        :param pulumi.Input[str] status: Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        :param pulumi.Input[dict] transfer_spec: Transfer specification. Structure documented below.

        The **schedule** object supports the following:

          * `scheduleEndDate` (`pulumi.Input[dict]`) - The last day the recurring transfer will be run. If `schedule_end_date` is the same as `schedule_start_date`, the transfer will be executed only once. Structure documented below.
            * `day` (`pulumi.Input[float]`) - Day of month. Must be from 1 to 31 and valid for the year and month.
            * `month` (`pulumi.Input[float]`) - Month of year. Must be from 1 to 12.
            * `year` (`pulumi.Input[float]`) - Year of date. Must be from 1 to 9999.

          * `scheduleStartDate` (`pulumi.Input[dict]`) - The first day the recurring transfer is scheduled to run. If `schedule_start_date` is in the past, the transfer will run for the first time on the following day. Structure documented below.
            * `day` (`pulumi.Input[float]`) - Day of month. Must be from 1 to 31 and valid for the year and month.
            * `month` (`pulumi.Input[float]`) - Month of year. Must be from 1 to 12.
            * `year` (`pulumi.Input[float]`) - Year of date. Must be from 1 to 9999.

          * `startTimeOfDay` (`pulumi.Input[dict]`) - The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer's start time in a day is specified in your local timezone. Structure documented below.
            * `hours` (`pulumi.Input[float]`) - Hours of day in 24 hour format. Should be from 0 to 23
            * `minutes` (`pulumi.Input[float]`) - Minutes of hour of day. Must be from 0 to 59.
            * `nanos` (`pulumi.Input[float]`) - Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
            * `seconds` (`pulumi.Input[float]`) - Seconds of minutes of the time. Must normally be from 0 to 59.

        The **transfer_spec** object supports the following:

          * `awsS3DataSource` (`pulumi.Input[dict]`) - An AWS S3 data source. Structure documented below.
            * `awsAccessKey` (`pulumi.Input[dict]`) - AWS credentials block.
              * `accessKeyId` (`pulumi.Input[str]`) - AWS Key ID.
              * `secretAccessKey` (`pulumi.Input[str]`) - AWS Secret Access Key.

            * `bucket_name` (`pulumi.Input[str]`) - S3 Bucket name.

          * `gcsDataSink` (`pulumi.Input[dict]`) - A Google Cloud Storage data sink. Structure documented below.
            * `bucket_name` (`pulumi.Input[str]`) - S3 Bucket name.

          * `gcsDataSource` (`pulumi.Input[dict]`) - A Google Cloud Storage data source. Structure documented below.
            * `bucket_name` (`pulumi.Input[str]`) - S3 Bucket name.

          * `httpDataSource` (`pulumi.Input[dict]`) - An HTTP URL data source. Structure documented below.
            * `listUrl` (`pulumi.Input[str]`) - The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.

          * `objectConditions` (`pulumi.Input[dict]`) - Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' `last_modification_time` do not exclude objects in a data sink. Structure documented below.
            * `excludePrefixes` (`pulumi.Input[list]`) - `exclude_prefixes` must follow the requirements described for `include_prefixes`. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).
            * `includePrefixes` (`pulumi.Input[list]`) - If `include_refixes` is specified, objects that satisfy the object conditions must have names that start with one of the `include_prefixes` and that do not start with any of the `exclude_prefixes`. If `include_prefixes` is not specified, all objects except those that have names starting with one of the `exclude_prefixes` must satisfy the object conditions. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).
            * `maxTimeElapsedSinceLastModification` (`pulumi.Input[str]`) - A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
            * `minTimeElapsedSinceLastModification` (`pulumi.Input[str]`) - 
              A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".

          * `transferOptions` (`pulumi.Input[dict]`) - Characteristics of how to treat files from datasource and sink during job. If the option `delete_objects_unique_in_sink` is true, object conditions based on objects' `last_modification_time` are ignored and do not exclude objects in a data source or a data sink. Structure documented below.
            * `deleteObjectsFromSourceAfterTransfer` (`pulumi.Input[bool]`) - Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and `delete_objects_unique_in_sink` are mutually exclusive.
            * `deleteObjectsUniqueInSink` (`pulumi.Input[bool]`) - Whether objects that exist only in the sink should be deleted. Note that this option and
              `delete_objects_from_source_after_transfer` are mutually exclusive.
            * `overwriteObjectsAlreadyExistingInSink` (`pulumi.Input[bool]`) - Whether overwriting objects that already exist in the sink is allowed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_time"] = creation_time
        __props__["deletion_time"] = deletion_time
        __props__["description"] = description
        __props__["last_modification_time"] = last_modification_time
        __props__["name"] = name
        __props__["project"] = project
        __props__["schedule"] = schedule
        __props__["status"] = status
        __props__["transfer_spec"] = transfer_spec
        return TransferJob(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
