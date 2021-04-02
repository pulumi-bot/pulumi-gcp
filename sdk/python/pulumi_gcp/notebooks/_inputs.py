# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'EnvironmentContainerImageArgs',
    'EnvironmentVmImageArgs',
    'InstanceAcceleratorConfigArgs',
    'InstanceContainerImageArgs',
    'InstanceIamBindingConditionArgs',
    'InstanceIamMemberConditionArgs',
    'InstanceShieldedInstanceConfigArgs',
    'InstanceVmImageArgs',
]

@pulumi.input_type
class EnvironmentContainerImageArgs:
    def __init__(__self__, *,
                 repository: pulumi.Input[str],
                 tag: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] repository: The path to the container image repository.
               For example: gcr.io/{project_id}/{imageName}
        :param pulumi.Input[str] tag: The tag of the container image. If not specified, this defaults to the latest tag.
        """
        pulumi.set(__self__, "repository", repository)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The path to the container image repository.
        For example: gcr.io/{project_id}/{imageName}
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        The tag of the container image. If not specified, this defaults to the latest tag.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)


@pulumi.input_type
class EnvironmentVmImageArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 image_family: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] project: The name of the Google Cloud project that this VM image belongs to.
               Format: projects/{project_id}
        :param pulumi.Input[str] image_family: Use this VM image family to find the image; the newest image in this family will be used.
        :param pulumi.Input[str] image_name: Use VM image name to find the image.
        """
        pulumi.set(__self__, "project", project)
        if image_family is not None:
            pulumi.set(__self__, "image_family", image_family)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The name of the Google Cloud project that this VM image belongs to.
        Format: projects/{project_id}
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="imageFamily")
    def image_family(self) -> Optional[pulumi.Input[str]]:
        """
        Use this VM image family to find the image; the newest image in this family will be used.
        """
        return pulumi.get(self, "image_family")

    @image_family.setter
    def image_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_family", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[pulumi.Input[str]]:
        """
        Use VM image name to find the image.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_name", value)


@pulumi.input_type
class InstanceAcceleratorConfigArgs:
    def __init__(__self__, *,
                 core_count: pulumi.Input[int],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[int] core_count: Count of cores of this accelerator.
        :param pulumi.Input[str] type: Type of this accelerator.
               Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, and `TPU_V3`.
        """
        pulumi.set(__self__, "core_count", core_count)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="coreCount")
    def core_count(self) -> pulumi.Input[int]:
        """
        Count of cores of this accelerator.
        """
        return pulumi.get(self, "core_count")

    @core_count.setter
    def core_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "core_count", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of this accelerator.
        Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, and `TPU_V3`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class InstanceContainerImageArgs:
    def __init__(__self__, *,
                 repository: pulumi.Input[str],
                 tag: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] repository: The path to the container image repository.
               For example: gcr.io/{project_id}/{imageName}
        :param pulumi.Input[str] tag: The tag of the container image. If not specified, this defaults to the latest tag.
        """
        pulumi.set(__self__, "repository", repository)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The path to the container image repository.
        For example: gcr.io/{project_id}/{imageName}
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        The tag of the container image. If not specified, this defaults to the latest tag.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)


@pulumi.input_type
class InstanceIamBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class InstanceIamMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class InstanceShieldedInstanceConfigArgs:
    def __init__(__self__, *,
                 enable_integrity_monitoring: Optional[pulumi.Input[bool]] = None,
                 enable_secure_boot: Optional[pulumi.Input[bool]] = None,
                 enable_vtpm: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] enable_integrity_monitoring: Defines whether the instance has integrity monitoring enabled. Enables monitoring and attestation of the
               boot integrity of the instance. The attestation is performed against the integrity policy baseline.
               This baseline is initially derived from the implicitly trusted boot image when the instance is created.
               Enabled by default.
        :param pulumi.Input[bool] enable_secure_boot: Defines whether the instance has Secure Boot enabled. Secure Boot helps ensure that the system only runs
               authentic software by verifying the digital signature of all boot components, and halting the boot process
               if signature verification fails.
               Disabled by default.
        :param pulumi.Input[bool] enable_vtpm: Defines whether the instance has the vTPM enabled.
               Enabled by default.
        """
        if enable_integrity_monitoring is not None:
            pulumi.set(__self__, "enable_integrity_monitoring", enable_integrity_monitoring)
        if enable_secure_boot is not None:
            pulumi.set(__self__, "enable_secure_boot", enable_secure_boot)
        if enable_vtpm is not None:
            pulumi.set(__self__, "enable_vtpm", enable_vtpm)

    @property
    @pulumi.getter(name="enableIntegrityMonitoring")
    def enable_integrity_monitoring(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines whether the instance has integrity monitoring enabled. Enables monitoring and attestation of the
        boot integrity of the instance. The attestation is performed against the integrity policy baseline.
        This baseline is initially derived from the implicitly trusted boot image when the instance is created.
        Enabled by default.
        """
        return pulumi.get(self, "enable_integrity_monitoring")

    @enable_integrity_monitoring.setter
    def enable_integrity_monitoring(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_integrity_monitoring", value)

    @property
    @pulumi.getter(name="enableSecureBoot")
    def enable_secure_boot(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines whether the instance has Secure Boot enabled. Secure Boot helps ensure that the system only runs
        authentic software by verifying the digital signature of all boot components, and halting the boot process
        if signature verification fails.
        Disabled by default.
        """
        return pulumi.get(self, "enable_secure_boot")

    @enable_secure_boot.setter
    def enable_secure_boot(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_secure_boot", value)

    @property
    @pulumi.getter(name="enableVtpm")
    def enable_vtpm(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines whether the instance has the vTPM enabled.
        Enabled by default.
        """
        return pulumi.get(self, "enable_vtpm")

    @enable_vtpm.setter
    def enable_vtpm(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_vtpm", value)


@pulumi.input_type
class InstanceVmImageArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 image_family: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] project: The name of the Google Cloud project that this VM image belongs to.
               Format: projects/{project_id}
        :param pulumi.Input[str] image_family: Use this VM image family to find the image; the newest image in this family will be used.
        :param pulumi.Input[str] image_name: Use VM image name to find the image.
        """
        pulumi.set(__self__, "project", project)
        if image_family is not None:
            pulumi.set(__self__, "image_family", image_family)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The name of the Google Cloud project that this VM image belongs to.
        Format: projects/{project_id}
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="imageFamily")
    def image_family(self) -> Optional[pulumi.Input[str]]:
        """
        Use this VM image family to find the image; the newest image in this family will be used.
        """
        return pulumi.get(self, "image_family")

    @image_family.setter
    def image_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_family", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[pulumi.Input[str]]:
        """
        Use VM image name to find the image.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_name", value)


