# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['UsageExportBucket']


class UsageExportBucket(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows creation and management of a Google Cloud Platform project.

        Projects created with this resource must be associated with an Organization.
        See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.

        The service account used to run this provider when creating a `organizations.Project`
        resource must have `roles/resourcemanager.projectCreator`. See the
        [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
        doc for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            org_id="1234567",
            project_id="your-project-id")
        ```

        To create a project under a specific folder

        ```python
        import pulumi
        import pulumi_gcp as gcp

        department1 = gcp.organizations.Folder("department1",
            display_name="Department 1",
            parent="organizations/1234567")
        my_project_in_a_folder = gcp.organizations.Project("myProject-in-a-folder",
            project_id="your-project-id",
            folder_id=department1.name)
        ```

        ## Import

        Projects can be imported using the `project_id`, e.g.

        ```sh
         $ pulumi import gcp:projects/usageExportBucket:UsageExportBucket my_project your-project-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: The bucket to store reports in.
        :param pulumi.Input[str] prefix: A prefix for the reports, for instance, the project name.
        :param pulumi.Input[str] project: The project to set the export bucket on. If it is not provided, the provider project is used.
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

            if bucket_name is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_name'")
            __props__['bucket_name'] = bucket_name
            __props__['prefix'] = prefix
            __props__['project'] = project
        super(UsageExportBucket, __self__).__init__(
            'gcp:projects/usageExportBucket:UsageExportBucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_name: Optional[pulumi.Input[str]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'UsageExportBucket':
        """
        Get an existing UsageExportBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: The bucket to store reports in.
        :param pulumi.Input[str] prefix: A prefix for the reports, for instance, the project name.
        :param pulumi.Input[str] project: The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket_name"] = bucket_name
        __props__["prefix"] = prefix
        __props__["project"] = project
        return UsageExportBucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Output[str]:
        """
        The bucket to store reports in.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[Optional[str]]:
        """
        A prefix for the reports, for instance, the project name.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

