# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ObjectACL']


class ObjectACL(pulumi.CustomResource):
    bucket: pulumi.Output[str] = pulumi.property("bucket")
    """
    The name of the bucket the object is stored in.
    """

    object: pulumi.Output[str] = pulumi.property("object")
    """
    The name of the object to apply the acl to.
    """

    predefined_acl: pulumi.Output[Optional[str]] = pulumi.property("predefinedAcl")
    """
    The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
    """

    role_entities: pulumi.Output[List[str]] = pulumi.property("roleEntities")
    """
    List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
    Must be set if `predefined_acl` is not.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 predefined_acl: Optional[pulumi.Input[str]] = None,
                 role_entities: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Authoritatively manages the access control list (ACL) for an object in a Google
        Cloud Storage (GCS) bucket. Removing a `storage.ObjectACL` sets the
        acl to the `private` [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl).

        For more information see
        [the official documentation](https://cloud.google.com/storage/docs/access-control/lists)
        and
        [API](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls).

        > Want fine-grained control over object ACLs? Use `storage.ObjectAccessControl` to control individual
        role entity pairs.

        ## Example Usage

        Create an object ACL with one owner and one reader.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_store = gcp.storage.Bucket("image-store", location="EU")
        image = gcp.storage.BucketObject("image",
            bucket=image_store.name,
            source=pulumi.FileAsset("image1.jpg"))
        image_store_acl = gcp.storage.ObjectACL("image-store-acl",
            bucket=image_store.name,
            object=image.output_name,
            role_entities=[
                "OWNER:user-my.email@gmail.com",
                "READER:group-mygroup",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket the object is stored in.
        :param pulumi.Input[str] object: The name of the object to apply the acl to.
        :param pulumi.Input[str] predefined_acl: The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
        :param pulumi.Input[List[pulumi.Input[str]]] role_entities: List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
               Must be set if `predefined_acl` is not.
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

            if bucket is None:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            if object is None:
                raise TypeError("Missing required property 'object'")
            __props__['object'] = object
            __props__['predefined_acl'] = predefined_acl
            __props__['role_entities'] = role_entities
        super(ObjectACL, __self__).__init__(
            'gcp:storage/objectACL:ObjectACL',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            object: Optional[pulumi.Input[str]] = None,
            predefined_acl: Optional[pulumi.Input[str]] = None,
            role_entities: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'ObjectACL':
        """
        Get an existing ObjectACL resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket the object is stored in.
        :param pulumi.Input[str] object: The name of the object to apply the acl to.
        :param pulumi.Input[str] predefined_acl: The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
        :param pulumi.Input[List[pulumi.Input[str]]] role_entities: List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
               Must be set if `predefined_acl` is not.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket"] = bucket
        __props__["object"] = object
        __props__["predefined_acl"] = predefined_acl
        __props__["role_entities"] = role_entities
        return ObjectACL(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

