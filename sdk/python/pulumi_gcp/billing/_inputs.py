# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'AccountIamBindingConditionArgs',
    'AccountIamMemberConditionArgs',
    'BudgetAllUpdatesRuleArgs',
    'BudgetAmountArgs',
    'BudgetAmountSpecifiedAmountArgs',
    'BudgetBudgetFilterArgs',
    'BudgetThresholdRuleArgs',
]

@pulumi.input_type
class AccountIamBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    title: pulumi.Input[str] = pulumi.input_property("title")
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class AccountIamMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    title: pulumi.Input[str] = pulumi.input_property("title")
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class BudgetAllUpdatesRuleArgs:
    pubsub_topic: pulumi.Input[str] = pulumi.input_property("pubsubTopic")
    """
    The name of the Cloud Pub/Sub topic where budget related
    messages will be published, in the form
    projects/{project_id}/topics/{topic_id}. Updates are sent
    at regular intervals to the topic.
    """
    schema_version: Optional[pulumi.Input[str]] = pulumi.input_property("schemaVersion")
    """
    The schema version of the notification. Only "1.0" is
    accepted. It represents the JSON schema as defined in
    https://cloud.google.com/billing/docs/how-to/budgets#notification_format.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, pubsub_topic: pulumi.Input[str], schema_version: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] pubsub_topic: The name of the Cloud Pub/Sub topic where budget related
               messages will be published, in the form
               projects/{project_id}/topics/{topic_id}. Updates are sent
               at regular intervals to the topic.
        :param pulumi.Input[str] schema_version: The schema version of the notification. Only "1.0" is
               accepted. It represents the JSON schema as defined in
               https://cloud.google.com/billing/docs/how-to/budgets#notification_format.
        """
        __self__.pubsub_topic = pubsub_topic
        __self__.schema_version = schema_version

@pulumi.input_type
class BudgetAmountArgs:
    specified_amount: pulumi.Input['BudgetAmountSpecifiedAmountArgs'] = pulumi.input_property("specifiedAmount")
    """
    A specified amount to use as the budget. currencyCode is
    optional. If specified, it must match the currency of the
    billing account. The currencyCode is provided on output.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, specified_amount: pulumi.Input['BudgetAmountSpecifiedAmountArgs']) -> None:
        """
        :param pulumi.Input['BudgetAmountSpecifiedAmountArgs'] specified_amount: A specified amount to use as the budget. currencyCode is
               optional. If specified, it must match the currency of the
               billing account. The currencyCode is provided on output.  Structure is documented below.
        """
        __self__.specified_amount = specified_amount

@pulumi.input_type
class BudgetAmountSpecifiedAmountArgs:
    currency_code: Optional[pulumi.Input[str]] = pulumi.input_property("currencyCode")
    """
    The 3-letter currency code defined in ISO 4217.
    """
    nanos: Optional[pulumi.Input[float]] = pulumi.input_property("nanos")
    """
    Number of nano (10^-9) units of the amount.
    The value must be between -999,999,999 and +999,999,999
    inclusive. If units is positive, nanos must be positive or
    zero. If units is zero, nanos can be positive, zero, or
    negative. If units is negative, nanos must be negative or
    zero. For example $-1.75 is represented as units=-1 and
    nanos=-750,000,000.
    """
    units: Optional[pulumi.Input[str]] = pulumi.input_property("units")
    """
    The whole units of the amount. For example if currencyCode
    is "USD", then 1 unit is one US dollar.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, currency_code: Optional[pulumi.Input[str]] = None, nanos: Optional[pulumi.Input[float]] = None, units: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] currency_code: The 3-letter currency code defined in ISO 4217.
        :param pulumi.Input[float] nanos: Number of nano (10^-9) units of the amount.
               The value must be between -999,999,999 and +999,999,999
               inclusive. If units is positive, nanos must be positive or
               zero. If units is zero, nanos can be positive, zero, or
               negative. If units is negative, nanos must be negative or
               zero. For example $-1.75 is represented as units=-1 and
               nanos=-750,000,000.
        :param pulumi.Input[str] units: The whole units of the amount. For example if currencyCode
               is "USD", then 1 unit is one US dollar.
        """
        __self__.currency_code = currency_code
        __self__.nanos = nanos
        __self__.units = units

@pulumi.input_type
class BudgetBudgetFilterArgs:
    credit_types_treatment: Optional[pulumi.Input[str]] = pulumi.input_property("creditTypesTreatment")
    """
    Specifies how credits should be treated when determining spend
    for threshold calculations.
    """
    projects: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("projects")
    """
    A set of projects of the form projects/{project_id},
    specifying that usage from only this set of projects should be
    included in the budget. If omitted, the report will include
    all usage for the billing account, regardless of which project
    the usage occurred on. Only zero or one project can be
    specified currently.
    """
    services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("services")
    """
    A set of services of the form services/{service_id},
    specifying that usage from only this set of services should be
    included in the budget. If omitted, the report will include
    usage for all the services. The service names are available
    through the Catalog API:
    https://cloud.google.com/billing/v1/how-tos/catalog-api.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, credit_types_treatment: Optional[pulumi.Input[str]] = None, projects: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input[str] credit_types_treatment: Specifies how credits should be treated when determining spend
               for threshold calculations.
        :param pulumi.Input[List[pulumi.Input[str]]] projects: A set of projects of the form projects/{project_id},
               specifying that usage from only this set of projects should be
               included in the budget. If omitted, the report will include
               all usage for the billing account, regardless of which project
               the usage occurred on. Only zero or one project can be
               specified currently.
        :param pulumi.Input[List[pulumi.Input[str]]] services: A set of services of the form services/{service_id},
               specifying that usage from only this set of services should be
               included in the budget. If omitted, the report will include
               usage for all the services. The service names are available
               through the Catalog API:
               https://cloud.google.com/billing/v1/how-tos/catalog-api.
        """
        __self__.credit_types_treatment = credit_types_treatment
        __self__.projects = projects
        __self__.services = services

@pulumi.input_type
class BudgetThresholdRuleArgs:
    threshold_percent: pulumi.Input[float] = pulumi.input_property("thresholdPercent")
    """
    Send an alert when this threshold is exceeded. This is a
    1.0-based percentage, so 0.5 = 50%. Must be >= 0.
    """
    spend_basis: Optional[pulumi.Input[str]] = pulumi.input_property("spendBasis")
    """
    The type of basis used to determine if spend has passed
    the threshold.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, threshold_percent: pulumi.Input[float], spend_basis: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[float] threshold_percent: Send an alert when this threshold is exceeded. This is a
               1.0-based percentage, so 0.5 = 50%. Must be >= 0.
        :param pulumi.Input[str] spend_basis: The type of basis used to determine if spend has passed
               the threshold.
        """
        __self__.threshold_percent = threshold_percent
        __self__.spend_basis = spend_basis

