# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ObjectACL(pulumi.CustomResource):
    bucket: pulumi.Output[str]
    """
    The name of the bucket the object is stored in.
    """
    object: pulumi.Output[str]
    """
    The name of the object to apply the acl to.
    """
    predefined_acl: pulumi.Output[str]
    """
    The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
    """
    role_entities: pulumi.Output[list]
    """
    List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
    Must be set if `predefined_acl` is not.
    """
    def __init__(__self__, resource_name, opts=None, bucket=None, object=None, predefined_acl=None, role_entities=None, __name__=None, __opts__=None):
        """
        Authoritatively manages the access control list (ACL) for an object in a Google
        Cloud Storage (GCS) bucket. Removing a `google_storage_object_acl` sets the
        acl to the `private` [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl).
        
        For more information see
        [the official documentation](https://cloud.google.com/storage/docs/access-control/lists) 
        and 
        [API](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls).
        
        > Want fine-grained control over object ACLs? Use `google_storage_object_access_control` to control individual
        role entity pairs.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket the object is stored in.
        :param pulumi.Input[str] object: The name of the object to apply the acl to.
        :param pulumi.Input[str] predefined_acl: The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
        :param pulumi.Input[list] role_entities: List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
               Must be set if `predefined_acl` is not.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_object_acl.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__

        __props__ = dict()

        if bucket is None:
            raise TypeError("Missing required property 'bucket'")
        __props__['bucket'] = bucket
        if object is None:
            raise TypeError("Missing required property 'object'")
        __props__['object'] = object
        __props__['predefined_acl'] = predefined_acl
        __props__['role_entities'] = role_entities
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ObjectACL, __self__).__init__(
            'gcp:storage/objectACL:ObjectACL',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

