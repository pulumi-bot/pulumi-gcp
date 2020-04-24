# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ServicePerimeter(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    """
    Time the AccessPolicy was created in UTC.
    """
    description: pulumi.Output[str]
    """
    -
    (Optional)
    Description of the ServicePerimeter and its use. Does not affect
    behavior.
    """
    name: pulumi.Output[str]
    """
    -
    (Required)
    Resource name for the ServicePerimeter. The short_name component must
    begin with a letter and only include alphanumeric and '_'.
    Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
    """
    parent: pulumi.Output[str]
    """
    -
    (Required)
    The AccessPolicy this ServicePerimeter lives in.
    Format: accessPolicies/{policy_id}
    """
    perimeter_type: pulumi.Output[str]
    """
    -
    (Optional)
    Specifies the type of the Perimeter. There are two types: regular and
    bridge. Regular Service Perimeter contains resources, access levels,
    and restricted services. Every resource can be in at most
    ONE regular Service Perimeter.
    In addition to being in a regular service perimeter, a resource can also
    be in zero or more perimeter bridges. A perimeter bridge only contains
    resources. Cross project operations are permitted if all effected
    resources share some perimeter (whether bridge or regular). Perimeter
    Bridge does not contain access levels or services: those are governed
    entirely by the regular perimeter that resource is in.
    Perimeter Bridges are typically useful when building more complex
    topologies with many independent perimeters that need to share some data
    with a common perimeter, but should not be able to share data among
    themselves.
    """
    spec: pulumi.Output[dict]
    """
    -
    (Optional)
    Proposed (or dry run) ServicePerimeter configuration.
    This configuration allows to specify and test ServicePerimeter configuration
    without enforcing actual access restrictions. Only allowed to be set when
    the `useExplicitDryRunSpec` flag is set.  Structure is documented below.

      * `accessLevels` (`list`) - -
        (Optional)
        A list of AccessLevel resource names that allow resources within
        the ServicePerimeter to be accessed from the internet.
        AccessLevels listed must be in the same policy as this
        ServicePerimeter. Referencing a nonexistent AccessLevel is a
        syntax error. If no AccessLevel names are listed, resources within
        the perimeter can only be accessed via GCP calls with request
        origins within the perimeter. For Service Perimeter Bridge, must
        be empty.
        Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
      * `resources` (`list`) - -
        (Optional)
        A list of GCP resources that are inside of the service perimeter.
        Currently only projects are allowed.
        Format: projects/{project_number}
      * `restrictedServices` (`list`) - -
        (Optional)
        GCP services that are subject to the Service Perimeter
        restrictions. Must contain a list of services. For example, if
        `storage.googleapis.com` is specified, access to the storage
        buckets inside the perimeter must meet the perimeter's access
        restrictions.
      * `vpcAccessibleServices` (`dict`) - -
        (Optional)
        Specifies how APIs are allowed to communicate within the Service
        Perimeter.  Structure is documented below.
        * `allowedServices` (`list`) - -
          (Optional)
          The list of APIs usable within the Service Perimeter.
          Must be empty unless `enableRestriction` is True.
        * `enableRestriction` (`bool`) - -
          (Optional)
          Whether to restrict API calls within the Service Perimeter to the
          list of APIs specified in 'allowedServices'.
    """
    status: pulumi.Output[dict]
    """
    -
    (Optional)
    ServicePerimeter configuration. Specifies sets of resources,
    restricted services and access levels that determine
    perimeter content and boundaries.  Structure is documented below.

      * `accessLevels` (`list`) - -
        (Optional)
        A list of AccessLevel resource names that allow resources within
        the ServicePerimeter to be accessed from the internet.
        AccessLevels listed must be in the same policy as this
        ServicePerimeter. Referencing a nonexistent AccessLevel is a
        syntax error. If no AccessLevel names are listed, resources within
        the perimeter can only be accessed via GCP calls with request
        origins within the perimeter. For Service Perimeter Bridge, must
        be empty.
        Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
      * `resources` (`list`) - -
        (Optional)
        A list of GCP resources that are inside of the service perimeter.
        Currently only projects are allowed.
        Format: projects/{project_number}
      * `restrictedServices` (`list`) - -
        (Optional)
        GCP services that are subject to the Service Perimeter
        restrictions. Must contain a list of services. For example, if
        `storage.googleapis.com` is specified, access to the storage
        buckets inside the perimeter must meet the perimeter's access
        restrictions.
      * `vpcAccessibleServices` (`dict`) - -
        (Optional)
        Specifies how APIs are allowed to communicate within the Service
        Perimeter.  Structure is documented below.
        * `allowedServices` (`list`) - -
          (Optional)
          The list of APIs usable within the Service Perimeter.
          Must be empty unless `enableRestriction` is True.
        * `enableRestriction` (`bool`) - -
          (Optional)
          Whether to restrict API calls within the Service Perimeter to the
          list of APIs specified in 'allowedServices'.
    """
    title: pulumi.Output[str]
    """
    -
    (Required)
    Human readable title. Must be unique within the Policy.
    """
    update_time: pulumi.Output[str]
    """
    Time the AccessPolicy was updated in UTC.
    """
    use_explicit_dry_run_spec: pulumi.Output[bool]
    """
    -
    (Optional)
    Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
    for all Service Perimeters, and that spec is identical to the status for those
    Service Perimeters. When this flag is set, it inhibits the generation of the
    implicit spec, thereby allowing the user to explicitly provide a
    configuration ("spec") to use in a dry-run version of the Service Perimeter.
    This allows the user to test changes to the enforced config ("status") without
    actually enforcing them. This testing is done through analyzing the differences
    between currently enforced and suggested restrictions. useExplicitDryRunSpec must
    bet set to True if any of the fields in the spec are set to non-default values.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, parent=None, perimeter_type=None, spec=None, status=None, title=None, use_explicit_dry_run_spec=None, __props__=None, __name__=None, __opts__=None):
        """
        ServicePerimeter describes a set of GCP resources which can freely import
        and export data amongst themselves, but not export outside of the
        ServicePerimeter. If a request with a source within this ServicePerimeter
        has a target outside of the ServicePerimeter, the request will be blocked.
        Otherwise the request is allowed. There are two types of Service Perimeter
        - Regular and Bridge. Regular Service Perimeters cannot overlap, a single
        GCP project can only belong to a single regular Service Perimeter. Service
        Perimeter Bridges can contain only GCP projects as members, a single GCP
        project may belong to multiple Service Perimeter Bridges.


        To get more information about ServicePerimeter, see:

        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
        * How-to Guides
            * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: -
               (Optional)
               Description of the ServicePerimeter and its use. Does not affect
               behavior.
        :param pulumi.Input[str] name: -
               (Required)
               Resource name for the ServicePerimeter. The short_name component must
               begin with a letter and only include alphanumeric and '_'.
               Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
        :param pulumi.Input[str] parent: -
               (Required)
               The AccessPolicy this ServicePerimeter lives in.
               Format: accessPolicies/{policy_id}
        :param pulumi.Input[str] perimeter_type: -
               (Optional)
               Specifies the type of the Perimeter. There are two types: regular and
               bridge. Regular Service Perimeter contains resources, access levels,
               and restricted services. Every resource can be in at most
               ONE regular Service Perimeter.
               In addition to being in a regular service perimeter, a resource can also
               be in zero or more perimeter bridges. A perimeter bridge only contains
               resources. Cross project operations are permitted if all effected
               resources share some perimeter (whether bridge or regular). Perimeter
               Bridge does not contain access levels or services: those are governed
               entirely by the regular perimeter that resource is in.
               Perimeter Bridges are typically useful when building more complex
               topologies with many independent perimeters that need to share some data
               with a common perimeter, but should not be able to share data among
               themselves.
        :param pulumi.Input[dict] spec: -
               (Optional)
               Proposed (or dry run) ServicePerimeter configuration.
               This configuration allows to specify and test ServicePerimeter configuration
               without enforcing actual access restrictions. Only allowed to be set when
               the `useExplicitDryRunSpec` flag is set.  Structure is documented below.
        :param pulumi.Input[dict] status: -
               (Optional)
               ServicePerimeter configuration. Specifies sets of resources,
               restricted services and access levels that determine
               perimeter content and boundaries.  Structure is documented below.
        :param pulumi.Input[str] title: -
               (Required)
               Human readable title. Must be unique within the Policy.
        :param pulumi.Input[bool] use_explicit_dry_run_spec: -
               (Optional)
               Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
               for all Service Perimeters, and that spec is identical to the status for those
               Service Perimeters. When this flag is set, it inhibits the generation of the
               implicit spec, thereby allowing the user to explicitly provide a
               configuration ("spec") to use in a dry-run version of the Service Perimeter.
               This allows the user to test changes to the enforced config ("status") without
               actually enforcing them. This testing is done through analyzing the differences
               between currently enforced and suggested restrictions. useExplicitDryRunSpec must
               bet set to True if any of the fields in the spec are set to non-default values.

        The **spec** object supports the following:

          * `accessLevels` (`pulumi.Input[list]`) - -
            (Optional)
            A list of AccessLevel resource names that allow resources within
            the ServicePerimeter to be accessed from the internet.
            AccessLevels listed must be in the same policy as this
            ServicePerimeter. Referencing a nonexistent AccessLevel is a
            syntax error. If no AccessLevel names are listed, resources within
            the perimeter can only be accessed via GCP calls with request
            origins within the perimeter. For Service Perimeter Bridge, must
            be empty.
            Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
          * `resources` (`pulumi.Input[list]`) - -
            (Optional)
            A list of GCP resources that are inside of the service perimeter.
            Currently only projects are allowed.
            Format: projects/{project_number}
          * `restrictedServices` (`pulumi.Input[list]`) - -
            (Optional)
            GCP services that are subject to the Service Perimeter
            restrictions. Must contain a list of services. For example, if
            `storage.googleapis.com` is specified, access to the storage
            buckets inside the perimeter must meet the perimeter's access
            restrictions.
          * `vpcAccessibleServices` (`pulumi.Input[dict]`) - -
            (Optional)
            Specifies how APIs are allowed to communicate within the Service
            Perimeter.  Structure is documented below.
            * `allowedServices` (`pulumi.Input[list]`) - -
              (Optional)
              The list of APIs usable within the Service Perimeter.
              Must be empty unless `enableRestriction` is True.
            * `enableRestriction` (`pulumi.Input[bool]`) - -
              (Optional)
              Whether to restrict API calls within the Service Perimeter to the
              list of APIs specified in 'allowedServices'.

        The **status** object supports the following:

          * `accessLevels` (`pulumi.Input[list]`) - -
            (Optional)
            A list of AccessLevel resource names that allow resources within
            the ServicePerimeter to be accessed from the internet.
            AccessLevels listed must be in the same policy as this
            ServicePerimeter. Referencing a nonexistent AccessLevel is a
            syntax error. If no AccessLevel names are listed, resources within
            the perimeter can only be accessed via GCP calls with request
            origins within the perimeter. For Service Perimeter Bridge, must
            be empty.
            Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
          * `resources` (`pulumi.Input[list]`) - -
            (Optional)
            A list of GCP resources that are inside of the service perimeter.
            Currently only projects are allowed.
            Format: projects/{project_number}
          * `restrictedServices` (`pulumi.Input[list]`) - -
            (Optional)
            GCP services that are subject to the Service Perimeter
            restrictions. Must contain a list of services. For example, if
            `storage.googleapis.com` is specified, access to the storage
            buckets inside the perimeter must meet the perimeter's access
            restrictions.
          * `vpcAccessibleServices` (`pulumi.Input[dict]`) - -
            (Optional)
            Specifies how APIs are allowed to communicate within the Service
            Perimeter.  Structure is documented below.
            * `allowedServices` (`pulumi.Input[list]`) - -
              (Optional)
              The list of APIs usable within the Service Perimeter.
              Must be empty unless `enableRestriction` is True.
            * `enableRestriction` (`pulumi.Input[bool]`) - -
              (Optional)
              Whether to restrict API calls within the Service Perimeter to the
              list of APIs specified in 'allowedServices'.
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
            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['perimeter_type'] = perimeter_type
            __props__['spec'] = spec
            __props__['status'] = status
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            __props__['use_explicit_dry_run_spec'] = use_explicit_dry_run_spec
            __props__['create_time'] = None
            __props__['update_time'] = None
        super(ServicePerimeter, __self__).__init__(
            'gcp:accesscontextmanager/servicePerimeter:ServicePerimeter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, description=None, name=None, parent=None, perimeter_type=None, spec=None, status=None, title=None, update_time=None, use_explicit_dry_run_spec=None):
        """
        Get an existing ServicePerimeter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Time the AccessPolicy was created in UTC.
        :param pulumi.Input[str] description: -
               (Optional)
               Description of the ServicePerimeter and its use. Does not affect
               behavior.
        :param pulumi.Input[str] name: -
               (Required)
               Resource name for the ServicePerimeter. The short_name component must
               begin with a letter and only include alphanumeric and '_'.
               Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
        :param pulumi.Input[str] parent: -
               (Required)
               The AccessPolicy this ServicePerimeter lives in.
               Format: accessPolicies/{policy_id}
        :param pulumi.Input[str] perimeter_type: -
               (Optional)
               Specifies the type of the Perimeter. There are two types: regular and
               bridge. Regular Service Perimeter contains resources, access levels,
               and restricted services. Every resource can be in at most
               ONE regular Service Perimeter.
               In addition to being in a regular service perimeter, a resource can also
               be in zero or more perimeter bridges. A perimeter bridge only contains
               resources. Cross project operations are permitted if all effected
               resources share some perimeter (whether bridge or regular). Perimeter
               Bridge does not contain access levels or services: those are governed
               entirely by the regular perimeter that resource is in.
               Perimeter Bridges are typically useful when building more complex
               topologies with many independent perimeters that need to share some data
               with a common perimeter, but should not be able to share data among
               themselves.
        :param pulumi.Input[dict] spec: -
               (Optional)
               Proposed (or dry run) ServicePerimeter configuration.
               This configuration allows to specify and test ServicePerimeter configuration
               without enforcing actual access restrictions. Only allowed to be set when
               the `useExplicitDryRunSpec` flag is set.  Structure is documented below.
        :param pulumi.Input[dict] status: -
               (Optional)
               ServicePerimeter configuration. Specifies sets of resources,
               restricted services and access levels that determine
               perimeter content and boundaries.  Structure is documented below.
        :param pulumi.Input[str] title: -
               (Required)
               Human readable title. Must be unique within the Policy.
        :param pulumi.Input[str] update_time: Time the AccessPolicy was updated in UTC.
        :param pulumi.Input[bool] use_explicit_dry_run_spec: -
               (Optional)
               Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
               for all Service Perimeters, and that spec is identical to the status for those
               Service Perimeters. When this flag is set, it inhibits the generation of the
               implicit spec, thereby allowing the user to explicitly provide a
               configuration ("spec") to use in a dry-run version of the Service Perimeter.
               This allows the user to test changes to the enforced config ("status") without
               actually enforcing them. This testing is done through analyzing the differences
               between currently enforced and suggested restrictions. useExplicitDryRunSpec must
               bet set to True if any of the fields in the spec are set to non-default values.

        The **spec** object supports the following:

          * `accessLevels` (`pulumi.Input[list]`) - -
            (Optional)
            A list of AccessLevel resource names that allow resources within
            the ServicePerimeter to be accessed from the internet.
            AccessLevels listed must be in the same policy as this
            ServicePerimeter. Referencing a nonexistent AccessLevel is a
            syntax error. If no AccessLevel names are listed, resources within
            the perimeter can only be accessed via GCP calls with request
            origins within the perimeter. For Service Perimeter Bridge, must
            be empty.
            Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
          * `resources` (`pulumi.Input[list]`) - -
            (Optional)
            A list of GCP resources that are inside of the service perimeter.
            Currently only projects are allowed.
            Format: projects/{project_number}
          * `restrictedServices` (`pulumi.Input[list]`) - -
            (Optional)
            GCP services that are subject to the Service Perimeter
            restrictions. Must contain a list of services. For example, if
            `storage.googleapis.com` is specified, access to the storage
            buckets inside the perimeter must meet the perimeter's access
            restrictions.
          * `vpcAccessibleServices` (`pulumi.Input[dict]`) - -
            (Optional)
            Specifies how APIs are allowed to communicate within the Service
            Perimeter.  Structure is documented below.
            * `allowedServices` (`pulumi.Input[list]`) - -
              (Optional)
              The list of APIs usable within the Service Perimeter.
              Must be empty unless `enableRestriction` is True.
            * `enableRestriction` (`pulumi.Input[bool]`) - -
              (Optional)
              Whether to restrict API calls within the Service Perimeter to the
              list of APIs specified in 'allowedServices'.

        The **status** object supports the following:

          * `accessLevels` (`pulumi.Input[list]`) - -
            (Optional)
            A list of AccessLevel resource names that allow resources within
            the ServicePerimeter to be accessed from the internet.
            AccessLevels listed must be in the same policy as this
            ServicePerimeter. Referencing a nonexistent AccessLevel is a
            syntax error. If no AccessLevel names are listed, resources within
            the perimeter can only be accessed via GCP calls with request
            origins within the perimeter. For Service Perimeter Bridge, must
            be empty.
            Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
          * `resources` (`pulumi.Input[list]`) - -
            (Optional)
            A list of GCP resources that are inside of the service perimeter.
            Currently only projects are allowed.
            Format: projects/{project_number}
          * `restrictedServices` (`pulumi.Input[list]`) - -
            (Optional)
            GCP services that are subject to the Service Perimeter
            restrictions. Must contain a list of services. For example, if
            `storage.googleapis.com` is specified, access to the storage
            buckets inside the perimeter must meet the perimeter's access
            restrictions.
          * `vpcAccessibleServices` (`pulumi.Input[dict]`) - -
            (Optional)
            Specifies how APIs are allowed to communicate within the Service
            Perimeter.  Structure is documented below.
            * `allowedServices` (`pulumi.Input[list]`) - -
              (Optional)
              The list of APIs usable within the Service Perimeter.
              Must be empty unless `enableRestriction` is True.
            * `enableRestriction` (`pulumi.Input[bool]`) - -
              (Optional)
              Whether to restrict API calls within the Service Perimeter to the
              list of APIs specified in 'allowedServices'.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["name"] = name
        __props__["parent"] = parent
        __props__["perimeter_type"] = perimeter_type
        __props__["spec"] = spec
        __props__["status"] = status
        __props__["title"] = title
        __props__["update_time"] = update_time
        __props__["use_explicit_dry_run_spec"] = use_explicit_dry_run_spec
        return ServicePerimeter(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

