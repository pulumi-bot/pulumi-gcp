# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Bucket(pulumi.CustomResource):
    bucket_policy_only: pulumi.Output[bool]
    """
    Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
    """
    cors: pulumi.Output[list]
    """
    The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
    """
    encryption: pulumi.Output[dict]
    """
    The bucket's encryption configuration.
    """
    force_destroy: pulumi.Output[bool]
    """
    When deleting a bucket, this
    boolean option will delete all contained objects. If you try to delete a
    bucket that contains objects, this provider will fail that run.
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to the bucket.
    """
    lifecycle_rules: pulumi.Output[list]
    """
    The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
    """
    location: pulumi.Output[str]
    """
    The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
    """
    logging: pulumi.Output[dict]
    """
    The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
    """
    name: pulumi.Output[str]
    """
    The name of the bucket.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    requester_pays: pulumi.Output[bool]
    """
    Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
    """
    retention_policy: pulumi.Output[dict]
    """
    Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    storage_class: pulumi.Output[str]
    """
    The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
    """
    url: pulumi.Output[str]
    """
    The base URL of the bucket, in the format `gs://<bucket-name>`.
    """
    versioning: pulumi.Output[dict]
    """
    The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
    """
    website: pulumi.Output[dict]
    """
    Configuration if the bucket acts as a website. Structure is documented below.
    """
    def __init__(__self__, resource_name, opts=None, bucket_policy_only=None, cors=None, encryption=None, force_destroy=None, labels=None, lifecycle_rules=None, location=None, logging=None, name=None, project=None, requester_pays=None, retention_policy=None, storage_class=None, versioning=None, website=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a new bucket in Google cloud storage service (GCS).
        Once a bucket has been created, its location can't be changed.
        [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied
        using the [`google_storage_bucket_acl` resource](https://www.terraform.io/docs/providers/google/r/storage_bucket_acl.html).
        
        For more information see
        [the official documentation](https://cloud.google.com/storage/docs/overview)
        and
        [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
        
        **Note**: If the project id is not set on the resource or in the provider block it will be dynamically
        determined which will require enabling the compute api.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bucket_policy_only: Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
        :param pulumi.Input[list] cors: The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[dict] encryption: The bucket's encryption configuration.
        :param pulumi.Input[bool] force_destroy: When deleting a bucket, this
               boolean option will delete all contained objects. If you try to delete a
               bucket that contains objects, this provider will fail that run.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the bucket.
        :param pulumi.Input[list] lifecycle_rules: The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[str] location: The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        :param pulumi.Input[dict] logging: The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        :param pulumi.Input[str] name: The name of the bucket.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[bool] requester_pays: Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        :param pulumi.Input[dict] retention_policy: Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        :param pulumi.Input[str] storage_class: The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        :param pulumi.Input[dict] versioning: The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        :param pulumi.Input[dict] website: Configuration if the bucket acts as a website. Structure is documented below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_bucket.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            __props__['bucket_policy_only'] = bucket_policy_only
            __props__['cors'] = cors
            __props__['encryption'] = encryption
            __props__['force_destroy'] = force_destroy
            __props__['labels'] = labels
            __props__['lifecycle_rules'] = lifecycle_rules
            __props__['location'] = location
            __props__['logging'] = logging
            __props__['name'] = name
            __props__['project'] = project
            __props__['requester_pays'] = requester_pays
            __props__['retention_policy'] = retention_policy
            __props__['storage_class'] = storage_class
            __props__['versioning'] = versioning
            __props__['website'] = website
            __props__['self_link'] = None
            __props__['url'] = None
        super(Bucket, __self__).__init__(
            'gcp:storage/bucket:Bucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bucket_policy_only=None, cors=None, encryption=None, force_destroy=None, labels=None, lifecycle_rules=None, location=None, logging=None, name=None, project=None, requester_pays=None, retention_policy=None, self_link=None, storage_class=None, url=None, versioning=None, website=None):
        """
        Get an existing Bucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bucket_policy_only: Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
        :param pulumi.Input[list] cors: The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[dict] encryption: The bucket's encryption configuration.
        :param pulumi.Input[bool] force_destroy: When deleting a bucket, this
               boolean option will delete all contained objects. If you try to delete a
               bucket that contains objects, this provider will fail that run.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the bucket.
        :param pulumi.Input[list] lifecycle_rules: The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[str] location: The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        :param pulumi.Input[dict] logging: The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        :param pulumi.Input[str] name: The name of the bucket.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[bool] requester_pays: Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        :param pulumi.Input[dict] retention_policy: Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] storage_class: The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        :param pulumi.Input[str] url: The base URL of the bucket, in the format `gs://<bucket-name>`.
        :param pulumi.Input[dict] versioning: The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        :param pulumi.Input[dict] website: Configuration if the bucket acts as a website. Structure is documented below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_bucket.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["bucket_policy_only"] = bucket_policy_only
        __props__["cors"] = cors
        __props__["encryption"] = encryption
        __props__["force_destroy"] = force_destroy
        __props__["labels"] = labels
        __props__["lifecycle_rules"] = lifecycle_rules
        __props__["location"] = location
        __props__["logging"] = logging
        __props__["name"] = name
        __props__["project"] = project
        __props__["requester_pays"] = requester_pays
        __props__["retention_policy"] = retention_policy
        __props__["self_link"] = self_link
        __props__["storage_class"] = storage_class
        __props__["url"] = url
        __props__["versioning"] = versioning
        __props__["website"] = website
        return Bucket(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

