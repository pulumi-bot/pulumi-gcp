# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Metric(pulumi.CustomResource):
    bucket_options: pulumi.Output[dict]
    description: pulumi.Output[str]
    filter: pulumi.Output[str]
    label_extractors: pulumi.Output[dict]
    metric_descriptor: pulumi.Output[dict]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    value_extractor: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, bucket_options=None, description=None, filter=None, label_extractors=None, metric_descriptor=None, name=None, project=None, value_extractor=None, __props__=None, __name__=None, __opts__=None):
        """
        Logs-based metric can also be used to extract values from logs and create a a distribution
        of the values. The distribution records the statistics of the extracted values along with
        an optional histogram of the values as specified by the bucket options.
        
        
        To get more information about Metric, see:
        
        * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics/create)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/logging/docs/apis)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_metric.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            __props__['bucket_options'] = bucket_options
            __props__['description'] = description
            if filter is None:
                raise TypeError("Missing required property 'filter'")
            __props__['filter'] = filter
            __props__['label_extractors'] = label_extractors
            if metric_descriptor is None:
                raise TypeError("Missing required property 'metric_descriptor'")
            __props__['metric_descriptor'] = metric_descriptor
            __props__['name'] = name
            __props__['project'] = project
            __props__['value_extractor'] = value_extractor
        super(Metric, __self__).__init__(
            'gcp:logging/metric:Metric',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bucket_options=None, description=None, filter=None, label_extractors=None, metric_descriptor=None, name=None, project=None, value_extractor=None):
        """
        Get an existing Metric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_metric.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["bucket_options"] = bucket_options
        __props__["description"] = description
        __props__["filter"] = filter
        __props__["label_extractors"] = label_extractors
        __props__["metric_descriptor"] = metric_descriptor
        __props__["name"] = name
        __props__["project"] = project
        __props__["value_extractor"] = value_extractor
        return Metric(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

