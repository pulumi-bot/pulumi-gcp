# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ProjectLocation(pulumi.CustomResource):
    location_id: pulumi.Output[str]
    """
    The ID of the default GCP resource location for the Project. The location must be one of the available GCP
    resource locations.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, location_id=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Sets the default Google Cloud Platform (GCP) resource location for the specified FirebaseProject.

        This method creates an App Engine application with a default Cloud Storage bucket, located in the specified
        locationId. This location must be one of the available GCP resource locations.

        After the default GCP resource location is finalized, or if it was already set, it cannot be changed.
        The default GCP resource location for the specified FirebaseProject might already be set because either the
        GCP Project already has an App Engine application or defaultLocation.finalize was previously called with a
        specified locationId. Any new calls to defaultLocation.finalize with a different specified locationId will
        return a 409 error.

        To get more information about ProjectLocation, see:

        * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.defaultLocation/finalize)
        * How-to Guides
            * [Official Documentation](https://firebase.google.com/)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location_id: The ID of the default GCP resource location for the Project. The location must be one of the available GCP
               resource locations.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

            if location_id is None:
                raise TypeError("Missing required property 'location_id'")
            __props__['location_id'] = location_id
            __props__['project'] = project
        super(ProjectLocation, __self__).__init__(
            'gcp:firebase/projectLocation:ProjectLocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location_id=None, project=None):
        """
        Get an existing ProjectLocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location_id: The ID of the default GCP resource location for the Project. The location must be one of the available GCP
               resource locations.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location_id"] = location_id
        __props__["project"] = project
        return ProjectLocation(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
