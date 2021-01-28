// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority.Outputs
{

    [OutputType]
    public sealed class AuthorityConfigSubjectConfig
    {
        /// <summary>
        /// The common name of the distinguished name.
        /// </summary>
        public readonly string CommonName;
        /// <summary>
        /// Contains distinguished name fields such as the location and organization.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.AuthorityConfigSubjectConfigSubject Subject;
        /// <summary>
        /// The subject alternative name fields.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.AuthorityConfigSubjectConfigSubjectAltName? SubjectAltName;

        [OutputConstructor]
        private AuthorityConfigSubjectConfig(
            string commonName,

            Outputs.AuthorityConfigSubjectConfigSubject subject,

            Outputs.AuthorityConfigSubjectConfigSubjectAltName? subjectAltName)
        {
            CommonName = commonName;
            Subject = subject;
            SubjectAltName = subjectAltName;
        }
    }
}