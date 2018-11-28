# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    """
    #     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
    #
    # ----------------------------------------------------------------------------
    #
    #     This file is automatically generated by Magic Modules and manual
    #     changes will be clobbered when the file is regenerated.
    #
    #     Please read more about how to change this file in
    #     .github/CONTRIBUTING.md.
    #
    # ----------------------------------------------------------------------------
    layout: "google"
    page_title: "Google: google_filestore_instance"
    sidebar_current: "docs-google-filestore-instance"
    description: |-
      A Google Cloud Filestore instance.
    ---
    
    # google\_filestore\_instance
    
    A Google Cloud Filestore instance.
    
    ~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
    
    To get more information about Instance, see:
    
    * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
    * How-to Guides
        * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
        * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
        * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)
    """
    def __init__(__self__, __name__, __opts__=None, description=None, file_shares=None, labels=None, name=None, networks=None, project=None, tier=None, zone=None):
        """Create a Instance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        if not file_shares:
            raise TypeError('Missing required property file_shares')
        __props__['file_shares'] = file_shares

        __props__['labels'] = labels

        __props__['name'] = name

        if not networks:
            raise TypeError('Missing required property networks')
        __props__['networks'] = networks

        __props__['project'] = project

        if not tier:
            raise TypeError('Missing required property tier')
        __props__['tier'] = tier

        if not zone:
            raise TypeError('Missing required property zone')
        __props__['zone'] = zone

        __props__['create_time'] = None
        __props__['etag'] = None

        super(Instance, __self__).__init__(
            'gcp:filestore/instance:Instance',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

