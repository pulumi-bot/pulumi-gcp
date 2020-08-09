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

__all__ = ['Topic']


class Topic(pulumi.CustomResource):
    kms_key_name: pulumi.Output[Optional[str]] = pulumi.property("kmsKeyName")
    """
    The resource name of the Cloud KMS CryptoKey to be used to protect access
    to messages published on this topic. Your project's PubSub service account
    (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
    `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
    The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
    """

    labels: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("labels")
    """
    A set of key/value label pairs to assign to this Topic.
    """

    message_storage_policy: pulumi.Output['outputs.TopicMessageStoragePolicy'] = pulumi.property("messageStoragePolicy")
    """
    Policy constraining the set of Google Cloud Platform regions where
    messages published to the topic may be stored. If not present, then no
    constraints are in effect.  Structure is documented below.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Name of the topic.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 message_storage_policy: Optional[pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A named resource to which messages are sent by publishers.

        To get more information about Topic, see:

        * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
        * How-to Guides
            * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)

        ## Example Usage
        ### Pubsub Topic Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.pubsub.Topic("example", labels={
            "foo": "bar",
        })
        ```
        ### Pubsub Topic Cmek

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRing("keyRing", location="global")
        crypto_key = gcp.kms.CryptoKey("cryptoKey", key_ring=key_ring.id)
        example = gcp.pubsub.Topic("example", kms_key_name=crypto_key.id)
        ```
        ### Pubsub Topic Geo Restricted

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.pubsub.Topic("example", message_storage_policy={
            "allowedPersistenceRegions": ["europe-west3"],
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kms_key_name: The resource name of the Cloud KMS CryptoKey to be used to protect access
               to messages published on this topic. Your project's PubSub service account
               (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
               `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
               The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this Topic.
        :param pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']] message_storage_policy: Policy constraining the set of Google Cloud Platform regions where
               messages published to the topic may be stored. If not present, then no
               constraints are in effect.  Structure is documented below.
        :param pulumi.Input[str] name: Name of the topic.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['kms_key_name'] = kms_key_name
            __props__['labels'] = labels
            __props__['message_storage_policy'] = message_storage_policy
            __props__['name'] = name
            __props__['project'] = project
        super(Topic, __self__).__init__(
            'gcp:pubsub/topic:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            kms_key_name: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            message_storage_policy: Optional[pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'Topic':
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kms_key_name: The resource name of the Cloud KMS CryptoKey to be used to protect access
               to messages published on this topic. Your project's PubSub service account
               (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
               `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
               The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this Topic.
        :param pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']] message_storage_policy: Policy constraining the set of Google Cloud Platform regions where
               messages published to the topic may be stored. If not present, then no
               constraints are in effect.  Structure is documented below.
        :param pulumi.Input[str] name: Name of the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["kms_key_name"] = kms_key_name
        __props__["labels"] = labels
        __props__["message_storage_policy"] = message_storage_policy
        __props__["name"] = name
        __props__["project"] = project
        return Topic(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

