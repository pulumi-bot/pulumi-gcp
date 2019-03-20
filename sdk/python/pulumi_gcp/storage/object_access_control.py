# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ObjectAccessControl(pulumi.CustomResource):
    bucket: pulumi.Output[str]
    domain: pulumi.Output[str]
    email: pulumi.Output[str]
    entity: pulumi.Output[str]
    entity_id: pulumi.Output[str]
    generation: pulumi.Output[float]
    object: pulumi.Output[str]
    project_team: pulumi.Output[dict]
    role: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, bucket=None, entity=None, object=None, role=None, __name__=None, __opts__=None):
        """
        The ObjectAccessControls resources represent the Access Control Lists
        (ACLs) for objects within Google Cloud Storage. ACLs let you specify
        who has access to your data and to what extent.
        
        There are two roles that can be assigned to an entity:
        
        READERs can get an object, though the acl property will not be revealed.
        OWNERs are READERs, and they can get the acl property, update an object,
        and call all objectAccessControls methods on the object. The owner of an
        object is always an OWNER.
        For more information, see Access Control, with the caveat that this API
        uses READER and OWNER instead of READ and FULL_CONTROL.
        
        
        To get more information about ObjectAccessControl, see:
        
        * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=storage_object_access_control_public_object&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if bucket is None:
            raise TypeError('Missing required property bucket')
        __props__['bucket'] = bucket

        if entity is None:
            raise TypeError('Missing required property entity')
        __props__['entity'] = entity

        if object is None:
            raise TypeError('Missing required property object')
        __props__['object'] = object

        if role is None:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        __props__['domain'] = None
        __props__['email'] = None
        __props__['entity_id'] = None
        __props__['generation'] = None
        __props__['project_team'] = None

        super(ObjectAccessControl, __self__).__init__(
            'gcp:storage/objectAccessControl:ObjectAccessControl',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

