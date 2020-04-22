# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Slo(pulumi.CustomResource):
    basic_sli: pulumi.Output[dict]
    """
    -
    (Required)
    Basic Service-Level Indicator (SLI) on a well-known service type.
    Performance will be computed on the basis of pre-defined metrics.
    SLIs are used to measure and calculate the quality of the Service's
    performance with respect to a single aspect of service quality.  Structure is documented below.

      * `latency` (`dict`) - -
        (Required)
        Parameters for a latency threshold SLI.  Structure is documented below.
        * `threshold` (`str`) - -
          (Required)
          A duration string, e.g. 10s.
          Good service is defined to be the count of requests made to
          this service that return in no more than threshold.

      * `locations` (`list`) - -
        (Optional)
        An optional set of locations to which this SLI is relevant.
        Telemetry from other locations will not be used to calculate
        performance for this SLI. If omitted, this SLI applies to all
        locations in which the Service has activity. For service types
        that don't support breaking down by location, setting this
        field will result in an error.
      * `methods` (`list`) - -
        (Optional)
        An optional set of RPCs to which this SLI is relevant.
        Telemetry from other methods will not be used to calculate
        performance for this SLI. If omitted, this SLI applies to all
        the Service's methods. For service types that don't support
        breaking down by method, setting this field will result in an
        error.
      * `versions` (`list`) - -
        (Optional)
        The set of API versions to which this SLI is relevant.
        Telemetry from other API versions will not be used to
        calculate performance for this SLI. If omitted,
        this SLI applies to all API versions. For service types
        that don't support breaking down by version, setting this
        field will result in an error.
    """
    calendar_period: pulumi.Output[str]
    """
    -
    (Optional)
    A calendar period, semantically "since the start of the current
    <calendarPeriod>".
    """
    display_name: pulumi.Output[str]
    """
    -
    (Optional)
    Name used for UI elements listing this SLO.
    """
    goal: pulumi.Output[float]
    """
    -
    (Required)
    The fraction of service that must be good in order for this objective
    to be met. 0 < goal <= 0.999
    """
    name: pulumi.Output[str]
    """
    The full resource name for this service. The syntax is:
    projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    rolling_period_days: pulumi.Output[float]
    """
    -
    (Optional)
    A rolling time period, semantically "in the past X days".
    Must be between 1 to 30 days, inclusive.
    """
    service: pulumi.Output[str]
    """
    -
    (Required)
    ID of the service to which this SLO belongs.
    """
    slo_id: pulumi.Output[str]
    """
    -
    (Optional)
    The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
    """
    def __init__(__self__, resource_name, opts=None, basic_sli=None, calendar_period=None, display_name=None, goal=None, project=None, rolling_period_days=None, service=None, slo_id=None, __props__=None, __name__=None, __opts__=None):
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

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] basic_sli: -
               (Required)
               Basic Service-Level Indicator (SLI) on a well-known service type.
               Performance will be computed on the basis of pre-defined metrics.
               SLIs are used to measure and calculate the quality of the Service's
               performance with respect to a single aspect of service quality.  Structure is documented below.
        :param pulumi.Input[str] calendar_period: -
               (Optional)
               A calendar period, semantically "since the start of the current
               <calendarPeriod>".
        :param pulumi.Input[str] display_name: -
               (Optional)
               Name used for UI elements listing this SLO.
        :param pulumi.Input[float] goal: -
               (Required)
               The fraction of service that must be good in order for this objective
               to be met. 0 < goal <= 0.999
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[float] rolling_period_days: -
               (Optional)
               A rolling time period, semantically "in the past X days".
               Must be between 1 to 30 days, inclusive.
        :param pulumi.Input[str] service: -
               (Required)
               ID of the service to which this SLO belongs.
        :param pulumi.Input[str] slo_id: -
               (Optional)
               The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.

        The **basic_sli** object supports the following:

          * `latency` (`pulumi.Input[dict]`) - -
            (Required)
            Parameters for a latency threshold SLI.  Structure is documented below.
            * `threshold` (`pulumi.Input[str]`) - -
              (Required)
              A duration string, e.g. 10s.
              Good service is defined to be the count of requests made to
              this service that return in no more than threshold.

          * `locations` (`pulumi.Input[list]`) - -
            (Optional)
            An optional set of locations to which this SLI is relevant.
            Telemetry from other locations will not be used to calculate
            performance for this SLI. If omitted, this SLI applies to all
            locations in which the Service has activity. For service types
            that don't support breaking down by location, setting this
            field will result in an error.
          * `methods` (`pulumi.Input[list]`) - -
            (Optional)
            An optional set of RPCs to which this SLI is relevant.
            Telemetry from other methods will not be used to calculate
            performance for this SLI. If omitted, this SLI applies to all
            the Service's methods. For service types that don't support
            breaking down by method, setting this field will result in an
            error.
          * `versions` (`pulumi.Input[list]`) - -
            (Optional)
            The set of API versions to which this SLI is relevant.
            Telemetry from other API versions will not be used to
            calculate performance for this SLI. If omitted,
            this SLI applies to all API versions. For service types
            that don't support breaking down by version, setting this
            field will result in an error.
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

            if basic_sli is None:
                raise TypeError("Missing required property 'basic_sli'")
            __props__['basic_sli'] = basic_sli
            __props__['calendar_period'] = calendar_period
            __props__['display_name'] = display_name
            if goal is None:
                raise TypeError("Missing required property 'goal'")
            __props__['goal'] = goal
            __props__['project'] = project
            __props__['rolling_period_days'] = rolling_period_days
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['slo_id'] = slo_id
            __props__['name'] = None
        super(Slo, __self__).__init__(
            'gcp:monitoring/slo:Slo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, basic_sli=None, calendar_period=None, display_name=None, goal=None, name=None, project=None, rolling_period_days=None, service=None, slo_id=None):
        """
        Get an existing Slo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] basic_sli: -
               (Required)
               Basic Service-Level Indicator (SLI) on a well-known service type.
               Performance will be computed on the basis of pre-defined metrics.
               SLIs are used to measure and calculate the quality of the Service's
               performance with respect to a single aspect of service quality.  Structure is documented below.
        :param pulumi.Input[str] calendar_period: -
               (Optional)
               A calendar period, semantically "since the start of the current
               <calendarPeriod>".
        :param pulumi.Input[str] display_name: -
               (Optional)
               Name used for UI elements listing this SLO.
        :param pulumi.Input[float] goal: -
               (Required)
               The fraction of service that must be good in order for this objective
               to be met. 0 < goal <= 0.999
        :param pulumi.Input[str] name: The full resource name for this service. The syntax is:
               projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[float] rolling_period_days: -
               (Optional)
               A rolling time period, semantically "in the past X days".
               Must be between 1 to 30 days, inclusive.
        :param pulumi.Input[str] service: -
               (Required)
               ID of the service to which this SLO belongs.
        :param pulumi.Input[str] slo_id: -
               (Optional)
               The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.

        The **basic_sli** object supports the following:

          * `latency` (`pulumi.Input[dict]`) - -
            (Required)
            Parameters for a latency threshold SLI.  Structure is documented below.
            * `threshold` (`pulumi.Input[str]`) - -
              (Required)
              A duration string, e.g. 10s.
              Good service is defined to be the count of requests made to
              this service that return in no more than threshold.

          * `locations` (`pulumi.Input[list]`) - -
            (Optional)
            An optional set of locations to which this SLI is relevant.
            Telemetry from other locations will not be used to calculate
            performance for this SLI. If omitted, this SLI applies to all
            locations in which the Service has activity. For service types
            that don't support breaking down by location, setting this
            field will result in an error.
          * `methods` (`pulumi.Input[list]`) - -
            (Optional)
            An optional set of RPCs to which this SLI is relevant.
            Telemetry from other methods will not be used to calculate
            performance for this SLI. If omitted, this SLI applies to all
            the Service's methods. For service types that don't support
            breaking down by method, setting this field will result in an
            error.
          * `versions` (`pulumi.Input[list]`) - -
            (Optional)
            The set of API versions to which this SLI is relevant.
            Telemetry from other API versions will not be used to
            calculate performance for this SLI. If omitted,
            this SLI applies to all API versions. For service types
            that don't support breaking down by version, setting this
            field will result in an error.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["basic_sli"] = basic_sli
        __props__["calendar_period"] = calendar_period
        __props__["display_name"] = display_name
        __props__["goal"] = goal
        __props__["name"] = name
        __props__["project"] = project
        __props__["rolling_period_days"] = rolling_period_days
        __props__["service"] = service
        __props__["slo_id"] = slo_id
        return Slo(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

