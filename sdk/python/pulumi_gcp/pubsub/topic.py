# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Topic(pulumi.CustomResource):
    kms_key_name: pulumi.Output[str]
    """
    The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. Your
    project's PubSub service account ('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have
    'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature. The expected format is
    'projects/*/locations/*/keyRings/*/cryptoKeys/*'
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to this Topic.
    """
    message_storage_policy: pulumi.Output[dict]
    """
    Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not
    present, then no constraints are in effect.

      * `allowedPersistenceRegions` (`list`)
    """
    name: pulumi.Output[str]
    """
    Name of the topic.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, kms_key_name=None, labels=None, message_storage_policy=None, name=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Topic resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kms_key_name: The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. Your
               project's PubSub service account ('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have
               'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature. The expected format is
               'projects/*/locations/*/keyRings/*/cryptoKeys/*'
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to this Topic.
        :param pulumi.Input[dict] message_storage_policy: Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not
               present, then no constraints are in effect.
        :param pulumi.Input[str] name: Name of the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **message_storage_policy** object supports the following:

          * `allowedPersistenceRegions` (`pulumi.Input[list]`)
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
    def get(resource_name, id, opts=None, kms_key_name=None, labels=None, message_storage_policy=None, name=None, project=None):
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kms_key_name: The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. Your
               project's PubSub service account ('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have
               'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature. The expected format is
               'projects/*/locations/*/keyRings/*/cryptoKeys/*'
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to this Topic.
        :param pulumi.Input[dict] message_storage_policy: Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not
               present, then no constraints are in effect.
        :param pulumi.Input[str] name: Name of the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **message_storage_policy** object supports the following:

          * `allowedPersistenceRegions` (`pulumi.Input[list]`)
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

