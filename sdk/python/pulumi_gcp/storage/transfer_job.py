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

__all__ = ['TransferJobArgs', 'TransferJob']

@pulumi.input_type
class TransferJobArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 schedule: pulumi.Input['TransferJobScheduleArgs'],
                 transfer_spec: pulumi.Input['TransferJobTransferSpecArgs'],
                 project: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TransferJob resource.
        :param pulumi.Input[str] description: Unique description to identify the Transfer Job.
        :param pulumi.Input['TransferJobScheduleArgs'] schedule: Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        :param pulumi.Input['TransferJobTransferSpecArgs'] transfer_spec: Transfer specification. Structure documented below.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] status: Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "transfer_spec", transfer_spec)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Unique description to identify the Transfer Job.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input['TransferJobScheduleArgs']:
        """
        Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input['TransferJobScheduleArgs']):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="transferSpec")
    def transfer_spec(self) -> pulumi.Input['TransferJobTransferSpecArgs']:
        """
        Transfer specification. Structure documented below.
        """
        return pulumi.get(self, "transfer_spec")

    @transfer_spec.setter
    def transfer_spec(self, value: pulumi.Input['TransferJobTransferSpecArgs']):
        pulumi.set(self, "transfer_spec", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _TransferJobState:
    def __init__(__self__, *,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 deletion_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 last_modification_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input['TransferJobScheduleArgs']] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transfer_spec: Optional[pulumi.Input['TransferJobTransferSpecArgs']] = None):
        """
        Input properties used for looking up and filtering TransferJob resources.
        :param pulumi.Input[str] creation_time: When the Transfer Job was created.
        :param pulumi.Input[str] deletion_time: When the Transfer Job was deleted.
        :param pulumi.Input[str] description: Unique description to identify the Transfer Job.
        :param pulumi.Input[str] last_modification_time: When the Transfer Job was last modified.
        :param pulumi.Input[str] name: The name of the Transfer Job.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input['TransferJobScheduleArgs'] schedule: Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        :param pulumi.Input[str] status: Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        :param pulumi.Input['TransferJobTransferSpecArgs'] transfer_spec: Transfer specification. Structure documented below.
        """
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if deletion_time is not None:
            pulumi.set(__self__, "deletion_time", deletion_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if last_modification_time is not None:
            pulumi.set(__self__, "last_modification_time", last_modification_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transfer_spec is not None:
            pulumi.set(__self__, "transfer_spec", transfer_spec)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        When the Transfer Job was created.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter(name="deletionTime")
    def deletion_time(self) -> Optional[pulumi.Input[str]]:
        """
        When the Transfer Job was deleted.
        """
        return pulumi.get(self, "deletion_time")

    @deletion_time.setter
    def deletion_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deletion_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Unique description to identify the Transfer Job.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="lastModificationTime")
    def last_modification_time(self) -> Optional[pulumi.Input[str]]:
        """
        When the Transfer Job was last modified.
        """
        return pulumi.get(self, "last_modification_time")

    @last_modification_time.setter
    def last_modification_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modification_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Transfer Job.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input['TransferJobScheduleArgs']]:
        """
        Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['TransferJobScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="transferSpec")
    def transfer_spec(self) -> Optional[pulumi.Input['TransferJobTransferSpecArgs']]:
        """
        Transfer specification. Structure documented below.
        """
        return pulumi.get(self, "transfer_spec")

    @transfer_spec.setter
    def transfer_spec(self, value: Optional[pulumi.Input['TransferJobTransferSpecArgs']]):
        pulumi.set(self, "transfer_spec", value)


class TransferJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['TransferJobScheduleArgs']]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transfer_spec: Optional[pulumi.Input[pulumi.InputType['TransferJobTransferSpecArgs']]] = None,
                 __props__=None):
        """
        Creates a new Transfer Job in Google Cloud Storage Transfer.

        To get more information about Google Cloud Storage Transfer, see:

        * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
        * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs#TransferJob)
        * How-to Guides
            * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)

        ## Example Usage

        Example creating a nightly Transfer Job from an AWS S3 Bucket to a GCS bucket.

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
            member=f"serviceAccount:{default.email}",
            opts=pulumi.ResourceOptions(depends_on=[s3_backup_bucket_bucket]))
        s3_bucket_nightly_backup = gcp.storage.TransferJob("s3-bucket-nightly-backup",
            description="Nightly backup of S3 bucket",
            project=var["project"],
            transfer_spec=gcp.storage.TransferJobTransferSpecArgs(
                object_conditions=gcp.storage.TransferJobTransferSpecObjectConditionsArgs(
                    max_time_elapsed_since_last_modification="600s",
                    exclude_prefixes=["requests.gz"],
                ),
                transfer_options=gcp.storage.TransferJobTransferSpecTransferOptionsArgs(
                    delete_objects_unique_in_sink=False,
                ),
                aws_s3_data_source=gcp.storage.TransferJobTransferSpecAwsS3DataSourceArgs(
                    bucket_name=var["aws_s3_bucket"],
                    aws_access_key=gcp.storage.TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs(
                        access_key_id=var["aws_access_key"],
                        secret_access_key=var["aws_secret_key"],
                    ),
                ),
                gcs_data_sink=gcp.storage.TransferJobTransferSpecGcsDataSinkArgs(
                    bucket_name=s3_backup_bucket_bucket.name,
                ),
            ),
            schedule=gcp.storage.TransferJobScheduleArgs(
                schedule_start_date=gcp.storage.TransferJobScheduleScheduleStartDateArgs(
                    year=2018,
                    month=10,
                    day=1,
                ),
                schedule_end_date=gcp.storage.TransferJobScheduleScheduleEndDateArgs(
                    year=2019,
                    month=1,
                    day=15,
                ),
                start_time_of_day=gcp.storage.TransferJobScheduleStartTimeOfDayArgs(
                    hours=23,
                    minutes=30,
                    seconds=0,
                    nanos=0,
                ),
            ),
            opts=pulumi.ResourceOptions(depends_on=[s3_backup_bucket_bucket_iam_member]))
        ```

        ## Import

        Storage buckets can be imported using the Transfer Job's `project` and `name` without the `transferJob/` prefix, e.g.

        ```sh
         $ pulumi import gcp:storage/transferJob:TransferJob nightly-backup-transfer-job my-project-1asd32/8422144862922355674
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Unique description to identify the Transfer Job.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['TransferJobScheduleArgs']] schedule: Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        :param pulumi.Input[str] status: Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        :param pulumi.Input[pulumi.InputType['TransferJobTransferSpecArgs']] transfer_spec: Transfer specification. Structure documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: TransferJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Transfer Job in Google Cloud Storage Transfer.

        To get more information about Google Cloud Storage Transfer, see:

        * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
        * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs#TransferJob)
        * How-to Guides
            * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)

        ## Example Usage

        Example creating a nightly Transfer Job from an AWS S3 Bucket to a GCS bucket.

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
            member=f"serviceAccount:{default.email}",
            opts=pulumi.ResourceOptions(depends_on=[s3_backup_bucket_bucket]))
        s3_bucket_nightly_backup = gcp.storage.TransferJob("s3-bucket-nightly-backup",
            description="Nightly backup of S3 bucket",
            project=var["project"],
            transfer_spec=gcp.storage.TransferJobTransferSpecArgs(
                object_conditions=gcp.storage.TransferJobTransferSpecObjectConditionsArgs(
                    max_time_elapsed_since_last_modification="600s",
                    exclude_prefixes=["requests.gz"],
                ),
                transfer_options=gcp.storage.TransferJobTransferSpecTransferOptionsArgs(
                    delete_objects_unique_in_sink=False,
                ),
                aws_s3_data_source=gcp.storage.TransferJobTransferSpecAwsS3DataSourceArgs(
                    bucket_name=var["aws_s3_bucket"],
                    aws_access_key=gcp.storage.TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs(
                        access_key_id=var["aws_access_key"],
                        secret_access_key=var["aws_secret_key"],
                    ),
                ),
                gcs_data_sink=gcp.storage.TransferJobTransferSpecGcsDataSinkArgs(
                    bucket_name=s3_backup_bucket_bucket.name,
                ),
            ),
            schedule=gcp.storage.TransferJobScheduleArgs(
                schedule_start_date=gcp.storage.TransferJobScheduleScheduleStartDateArgs(
                    year=2018,
                    month=10,
                    day=1,
                ),
                schedule_end_date=gcp.storage.TransferJobScheduleScheduleEndDateArgs(
                    year=2019,
                    month=1,
                    day=15,
                ),
                start_time_of_day=gcp.storage.TransferJobScheduleStartTimeOfDayArgs(
                    hours=23,
                    minutes=30,
                    seconds=0,
                    nanos=0,
                ),
            ),
            opts=pulumi.ResourceOptions(depends_on=[s3_backup_bucket_bucket_iam_member]))
        ```

        ## Import

        Storage buckets can be imported using the Transfer Job's `project` and `name` without the `transferJob/` prefix, e.g.

        ```sh
         $ pulumi import gcp:storage/transferJob:TransferJob nightly-backup-transfer-job my-project-1asd32/8422144862922355674
        ```

        :param str resource_name_: The name of the resource.
        :param TransferJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransferJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['TransferJobScheduleArgs']]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transfer_spec: Optional[pulumi.Input[pulumi.InputType['TransferJobTransferSpecArgs']]] = None,
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
            __props__ = TransferJobArgs.__new__(TransferJobArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["project"] = project
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["status"] = status
            if transfer_spec is None and not opts.urn:
                raise TypeError("Missing required property 'transfer_spec'")
            __props__.__dict__["transfer_spec"] = transfer_spec
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["deletion_time"] = None
            __props__.__dict__["last_modification_time"] = None
            __props__.__dict__["name"] = None
        super(TransferJob, __self__).__init__(
            'gcp:storage/transferJob:TransferJob',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            deletion_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            last_modification_time: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[pulumi.InputType['TransferJobScheduleArgs']]] = None,
            status: Optional[pulumi.Input[str]] = None,
            transfer_spec: Optional[pulumi.Input[pulumi.InputType['TransferJobTransferSpecArgs']]] = None) -> 'TransferJob':
        """
        Get an existing TransferJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_time: When the Transfer Job was created.
        :param pulumi.Input[str] deletion_time: When the Transfer Job was deleted.
        :param pulumi.Input[str] description: Unique description to identify the Transfer Job.
        :param pulumi.Input[str] last_modification_time: When the Transfer Job was last modified.
        :param pulumi.Input[str] name: The name of the Transfer Job.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['TransferJobScheduleArgs']] schedule: Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        :param pulumi.Input[str] status: Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        :param pulumi.Input[pulumi.InputType['TransferJobTransferSpecArgs']] transfer_spec: Transfer specification. Structure documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransferJobState.__new__(_TransferJobState)

        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["deletion_time"] = deletion_time
        __props__.__dict__["description"] = description
        __props__.__dict__["last_modification_time"] = last_modification_time
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["status"] = status
        __props__.__dict__["transfer_spec"] = transfer_spec
        return TransferJob(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        When the Transfer Job was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="deletionTime")
    def deletion_time(self) -> pulumi.Output[str]:
        """
        When the Transfer Job was deleted.
        """
        return pulumi.get(self, "deletion_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Unique description to identify the Transfer Job.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModificationTime")
    def last_modification_time(self) -> pulumi.Output[str]:
        """
        When the Transfer Job was last modified.
        """
        return pulumi.get(self, "last_modification_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Transfer Job.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output['outputs.TransferJobSchedule']:
        """
        Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transferSpec")
    def transfer_spec(self) -> pulumi.Output['outputs.TransferJobTransferSpec']:
        """
        Transfer specification. Structure documented below.
        """
        return pulumi.get(self, "transfer_spec")

