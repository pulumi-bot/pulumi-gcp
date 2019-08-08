# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Application(pulumi.CustomResource):
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
    """
    gcr_domain: pulumi.Output[str]
    """
    The GCR domain used for storing managed Docker images for this app.
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
    """
    def __init__(__self__, resource_name, opts=None, auth_domain=None, feature_settings=None, location_id=None, project=None, serving_status=None, __name__=None, __opts__=None):
        """
        Allows creation and management of an App Engine application.
        
        > App Engine applications cannot be deleted once they're created; you have to delete the
           entire project to delete the application. This provider will report the application has been
           successfully deleted; this is a limitation of this provider, and will go away in the future.
           This provider is not able to delete App Engine applications.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input[dict] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE**: GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/app_engine_application.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['auth_domain'] = auth_domain
        __props__['feature_settings'] = feature_settings
        if location_id is None:
            raise TypeError("Missing required property 'location_id'")
        __props__['location_id'] = location_id
        __props__['project'] = project
        __props__['serving_status'] = serving_status
        __props__['code_bucket'] = None
        __props__['default_bucket'] = None
        __props__['default_hostname'] = None
        __props__['gcr_domain'] = None
        __props__['name'] = None
        __props__['url_dispatch_rules'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Application, __self__).__init__(
            'gcp:appengine/application:Application',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

