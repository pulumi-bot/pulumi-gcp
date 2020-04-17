# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RegionAutoscaler(pulumi.CustomResource):
    autoscaling_policy: pulumi.Output[dict]
    """
    -
    (Required)
    The configuration parameters for the autoscaling algorithm. You can
    define one or more of the policies for an autoscaler: cpuUtilization,
    customMetricUtilizations, and loadBalancingUtilization.
    If none of these are specified, the default will be to autoscale based
    on cpuUtilization to 0.6 or 60%.  Structure is documented below.

      * `cooldownPeriod` (`float`) - -
        (Optional)
        The number of seconds that the autoscaler should wait before it
        starts collecting information from a new instance. This prevents
        the autoscaler from collecting information when the instance is
        initializing, during which the collected usage would not be
        reliable. The default time autoscaler waits is 60 seconds.
        Virtual machine initialization times might vary because of
        numerous factors. We recommend that you test how long an
        instance may take to initialize. To do this, create an instance
        and time the startup process.
      * `cpuUtilization` (`dict`) - -
        (Optional)
        Defines the CPU utilization policy that allows the autoscaler to
        scale based on the average CPU utilization of a managed instance
        group.  Structure is documented below.
        * `target` (`float`) - -
          (Required)
          URL of the managed instance group that this autoscaler will scale.

      * `loadBalancingUtilization` (`dict`) - -
        (Optional)
        Configuration parameters of autoscaling based on a load balancer.  Structure is documented below.
        * `target` (`float`) - -
          (Required)
          URL of the managed instance group that this autoscaler will scale.

      * `maxReplicas` (`float`) - -
        (Required)
        The maximum number of instances that the autoscaler can scale up
        to. This is required when creating or updating an autoscaler. The
        maximum number of replicas should not be lower than minimal number
        of replicas.
      * `metrics` (`list`) - -
        (Optional)
        Configuration parameters of autoscaling based on a custom metric.  Structure is documented below.
        * `filter` (`str`) - -
          (Optional)
          A filter string to be used as the filter string for
          a Stackdriver Monitoring TimeSeries.list API call.
          This filter is used to select a specific TimeSeries for
          the purpose of autoscaling and to determine whether the metric
          is exporting per-instance or per-group data.
          You can only use the AND operator for joining selectors.
          You can only use direct equality comparison operator (=) without
          any functions for each selector.
          You can specify the metric in both the filter string and in the
          metric field. However, if specified in both places, the metric must
          be identical.
          The monitored resource type determines what kind of values are
          expected for the metric. If it is a gce_instance, the autoscaler
          expects the metric to include a separate TimeSeries for each
          instance in a group. In such a case, you cannot filter on resource
          labels.
          If the resource type is any other value, the autoscaler expects
          this metric to contain values that apply to the entire autoscaled
          instance group and resource label filtering can be performed to
          point autoscaler at the correct TimeSeries to scale upon.
          This is called a per-group metric for the purpose of autoscaling.
          If not specified, the type defaults to gce_instance.
          You should provide a filter that is selective enough to pick just
          one TimeSeries for the autoscaled group or for each of the instances
          (if you are using gce_instance resource type). If multiple
          TimeSeries are returned upon the query execution, the autoscaler
          will sum their respective values to obtain its scaling value.
        * `name` (`str`) - -
          (Required)
          Name of the resource. The name must be 1-63 characters long and match
          the regular expression `a-z?` which means the
          first character must be a lowercase letter, and all following
          characters must be a dash, lowercase letter, or digit, except the last
          character, which cannot be a dash.
        * `singleInstanceAssignment` (`float`) - -
          (Optional)
          If scaling is based on a per-group metric value that represents the
          total amount of work to be done or resource usage, set this value to
          an amount assigned for a single instance of the scaled group.
          The autoscaler will keep the number of instances proportional to the
          value of this metric, the metric itself should not change value due
          to group resizing.
          For example, a good metric to use with the target is
          `pubsub.googleapis.com/subscription/num_undelivered_messages`
          or a custom metric exporting the total number of requests coming to
          your instances.
          A bad example would be a metric exporting an average or median
          latency, since this value can't include a chunk assignable to a
          single instance, it could be better used with utilization_target
          instead.
        * `target` (`float`) - -
          (Required)
          URL of the managed instance group that this autoscaler will scale.
        * `type` (`str`) - -
          (Optional)
          Defines how target utilization value is expressed for a
          Stackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND,
          or DELTA_PER_MINUTE.

      * `minReplicas` (`float`) - -
        (Required)
        The minimum number of replicas that the autoscaler can scale down
        to. This cannot be less than 0. If not provided, autoscaler will
        choose a default value depending on maximum number of instances
        allowed.
    """
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    -
    (Optional)
    An optional description of this resource.
    """
    name: pulumi.Output[str]
    """
    -
    (Required)
    Name of the resource. The name must be 1-63 characters long and match
    the regular expression `a-z?` which means the
    first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last
    character, which cannot be a dash.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    -
    (Optional)
    URL of the region where the instance group resides.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    target: pulumi.Output[str]
    """
    -
    (Required)
    URL of the managed instance group that this autoscaler will scale.
    """
    def __init__(__self__, resource_name, opts=None, autoscaling_policy=None, description=None, name=None, project=None, region=None, target=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents an Autoscaler resource.

        Autoscalers allow you to automatically scale virtual machine instances in
        managed instance groups according to an autoscaling policy that you
        define.


        To get more information about RegionAutoscaler, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionAutoscalers)
        * How-to Guides
            * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] autoscaling_policy: -
               (Required)
               The configuration parameters for the autoscaling algorithm. You can
               define one or more of the policies for an autoscaler: cpuUtilization,
               customMetricUtilizations, and loadBalancingUtilization.
               If none of these are specified, the default will be to autoscale based
               on cpuUtilization to 0.6 or 60%.  Structure is documented below.
        :param pulumi.Input[str] description: -
               (Optional)
               An optional description of this resource.
        :param pulumi.Input[str] name: -
               (Required)
               Name of the resource. The name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: -
               (Optional)
               URL of the region where the instance group resides.
        :param pulumi.Input[str] target: -
               (Required)
               URL of the managed instance group that this autoscaler will scale.

        The **autoscaling_policy** object supports the following:

          * `cooldownPeriod` (`pulumi.Input[float]`) - -
            (Optional)
            The number of seconds that the autoscaler should wait before it
            starts collecting information from a new instance. This prevents
            the autoscaler from collecting information when the instance is
            initializing, during which the collected usage would not be
            reliable. The default time autoscaler waits is 60 seconds.
            Virtual machine initialization times might vary because of
            numerous factors. We recommend that you test how long an
            instance may take to initialize. To do this, create an instance
            and time the startup process.
          * `cpuUtilization` (`pulumi.Input[dict]`) - -
            (Optional)
            Defines the CPU utilization policy that allows the autoscaler to
            scale based on the average CPU utilization of a managed instance
            group.  Structure is documented below.
            * `target` (`pulumi.Input[float]`) - -
              (Required)
              URL of the managed instance group that this autoscaler will scale.

          * `loadBalancingUtilization` (`pulumi.Input[dict]`) - -
            (Optional)
            Configuration parameters of autoscaling based on a load balancer.  Structure is documented below.
            * `target` (`pulumi.Input[float]`) - -
              (Required)
              URL of the managed instance group that this autoscaler will scale.

          * `maxReplicas` (`pulumi.Input[float]`) - -
            (Required)
            The maximum number of instances that the autoscaler can scale up
            to. This is required when creating or updating an autoscaler. The
            maximum number of replicas should not be lower than minimal number
            of replicas.
          * `metrics` (`pulumi.Input[list]`) - -
            (Optional)
            Configuration parameters of autoscaling based on a custom metric.  Structure is documented below.
            * `filter` (`pulumi.Input[str]`) - -
              (Optional)
              A filter string to be used as the filter string for
              a Stackdriver Monitoring TimeSeries.list API call.
              This filter is used to select a specific TimeSeries for
              the purpose of autoscaling and to determine whether the metric
              is exporting per-instance or per-group data.
              You can only use the AND operator for joining selectors.
              You can only use direct equality comparison operator (=) without
              any functions for each selector.
              You can specify the metric in both the filter string and in the
              metric field. However, if specified in both places, the metric must
              be identical.
              The monitored resource type determines what kind of values are
              expected for the metric. If it is a gce_instance, the autoscaler
              expects the metric to include a separate TimeSeries for each
              instance in a group. In such a case, you cannot filter on resource
              labels.
              If the resource type is any other value, the autoscaler expects
              this metric to contain values that apply to the entire autoscaled
              instance group and resource label filtering can be performed to
              point autoscaler at the correct TimeSeries to scale upon.
              This is called a per-group metric for the purpose of autoscaling.
              If not specified, the type defaults to gce_instance.
              You should provide a filter that is selective enough to pick just
              one TimeSeries for the autoscaled group or for each of the instances
              (if you are using gce_instance resource type). If multiple
              TimeSeries are returned upon the query execution, the autoscaler
              will sum their respective values to obtain its scaling value.
            * `name` (`pulumi.Input[str]`) - -
              (Required)
              Name of the resource. The name must be 1-63 characters long and match
              the regular expression `a-z?` which means the
              first character must be a lowercase letter, and all following
              characters must be a dash, lowercase letter, or digit, except the last
              character, which cannot be a dash.
            * `singleInstanceAssignment` (`pulumi.Input[float]`) - -
              (Optional)
              If scaling is based on a per-group metric value that represents the
              total amount of work to be done or resource usage, set this value to
              an amount assigned for a single instance of the scaled group.
              The autoscaler will keep the number of instances proportional to the
              value of this metric, the metric itself should not change value due
              to group resizing.
              For example, a good metric to use with the target is
              `pubsub.googleapis.com/subscription/num_undelivered_messages`
              or a custom metric exporting the total number of requests coming to
              your instances.
              A bad example would be a metric exporting an average or median
              latency, since this value can't include a chunk assignable to a
              single instance, it could be better used with utilization_target
              instead.
            * `target` (`pulumi.Input[float]`) - -
              (Required)
              URL of the managed instance group that this autoscaler will scale.
            * `type` (`pulumi.Input[str]`) - -
              (Optional)
              Defines how target utilization value is expressed for a
              Stackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND,
              or DELTA_PER_MINUTE.

          * `minReplicas` (`pulumi.Input[float]`) - -
            (Required)
            The minimum number of replicas that the autoscaler can scale down
            to. This cannot be less than 0. If not provided, autoscaler will
            choose a default value depending on maximum number of instances
            allowed.
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

            if autoscaling_policy is None:
                raise TypeError("Missing required property 'autoscaling_policy'")
            __props__['autoscaling_policy'] = autoscaling_policy
            __props__['description'] = description
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
            if target is None:
                raise TypeError("Missing required property 'target'")
            __props__['target'] = target
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(RegionAutoscaler, __self__).__init__(
            'gcp:compute/regionAutoscaler:RegionAutoscaler',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, autoscaling_policy=None, creation_timestamp=None, description=None, name=None, project=None, region=None, self_link=None, target=None):
        """
        Get an existing RegionAutoscaler resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] autoscaling_policy: -
               (Required)
               The configuration parameters for the autoscaling algorithm. You can
               define one or more of the policies for an autoscaler: cpuUtilization,
               customMetricUtilizations, and loadBalancingUtilization.
               If none of these are specified, the default will be to autoscale based
               on cpuUtilization to 0.6 or 60%.  Structure is documented below.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: -
               (Optional)
               An optional description of this resource.
        :param pulumi.Input[str] name: -
               (Required)
               Name of the resource. The name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: -
               (Optional)
               URL of the region where the instance group resides.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] target: -
               (Required)
               URL of the managed instance group that this autoscaler will scale.

        The **autoscaling_policy** object supports the following:

          * `cooldownPeriod` (`pulumi.Input[float]`) - -
            (Optional)
            The number of seconds that the autoscaler should wait before it
            starts collecting information from a new instance. This prevents
            the autoscaler from collecting information when the instance is
            initializing, during which the collected usage would not be
            reliable. The default time autoscaler waits is 60 seconds.
            Virtual machine initialization times might vary because of
            numerous factors. We recommend that you test how long an
            instance may take to initialize. To do this, create an instance
            and time the startup process.
          * `cpuUtilization` (`pulumi.Input[dict]`) - -
            (Optional)
            Defines the CPU utilization policy that allows the autoscaler to
            scale based on the average CPU utilization of a managed instance
            group.  Structure is documented below.
            * `target` (`pulumi.Input[float]`) - -
              (Required)
              URL of the managed instance group that this autoscaler will scale.

          * `loadBalancingUtilization` (`pulumi.Input[dict]`) - -
            (Optional)
            Configuration parameters of autoscaling based on a load balancer.  Structure is documented below.
            * `target` (`pulumi.Input[float]`) - -
              (Required)
              URL of the managed instance group that this autoscaler will scale.

          * `maxReplicas` (`pulumi.Input[float]`) - -
            (Required)
            The maximum number of instances that the autoscaler can scale up
            to. This is required when creating or updating an autoscaler. The
            maximum number of replicas should not be lower than minimal number
            of replicas.
          * `metrics` (`pulumi.Input[list]`) - -
            (Optional)
            Configuration parameters of autoscaling based on a custom metric.  Structure is documented below.
            * `filter` (`pulumi.Input[str]`) - -
              (Optional)
              A filter string to be used as the filter string for
              a Stackdriver Monitoring TimeSeries.list API call.
              This filter is used to select a specific TimeSeries for
              the purpose of autoscaling and to determine whether the metric
              is exporting per-instance or per-group data.
              You can only use the AND operator for joining selectors.
              You can only use direct equality comparison operator (=) without
              any functions for each selector.
              You can specify the metric in both the filter string and in the
              metric field. However, if specified in both places, the metric must
              be identical.
              The monitored resource type determines what kind of values are
              expected for the metric. If it is a gce_instance, the autoscaler
              expects the metric to include a separate TimeSeries for each
              instance in a group. In such a case, you cannot filter on resource
              labels.
              If the resource type is any other value, the autoscaler expects
              this metric to contain values that apply to the entire autoscaled
              instance group and resource label filtering can be performed to
              point autoscaler at the correct TimeSeries to scale upon.
              This is called a per-group metric for the purpose of autoscaling.
              If not specified, the type defaults to gce_instance.
              You should provide a filter that is selective enough to pick just
              one TimeSeries for the autoscaled group or for each of the instances
              (if you are using gce_instance resource type). If multiple
              TimeSeries are returned upon the query execution, the autoscaler
              will sum their respective values to obtain its scaling value.
            * `name` (`pulumi.Input[str]`) - -
              (Required)
              Name of the resource. The name must be 1-63 characters long and match
              the regular expression `a-z?` which means the
              first character must be a lowercase letter, and all following
              characters must be a dash, lowercase letter, or digit, except the last
              character, which cannot be a dash.
            * `singleInstanceAssignment` (`pulumi.Input[float]`) - -
              (Optional)
              If scaling is based on a per-group metric value that represents the
              total amount of work to be done or resource usage, set this value to
              an amount assigned for a single instance of the scaled group.
              The autoscaler will keep the number of instances proportional to the
              value of this metric, the metric itself should not change value due
              to group resizing.
              For example, a good metric to use with the target is
              `pubsub.googleapis.com/subscription/num_undelivered_messages`
              or a custom metric exporting the total number of requests coming to
              your instances.
              A bad example would be a metric exporting an average or median
              latency, since this value can't include a chunk assignable to a
              single instance, it could be better used with utilization_target
              instead.
            * `target` (`pulumi.Input[float]`) - -
              (Required)
              URL of the managed instance group that this autoscaler will scale.
            * `type` (`pulumi.Input[str]`) - -
              (Optional)
              Defines how target utilization value is expressed for a
              Stackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND,
              or DELTA_PER_MINUTE.

          * `minReplicas` (`pulumi.Input[float]`) - -
            (Required)
            The minimum number of replicas that the autoscaler can scale down
            to. This cannot be less than 0. If not provided, autoscaler will
            choose a default value depending on maximum number of instances
            allowed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["autoscaling_policy"] = autoscaling_policy
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["target"] = target
        return RegionAutoscaler(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

