# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['MetricDescriptor']


class MetricDescriptor(pulumi.CustomResource):
    description: pulumi.Output[str] = pulumi.output_property("description")
    """
    A human-readable description for the label.
    """
    display_name: pulumi.Output[str] = pulumi.output_property("displayName")
    """
    A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
    """
    labels: pulumi.Output[Optional[List['outputs.MetricDescriptorLabel']]] = pulumi.output_property("labels")
    """
    The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.  Structure is documented below.
    """
    launch_stage: pulumi.Output[Optional[str]] = pulumi.output_property("launchStage")
    """
    The launch stage of the metric definition.
    """
    metadata: pulumi.Output[Optional['outputs.MetricDescriptorMetadata']] = pulumi.output_property("metadata")
    """
    Metadata which can be used to guide usage of the metric.  Structure is documented below.
    """
    metric_kind: pulumi.Output[str] = pulumi.output_property("metricKind")
    """
    Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
    """
    monitored_resource_types: pulumi.Output[List[str]] = pulumi.output_property("monitoredResourceTypes")
    """
    If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
    is associated with this metric type can only be associated with one of the monitored resource types listed here. This
    field allows time series to be associated with the intersection of this metric type and the monitored resource types in
    this list.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The resource name of the metric descriptor.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    type: pulumi.Output[str] = pulumi.output_property("type")
    """
    The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
    """
    unit: pulumi.Output[Optional[str]] = pulumi.output_property("unit")
    """
    The units in which the metric value is reported. It is only applicable if the
    valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
    the stored metric values.
    Different systems may scale the values to be more easily displayed (so a value of
    0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
    3.5MBy). However, if the unit is KBy, then the value of the metric is always in
    thousands of bytes, no matter how it may be displayed.
    If you want a custom metric to record the exact number of CPU-seconds used by a job,
    you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
    1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
    12005.
    Alternatively, if you want a custom metric to record data in a more granular way, you
    can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
    12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
    The supported units are a subset of The Unified Code for Units of Measure standard.
    More info can be found in the API documentation
    (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
    """
    value_type: pulumi.Output[str] = pulumi.output_property("valueType")
    """
    The type of data that can be assigned to the label.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, description: Optional[pulumi.Input[str]] = None, display_name: Optional[pulumi.Input[str]] = None, labels: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MetricDescriptorLabelArgs']]]]] = None, launch_stage: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input[pulumi.InputType['MetricDescriptorMetadataArgs']]] = None, metric_kind: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, type: Optional[pulumi.Input[str]] = None, unit: Optional[pulumi.Input[str]] = None, value_type: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Defines a metric type and its schema. Once a metric descriptor is created, deleting or altering it stops data collection and makes the metric type's existing data unusable.

        To get more information about MetricDescriptor, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/monitoring/custom-metrics/)

        ## Example Usage
        ### Monitoring Metric Descriptor Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.monitoring.MetricDescriptor("basic",
            description="Daily sales records from all branch stores.",
            display_name="metric-descriptor",
            labels=[{
                "description": "The ID of the store.",
                "key": "store_id",
                "value_type": "STRING",
            }],
            launch_stage="BETA",
            metadata={
                "ingestDelay": "30s",
                "samplePeriod": "60s",
            },
            metric_kind="GAUGE",
            type="custom.googleapis.com/stores/daily_sales",
            unit="{USD}",
            value_type="DOUBLE")
        ```
        ### Monitoring Metric Descriptor Alert

        ```python
        import pulumi
        import pulumi_gcp as gcp

        with_alert = gcp.monitoring.MetricDescriptor("withAlert",
            description="Daily sales records from all branch stores.",
            display_name="metric-descriptor",
            metric_kind="GAUGE",
            type="custom.googleapis.com/stores/daily_sales",
            unit="{USD}",
            value_type="DOUBLE")
        alert_policy = gcp.monitoring.AlertPolicy("alertPolicy",
            combiner="OR",
            conditions=[{
                "conditionThreshold": {
                    "comparison": "COMPARISON_GT",
                    "duration": "60s",
                    "filter": with_alert.type.apply(lambda type: f"metric.type=\"{type}\" AND resource.type=\"gce_instance\""),
                },
                "display_name": "test condition",
            }],
            display_name="metric-descriptor")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description for the label.
        :param pulumi.Input[str] display_name: A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['MetricDescriptorLabelArgs']]]] labels: The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.  Structure is documented below.
        :param pulumi.Input[str] launch_stage: The launch stage of the metric definition.
        :param pulumi.Input[pulumi.InputType['MetricDescriptorMetadataArgs']] metadata: Metadata which can be used to guide usage of the metric.  Structure is documented below.
        :param pulumi.Input[str] metric_kind: Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] type: The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
        :param pulumi.Input[str] unit: The units in which the metric value is reported. It is only applicable if the
               valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
               the stored metric values.
               Different systems may scale the values to be more easily displayed (so a value of
               0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
               3.5MBy). However, if the unit is KBy, then the value of the metric is always in
               thousands of bytes, no matter how it may be displayed.
               If you want a custom metric to record the exact number of CPU-seconds used by a job,
               you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
               1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
               12005.
               Alternatively, if you want a custom metric to record data in a more granular way, you
               can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
               12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
               The supported units are a subset of The Unified Code for Units of Measure standard.
               More info can be found in the API documentation
               (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
        :param pulumi.Input[str] value_type: The type of data that can be assigned to the label.
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

            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['labels'] = labels
            __props__['launch_stage'] = launch_stage
            __props__['metadata'] = metadata
            if metric_kind is None:
                raise TypeError("Missing required property 'metric_kind'")
            __props__['metric_kind'] = metric_kind
            __props__['project'] = project
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['unit'] = unit
            if value_type is None:
                raise TypeError("Missing required property 'value_type'")
            __props__['value_type'] = value_type
            __props__['monitored_resource_types'] = None
            __props__['name'] = None
        super(MetricDescriptor, __self__).__init__(
            'gcp:monitoring/metricDescriptor:MetricDescriptor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, description: Optional[pulumi.Input[str]] = None, display_name: Optional[pulumi.Input[str]] = None, labels: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MetricDescriptorLabelArgs']]]]] = None, launch_stage: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input[pulumi.InputType['MetricDescriptorMetadataArgs']]] = None, metric_kind: Optional[pulumi.Input[str]] = None, monitored_resource_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, type: Optional[pulumi.Input[str]] = None, unit: Optional[pulumi.Input[str]] = None, value_type: Optional[pulumi.Input[str]] = None) -> 'MetricDescriptor':
        """
        Get an existing MetricDescriptor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description for the label.
        :param pulumi.Input[str] display_name: A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['MetricDescriptorLabelArgs']]]] labels: The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.  Structure is documented below.
        :param pulumi.Input[str] launch_stage: The launch stage of the metric definition.
        :param pulumi.Input[pulumi.InputType['MetricDescriptorMetadataArgs']] metadata: Metadata which can be used to guide usage of the metric.  Structure is documented below.
        :param pulumi.Input[str] metric_kind: Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
        :param pulumi.Input[List[pulumi.Input[str]]] monitored_resource_types: If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that
               is associated with this metric type can only be associated with one of the monitored resource types listed here. This
               field allows time series to be associated with the intersection of this metric type and the monitored resource types in
               this list.
        :param pulumi.Input[str] name: The resource name of the metric descriptor.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] type: The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
        :param pulumi.Input[str] unit: The units in which the metric value is reported. It is only applicable if the
               valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
               the stored metric values.
               Different systems may scale the values to be more easily displayed (so a value of
               0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
               3.5MBy). However, if the unit is KBy, then the value of the metric is always in
               thousands of bytes, no matter how it may be displayed.
               If you want a custom metric to record the exact number of CPU-seconds used by a job,
               you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
               1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
               12005.
               Alternatively, if you want a custom metric to record data in a more granular way, you
               can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
               12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
               The supported units are a subset of The Unified Code for Units of Measure standard.
               More info can be found in the API documentation
               (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
        :param pulumi.Input[str] value_type: The type of data that can be assigned to the label.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["labels"] = labels
        __props__["launch_stage"] = launch_stage
        __props__["metadata"] = metadata
        __props__["metric_kind"] = metric_kind
        __props__["monitored_resource_types"] = monitored_resource_types
        __props__["name"] = name
        __props__["project"] = project
        __props__["type"] = type
        __props__["unit"] = unit
        __props__["value_type"] = value_type
        return MetricDescriptor(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

