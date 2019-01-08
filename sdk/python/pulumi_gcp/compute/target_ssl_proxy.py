# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class TargetSSLProxy(pulumi.CustomResource):
    """
    Represents a TargetSslProxy resource, which is used by one or more
    global forwarding rule to route incoming SSL requests to a backend
    service.
    
    
    To get more information about TargetSslProxy, see:
    
    * [API documentation](https://cloud.google.com/compute/docs/reference/latest/targetSslProxies)
    * How-to Guides
        * [Setting Up SSL proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/)
    """
    def __init__(__self__, __name__, __opts__=None, backend_service=None, description=None, name=None, project=None, proxy_header=None, ssl_certificates=None, ssl_policy=None):
        """Create a TargetSSLProxy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not backend_service:
            raise TypeError('Missing required property backend_service')
        __props__['backend_service'] = backend_service

        __props__['description'] = description

        __props__['name'] = name

        __props__['project'] = project

        __props__['proxy_header'] = proxy_header

        if not ssl_certificates:
            raise TypeError('Missing required property ssl_certificates')
        __props__['ssl_certificates'] = ssl_certificates

        __props__['ssl_policy'] = ssl_policy

        __props__['creation_timestamp'] = None
        __props__['proxy_id'] = None
        __props__['self_link'] = None

        super(TargetSSLProxy, __self__).__init__(
            'gcp:compute/targetSSLProxy:TargetSSLProxy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

