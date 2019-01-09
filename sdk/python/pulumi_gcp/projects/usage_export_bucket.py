# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class UsageExportBucket(pulumi.CustomResource):
    """
    Allows creation and management of a Google Cloud Platform project.
    
    Projects created with this resource must be associated with an Organization.
    See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
    
    The service account used to run Terraform when creating a `google_project`
    resource must have `roles/resourcemanager.projectCreator`. See the
    [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
    doc for more information.
    
    Note that prior to 0.8.5, `google_project` functioned like a data source,
    meaning any project referenced by it had to be created and managed outside
    Terraform. As of 0.8.5, `google_project` functions like any other Terraform
    resource, with Terraform creating and managing the project. To replicate the old
    behavior, either:
    
    * Use the project ID directly in whatever is referencing the project, using the
      [google_project_iam_policy](https://www.terraform.io/docs/providers/google/r/google_project_iam.html)
      to replace the old `policy_data` property.
    * Use the [import](https://www.terraform.io/docs/import/usage.html) functionality
      to import your pre-existing project into Terraform, where it can be referenced and
      used just like always, keeping in mind that Terraform will attempt to undo any changes
      made outside Terraform.
    
    > It's important to note that any project resources that were added to your Terraform config
    prior to 0.8.5 will continue to function as they always have, and will not be managed by
    Terraform. Only newly added projects are affected.
    """
    def __init__(__self__, __name__, __opts__=None, bucket_name=None, prefix=None, project=None):
        """Create a UsageExportBucket resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not bucket_name:
            raise TypeError('Missing required property bucket_name')
        __props__['bucket_name'] = bucket_name

        __props__['prefix'] = prefix

        __props__['project'] = project

        super(UsageExportBucket, __self__).__init__(
            'gcp:projects/usageExportBucket:UsageExportBucket',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

