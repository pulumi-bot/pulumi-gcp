# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Job']


class Job(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 hadoop_config: Optional[pulumi.Input[pulumi.InputType['JobHadoopConfigArgs']]] = None,
                 hive_config: Optional[pulumi.Input[pulumi.InputType['JobHiveConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 pig_config: Optional[pulumi.Input[pulumi.InputType['JobPigConfigArgs']]] = None,
                 placement: Optional[pulumi.Input[pulumi.InputType['JobPlacementArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 pyspark_config: Optional[pulumi.Input[pulumi.InputType['JobPysparkConfigArgs']]] = None,
                 reference: Optional[pulumi.Input[pulumi.InputType['JobReferenceArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scheduling: Optional[pulumi.Input[pulumi.InputType['JobSchedulingArgs']]] = None,
                 spark_config: Optional[pulumi.Input[pulumi.InputType['JobSparkConfigArgs']]] = None,
                 sparksql_config: Optional[pulumi.Input[pulumi.InputType['JobSparksqlConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a job resource within a Dataproc cluster within GCE. For more information see
        [the official dataproc documentation](https://cloud.google.com/dataproc/).

        !> **Note:** This resource does not support 'update' and changing any attributes will cause the resource to be recreated.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] force_delete: By default, you can only delete inactive jobs within
               Dataproc. Setting this to true, and calling destroy, will ensure that the
               job is first cancelled before issuing the delete.
        :param pulumi.Input[pulumi.InputType['JobHadoopConfigArgs']] hadoop_config: The config of Hadoop job
        :param pulumi.Input[pulumi.InputType['JobHiveConfigArgs']] hive_config: The config of hive job
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The list of labels (key/value pairs) to add to the job.
        :param pulumi.Input[pulumi.InputType['JobPigConfigArgs']] pig_config: The config of pag job.
        :param pulumi.Input[pulumi.InputType['JobPlacementArgs']] placement: The config of job placement.
        :param pulumi.Input[str] project: The project in which the `cluster` can be found and jobs
               subsequently run against. If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['JobPysparkConfigArgs']] pyspark_config: The config of pySpark job.
        :param pulumi.Input[pulumi.InputType['JobReferenceArgs']] reference: The reference of the job
        :param pulumi.Input[str] region: The Cloud Dataproc region. This essentially determines which clusters are available
               for this job to be submitted to. If not specified, defaults to `global`.
        :param pulumi.Input[pulumi.InputType['JobSchedulingArgs']] scheduling: Optional. Job scheduling configuration.
        :param pulumi.Input[pulumi.InputType['JobSparkConfigArgs']] spark_config: The config of the Spark job.
        :param pulumi.Input[pulumi.InputType['JobSparksqlConfigArgs']] sparksql_config: The config of SparkSql job
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['force_delete'] = force_delete
            __props__['hadoop_config'] = hadoop_config
            __props__['hive_config'] = hive_config
            __props__['labels'] = labels
            __props__['pig_config'] = pig_config
            if placement is None:
                raise TypeError("Missing required property 'placement'")
            __props__['placement'] = placement
            __props__['project'] = project
            __props__['pyspark_config'] = pyspark_config
            __props__['reference'] = reference
            __props__['region'] = region
            __props__['scheduling'] = scheduling
            __props__['spark_config'] = spark_config
            __props__['sparksql_config'] = sparksql_config
            __props__['driver_controls_files_uri'] = None
            __props__['driver_output_resource_uri'] = None
            __props__['status'] = None
        super(Job, __self__).__init__(
            'gcp:dataproc/job:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            driver_controls_files_uri: Optional[pulumi.Input[str]] = None,
            driver_output_resource_uri: Optional[pulumi.Input[str]] = None,
            force_delete: Optional[pulumi.Input[bool]] = None,
            hadoop_config: Optional[pulumi.Input[pulumi.InputType['JobHadoopConfigArgs']]] = None,
            hive_config: Optional[pulumi.Input[pulumi.InputType['JobHiveConfigArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            pig_config: Optional[pulumi.Input[pulumi.InputType['JobPigConfigArgs']]] = None,
            placement: Optional[pulumi.Input[pulumi.InputType['JobPlacementArgs']]] = None,
            project: Optional[pulumi.Input[str]] = None,
            pyspark_config: Optional[pulumi.Input[pulumi.InputType['JobPysparkConfigArgs']]] = None,
            reference: Optional[pulumi.Input[pulumi.InputType['JobReferenceArgs']]] = None,
            region: Optional[pulumi.Input[str]] = None,
            scheduling: Optional[pulumi.Input[pulumi.InputType['JobSchedulingArgs']]] = None,
            spark_config: Optional[pulumi.Input[pulumi.InputType['JobSparkConfigArgs']]] = None,
            sparksql_config: Optional[pulumi.Input[pulumi.InputType['JobSparksqlConfigArgs']]] = None,
            status: Optional[pulumi.Input[pulumi.InputType['JobStatusArgs']]] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] driver_controls_files_uri: If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        :param pulumi.Input[str] driver_output_resource_uri: A URI pointing to the location of the stdout of the job's driver program.
        :param pulumi.Input[bool] force_delete: By default, you can only delete inactive jobs within
               Dataproc. Setting this to true, and calling destroy, will ensure that the
               job is first cancelled before issuing the delete.
        :param pulumi.Input[pulumi.InputType['JobHadoopConfigArgs']] hadoop_config: The config of Hadoop job
        :param pulumi.Input[pulumi.InputType['JobHiveConfigArgs']] hive_config: The config of hive job
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The list of labels (key/value pairs) to add to the job.
        :param pulumi.Input[pulumi.InputType['JobPigConfigArgs']] pig_config: The config of pag job.
        :param pulumi.Input[pulumi.InputType['JobPlacementArgs']] placement: The config of job placement.
        :param pulumi.Input[str] project: The project in which the `cluster` can be found and jobs
               subsequently run against. If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['JobPysparkConfigArgs']] pyspark_config: The config of pySpark job.
        :param pulumi.Input[pulumi.InputType['JobReferenceArgs']] reference: The reference of the job
        :param pulumi.Input[str] region: The Cloud Dataproc region. This essentially determines which clusters are available
               for this job to be submitted to. If not specified, defaults to `global`.
        :param pulumi.Input[pulumi.InputType['JobSchedulingArgs']] scheduling: Optional. Job scheduling configuration.
        :param pulumi.Input[pulumi.InputType['JobSparkConfigArgs']] spark_config: The config of the Spark job.
        :param pulumi.Input[pulumi.InputType['JobSparksqlConfigArgs']] sparksql_config: The config of SparkSql job
        :param pulumi.Input[pulumi.InputType['JobStatusArgs']] status: The status of the job.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["driver_controls_files_uri"] = driver_controls_files_uri
        __props__["driver_output_resource_uri"] = driver_output_resource_uri
        __props__["force_delete"] = force_delete
        __props__["hadoop_config"] = hadoop_config
        __props__["hive_config"] = hive_config
        __props__["labels"] = labels
        __props__["pig_config"] = pig_config
        __props__["placement"] = placement
        __props__["project"] = project
        __props__["pyspark_config"] = pyspark_config
        __props__["reference"] = reference
        __props__["region"] = region
        __props__["scheduling"] = scheduling
        __props__["spark_config"] = spark_config
        __props__["sparksql_config"] = sparksql_config
        __props__["status"] = status
        return Job(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="driverControlsFilesUri")
    def driver_controls_files_uri(self) -> str:
        """
        If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        """
        return pulumi.get(self, "driver_controls_files_uri")

    @property
    @pulumi.getter(name="driverOutputResourceUri")
    def driver_output_resource_uri(self) -> str:
        """
        A URI pointing to the location of the stdout of the job's driver program.
        """
        return pulumi.get(self, "driver_output_resource_uri")

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[bool]:
        """
        By default, you can only delete inactive jobs within
        Dataproc. Setting this to true, and calling destroy, will ensure that the
        job is first cancelled before issuing the delete.
        """
        return pulumi.get(self, "force_delete")

    @property
    @pulumi.getter(name="hadoopConfig")
    def hadoop_config(self) -> Optional['outputs.JobHadoopConfig']:
        """
        The config of Hadoop job
        """
        return pulumi.get(self, "hadoop_config")

    @property
    @pulumi.getter(name="hiveConfig")
    def hive_config(self) -> Optional['outputs.JobHiveConfig']:
        """
        The config of hive job
        """
        return pulumi.get(self, "hive_config")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        The list of labels (key/value pairs) to add to the job.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="pigConfig")
    def pig_config(self) -> Optional['outputs.JobPigConfig']:
        """
        The config of pag job.
        """
        return pulumi.get(self, "pig_config")

    @property
    @pulumi.getter
    def placement(self) -> 'outputs.JobPlacement':
        """
        The config of job placement.
        """
        return pulumi.get(self, "placement")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The project in which the `cluster` can be found and jobs
        subsequently run against. If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="pysparkConfig")
    def pyspark_config(self) -> Optional['outputs.JobPysparkConfig']:
        """
        The config of pySpark job.
        """
        return pulumi.get(self, "pyspark_config")

    @property
    @pulumi.getter
    def reference(self) -> 'outputs.JobReference':
        """
        The reference of the job
        """
        return pulumi.get(self, "reference")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The Cloud Dataproc region. This essentially determines which clusters are available
        for this job to be submitted to. If not specified, defaults to `global`.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def scheduling(self) -> Optional['outputs.JobScheduling']:
        """
        Optional. Job scheduling configuration.
        """
        return pulumi.get(self, "scheduling")

    @property
    @pulumi.getter(name="sparkConfig")
    def spark_config(self) -> Optional['outputs.JobSparkConfig']:
        """
        The config of the Spark job.
        """
        return pulumi.get(self, "spark_config")

    @property
    @pulumi.getter(name="sparksqlConfig")
    def sparksql_config(self) -> Optional['outputs.JobSparksqlConfig']:
        """
        The config of SparkSql job
        """
        return pulumi.get(self, "sparksql_config")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.JobStatus':
        """
        The status of the job.
        """
        return pulumi.get(self, "status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

