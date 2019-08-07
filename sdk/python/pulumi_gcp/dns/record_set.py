# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RecordSet(pulumi.CustomResource):
    managed_zone: pulumi.Output[str]
    """
    The name of the zone in which this record set will
    reside.
    """
    name: pulumi.Output[str]
    """
    The DNS name this record set will apply to.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    rrdatas: pulumi.Output[list]
    """
    The string data for the records in this record set
    whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside this provider's configuration string (e.g. `"first255characters\"\"morecharacters"`).
    """
    ttl: pulumi.Output[float]
    """
    The time-to-live of this record set (seconds).
    """
    type: pulumi.Output[str]
    """
    The DNS record set type.
    """
    def __init__(__self__, resource_name, opts=None, managed_zone=None, name=None, project=None, rrdatas=None, ttl=None, type=None, __name__=None, __opts__=None):
        """
        Manages a set of DNS records within Google Cloud DNS. For more information see [the official documentation](https://cloud.google.com/dns/records/) and
        [API](https://cloud.google.com/dns/api/v1/resourceRecordSets).
        
        > **Note:** The provider treats this resource as an authoritative record set. This means existing records (including the default records) for the given type will be overwritten when you create this resource with this provider. In addition, the Google Cloud DNS API requires NS records to be present at all times, so this provider will not actually remove NS records during destroy but will report that it did.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] managed_zone: The name of the zone in which this record set will
               reside.
        :param pulumi.Input[str] name: The DNS name this record set will apply to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[list] rrdatas: The string data for the records in this record set
               whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside this provider's configuration string (e.g. `"first255characters\"\"morecharacters"`).
        :param pulumi.Input[float] ttl: The time-to-live of this record set (seconds).
        :param pulumi.Input[str] type: The DNS record set type.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_record_set.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if managed_zone is None:
            raise TypeError("Missing required property 'managed_zone'")
        __props__['managed_zone'] = managed_zone
        __props__['name'] = name
        __props__['project'] = project
        if rrdatas is None:
            raise TypeError("Missing required property 'rrdatas'")
        __props__['rrdatas'] = rrdatas
        if ttl is None:
            raise TypeError("Missing required property 'ttl'")
        __props__['ttl'] = ttl
        if type is None:
            raise TypeError("Missing required property 'type'")
        __props__['type'] = type
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(RecordSet, __self__).__init__(
            'gcp:dns/recordSet:RecordSet',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

