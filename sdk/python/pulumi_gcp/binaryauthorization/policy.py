# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Policy(pulumi.CustomResource):
    admission_whitelist_patterns: pulumi.Output[list]
    cluster_admission_rules: pulumi.Output[list]
    default_admission_rule: pulumi.Output[dict]
    description: pulumi.Output[str]
    project: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, admission_whitelist_patterns=None, cluster_admission_rules=None, default_admission_rule=None, description=None, project=None, __name__=None, __opts__=None):
        """
        A policy for container image binary authorization.
        
        > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
        See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
        
        To get more information about Policy, see:
        
        * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/binary-authorization/)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['admission_whitelist_patterns'] = admission_whitelist_patterns

        __props__['cluster_admission_rules'] = cluster_admission_rules

        if default_admission_rule is None:
            raise TypeError("Missing required property 'default_admission_rule'")
        __props__['default_admission_rule'] = default_admission_rule

        __props__['description'] = description

        __props__['project'] = project

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Policy, __self__).__init__(
            'gcp:binaryauthorization/policy:Policy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

