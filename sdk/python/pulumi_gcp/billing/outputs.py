# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'AccountIamBindingCondition',
    'AccountIamMemberCondition',
    'BudgetAllUpdatesRule',
    'BudgetAmount',
    'BudgetAmountSpecifiedAmount',
    'BudgetBudgetFilter',
    'BudgetThresholdRule',
]

@pulumi.output_type
class AccountIamBindingCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AccountIamMemberCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BudgetAllUpdatesRule(dict):
    pubsub_topic: str = pulumi.output_property("pubsubTopic")
    """
    The name of the Cloud Pub/Sub topic where budget related
    messages will be published, in the form
    projects/{project_id}/topics/{topic_id}. Updates are sent
    at regular intervals to the topic.
    """
    schema_version: Optional[str] = pulumi.output_property("schemaVersion")
    """
    The schema version of the notification. Only "1.0" is
    accepted. It represents the JSON schema as defined in
    https://cloud.google.com/billing/docs/how-to/budgets#notification_format.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BudgetAmount(dict):
    specified_amount: 'outputs.BudgetAmountSpecifiedAmount' = pulumi.output_property("specifiedAmount")
    """
    A specified amount to use as the budget. currencyCode is
    optional. If specified, it must match the currency of the
    billing account. The currencyCode is provided on output.  Structure is documented below.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BudgetAmountSpecifiedAmount(dict):
    currency_code: Optional[str] = pulumi.output_property("currencyCode")
    """
    The 3-letter currency code defined in ISO 4217.
    """
    nanos: Optional[float] = pulumi.output_property("nanos")
    """
    Number of nano (10^-9) units of the amount.
    The value must be between -999,999,999 and +999,999,999
    inclusive. If units is positive, nanos must be positive or
    zero. If units is zero, nanos can be positive, zero, or
    negative. If units is negative, nanos must be negative or
    zero. For example $-1.75 is represented as units=-1 and
    nanos=-750,000,000.
    """
    units: Optional[str] = pulumi.output_property("units")
    """
    The whole units of the amount. For example if currencyCode
    is "USD", then 1 unit is one US dollar.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BudgetBudgetFilter(dict):
    credit_types_treatment: Optional[str] = pulumi.output_property("creditTypesTreatment")
    """
    Specifies how credits should be treated when determining spend
    for threshold calculations.
    """
    projects: Optional[List[str]] = pulumi.output_property("projects")
    """
    A set of projects of the form projects/{project_id},
    specifying that usage from only this set of projects should be
    included in the budget. If omitted, the report will include
    all usage for the billing account, regardless of which project
    the usage occurred on. Only zero or one project can be
    specified currently.
    """
    services: Optional[List[str]] = pulumi.output_property("services")
    """
    A set of services of the form services/{service_id},
    specifying that usage from only this set of services should be
    included in the budget. If omitted, the report will include
    usage for all the services. The service names are available
    through the Catalog API:
    https://cloud.google.com/billing/v1/how-tos/catalog-api.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BudgetThresholdRule(dict):
    spend_basis: Optional[str] = pulumi.output_property("spendBasis")
    """
    The type of basis used to determine if spend has passed
    the threshold.
    """
    threshold_percent: float = pulumi.output_property("thresholdPercent")
    """
    Send an alert when this threshold is exceeded. This is a
    1.0-based percentage, so 0.5 = 50%. Must be >= 0.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


