# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Hl7Store(pulumi.CustomResource):
    dataset: pulumi.Output[str]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    notification_config: pulumi.Output[dict]
    parser_config: pulumi.Output[dict]
    self_link: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, dataset=None, labels=None, name=None, notification_config=None, parser_config=None, __name__=None, __opts__=None):
        """
        Create a Hl7Store resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_hl7_v2_store.html.markdown.
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

        if dataset is None:
            raise TypeError("Missing required property 'dataset'")
        __props__['dataset'] = dataset

        __props__['labels'] = labels

        __props__['name'] = name

        __props__['notification_config'] = notification_config

        __props__['parser_config'] = parser_config

        __props__['self_link'] = None

        super(Hl7Store, __self__).__init__(
            'gcp:healthcare/hl7Store:Hl7Store',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

