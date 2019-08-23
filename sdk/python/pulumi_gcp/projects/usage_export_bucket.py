# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UsageExportBucket(pulumi.CustomResource):
    bucket_name: pulumi.Output[str]
    prefix: pulumi.Output[str]
    project: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, bucket_name=None, prefix=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows creation and management of a Google Cloud Platform project.
        
        Projects created with this resource must be associated with an Organization.
        See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
        
        The service account used to run this provider when creating a `organizations.Project`
        resource must have `roles/resourcemanager.projectCreator`. See the
        [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
        doc for more information.
        
        Note that prior to 0.8.5, `organizations.Project` functioned like a data source,
        meaning any project referenced by it had to be created and managed outside
        this provider. As of 0.8.5, `organizations.Project` functions like any other
        resource, with this provider creating and managing the project. To replicate the old
        behavior, either:
        
        * Use the project ID directly in whatever is referencing the project, using the
          [projects.IAMPolicy](https://www.terraform.io/docs/providers/google/r/google_project_iam.html)
          to replace the old `policy_data` property.
        * Use the [import](https://www.terraform.io/docs/import/usage.html) functionality
          to import your pre-existing project into this provider, where it can be referenced and
          used just like always, keeping in mind that this provider will attempt to undo any changes
          made outside this provider.
        
        > It's important to note that any project resources that were added to your config
        prior to 0.8.5 will continue to function as they always have, and will not be managed by
        this provider. Only newly added projects are affected.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project_usage_export_bucket.html.markdown.
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

            if bucket_name is None:
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
    def get(resource_name, id, opts=None, bucket_name=None, prefix=None, project=None):
        """
        Get an existing UsageExportBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project_usage_export_bucket.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["bucket_name"] = bucket_name
        __props__["prefix"] = prefix
        __props__["project"] = project
        return UsageExportBucket(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

