# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['FlexTemplateJob']


class FlexTemplateJob(pulumi.CustomResource):
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
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel". Specifies behavior of
               deletion during `pulumi destroy`.  See above note.
        :param pulumi.Input[Mapping[str, Any]] parameters: Key/Value pairs to be passed to the Dataflow job (as
               used in the template).
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not
               provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created job should run.
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

            if container_spec_gcs_path is None:
                raise TypeError("Missing required property 'container_spec_gcs_path'")
            __props__['container_spec_gcs_path'] = container_spec_gcs_path
            if labels is not None:
                warnings.warn("""Deprecated until the API supports this field""", DeprecationWarning)
                pulumi.log.warn("labels is deprecated: Deprecated until the API supports this field")
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['on_delete'] = on_delete
            __props__['parameters'] = parameters
            __props__['project'] = project
            __props__['region'] = region
            __props__['job_id'] = None
            __props__['state'] = None
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
        :param pulumi.Input[str] name: A unique name for the resource, required by Dataflow.
        :param pulumi.Input[str] on_delete: One of "drain" or "cancel". Specifies behavior of
               deletion during `pulumi destroy`.  See above note.
        :param pulumi.Input[Mapping[str, Any]] parameters: Key/Value pairs to be passed to the Dataflow job (as
               used in the template).
        :param pulumi.Input[str] project: The project in which the resource belongs. If it is not
               provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created job should run.
        :param pulumi.Input[str] state: The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["container_spec_gcs_path"] = container_spec_gcs_path
        __props__["job_id"] = job_id
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["on_delete"] = on_delete
        __props__["parameters"] = parameters
        __props__["project"] = project
        __props__["region"] = region
        __props__["state"] = state
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
        used in the template).
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

