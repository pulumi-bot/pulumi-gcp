# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Policy(pulumi.CustomResource):
    admission_whitelist_patterns: pulumi.Output[list]
    cluster_admission_rules: pulumi.Output[list]
    default_admission_rule: pulumi.Output[dict]
    description: pulumi.Output[str]
    project: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, admission_whitelist_patterns=None, cluster_admission_rules=None, default_admission_rule=None, description=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        A policy for container image binary authorization.
        
        
        To get more information about Policy, see:
        
        * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/binary-authorization/)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_policy.html.markdown.
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

            __props__['admission_whitelist_patterns'] = admission_whitelist_patterns
            __props__['cluster_admission_rules'] = cluster_admission_rules
            if default_admission_rule is None:
                raise TypeError("Missing required property 'default_admission_rule'")
            __props__['default_admission_rule'] = default_admission_rule
            __props__['description'] = description
            __props__['project'] = project
        super(Policy, __self__).__init__(
            'gcp:binaryauthorization/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admission_whitelist_patterns=None, cluster_admission_rules=None, default_admission_rule=None, description=None, project=None):
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_policy.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["admission_whitelist_patterns"] = admission_whitelist_patterns
        __props__["cluster_admission_rules"] = cluster_admission_rules
        __props__["default_admission_rule"] = default_admission_rule
        __props__["description"] = description
        __props__["project"] = project
        return Policy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

