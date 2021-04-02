# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['BackendBucketSignedUrlKeyArgs', 'BackendBucketSignedUrlKey']

@pulumi.input_type
class BackendBucketSignedUrlKeyArgs:
    def __init__(__self__, *,
                 backend_bucket: pulumi.Input[str],
                 key_value: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BackendBucketSignedUrlKey resource.
        :param pulumi.Input[str] backend_bucket: The backend bucket this signed URL key belongs.
        :param pulumi.Input[str] key_value: 128-bit key value used for signing the URL. The key value must be a
               valid RFC 4648 Section 5 base64url encoded string.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] name: Name of the signed URL key.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "backend_bucket", backend_bucket)
        pulumi.set(__self__, "key_value", key_value)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="backendBucket")
    def backend_bucket(self) -> pulumi.Input[str]:
        """
        The backend bucket this signed URL key belongs.
        """
        return pulumi.get(self, "backend_bucket")

    @backend_bucket.setter
    def backend_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend_bucket", value)

    @property
    @pulumi.getter(name="keyValue")
    def key_value(self) -> pulumi.Input[str]:
        """
        128-bit key value used for signing the URL. The key value must be a
        valid RFC 4648 Section 5 base64url encoded string.
        **Note**: This property is sensitive and will not be displayed in the plan.
        """
        return pulumi.get(self, "key_value")

    @key_value.setter
    def key_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the signed URL key.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class BackendBucketSignedUrlKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_bucket: Optional[pulumi.Input[str]] = None,
                 key_value: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A key for signing Cloud CDN signed URLs for BackendBuckets.

        To get more information about BackendBucketSignedUrlKey, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendBuckets)
        * How-to Guides
            * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)

        > **Warning:** All arguments including `key_value` will be stored in the raw
        state as plain-text.

        ## Example Usage
        ### Backend Bucket Signed Url Key

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        url_signature = random.RandomId("urlSignature", byte_length=16)
        bucket = gcp.storage.Bucket("bucket", location="EU")
        test_backend = gcp.compute.BackendBucket("testBackend",
            description="Contains beautiful images",
            bucket_name=bucket.name,
            enable_cdn=True)
        backend_key = gcp.compute.BackendBucketSignedUrlKey("backendKey",
            key_value=url_signature.b64_url,
            backend_bucket=test_backend.name)
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend_bucket: The backend bucket this signed URL key belongs.
        :param pulumi.Input[str] key_value: 128-bit key value used for signing the URL. The key value must be a
               valid RFC 4648 Section 5 base64url encoded string.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] name: Name of the signed URL key.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackendBucketSignedUrlKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A key for signing Cloud CDN signed URLs for BackendBuckets.

        To get more information about BackendBucketSignedUrlKey, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendBuckets)
        * How-to Guides
            * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)

        > **Warning:** All arguments including `key_value` will be stored in the raw
        state as plain-text.

        ## Example Usage
        ### Backend Bucket Signed Url Key

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        url_signature = random.RandomId("urlSignature", byte_length=16)
        bucket = gcp.storage.Bucket("bucket", location="EU")
        test_backend = gcp.compute.BackendBucket("testBackend",
            description="Contains beautiful images",
            bucket_name=bucket.name,
            enable_cdn=True)
        backend_key = gcp.compute.BackendBucketSignedUrlKey("backendKey",
            key_value=url_signature.b64_url,
            backend_bucket=test_backend.name)
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param BackendBucketSignedUrlKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackendBucketSignedUrlKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_bucket: Optional[pulumi.Input[str]] = None,
                 key_value: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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

            if backend_bucket is None and not opts.urn:
                raise TypeError("Missing required property 'backend_bucket'")
            __props__['backend_bucket'] = backend_bucket
            if key_value is None and not opts.urn:
                raise TypeError("Missing required property 'key_value'")
            __props__['key_value'] = key_value
            __props__['name'] = name
            __props__['project'] = project
        super(BackendBucketSignedUrlKey, __self__).__init__(
            'gcp:compute/backendBucketSignedUrlKey:BackendBucketSignedUrlKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend_bucket: Optional[pulumi.Input[str]] = None,
            key_value: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'BackendBucketSignedUrlKey':
        """
        Get an existing BackendBucketSignedUrlKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend_bucket: The backend bucket this signed URL key belongs.
        :param pulumi.Input[str] key_value: 128-bit key value used for signing the URL. The key value must be a
               valid RFC 4648 Section 5 base64url encoded string.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] name: Name of the signed URL key.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend_bucket"] = backend_bucket
        __props__["key_value"] = key_value
        __props__["name"] = name
        __props__["project"] = project
        return BackendBucketSignedUrlKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendBucket")
    def backend_bucket(self) -> pulumi.Output[str]:
        """
        The backend bucket this signed URL key belongs.
        """
        return pulumi.get(self, "backend_bucket")

    @property
    @pulumi.getter(name="keyValue")
    def key_value(self) -> pulumi.Output[str]:
        """
        128-bit key value used for signing the URL. The key value must be a
        valid RFC 4648 Section 5 base64url encoded string.
        **Note**: This property is sensitive and will not be displayed in the plan.
        """
        return pulumi.get(self, "key_value")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the signed URL key.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

