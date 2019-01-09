# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Bucket(pulumi.CustomResource):
    """
    Creates a new bucket in Google cloud storage service (GCS).
    Once a bucket has been created, its location can't be changed.
    [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied
    using the [`google_storage_bucket_acl` resource](https://www.terraform.io/docs/providers/google/r/storage_bucket_acl.html).
    
    For more information see
    [the official documentation](https://cloud.google.com/storage/docs/overview)
    and
    [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
    
    """
    def __init__(__self__, __name__, __opts__=None, cors=None, encryption=None, force_destroy=None, labels=None, lifecycle_rules=None, location=None, logging=None, name=None, project=None, storage_class=None, versioning=None, websites=None):
        """Create a Bucket resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['cors'] = cors

        __props__['encryption'] = encryption

        __props__['force_destroy'] = force_destroy

        __props__['labels'] = labels

        __props__['lifecycle_rules'] = lifecycle_rules

        __props__['location'] = location

        __props__['logging'] = logging

        __props__['name'] = name

        __props__['project'] = project

        __props__['storage_class'] = storage_class

        __props__['versioning'] = versioning

        __props__['websites'] = websites

        __props__['self_link'] = None
        __props__['url'] = None

        super(Bucket, __self__).__init__(
            'gcp:storage/bucket:Bucket',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

