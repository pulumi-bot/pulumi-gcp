# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ApplicationUrlDispatchRules(pulumi.CustomResource):
    dispatch_rules: pulumi.Output[list]
    """
    Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.

      * `domain` (`str`) - Domain name to match against. The wildcard "*" is supported if specified before a period: "*.".
        Defaults to matching all domains: "*".
      * `path` (`str`) - Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
        The sum of the lengths of the domain and path may not exceed 100 characters.
      * `service` (`str`) - Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
        The sum of the lengths of the domain and path may not exceed 100 characters.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, dispatch_rules=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Rules to match an HTTP request and dispatch that request to a service.


        To get more information about ApplicationUrlDispatchRules, see:

        * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)

        ## Example Usage

        ### App Engine Application Url Dispatch Rules Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        bucket = gcp.storage.Bucket("bucket")
        object = gcp.storage.BucketObject("object",
            bucket=bucket.name,
            source=pulumi.FileAsset("./test-fixtures/appengine/hello-world.zip"))
        admin_v3 = gcp.appengine.StandardAppVersion("adminV3",
            version_id="v3",
            service="admin",
            runtime="nodejs10",
            entrypoint={
                "shell": "node ./app.js",
            },
            deployment={
                "zip": {
                    "sourceUrl": pulumi.Output.all(bucket.name, object.name).apply(lambda bucketName, objectName: f"https://storage.googleapis.com/{bucket_name}/{object_name}"),
                },
            },
            env_variables={
                "port": "8080",
            },
            noop_on_destroy=True)
        web_service = gcp.appengine.ApplicationUrlDispatchRules("webService", dispatch_rules=[
            {
                "domain": "*",
                "path": "/*",
                "service": "default",
            },
            {
                "domain": "*",
                "path": "/admin/*",
                "service": admin_v3.service,
            },
        ])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] dispatch_rules: Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **dispatch_rules** object supports the following:

          * `domain` (`pulumi.Input[str]`) - Domain name to match against. The wildcard "*" is supported if specified before a period: "*.".
            Defaults to matching all domains: "*".
          * `path` (`pulumi.Input[str]`) - Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
            The sum of the lengths of the domain and path may not exceed 100 characters.
          * `service` (`pulumi.Input[str]`) - Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
            The sum of the lengths of the domain and path may not exceed 100 characters.
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

            if dispatch_rules is None:
                raise TypeError("Missing required property 'dispatch_rules'")
            __props__['dispatch_rules'] = dispatch_rules
            __props__['project'] = project
        super(ApplicationUrlDispatchRules, __self__).__init__(
            'gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dispatch_rules=None, project=None):
        """
        Get an existing ApplicationUrlDispatchRules resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] dispatch_rules: Rules to match an HTTP request and dispatch that request to a service.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **dispatch_rules** object supports the following:

          * `domain` (`pulumi.Input[str]`) - Domain name to match against. The wildcard "*" is supported if specified before a period: "*.".
            Defaults to matching all domains: "*".
          * `path` (`pulumi.Input[str]`) - Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
            The sum of the lengths of the domain and path may not exceed 100 characters.
          * `service` (`pulumi.Input[str]`) - Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
            The sum of the lengths of the domain and path may not exceed 100 characters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dispatch_rules"] = dispatch_rules
        __props__["project"] = project
        return ApplicationUrlDispatchRules(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

