# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables


class ConsumerQuotaOverride(pulumi.CustomResource):
    dimensions: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("dimensions")
    """
    If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
    """
    force: pulumi.Output[Optional[bool]] = pulumi.output_property("force")
    """
    If the new quota would decrease the existing quota by more than 10%, the request is rejected.
    If `force` is `true`, that safety check is ignored.
    """
    limit: pulumi.Output[str] = pulumi.output_property("limit")
    """
    The limit on the metric, e.g. `/project/region`.
    """
    metric: pulumi.Output[str] = pulumi.output_property("metric")
    """
    The metric that should be limited, e.g. `compute.googleapis.com/cpus`.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The server-generated name of the quota override.
    """
    override_value: pulumi.Output[str] = pulumi.output_property("overrideValue")
    """
    The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    service: pulumi.Output[str] = pulumi.output_property("service")
    """
    The service that the metrics belong to, e.g. `compute.googleapis.com`.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, dimensions=None, force=None, limit=None, metric=None, override_value=None, project=None, service=None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A consumer override is applied to the consumer on its own authority to limit its own quota usage.
        Consumer overrides cannot be used to grant more quota than would be allowed by admin overrides,
        producer overrides, or the default limit of the service.

        To get more information about ConsumerQuotaOverride, see:

        * How-to Guides
            * [Getting Started](https://cloud.google.com/service-usage/docs/getting-started)
            * [REST API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1beta1/services.consumerQuotaMetrics.limits.consumerOverrides)

        ## Example Usage
        ### Consumer Quota Override

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            project_id="quota",
            org_id="123456789",
            opts=ResourceOptions(provider=google_beta))
        override = gcp.serviceusage.ConsumerQuotaOverride("override",
            project=my_project.project_id,
            service="servicemanagement.googleapis.com",
            metric="servicemanagement.googleapis.com%2Fdefault_requests",
            limit="%2Fmin%2Fproject",
            override_value="95",
            force=True,
            opts=ResourceOptions(provider=google_beta))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] dimensions: If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
        :param pulumi.Input[bool] force: If the new quota would decrease the existing quota by more than 10%, the request is rejected.
               If `force` is `true`, that safety check is ignored.
        :param pulumi.Input[str] limit: The limit on the metric, e.g. `/project/region`.
        :param pulumi.Input[str] metric: The metric that should be limited, e.g. `compute.googleapis.com/cpus`.
        :param pulumi.Input[str] override_value: The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service: The service that the metrics belong to, e.g. `compute.googleapis.com`.
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

            __props__['dimensions'] = dimensions
            __props__['force'] = force
            if limit is None:
                raise TypeError("Missing required property 'limit'")
            __props__['limit'] = limit
            if metric is None:
                raise TypeError("Missing required property 'metric'")
            __props__['metric'] = metric
            if override_value is None:
                raise TypeError("Missing required property 'override_value'")
            __props__['override_value'] = override_value
            __props__['project'] = project
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['name'] = None
        super(ConsumerQuotaOverride, __self__).__init__(
            'gcp:serviceusage/consumerQuotaOverride:ConsumerQuotaOverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dimensions=None, force=None, limit=None, metric=None, name=None, override_value=None, project=None, service=None):
        """
        Get an existing ConsumerQuotaOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] dimensions: If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.
        :param pulumi.Input[bool] force: If the new quota would decrease the existing quota by more than 10%, the request is rejected.
               If `force` is `true`, that safety check is ignored.
        :param pulumi.Input[str] limit: The limit on the metric, e.g. `/project/region`.
        :param pulumi.Input[str] metric: The metric that should be limited, e.g. `compute.googleapis.com/cpus`.
        :param pulumi.Input[str] name: The server-generated name of the quota override.
        :param pulumi.Input[str] override_value: The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service: The service that the metrics belong to, e.g. `compute.googleapis.com`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dimensions"] = dimensions
        __props__["force"] = force
        __props__["limit"] = limit
        __props__["metric"] = metric
        __props__["name"] = name
        __props__["override_value"] = override_value
        __props__["project"] = project
        __props__["service"] = service
        return ConsumerQuotaOverride(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

