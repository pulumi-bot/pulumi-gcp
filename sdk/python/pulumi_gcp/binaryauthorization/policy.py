# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Policy(pulumi.CustomResource):
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
    page_title: "Google: google_binary_authorization_policy"
    sidebar_current: "docs-google-binary-authorization-policy"
    description: |-
      A policy for container image binary authorization.
    ---
    
    # google\_binary\_authorization\_policy
    
    A policy for container image binary authorization.
    
    > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
    
    To get more information about Policy, see:
    
    * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
    * How-to Guides
        * [Official Documentation](https://cloud.google.com/binary-authorization/)
    """
    def __init__(__self__, __name__, __opts__=None, admission_whitelist_patterns=None, cluster_admission_rules=None, default_admission_rule=None, description=None, project=None):
        """Create a Policy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['admission_whitelist_patterns'] = admission_whitelist_patterns

        __props__['cluster_admission_rules'] = cluster_admission_rules

        if not default_admission_rule:
            raise TypeError('Missing required property default_admission_rule')
        __props__['default_admission_rule'] = default_admission_rule

        __props__['description'] = description

        __props__['project'] = project

        super(Policy, __self__).__init__(
            'gcp:binaryauthorization/policy:Policy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

