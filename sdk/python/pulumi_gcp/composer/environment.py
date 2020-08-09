# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Environment']


class Environment(pulumi.CustomResource):
    config: pulumi.Output['outputs.EnvironmentConfig'] = pulumi.property("config")
    """
    Configuration parameters for this environment  Structure is documented below.
    """

    labels: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("labels")
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

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Name of the environment
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """

    region: pulumi.Output[Optional[str]] = pulumi.property("region")
    """
    The location or Compute Engine region for the environment.
    """

    def __init__(__self__,
                 resource_name,
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
        ### Basic Usage
        ```python
        import pulumi
        import pulumi_gcp as gcp

        test = gcp.composer.Environment("test", region="us-central1")
        ```
        ### With GKE and Compute Resource Dependencies

        **NOTE** To use service accounts, you need to give `role/composer.worker` to the service account on any resources that may be created for the environment
        (i.e. at a project level). This will probably require an explicit dependency
        on the IAM policy binding (see `projects.IAMMember` below).

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test_network = gcp.compute.Network("testNetwork", auto_create_subnetworks=False)
        test_subnetwork = gcp.compute.Subnetwork("testSubnetwork",
            ip_cidr_range="10.2.0.0/16",
            region="us-central1",
            network=test_network.id)
        test_account = gcp.service_account.Account("testAccount",
            account_id="composer-env-account",
            display_name="Test Service Account for Composer Environment")
        composer_worker = gcp.projects.IAMMember("composer-worker",
            role="roles/composer.worker",
            member=test_account.email.apply(lambda email: f"serviceAccount:{email}"))
        test_environment = gcp.composer.Environment("testEnvironment",
            region="us-central1",
            config={
                "node_count": 4,
                "node_config": {
                    "zone": "us-central1-a",
                    "machine_type": "n1-standard-1",
                    "network": test_network.id,
                    "subnetwork": test_subnetwork.id,
                    "service_account": test_account.name,
                },
            },
            opts=ResourceOptions(depends_on=[composer_worker]))
        ```
        ### With Software (Airflow) Config
        ```python
        import pulumi
        import pulumi_gcp as gcp

        test = gcp.composer.Environment("test",
            config={
                "softwareConfig": {
                    "airflowConfigOverrides": {
                        "core-loadExample": "True",
                    },
                    "env_variables": {
                        "FOO": "bar",
                    },
                    "pypiPackages": {
                        "numpy": "",
                        "scipy": "==1.1.0",
                    },
                },
            },
            region="us-central1")
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
            id: str,
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
        :param str id: The unique provider ID of the resource to lookup.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

