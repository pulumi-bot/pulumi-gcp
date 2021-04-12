# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ProjectBucketConfigArgs', 'ProjectBucketConfig']

@pulumi.input_type
class ProjectBucketConfigArgs:
    def __init__(__self__, *,
                 bucket_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ProjectBucketConfig resource.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[str] project: The parent resource that contains the logging bucket.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        pulumi.set(__self__, "bucket_id", bucket_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)

    @property
    @pulumi.getter(name="bucketId")
    def bucket_id(self) -> pulumi.Input[str]:
        """
        The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        """
        return pulumi.get(self, "bucket_id")

    @bucket_id.setter
    def bucket_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location of the bucket.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The parent resource that contains the logging bucket.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Describes this bucket.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)


class ProjectBucketConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a project-level logging bucket config. For more information see
        [the official logging documentation](https://cloud.google.com/logging/docs/) and
        [Storing Logs](https://cloud.google.com/logging/docs/storage).

        > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.organizations.Project("default",
            project_id="your-project-id",
            org_id="123456789")
        basic = gcp.logging.ProjectBucketConfig("basic",
            project=default.name,
            location="global",
            retention_days=30,
            bucket_id="_Default")
        ```

        Create logging bucket with customId

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.logging.ProjectBucketConfig("basic",
            bucket_id="custom-bucket",
            location="global",
            project="project_id",
            retention_days=30)
        ```

        ## Import

        This resource can be imported using the following format

        ```sh
         $ pulumi import gcp:logging/projectBucketConfig:ProjectBucketConfig default projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[str] project: The parent resource that contains the logging bucket.
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectBucketConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a project-level logging bucket config. For more information see
        [the official logging documentation](https://cloud.google.com/logging/docs/) and
        [Storing Logs](https://cloud.google.com/logging/docs/storage).

        > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.organizations.Project("default",
            project_id="your-project-id",
            org_id="123456789")
        basic = gcp.logging.ProjectBucketConfig("basic",
            project=default.name,
            location="global",
            retention_days=30,
            bucket_id="_Default")
        ```

        Create logging bucket with customId

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.logging.ProjectBucketConfig("basic",
            bucket_id="custom-bucket",
            location="global",
            project="project_id",
            retention_days=30)
        ```

        ## Import

        This resource can be imported using the following format

        ```sh
         $ pulumi import gcp:logging/projectBucketConfig:ProjectBucketConfig default projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
        ```

        :param str resource_name: The name of the resource.
        :param ProjectBucketConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectBucketConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
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
            __props__ = dict()

            if bucket_id is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_id'")
            __props__['bucket_id'] = bucket_id
            __props__['description'] = description
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            __props__['retention_days'] = retention_days
            __props__['lifecycle_state'] = None
            __props__['name'] = None
        super(ProjectBucketConfig, __self__).__init__(
            'gcp:logging/projectBucketConfig:ProjectBucketConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lifecycle_state: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            retention_days: Optional[pulumi.Input[int]] = None) -> 'ProjectBucketConfig':
        """
        Get an existing ProjectBucketConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] lifecycle_state: The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[str] name: The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
        :param pulumi.Input[str] project: The parent resource that contains the logging bucket.
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket_id"] = bucket_id
        __props__["description"] = description
        __props__["lifecycle_state"] = lifecycle_state
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["retention_days"] = retention_days
        return ProjectBucketConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bucketId")
    def bucket_id(self) -> pulumi.Output[str]:
        """
        The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        """
        return pulumi.get(self, "bucket_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Describes this bucket.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lifecycleState")
    def lifecycle_state(self) -> pulumi.Output[str]:
        """
        The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        """
        return pulumi.get(self, "lifecycle_state")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the bucket.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The parent resource that contains the logging bucket.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        """
        return pulumi.get(self, "retention_days")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

