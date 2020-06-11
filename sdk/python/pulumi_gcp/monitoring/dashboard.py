# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Dashboard(pulumi.CustomResource):
    dashboard_json: pulumi.Output[str]
    """
    The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
    The representation of an existing dashboard can be found by using the [API Explorer](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards/get)
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, dashboard_json=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        A Google Stackdriver dashboard. Dashboards define the content and layout of pages in the Stackdriver web application.

        To get more information about Dashboards, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/monitoring/dashboards)

        ## Example Usage - Monitoring Dashboard Basic


        ```python
        import pulumi
        import pulumi_gcp as gcp

        dashboard = gcp.monitoring.Dashboard("dashboard", dashboard_json=\"\"\"{
          "displayName": "Demo Dashboard",
          "gridLayout": {
            "widgets": [
              {
                "blank": {}
              }
            ]
          }
        }


        \"\"\")
        ```

        ## Example Usage - Monitoring Dashboard GridLayout


        ```python
        import pulumi
        import pulumi_gcp as gcp

        dashboard = gcp.monitoring.Dashboard("dashboard", dashboard_json=\"\"\"{
          "displayName": "Grid Layout Example",
          "gridLayout": {
            "columns": "2",
            "widgets": [
              {
                "title": "Widget 1",
                "xyChart": {
                  "dataSets": [{
                    "timeSeriesQuery": {
                      "timeSeriesFilter": {
                        "filter": "metric.type=\"agent.googleapis.com/nginx/connections/accepted_count\"",
                        "aggregation": {
                          "perSeriesAligner": "ALIGN_RATE"
                        }
                      },
                      "unitOverride": "1"
                    },
                    "plotType": "LINE"
                  }],
                  "timeshiftDuration": "0s",
                  "yAxis": {
                    "label": "y1Axis",
                    "scale": "LINEAR"
                  }
                }
              },
              {
                "text": {
                  "content": "Widget 2",
                  "format": "MARKDOWN"
                }
              },
              {
                "title": "Widget 3",
                "xyChart": {
                  "dataSets": [{
                    "timeSeriesQuery": {
                      "timeSeriesFilter": {
                        "filter": "metric.type=\"agent.googleapis.com/nginx/connections/accepted_count\"",
                        "aggregation": {
                          "perSeriesAligner": "ALIGN_RATE"
                        }
                      },
                      "unitOverride": "1"
                    },
                    "plotType": "STACKED_BAR"
                  }],
                  "timeshiftDuration": "0s",
                  "yAxis": {
                    "label": "y1Axis",
                    "scale": "LINEAR"
                  }
                }
              }
            ]
          }
        }


        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dashboard_json: The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
               The representation of an existing dashboard can be found by using the [API Explorer](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards/get)
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

            if dashboard_json is None:
                raise TypeError("Missing required property 'dashboard_json'")
            __props__['dashboard_json'] = dashboard_json
            __props__['project'] = project
        super(Dashboard, __self__).__init__(
            'gcp:monitoring/dashboard:Dashboard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dashboard_json=None, project=None):
        """
        Get an existing Dashboard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dashboard_json: The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
               The representation of an existing dashboard can be found by using the [API Explorer](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards/get)
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dashboard_json"] = dashboard_json
        __props__["project"] = project
        return Dashboard(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
