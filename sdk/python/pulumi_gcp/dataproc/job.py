# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Job(pulumi.CustomResource):
    """
    Manages a job resource within a Dataproc cluster within GCE. For more information see
    [the official dataproc documentation](https://cloud.google.com/dataproc/).
    
    !> **Note:** This resource does not support 'update' and changing any attributes will cause the resource to be recreated.
    """
    def __init__(__self__, __name__, __opts__=None, force_delete=None, hadoop_config=None, hive_config=None, labels=None, pig_config=None, placement=None, project=None, pyspark_config=None, reference=None, region=None, scheduling=None, spark_config=None, sparksql_config=None):
        """Create a Job resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if force_delete and not isinstance(force_delete, bool):
            raise TypeError('Expected property force_delete to be a bool')
        __self__.force_delete = force_delete
        """
        By default, you can only delete inactive jobs within
        Dataproc. Setting this to true, and calling destroy, will ensure that the
        job is first cancelled before issuing the delete.
        """
        __props__['forceDelete'] = force_delete

        if hadoop_config and not isinstance(hadoop_config, dict):
            raise TypeError('Expected property hadoop_config to be a dict')
        __self__.hadoop_config = hadoop_config
        __props__['hadoopConfig'] = hadoop_config

        if hive_config and not isinstance(hive_config, dict):
            raise TypeError('Expected property hive_config to be a dict')
        __self__.hive_config = hive_config
        __props__['hiveConfig'] = hive_config

        if labels and not isinstance(labels, dict):
            raise TypeError('Expected property labels to be a dict')
        __self__.labels = labels
        """
        The list of labels (key/value pairs) to add to the job.
        """
        __props__['labels'] = labels

        if pig_config and not isinstance(pig_config, dict):
            raise TypeError('Expected property pig_config to be a dict')
        __self__.pig_config = pig_config
        __props__['pigConfig'] = pig_config

        if not placement:
            raise TypeError('Missing required property placement')
        elif not isinstance(placement, dict):
            raise TypeError('Expected property placement to be a dict')
        __self__.placement = placement
        __props__['placement'] = placement

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The project in which the `cluster` can be found and jobs
        subsequently run against. If it is not provided, the provider project is used.
        """
        __props__['project'] = project

        if pyspark_config and not isinstance(pyspark_config, dict):
            raise TypeError('Expected property pyspark_config to be a dict')
        __self__.pyspark_config = pyspark_config
        __props__['pysparkConfig'] = pyspark_config

        if reference and not isinstance(reference, dict):
            raise TypeError('Expected property reference to be a dict')
        __self__.reference = reference
        __props__['reference'] = reference

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The Cloud Dataproc region. This essentially determines which clusters are available
        for this job to be submitted to. If not specified, defaults to `global`.
        """
        __props__['region'] = region

        if scheduling and not isinstance(scheduling, dict):
            raise TypeError('Expected property scheduling to be a dict')
        __self__.scheduling = scheduling
        """
        Optional. Job scheduling configuration.
        """
        __props__['scheduling'] = scheduling

        if spark_config and not isinstance(spark_config, dict):
            raise TypeError('Expected property spark_config to be a dict')
        __self__.spark_config = spark_config
        __props__['sparkConfig'] = spark_config

        if sparksql_config and not isinstance(sparksql_config, dict):
            raise TypeError('Expected property sparksql_config to be a dict')
        __self__.sparksql_config = sparksql_config
        __props__['sparksqlConfig'] = sparksql_config

        __self__.driver_controls_files_uri = pulumi.runtime.UNKNOWN
        """
        If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        """
        __self__.driver_output_resource_uri = pulumi.runtime.UNKNOWN
        """
        A URI pointing to the location of the stdout of the job's driver program.
        """
        __self__.status = pulumi.runtime.UNKNOWN

        super(Job, __self__).__init__(
            'gcp:dataproc/job:Job',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'driverControlsFilesUri' in outs:
            self.driver_controls_files_uri = outs['driverControlsFilesUri']
        if 'driverOutputResourceUri' in outs:
            self.driver_output_resource_uri = outs['driverOutputResourceUri']
        if 'forceDelete' in outs:
            self.force_delete = outs['forceDelete']
        if 'hadoopConfig' in outs:
            self.hadoop_config = outs['hadoopConfig']
        if 'hiveConfig' in outs:
            self.hive_config = outs['hiveConfig']
        if 'labels' in outs:
            self.labels = outs['labels']
        if 'pigConfig' in outs:
            self.pig_config = outs['pigConfig']
        if 'placement' in outs:
            self.placement = outs['placement']
        if 'project' in outs:
            self.project = outs['project']
        if 'pysparkConfig' in outs:
            self.pyspark_config = outs['pysparkConfig']
        if 'reference' in outs:
            self.reference = outs['reference']
        if 'region' in outs:
            self.region = outs['region']
        if 'scheduling' in outs:
            self.scheduling = outs['scheduling']
        if 'sparkConfig' in outs:
            self.spark_config = outs['sparkConfig']
        if 'sparksqlConfig' in outs:
            self.sparksql_config = outs['sparksqlConfig']
        if 'status' in outs:
            self.status = outs['status']
