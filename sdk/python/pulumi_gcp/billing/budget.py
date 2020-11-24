# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Budget']


class Budget(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_updates_rule: Optional[pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']]] = None,
                 amount: Optional[pulumi.Input[pulumi.InputType['BudgetAmountArgs']]] = None,
                 billing_account: Optional[pulumi.Input[str]] = None,
                 budget_filter: Optional[pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 threshold_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Budget configuration for a billing account.

        To get more information about Budget, see:

        * [API documentation](https://cloud.google.com/billing/docs/reference/budget/rest/v1beta1/billingAccounts.budgets)
        * How-to Guides
            * [Creating a budget](https://cloud.google.com/billing/docs/how-to/budgets)

        ## Example Usage
        ### Billing Budget Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            amount={
                "specifiedAmount": {
                    "currencyCode": "USD",
                    "units": "100000",
                },
            },
            threshold_rules=[{
                "thresholdPercent": 0.5,
            }],
            opts=ResourceOptions(provider=google_beta))
        ```
        ### Billing Budget Lastperiod

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter={
                "projects": ["projects/my-project-name"],
            },
            amount={
                "lastPeriodAmount": True,
            },
            threshold_rules=[{
                "thresholdPercent": 10,
            }],
            opts=ResourceOptions(provider=google_beta))
        ```
        ### Billing Budget Filter

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter={
                "projects": ["projects/my-project-name"],
                "creditTypesTreatment": "EXCLUDE_ALL_CREDITS",
                "services": ["services/24E6-581D-38E5"],
            },
            amount={
                "specifiedAmount": {
                    "currencyCode": "USD",
                    "units": "100000",
                },
            },
            threshold_rules=[
                {
                    "thresholdPercent": 0.5,
                },
                {
                    "thresholdPercent": 0.9,
                    "spendBasis": "FORECASTED_SPEND",
                },
            ],
            opts=ResourceOptions(provider=google_beta))
        ```
        ### Billing Budget Notify

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        notification_channel = gcp.monitoring.NotificationChannel("notificationChannel",
            display_name="Example Notification Channel",
            type="email",
            labels={
                "email_address": "address@example.com",
            },
            opts=ResourceOptions(provider=google_beta))
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter={
                "projects": ["projects/my-project-name"],
            },
            amount={
                "specifiedAmount": {
                    "currencyCode": "USD",
                    "units": "100000",
                },
            },
            threshold_rules=[
                {
                    "thresholdPercent": 1,
                },
                {
                    "thresholdPercent": 1,
                    "spendBasis": "FORECASTED_SPEND",
                },
            ],
            all_updates_rule={
                "monitoringNotificationChannels": [notification_channel.id],
                "disableDefaultIamRecipients": True,
            },
            opts=ResourceOptions(provider=google_beta))
        ```

        ## Import

        Budget can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:billing/budget:Budget default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']] all_updates_rule: Defines notifications that are sent on every update to the
               billing account's spend, regardless of the thresholds defined
               using threshold rules.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['BudgetAmountArgs']] amount: The budgeted amount for each usage period.
               Structure is documented below.
        :param pulumi.Input[str] billing_account: ID of the billing account to set a budget on.
        :param pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']] budget_filter: Filters that define which resources are used to compute the actual
               spend against the budget.
               Structure is documented below.
        :param pulumi.Input[str] display_name: User data for display name in UI. Must be <= 60 chars.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]] threshold_rules: Rules that trigger alerts (notifications of thresholds being
               crossed) when spend exceeds the specified percentages of the
               budget.
               Structure is documented below.
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

            __props__['all_updates_rule'] = all_updates_rule
            if amount is None:
                raise TypeError("Missing required property 'amount'")
            __props__['amount'] = amount
            if billing_account is None:
                raise TypeError("Missing required property 'billing_account'")
            __props__['billing_account'] = billing_account
            __props__['budget_filter'] = budget_filter
            __props__['display_name'] = display_name
            if threshold_rules is None:
                raise TypeError("Missing required property 'threshold_rules'")
            __props__['threshold_rules'] = threshold_rules
            __props__['name'] = None
        super(Budget, __self__).__init__(
            'gcp:billing/budget:Budget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_updates_rule: Optional[pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']]] = None,
            amount: Optional[pulumi.Input[pulumi.InputType['BudgetAmountArgs']]] = None,
            billing_account: Optional[pulumi.Input[str]] = None,
            budget_filter: Optional[pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            threshold_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]]] = None) -> 'Budget':
        """
        Get an existing Budget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']] all_updates_rule: Defines notifications that are sent on every update to the
               billing account's spend, regardless of the thresholds defined
               using threshold rules.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['BudgetAmountArgs']] amount: The budgeted amount for each usage period.
               Structure is documented below.
        :param pulumi.Input[str] billing_account: ID of the billing account to set a budget on.
        :param pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']] budget_filter: Filters that define which resources are used to compute the actual
               spend against the budget.
               Structure is documented below.
        :param pulumi.Input[str] display_name: User data for display name in UI. Must be <= 60 chars.
        :param pulumi.Input[str] name: Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
               billingAccounts/{billingAccountId}/budgets/{budgetId}.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]] threshold_rules: Rules that trigger alerts (notifications of thresholds being
               crossed) when spend exceeds the specified percentages of the
               budget.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["all_updates_rule"] = all_updates_rule
        __props__["amount"] = amount
        __props__["billing_account"] = billing_account
        __props__["budget_filter"] = budget_filter
        __props__["display_name"] = display_name
        __props__["name"] = name
        __props__["threshold_rules"] = threshold_rules
        return Budget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allUpdatesRule")
    def all_updates_rule(self) -> pulumi.Output[Optional['outputs.BudgetAllUpdatesRule']]:
        """
        Defines notifications that are sent on every update to the
        billing account's spend, regardless of the thresholds defined
        using threshold rules.
        Structure is documented below.
        """
        return pulumi.get(self, "all_updates_rule")

    @property
    @pulumi.getter
    def amount(self) -> pulumi.Output['outputs.BudgetAmount']:
        """
        The budgeted amount for each usage period.
        Structure is documented below.
        """
        return pulumi.get(self, "amount")

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> pulumi.Output[str]:
        """
        ID of the billing account to set a budget on.
        """
        return pulumi.get(self, "billing_account")

    @property
    @pulumi.getter(name="budgetFilter")
    def budget_filter(self) -> pulumi.Output[Optional['outputs.BudgetBudgetFilter']]:
        """
        Filters that define which resources are used to compute the actual
        spend against the budget.
        Structure is documented below.
        """
        return pulumi.get(self, "budget_filter")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        User data for display name in UI. Must be <= 60 chars.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
        billingAccounts/{billingAccountId}/budgets/{budgetId}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="thresholdRules")
    def threshold_rules(self) -> pulumi.Output[Sequence['outputs.BudgetThresholdRule']]:
        """
        Rules that trigger alerts (notifications of thresholds being
        crossed) when spend exceeds the specified percentages of the
        budget.
        Structure is documented below.
        """
        return pulumi.get(self, "threshold_rules")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

