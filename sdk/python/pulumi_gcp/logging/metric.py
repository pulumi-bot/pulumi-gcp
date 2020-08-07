# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Metric']


class Metric(pulumi.CustomResource):
    bucket_options: pulumi.Output[Optional['outputs.MetricBucketOptions']] = pulumi.property("bucketOptions")
    """
    The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
    describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    A description of this metric, which is used in documentation. The maximum length of the
    description is 8000 characters.
    """

    filter: pulumi.Output[str] = pulumi.property("filter")
    """
    An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
    is used to match log entries.
    """

    label_extractors: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("labelExtractors")
    """
    A map from a label key string to an extractor expression which is used to extract data from a log
    entry field and assign as the label value. Each label key specified in the LabelDescriptor must
    have an associated extractor expression in this map. The syntax of the extractor expression is
    the same as for the valueExtractor field.
    """

    metric_descriptor: pulumi.Output['outputs.MetricMetricDescriptor'] = pulumi.property("metricDescriptor")
    """
    The metric descriptor associated with the logs-based metric.  Structure is documented below.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The client-assigned metric identifier. Examples - "error_count", "nginx/requests".
    Metric identifiers are limited to 100 characters and can include only the following
    characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
    character (/) denotes a hierarchy of name pieces, and it cannot be the first character
    of the name.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """

    value_extractor: pulumi.Output[Optional[str]] = pulumi.property("valueExtractor")
    """
    A valueExtractor is required when using a distribution logs-based metric to extract the values to
    record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
    REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
    the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
    (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
    log entry field. The value of the field is converted to a string before applying the regex. It is an
    error to specify a regex that does not include exactly one capture group.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_options: Optional[pulumi.Input[pulumi.InputType['MetricBucketOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 label_extractors: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 metric_descriptor: Optional[pulumi.Input[pulumi.InputType['MetricMetricDescriptorArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 value_extractor: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Logs-based metric can also be used to extract values from logs and create a a distribution
        of the values. The distribution records the statistics of the extracted values along with
        an optional histogram of the values as specified by the bucket options.

        To get more information about Metric, see:

        * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics/create)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/logging/docs/apis)

        ## Example Usage
        ### Logging Metric Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        logging_metric = gcp.logging.Metric("loggingMetric",
            bucket_options={
                "linearBuckets": {
                    "numFiniteBuckets": 3,
                    "offset": 1,
                    "width": 1,
                },
            },
            filter="resource.type=gae_app AND severity>=ERROR",
            label_extractors={
                "mass": "EXTRACT(jsonPayload.request)",
                "sku": "EXTRACT(jsonPayload.id)",
            },
            metric_descriptor={
                "display_name": "My metric",
                "labels": [
                    {
                        "description": "amount of matter",
                        "key": "mass",
                        "value_type": "STRING",
                    },
                    {
                        "description": "Identifying number for item",
                        "key": "sku",
                        "value_type": "INT64",
                    },
                ],
                "metric_kind": "DELTA",
                "unit": "1",
                "value_type": "DISTRIBUTION",
            },
            value_extractor="EXTRACT(jsonPayload.request)")
        ```
        ### Logging Metric Counter Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        logging_metric = gcp.logging.Metric("loggingMetric",
            filter="resource.type=gae_app AND severity>=ERROR",
            metric_descriptor={
                "metric_kind": "DELTA",
                "value_type": "INT64",
            })
        ```
        ### Logging Metric Counter Labels

        ```python
        import pulumi
        import pulumi_gcp as gcp

        logging_metric = gcp.logging.Metric("loggingMetric",
            filter="resource.type=gae_app AND severity>=ERROR",
            label_extractors={
                "mass": "EXTRACT(jsonPayload.request)",
            },
            metric_descriptor={
                "labels": [{
                    "description": "amount of matter",
                    "key": "mass",
                    "value_type": "STRING",
                }],
                "metric_kind": "DELTA",
                "value_type": "INT64",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MetricBucketOptionsArgs']] bucket_options: The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
               describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.
        :param pulumi.Input[str] description: A description of this metric, which is used in documentation. The maximum length of the
               description is 8000 characters.
        :param pulumi.Input[str] filter: An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
               is used to match log entries.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] label_extractors: A map from a label key string to an extractor expression which is used to extract data from a log
               entry field and assign as the label value. Each label key specified in the LabelDescriptor must
               have an associated extractor expression in this map. The syntax of the extractor expression is
               the same as for the valueExtractor field.
        :param pulumi.Input[pulumi.InputType['MetricMetricDescriptorArgs']] metric_descriptor: The metric descriptor associated with the logs-based metric.  Structure is documented below.
        :param pulumi.Input[str] name: The client-assigned metric identifier. Examples - "error_count", "nginx/requests".
               Metric identifiers are limited to 100 characters and can include only the following
               characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
               character (/) denotes a hierarchy of name pieces, and it cannot be the first character
               of the name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] value_extractor: A valueExtractor is required when using a distribution logs-based metric to extract the values to
               record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
               REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
               the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
               (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
               log entry field. The value of the field is converted to a string before applying the regex. It is an
               error to specify a regex that does not include exactly one capture group.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
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
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_options: Optional[pulumi.Input[pulumi.InputType['MetricBucketOptionsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            label_extractors: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            metric_descriptor: Optional[pulumi.Input[pulumi.InputType['MetricMetricDescriptorArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            value_extractor: Optional[pulumi.Input[str]] = None) -> 'Metric':
        """
        Get an existing Metric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MetricBucketOptionsArgs']] bucket_options: The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
               describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.
        :param pulumi.Input[str] description: A description of this metric, which is used in documentation. The maximum length of the
               description is 8000 characters.
        :param pulumi.Input[str] filter: An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
               is used to match log entries.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] label_extractors: A map from a label key string to an extractor expression which is used to extract data from a log
               entry field and assign as the label value. Each label key specified in the LabelDescriptor must
               have an associated extractor expression in this map. The syntax of the extractor expression is
               the same as for the valueExtractor field.
        :param pulumi.Input[pulumi.InputType['MetricMetricDescriptorArgs']] metric_descriptor: The metric descriptor associated with the logs-based metric.  Structure is documented below.
        :param pulumi.Input[str] name: The client-assigned metric identifier. Examples - "error_count", "nginx/requests".
               Metric identifiers are limited to 100 characters and can include only the following
               characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
               character (/) denotes a hierarchy of name pieces, and it cannot be the first character
               of the name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] value_extractor: A valueExtractor is required when using a distribution logs-based metric to extract the values to
               record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
               REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
               the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
               (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
               log entry field. The value of the field is converted to a string before applying the regex. It is an
               error to specify a regex that does not include exactly one capture group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

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
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

