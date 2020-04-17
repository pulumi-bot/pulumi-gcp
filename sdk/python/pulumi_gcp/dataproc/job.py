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
    """
    Optional. Job scheduling configuration.

      * `maxFailuresPerHour` (`float`)
    """
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
        :param pulumi.Input[dict] scheduling: Optional. Job scheduling configuration.

        The **hadoop_config** object supports the following:

          * `archiveUris` (`pulumi.Input[list]`) - HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
          * `args` (`pulumi.Input[list]`) - The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
          * `fileUris` (`pulumi.Input[list]`) - HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `mainClass` (`pulumi.Input[str]`) - The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`. Conflicts with `main_jar_file_uri`
          * `mainJarFileUri` (`pulumi.Input[str]`) - The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'. Conflicts with `main_class`
          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.

        The **hive_config** object supports the following:

          * `continueOnFailure` (`pulumi.Input[bool]`) - Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
          * `queryFileUri` (`pulumi.Input[str]`) - The HCFS URI of the script that contains SQL queries.
            Conflicts with `query_list`
          * `queryLists` (`pulumi.Input[list]`) - The list of SQL queries or statements to execute as part of the job.
            Conflicts with `query_file_uri`
          * `scriptVariables` (`pulumi.Input[dict]`) - Mapping of query variable names to values (equivalent to the Spark SQL command: `SET name="value";`).

        The **pig_config** object supports the following:

          * `continueOnFailure` (`pulumi.Input[bool]`) - Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
          * `queryFileUri` (`pulumi.Input[str]`) - The HCFS URI of the script that contains SQL queries.
            Conflicts with `query_list`
          * `queryLists` (`pulumi.Input[list]`) - The list of SQL queries or statements to execute as part of the job.
            Conflicts with `query_file_uri`
          * `scriptVariables` (`pulumi.Input[dict]`) - Mapping of query variable names to values (equivalent to the Spark SQL command: `SET name="value";`).

        The **placement** object supports the following:

          * `clusterName` (`pulumi.Input[str]`)
          * `clusterUuid` (`pulumi.Input[str]`)

        The **pyspark_config** object supports the following:

          * `archiveUris` (`pulumi.Input[list]`) - HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
          * `args` (`pulumi.Input[list]`) - The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
          * `fileUris` (`pulumi.Input[list]`) - HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `mainPythonFileUri` (`pulumi.Input[str]`) - The HCFS URI of the main Python file to use as the driver. Must be a .py file.
          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
          * `pythonFileUris` (`pulumi.Input[list]`) - HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.

        The **reference** object supports the following:

          * `job_id` (`pulumi.Input[str]`)

        The **scheduling** object supports the following:

          * `maxFailuresPerHour` (`pulumi.Input[float]`)

        The **spark_config** object supports the following:

          * `archiveUris` (`pulumi.Input[list]`) - HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
          * `args` (`pulumi.Input[list]`) - The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
          * `fileUris` (`pulumi.Input[list]`) - HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `mainClass` (`pulumi.Input[str]`) - The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`. Conflicts with `main_jar_file_uri`
          * `mainJarFileUri` (`pulumi.Input[str]`) - The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'. Conflicts with `main_class`
          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.

        The **sparksql_config** object supports the following:

          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
          * `queryFileUri` (`pulumi.Input[str]`) - The HCFS URI of the script that contains SQL queries.
            Conflicts with `query_list`
          * `queryLists` (`pulumi.Input[list]`) - The list of SQL queries or statements to execute as part of the job.
            Conflicts with `query_file_uri`
          * `scriptVariables` (`pulumi.Input[dict]`) - Mapping of query variable names to values (equivalent to the Spark SQL command: `SET name="value";`).
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
        :param pulumi.Input[dict] scheduling: Optional. Job scheduling configuration.

        The **hadoop_config** object supports the following:

          * `archiveUris` (`pulumi.Input[list]`) - HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
          * `args` (`pulumi.Input[list]`) - The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
          * `fileUris` (`pulumi.Input[list]`) - HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `mainClass` (`pulumi.Input[str]`) - The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`. Conflicts with `main_jar_file_uri`
          * `mainJarFileUri` (`pulumi.Input[str]`) - The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'. Conflicts with `main_class`
          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.

        The **hive_config** object supports the following:

          * `continueOnFailure` (`pulumi.Input[bool]`) - Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
          * `queryFileUri` (`pulumi.Input[str]`) - The HCFS URI of the script that contains SQL queries.
            Conflicts with `query_list`
          * `queryLists` (`pulumi.Input[list]`) - The list of SQL queries or statements to execute as part of the job.
            Conflicts with `query_file_uri`
          * `scriptVariables` (`pulumi.Input[dict]`) - Mapping of query variable names to values (equivalent to the Spark SQL command: `SET name="value";`).

        The **pig_config** object supports the following:

          * `continueOnFailure` (`pulumi.Input[bool]`) - Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
          * `queryFileUri` (`pulumi.Input[str]`) - The HCFS URI of the script that contains SQL queries.
            Conflicts with `query_list`
          * `queryLists` (`pulumi.Input[list]`) - The list of SQL queries or statements to execute as part of the job.
            Conflicts with `query_file_uri`
          * `scriptVariables` (`pulumi.Input[dict]`) - Mapping of query variable names to values (equivalent to the Spark SQL command: `SET name="value";`).

        The **placement** object supports the following:

          * `clusterName` (`pulumi.Input[str]`)
          * `clusterUuid` (`pulumi.Input[str]`)

        The **pyspark_config** object supports the following:

          * `archiveUris` (`pulumi.Input[list]`) - HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
          * `args` (`pulumi.Input[list]`) - The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
          * `fileUris` (`pulumi.Input[list]`) - HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `mainPythonFileUri` (`pulumi.Input[str]`) - The HCFS URI of the main Python file to use as the driver. Must be a .py file.
          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
          * `pythonFileUris` (`pulumi.Input[list]`) - HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.

        The **reference** object supports the following:

          * `job_id` (`pulumi.Input[str]`)

        The **scheduling** object supports the following:

          * `maxFailuresPerHour` (`pulumi.Input[float]`)

        The **spark_config** object supports the following:

          * `archiveUris` (`pulumi.Input[list]`) - HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
          * `args` (`pulumi.Input[list]`) - The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
          * `fileUris` (`pulumi.Input[list]`) - HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `mainClass` (`pulumi.Input[str]`) - The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`. Conflicts with `main_jar_file_uri`
          * `mainJarFileUri` (`pulumi.Input[str]`) - The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'. Conflicts with `main_class`
          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.

        The **sparksql_config** object supports the following:

          * `jarFileUris` (`pulumi.Input[list]`) - HCFS URIs of jar files to be added to the Spark CLASSPATH.
          * `loggingConfig` (`pulumi.Input[dict]`)
            * `driverLogLevels` (`pulumi.Input[dict]`)

          * `properties` (`pulumi.Input[dict]`) - A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
          * `queryFileUri` (`pulumi.Input[str]`) - The HCFS URI of the script that contains SQL queries.
            Conflicts with `query_list`
          * `queryLists` (`pulumi.Input[list]`) - The list of SQL queries or statements to execute as part of the job.
            Conflicts with `query_file_uri`
          * `scriptVariables` (`pulumi.Input[dict]`) - Mapping of query variable names to values (equivalent to the Spark SQL command: `SET name="value";`).

        The **status** object supports the following:

          * `details` (`pulumi.Input[str]`)
          * `state` (`pulumi.Input[str]`)
          * `stateStartTime` (`pulumi.Input[str]`)
          * `substate` (`pulumi.Input[str]`)
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

