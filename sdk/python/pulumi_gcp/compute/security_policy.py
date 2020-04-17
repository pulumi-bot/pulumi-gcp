# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecurityPolicy(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    An optional description of this security policy. Max size is 2048.
    """
    fingerprint: pulumi.Output[str]
    """
    Fingerprint of this resource.
    """
    name: pulumi.Output[str]
    """
    The name of the security policy.
    """
    project: pulumi.Output[str]
    """
    The project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    rules: pulumi.Output[list]
    """
    The set of rules that belong to this policy. There must always be a default
    rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
    security policy, a default rule with action "allow" will be added. Structure is documented below.

      * `action` (`str`) - Action to take when `match` matches the request. Valid values:
        * "allow" : allow access to target
        * "deny(status)" : deny access to target, returns the  HTTP response code specified (valid values are 403, 404 and 502)
      * `description` (`str`) - An optional description of this security policy. Max size is 2048.
      * `match` (`dict`) - A match condition that incoming traffic is evaluated against.
        If it evaluates to true, the corresponding `action` is enforced. Structure is documented below.
        * `config` (`dict`) - The configuration options available when specifying `versioned_expr`.
          This field must be specified if `versioned_expr` is specified and cannot be specified if `versioned_expr` is not specified.
          Structure is documented below.
          * `srcIpRanges` (`list`) - Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation
            to match against inbound traffic. There is a limit of 5 IP ranges per rule. A value of '\*' matches all IPs
            (can be used to override the default behavior).

        * `expr` (`dict`) - User defined CEVAL expression. A CEVAL expression is used to specify match criteria
          such as origin.ip, source.region_code and contents in the request header.
          Structure is documented below.
          * `expression` (`str`) - Textual representation of an expression in Common Expression Language syntax.
            The application context of the containing message determines which well-known feature set of CEL is supported.

        * `versionedExpr` (`str`) - Predefined rule expression. If this field is specified, `config` must also be specified.
          Available options:
          * SRC_IPS_V1: Must specify the corresponding `src_ip_ranges` field in `config`.

      * `preview` (`bool`) - When set to true, the `action` specified above is not enforced.
        Stackdriver logs for requests that trigger a preview action are annotated as such.
      * `priority` (`float`) - An unique positive integer indicating the priority of evaluation for a rule.
        Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, project=None, rules=None, __props__=None, __name__=None, __opts__=None):
        """
        A Security Policy defines an IP blacklist or whitelist that protects load balanced Google Cloud services by denying or permitting traffic from specified IP ranges. For more information
        see the [official documentation](https://cloud.google.com/armor/docs/configure-security-policies)
        and the [API](https://cloud.google.com/compute/docs/reference/rest/beta/securityPolicies).



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this security policy. Max size is 2048.
        :param pulumi.Input[str] name: The name of the security policy.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[list] rules: The set of rules that belong to this policy. There must always be a default
               rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
               security policy, a default rule with action "allow" will be added. Structure is documented below.

        The **rules** object supports the following:

          * `action` (`pulumi.Input[str]`) - Action to take when `match` matches the request. Valid values:
            * "allow" : allow access to target
            * "deny(status)" : deny access to target, returns the  HTTP response code specified (valid values are 403, 404 and 502)
          * `description` (`pulumi.Input[str]`) - An optional description of this security policy. Max size is 2048.
          * `match` (`pulumi.Input[dict]`) - A match condition that incoming traffic is evaluated against.
            If it evaluates to true, the corresponding `action` is enforced. Structure is documented below.
            * `config` (`pulumi.Input[dict]`) - The configuration options available when specifying `versioned_expr`.
              This field must be specified if `versioned_expr` is specified and cannot be specified if `versioned_expr` is not specified.
              Structure is documented below.
              * `srcIpRanges` (`pulumi.Input[list]`) - Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation
                to match against inbound traffic. There is a limit of 5 IP ranges per rule. A value of '\*' matches all IPs
                (can be used to override the default behavior).

            * `expr` (`pulumi.Input[dict]`) - User defined CEVAL expression. A CEVAL expression is used to specify match criteria
              such as origin.ip, source.region_code and contents in the request header.
              Structure is documented below.
              * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
                The application context of the containing message determines which well-known feature set of CEL is supported.

            * `versionedExpr` (`pulumi.Input[str]`) - Predefined rule expression. If this field is specified, `config` must also be specified.
              Available options:
              * SRC_IPS_V1: Must specify the corresponding `src_ip_ranges` field in `config`.

          * `preview` (`pulumi.Input[bool]`) - When set to true, the `action` specified above is not enforced.
            Stackdriver logs for requests that trigger a preview action are annotated as such.
          * `priority` (`pulumi.Input[float]`) - An unique positive integer indicating the priority of evaluation for a rule.
            Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
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

            __props__['description'] = description
            __props__['name'] = name
            __props__['project'] = project
            __props__['rules'] = rules
            __props__['fingerprint'] = None
            __props__['self_link'] = None
        super(SecurityPolicy, __self__).__init__(
            'gcp:compute/securityPolicy:SecurityPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, fingerprint=None, name=None, project=None, rules=None, self_link=None):
        """
        Get an existing SecurityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this security policy. Max size is 2048.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource.
        :param pulumi.Input[str] name: The name of the security policy.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[list] rules: The set of rules that belong to this policy. There must always be a default
               rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
               security policy, a default rule with action "allow" will be added. Structure is documented below.
        :param pulumi.Input[str] self_link: The URI of the created resource.

        The **rules** object supports the following:

          * `action` (`pulumi.Input[str]`) - Action to take when `match` matches the request. Valid values:
            * "allow" : allow access to target
            * "deny(status)" : deny access to target, returns the  HTTP response code specified (valid values are 403, 404 and 502)
          * `description` (`pulumi.Input[str]`) - An optional description of this security policy. Max size is 2048.
          * `match` (`pulumi.Input[dict]`) - A match condition that incoming traffic is evaluated against.
            If it evaluates to true, the corresponding `action` is enforced. Structure is documented below.
            * `config` (`pulumi.Input[dict]`) - The configuration options available when specifying `versioned_expr`.
              This field must be specified if `versioned_expr` is specified and cannot be specified if `versioned_expr` is not specified.
              Structure is documented below.
              * `srcIpRanges` (`pulumi.Input[list]`) - Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation
                to match against inbound traffic. There is a limit of 5 IP ranges per rule. A value of '\*' matches all IPs
                (can be used to override the default behavior).

            * `expr` (`pulumi.Input[dict]`) - User defined CEVAL expression. A CEVAL expression is used to specify match criteria
              such as origin.ip, source.region_code and contents in the request header.
              Structure is documented below.
              * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
                The application context of the containing message determines which well-known feature set of CEL is supported.

            * `versionedExpr` (`pulumi.Input[str]`) - Predefined rule expression. If this field is specified, `config` must also be specified.
              Available options:
              * SRC_IPS_V1: Must specify the corresponding `src_ip_ranges` field in `config`.

          * `preview` (`pulumi.Input[bool]`) - When set to true, the `action` specified above is not enforced.
            Stackdriver logs for requests that trigger a preview action are annotated as such.
          * `priority` (`pulumi.Input[float]`) - An unique positive integer indicating the priority of evaluation for a rule.
            Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["fingerprint"] = fingerprint
        __props__["name"] = name
        __props__["project"] = project
        __props__["rules"] = rules
        __props__["self_link"] = self_link
        return SecurityPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

