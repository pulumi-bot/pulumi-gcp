# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Job(pulumi.CustomResource):
    """
    Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
    the official documentation for
    [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).
    
    """
    def __init__(__self__, __name__, __opts__=None, max_workers=None, name=None, on_delete=None, parameters=None, project=None, region=None, temp_gcs_location=None, template_gcs_path=None, zone=None):
        """Create a Job resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['maxWorkers'] = max_workers

        __props__['name'] = name

        __props__['onDelete'] = on_delete

        __props__['parameters'] = parameters

        __props__['project'] = project

        __props__['region'] = region

        if not temp_gcs_location:
            raise TypeError('Missing required property temp_gcs_location')
        __props__['tempGcsLocation'] = temp_gcs_location

        if not template_gcs_path:
            raise TypeError('Missing required property template_gcs_path')
        __props__['templateGcsPath'] = template_gcs_path

        __props__['zone'] = zone

        __props__['state'] = None

        super(Job, __self__).__init__(
            'gcp:dataflow/job:Job',
            __name__,
            __props__,
            __opts__)

