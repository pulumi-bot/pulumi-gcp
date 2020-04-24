// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class DiskSourceSnapshotEncryptionKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
        /// in the cloud console. Your project's Compute Engine System service account
        /// (`service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com`) must have
        /// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
        /// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
        /// </summary>
        [Input("kmsKeySelfLink")]
        public Input<string>? KmsKeySelfLink { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a 256-bit customer-supplied encryption key, encoded in
        /// RFC 4648 base64 to either encrypt or decrypt this resource.
        /// </summary>
        [Input("rawKey")]
        public Input<string>? RawKey { get; set; }

        /// <summary>
        /// -
        /// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
        /// encryption key that protects this resource.
        /// </summary>
        [Input("sha256")]
        public Input<string>? Sha256 { get; set; }

        public DiskSourceSnapshotEncryptionKeyArgs()
        {
        }
    }
}
