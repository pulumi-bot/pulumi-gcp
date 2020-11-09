# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Bucket']


class Bucket(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_policy_only: Optional[pulumi.Input[bool]] = None,
                 cors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorArgs']]]]] = None,
                 default_event_based_hold: Optional[pulumi.Input[bool]] = None,
                 encryption: Optional[pulumi.Input[pulumi.InputType['BucketEncryptionArgs']]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lifecycle_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleRuleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 logging: Optional[pulumi.Input[pulumi.InputType['BucketLoggingArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 requester_pays: Optional[pulumi.Input[bool]] = None,
                 retention_policy: Optional[pulumi.Input[pulumi.InputType['BucketRetentionPolicyArgs']]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 uniform_bucket_level_access: Optional[pulumi.Input[bool]] = None,
                 versioning: Optional[pulumi.Input[pulumi.InputType['BucketVersioningArgs']]] = None,
                 website: Optional[pulumi.Input[pulumi.InputType['BucketWebsiteArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new bucket in Google cloud storage service (GCS).
        Once a bucket has been created, its location can't be changed.
        [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied
        using the [`storage.BucketACL`](https://www.terraform.io/docs/providers/google/r/storage_bucket_acl.html) resource.

        For more information see
        [the official documentation](https://cloud.google.com/storage/docs/overview)
        and
        [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).

        **Note**: If the project id is not set on the resource or in the provider block it will be dynamically
        determined which will require enabling the compute api.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bucket_policy_only: Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. This field will be removed in the next major release of the provider.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorArgs']]]] cors: The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[pulumi.InputType['BucketEncryptionArgs']] encryption: The bucket's encryption configuration.
        :param pulumi.Input[bool] force_destroy: When deleting a bucket, this
               boolean option will delete all contained objects. If you try to delete a
               bucket that contains objects, the provider will fail that run.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleRuleArgs']]]] lifecycle_rules: The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[str] location: The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        :param pulumi.Input[pulumi.InputType['BucketLoggingArgs']] logging: The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        :param pulumi.Input[str] name: The name of the bucket.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[bool] requester_pays: Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        :param pulumi.Input[pulumi.InputType['BucketRetentionPolicyArgs']] retention_policy: Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        :param pulumi.Input[str] storage_class: The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
        :param pulumi.Input[bool] uniform_bucket_level_access: Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
        :param pulumi.Input[pulumi.InputType['BucketVersioningArgs']] versioning: The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        :param pulumi.Input[pulumi.InputType['BucketWebsiteArgs']] website: Configuration if the bucket acts as a website. Structure is documented below.
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

            if bucket_policy_only is not None:
                warnings.warn("""Please use the uniform_bucket_level_access as this field has been renamed by Google.""", DeprecationWarning)
                pulumi.log.warn("bucket_policy_only is deprecated: Please use the uniform_bucket_level_access as this field has been renamed by Google.")
            __props__['bucket_policy_only'] = bucket_policy_only
            __props__['cors'] = cors
            __props__['default_event_based_hold'] = default_event_based_hold
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
            __props__['uniform_bucket_level_access'] = uniform_bucket_level_access
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_policy_only: Optional[pulumi.Input[bool]] = None,
            cors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorArgs']]]]] = None,
            default_event_based_hold: Optional[pulumi.Input[bool]] = None,
            encryption: Optional[pulumi.Input[pulumi.InputType['BucketEncryptionArgs']]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            lifecycle_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleRuleArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            logging: Optional[pulumi.Input[pulumi.InputType['BucketLoggingArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            requester_pays: Optional[pulumi.Input[bool]] = None,
            retention_policy: Optional[pulumi.Input[pulumi.InputType['BucketRetentionPolicyArgs']]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            storage_class: Optional[pulumi.Input[str]] = None,
            uniform_bucket_level_access: Optional[pulumi.Input[bool]] = None,
            url: Optional[pulumi.Input[str]] = None,
            versioning: Optional[pulumi.Input[pulumi.InputType['BucketVersioningArgs']]] = None,
            website: Optional[pulumi.Input[pulumi.InputType['BucketWebsiteArgs']]] = None) -> 'Bucket':
        """
        Get an existing Bucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bucket_policy_only: Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. This field will be removed in the next major release of the provider.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorArgs']]]] cors: The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[pulumi.InputType['BucketEncryptionArgs']] encryption: The bucket's encryption configuration.
        :param pulumi.Input[bool] force_destroy: When deleting a bucket, this
               boolean option will delete all contained objects. If you try to delete a
               bucket that contains objects, the provider will fail that run.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketLifecycleRuleArgs']]]] lifecycle_rules: The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[str] location: The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        :param pulumi.Input[pulumi.InputType['BucketLoggingArgs']] logging: The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        :param pulumi.Input[str] name: The name of the bucket.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[bool] requester_pays: Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        :param pulumi.Input[pulumi.InputType['BucketRetentionPolicyArgs']] retention_policy: Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] storage_class: The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
        :param pulumi.Input[bool] uniform_bucket_level_access: Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
        :param pulumi.Input[str] url: The base URL of the bucket, in the format `gs://<bucket-name>`.
        :param pulumi.Input[pulumi.InputType['BucketVersioningArgs']] versioning: The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        :param pulumi.Input[pulumi.InputType['BucketWebsiteArgs']] website: Configuration if the bucket acts as a website. Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket_policy_only"] = bucket_policy_only
        __props__["cors"] = cors
        __props__["default_event_based_hold"] = default_event_based_hold
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
        __props__["uniform_bucket_level_access"] = uniform_bucket_level_access
        __props__["url"] = url
        __props__["versioning"] = versioning
        __props__["website"] = website
        return Bucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bucketPolicyOnly")
    def bucket_policy_only(self) -> pulumi.Output[bool]:
        """
        Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. This field will be removed in the next major release of the provider.
        """
        return pulumi.get(self, "bucket_policy_only")

    @property
    @pulumi.getter
    def cors(self) -> pulumi.Output[Optional[Sequence['outputs.BucketCor']]]:
        """
        The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        """
        return pulumi.get(self, "cors")

    @property
    @pulumi.getter(name="defaultEventBasedHold")
    def default_event_based_hold(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "default_event_based_hold")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[Optional['outputs.BucketEncryption']]:
        """
        The bucket's encryption configuration.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        When deleting a bucket, this
        boolean option will delete all contained objects. If you try to delete a
        bucket that contains objects, the provider will fail that run.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to the bucket.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lifecycleRules")
    def lifecycle_rules(self) -> pulumi.Output[Optional[Sequence['outputs.BucketLifecycleRule']]]:
        """
        The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        """
        return pulumi.get(self, "lifecycle_rules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def logging(self) -> pulumi.Output[Optional['outputs.BucketLogging']]:
        """
        The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        """
        return pulumi.get(self, "logging")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="requesterPays")
    def requester_pays(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        """
        return pulumi.get(self, "requester_pays")

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> pulumi.Output[Optional['outputs.BucketRetentionPolicy']]:
        """
        Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        """
        return pulumi.get(self, "retention_policy")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> pulumi.Output[Optional[str]]:
        """
        The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`.
        """
        return pulumi.get(self, "storage_class")

    @property
    @pulumi.getter(name="uniformBucketLevelAccess")
    def uniform_bucket_level_access(self) -> pulumi.Output[bool]:
        """
        Enables [Uniform bucket-level access](https://cloud.google.com/storage/docs/uniform-bucket-level-access) access to a bucket.
        """
        return pulumi.get(self, "uniform_bucket_level_access")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The base URL of the bucket, in the format `gs://<bucket-name>`.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def versioning(self) -> pulumi.Output[Optional['outputs.BucketVersioning']]:
        """
        The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        """
        return pulumi.get(self, "versioning")

    @property
    @pulumi.getter
    def website(self) -> pulumi.Output[Optional['outputs.BucketWebsite']]:
        """
        Configuration if the bucket acts as a website. Structure is documented below.
        """
        return pulumi.get(self, "website")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

