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

__all__ = ['DicomStore']


class DicomStore(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
        (https://www.dicomstandard.org/about/) standard for Healthcare information exchange

        To get more information about DicomStore, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.dicomStores)
        * How-to Guides
            * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize DICOM stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the DicomStore.
               ** Changing this property may recreate the Dicom store (removing all data) **
        :param pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']] notification_config: A nested object resource
               Structure is documented below.
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

            if dataset is None:
                raise TypeError("Missing required property 'dataset'")
            __props__['dataset'] = dataset
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['notification_config'] = notification_config
            __props__['self_link'] = None
        super(DicomStore, __self__).__init__(
            'gcp:healthcare/dicomStore:DicomStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dataset: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_config: Optional[pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'DicomStore':
        """
        Get an existing DicomStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize DICOM stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the DicomStore.
               ** Changing this property may recreate the Dicom store (removing all data) **
        :param pulumi.Input[pulumi.InputType['DicomStoreNotificationConfigArgs']] notification_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dataset"] = dataset
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["notification_config"] = notification_config
        __props__["self_link"] = self_link
        return DicomStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Output[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        """
        return pulumi.get(self, "dataset")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        User-supplied key-value pairs used to organize DICOM stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the DicomStore.
        ** Changing this property may recreate the Dicom store (removing all data) **
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> pulumi.Output[Optional['outputs.DicomStoreNotificationConfig']]:
        """
        A nested object resource
        Structure is documented below.
        """
        return pulumi.get(self, "notification_config")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The fully qualified name of this dataset
        """
        return pulumi.get(self, "self_link")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

