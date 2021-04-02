# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BillingAccountBucketConfigArgs', 'BillingAccountBucketConfig']

@pulumi.input_type
class BillingAccountBucketConfigArgs:
    def __init__(__self__, *,
                 billing_account: pulumi.Input[str],
                 bucket_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a BillingAccountBucketConfig resource.
        :param pulumi.Input[str] billing_account: The parent resource that contains the logging bucket.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        pulumi.set(__self__, "billing_account", billing_account)
        pulumi.set(__self__, "bucket_id", bucket_id)
        pulumi.set(__self__, "location", location)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> pulumi.Input[str]:
        """
        The parent resource that contains the logging bucket.
        """
        return pulumi.get(self, "billing_account")

    @billing_account.setter
    def billing_account(self, value: pulumi.Input[str]):
        pulumi.set(self, "billing_account", value)

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
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)


@pulumi.input_type
class _BillingAccountBucketConfigState:
    def __init__(__self__, *,
                 billing_account: Optional[pulumi.Input[str]] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lifecycle_state: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering BillingAccountBucketConfig resources.
        :param pulumi.Input[str] billing_account: The parent resource that contains the logging bucket.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] lifecycle_state: The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[str] name: The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        if billing_account is not None:
            pulumi.set(__self__, "billing_account", billing_account)
        if bucket_id is not None:
            pulumi.set(__self__, "bucket_id", bucket_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if lifecycle_state is not None:
            pulumi.set(__self__, "lifecycle_state", lifecycle_state)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> Optional[pulumi.Input[str]]:
        """
        The parent resource that contains the logging bucket.
        """
        return pulumi.get(self, "billing_account")

    @billing_account.setter
    def billing_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_account", value)

    @property
    @pulumi.getter(name="bucketId")
    def bucket_id(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        """
        return pulumi.get(self, "bucket_id")

    @bucket_id.setter
    def bucket_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_id", value)

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
    @pulumi.getter(name="lifecycleState")
    def lifecycle_state(self) -> Optional[pulumi.Input[str]]:
        """
        The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        """
        return pulumi.get(self, "lifecycle_state")

    @lifecycle_state.setter
    def lifecycle_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lifecycle_state", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location of the bucket.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)


class BillingAccountBucketConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account: Optional[pulumi.Input[str]] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a billing account level logging bucket config. For more information see
        [the official logging documentation](https://cloud.google.com/logging/docs/) and
        [Storing Logs](https://cloud.google.com/logging/docs/storage).

        > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.organizations.get_billing_account(billing_account="00AA00-000AAA-00AA0A")
        basic = gcp.logging.BillingAccountBucketConfig("basic",
            billing_account=default.billing_account,
            location="global",
            retention_days=30,
            bucket_id="_Default")
        ```

        ## Import

        This resource can be imported using the following format

        ```sh
         $ pulumi import gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig default billingAccounts/{{billingAccount}}/locations/{{location}}/buckets/{{bucket_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account: The parent resource that contains the logging bucket.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BillingAccountBucketConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a billing account level logging bucket config. For more information see
        [the official logging documentation](https://cloud.google.com/logging/docs/) and
        [Storing Logs](https://cloud.google.com/logging/docs/storage).

        > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.organizations.get_billing_account(billing_account="00AA00-000AAA-00AA0A")
        basic = gcp.logging.BillingAccountBucketConfig("basic",
            billing_account=default.billing_account,
            location="global",
            retention_days=30,
            bucket_id="_Default")
        ```

        ## Import

        This resource can be imported using the following format

        ```sh
         $ pulumi import gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig default billingAccounts/{{billingAccount}}/locations/{{location}}/buckets/{{bucket_id}}
        ```

        :param str resource_name: The name of the resource.
        :param BillingAccountBucketConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BillingAccountBucketConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account: Optional[pulumi.Input[str]] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
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
            __props__ = BillingAccountBucketConfigArgs.__new__(BillingAccountBucketConfigArgs)

            if billing_account is None and not opts.urn:
                raise TypeError("Missing required property 'billing_account'")
            __props__.__dict__['billing_account'] = billing_account
            if bucket_id is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_id'")
            __props__.__dict__['bucket_id'] = bucket_id
            __props__.__dict__['description'] = description
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__['location'] = location
            __props__.__dict__['retention_days'] = retention_days
            __props__.__dict__['lifecycle_state'] = None
            __props__.__dict__['name'] = None
        super(BillingAccountBucketConfig, __self__).__init__(
            'gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            billing_account: Optional[pulumi.Input[str]] = None,
            bucket_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lifecycle_state: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            retention_days: Optional[pulumi.Input[int]] = None) -> 'BillingAccountBucketConfig':
        """
        Get an existing BillingAccountBucketConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account: The parent resource that contains the logging bucket.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] lifecycle_state: The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[str] name: The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BillingAccountBucketConfigState.__new__(_BillingAccountBucketConfigState)

        __props__.__dict__['billing_account'] = billing_account
        __props__.__dict__['bucket_id'] = bucket_id
        __props__.__dict__['description'] = description
        __props__.__dict__['lifecycle_state'] = lifecycle_state
        __props__.__dict__['location'] = location
        __props__.__dict__['name'] = name
        __props__.__dict__['retention_days'] = retention_days
        return BillingAccountBucketConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> pulumi.Output[str]:
        """
        The parent resource that contains the logging bucket.
        """
        return pulumi.get(self, "billing_account")

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
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        return pulumi.get(self, "retention_days")

