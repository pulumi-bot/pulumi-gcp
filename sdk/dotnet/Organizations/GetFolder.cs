// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static class GetFolder
    {
        /// <summary>
        /// Use this data source to get information about a Google Cloud Folder.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var myFolder1 = Output.Create(Gcp.Organizations.GetFolder.InvokeAsync(new Gcp.Organizations.GetFolderArgs
        ///         {
        ///             Folder = "folders/12345",
        ///             LookupOrganization = true,
        ///         }));
        ///         var myFolder2 = Output.Create(Gcp.Organizations.GetFolder.InvokeAsync(new Gcp.Organizations.GetFolderArgs
        ///         {
        ///             Folder = "folders/23456",
        ///         }));
        ///         this.MyFolder1Organization = myFolder1.Apply(myFolder1 =&gt; myFolder1.Organization);
        ///         this.MyFolder2Parent = myFolder2.Apply(myFolder2 =&gt; myFolder2.Parent);
        ///     }
        /// 
        ///     [Output("myFolder1Organization")]
        ///     public Output&lt;string&gt; MyFolder1Organization { get; set; }
        ///     [Output("myFolder2Parent")]
        ///     public Output&lt;string&gt; MyFolder2Parent { get; set; }
        /// }
        /// ```
        /// </summary>
        public static Task<GetFolderResult> InvokeAsync(GetFolderArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFolderResult>("gcp:organizations/getFolder:getFolder", args ?? new GetFolderArgs(), options.WithVersion());

        public static Output<GetFolderResult> Invoke(GetFolderOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Folder.Box(),
                args.LookupOrganization.Box()
            ).Apply(a => {
                    var args = new GetFolderArgs();
                    a[0].Set(args, nameof(args.Folder));
                    a[1].Set(args, nameof(args.LookupOrganization));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetFolderArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Folder in the form `{folder_id}` or `folders/{folder_id}`.
        /// </summary>
        [Input("folder", required: true)]
        public string Folder { get; set; } = null!;

        /// <summary>
        /// `true` to find the organization that the folder belongs, `false` to avoid the lookup. It searches up the tree. (defaults to `false`)
        /// </summary>
        [Input("lookupOrganization")]
        public bool? LookupOrganization { get; set; }

        public GetFolderArgs()
        {
        }
    }

    public sealed class GetFolderOutputArgs
    {
        /// <summary>
        /// The name of the Folder in the form `{folder_id}` or `folders/{folder_id}`.
        /// </summary>
        [Input("folder", required: true)]
        public Input<string> Folder { get; set; } = null!;

        /// <summary>
        /// `true` to find the organization that the folder belongs, `false` to avoid the lookup. It searches up the tree. (defaults to `false`)
        /// </summary>
        [Input("lookupOrganization")]
        public Input<bool>? LookupOrganization { get; set; }

        public GetFolderOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFolderResult
    {
        /// <summary>
        /// Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The folder's display name.
        /// </summary>
        public readonly string DisplayName;
        public readonly string Folder;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Folder's current lifecycle state.
        /// </summary>
        public readonly string LifecycleState;
        public readonly bool? LookupOrganization;
        /// <summary>
        /// The resource name of the Folder in the form `folders/{folder_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If `lookup_organization` is enable, the resource name of the Organization that the folder belongs.
        /// </summary>
        public readonly string Organization;
        /// <summary>
        /// The resource name of the parent Folder or Organization.
        /// </summary>
        public readonly string Parent;

        [OutputConstructor]
        private GetFolderResult(
            string createTime,

            string displayName,

            string folder,

            string folderId,

            string id,

            string lifecycleState,

            bool? lookupOrganization,

            string name,

            string organization,

            string parent)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            Folder = folder;
            FolderId = folderId;
            Id = id;
            LifecycleState = lifecycleState;
            LookupOrganization = lookupOrganization;
            Name = name;
            Organization = organization;
            Parent = parent;
        }
    }
}
