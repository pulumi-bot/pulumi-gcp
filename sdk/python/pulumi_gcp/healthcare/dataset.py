# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Dataset']


class Dataset(pulumi.CustomResource):
    location: pulumi.Output[str] = pulumi.output_property("location")
    """
    The location for the Dataset.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The resource name for the Dataset.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str] = pulumi.output_property("selfLink")
    """
    The fully qualified name of this dataset
    """
    time_zone: pulumi.Output[str] = pulumi.output_property("timeZone")
    """
    The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
    "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
    (e.g., HL7 messages) where no explicit timezone is specified.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, time_zone: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A Healthcare `Dataset` is a toplevel logical grouping of `dicomStores`, `fhirStores` and `hl7V2Stores`.

        To get more information about Dataset, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets)
        * How-to Guides
            * [Creating a dataset](https://cloud.google.com/healthcare/docs/how-tos/datasets)

        ## Example Usage
        ### Healthcare Dataset Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.healthcare.Dataset("default",
            location="us-central1",
            time_zone="UTC")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location for the Dataset.
        :param pulumi.Input[str] name: The resource name for the Dataset.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] time_zone: The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
               "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
               (e.g., HL7 messages) where no explicit timezone is specified.
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['name'] = name
            __props__['project'] = project
            __props__['time_zone'] = time_zone
            __props__['self_link'] = None
        super(Dataset, __self__).__init__(
            'gcp:healthcare/dataset:Dataset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, location: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, self_link: Optional[pulumi.Input[str]] = None, time_zone: Optional[pulumi.Input[str]] = None) -> 'Dataset':
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location for the Dataset.
        :param pulumi.Input[str] name: The resource name for the Dataset.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset
        :param pulumi.Input[str] time_zone: The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
               "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
               (e.g., HL7 messages) where no explicit timezone is specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["time_zone"] = time_zone
        return Dataset(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

