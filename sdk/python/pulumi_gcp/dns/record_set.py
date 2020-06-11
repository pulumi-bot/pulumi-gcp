# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
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
    whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g. `"first255characters\"\"morecharacters"`).
    """
    ttl: pulumi.Output[float]
    """
    The time-to-live of this record set (seconds).
    """
    type: pulumi.Output[str]
    """
    The DNS record set type.
    """
    def __init__(__self__, resource_name, opts=None, managed_zone=None, name=None, project=None, rrdatas=None, ttl=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a set of DNS records within Google Cloud DNS. For more information see [the official documentation](https://cloud.google.com/dns/records/) and
        [API](https://cloud.google.com/dns/api/v1/resourceRecordSets).

        > **Note:** The provider treats this resource as an authoritative record set. This means existing records (including 
        the default records) for the given type will be overwritten when you create this resource in the provider. 
        In addition, the Google Cloud DNS API requires NS records to be present at all times, so the provider 
        will not actually remove NS records during destroy but will report that it did.

        ## Example Usage

        ### Binding a DNS name to the ephemeral IP of a new instance:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        frontend_instance = gcp.compute.Instance("frontendInstance",
            machine_type="g1-small",
            zone="us-central1-b",
            boot_disk={
                "initialize_params": {
                    "image": "debian-cloud/debian-9",
                },
            },
            network_interface=[{
                "network": "default",
                "access_config": [{}],
            }])
        prod = gcp.dns.ManagedZone("prod", dns_name="prod.mydomain.com.")
        frontend_record_set = gcp.dns.RecordSet("frontendRecordSet",
            type="A",
            ttl=300,
            managed_zone=prod.name,
            rrdatas=[frontend_instance.network_interfaces[0]["accessConfigs"][0]["natIp"]])
        ```

        ### Adding an A record

        ```python
        import pulumi
        import pulumi_gcp as gcp

        prod = gcp.dns.ManagedZone("prod", dns_name="prod.mydomain.com.")
        record_set = gcp.dns.RecordSet("recordSet",
            managed_zone=prod.name,
            type="A",
            ttl=300,
            rrdatas=["8.8.8.8"])
        ```

        ### Adding an MX record

        ```python
        import pulumi
        import pulumi_gcp as gcp

        prod = gcp.dns.ManagedZone("prod", dns_name="prod.mydomain.com.")
        mx = gcp.dns.RecordSet("mx",
            managed_zone=prod.name,
            type="MX",
            ttl=3600,
            rrdatas=[
                "1 aspmx.l.google.com.",
                "5 alt1.aspmx.l.google.com.",
                "5 alt2.aspmx.l.google.com.",
                "10 alt3.aspmx.l.google.com.",
                "10 alt4.aspmx.l.google.com.",
            ])
        ```

        ### Adding an SPF record

        ```python
        import pulumi
        import pulumi_gcp as gcp

        prod = gcp.dns.ManagedZone("prod", dns_name="prod.mydomain.com.")
        spf = gcp.dns.RecordSet("spf",
            managed_zone=prod.name,
            type="TXT",
            ttl=300,
            rrdatas=["\"v=spf1 ip4:111.111.111.111 include:backoff.email-example.com -all\""])
        ```

        ### Adding a CNAME record

        ```python
        import pulumi
        import pulumi_gcp as gcp

        prod = gcp.dns.ManagedZone("prod", dns_name="prod.mydomain.com.")
        cname = gcp.dns.RecordSet("cname",
            managed_zone=prod.name,
            type="CNAME",
            ttl=300,
            rrdatas=["frontend.mydomain.com."])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] managed_zone: The name of the zone in which this record set will
               reside.
        :param pulumi.Input[str] name: The DNS name this record set will apply to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[list] rrdatas: The string data for the records in this record set
               whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g. `"first255characters\"\"morecharacters"`).
        :param pulumi.Input[float] ttl: The time-to-live of this record set (seconds).
        :param pulumi.Input[str] type: The DNS record set type.
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
        super(RecordSet, __self__).__init__(
            'gcp:dns/recordSet:RecordSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, managed_zone=None, name=None, project=None, rrdatas=None, ttl=None, type=None):
        """
        Get an existing RecordSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] managed_zone: The name of the zone in which this record set will
               reside.
        :param pulumi.Input[str] name: The DNS name this record set will apply to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[list] rrdatas: The string data for the records in this record set
               whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g. `"first255characters\"\"morecharacters"`).
        :param pulumi.Input[float] ttl: The time-to-live of this record set (seconds).
        :param pulumi.Input[str] type: The DNS record set type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["managed_zone"] = managed_zone
        __props__["name"] = name
        __props__["project"] = project
        __props__["rrdatas"] = rrdatas
        __props__["ttl"] = ttl
        __props__["type"] = type
        return RecordSet(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

