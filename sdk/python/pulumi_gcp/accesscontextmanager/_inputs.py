# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs

@pulumi.input_type
class AccessLevelBasicArgs:
    conditions: pulumi.Input[List[pulumi.Input['AccessLevelBasicConditionArgs']]] = pulumi.input_property("conditions")
    """
    A set of requirements for the AccessLevel to be granted.  Structure is documented below.
    """
    combining_function: Optional[pulumi.Input[str]] = pulumi.input_property("combiningFunction")
    """
    How the conditions list should be combined to determine if a request
    is granted this AccessLevel. If AND is used, each Condition in
    conditions must be satisfied for the AccessLevel to be applied. If
    OR is used, at least one Condition in conditions must be satisfied
    for the AccessLevel to be applied.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, conditions: pulumi.Input[List[pulumi.Input['AccessLevelBasicConditionArgs']]], combining_function: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input['AccessLevelBasicConditionArgs']]] conditions: A set of requirements for the AccessLevel to be granted.  Structure is documented below.
        :param pulumi.Input[str] combining_function: How the conditions list should be combined to determine if a request
               is granted this AccessLevel. If AND is used, each Condition in
               conditions must be satisfied for the AccessLevel to be applied. If
               OR is used, at least one Condition in conditions must be satisfied
               for the AccessLevel to be applied.
        """
        __self__.conditions = conditions
        __self__.combining_function = combining_function

@pulumi.input_type
class AccessLevelBasicConditionArgs:
    device_policy: Optional[pulumi.Input['AccessLevelBasicConditionDevicePolicyArgs']] = pulumi.input_property("devicePolicy")
    """
    Device specific restrictions, all restrictions must hold for
    the Condition to be true. If not specified, all devices are
    allowed.  Structure is documented below.
    """
    ip_subnetworks: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("ipSubnetworks")
    """
    A list of CIDR block IP subnetwork specification. May be IPv4
    or IPv6.
    Note that for a CIDR IP address block, the specified IP address
    portion must be properly truncated (i.e. all the host bits must
    be zero) or the input is considered malformed. For example,
    "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
    for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
    is not. The originating IP of a request must be in one of the
    listed subnets in order for this Condition to be true.
    If empty, all IP addresses are allowed.
    """
    members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("members")
    """
    An allowed list of members (users, service accounts).
    Using groups is not supported yet.
    The signed-in user originating the request must be a part of one
    of the provided members. If not specified, a request may come
    from any user (logged in/not logged in, not present in any
    groups, etc.).
    Formats: `user:{emailid}`, `serviceAccount:{emailid}`
    """
    negate: Optional[pulumi.Input[bool]] = pulumi.input_property("negate")
    """
    Whether to negate the Condition. If true, the Condition becomes
    a NAND over its non-empty fields, each field must be false for
    the Condition overall to be satisfied. Defaults to false.
    """
    regions: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("regions")
    """
    The request must originate from one of the provided
    countries/regions.
    Format: A valid ISO 3166-1 alpha-2 code.
    """
    required_access_levels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("requiredAccessLevels")
    """
    A list of other access levels defined in the same Policy,
    referenced by resource name. Referencing an AccessLevel which
    does not exist is an error. All access levels listed must be
    granted for the Condition to be true.
    Format: accessPolicies/{policy_id}/accessLevels/{short_name}
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, device_policy: Optional[pulumi.Input['AccessLevelBasicConditionDevicePolicyArgs']] = None, ip_subnetworks: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, negate: Optional[pulumi.Input[bool]] = None, regions: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, required_access_levels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input['AccessLevelBasicConditionDevicePolicyArgs'] device_policy: Device specific restrictions, all restrictions must hold for
               the Condition to be true. If not specified, all devices are
               allowed.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] ip_subnetworks: A list of CIDR block IP subnetwork specification. May be IPv4
               or IPv6.
               Note that for a CIDR IP address block, the specified IP address
               portion must be properly truncated (i.e. all the host bits must
               be zero) or the input is considered malformed. For example,
               "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
               for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
               is not. The originating IP of a request must be in one of the
               listed subnets in order for this Condition to be true.
               If empty, all IP addresses are allowed.
        :param pulumi.Input[List[pulumi.Input[str]]] members: An allowed list of members (users, service accounts).
               Using groups is not supported yet.
               The signed-in user originating the request must be a part of one
               of the provided members. If not specified, a request may come
               from any user (logged in/not logged in, not present in any
               groups, etc.).
               Formats: `user:{emailid}`, `serviceAccount:{emailid}`
        :param pulumi.Input[bool] negate: Whether to negate the Condition. If true, the Condition becomes
               a NAND over its non-empty fields, each field must be false for
               the Condition overall to be satisfied. Defaults to false.
        :param pulumi.Input[List[pulumi.Input[str]]] regions: The request must originate from one of the provided
               countries/regions.
               Format: A valid ISO 3166-1 alpha-2 code.
        :param pulumi.Input[List[pulumi.Input[str]]] required_access_levels: A list of other access levels defined in the same Policy,
               referenced by resource name. Referencing an AccessLevel which
               does not exist is an error. All access levels listed must be
               granted for the Condition to be true.
               Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        """
        __self__.device_policy = device_policy
        __self__.ip_subnetworks = ip_subnetworks
        __self__.members = members
        __self__.negate = negate
        __self__.regions = regions
        __self__.required_access_levels = required_access_levels

@pulumi.input_type
class AccessLevelBasicConditionDevicePolicyArgs:
    allowed_device_management_levels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("allowedDeviceManagementLevels")
    """
    A list of allowed device management levels.
    An empty list allows all management levels.
    """
    allowed_encryption_statuses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("allowedEncryptionStatuses")
    """
    A list of allowed encryptions statuses.
    An empty list allows all statuses.
    """
    os_constraints: Optional[pulumi.Input[List[pulumi.Input['AccessLevelBasicConditionDevicePolicyOsConstraintArgs']]]] = pulumi.input_property("osConstraints")
    """
    A list of allowed OS versions.
    An empty list allows all types and all versions.  Structure is documented below.
    """
    require_admin_approval: Optional[pulumi.Input[bool]] = pulumi.input_property("requireAdminApproval")
    """
    Whether the device needs to be approved by the customer admin.
    """
    require_corp_owned: Optional[pulumi.Input[bool]] = pulumi.input_property("requireCorpOwned")
    """
    Whether the device needs to be corp owned.
    """
    require_screen_lock: Optional[pulumi.Input[bool]] = pulumi.input_property("requireScreenLock")
    """
    Whether or not screenlock is required for the DevicePolicy
    to be true. Defaults to false.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, allowed_device_management_levels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, allowed_encryption_statuses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, os_constraints: Optional[pulumi.Input[List[pulumi.Input['AccessLevelBasicConditionDevicePolicyOsConstraintArgs']]]] = None, require_admin_approval: Optional[pulumi.Input[bool]] = None, require_corp_owned: Optional[pulumi.Input[bool]] = None, require_screen_lock: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_device_management_levels: A list of allowed device management levels.
               An empty list allows all management levels.
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_encryption_statuses: A list of allowed encryptions statuses.
               An empty list allows all statuses.
        :param pulumi.Input[List[pulumi.Input['AccessLevelBasicConditionDevicePolicyOsConstraintArgs']]] os_constraints: A list of allowed OS versions.
               An empty list allows all types and all versions.  Structure is documented below.
        :param pulumi.Input[bool] require_admin_approval: Whether the device needs to be approved by the customer admin.
        :param pulumi.Input[bool] require_corp_owned: Whether the device needs to be corp owned.
        :param pulumi.Input[bool] require_screen_lock: Whether or not screenlock is required for the DevicePolicy
               to be true. Defaults to false.
        """
        __self__.allowed_device_management_levels = allowed_device_management_levels
        __self__.allowed_encryption_statuses = allowed_encryption_statuses
        __self__.os_constraints = os_constraints
        __self__.require_admin_approval = require_admin_approval
        __self__.require_corp_owned = require_corp_owned
        __self__.require_screen_lock = require_screen_lock

@pulumi.input_type
class AccessLevelBasicConditionDevicePolicyOsConstraintArgs:
    os_type: pulumi.Input[str] = pulumi.input_property("osType")
    """
    The operating system type of the device.
    """
    minimum_version: Optional[pulumi.Input[str]] = pulumi.input_property("minimumVersion")
    """
    The minimum allowed OS version. If not set, any version
    of this OS satisfies the constraint.
    Format: "major.minor.patch" such as "10.5.301", "9.2.1".
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, os_type: pulumi.Input[str], minimum_version: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] os_type: The operating system type of the device.
        :param pulumi.Input[str] minimum_version: The minimum allowed OS version. If not set, any version
               of this OS satisfies the constraint.
               Format: "major.minor.patch" such as "10.5.301", "9.2.1".
        """
        __self__.os_type = os_type
        __self__.minimum_version = minimum_version

@pulumi.input_type
class AccessLevelCustomArgs:
    expr: pulumi.Input['AccessLevelCustomExprArgs'] = pulumi.input_property("expr")
    """
    Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
    This page details the objects and attributes that are used to the build the CEL expressions for
    custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expr: pulumi.Input['AccessLevelCustomExprArgs']) -> None:
        """
        :param pulumi.Input['AccessLevelCustomExprArgs'] expr: Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
               This page details the objects and attributes that are used to the build the CEL expressions for
               custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.  Structure is documented below.
        """
        __self__.expr = expr

@pulumi.input_type
class AccessLevelCustomExprArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    Description of the expression
    """
    location: Optional[pulumi.Input[str]] = pulumi.input_property("location")
    """
    String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
    """
    title: Optional[pulumi.Input[str]] = pulumi.input_property("title")
    """
    Title for the expression, i.e. a short string describing its purpose.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None, location: Optional[pulumi.Input[str]] = None, title: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] description: Description of the expression
        :param pulumi.Input[str] location: String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
        :param pulumi.Input[str] title: Title for the expression, i.e. a short string describing its purpose.
        """
        __self__.expression = expression
        __self__.description = description
        __self__.location = location
        __self__.title = title

@pulumi.input_type
class ServicePerimeterSpecArgs:
    access_levels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("accessLevels")
    """
    A list of AccessLevel resource names that allow resources within
    the ServicePerimeter to be accessed from the internet.
    AccessLevels listed must be in the same policy as this
    ServicePerimeter. Referencing a nonexistent AccessLevel is a
    syntax error. If no AccessLevel names are listed, resources within
    the perimeter can only be accessed via GCP calls with request
    origins within the perimeter. For Service Perimeter Bridge, must
    be empty.
    Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
    """
    resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("resources")
    """
    A list of GCP resources that are inside of the service perimeter.
    Currently only projects are allowed.
    Format: projects/{project_number}
    """
    restricted_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("restrictedServices")
    """
    GCP services that are subject to the Service Perimeter
    restrictions. Must contain a list of services. For example, if
    `storage.googleapis.com` is specified, access to the storage
    buckets inside the perimeter must meet the perimeter's access
    restrictions.
    """
    vpc_accessible_services: Optional[pulumi.Input['ServicePerimeterSpecVpcAccessibleServicesArgs']] = pulumi.input_property("vpcAccessibleServices")
    """
    Specifies how APIs are allowed to communicate within the Service
    Perimeter.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, access_levels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, restricted_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, vpc_accessible_services: Optional[pulumi.Input['ServicePerimeterSpecVpcAccessibleServicesArgs']] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] access_levels: A list of AccessLevel resource names that allow resources within
               the ServicePerimeter to be accessed from the internet.
               AccessLevels listed must be in the same policy as this
               ServicePerimeter. Referencing a nonexistent AccessLevel is a
               syntax error. If no AccessLevel names are listed, resources within
               the perimeter can only be accessed via GCP calls with request
               origins within the perimeter. For Service Perimeter Bridge, must
               be empty.
               Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
        :param pulumi.Input[List[pulumi.Input[str]]] resources: A list of GCP resources that are inside of the service perimeter.
               Currently only projects are allowed.
               Format: projects/{project_number}
        :param pulumi.Input[List[pulumi.Input[str]]] restricted_services: GCP services that are subject to the Service Perimeter
               restrictions. Must contain a list of services. For example, if
               `storage.googleapis.com` is specified, access to the storage
               buckets inside the perimeter must meet the perimeter's access
               restrictions.
        :param pulumi.Input['ServicePerimeterSpecVpcAccessibleServicesArgs'] vpc_accessible_services: Specifies how APIs are allowed to communicate within the Service
               Perimeter.  Structure is documented below.
        """
        __self__.access_levels = access_levels
        __self__.resources = resources
        __self__.restricted_services = restricted_services
        __self__.vpc_accessible_services = vpc_accessible_services

@pulumi.input_type
class ServicePerimeterSpecVpcAccessibleServicesArgs:
    allowed_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("allowedServices")
    """
    The list of APIs usable within the Service Perimeter.
    Must be empty unless `enableRestriction` is True.
    """
    enable_restriction: Optional[pulumi.Input[bool]] = pulumi.input_property("enableRestriction")
    """
    Whether to restrict API calls within the Service Perimeter to the
    list of APIs specified in 'allowedServices'.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, allowed_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, enable_restriction: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_services: The list of APIs usable within the Service Perimeter.
               Must be empty unless `enableRestriction` is True.
        :param pulumi.Input[bool] enable_restriction: Whether to restrict API calls within the Service Perimeter to the
               list of APIs specified in 'allowedServices'.
        """
        __self__.allowed_services = allowed_services
        __self__.enable_restriction = enable_restriction

@pulumi.input_type
class ServicePerimeterStatusArgs:
    access_levels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("accessLevels")
    """
    A list of AccessLevel resource names that allow resources within
    the ServicePerimeter to be accessed from the internet.
    AccessLevels listed must be in the same policy as this
    ServicePerimeter. Referencing a nonexistent AccessLevel is a
    syntax error. If no AccessLevel names are listed, resources within
    the perimeter can only be accessed via GCP calls with request
    origins within the perimeter. For Service Perimeter Bridge, must
    be empty.
    Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
    """
    resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("resources")
    """
    A list of GCP resources that are inside of the service perimeter.
    Currently only projects are allowed.
    Format: projects/{project_number}
    """
    restricted_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("restrictedServices")
    """
    GCP services that are subject to the Service Perimeter
    restrictions. Must contain a list of services. For example, if
    `storage.googleapis.com` is specified, access to the storage
    buckets inside the perimeter must meet the perimeter's access
    restrictions.
    """
    vpc_accessible_services: Optional[pulumi.Input['ServicePerimeterStatusVpcAccessibleServicesArgs']] = pulumi.input_property("vpcAccessibleServices")
    """
    Specifies how APIs are allowed to communicate within the Service
    Perimeter.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, access_levels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, restricted_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, vpc_accessible_services: Optional[pulumi.Input['ServicePerimeterStatusVpcAccessibleServicesArgs']] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] access_levels: A list of AccessLevel resource names that allow resources within
               the ServicePerimeter to be accessed from the internet.
               AccessLevels listed must be in the same policy as this
               ServicePerimeter. Referencing a nonexistent AccessLevel is a
               syntax error. If no AccessLevel names are listed, resources within
               the perimeter can only be accessed via GCP calls with request
               origins within the perimeter. For Service Perimeter Bridge, must
               be empty.
               Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
        :param pulumi.Input[List[pulumi.Input[str]]] resources: A list of GCP resources that are inside of the service perimeter.
               Currently only projects are allowed.
               Format: projects/{project_number}
        :param pulumi.Input[List[pulumi.Input[str]]] restricted_services: GCP services that are subject to the Service Perimeter
               restrictions. Must contain a list of services. For example, if
               `storage.googleapis.com` is specified, access to the storage
               buckets inside the perimeter must meet the perimeter's access
               restrictions.
        :param pulumi.Input['ServicePerimeterStatusVpcAccessibleServicesArgs'] vpc_accessible_services: Specifies how APIs are allowed to communicate within the Service
               Perimeter.  Structure is documented below.
        """
        __self__.access_levels = access_levels
        __self__.resources = resources
        __self__.restricted_services = restricted_services
        __self__.vpc_accessible_services = vpc_accessible_services

@pulumi.input_type
class ServicePerimeterStatusVpcAccessibleServicesArgs:
    allowed_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("allowedServices")
    """
    The list of APIs usable within the Service Perimeter.
    Must be empty unless `enableRestriction` is True.
    """
    enable_restriction: Optional[pulumi.Input[bool]] = pulumi.input_property("enableRestriction")
    """
    Whether to restrict API calls within the Service Perimeter to the
    list of APIs specified in 'allowedServices'.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, allowed_services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, enable_restriction: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_services: The list of APIs usable within the Service Perimeter.
               Must be empty unless `enableRestriction` is True.
        :param pulumi.Input[bool] enable_restriction: Whether to restrict API calls within the Service Perimeter to the
               list of APIs specified in 'allowedServices'.
        """
        __self__.allowed_services = allowed_services
        __self__.enable_restriction = enable_restriction

