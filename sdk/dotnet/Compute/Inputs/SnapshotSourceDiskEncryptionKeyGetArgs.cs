// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SnapshotSourceDiskEncryptionKeyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a 256-bit customer-supplied encryption key, encoded in
        /// RFC 4648 base64 to either encrypt or decrypt this resource.
        /// </summary>
        [Input("rawKey")]
        public Input<string>? RawKey { get; set; }

        public SnapshotSourceDiskEncryptionKeyGetArgs()
        {
        }
    }
}
