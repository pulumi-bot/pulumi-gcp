# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FlexTemplateJobArgs', 'FlexTemplateJob']

@pulumi.input_type
class FlexTemplateJobArgs:
    def __init__(__self__, *,
                 container_spec_gcs_path: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 on_delete: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FlexTemplateJob resource.
        :param pulumi.Input[str] container_spec_gcs_path: The GCS path to the Dataflow job Flex
               Template.
        :param pulumi.Input[Mapping[str, Any]] labels: User labels to be specified for the job. Keys and values
               should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
               page. **Note**: This field is marked as deprecated as the API does not currently
               support adding labels.
               **NOTE**: Google-provided Dataflow templates often provide default labels
               that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
               labels will be ignored to prevent diffs on re-apply.
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel". Specifies behavior of
               deletion during `pulumi destroy`.  See above note.
        :param pulumi.Input[Mapping[str, Any]] parameters: Key/Value pairs to be passed to the Dataflow job (as
               used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
               such as `serviceAccount`, `workerMachineType`, etc can be specified here.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not
               provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created job should run.
        """
        pulumi.set(__self__, "container_spec_gcs_path", container_spec_gcs_path)
        if labels is not None:
            warnings.warn("""Deprecated until the API supports this field""", DeprecationWarning)
            pulumi.log.warn("""labels is deprecated: Deprecated until the API supports this field""")
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if on_delete is not None:
            pulumi.set(__self__, "on_delete", on_delete)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="containerSpecGcsPath")
    def container_spec_gcs_path(self) -> pulumi.Input[str]:
        """
        The GCS path to the Dataflow job Flex
        Template.
        """
        return pulumi.get(self, "container_spec_gcs_path")

    @container_spec_gcs_path.setter
    def container_spec_gcs_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_spec_gcs_path", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        User labels to be specified for the job. Keys and values
        should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
        page. **Note**: This field is marked as deprecated as the API does not currently
        support adding labels.
        **NOTE**: Google-provided Dataflow templates often provide default labels
        that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
        labels will be ignored to prevent diffs on re-apply.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the resource, required by Dataflow.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="onDelete")
    def on_delete(self) -> Optional[pulumi.Input[str]]:
        """
        One of "drain" or "cancel". Specifies behavior of
        deletion during `pulumi destroy`.  See above note.
        """
        return pulumi.get(self, "on_delete")

    @on_delete.setter
    def on_delete(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "on_delete", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key/Value pairs to be passed to the Dataflow job (as
        used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
        such as `serviceAccount`, `workerMachineType`, etc can be specified here.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the resource belongs. If it is not
        provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which the created job should run.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _FlexTemplateJobState:
    def __init__(__self__, *,
                 container_spec_gcs_path: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 on_delete: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FlexTemplateJob resources.
        :param pulumi.Input[str] container_spec_gcs_path: The GCS path to the Dataflow job Flex
               Template.
        :param pulumi.Input[str] job_id: The unique ID of this job.
        :param pulumi.Input[Mapping[str, Any]] labels: User labels to be specified for the job. Keys and values
               should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
               page. **Note**: This field is marked as deprecated as the API does not currently
               support adding labels.
               **NOTE**: Google-provided Dataflow templates often provide default labels
               that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
               labels will be ignored to prevent diffs on re-apply.
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel". Specifies behavior of
               deletion during `pulumi destroy`.  See above note.
        :param pulumi.Input[Mapping[str, Any]] parameters: Key/Value pairs to be passed to the Dataflow job (as
               used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
               such as `serviceAccount`, `workerMachineType`, etc can be specified here.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not
               provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created job should run.
        :param pulumi.Input[str] state: The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
        """
        if container_spec_gcs_path is not None:
            pulumi.set(__self__, "container_spec_gcs_path", container_spec_gcs_path)
        if job_id is not None:
            pulumi.set(__self__, "job_id", job_id)
        if labels is not None:
            warnings.warn("""Deprecated until the API supports this field""", DeprecationWarning)
            pulumi.log.warn("""labels is deprecated: Deprecated until the API supports this field""")
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if on_delete is not None:
            pulumi.set(__self__, "on_delete", on_delete)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="containerSpecGcsPath")
    def container_spec_gcs_path(self) -> Optional[pulumi.Input[str]]:
        """
        The GCS path to the Dataflow job Flex
        Template.
        """
        return pulumi.get(self, "container_spec_gcs_path")

    @container_spec_gcs_path.setter
    def container_spec_gcs_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_spec_gcs_path", value)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID of this job.
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        User labels to be specified for the job. Keys and values
        should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
        page. **Note**: This field is marked as deprecated as the API does not currently
        support adding labels.
        **NOTE**: Google-provided Dataflow templates often provide default labels
        that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
        labels will be ignored to prevent diffs on re-apply.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the resource, required by Dataflow.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="onDelete")
    def on_delete(self) -> Optional[pulumi.Input[str]]:
        """
        One of "drain" or "cancel". Specifies behavior of
        deletion during `pulumi destroy`.  See above note.
        """
        return pulumi.get(self, "on_delete")

    @on_delete.setter
    def on_delete(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "on_delete", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key/Value pairs to be passed to the Dataflow job (as
        used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
        such as `serviceAccount`, `workerMachineType`, etc can be specified here.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the resource belongs. If it is not
        provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which the created job should run.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class FlexTemplateJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_spec_gcs_path: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 on_delete: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a [Flex Template](https://cloud.google.com/dataflow/docs/guides/templates/using-flex-templates)
        job on Dataflow, which is an implementation of Apache Beam running on Google
        Compute Engine. For more information see the official documentation for [Beam](https://beam.apache.org)
        and [Dataflow](https://cloud.google.com/dataflow/).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        big_data_job = gcp.dataflow.FlexTemplateJob("bigDataJob",
            container_spec_gcs_path="gs://my-bucket/templates/template.json",
            parameters={
                "inputSubscription": "messages",
            },
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ## Note on "destroy" / "apply"

        There are many types of Dataflow jobs.  Some Dataflow jobs run constantly,
        getting new data from (e.g.) a GCS bucket, and outputting data continuously.
        Some jobs process a set amount of data then terminate. All jobs can fail while
        running due to programming errors or other issues. In this way, Dataflow jobs
        are different from most other provider / Google resources.

        The Dataflow resource is considered 'existing' while it is in a nonterminal
        state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE',
        'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for
        jobs which run continuously, but may surprise users who use this resource for
        other kinds of Dataflow jobs.

        A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If
        "cancelled", the job terminates - any data written remains where it is, but no
        new data will be processed.  If "drained", no new data will enter the pipeline,
        but any data currently in the pipeline will finish being processed.  The default
        is "cancelled", but if a user sets `on_delete` to `"drain"` in the
        configuration, you may experience a long wait for your `pulumi destroy` to
        complete.

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_spec_gcs_path: The GCS path to the Dataflow job Flex
               Template.
        :param pulumi.Input[Mapping[str, Any]] labels: User labels to be specified for the job. Keys and values
               should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
               page. **Note**: This field is marked as deprecated as the API does not currently
               support adding labels.
               **NOTE**: Google-provided Dataflow templates often provide default labels
               that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
               labels will be ignored to prevent diffs on re-apply.
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel". Specifies behavior of
               deletion during `pulumi destroy`.  See above note.
        :param pulumi.Input[Mapping[str, Any]] parameters: Key/Value pairs to be passed to the Dataflow job (as
               used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
               such as `serviceAccount`, `workerMachineType`, etc can be specified here.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not
               provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created job should run.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlexTemplateJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a [Flex Template](https://cloud.google.com/dataflow/docs/guides/templates/using-flex-templates)
        job on Dataflow, which is an implementation of Apache Beam running on Google
        Compute Engine. For more information see the official documentation for [Beam](https://beam.apache.org)
        and [Dataflow](https://cloud.google.com/dataflow/).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        big_data_job = gcp.dataflow.FlexTemplateJob("bigDataJob",
            container_spec_gcs_path="gs://my-bucket/templates/template.json",
            parameters={
                "inputSubscription": "messages",
            },
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ## Note on "destroy" / "apply"

        There are many types of Dataflow jobs.  Some Dataflow jobs run constantly,
        getting new data from (e.g.) a GCS bucket, and outputting data continuously.
        Some jobs process a set amount of data then terminate. All jobs can fail while
        running due to programming errors or other issues. In this way, Dataflow jobs
        are different from most other provider / Google resources.

        The Dataflow resource is considered 'existing' while it is in a nonterminal
        state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE',
        'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for
        jobs which run continuously, but may surprise users who use this resource for
        other kinds of Dataflow jobs.

        A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If
        "cancelled", the job terminates - any data written remains where it is, but no
        new data will be processed.  If "drained", no new data will enter the pipeline,
        but any data currently in the pipeline will finish being processed.  The default
        is "cancelled", but if a user sets `on_delete` to `"drain"` in the
        configuration, you may experience a long wait for your `pulumi destroy` to
        complete.

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param FlexTemplateJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlexTemplateJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_spec_gcs_path: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 on_delete: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = FlexTemplateJobArgs.__new__(FlexTemplateJobArgs)

            if container_spec_gcs_path is None and not opts.urn:
                raise TypeError("Missing required property 'container_spec_gcs_path'")
            __props__.__dict__['container_spec_gcs_path'] = container_spec_gcs_path
            if labels is not None and not opts.urn:
                warnings.warn("""Deprecated until the API supports this field""", DeprecationWarning)
                pulumi.log.warn("""labels is deprecated: Deprecated until the API supports this field""")
            __props__.__dict__['labels'] = labels
            __props__.__dict__['name'] = name
            __props__.__dict__['on_delete'] = on_delete
            __props__.__dict__['parameters'] = parameters
            __props__.__dict__['project'] = project
            __props__.__dict__['region'] = region
            __props__.__dict__['job_id'] = None
            __props__.__dict__['state'] = None
        super(FlexTemplateJob, __self__).__init__(
            'gcp:dataflow/flexTemplateJob:FlexTemplateJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_spec_gcs_path: Optional[pulumi.Input[str]] = None,
            job_id: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            on_delete: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'FlexTemplateJob':
        """
        Get an existing FlexTemplateJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_spec_gcs_path: The GCS path to the Dataflow job Flex
               Template.
        :param pulumi.Input[str] job_id: The unique ID of this job.
        :param pulumi.Input[Mapping[str, Any]] labels: User labels to be specified for the job. Keys and values
               should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
               page. **Note**: This field is marked as deprecated as the API does not currently
               support adding labels.
               **NOTE**: Google-provided Dataflow templates often provide default labels
               that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
               labels will be ignored to prevent diffs on re-apply.
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel". Specifies behavior of
               deletion during `pulumi destroy`.  See above note.
        :param pulumi.Input[Mapping[str, Any]] parameters: Key/Value pairs to be passed to the Dataflow job (as
               used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
               such as `serviceAccount`, `workerMachineType`, etc can be specified here.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not
               provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created job should run.
        :param pulumi.Input[str] state: The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlexTemplateJobState.__new__(_FlexTemplateJobState)

        __props__.__dict__['container_spec_gcs_path'] = container_spec_gcs_path
        __props__.__dict__['job_id'] = job_id
        __props__.__dict__['labels'] = labels
        __props__.__dict__['name'] = name
        __props__.__dict__['on_delete'] = on_delete
        __props__.__dict__['parameters'] = parameters
        __props__.__dict__['project'] = project
        __props__.__dict__['region'] = region
        __props__.__dict__['state'] = state
        return FlexTemplateJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerSpecGcsPath")
    def container_spec_gcs_path(self) -> pulumi.Output[str]:
        """
        The GCS path to the Dataflow job Flex
        Template.
        """
        return pulumi.get(self, "container_spec_gcs_path")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Output[str]:
        """
        The unique ID of this job.
        """
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        User labels to be specified for the job. Keys and values
        should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
        page. **Note**: This field is marked as deprecated as the API does not currently
        support adding labels.
        **NOTE**: Google-provided Dataflow templates often provide default labels
        that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
        labels will be ignored to prevent diffs on re-apply.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the resource, required by Dataflow.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="onDelete")
    def on_delete(self) -> pulumi.Output[Optional[str]]:
        """
        One of "drain" or "cancel". Specifies behavior of
        deletion during `pulumi destroy`.  See above note.
        """
        return pulumi.get(self, "on_delete")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Key/Value pairs to be passed to the Dataflow job (as
        used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
        such as `serviceAccount`, `workerMachineType`, etc can be specified here.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project in which the resource belongs. If it is not
        provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which the created job should run.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
        """
        return pulumi.get(self, "state")

