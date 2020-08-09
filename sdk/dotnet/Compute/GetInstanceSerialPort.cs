// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetInstanceSerialPort
    {
        /// <summary>
        /// Get the serial port output from a Compute Instance. For more information see
        /// the official [API](https://cloud.google.com/compute/docs/instances/viewing-serial-port-output) documentation.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var serial = Output.Create(Gcp.Compute.GetInstanceSerialPort.InvokeAsync(new Gcp.Compute.GetInstanceSerialPortArgs
        ///         {
        ///             Instance = "my-instance",
        ///             Zone = "us-central1-a",
        ///             Port = 1,
        ///         }));
        ///         this.SerialOut = serial.Apply(serial =&gt; serial.Contents);
        ///     }
        /// 
        ///     [Output("serialOut")]
        ///     public Output&lt;string&gt; SerialOut { get; set; }
        /// }
        /// ```
        /// 
        /// Using the serial port output to generate a windows password, derived from the [official guide](https://cloud.google.com/compute/docs/instances/windows/automate-pw-generation):
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Text.Json;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var windows = new Gcp.Compute.Instance("windows", new Gcp.Compute.InstanceArgs
        ///         {
        ///             MachineType = "n1-standard-1",
        ///             Zone = "us-central1-a",
        ///             BootDisk = new Gcp.Compute.Inputs.InstanceBootDiskArgs
        ///             {
        ///                 InitializeParams = new Gcp.Compute.Inputs.InstanceBootDiskInitializeParamsArgs
        ///                 {
        ///                     Image = "gce-uefi-images/windows-2019",
        ///                 },
        ///             },
        ///             NetworkInterfaces = 
        ///             {
        ///                 new Gcp.Compute.Inputs.InstanceNetworkInterfaceArgs
        ///                 {
        ///                     Network = "default",
        ///                     AccessConfigs = 
        ///                     {
        ///                         ,
        ///                     },
        ///                 },
        ///             },
        ///             Metadata = 
        ///             {
        ///                 { "serial-port-logging-enable", "TRUE" },
        ///                 { "windows-keys", JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///                 {
        ///                     { "email", "example.user@example.com" },
        ///                     { "expireOn", "2020-04-14T01:37:19Z" },
        ///                     { "exponent", "AQAB" },
        ///                     { "modulus", "wgsquN4IBNPqIUnu+h/5Za1kujb2YRhX1vCQVQAkBwnWigcCqOBVfRa5JoZfx6KIvEXjWqa77jPvlsxM4WPqnDIM2qiK36up3SKkYwFjff6F2ni/ry8vrwXCX3sGZ1hbIHlK0O012HpA3ISeEswVZmX2X67naOvJXfY5v0hGPWqCADao+xVxrmxsZD4IWnKl1UaZzI5lhAzr8fw6utHwx1EZ/MSgsEki6tujcZfN+GUDRnmJGQSnPTXmsf7Q4DKreTZk49cuyB3prV91S0x3DYjCUpSXrkVy1Ha5XicGD/q+ystuFsJnrrhbNXJbpSjM6sjo/aduAkZJl4FmOt0R7Q==" },
        ///                     { "userName", "example-user" },
        ///                 }) },
        ///             },
        ///             ServiceAccount = new Gcp.Compute.Inputs.InstanceServiceAccountArgs
        ///             {
        ///                 Scopes = 
        ///                 {
        ///                     "userinfo-email",
        ///                     "compute-ro",
        ///                     "storage-ro",
        ///                 },
        ///             },
        ///         });
        ///         var serial = Output.Tuple(windows.Name, windows.Zone).Apply(values =&gt;
        ///         {
        ///             var name = values.Item1;
        ///             var zone = values.Item2;
        ///             return Gcp.Compute.GetInstanceSerialPort.InvokeAsync(new Gcp.Compute.GetInstanceSerialPortArgs
        ///             {
        ///                 Instance = name,
        ///                 Zone = zone,
        ///                 Port = 4,
        ///             });
        ///         });
        ///         this.SerialOut = serial.Apply(serial =&gt; serial.Contents);
        ///     }
        /// 
        ///     [Output("serialOut")]
        ///     public Output&lt;string&gt; SerialOut { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceSerialPortResult> InvokeAsync(GetInstanceSerialPortArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceSerialPortResult>("gcp:compute/getInstanceSerialPort:getInstanceSerialPort", args ?? new GetInstanceSerialPortArgs(), options.WithVersion());
    }


    public sealed class GetInstanceSerialPortArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Compute Instance to read output from.
        /// </summary>
        [Input("instance", required: true)]
        public string Instance { get; set; } = null!;

        /// <summary>
        /// The number of the serial port to read output from. Possible values are 1-4.
        /// </summary>
        [Input("port", required: true)]
        public int Port { get; set; }

        /// <summary>
        /// The project in which the Compute Instance exists. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The zone in which the Compute Instance exists.
        /// If it is not provided, the provider zone is used.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstanceSerialPortArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceSerialPortResult
    {
        /// <summary>
        /// The output of the serial port. Serial port output is available only when the VM instance is running, and logs are limited to the most recent 1 MB of output per port.
        /// </summary>
        public readonly string Contents;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Instance;
        public readonly int Port;
        public readonly string Project;
        public readonly string Zone;

        [OutputConstructor]
        private GetInstanceSerialPortResult(
            string contents,

            string id,

            string instance,

            int port,

            string project,

            string zone)
        {
            Contents = contents;
            Id = id;
            Instance = instance;
            Port = port;
            Project = project;
            Zone = zone;
        }
    }
}
