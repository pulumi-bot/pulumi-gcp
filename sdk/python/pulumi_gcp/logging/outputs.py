# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs

@pulumi.output_type
class BillingAccountSinkBigqueryOptions(dict):
    use_partitioned_tables: bool = pulumi.output_property("usePartitionedTables")
    """
    Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
    By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
    tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
    has to be used instead. In both cases, tables are sharded based on UTC timezone.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FolderSinkBigqueryOptions(dict):
    use_partitioned_tables: bool = pulumi.output_property("usePartitionedTables")
    """
    Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
    By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
    tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
    has to be used instead. In both cases, tables are sharded based on UTC timezone.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricBucketOptions(dict):
    explicit_buckets: Optional['outputs.MetricBucketOptionsExplicitBuckets'] = pulumi.output_property("explicitBuckets")
    """
    Specifies a set of buckets with arbitrary widths.  Structure is documented below.
    """
    exponential_buckets: Optional['outputs.MetricBucketOptionsExponentialBuckets'] = pulumi.output_property("exponentialBuckets")
    """
    Specifies an exponential sequence of buckets that have a width that is proportional to the value of
    the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.  Structure is documented below.
    """
    linear_buckets: Optional['outputs.MetricBucketOptionsLinearBuckets'] = pulumi.output_property("linearBuckets")
    """
    Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
    Each bucket represents a constant absolute uncertainty on the specific value in the bucket.  Structure is documented below.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricBucketOptionsExplicitBuckets(dict):
    bounds: List[float] = pulumi.output_property("bounds")
    """
    The values must be monotonically increasing.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricBucketOptionsExponentialBuckets(dict):
    growth_factor: Optional[float] = pulumi.output_property("growthFactor")
    """
    Must be greater than 1.
    """
    num_finite_buckets: Optional[float] = pulumi.output_property("numFiniteBuckets")
    """
    Must be greater than 0.
    """
    scale: Optional[float] = pulumi.output_property("scale")
    """
    Must be greater than 0.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricBucketOptionsLinearBuckets(dict):
    num_finite_buckets: Optional[float] = pulumi.output_property("numFiniteBuckets")
    """
    Must be greater than 0.
    """
    offset: Optional[float] = pulumi.output_property("offset")
    """
    Lower bound of the first bucket.
    """
    width: Optional[float] = pulumi.output_property("width")
    """
    Must be greater than 0.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricMetricDescriptor(dict):
    display_name: Optional[str] = pulumi.output_property("displayName")
    """
    A concise name for the metric, which can be displayed in user interfaces. Use sentence case
    without an ending period, for example "Request count". This field is optional but it is
    recommended to be set for any metrics associated with user-visible concepts, such as Quota.
    """
    labels: Optional[List['outputs.MetricMetricDescriptorLabel']] = pulumi.output_property("labels")
    """
    The set of labels that can be used to describe a specific instance of this metric type. For
    example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
    for the HTTP response code, response_code, so you can look at latencies for successful responses
    or just for responses that failed.  Structure is documented below.
    """
    metric_kind: str = pulumi.output_property("metricKind")
    """
    Whether the metric records instantaneous values, changes to a value, etc.
    Some combinations of metricKind and valueType might not be supported.
    For counter metrics, set this to DELTA.
    """
    unit: Optional[str] = pulumi.output_property("unit")
    """
    The unit in which the metric value is reported. It is only applicable if the valueType is
    `INT64`, `DOUBLE`, or `DISTRIBUTION`. The supported units are a subset of
    [The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard
    """
    value_type: str = pulumi.output_property("valueType")
    """
    The type of data that can be assigned to the label.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricMetricDescriptorLabel(dict):
    description: Optional[str] = pulumi.output_property("description")
    """
    A description of this metric, which is used in documentation. The maximum length of the
    description is 8000 characters.
    """
    key: str = pulumi.output_property("key")
    """
    The label key.
    """
    value_type: Optional[str] = pulumi.output_property("valueType")
    """
    The type of data that can be assigned to the label.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationSinkBigqueryOptions(dict):
    use_partitioned_tables: bool = pulumi.output_property("usePartitionedTables")
    """
    Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
    By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
    tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
    has to be used instead. In both cases, tables are sharded based on UTC timezone.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProjectSinkBigqueryOptions(dict):
    use_partitioned_tables: bool = pulumi.output_property("usePartitionedTables")
    """
    Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
    By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
    tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
    has to be used instead. In both cases, tables are sharded based on UTC timezone.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


