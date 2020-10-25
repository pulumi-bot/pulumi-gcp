# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'BillingAccountSinkBigqueryOptions',
    'BillingAccountSinkExclusion',
    'FolderSinkBigqueryOptions',
    'FolderSinkExclusion',
    'MetricBucketOptions',
    'MetricBucketOptionsExplicitBuckets',
    'MetricBucketOptionsExponentialBuckets',
    'MetricBucketOptionsLinearBuckets',
    'MetricMetricDescriptor',
    'MetricMetricDescriptorLabel',
    'OrganizationSinkBigqueryOptions',
    'OrganizationSinkExclusion',
    'ProjectSinkBigqueryOptions',
    'ProjectSinkExclusion',
]

@pulumi.output_type
class BillingAccountSinkBigqueryOptions(dict):
    def __init__(__self__, *,
                 use_partitioned_tables: bool):
        """
        :param bool use_partitioned_tables: Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
               By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
               tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
               has to be used instead. In both cases, tables are sharded based on UTC timezone.
        """
        pulumi.set(__self__, "use_partitioned_tables", use_partitioned_tables)

    @property
    @pulumi.getter(name="usePartitionedTables")
    def use_partitioned_tables(self) -> bool:
        """
        Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
        By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
        tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
        has to be used instead. In both cases, tables are sharded based on UTC timezone.
        """
        return pulumi.get(self, "use_partitioned_tables")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BillingAccountSinkExclusion(dict):
    def __init__(__self__, *,
                 filter: str,
                 name: str,
                 description: Optional[str] = None,
                 disabled: Optional[bool] = None):
        """
        :param str filter: The filter to apply when exporting logs. Only log entries that match the filter are exported.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param str name: The name of the logging sink.
        """
        pulumi.set(__self__, "filter", filter)
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def filter(self) -> str:
        """
        The filter to apply when exporting logs. Only log entries that match the filter are exported.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the logging sink.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> Optional[bool]:
        return pulumi.get(self, "disabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FolderSinkBigqueryOptions(dict):
    def __init__(__self__, *,
                 use_partitioned_tables: bool):
        """
        :param bool use_partitioned_tables: Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
               By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
               tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
               has to be used instead. In both cases, tables are sharded based on UTC timezone.
        """
        pulumi.set(__self__, "use_partitioned_tables", use_partitioned_tables)

    @property
    @pulumi.getter(name="usePartitionedTables")
    def use_partitioned_tables(self) -> bool:
        """
        Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
        By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
        tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
        has to be used instead. In both cases, tables are sharded based on UTC timezone.
        """
        return pulumi.get(self, "use_partitioned_tables")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FolderSinkExclusion(dict):
    def __init__(__self__, *,
                 filter: str,
                 name: str,
                 description: Optional[str] = None,
                 disabled: Optional[bool] = None):
        """
        :param str filter: The filter to apply when exporting logs. Only log entries that match the filter are exported.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param str name: The name of the logging sink.
        """
        pulumi.set(__self__, "filter", filter)
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def filter(self) -> str:
        """
        The filter to apply when exporting logs. Only log entries that match the filter are exported.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the logging sink.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> Optional[bool]:
        return pulumi.get(self, "disabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricBucketOptions(dict):
    def __init__(__self__, *,
                 explicit_buckets: Optional['outputs.MetricBucketOptionsExplicitBuckets'] = None,
                 exponential_buckets: Optional['outputs.MetricBucketOptionsExponentialBuckets'] = None,
                 linear_buckets: Optional['outputs.MetricBucketOptionsLinearBuckets'] = None):
        """
        :param 'MetricBucketOptionsExplicitBucketsArgs' explicit_buckets: Specifies a set of buckets with arbitrary widths.
               Structure is documented below.
        :param 'MetricBucketOptionsExponentialBucketsArgs' exponential_buckets: Specifies an exponential sequence of buckets that have a width that is proportional to the value of
               the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.
               Structure is documented below.
        :param 'MetricBucketOptionsLinearBucketsArgs' linear_buckets: Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
               Each bucket represents a constant absolute uncertainty on the specific value in the bucket.
               Structure is documented below.
        """
        if explicit_buckets is not None:
            pulumi.set(__self__, "explicit_buckets", explicit_buckets)
        if exponential_buckets is not None:
            pulumi.set(__self__, "exponential_buckets", exponential_buckets)
        if linear_buckets is not None:
            pulumi.set(__self__, "linear_buckets", linear_buckets)

    @property
    @pulumi.getter(name="explicitBuckets")
    def explicit_buckets(self) -> Optional['outputs.MetricBucketOptionsExplicitBuckets']:
        """
        Specifies a set of buckets with arbitrary widths.
        Structure is documented below.
        """
        return pulumi.get(self, "explicit_buckets")

    @property
    @pulumi.getter(name="exponentialBuckets")
    def exponential_buckets(self) -> Optional['outputs.MetricBucketOptionsExponentialBuckets']:
        """
        Specifies an exponential sequence of buckets that have a width that is proportional to the value of
        the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.
        Structure is documented below.
        """
        return pulumi.get(self, "exponential_buckets")

    @property
    @pulumi.getter(name="linearBuckets")
    def linear_buckets(self) -> Optional['outputs.MetricBucketOptionsLinearBuckets']:
        """
        Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
        Each bucket represents a constant absolute uncertainty on the specific value in the bucket.
        Structure is documented below.
        """
        return pulumi.get(self, "linear_buckets")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricBucketOptionsExplicitBuckets(dict):
    def __init__(__self__, *,
                 bounds: Sequence[float]):
        """
        :param Sequence[float] bounds: The values must be monotonically increasing.
        """
        pulumi.set(__self__, "bounds", bounds)

    @property
    @pulumi.getter
    def bounds(self) -> Sequence[float]:
        """
        The values must be monotonically increasing.
        """
        return pulumi.get(self, "bounds")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricBucketOptionsExponentialBuckets(dict):
    def __init__(__self__, *,
                 growth_factor: Optional[float] = None,
                 num_finite_buckets: Optional[int] = None,
                 scale: Optional[float] = None):
        """
        :param float growth_factor: Must be greater than 1.
        :param int num_finite_buckets: Must be greater than 0.
        :param float scale: Must be greater than 0.
        """
        if growth_factor is not None:
            pulumi.set(__self__, "growth_factor", growth_factor)
        if num_finite_buckets is not None:
            pulumi.set(__self__, "num_finite_buckets", num_finite_buckets)
        if scale is not None:
            pulumi.set(__self__, "scale", scale)

    @property
    @pulumi.getter(name="growthFactor")
    def growth_factor(self) -> Optional[float]:
        """
        Must be greater than 1.
        """
        return pulumi.get(self, "growth_factor")

    @property
    @pulumi.getter(name="numFiniteBuckets")
    def num_finite_buckets(self) -> Optional[int]:
        """
        Must be greater than 0.
        """
        return pulumi.get(self, "num_finite_buckets")

    @property
    @pulumi.getter
    def scale(self) -> Optional[float]:
        """
        Must be greater than 0.
        """
        return pulumi.get(self, "scale")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricBucketOptionsLinearBuckets(dict):
    def __init__(__self__, *,
                 num_finite_buckets: Optional[int] = None,
                 offset: Optional[float] = None,
                 width: Optional[int] = None):
        """
        :param int num_finite_buckets: Must be greater than 0.
        :param float offset: Lower bound of the first bucket.
        :param int width: Must be greater than 0.
        """
        if num_finite_buckets is not None:
            pulumi.set(__self__, "num_finite_buckets", num_finite_buckets)
        if offset is not None:
            pulumi.set(__self__, "offset", offset)
        if width is not None:
            pulumi.set(__self__, "width", width)

    @property
    @pulumi.getter(name="numFiniteBuckets")
    def num_finite_buckets(self) -> Optional[int]:
        """
        Must be greater than 0.
        """
        return pulumi.get(self, "num_finite_buckets")

    @property
    @pulumi.getter
    def offset(self) -> Optional[float]:
        """
        Lower bound of the first bucket.
        """
        return pulumi.get(self, "offset")

    @property
    @pulumi.getter
    def width(self) -> Optional[int]:
        """
        Must be greater than 0.
        """
        return pulumi.get(self, "width")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricMetricDescriptor(dict):
    def __init__(__self__, *,
                 metric_kind: str,
                 value_type: str,
                 display_name: Optional[str] = None,
                 labels: Optional[Sequence['outputs.MetricMetricDescriptorLabel']] = None,
                 unit: Optional[str] = None):
        """
        :param str metric_kind: Whether the metric records instantaneous values, changes to a value, etc.
               Some combinations of metricKind and valueType might not be supported.
               For counter metrics, set this to DELTA.
               Possible values are `DELTA`, `GAUGE`, and `CUMULATIVE`.
        :param str value_type: The type of data that can be assigned to the label.
               Default value is `STRING`.
               Possible values are `BOOL`, `INT64`, and `STRING`.
        :param str display_name: A concise name for the metric, which can be displayed in user interfaces. Use sentence case
               without an ending period, for example "Request count". This field is optional but it is
               recommended to be set for any metrics associated with user-visible concepts, such as Quota.
        :param Sequence['MetricMetricDescriptorLabelArgs'] labels: The set of labels that can be used to describe a specific instance of this metric type. For
               example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
               for the HTTP response code, response_code, so you can look at latencies for successful responses
               or just for responses that failed.
               Structure is documented below.
        :param str unit: The unit in which the metric value is reported. It is only applicable if the valueType is
               `INT64`, `DOUBLE`, or `DISTRIBUTION`. The supported units are a subset of
               [The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard
        """
        pulumi.set(__self__, "metric_kind", metric_kind)
        pulumi.set(__self__, "value_type", value_type)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if unit is not None:
            pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter(name="metricKind")
    def metric_kind(self) -> str:
        """
        Whether the metric records instantaneous values, changes to a value, etc.
        Some combinations of metricKind and valueType might not be supported.
        For counter metrics, set this to DELTA.
        Possible values are `DELTA`, `GAUGE`, and `CUMULATIVE`.
        """
        return pulumi.get(self, "metric_kind")

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> str:
        """
        The type of data that can be assigned to the label.
        Default value is `STRING`.
        Possible values are `BOOL`, `INT64`, and `STRING`.
        """
        return pulumi.get(self, "value_type")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        A concise name for the metric, which can be displayed in user interfaces. Use sentence case
        without an ending period, for example "Request count". This field is optional but it is
        recommended to be set for any metrics associated with user-visible concepts, such as Quota.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence['outputs.MetricMetricDescriptorLabel']]:
        """
        The set of labels that can be used to describe a specific instance of this metric type. For
        example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
        for the HTTP response code, response_code, so you can look at latencies for successful responses
        or just for responses that failed.
        Structure is documented below.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def unit(self) -> Optional[str]:
        """
        The unit in which the metric value is reported. It is only applicable if the valueType is
        `INT64`, `DOUBLE`, or `DISTRIBUTION`. The supported units are a subset of
        [The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard
        """
        return pulumi.get(self, "unit")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricMetricDescriptorLabel(dict):
    def __init__(__self__, *,
                 key: str,
                 description: Optional[str] = None,
                 value_type: Optional[str] = None):
        """
        :param str key: The label key.
        :param str description: A description of this metric, which is used in documentation. The maximum length of the
               description is 8000 characters.
        :param str value_type: The type of data that can be assigned to the label.
               Default value is `STRING`.
               Possible values are `BOOL`, `INT64`, and `STRING`.
        """
        pulumi.set(__self__, "key", key)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if value_type is not None:
            pulumi.set(__self__, "value_type", value_type)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The label key.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of this metric, which is used in documentation. The maximum length of the
        description is 8000 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> Optional[str]:
        """
        The type of data that can be assigned to the label.
        Default value is `STRING`.
        Possible values are `BOOL`, `INT64`, and `STRING`.
        """
        return pulumi.get(self, "value_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationSinkBigqueryOptions(dict):
    def __init__(__self__, *,
                 use_partitioned_tables: bool):
        """
        :param bool use_partitioned_tables: Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
               By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
               tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
               has to be used instead. In both cases, tables are sharded based on UTC timezone.
        """
        pulumi.set(__self__, "use_partitioned_tables", use_partitioned_tables)

    @property
    @pulumi.getter(name="usePartitionedTables")
    def use_partitioned_tables(self) -> bool:
        """
        Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
        By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
        tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
        has to be used instead. In both cases, tables are sharded based on UTC timezone.
        """
        return pulumi.get(self, "use_partitioned_tables")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationSinkExclusion(dict):
    def __init__(__self__, *,
                 filter: str,
                 name: str,
                 description: Optional[str] = None,
                 disabled: Optional[bool] = None):
        """
        :param str filter: The filter to apply when exporting logs. Only log entries that match the filter are exported.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param str name: The name of the logging sink.
        """
        pulumi.set(__self__, "filter", filter)
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def filter(self) -> str:
        """
        The filter to apply when exporting logs. Only log entries that match the filter are exported.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the logging sink.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> Optional[bool]:
        return pulumi.get(self, "disabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProjectSinkBigqueryOptions(dict):
    def __init__(__self__, *,
                 use_partitioned_tables: bool):
        """
        :param bool use_partitioned_tables: Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
               By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
               tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
               has to be used instead. In both cases, tables are sharded based on UTC timezone.
        """
        pulumi.set(__self__, "use_partitioned_tables", use_partitioned_tables)

    @property
    @pulumi.getter(name="usePartitionedTables")
    def use_partitioned_tables(self) -> bool:
        """
        Whether to use [BigQuery's partition tables](https://cloud.google.com/bigquery/docs/partitioned-tables).
        By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
        tables the date suffix is no longer present and [special query syntax](https://cloud.google.com/bigquery/docs/querying-partitioned-tables)
        has to be used instead. In both cases, tables are sharded based on UTC timezone.
        """
        return pulumi.get(self, "use_partitioned_tables")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProjectSinkExclusion(dict):
    def __init__(__self__, *,
                 filter: str,
                 name: str,
                 description: Optional[str] = None,
                 disabled: Optional[bool] = None):
        """
        :param str filter: An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param str name: A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        :param str description: A description of this exclusion.
        :param bool disabled: If set to True, then this exclusion is disabled and it does not exclude any log entries.
        """
        pulumi.set(__self__, "filter", filter)
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def filter(self) -> str:
        """
        An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of this exclusion.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> Optional[bool]:
        """
        If set to True, then this exclusion is disabled and it does not exclude any log entries.
        """
        return pulumi.get(self, "disabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


