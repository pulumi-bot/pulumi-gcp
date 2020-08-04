# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class EngineSplitTraffic(pulumi.CustomResource):
    migrate_traffic: pulumi.Output[bool]
    """
    If set to true traffic will be migrated to this version.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    service: pulumi.Output[str]
    """
    The name of the service these settings apply to.
    """
    split: pulumi.Output[dict]
    """
    Mapping that defines fractional HTTP traffic diversion to different versions within the service.  Structure is documented below.

      * `allocations` (`dict`) - Mapping from version IDs within the service to fractional (0.000, 1] allocations of traffic for that version. Each version can be specified only once, but some versions in the service may not have any traffic allocation. Services that have traffic allocated cannot be deleted until either the service is deleted or their traffic allocation is removed. Allocations must sum to 1. Up to two decimal place precision is supported for IP-based splits and up to three decimal places is supported for cookie-based splits.
      * `shardBy` (`str`) - Mechanism used to determine which version a request is sent to. The traffic selection algorithm will be stable for either type until allocations are changed.
    """
    def __init__(__self__, resource_name, opts=None, migrate_traffic=None, project=None, service=None, split=None, __props__=None, __name__=None, __opts__=None):
        """
        Traffic routing configuration for versions within a single service. Traffic splits define how traffic directed to the service is assigned to versions.

        To get more information about ServiceSplitTraffic, see:

        * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services)

        ## Example Usage
        ### App Engine Service Split Traffic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        bucket = gcp.storage.Bucket("bucket")
        object = gcp.storage.BucketObject("object",
            bucket=bucket.name,
            source=pulumi.FileAsset("./test-fixtures/appengine/hello-world.zip"))
        liveapp_v1 = gcp.appengine.StandardAppVersion("liveappV1",
            version_id="v1",
            service="liveapp",
            delete_service_on_destroy=True,
            runtime="nodejs10",
            entrypoint=gcp.appengine.StandardAppVersionEntrypointArgs(
                shell="node ./app.js",
            ),
            deployment=gcp.appengine.StandardAppVersionDeploymentArgs(
                zip=gcp.appengine.StandardAppVersionDeploymentZipArgs(
                    source_url=pulumi.Output.all(bucket.name, object.name).apply(lambda bucketName, objectName: f"https://storage.googleapis.com/{bucket_name}/{object_name}"),
                ),
            ),
            env_variables={
                "port": "8080",
            })
        liveapp_v2 = gcp.appengine.StandardAppVersion("liveappV2",
            version_id="v2",
            service="liveapp",
            noop_on_destroy=True,
            runtime="nodejs10",
            entrypoint=gcp.appengine.StandardAppVersionEntrypointArgs(
                shell="node ./app.js",
            ),
            deployment=gcp.appengine.StandardAppVersionDeploymentArgs(
                zip=gcp.appengine.StandardAppVersionDeploymentZipArgs(
                    source_url=pulumi.Output.all(bucket.name, object.name).apply(lambda bucketName, objectName: f"https://storage.googleapis.com/{bucket_name}/{object_name}"),
                ),
            ),
            env_variables={
                "port": "8080",
            })
        liveapp = gcp.appengine.EngineSplitTraffic("liveapp",
            service=liveapp_v2.service,
            migrate_traffic=False,
            split=gcp.appengine.EngineSplitTrafficSplitArgs(
                shard_by="IP",
                allocations=pulumi.Output.all(liveapp_v1.version_id, liveapp_v2.version_id).apply(lambda liveappV1Version_id, liveappV2Version_id: {
                    liveapp_v1_version_id: 0.75,
                    liveapp_v2_version_id: 0.25,
                }),
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] migrate_traffic: If set to true traffic will be migrated to this version.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service: The name of the service these settings apply to.
        :param pulumi.Input[dict] split: Mapping that defines fractional HTTP traffic diversion to different versions within the service.  Structure is documented below.

        The **split** object supports the following:

          * `allocations` (`pulumi.Input[dict]`) - Mapping from version IDs within the service to fractional (0.000, 1] allocations of traffic for that version. Each version can be specified only once, but some versions in the service may not have any traffic allocation. Services that have traffic allocated cannot be deleted until either the service is deleted or their traffic allocation is removed. Allocations must sum to 1. Up to two decimal place precision is supported for IP-based splits and up to three decimal places is supported for cookie-based splits.
          * `shardBy` (`pulumi.Input[str]`) - Mechanism used to determine which version a request is sent to. The traffic selection algorithm will be stable for either type until allocations are changed.
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

            __props__['migrate_traffic'] = migrate_traffic
            __props__['project'] = project
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            if split is None:
                raise TypeError("Missing required property 'split'")
            __props__['split'] = split
        super(EngineSplitTraffic, __self__).__init__(
            'gcp:appengine/engineSplitTraffic:EngineSplitTraffic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, migrate_traffic=None, project=None, service=None, split=None):
        """
        Get an existing EngineSplitTraffic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] migrate_traffic: If set to true traffic will be migrated to this version.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service: The name of the service these settings apply to.
        :param pulumi.Input[dict] split: Mapping that defines fractional HTTP traffic diversion to different versions within the service.  Structure is documented below.

        The **split** object supports the following:

          * `allocations` (`pulumi.Input[dict]`) - Mapping from version IDs within the service to fractional (0.000, 1] allocations of traffic for that version. Each version can be specified only once, but some versions in the service may not have any traffic allocation. Services that have traffic allocated cannot be deleted until either the service is deleted or their traffic allocation is removed. Allocations must sum to 1. Up to two decimal place precision is supported for IP-based splits and up to three decimal places is supported for cookie-based splits.
          * `shardBy` (`pulumi.Input[str]`) - Mechanism used to determine which version a request is sent to. The traffic selection algorithm will be stable for either type until allocations are changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["migrate_traffic"] = migrate_traffic
        __props__["project"] = project
        __props__["service"] = service
        __props__["split"] = split
        return EngineSplitTraffic(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
