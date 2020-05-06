# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FirewallRule(pulumi.CustomResource):
    action: pulumi.Output[str]
    """
    The action to take if this rule matches.
    """
    description: pulumi.Output[str]
    """
    An optional string description of this rule.
    """
    priority: pulumi.Output[float]
    """
    A positive integer that defines the order of rule evaluation.
    Rules with the lowest priority are evaluated first.
    A default rule at priority Int32.MaxValue matches all IPv4 and
    IPv6 traffic when no previous rule matches. Only the action of
    this rule can be modified by the user.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    source_range: pulumi.Output[str]
    """
    IP address or range, defined using CIDR notation, of requests that this rule applies to.
    """
    def __init__(__self__, resource_name, opts=None, action=None, description=None, priority=None, project=None, source_range=None, __props__=None, __name__=None, __opts__=None):
        """
        A single firewall rule that is evaluated against incoming traffic
        and provides an action to take on matched requests.


        To get more information about FirewallRule, see:

        * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/appengine/docs/standard/python/creating-firewalls#creating_firewall_rules)

        ## Example Usage - App Engine Firewall Rule Basic


        {{ % example python % }}
        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            project_id="ae-project",
            org_id="123456789")
        app = gcp.appengine.Application("app",
            project=my_project.project_id,
            location_id="us-central")
        rule = gcp.appengine.FirewallRule("rule",
            project=app.project,
            priority=1000,
            action="ALLOW",
            source_range="*")
        ```
        {{ % /example % }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action to take if this rule matches.
        :param pulumi.Input[str] description: An optional string description of this rule.
        :param pulumi.Input[float] priority: A positive integer that defines the order of rule evaluation.
               Rules with the lowest priority are evaluated first.
               A default rule at priority Int32.MaxValue matches all IPv4 and
               IPv6 traffic when no previous rule matches. Only the action of
               this rule can be modified by the user.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] source_range: IP address or range, defined using CIDR notation, of requests that this rule applies to.
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

            if action is None:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            __props__['description'] = description
            __props__['priority'] = priority
            __props__['project'] = project
            if source_range is None:
                raise TypeError("Missing required property 'source_range'")
            __props__['source_range'] = source_range
        super(FirewallRule, __self__).__init__(
            'gcp:appengine/firewallRule:FirewallRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, action=None, description=None, priority=None, project=None, source_range=None):
        """
        Get an existing FirewallRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action to take if this rule matches.
        :param pulumi.Input[str] description: An optional string description of this rule.
        :param pulumi.Input[float] priority: A positive integer that defines the order of rule evaluation.
               Rules with the lowest priority are evaluated first.
               A default rule at priority Int32.MaxValue matches all IPv4 and
               IPv6 traffic when no previous rule matches. Only the action of
               this rule can be modified by the user.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] source_range: IP address or range, defined using CIDR notation, of requests that this rule applies to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["description"] = description
        __props__["priority"] = priority
        __props__["project"] = project
        __props__["source_range"] = source_range
        return FirewallRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

