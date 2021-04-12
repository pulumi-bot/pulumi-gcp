# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DatasetArgs', 'Dataset']

@pulumi.input_type
class DatasetArgs:
    def __init__(__self__, *,
                 dataset_id: pulumi.Input[str],
                 accesses: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetAccessArgs']]]] = None,
                 default_encryption_configuration: Optional[pulumi.Input['DatasetDefaultEncryptionConfigurationArgs']] = None,
                 default_partition_expiration_ms: Optional[pulumi.Input[int]] = None,
                 default_table_expiration_ms: Optional[pulumi.Input[int]] = None,
                 delete_contents_on_destroy: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Dataset resource.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this table.
        :param pulumi.Input[Sequence[pulumi.Input['DatasetAccessArgs']]] accesses: An array of objects that define dataset access for one or more entities.
               Structure is documented below.
        :param pulumi.Input['DatasetDefaultEncryptionConfigurationArgs'] default_encryption_configuration: The default encryption key for all tables in the dataset. Once this property is set,
               all newly-created partitioned tables in the dataset will have encryption key set to
               this value, unless table creation request (or query) overrides the key.
               Structure is documented below.
        :param pulumi.Input[int] default_partition_expiration_ms: The default partition expiration for all partitioned tables in
               the dataset, in milliseconds.
        :param pulumi.Input[int] default_table_expiration_ms: The default lifetime of all tables in the dataset, in milliseconds.
               The minimum value is 3600000 milliseconds (one hour).
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the tables in the
               dataset when destroying the resource; otherwise,
               destroying the resource will fail if tables are present.
        :param pulumi.Input[str] description: A user-friendly description of the dataset
        :param pulumi.Input[str] friendly_name: A descriptive name for the dataset
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this dataset. You can use these to
               organize and group your datasets
        :param pulumi.Input[str] location: The geographic location where the dataset should reside.
               See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "dataset_id", dataset_id)
        if accesses is not None:
            pulumi.set(__self__, "accesses", accesses)
        if default_encryption_configuration is not None:
            pulumi.set(__self__, "default_encryption_configuration", default_encryption_configuration)
        if default_partition_expiration_ms is not None:
            pulumi.set(__self__, "default_partition_expiration_ms", default_partition_expiration_ms)
        if default_table_expiration_ms is not None:
            pulumi.set(__self__, "default_table_expiration_ms", default_table_expiration_ms)
        if delete_contents_on_destroy is not None:
            pulumi.set(__self__, "delete_contents_on_destroy", delete_contents_on_destroy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Input[str]:
        """
        The ID of the dataset containing this table.
        """
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset_id", value)

    @property
    @pulumi.getter
    def accesses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetAccessArgs']]]]:
        """
        An array of objects that define dataset access for one or more entities.
        Structure is documented below.
        """
        return pulumi.get(self, "accesses")

    @accesses.setter
    def accesses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetAccessArgs']]]]):
        pulumi.set(self, "accesses", value)

    @property
    @pulumi.getter(name="defaultEncryptionConfiguration")
    def default_encryption_configuration(self) -> Optional[pulumi.Input['DatasetDefaultEncryptionConfigurationArgs']]:
        """
        The default encryption key for all tables in the dataset. Once this property is set,
        all newly-created partitioned tables in the dataset will have encryption key set to
        this value, unless table creation request (or query) overrides the key.
        Structure is documented below.
        """
        return pulumi.get(self, "default_encryption_configuration")

    @default_encryption_configuration.setter
    def default_encryption_configuration(self, value: Optional[pulumi.Input['DatasetDefaultEncryptionConfigurationArgs']]):
        pulumi.set(self, "default_encryption_configuration", value)

    @property
    @pulumi.getter(name="defaultPartitionExpirationMs")
    def default_partition_expiration_ms(self) -> Optional[pulumi.Input[int]]:
        """
        The default partition expiration for all partitioned tables in
        the dataset, in milliseconds.
        """
        return pulumi.get(self, "default_partition_expiration_ms")

    @default_partition_expiration_ms.setter
    def default_partition_expiration_ms(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_partition_expiration_ms", value)

    @property
    @pulumi.getter(name="defaultTableExpirationMs")
    def default_table_expiration_ms(self) -> Optional[pulumi.Input[int]]:
        """
        The default lifetime of all tables in the dataset, in milliseconds.
        The minimum value is 3600000 milliseconds (one hour).
        """
        return pulumi.get(self, "default_table_expiration_ms")

    @default_table_expiration_ms.setter
    def default_table_expiration_ms(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_table_expiration_ms", value)

    @property
    @pulumi.getter(name="deleteContentsOnDestroy")
    def delete_contents_on_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, delete all the tables in the
        dataset when destroying the resource; otherwise,
        destroying the resource will fail if tables are present.
        """
        return pulumi.get(self, "delete_contents_on_destroy")

    @delete_contents_on_destroy.setter
    def delete_contents_on_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_contents_on_destroy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A user-friendly description of the dataset
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive name for the dataset
        """
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels associated with this dataset. You can use these to
        organize and group your datasets
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The geographic location where the dataset should reside.
        See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Dataset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accesses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetAccessArgs']]]]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 default_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['DatasetDefaultEncryptionConfigurationArgs']]] = None,
                 default_partition_expiration_ms: Optional[pulumi.Input[int]] = None,
                 default_table_expiration_ms: Optional[pulumi.Input[int]] = None,
                 delete_contents_on_destroy: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Dataset can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:bigquery/dataset:Dataset default projects/{{project}}/datasets/{{dataset_id}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/dataset:Dataset default {{project}}/{{dataset_id}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/dataset:Dataset default {{dataset_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetAccessArgs']]]] accesses: An array of objects that define dataset access for one or more entities.
               Structure is documented below.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this table.
        :param pulumi.Input[pulumi.InputType['DatasetDefaultEncryptionConfigurationArgs']] default_encryption_configuration: The default encryption key for all tables in the dataset. Once this property is set,
               all newly-created partitioned tables in the dataset will have encryption key set to
               this value, unless table creation request (or query) overrides the key.
               Structure is documented below.
        :param pulumi.Input[int] default_partition_expiration_ms: The default partition expiration for all partitioned tables in
               the dataset, in milliseconds.
        :param pulumi.Input[int] default_table_expiration_ms: The default lifetime of all tables in the dataset, in milliseconds.
               The minimum value is 3600000 milliseconds (one hour).
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the tables in the
               dataset when destroying the resource; otherwise,
               destroying the resource will fail if tables are present.
        :param pulumi.Input[str] description: A user-friendly description of the dataset
        :param pulumi.Input[str] friendly_name: A descriptive name for the dataset
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this dataset. You can use these to
               organize and group your datasets
        :param pulumi.Input[str] location: The geographic location where the dataset should reside.
               See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatasetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Dataset can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:bigquery/dataset:Dataset default projects/{{project}}/datasets/{{dataset_id}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/dataset:Dataset default {{project}}/{{dataset_id}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/dataset:Dataset default {{dataset_id}}
        ```

        :param str resource_name: The name of the resource.
        :param DatasetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accesses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetAccessArgs']]]]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 default_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['DatasetDefaultEncryptionConfigurationArgs']]] = None,
                 default_partition_expiration_ms: Optional[pulumi.Input[int]] = None,
                 default_table_expiration_ms: Optional[pulumi.Input[int]] = None,
                 delete_contents_on_destroy: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['accesses'] = accesses
            if dataset_id is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_id'")
            __props__['dataset_id'] = dataset_id
            __props__['default_encryption_configuration'] = default_encryption_configuration
            __props__['default_partition_expiration_ms'] = default_partition_expiration_ms
            __props__['default_table_expiration_ms'] = default_table_expiration_ms
            __props__['delete_contents_on_destroy'] = delete_contents_on_destroy
            __props__['description'] = description
            __props__['friendly_name'] = friendly_name
            __props__['labels'] = labels
            __props__['location'] = location
            __props__['project'] = project
            __props__['creation_time'] = None
            __props__['etag'] = None
            __props__['last_modified_time'] = None
            __props__['self_link'] = None
        super(Dataset, __self__).__init__(
            'gcp:bigquery/dataset:Dataset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accesses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetAccessArgs']]]]] = None,
            creation_time: Optional[pulumi.Input[int]] = None,
            dataset_id: Optional[pulumi.Input[str]] = None,
            default_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['DatasetDefaultEncryptionConfigurationArgs']]] = None,
            default_partition_expiration_ms: Optional[pulumi.Input[int]] = None,
            default_table_expiration_ms: Optional[pulumi.Input[int]] = None,
            delete_contents_on_destroy: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            last_modified_time: Optional[pulumi.Input[int]] = None,
            location: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'Dataset':
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetAccessArgs']]]] accesses: An array of objects that define dataset access for one or more entities.
               Structure is documented below.
        :param pulumi.Input[int] creation_time: The time when this dataset was created, in milliseconds since the epoch.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this table.
        :param pulumi.Input[pulumi.InputType['DatasetDefaultEncryptionConfigurationArgs']] default_encryption_configuration: The default encryption key for all tables in the dataset. Once this property is set,
               all newly-created partitioned tables in the dataset will have encryption key set to
               this value, unless table creation request (or query) overrides the key.
               Structure is documented below.
        :param pulumi.Input[int] default_partition_expiration_ms: The default partition expiration for all partitioned tables in
               the dataset, in milliseconds.
        :param pulumi.Input[int] default_table_expiration_ms: The default lifetime of all tables in the dataset, in milliseconds.
               The minimum value is 3600000 milliseconds (one hour).
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the tables in the
               dataset when destroying the resource; otherwise,
               destroying the resource will fail if tables are present.
        :param pulumi.Input[str] description: A user-friendly description of the dataset
        :param pulumi.Input[str] etag: A hash of the resource.
        :param pulumi.Input[str] friendly_name: A descriptive name for the dataset
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this dataset. You can use these to
               organize and group your datasets
        :param pulumi.Input[int] last_modified_time: The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
        :param pulumi.Input[str] location: The geographic location where the dataset should reside.
               See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accesses"] = accesses
        __props__["creation_time"] = creation_time
        __props__["dataset_id"] = dataset_id
        __props__["default_encryption_configuration"] = default_encryption_configuration
        __props__["default_partition_expiration_ms"] = default_partition_expiration_ms
        __props__["default_table_expiration_ms"] = default_table_expiration_ms
        __props__["delete_contents_on_destroy"] = delete_contents_on_destroy
        __props__["description"] = description
        __props__["etag"] = etag
        __props__["friendly_name"] = friendly_name
        __props__["labels"] = labels
        __props__["last_modified_time"] = last_modified_time
        __props__["location"] = location
        __props__["project"] = project
        __props__["self_link"] = self_link
        return Dataset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accesses(self) -> pulumi.Output[Sequence['outputs.DatasetAccess']]:
        """
        An array of objects that define dataset access for one or more entities.
        Structure is documented below.
        """
        return pulumi.get(self, "accesses")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[int]:
        """
        The time when this dataset was created, in milliseconds since the epoch.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Output[str]:
        """
        The ID of the dataset containing this table.
        """
        return pulumi.get(self, "dataset_id")

    @property
    @pulumi.getter(name="defaultEncryptionConfiguration")
    def default_encryption_configuration(self) -> pulumi.Output[Optional['outputs.DatasetDefaultEncryptionConfiguration']]:
        """
        The default encryption key for all tables in the dataset. Once this property is set,
        all newly-created partitioned tables in the dataset will have encryption key set to
        this value, unless table creation request (or query) overrides the key.
        Structure is documented below.
        """
        return pulumi.get(self, "default_encryption_configuration")

    @property
    @pulumi.getter(name="defaultPartitionExpirationMs")
    def default_partition_expiration_ms(self) -> pulumi.Output[Optional[int]]:
        """
        The default partition expiration for all partitioned tables in
        the dataset, in milliseconds.
        """
        return pulumi.get(self, "default_partition_expiration_ms")

    @property
    @pulumi.getter(name="defaultTableExpirationMs")
    def default_table_expiration_ms(self) -> pulumi.Output[Optional[int]]:
        """
        The default lifetime of all tables in the dataset, in milliseconds.
        The minimum value is 3600000 milliseconds (one hour).
        """
        return pulumi.get(self, "default_table_expiration_ms")

    @property
    @pulumi.getter(name="deleteContentsOnDestroy")
    def delete_contents_on_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, delete all the tables in the
        dataset when destroying the resource; otherwise,
        destroying the resource will fail if tables are present.
        """
        return pulumi.get(self, "delete_contents_on_destroy")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A user-friendly description of the dataset
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A hash of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[Optional[str]]:
        """
        A descriptive name for the dataset
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The labels associated with this dataset. You can use these to
        organize and group your datasets
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[int]:
        """
        The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The geographic location where the dataset should reside.
        See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

