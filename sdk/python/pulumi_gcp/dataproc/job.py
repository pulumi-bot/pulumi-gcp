# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Job(pulumi.CustomResource):
    driver_controls_files_uri: pulumi.Output[str]
    """
    If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
    """
    driver_output_resource_uri: pulumi.Output[str]
    """
    A URI pointing to the location of the stdout of the job's driver program.
    """
    force_delete: pulumi.Output[bool]
    """
    By default, you can only delete inactive jobs within
    Dataproc. Setting this to true, and calling destroy, will ensure that the
    job is first cancelled before issuing the delete.
    """
    hadoop_config: pulumi.Output[dict]
    hive_config: pulumi.Output[dict]
    labels: pulumi.Output[dict]
    """
    The list of labels (key/value pairs) to add to the job.
    """
    pig_config: pulumi.Output[dict]
    placement: pulumi.Output[dict]
    project: pulumi.Output[str]
    """
    The project in which the `cluster` can be found and jobs
    subsequently run against. If it is not provided, the provider project is used.
    """
    pyspark_config: pulumi.Output[dict]
    reference: pulumi.Output[dict]
    region: pulumi.Output[str]
    """
    The Cloud Dataproc region. This essentially determines which clusters are available
    for this job to be submitted to. If not specified, defaults to `global`.
    """
    scheduling: pulumi.Output[dict]
    spark_config: pulumi.Output[dict]
    sparksql_config: pulumi.Output[dict]
    status: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, force_delete=None, hadoop_config=None, hive_config=None, labels=None, pig_config=None, placement=None, project=None, pyspark_config=None, reference=None, region=None, scheduling=None, spark_config=None, sparksql_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a job resource within a Dataproc cluster within GCE. For more information see
        [the official dataproc documentation](https://cloud.google.com/dataproc/).
        
        !> **Note:** This resource does not support 'update' and changing any attributes will cause the resource to be recreated.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] force_delete: By default, you can only delete inactive jobs within
               Dataproc. Setting this to true, and calling destroy, will ensure that the
               job is first cancelled before issuing the delete.
        :param pulumi.Input[dict] labels: The list of labels (key/value pairs) to add to the job.
        :param pulumi.Input[str] project: The project in which the `cluster` can be found and jobs
               subsequently run against. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Cloud Dataproc region. This essentially determines which clusters are available
               for this job to be submitted to. If not specified, defaults to `global`.
        
        The **pig_config** object supports the following:
        
          * `continue_on_failure` (`pulumi.Input[bool]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `properties` (`pulumi.Input[dict]`)
          * `query_file_uri` (`pulumi.Input[str]`)
          * `query_lists` (`pulumi.Input[list]`)
          * `script_variables` (`pulumi.Input[dict]`)
        
        The **placement** object supports the following:
        
          * `cluster_name` (`pulumi.Input[str]`)
          * `cluster_uuid` (`pulumi.Input[str]`)
        
        The **reference** object supports the following:
        
          * `job_id` (`pulumi.Input[str]`)
        
        The **scheduling** object supports the following:
        
          * `max_failures_per_hour` (`pulumi.Input[float]`)
        
        The **spark_config** object supports the following:
        
          * `archive_uris` (`pulumi.Input[list]`)
          * `args` (`pulumi.Input[list]`)
          * `file_uris` (`pulumi.Input[list]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `main_class` (`pulumi.Input[str]`)
          * `main_jar_file_uri` (`pulumi.Input[str]`)
          * `properties` (`pulumi.Input[dict]`)
        
        The **hadoop_config** object supports the following:
        
          * `archive_uris` (`pulumi.Input[list]`)
          * `args` (`pulumi.Input[list]`)
          * `file_uris` (`pulumi.Input[list]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `main_class` (`pulumi.Input[str]`)
          * `main_jar_file_uri` (`pulumi.Input[str]`)
          * `properties` (`pulumi.Input[dict]`)
        
        The **hive_config** object supports the following:
        
          * `continue_on_failure` (`pulumi.Input[bool]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `properties` (`pulumi.Input[dict]`)
          * `query_file_uri` (`pulumi.Input[str]`)
          * `query_lists` (`pulumi.Input[list]`)
          * `script_variables` (`pulumi.Input[dict]`)
        
        The **pyspark_config** object supports the following:
        
          * `archive_uris` (`pulumi.Input[list]`)
          * `args` (`pulumi.Input[list]`)
          * `file_uris` (`pulumi.Input[list]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `main_python_file_uri` (`pulumi.Input[str]`)
          * `properties` (`pulumi.Input[dict]`)
          * `python_file_uris` (`pulumi.Input[list]`)
        
        The **sparksql_config** object supports the following:
        
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `properties` (`pulumi.Input[dict]`)
          * `query_file_uri` (`pulumi.Input[str]`)
          * `query_lists` (`pulumi.Input[list]`)
          * `script_variables` (`pulumi.Input[dict]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_job.html.markdown.
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
    def get(resource_name, id, opts=None, driver_controls_files_uri=None, driver_output_resource_uri=None, force_delete=None, hadoop_config=None, hive_config=None, labels=None, pig_config=None, placement=None, project=None, pyspark_config=None, reference=None, region=None, scheduling=None, spark_config=None, sparksql_config=None, status=None):
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] driver_controls_files_uri: If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        :param pulumi.Input[str] driver_output_resource_uri: A URI pointing to the location of the stdout of the job's driver program.
        :param pulumi.Input[bool] force_delete: By default, you can only delete inactive jobs within
               Dataproc. Setting this to true, and calling destroy, will ensure that the
               job is first cancelled before issuing the delete.
        :param pulumi.Input[dict] labels: The list of labels (key/value pairs) to add to the job.
        :param pulumi.Input[str] project: The project in which the `cluster` can be found and jobs
               subsequently run against. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Cloud Dataproc region. This essentially determines which clusters are available
               for this job to be submitted to. If not specified, defaults to `global`.
        
        The **sparksql_config** object supports the following:
        
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `properties` (`pulumi.Input[dict]`)
          * `query_file_uri` (`pulumi.Input[str]`)
          * `query_lists` (`pulumi.Input[list]`)
          * `script_variables` (`pulumi.Input[dict]`)
        
        The **status** object supports the following:
        
          * `details` (`pulumi.Input[str]`)
          * `state` (`pulumi.Input[str]`)
          * `state_start_time` (`pulumi.Input[str]`)
          * `substate` (`pulumi.Input[str]`)
        
        The **hive_config** object supports the following:
        
          * `continue_on_failure` (`pulumi.Input[bool]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `properties` (`pulumi.Input[dict]`)
          * `query_file_uri` (`pulumi.Input[str]`)
          * `query_lists` (`pulumi.Input[list]`)
          * `script_variables` (`pulumi.Input[dict]`)
        
        The **pig_config** object supports the following:
        
          * `continue_on_failure` (`pulumi.Input[bool]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `properties` (`pulumi.Input[dict]`)
          * `query_file_uri` (`pulumi.Input[str]`)
          * `query_lists` (`pulumi.Input[list]`)
          * `script_variables` (`pulumi.Input[dict]`)
        
        The **pyspark_config** object supports the following:
        
          * `archive_uris` (`pulumi.Input[list]`)
          * `args` (`pulumi.Input[list]`)
          * `file_uris` (`pulumi.Input[list]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `main_python_file_uri` (`pulumi.Input[str]`)
          * `properties` (`pulumi.Input[dict]`)
          * `python_file_uris` (`pulumi.Input[list]`)
        
        The **reference** object supports the following:
        
          * `job_id` (`pulumi.Input[str]`)
        
        The **hadoop_config** object supports the following:
        
          * `archive_uris` (`pulumi.Input[list]`)
          * `args` (`pulumi.Input[list]`)
          * `file_uris` (`pulumi.Input[list]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `main_class` (`pulumi.Input[str]`)
          * `main_jar_file_uri` (`pulumi.Input[str]`)
          * `properties` (`pulumi.Input[dict]`)
        
        The **placement** object supports the following:
        
          * `cluster_name` (`pulumi.Input[str]`)
          * `cluster_uuid` (`pulumi.Input[str]`)
        
        The **scheduling** object supports the following:
        
          * `max_failures_per_hour` (`pulumi.Input[float]`)
        
        The **spark_config** object supports the following:
        
          * `archive_uris` (`pulumi.Input[list]`)
          * `args` (`pulumi.Input[list]`)
          * `file_uris` (`pulumi.Input[list]`)
          * `jar_file_uris` (`pulumi.Input[list]`)
          * `logging_config` (`pulumi.Input[dict]`)
        
            * `driver_log_levels` (`pulumi.Input[dict]`)
        
          * `main_class` (`pulumi.Input[str]`)
          * `main_jar_file_uri` (`pulumi.Input[str]`)
          * `properties` (`pulumi.Input[dict]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_job.html.markdown.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

