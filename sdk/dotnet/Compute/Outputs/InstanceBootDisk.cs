// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class InstanceBootDisk
    {
        /// <summary>
        /// Whether the disk will be auto-deleted when the instance
        /// is deleted. Defaults to true.
        /// </summary>
        public readonly bool? AutoDelete;
        /// <summary>
        /// Name with which attached disk will be accessible.
        /// On the instance, this device will be `/dev/disk/by-id/google-{{device_name}}`.
        /// </summary>
        public readonly string? DeviceName;
        /// <summary>
        /// A 256-bit [customer-supplied encryption key]
        /// (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
        /// encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
        /// to encrypt this disk. Only one of `kms_key_self_link` and `disk_encryption_key_raw`
        /// may be set.
        /// </summary>
        public readonly string? DiskEncryptionKeyRaw;
        public readonly string? DiskEncryptionKeySha256;
        /// <summary>
        /// Parameters for a new disk that will be created
        /// alongside the new instance. Either `initialize_params` or `source` must be set.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.InstanceBootDiskInitializeParams? InitializeParams;
        /// <summary>
        /// The self_link of the encryption key that is
        /// stored in Google Cloud KMS to encrypt this disk. Only one of `kms_key_self_link`
        /// and `disk_encryption_key_raw` may be set.
        /// </summary>
        public readonly string? KmsKeySelfLink;
        /// <summary>
        /// The mode in which to attach this disk, either `READ_WRITE`
        /// or `READ_ONLY`. If not specified, the default is to attach the disk in `READ_WRITE` mode.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// The name or self_link of the existing disk (such as those managed by
        /// `gcp.compute.Disk`) or disk image. To create an instance from a snapshot, first create a
        /// `gcp.compute.Disk` from a snapshot and reference it here.
        /// </summary>
        public readonly string? Source;

        [OutputConstructor]
        private InstanceBootDisk(
            bool? autoDelete,

            string? deviceName,

            string? diskEncryptionKeyRaw,

            string? diskEncryptionKeySha256,

            Outputs.InstanceBootDiskInitializeParams? initializeParams,

            string? kmsKeySelfLink,

            string? mode,

            string? source)
        {
            AutoDelete = autoDelete;
            DeviceName = deviceName;
            DiskEncryptionKeyRaw = diskEncryptionKeyRaw;
            DiskEncryptionKeySha256 = diskEncryptionKeySha256;
            InitializeParams = initializeParams;
            KmsKeySelfLink = kmsKeySelfLink;
            Mode = mode;
            Source = source;
        }
    }
}
