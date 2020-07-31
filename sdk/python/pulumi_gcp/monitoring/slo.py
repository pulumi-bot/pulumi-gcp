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

__all__ = ['Slo']


class Slo(pulumi.CustomResource):
    basic_sli: pulumi.Output[Optional['outputs.SloBasicSli']] = pulumi.output_property("basicSli")
    """
    Basic Service-Level Indicator (SLI) on a well-known service type.
    Performance will be computed on the basis of pre-defined metrics.
    SLIs are used to measure and calculate the quality of the Service's
    performance with respect to a single aspect of service quality.
    Exactly one of the following must be set:
    `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
    """
    calendar_period: pulumi.Output[Optional[str]] = pulumi.output_property("calendarPeriod")
    """
    A calendar period, semantically "since the start of the current
    <calendarPeriod>".
    """
    display_name: pulumi.Output[Optional[str]] = pulumi.output_property("displayName")
    """
    Name used for UI elements listing this SLO.
    """
    goal: pulumi.Output[float] = pulumi.output_property("goal")
    """
    The fraction of service that must be good in order for this objective
    to be met. 0 < goal <= 0.999
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The full resource name for this service. The syntax is:
    projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    request_based_sli: pulumi.Output[Optional['outputs.SloRequestBasedSli']] = pulumi.output_property("requestBasedSli")
    """
    A request-based SLI defines a SLI for which atomic units of
    service are counted directly.
    A SLI describes a good service.
    It is used to measure and calculate the quality of the Service's
    performance with respect to a single aspect of service quality.
    Exactly one of the following must be set:
    `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
    """
    rolling_period_days: pulumi.Output[Optional[float]] = pulumi.output_property("rollingPeriodDays")
    """
    A rolling time period, semantically "in the past X days".
    Must be between 1 to 30 days, inclusive.
    """
    service: pulumi.Output[str] = pulumi.output_property("service")
    """
    ID of the service to which this SLO belongs.
    """
    slo_id: pulumi.Output[str] = pulumi.output_property("sloId")
    """
    The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
    """
    windows_based_sli: pulumi.Output[Optional['outputs.SloWindowsBasedSli']] = pulumi.output_property("windowsBasedSli")
    """
    A windows-based SLI defines the criteria for time windows.
    good_service is defined based off the count of these time windows
    for which the provided service was of good quality.
    A SLI describes a good service. It is used to measure and calculate
    the quality of the Service's performance with respect to a single
    aspect of service quality.
    Exactly one of the following must be set:
    `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, basic_sli: Optional[pulumi.Input[pulumi.InputType['SloBasicSliArgs']]] = None, calendar_period: Optional[pulumi.Input[str]] = None, display_name: Optional[pulumi.Input[str]] = None, goal: Optional[pulumi.Input[float]] = None, project: Optional[pulumi.Input[str]] = None, request_based_sli: Optional[pulumi.Input[pulumi.InputType['SloRequestBasedSliArgs']]] = None, rolling_period_days: Optional[pulumi.Input[float]] = None, service: Optional[pulumi.Input[str]] = None, slo_id: Optional[pulumi.Input[str]] = None, windows_based_sli: Optional[pulumi.Input[pulumi.InputType['SloWindowsBasedSliArgs']]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A Service-Level Objective (SLO) describes the level of desired good
        service. It consists of a service-level indicator (SLI), a performance
        goal, and a period over which the objective is to be evaluated against
        that goal. The SLO can use SLIs defined in a number of different manners.
        Typical SLOs might include "99% of requests in each rolling week have
        latency below 200 milliseconds" or "99.5% of requests in each calendar
        month return successfully."

        To get more information about Slo, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services.serviceLevelObjectives)
        * How-to Guides
            * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
            * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)

        ## Example Usage
        ### Monitoring Slo Appengine

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.monitoring.get_app_engine_service(module_id="default")
        appeng_slo = gcp.monitoring.Slo("appengSlo",
            service=default.service_id,
            slo_id="ae-slo",
            display_name="Test SLO for App Engine",
            goal=0.9,
            calendar_period="DAY",
            basic_sli={
                "latency": {
                    "threshold": "1s",
                },
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SloBasicSliArgs']] basic_sli: Basic Service-Level Indicator (SLI) on a well-known service type.
               Performance will be computed on the basis of pre-defined metrics.
               SLIs are used to measure and calculate the quality of the Service's
               performance with respect to a single aspect of service quality.
               Exactly one of the following must be set:
               `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
        :param pulumi.Input[str] calendar_period: A calendar period, semantically "since the start of the current
               <calendarPeriod>".
        :param pulumi.Input[str] display_name: Name used for UI elements listing this SLO.
        :param pulumi.Input[float] goal: The fraction of service that must be good in order for this objective
               to be met. 0 < goal <= 0.999
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['SloRequestBasedSliArgs']] request_based_sli: A request-based SLI defines a SLI for which atomic units of
               service are counted directly.
               A SLI describes a good service.
               It is used to measure and calculate the quality of the Service's
               performance with respect to a single aspect of service quality.
               Exactly one of the following must be set:
               `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
        :param pulumi.Input[float] rolling_period_days: A rolling time period, semantically "in the past X days".
               Must be between 1 to 30 days, inclusive.
        :param pulumi.Input[str] service: ID of the service to which this SLO belongs.
        :param pulumi.Input[str] slo_id: The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
        :param pulumi.Input[pulumi.InputType['SloWindowsBasedSliArgs']] windows_based_sli: A windows-based SLI defines the criteria for time windows.
               good_service is defined based off the count of these time windows
               for which the provided service was of good quality.
               A SLI describes a good service. It is used to measure and calculate
               the quality of the Service's performance with respect to a single
               aspect of service quality.
               Exactly one of the following must be set:
               `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
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

            __props__['basic_sli'] = basic_sli
            __props__['calendar_period'] = calendar_period
            __props__['display_name'] = display_name
            if goal is None:
                raise TypeError("Missing required property 'goal'")
            __props__['goal'] = goal
            __props__['project'] = project
            __props__['request_based_sli'] = request_based_sli
            __props__['rolling_period_days'] = rolling_period_days
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['slo_id'] = slo_id
            __props__['windows_based_sli'] = windows_based_sli
            __props__['name'] = None
        super(Slo, __self__).__init__(
            'gcp:monitoring/slo:Slo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, basic_sli: Optional[pulumi.Input[pulumi.InputType['SloBasicSliArgs']]] = None, calendar_period: Optional[pulumi.Input[str]] = None, display_name: Optional[pulumi.Input[str]] = None, goal: Optional[pulumi.Input[float]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, request_based_sli: Optional[pulumi.Input[pulumi.InputType['SloRequestBasedSliArgs']]] = None, rolling_period_days: Optional[pulumi.Input[float]] = None, service: Optional[pulumi.Input[str]] = None, slo_id: Optional[pulumi.Input[str]] = None, windows_based_sli: Optional[pulumi.Input[pulumi.InputType['SloWindowsBasedSliArgs']]] = None) -> 'Slo':
        """
        Get an existing Slo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SloBasicSliArgs']] basic_sli: Basic Service-Level Indicator (SLI) on a well-known service type.
               Performance will be computed on the basis of pre-defined metrics.
               SLIs are used to measure and calculate the quality of the Service's
               performance with respect to a single aspect of service quality.
               Exactly one of the following must be set:
               `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
        :param pulumi.Input[str] calendar_period: A calendar period, semantically "since the start of the current
               <calendarPeriod>".
        :param pulumi.Input[str] display_name: Name used for UI elements listing this SLO.
        :param pulumi.Input[float] goal: The fraction of service that must be good in order for this objective
               to be met. 0 < goal <= 0.999
        :param pulumi.Input[str] name: The full resource name for this service. The syntax is:
               projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['SloRequestBasedSliArgs']] request_based_sli: A request-based SLI defines a SLI for which atomic units of
               service are counted directly.
               A SLI describes a good service.
               It is used to measure and calculate the quality of the Service's
               performance with respect to a single aspect of service quality.
               Exactly one of the following must be set:
               `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
        :param pulumi.Input[float] rolling_period_days: A rolling time period, semantically "in the past X days".
               Must be between 1 to 30 days, inclusive.
        :param pulumi.Input[str] service: ID of the service to which this SLO belongs.
        :param pulumi.Input[str] slo_id: The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
        :param pulumi.Input[pulumi.InputType['SloWindowsBasedSliArgs']] windows_based_sli: A windows-based SLI defines the criteria for time windows.
               good_service is defined based off the count of these time windows
               for which the provided service was of good quality.
               A SLI describes a good service. It is used to measure and calculate
               the quality of the Service's performance with respect to a single
               aspect of service quality.
               Exactly one of the following must be set:
               `basic_sli`, `request_based_sli`, `windows_based_sli`  Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["basic_sli"] = basic_sli
        __props__["calendar_period"] = calendar_period
        __props__["display_name"] = display_name
        __props__["goal"] = goal
        __props__["name"] = name
        __props__["project"] = project
        __props__["request_based_sli"] = request_based_sli
        __props__["rolling_period_days"] = rolling_period_days
        __props__["service"] = service
        __props__["slo_id"] = slo_id
        __props__["windows_based_sli"] = windows_based_sli
        return Slo(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

