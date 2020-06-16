# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Application(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    Identifier of the app, usually `{PROJECT_ID}`
    """
    auth_domain: pulumi.Output[str]
    """
    The domain to authenticate users with when using App Engine's User API.
    """
    code_bucket: pulumi.Output[str]
    """
    The GCS bucket code is being stored in for this app.
    """
    default_bucket: pulumi.Output[str]
    """
    The GCS bucket content is being stored in for this app.
    """
    default_hostname: pulumi.Output[str]
    """
    The default hostname for this app.
    """
    feature_settings: pulumi.Output[dict]
    """
    A block of optional settings to configure specific App Engine features:

      * `splitHealthChecks` (`bool`) - Set to false to use the legacy health check instead of the readiness
        and liveness checks.
    """
    gcr_domain: pulumi.Output[str]
    """
    The GCR domain used for storing managed Docker images for this app.
    """
    iap: pulumi.Output[dict]
    """
    Settings for enabling Cloud Identity Aware Proxy

      * `enabled` (`bool`)
      * `oauth2ClientId` (`str`) - OAuth2 client ID to use for the authentication flow.
      * `oauth2ClientSecret` (`str`) - OAuth2 client secret to use for the authentication flow.
        The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
      * `oauth2ClientSecretSha256` (`str`) - Hex-encoded SHA-256 hash of the client secret.
    """
    location_id: pulumi.Output[str]
    """
    The [location](https://cloud.google.com/appengine/docs/locations)
    to serve the app from.
    """
    name: pulumi.Output[str]
    """
    Unique name of the app, usually `apps/{PROJECT_ID}`
    """
    project: pulumi.Output[str]
    """
    The project ID to create the application under.
    ~>**NOTE**: GCP only accepts project ID, not project number. If you are using number,
    you may get a "Permission denied" error.
    """
    serving_status: pulumi.Output[str]
    """
    The serving status of the app.
    """
    url_dispatch_rules: pulumi.Output[list]
    """
    A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.

      * `domain` (`str`)
      * `path` (`str`)
      * `service` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, auth_domain=None, feature_settings=None, iap=None, location_id=None, project=None, serving_status=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows creation and management of an App Engine application.

        > App Engine applications cannot be deleted once they're created; you have to delete the
           entire project to delete the application. This provider will report the application has been
           successfully deleted; this is a limitation of the provider, and will go away in the future.
           This provider is not able to delete App Engine applications.

        > **Warning:** All arguments including `iap.oauth2_client_secret` will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

        ## Example Usage



        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            project_id="your-project-id",
            org_id="1234567")
        app = gcp.appengine.Application("app",
            project=my_project.project_id,
            location_id="us-central")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input[dict] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input[dict] iap: Settings for enabling Cloud Identity Aware Proxy
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE**: GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.

        The **feature_settings** object supports the following:

          * `splitHealthChecks` (`pulumi.Input[bool]`) - Set to false to use the legacy health check instead of the readiness
            and liveness checks.

        The **iap** object supports the following:

          * `enabled` (`pulumi.Input[bool]`)
          * `oauth2ClientId` (`pulumi.Input[str]`) - OAuth2 client ID to use for the authentication flow.
          * `oauth2ClientSecret` (`pulumi.Input[str]`) - OAuth2 client secret to use for the authentication flow.
            The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
          * `oauth2ClientSecretSha256` (`pulumi.Input[str]`) - Hex-encoded SHA-256 hash of the client secret.
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

            __props__['auth_domain'] = auth_domain
            __props__['feature_settings'] = feature_settings
            __props__['iap'] = iap
            if location_id is None:
                raise TypeError("Missing required property 'location_id'")
            __props__['location_id'] = location_id
            __props__['project'] = project
            __props__['serving_status'] = serving_status
            __props__['app_id'] = None
            __props__['code_bucket'] = None
            __props__['default_bucket'] = None
            __props__['default_hostname'] = None
            __props__['gcr_domain'] = None
            __props__['name'] = None
            __props__['url_dispatch_rules'] = None
        super(Application, __self__).__init__(
            'gcp:appengine/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, auth_domain=None, code_bucket=None, default_bucket=None, default_hostname=None, feature_settings=None, gcr_domain=None, iap=None, location_id=None, name=None, project=None, serving_status=None, url_dispatch_rules=None):
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Identifier of the app, usually `{PROJECT_ID}`
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input[str] code_bucket: The GCS bucket code is being stored in for this app.
        :param pulumi.Input[str] default_bucket: The GCS bucket content is being stored in for this app.
        :param pulumi.Input[str] default_hostname: The default hostname for this app.
        :param pulumi.Input[dict] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input[str] gcr_domain: The GCR domain used for storing managed Docker images for this app.
        :param pulumi.Input[dict] iap: Settings for enabling Cloud Identity Aware Proxy
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] name: Unique name of the app, usually `apps/{PROJECT_ID}`
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE**: GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.
        :param pulumi.Input[list] url_dispatch_rules: A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.

        The **feature_settings** object supports the following:

          * `splitHealthChecks` (`pulumi.Input[bool]`) - Set to false to use the legacy health check instead of the readiness
            and liveness checks.

        The **iap** object supports the following:

          * `enabled` (`pulumi.Input[bool]`)
          * `oauth2ClientId` (`pulumi.Input[str]`) - OAuth2 client ID to use for the authentication flow.
          * `oauth2ClientSecret` (`pulumi.Input[str]`) - OAuth2 client secret to use for the authentication flow.
            The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
          * `oauth2ClientSecretSha256` (`pulumi.Input[str]`) - Hex-encoded SHA-256 hash of the client secret.

        The **url_dispatch_rules** object supports the following:

          * `domain` (`pulumi.Input[str]`)
          * `path` (`pulumi.Input[str]`)
          * `service` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["auth_domain"] = auth_domain
        __props__["code_bucket"] = code_bucket
        __props__["default_bucket"] = default_bucket
        __props__["default_hostname"] = default_hostname
        __props__["feature_settings"] = feature_settings
        __props__["gcr_domain"] = gcr_domain
        __props__["iap"] = iap
        __props__["location_id"] = location_id
        __props__["name"] = name
        __props__["project"] = project
        __props__["serving_status"] = serving_status
        __props__["url_dispatch_rules"] = url_dispatch_rules
        return Application(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
