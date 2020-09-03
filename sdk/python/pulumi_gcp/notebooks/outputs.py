# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'EnvironmentContainerImage',
    'EnvironmentVmImage',
    'InstanceAcceleratorConfig',
    'InstanceContainerImage',
    'InstanceVmImage',
]

@pulumi.output_type
class EnvironmentContainerImage(dict):
    def __init__(__self__, *,
                 repository: str,
                 tag: Optional[str] = None):
        """
        :param str repository: The path to the container image repository.
               For example: gcr.io/{project_id}/{imageName}
        :param str tag: The tag of the container image. If not specified, this defaults to the latest tag.
        """
        pulumi.set(__self__, "repository", repository)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def repository(self) -> str:
        """
        The path to the container image repository.
        For example: gcr.io/{project_id}/{imageName}
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        The tag of the container image. If not specified, this defaults to the latest tag.
        """
        return pulumi.get(self, "tag")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EnvironmentVmImage(dict):
    def __init__(__self__, *,
                 project: str,
                 image_family: Optional[str] = None,
                 image_name: Optional[str] = None):
        """
        :param str project: The name of the Google Cloud project that this VM image belongs to.
               Format: projects/{project_id}
        :param str image_family: Use this VM image family to find the image; the newest image in this family will be used.
        :param str image_name: Use VM image name to find the image.
        """
        pulumi.set(__self__, "project", project)
        if image_family is not None:
            pulumi.set(__self__, "image_family", image_family)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The name of the Google Cloud project that this VM image belongs to.
        Format: projects/{project_id}
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="imageFamily")
    def image_family(self) -> Optional[str]:
        """
        Use this VM image family to find the image; the newest image in this family will be used.
        """
        return pulumi.get(self, "image_family")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[str]:
        """
        Use VM image name to find the image.
        """
        return pulumi.get(self, "image_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InstanceAcceleratorConfig(dict):
    def __init__(__self__, *,
                 core_count: float,
                 type: str):
        """
        :param float core_count: Count of cores of this accelerator.
        :param str type: Type of this accelerator.
               Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `TPU_V2`, and `TPU_V3`.
        """
        pulumi.set(__self__, "core_count", core_count)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="coreCount")
    def core_count(self) -> float:
        """
        Count of cores of this accelerator.
        """
        return pulumi.get(self, "core_count")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of this accelerator.
        Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `TPU_V2`, and `TPU_V3`.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InstanceContainerImage(dict):
    def __init__(__self__, *,
                 repository: str,
                 tag: Optional[str] = None):
        """
        :param str repository: The path to the container image repository.
               For example: gcr.io/{project_id}/{imageName}
        :param str tag: The tag of the container image. If not specified, this defaults to the latest tag.
        """
        pulumi.set(__self__, "repository", repository)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def repository(self) -> str:
        """
        The path to the container image repository.
        For example: gcr.io/{project_id}/{imageName}
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        The tag of the container image. If not specified, this defaults to the latest tag.
        """
        return pulumi.get(self, "tag")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InstanceVmImage(dict):
    def __init__(__self__, *,
                 project: str,
                 image_family: Optional[str] = None,
                 image_name: Optional[str] = None):
        """
        :param str project: The name of the Google Cloud project that this VM image belongs to.
               Format: projects/{project_id}
        :param str image_family: Use this VM image family to find the image; the newest image in this family will be used.
        :param str image_name: Use VM image name to find the image.
        """
        pulumi.set(__self__, "project", project)
        if image_family is not None:
            pulumi.set(__self__, "image_family", image_family)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The name of the Google Cloud project that this VM image belongs to.
        Format: projects/{project_id}
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="imageFamily")
    def image_family(self) -> Optional[str]:
        """
        Use this VM image family to find the image; the newest image in this family will be used.
        """
        return pulumi.get(self, "image_family")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[str]:
        """
        Use VM image name to find the image.
        """
        return pulumi.get(self, "image_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


