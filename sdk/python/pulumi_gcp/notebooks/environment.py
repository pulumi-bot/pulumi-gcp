# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Environment(pulumi.CustomResource):
    container_image: pulumi.Output[dict]
    """
    Use a container image to start the notebook instance.
    Structure is documented below.

      * `repository` (`str`) - The path to the container image repository.
        For example: gcr.io/{project_id}/{imageName}
      * `tag` (`str`) - The tag of the container image. If not specified, this defaults to the latest tag.
    """
    create_time: pulumi.Output[str]
    """
    Instance creation time
    """
    description: pulumi.Output[str]
    """
    A brief description of this environment.
    """
    display_name: pulumi.Output[str]
    """
    Display name of this environment for the UI.
    """
    location: pulumi.Output[str]
    """
    A reference to the zone where the machine resides.
    """
    name: pulumi.Output[str]
    """
    The name specified for the Environment instance.
    Format: projects/{project_id}/locations/{location}/environments/{environmentId}
    """
    post_startup_script: pulumi.Output[str]
    """
    Path to a Bash script that automatically runs after a notebook instance fully boots up.
    The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
    """
    project: pulumi.Output[str]
    """
    The name of the Google Cloud project that this VM image belongs to.
    Format: projects/{project_id}
    """
    vm_image: pulumi.Output[dict]
    """
    Use a Compute Engine VM image to start the notebook instance.
    Structure is documented below.

      * `imageFamily` (`str`) - Use this VM image family to find the image; the newest image in this family will be used.
      * `imageName` (`str`) - Use VM image name to find the image.
      * `project` (`str`) - The name of the Google Cloud project that this VM image belongs to.
        Format: projects/{project_id}
    """
    def __init__(__self__, resource_name, opts=None, container_image=None, description=None, display_name=None, location=None, name=None, post_startup_script=None, project=None, vm_image=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Environment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] container_image: Use a container image to start the notebook instance.
               Structure is documented below.
        :param pulumi.Input[str] description: A brief description of this environment.
        :param pulumi.Input[str] display_name: Display name of this environment for the UI.
        :param pulumi.Input[str] location: A reference to the zone where the machine resides.
        :param pulumi.Input[str] name: The name specified for the Environment instance.
               Format: projects/{project_id}/locations/{location}/environments/{environmentId}
        :param pulumi.Input[str] post_startup_script: Path to a Bash script that automatically runs after a notebook instance fully boots up.
               The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
        :param pulumi.Input[str] project: The name of the Google Cloud project that this VM image belongs to.
               Format: projects/{project_id}
        :param pulumi.Input[dict] vm_image: Use a Compute Engine VM image to start the notebook instance.
               Structure is documented below.

        The **container_image** object supports the following:

          * `repository` (`pulumi.Input[str]`) - The path to the container image repository.
            For example: gcr.io/{project_id}/{imageName}
          * `tag` (`pulumi.Input[str]`) - The tag of the container image. If not specified, this defaults to the latest tag.

        The **vm_image** object supports the following:

          * `imageFamily` (`pulumi.Input[str]`) - Use this VM image family to find the image; the newest image in this family will be used.
          * `imageName` (`pulumi.Input[str]`) - Use VM image name to find the image.
          * `project` (`pulumi.Input[str]`) - The name of the Google Cloud project that this VM image belongs to.
            Format: projects/{project_id}
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

            __props__['container_image'] = container_image
            __props__['description'] = description
            __props__['display_name'] = display_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['name'] = name
            __props__['post_startup_script'] = post_startup_script
            __props__['project'] = project
            __props__['vm_image'] = vm_image
            __props__['create_time'] = None
        super(Environment, __self__).__init__(
            'gcp:notebooks/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, container_image=None, create_time=None, description=None, display_name=None, location=None, name=None, post_startup_script=None, project=None, vm_image=None):
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] container_image: Use a container image to start the notebook instance.
               Structure is documented below.
        :param pulumi.Input[str] create_time: Instance creation time
        :param pulumi.Input[str] description: A brief description of this environment.
        :param pulumi.Input[str] display_name: Display name of this environment for the UI.
        :param pulumi.Input[str] location: A reference to the zone where the machine resides.
        :param pulumi.Input[str] name: The name specified for the Environment instance.
               Format: projects/{project_id}/locations/{location}/environments/{environmentId}
        :param pulumi.Input[str] post_startup_script: Path to a Bash script that automatically runs after a notebook instance fully boots up.
               The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
        :param pulumi.Input[str] project: The name of the Google Cloud project that this VM image belongs to.
               Format: projects/{project_id}
        :param pulumi.Input[dict] vm_image: Use a Compute Engine VM image to start the notebook instance.
               Structure is documented below.

        The **container_image** object supports the following:

          * `repository` (`pulumi.Input[str]`) - The path to the container image repository.
            For example: gcr.io/{project_id}/{imageName}
          * `tag` (`pulumi.Input[str]`) - The tag of the container image. If not specified, this defaults to the latest tag.

        The **vm_image** object supports the following:

          * `imageFamily` (`pulumi.Input[str]`) - Use this VM image family to find the image; the newest image in this family will be used.
          * `imageName` (`pulumi.Input[str]`) - Use VM image name to find the image.
          * `project` (`pulumi.Input[str]`) - The name of the Google Cloud project that this VM image belongs to.
            Format: projects/{project_id}
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["container_image"] = container_image
        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["location"] = location
        __props__["name"] = name
        __props__["post_startup_script"] = post_startup_script
        __props__["project"] = project
        __props__["vm_image"] = vm_image
        return Environment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
