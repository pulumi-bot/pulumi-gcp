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

__all__ = ['Environment']


class Environment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['EnvironmentConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An environment for running orchestration tasks.

        Environments run Apache Airflow software on Google infrastructure.

        To get more information about Environments, see:

        * [API documentation](https://cloud.google.com/composer/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/composer/docs)
            * [Configuring Shared VPC for Composer Environments](https://cloud.google.com/composer/docs/how-to/managing/configuring-shared-vpc)
        * [Apache Airflow Documentation](http://airflow.apache.org/)

        > **Warning:** We **STRONGLY** recommend  you read the [GCP guides](https://cloud.google.com/composer/docs/how-to)
          as the Environment resource requires a long deployment process and involves several layers of GCP infrastructure,
          including a Kubernetes Engine cluster, Cloud Storage, and Compute networking resources. Due to limitations of the API,
          This provider will not be able to automatically find or manage many of these underlying resources. In particular:
          * It can take up to one hour to create or update an environment resource. In addition, GCP may only detect some
            errors in configuration when they are used (e.g. ~40-50 minutes into the creation process), and is prone to limited
            error reporting. If you encounter confusing or uninformative errors, please verify your configuration is valid
            against GCP Cloud Composer before filing bugs against this provider.
          * **Environments create Google Cloud Storage buckets that do not get cleaned up automatically** on environment
            deletion. [More about Composer's use of Cloud Storage](https://cloud.google.com/composer/docs/concepts/cloud-storage).

        ## Example Usage

        ## Import

        Environment can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:composer/environment:Environment default projects/{{project}}/locations/{{region}}/environments/{{name}}
        ```

        ```sh
         $ pulumi import gcp:composer/environment:Environment default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:composer/environment:Environment default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EnvironmentConfigArgs']] config: Configuration parameters for this environment  Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-defined labels for this environment. The labels map can contain
               no more than 64 entries. Entries of the labels map are UTF8 strings
               that comply with the following restrictions:
               Label keys must be between 1 and 63 characters long and must conform
               to the following regular expression: `a-z?`.
               Label values must be between 0 and 63 characters long and must
               conform to the regular expression `(a-z?)?`.
               No more than 64 labels can be associated with a given environment.
               Both keys and values must be <= 128 bytes in size.
        :param pulumi.Input[str] name: Name of the environment
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The location or Compute Engine region for the environment.
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

            __props__['config'] = config
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
        super(Environment, __self__).__init__(
            'gcp:composer/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[pulumi.InputType['EnvironmentConfigArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EnvironmentConfigArgs']] config: Configuration parameters for this environment  Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-defined labels for this environment. The labels map can contain
               no more than 64 entries. Entries of the labels map are UTF8 strings
               that comply with the following restrictions:
               Label keys must be between 1 and 63 characters long and must conform
               to the following regular expression: `a-z?`.
               Label values must be between 0 and 63 characters long and must
               conform to the regular expression `(a-z?)?`.
               No more than 64 labels can be associated with a given environment.
               Both keys and values must be <= 128 bytes in size.
        :param pulumi.Input[str] name: Name of the environment
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The location or Compute Engine region for the environment.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["config"] = config
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.EnvironmentConfig']:
        """
        Configuration parameters for this environment  Structure is documented below.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        User-defined labels for this environment. The labels map can contain
        no more than 64 entries. Entries of the labels map are UTF8 strings
        that comply with the following restrictions:
        Label keys must be between 1 and 63 characters long and must conform
        to the following regular expression: `a-z?`.
        Label values must be between 0 and 63 characters long and must
        conform to the regular expression `(a-z?)?`.
        No more than 64 labels can be associated with a given environment.
        Both keys and values must be <= 128 bytes in size.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the environment
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The location or Compute Engine region for the environment.
        """
        return pulumi.get(self, "region")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

