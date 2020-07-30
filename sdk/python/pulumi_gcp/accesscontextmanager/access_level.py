# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class AccessLevel(pulumi.CustomResource):
    basic: pulumi.Output[dict]
    """
    A set of predefined conditions for the access level and a combining function.  Structure is documented below.

      * `combiningFunction` (`str`) - How the conditions list should be combined to determine if a request
        is granted this AccessLevel. If AND is used, each Condition in
        conditions must be satisfied for the AccessLevel to be applied. If
        OR is used, at least one Condition in conditions must be satisfied
        for the AccessLevel to be applied.
      * `conditions` (`list`) - A set of requirements for the AccessLevel to be granted.  Structure is documented below.
        * `devicePolicy` (`dict`) - Device specific restrictions, all restrictions must hold for
          the Condition to be true. If not specified, all devices are
          allowed.  Structure is documented below.
          * `allowedDeviceManagementLevels` (`list`) - A list of allowed device management levels.
            An empty list allows all management levels.
          * `allowedEncryptionStatuses` (`list`) - A list of allowed encryptions statuses.
            An empty list allows all statuses.
          * `osConstraints` (`list`) - A list of allowed OS versions.
            An empty list allows all types and all versions.  Structure is documented below.
            * `minimumVersion` (`str`) - The minimum allowed OS version. If not set, any version
              of this OS satisfies the constraint.
              Format: "major.minor.patch" such as "10.5.301", "9.2.1".
            * `osType` (`str`) - The operating system type of the device.

          * `requireAdminApproval` (`bool`) - Whether the device needs to be approved by the customer admin.
          * `requireCorpOwned` (`bool`) - Whether the device needs to be corp owned.
          * `requireScreenLock` (`bool`) - Whether or not screenlock is required for the DevicePolicy
            to be true. Defaults to false.

        * `ipSubnetworks` (`list`) - A list of CIDR block IP subnetwork specification. May be IPv4
          or IPv6.
          Note that for a CIDR IP address block, the specified IP address
          portion must be properly truncated (i.e. all the host bits must
          be zero) or the input is considered malformed. For example,
          "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
          for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
          is not. The originating IP of a request must be in one of the
          listed subnets in order for this Condition to be true.
          If empty, all IP addresses are allowed.
        * `members` (`list`) - An allowed list of members (users, service accounts).
          Using groups is not supported yet.
          The signed-in user originating the request must be a part of one
          of the provided members. If not specified, a request may come
          from any user (logged in/not logged in, not present in any
          groups, etc.).
          Formats: `user:{emailid}`, `serviceAccount:{emailid}`
        * `negate` (`bool`) - Whether to negate the Condition. If true, the Condition becomes
          a NAND over its non-empty fields, each field must be false for
          the Condition overall to be satisfied. Defaults to false.
        * `regions` (`list`) - The request must originate from one of the provided
          countries/regions.
          Format: A valid ISO 3166-1 alpha-2 code.
        * `requiredAccessLevels` (`list`) - A list of other access levels defined in the same Policy,
          referenced by resource name. Referencing an AccessLevel which
          does not exist is an error. All access levels listed must be
          granted for the Condition to be true.
          Format: accessPolicies/{policy_id}/accessLevels/{short_name}
    """
    custom: pulumi.Output[dict]
    """
    Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
    See CEL spec at: https://github.com/google/cel-spec.  Structure is documented below.

      * `expr` (`dict`) - Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
        This page details the objects and attributes that are used to the build the CEL expressions for
        custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.  Structure is documented below.
        * `description` (`str`) - Description of the expression
        * `expression` (`str`) - Textual representation of an expression in Common Expression Language syntax.
        * `location` (`str`) - String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
        * `title` (`str`) - Title for the expression, i.e. a short string describing its purpose.
    """
    description: pulumi.Output[str]
    """
    Description of the expression
    """
    name: pulumi.Output[str]
    """
    Resource name for the Access Level. The short_name component must begin
    with a letter and only include alphanumeric and '_'.
    Format: accessPolicies/{policy_id}/accessLevels/{short_name}
    """
    parent: pulumi.Output[str]
    """
    The AccessPolicy this AccessLevel lives in.
    Format: accessPolicies/{policy_id}
    """
    title: pulumi.Output[str]
    """
    Title for the expression, i.e. a short string describing its purpose.
    """
    def __init__(__self__, resource_name, opts=None, basic=None, custom=None, description=None, name=None, parent=None, title=None, __props__=None, __name__=None, __opts__=None):
        """
        An AccessLevel is a label that can be applied to requests to GCP services,
        along with a list of requirements necessary for the label to be applied.

        To get more information about AccessLevel, see:

        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
        * How-to Guides
            * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)

        ## Example Usage
        ### Access Context Manager Access Level Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        access_policy = gcp.accesscontextmanager.AccessPolicy("access-policy",
            parent="organizations/123456789",
            title="my policy")
        access_level = gcp.accesscontextmanager.AccessLevel("access-level",
            basic=gcp.accesscontextmanager.AccessLevelBasicArgs(
                conditions=[gcp.accesscontextmanager.AccessLevelBasicConditionArgs(
                    device_policy=gcp.accesscontextmanager.AccessLevelBasicConditionDevicePolicyArgs(
                        os_constraints=[gcp.accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArgs(
                            os_type="DESKTOP_CHROME_OS",
                        )],
                        require_screen_lock=True,
                    ),
                    regions=[
                        "CH",
                        "IT",
                        "US",
                    ],
                )],
            ),
            parent=access_policy.name.apply(lambda name: f"accessPolicies/{name}"),
            title="chromeos_no_lock")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] basic: A set of predefined conditions for the access level and a combining function.  Structure is documented below.
        :param pulumi.Input[dict] custom: Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
               See CEL spec at: https://github.com/google/cel-spec.  Structure is documented below.
        :param pulumi.Input[str] description: Description of the expression
        :param pulumi.Input[str] name: Resource name for the Access Level. The short_name component must begin
               with a letter and only include alphanumeric and '_'.
               Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        :param pulumi.Input[str] parent: The AccessPolicy this AccessLevel lives in.
               Format: accessPolicies/{policy_id}
        :param pulumi.Input[str] title: Title for the expression, i.e. a short string describing its purpose.

        The **basic** object supports the following:

          * `combiningFunction` (`pulumi.Input[str]`) - How the conditions list should be combined to determine if a request
            is granted this AccessLevel. If AND is used, each Condition in
            conditions must be satisfied for the AccessLevel to be applied. If
            OR is used, at least one Condition in conditions must be satisfied
            for the AccessLevel to be applied.
          * `conditions` (`pulumi.Input[list]`) - A set of requirements for the AccessLevel to be granted.  Structure is documented below.
            * `devicePolicy` (`pulumi.Input[dict]`) - Device specific restrictions, all restrictions must hold for
              the Condition to be true. If not specified, all devices are
              allowed.  Structure is documented below.
              * `allowedDeviceManagementLevels` (`pulumi.Input[list]`) - A list of allowed device management levels.
                An empty list allows all management levels.
              * `allowedEncryptionStatuses` (`pulumi.Input[list]`) - A list of allowed encryptions statuses.
                An empty list allows all statuses.
              * `osConstraints` (`pulumi.Input[list]`) - A list of allowed OS versions.
                An empty list allows all types and all versions.  Structure is documented below.
                * `minimumVersion` (`pulumi.Input[str]`) - The minimum allowed OS version. If not set, any version
                  of this OS satisfies the constraint.
                  Format: "major.minor.patch" such as "10.5.301", "9.2.1".
                * `osType` (`pulumi.Input[str]`) - The operating system type of the device.

              * `requireAdminApproval` (`pulumi.Input[bool]`) - Whether the device needs to be approved by the customer admin.
              * `requireCorpOwned` (`pulumi.Input[bool]`) - Whether the device needs to be corp owned.
              * `requireScreenLock` (`pulumi.Input[bool]`) - Whether or not screenlock is required for the DevicePolicy
                to be true. Defaults to false.

            * `ipSubnetworks` (`pulumi.Input[list]`) - A list of CIDR block IP subnetwork specification. May be IPv4
              or IPv6.
              Note that for a CIDR IP address block, the specified IP address
              portion must be properly truncated (i.e. all the host bits must
              be zero) or the input is considered malformed. For example,
              "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
              for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
              is not. The originating IP of a request must be in one of the
              listed subnets in order for this Condition to be true.
              If empty, all IP addresses are allowed.
            * `members` (`pulumi.Input[list]`) - An allowed list of members (users, service accounts).
              Using groups is not supported yet.
              The signed-in user originating the request must be a part of one
              of the provided members. If not specified, a request may come
              from any user (logged in/not logged in, not present in any
              groups, etc.).
              Formats: `user:{emailid}`, `serviceAccount:{emailid}`
            * `negate` (`pulumi.Input[bool]`) - Whether to negate the Condition. If true, the Condition becomes
              a NAND over its non-empty fields, each field must be false for
              the Condition overall to be satisfied. Defaults to false.
            * `regions` (`pulumi.Input[list]`) - The request must originate from one of the provided
              countries/regions.
              Format: A valid ISO 3166-1 alpha-2 code.
            * `requiredAccessLevels` (`pulumi.Input[list]`) - A list of other access levels defined in the same Policy,
              referenced by resource name. Referencing an AccessLevel which
              does not exist is an error. All access levels listed must be
              granted for the Condition to be true.
              Format: accessPolicies/{policy_id}/accessLevels/{short_name}

        The **custom** object supports the following:

          * `expr` (`pulumi.Input[dict]`) - Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
            This page details the objects and attributes that are used to the build the CEL expressions for
            custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.  Structure is documented below.
            * `description` (`pulumi.Input[str]`) - Description of the expression
            * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
            * `location` (`pulumi.Input[str]`) - String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
            * `title` (`pulumi.Input[str]`) - Title for the expression, i.e. a short string describing its purpose.
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

            __props__['basic'] = basic
            __props__['custom'] = custom
            __props__['description'] = description
            __props__['name'] = name
            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
        super(AccessLevel, __self__).__init__(
            'gcp:accesscontextmanager/accessLevel:AccessLevel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, basic=None, custom=None, description=None, name=None, parent=None, title=None):
        """
        Get an existing AccessLevel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] basic: A set of predefined conditions for the access level and a combining function.  Structure is documented below.
        :param pulumi.Input[dict] custom: Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
               See CEL spec at: https://github.com/google/cel-spec.  Structure is documented below.
        :param pulumi.Input[str] description: Description of the expression
        :param pulumi.Input[str] name: Resource name for the Access Level. The short_name component must begin
               with a letter and only include alphanumeric and '_'.
               Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        :param pulumi.Input[str] parent: The AccessPolicy this AccessLevel lives in.
               Format: accessPolicies/{policy_id}
        :param pulumi.Input[str] title: Title for the expression, i.e. a short string describing its purpose.

        The **basic** object supports the following:

          * `combiningFunction` (`pulumi.Input[str]`) - How the conditions list should be combined to determine if a request
            is granted this AccessLevel. If AND is used, each Condition in
            conditions must be satisfied for the AccessLevel to be applied. If
            OR is used, at least one Condition in conditions must be satisfied
            for the AccessLevel to be applied.
          * `conditions` (`pulumi.Input[list]`) - A set of requirements for the AccessLevel to be granted.  Structure is documented below.
            * `devicePolicy` (`pulumi.Input[dict]`) - Device specific restrictions, all restrictions must hold for
              the Condition to be true. If not specified, all devices are
              allowed.  Structure is documented below.
              * `allowedDeviceManagementLevels` (`pulumi.Input[list]`) - A list of allowed device management levels.
                An empty list allows all management levels.
              * `allowedEncryptionStatuses` (`pulumi.Input[list]`) - A list of allowed encryptions statuses.
                An empty list allows all statuses.
              * `osConstraints` (`pulumi.Input[list]`) - A list of allowed OS versions.
                An empty list allows all types and all versions.  Structure is documented below.
                * `minimumVersion` (`pulumi.Input[str]`) - The minimum allowed OS version. If not set, any version
                  of this OS satisfies the constraint.
                  Format: "major.minor.patch" such as "10.5.301", "9.2.1".
                * `osType` (`pulumi.Input[str]`) - The operating system type of the device.

              * `requireAdminApproval` (`pulumi.Input[bool]`) - Whether the device needs to be approved by the customer admin.
              * `requireCorpOwned` (`pulumi.Input[bool]`) - Whether the device needs to be corp owned.
              * `requireScreenLock` (`pulumi.Input[bool]`) - Whether or not screenlock is required for the DevicePolicy
                to be true. Defaults to false.

            * `ipSubnetworks` (`pulumi.Input[list]`) - A list of CIDR block IP subnetwork specification. May be IPv4
              or IPv6.
              Note that for a CIDR IP address block, the specified IP address
              portion must be properly truncated (i.e. all the host bits must
              be zero) or the input is considered malformed. For example,
              "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
              for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
              is not. The originating IP of a request must be in one of the
              listed subnets in order for this Condition to be true.
              If empty, all IP addresses are allowed.
            * `members` (`pulumi.Input[list]`) - An allowed list of members (users, service accounts).
              Using groups is not supported yet.
              The signed-in user originating the request must be a part of one
              of the provided members. If not specified, a request may come
              from any user (logged in/not logged in, not present in any
              groups, etc.).
              Formats: `user:{emailid}`, `serviceAccount:{emailid}`
            * `negate` (`pulumi.Input[bool]`) - Whether to negate the Condition. If true, the Condition becomes
              a NAND over its non-empty fields, each field must be false for
              the Condition overall to be satisfied. Defaults to false.
            * `regions` (`pulumi.Input[list]`) - The request must originate from one of the provided
              countries/regions.
              Format: A valid ISO 3166-1 alpha-2 code.
            * `requiredAccessLevels` (`pulumi.Input[list]`) - A list of other access levels defined in the same Policy,
              referenced by resource name. Referencing an AccessLevel which
              does not exist is an error. All access levels listed must be
              granted for the Condition to be true.
              Format: accessPolicies/{policy_id}/accessLevels/{short_name}

        The **custom** object supports the following:

          * `expr` (`pulumi.Input[dict]`) - Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
            This page details the objects and attributes that are used to the build the CEL expressions for
            custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.  Structure is documented below.
            * `description` (`pulumi.Input[str]`) - Description of the expression
            * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
            * `location` (`pulumi.Input[str]`) - String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
            * `title` (`pulumi.Input[str]`) - Title for the expression, i.e. a short string describing its purpose.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["basic"] = basic
        __props__["custom"] = custom
        __props__["description"] = description
        __props__["name"] = name
        __props__["parent"] = parent
        __props__["title"] = title
        return AccessLevel(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
