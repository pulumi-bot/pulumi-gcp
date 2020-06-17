# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Metric(pulumi.CustomResource):
    bucket_options: pulumi.Output[dict]
    """
    The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
    describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.

      * `explicitBuckets` (`dict`) - Specifies a set of buckets with arbitrary widths.  Structure is documented below.
        * `bounds` (`list`) - The values must be monotonically increasing.

      * `exponentialBuckets` (`dict`) - Specifies an exponential sequence of buckets that have a width that is proportional to the value of
        the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.  Structure is documented below.
        * `growthFactor` (`float`) - Must be greater than 1.
        * `numFiniteBuckets` (`float`) - Must be greater than 0.
        * `scale` (`float`) - Must be greater than 0.

      * `linearBuckets` (`dict`) - Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
        Each bucket represents a constant absolute uncertainty on the specific value in the bucket.  Structure is documented below.
        * `numFiniteBuckets` (`float`) - Must be greater than 0.
        * `offset` (`float`) - Lower bound of the first bucket.
        * `width` (`float`) - Must be greater than 0.
    """
    description: pulumi.Output[str]
    """
    A description of this metric, which is used in documentation. The maximum length of the
    description is 8000 characters.
    """
    filter: pulumi.Output[str]
    """
    An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
    is used to match log entries.
    """
    label_extractors: pulumi.Output[dict]
    """
    A map from a label key string to an extractor expression which is used to extract data from a log
    entry field and assign as the label value. Each label key specified in the LabelDescriptor must
    have an associated extractor expression in this map. The syntax of the extractor expression is
    the same as for the valueExtractor field.
    """
    metric_descriptor: pulumi.Output[dict]
    """
    The metric descriptor associated with the logs-based metric.  Structure is documented below.

      * `display_name` (`str`) - A concise name for the metric, which can be displayed in user interfaces. Use sentence case
        without an ending period, for example "Request count". This field is optional but it is
        recommended to be set for any metrics associated with user-visible concepts, such as Quota.
      * `labels` (`list`) - The set of labels that can be used to describe a specific instance of this metric type. For
        example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
        for the HTTP response code, response_code, so you can look at latencies for successful responses
        or just for responses that failed.  Structure is documented below.
        * `description` (`str`) - A description of this metric, which is used in documentation. The maximum length of the
          description is 8000 characters.
        * `key` (`str`) - The label key.
        * `valueType` (`str`) - The type of data that can be assigned to the label.

      * `metricKind` (`str`) - Whether the metric records instantaneous values, changes to a value, etc.
        Some combinations of metricKind and valueType might not be supported.
        For counter metrics, set this to DELTA.
      * `unit` (`str`) - The unit in which the metric value is reported. It is only applicable if the valueType is
        `INT64`, `DOUBLE`, or `DISTRIBUTION`. The supported units are a subset of
        [The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard
      * `valueType` (`str`) - The type of data that can be assigned to the label.
    """
    name: pulumi.Output[str]
    """
    The client-assigned metric identifier. Examples - "error_count", "nginx/requests".
    Metric identifiers are limited to 100 characters and can include only the following
    characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
    character (/) denotes a hierarchy of name pieces, and it cannot be the first character
    of the name.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    value_extractor: pulumi.Output[str]
    """
    A valueExtractor is required when using a distribution logs-based metric to extract the values to
    record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
    REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
    the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
    (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
    log entry field. The value of the field is converted to a string before applying the regex. It is an
    error to specify a regex that does not include exactly one capture group.
    """
    def __init__(__self__, resource_name, opts=None, bucket_options=None, description=None, filter=None, label_extractors=None, metric_descriptor=None, name=None, project=None, value_extractor=None, __props__=None, __name__=None, __opts__=None):
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
                        "valueType": "STRING",
                    },
                    {
                        "description": "Identifying number for item",
                        "key": "sku",
                        "valueType": "INT64",
                    },
                ],
                "metricKind": "DELTA",
                "unit": "1",
                "valueType": "DISTRIBUTION",
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
                "metricKind": "DELTA",
                "valueType": "INT64",
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
                    "valueType": "STRING",
                }],
                "metricKind": "DELTA",
                "valueType": "INT64",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] bucket_options: The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
               describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.
        :param pulumi.Input[str] description: A description of this metric, which is used in documentation. The maximum length of the
               description is 8000 characters.
        :param pulumi.Input[str] filter: An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
               is used to match log entries.
        :param pulumi.Input[dict] label_extractors: A map from a label key string to an extractor expression which is used to extract data from a log
               entry field and assign as the label value. Each label key specified in the LabelDescriptor must
               have an associated extractor expression in this map. The syntax of the extractor expression is
               the same as for the valueExtractor field.
        :param pulumi.Input[dict] metric_descriptor: The metric descriptor associated with the logs-based metric.  Structure is documented below.
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

        The **bucket_options** object supports the following:

          * `explicitBuckets` (`pulumi.Input[dict]`) - Specifies a set of buckets with arbitrary widths.  Structure is documented below.
            * `bounds` (`pulumi.Input[list]`) - The values must be monotonically increasing.

          * `exponentialBuckets` (`pulumi.Input[dict]`) - Specifies an exponential sequence of buckets that have a width that is proportional to the value of
            the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.  Structure is documented below.
            * `growthFactor` (`pulumi.Input[float]`) - Must be greater than 1.
            * `numFiniteBuckets` (`pulumi.Input[float]`) - Must be greater than 0.
            * `scale` (`pulumi.Input[float]`) - Must be greater than 0.

          * `linearBuckets` (`pulumi.Input[dict]`) - Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
            Each bucket represents a constant absolute uncertainty on the specific value in the bucket.  Structure is documented below.
            * `numFiniteBuckets` (`pulumi.Input[float]`) - Must be greater than 0.
            * `offset` (`pulumi.Input[float]`) - Lower bound of the first bucket.
            * `width` (`pulumi.Input[float]`) - Must be greater than 0.

        The **metric_descriptor** object supports the following:

          * `display_name` (`pulumi.Input[str]`) - A concise name for the metric, which can be displayed in user interfaces. Use sentence case
            without an ending period, for example "Request count". This field is optional but it is
            recommended to be set for any metrics associated with user-visible concepts, such as Quota.
          * `labels` (`pulumi.Input[list]`) - The set of labels that can be used to describe a specific instance of this metric type. For
            example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
            for the HTTP response code, response_code, so you can look at latencies for successful responses
            or just for responses that failed.  Structure is documented below.
            * `description` (`pulumi.Input[str]`) - A description of this metric, which is used in documentation. The maximum length of the
              description is 8000 characters.
            * `key` (`pulumi.Input[str]`) - The label key.
            * `valueType` (`pulumi.Input[str]`) - The type of data that can be assigned to the label.

          * `metricKind` (`pulumi.Input[str]`) - Whether the metric records instantaneous values, changes to a value, etc.
            Some combinations of metricKind and valueType might not be supported.
            For counter metrics, set this to DELTA.
          * `unit` (`pulumi.Input[str]`) - The unit in which the metric value is reported. It is only applicable if the valueType is
            `INT64`, `DOUBLE`, or `DISTRIBUTION`. The supported units are a subset of
            [The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard
          * `valueType` (`pulumi.Input[str]`) - The type of data that can be assigned to the label.
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
        :param pulumi.Input[dict] bucket_options: The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
               describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.
        :param pulumi.Input[str] description: A description of this metric, which is used in documentation. The maximum length of the
               description is 8000 characters.
        :param pulumi.Input[str] filter: An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
               is used to match log entries.
        :param pulumi.Input[dict] label_extractors: A map from a label key string to an extractor expression which is used to extract data from a log
               entry field and assign as the label value. Each label key specified in the LabelDescriptor must
               have an associated extractor expression in this map. The syntax of the extractor expression is
               the same as for the valueExtractor field.
        :param pulumi.Input[dict] metric_descriptor: The metric descriptor associated with the logs-based metric.  Structure is documented below.
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

        The **bucket_options** object supports the following:

          * `explicitBuckets` (`pulumi.Input[dict]`) - Specifies a set of buckets with arbitrary widths.  Structure is documented below.
            * `bounds` (`pulumi.Input[list]`) - The values must be monotonically increasing.

          * `exponentialBuckets` (`pulumi.Input[dict]`) - Specifies an exponential sequence of buckets that have a width that is proportional to the value of
            the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.  Structure is documented below.
            * `growthFactor` (`pulumi.Input[float]`) - Must be greater than 1.
            * `numFiniteBuckets` (`pulumi.Input[float]`) - Must be greater than 0.
            * `scale` (`pulumi.Input[float]`) - Must be greater than 0.

          * `linearBuckets` (`pulumi.Input[dict]`) - Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
            Each bucket represents a constant absolute uncertainty on the specific value in the bucket.  Structure is documented below.
            * `numFiniteBuckets` (`pulumi.Input[float]`) - Must be greater than 0.
            * `offset` (`pulumi.Input[float]`) - Lower bound of the first bucket.
            * `width` (`pulumi.Input[float]`) - Must be greater than 0.

        The **metric_descriptor** object supports the following:

          * `display_name` (`pulumi.Input[str]`) - A concise name for the metric, which can be displayed in user interfaces. Use sentence case
            without an ending period, for example "Request count". This field is optional but it is
            recommended to be set for any metrics associated with user-visible concepts, such as Quota.
          * `labels` (`pulumi.Input[list]`) - The set of labels that can be used to describe a specific instance of this metric type. For
            example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
            for the HTTP response code, response_code, so you can look at latencies for successful responses
            or just for responses that failed.  Structure is documented below.
            * `description` (`pulumi.Input[str]`) - A description of this metric, which is used in documentation. The maximum length of the
              description is 8000 characters.
            * `key` (`pulumi.Input[str]`) - The label key.
            * `valueType` (`pulumi.Input[str]`) - The type of data that can be assigned to the label.

          * `metricKind` (`pulumi.Input[str]`) - Whether the metric records instantaneous values, changes to a value, etc.
            Some combinations of metricKind and valueType might not be supported.
            For counter metrics, set this to DELTA.
          * `unit` (`pulumi.Input[str]`) - The unit in which the metric value is reported. It is only applicable if the valueType is
            `INT64`, `DOUBLE`, or `DISTRIBUTION`. The supported units are a subset of
            [The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard
          * `valueType` (`pulumi.Input[str]`) - The type of data that can be assigned to the label.
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
