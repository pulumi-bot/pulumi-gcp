// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class JobSparkConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("archiveUris")]
        private InputList<string>? _archiveUris;

        /// <summary>
        /// HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
        /// </summary>
        public InputList<string> ArchiveUris
        {
            get => _archiveUris ?? (_archiveUris = new InputList<string>());
            set => _archiveUris = value;
        }

        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// The arguments to pass to the driver.
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("fileUris")]
        private InputList<string>? _fileUris;

        /// <summary>
        /// HCFS URIs of files to be copied to the working directory of Python drivers and distributed tasks. Useful for naively parallel tasks.
        /// </summary>
        public InputList<string> FileUris
        {
            get => _fileUris ?? (_fileUris = new InputList<string>());
            set => _fileUris = value;
        }

        [Input("jarFileUris")]
        private InputList<string>? _jarFileUris;

        /// <summary>
        /// HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.
        /// </summary>
        public InputList<string> JarFileUris
        {
            get => _jarFileUris ?? (_jarFileUris = new InputList<string>());
            set => _jarFileUris = value;
        }

        [Input("loggingConfig")]
        public Input<Inputs.JobSparkConfigLoggingConfigGetArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// The class containing the main method of the driver. Must be in a
        /// provided jar or jar that is already on the classpath. Conflicts with `main_jar_file_uri`
        /// </summary>
        [Input("mainClass")]
        public Input<string>? MainClass { get; set; }

        /// <summary>
        /// The HCFS URI of jar file containing
        /// the driver jar. Conflicts with `main_class`
        /// </summary>
        [Input("mainJarFileUri")]
        public Input<string>? MainJarFileUri { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in `/etc/spark/conf/spark-defaults.conf` and classes in user code.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        public JobSparkConfigGetArgs()
        {
        }
    }
}
