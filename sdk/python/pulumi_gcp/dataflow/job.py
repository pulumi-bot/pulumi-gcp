# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Job']


class Job(pulumi.CustomResource):
    additional_experiments: pulumi.Output[Optional[List[str]]] = pulumi.output_property("additionalExperiments")
    """
    List of experiments that should be used by the job. An example value is `["enable_stackdriver_agent_metrics"]`.
    """
    ip_configuration: pulumi.Output[Optional[str]] = pulumi.output_property("ipConfiguration")
    """
    The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
    """
    job_id: pulumi.Output[str] = pulumi.output_property("jobId")
    """
    The unique ID of this job.
    """
    labels: pulumi.Output[Optional[Dict[str, Any]]] = pulumi.output_property("labels")
    """
    User labels to be specified for the job. Keys and values should follow the restrictions
    specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
    **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
    Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
    """
    machine_type: pulumi.Output[Optional[str]] = pulumi.output_property("machineType")
    """
    The machine type to use for the job.
    """
    max_workers: pulumi.Output[Optional[float]] = pulumi.output_property("maxWorkers")
    """
    The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    A unique name for the resource, required by Dataflow.
    """
    network: pulumi.Output[Optional[str]] = pulumi.output_property("network")
    """
    The network to which VMs will be assigned. If it is not provided, "default" will be used.
    """
    on_delete: pulumi.Output[Optional[str]] = pulumi.output_property("onDelete")
    """
    One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
    """
    parameters: pulumi.Output[Optional[Dict[str, Any]]] = pulumi.output_property("parameters")
    """
    Key/Value pairs to be passed to the Dataflow job (as used in the template).
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The project in which the resource belongs. If it is not provided, the provider project is used.
    """
    region: pulumi.Output[Optional[str]] = pulumi.output_property("region")
    """
    The region in which the created job should run.
    """
    service_account_email: pulumi.Output[Optional[str]] = pulumi.output_property("serviceAccountEmail")
    """
    The Service Account email used to create the job.
    """
    state: pulumi.Output[str] = pulumi.output_property("state")
    """
    The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
    """
    subnetwork: pulumi.Output[Optional[str]] = pulumi.output_property("subnetwork")
    """
    The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
    """
    temp_gcs_location: pulumi.Output[str] = pulumi.output_property("tempGcsLocation")
    """
    A writeable location on GCS for the Dataflow job to dump its temporary data.
    """
    template_gcs_path: pulumi.Output[str] = pulumi.output_property("templateGcsPath")
    """
    The GCS path to the Dataflow job template.
    """
    type: pulumi.Output[str] = pulumi.output_property("type")
    """
    The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
    """
    zone: pulumi.Output[Optional[str]] = pulumi.output_property("zone")
    """
    The zone in which the created job should run. If it is not provided, the provider zone is used.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, additional_experiments: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, ip_configuration: Optional[pulumi.Input[str]] = None, labels: Optional[pulumi.Input[Dict[str, Any]]] = None, machine_type: Optional[pulumi.Input[str]] = None, max_workers: Optional[pulumi.Input[float]] = None, name: Optional[pulumi.Input[str]] = None, network: Optional[pulumi.Input[str]] = None, on_delete: Optional[pulumi.Input[str]] = None, parameters: Optional[pulumi.Input[Dict[str, Any]]] = None, project: Optional[pulumi.Input[str]] = None, region: Optional[pulumi.Input[str]] = None, service_account_email: Optional[pulumi.Input[str]] = None, subnetwork: Optional[pulumi.Input[str]] = None, temp_gcs_location: Optional[pulumi.Input[str]] = None, template_gcs_path: Optional[pulumi.Input[str]] = None, zone: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
        the official documentation for
        [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        big_data_job = gcp.dataflow.Job("bigDataJob",
            parameters={
                "baz": "qux",
                "foo": "bar",
            },
            temp_gcs_location="gs://my-bucket/tmp_dir",
            template_gcs_path="gs://my-bucket/templates/template_file")
        ```
        ## Note on "destroy" / "apply"

        There are many types of Dataflow jobs.  Some Dataflow jobs run constantly, getting new data from (e.g.) a GCS bucket, and outputting data continuously.  Some jobs process a set amount of data then terminate.  All jobs can fail while running due to programming errors or other issues.  In this way, Dataflow jobs are different from most other Google resources.

        The Dataflow resource is considered 'existing' while it is in a nonterminal state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE', 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for jobs which run continuously, but may surprise users who use this resource for other kinds of Dataflow jobs.

        A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If "cancelled", the job terminates - any data written remains where it is, but no new data will be processed.  If "drained", no new data will enter the pipeline, but any data currently in the pipeline will finish being processed.  The default is "cancelled", but if a user sets `on_delete` to `"drain"` in the configuration, you may experience a long wait for your `pulumi destroy` to complete.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] additional_experiments: List of experiments that should be used by the job. An example value is `["enable_stackdriver_agent_metrics"]`.
        :param pulumi.Input[str] ip_configuration: The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
        :param pulumi.Input[Dict[str, Any]] labels: User labels to be specified for the job. Keys and values should follow the restrictions
               specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
               **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
               Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
        :param pulumi.Input[str] machine_type: The machine type to use for the job.
        :param pulumi.Input[float] max_workers: The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] network: The network to which VMs will be assigned. If it is not provided, "default" will be used.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
        :param pulumi.Input[Dict[str, Any]] parameters: Key/Value pairs to be passed to the Dataflow job (as used in the template).
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created job should run.
        :param pulumi.Input[str] service_account_email: The Service Account email used to create the job.
        :param pulumi.Input[str] subnetwork: The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
        :param pulumi.Input[str] temp_gcs_location: A writeable location on GCS for the Dataflow job to dump its temporary data.
        :param pulumi.Input[str] template_gcs_path: The GCS path to the Dataflow job template.
        :param pulumi.Input[str] zone: The zone in which the created job should run. If it is not provided, the provider zone is used.
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

            __props__['additional_experiments'] = additional_experiments
            __props__['ip_configuration'] = ip_configuration
            __props__['labels'] = labels
            __props__['machine_type'] = machine_type
            __props__['max_workers'] = max_workers
            __props__['name'] = name
            __props__['network'] = network
            __props__['on_delete'] = on_delete
            __props__['parameters'] = parameters
            __props__['project'] = project
            __props__['region'] = region
            __props__['service_account_email'] = service_account_email
            __props__['subnetwork'] = subnetwork
            if temp_gcs_location is None:
                raise TypeError("Missing required property 'temp_gcs_location'")
            __props__['temp_gcs_location'] = temp_gcs_location
            if template_gcs_path is None:
                raise TypeError("Missing required property 'template_gcs_path'")
            __props__['template_gcs_path'] = template_gcs_path
            __props__['zone'] = zone
            __props__['job_id'] = None
            __props__['state'] = None
            __props__['type'] = None
        super(Job, __self__).__init__(
            'gcp:dataflow/job:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, additional_experiments: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, ip_configuration: Optional[pulumi.Input[str]] = None, job_id: Optional[pulumi.Input[str]] = None, labels: Optional[pulumi.Input[Dict[str, Any]]] = None, machine_type: Optional[pulumi.Input[str]] = None, max_workers: Optional[pulumi.Input[float]] = None, name: Optional[pulumi.Input[str]] = None, network: Optional[pulumi.Input[str]] = None, on_delete: Optional[pulumi.Input[str]] = None, parameters: Optional[pulumi.Input[Dict[str, Any]]] = None, project: Optional[pulumi.Input[str]] = None, region: Optional[pulumi.Input[str]] = None, service_account_email: Optional[pulumi.Input[str]] = None, state: Optional[pulumi.Input[str]] = None, subnetwork: Optional[pulumi.Input[str]] = None, temp_gcs_location: Optional[pulumi.Input[str]] = None, template_gcs_path: Optional[pulumi.Input[str]] = None, type: Optional[pulumi.Input[str]] = None, zone: Optional[pulumi.Input[str]] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] additional_experiments: List of experiments that should be used by the job. An example value is `["enable_stackdriver_agent_metrics"]`.
        :param pulumi.Input[str] ip_configuration: The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
        :param pulumi.Input[str] job_id: The unique ID of this job.
        :param pulumi.Input[Dict[str, Any]] labels: User labels to be specified for the job. Keys and values should follow the restrictions
               specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
               **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
               Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
        :param pulumi.Input[str] machine_type: The machine type to use for the job.
        :param pulumi.Input[float] max_workers: The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] network: The network to which VMs will be assigned. If it is not provided, "default" will be used.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
        :param pulumi.Input[Dict[str, Any]] parameters: Key/Value pairs to be passed to the Dataflow job (as used in the template).
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created job should run.
        :param pulumi.Input[str] service_account_email: The Service Account email used to create the job.
        :param pulumi.Input[str] state: The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
        :param pulumi.Input[str] subnetwork: The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
        :param pulumi.Input[str] temp_gcs_location: A writeable location on GCS for the Dataflow job to dump its temporary data.
        :param pulumi.Input[str] template_gcs_path: The GCS path to the Dataflow job template.
        :param pulumi.Input[str] type: The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
        :param pulumi.Input[str] zone: The zone in which the created job should run. If it is not provided, the provider zone is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_experiments"] = additional_experiments
        __props__["ip_configuration"] = ip_configuration
        __props__["job_id"] = job_id
        __props__["labels"] = labels
        __props__["machine_type"] = machine_type
        __props__["max_workers"] = max_workers
        __props__["name"] = name
        __props__["network"] = network
        __props__["on_delete"] = on_delete
        __props__["parameters"] = parameters
        __props__["project"] = project
        __props__["region"] = region
        __props__["service_account_email"] = service_account_email
        __props__["state"] = state
        __props__["subnetwork"] = subnetwork
        __props__["temp_gcs_location"] = temp_gcs_location
        __props__["template_gcs_path"] = template_gcs_path
        __props__["type"] = type
        __props__["zone"] = zone
        return Job(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

