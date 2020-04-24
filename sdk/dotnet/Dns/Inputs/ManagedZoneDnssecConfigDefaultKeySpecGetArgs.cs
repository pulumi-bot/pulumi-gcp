// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns.Inputs
{

    public sealed class ManagedZoneDnssecConfigDefaultKeySpecGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String mnemonic specifying the DNSSEC algorithm of this key
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Length of the keys in bits
        /// </summary>
        [Input("keyLength")]
        public Input<int>? KeyLength { get; set; }

        /// <summary>
        /// Specifies whether this is a key signing key (KSK) or a zone
        /// signing key (ZSK). Key signing keys have the Secure Entry
        /// Point flag set and, when active, will only be used to sign
        /// resource record sets of type DNSKEY. Zone signing keys do
        /// not have the Secure Entry Point flag set and will be used
        /// to sign all other types of resource record sets.
        /// </summary>
        [Input("keyType")]
        public Input<string>? KeyType { get; set; }

        /// <summary>
        /// Identifies what kind of resource this is
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        public ManagedZoneDnssecConfigDefaultKeySpecGetArgs()
        {
        }
    }
}
