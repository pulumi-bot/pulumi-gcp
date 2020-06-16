# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Realm(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    Human readable description of the realm.
    """
    etag: pulumi.Output[str]
    """
    ETag of the resource.
    """
    labels: pulumi.Output[dict]
    """
    The labels associated with this realm. Each label is a key-value pair.
    """
    location: pulumi.Output[str]
    """
    Location of the Realm.
    """
    name: pulumi.Output[str]
    """
    The resource id of the realm, of the form: 'projects/{project_id}/locations/{location}/realms/{realm_id}'. For example,
    'projects/my-project/locations/{location}/realms/my-realm'.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    realm_id: pulumi.Output[str]
    """
    GCP region of the Realm.
    """
    time_zone: pulumi.Output[str]
    """
    Required. Time zone where all realm-specific policies are evaluated. The value of
    this field must be from the IANA time zone database:
    https://www.iana.org/time-zones.
    """
    def __init__(__self__, resource_name, opts=None, description=None, labels=None, location=None, project=None, realm_id=None, time_zone=None, __props__=None, __name__=None, __opts__=None):
        """
        A Realm resource.

        To get more information about Realm, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        ## Example Usage

        ### Game Service Realm Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.gameservices.Realm("default",
            realm_id="tf-test-realm",
            time_zone="EST",
            location="global",
            description="one of the nine")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Human readable description of the realm.
        :param pulumi.Input[dict] labels: The labels associated with this realm. Each label is a key-value pair.
        :param pulumi.Input[str] location: Location of the Realm.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] realm_id: GCP region of the Realm.
        :param pulumi.Input[str] time_zone: Required. Time zone where all realm-specific policies are evaluated. The value of
               this field must be from the IANA time zone database:
               https://www.iana.org/time-zones.
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

            __props__['description'] = description
            __props__['labels'] = labels
            __props__['location'] = location
            __props__['project'] = project
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if time_zone is None:
                raise TypeError("Missing required property 'time_zone'")
            __props__['time_zone'] = time_zone
            __props__['etag'] = None
            __props__['name'] = None
        super(Realm, __self__).__init__(
            'gcp:gameservices/realm:Realm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, etag=None, labels=None, location=None, name=None, project=None, realm_id=None, time_zone=None):
        """
        Get an existing Realm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Human readable description of the realm.
        :param pulumi.Input[str] etag: ETag of the resource.
        :param pulumi.Input[dict] labels: The labels associated with this realm. Each label is a key-value pair.
        :param pulumi.Input[str] location: Location of the Realm.
        :param pulumi.Input[str] name: The resource id of the realm, of the form: 'projects/{project_id}/locations/{location}/realms/{realm_id}'. For example,
               'projects/my-project/locations/{location}/realms/my-realm'.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] realm_id: GCP region of the Realm.
        :param pulumi.Input[str] time_zone: Required. Time zone where all realm-specific policies are evaluated. The value of
               this field must be from the IANA time zone database:
               https://www.iana.org/time-zones.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["etag"] = etag
        __props__["labels"] = labels
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["realm_id"] = realm_id
        __props__["time_zone"] = time_zone
        return Realm(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
