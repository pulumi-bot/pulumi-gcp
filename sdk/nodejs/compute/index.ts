// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./address";
export * from "./attachedDisk";
export * from "./autoscalar";
export * from "./autoscaler";
export * from "./backendBucket";
export * from "./backendBucketSignedUrlKey";
export * from "./backendService";
export * from "./backendServiceSignedUrlKey";
export * from "./disk";
export * from "./diskIamBinding";
export * from "./diskIamMember";
export * from "./diskIamPolicy";
export * from "./diskResourcePolicyAttachment";
export * from "./externalVpnGateway";
export * from "./firewall";
export * from "./forwardingRule";
export * from "./getAddress";
export * from "./getBackendBucket";
export * from "./getBackendService";
export * from "./getCertificate";
export * from "./getDefaultServiceAccount";
export * from "./getForwardingRule";
export * from "./getGlobalAddress";
export * from "./getGlobalForwardingRule";
export * from "./getImage";
export * from "./getInstance";
export * from "./getInstanceGroup";
export * from "./getInstanceSerialPort";
export * from "./getInstanceTemplate";
export * from "./getLBIPRanges";
export * from "./getNetblockIPRanges";
export * from "./getNetwork";
export * from "./getNetworkEndpointGroup";
export * from "./getNodeTypes";
export * from "./getRegionInstanceGroup";
export * from "./getRegionSslCertificate";
export * from "./getRegions";
export * from "./getResourcePolicy";
export * from "./getRouter";
export * from "./getSSLPolicy";
export * from "./getSubnetwork";
export * from "./getVPNGateway";
export * from "./getZones";
export * from "./globalAddress";
export * from "./globalForwardingRule";
export * from "./globalNetworkEndpoint";
export * from "./globalNetworkEndpointGroup";
export * from "./haVpnGateway";
export * from "./healthCheck";
export * from "./httpHealthCheck";
export * from "./httpsHealthCheck";
export * from "./image";
export * from "./imageIamBinding";
export * from "./imageIamMember";
export * from "./imageIamPolicy";
export * from "./instance";
export * from "./instanceFromMachineImage";
export * from "./instanceFromTemplate";
export * from "./instanceGroup";
export * from "./instanceGroupManager";
export * from "./instanceGroupNamedPort";
export * from "./instanceIAMBinding";
export * from "./instanceIAMMember";
export * from "./instanceIAMPolicy";
export * from "./instanceTemplate";
export * from "./interconnectAttachment";
export * from "./machineImage";
export * from "./machineImageIamBinding";
export * from "./machineImageIamMember";
export * from "./machineImageIamPolicy";
export * from "./managedSslCertificate";
export * from "./mangedSslCertificate";
export * from "./network";
export * from "./networkEndpoint";
export * from "./networkEndpointGroup";
export * from "./networkPeering";
export * from "./networkPeeringRoutesConfig";
export * from "./nodeGroup";
export * from "./nodeTemplate";
export * from "./organizationSecurityPolicy";
export * from "./organizationSecurityPolicyAssociation";
export * from "./organizationSecurityPolicyRule";
export * from "./packetMirroring";
export * from "./perInstanceConfig";
export * from "./projectDefaultNetworkTier";
export * from "./projectMetadata";
export * from "./projectMetadataItem";
export * from "./regionAutoscaler";
export * from "./regionBackendService";
export * from "./regionDisk";
export * from "./regionDiskIamBinding";
export * from "./regionDiskIamMember";
export * from "./regionDiskIamPolicy";
export * from "./regionDiskResourcePolicyAttachment";
export * from "./regionHealthCheck";
export * from "./regionInstanceGroupManager";
export * from "./regionNetworkEndpointGroup";
export * from "./regionPerInstanceConfig";
export * from "./regionSslCertificate";
export * from "./regionTargetHttpProxy";
export * from "./regionTargetHttpsProxy";
export * from "./regionUrlMap";
export * from "./reservation";
export * from "./resourcePolicy";
export * from "./route";
export * from "./router";
export * from "./routerInterface";
export * from "./routerNat";
export * from "./routerPeer";
export * from "./securityPolicy";
export * from "./securityScanConfig";
export * from "./sharedVPCHostProject";
export * from "./sharedVPCServiceProject";
export * from "./snapshot";
export * from "./sslcertificate";
export * from "./sslpolicy";
export * from "./subnetwork";
export * from "./subnetworkIAMBinding";
export * from "./subnetworkIAMMember";
export * from "./subnetworkIAMPolicy";
export * from "./targetGrpcProxy";
export * from "./targetHttpProxy";
export * from "./targetHttpsProxy";
export * from "./targetInstance";
export * from "./targetPool";
export * from "./targetSSLProxy";
export * from "./targetTCPProxy";
export * from "./urlmap";
export * from "./vpngateway";
export * from "./vpntunnel";

// Import resources to register:
import { Address } from "./address";
import { AttachedDisk } from "./attachedDisk";
import { Autoscalar } from "./autoscalar";
import { Autoscaler } from "./autoscaler";
import { BackendBucket } from "./backendBucket";
import { BackendBucketSignedUrlKey } from "./backendBucketSignedUrlKey";
import { BackendService } from "./backendService";
import { BackendServiceSignedUrlKey } from "./backendServiceSignedUrlKey";
import { Disk } from "./disk";
import { DiskIamBinding } from "./diskIamBinding";
import { DiskIamMember } from "./diskIamMember";
import { DiskIamPolicy } from "./diskIamPolicy";
import { DiskResourcePolicyAttachment } from "./diskResourcePolicyAttachment";
import { ExternalVpnGateway } from "./externalVpnGateway";
import { Firewall } from "./firewall";
import { ForwardingRule } from "./forwardingRule";
import { GlobalAddress } from "./globalAddress";
import { GlobalForwardingRule } from "./globalForwardingRule";
import { GlobalNetworkEndpoint } from "./globalNetworkEndpoint";
import { GlobalNetworkEndpointGroup } from "./globalNetworkEndpointGroup";
import { HaVpnGateway } from "./haVpnGateway";
import { HealthCheck } from "./healthCheck";
import { HttpHealthCheck } from "./httpHealthCheck";
import { HttpsHealthCheck } from "./httpsHealthCheck";
import { Image } from "./image";
import { ImageIamBinding } from "./imageIamBinding";
import { ImageIamMember } from "./imageIamMember";
import { ImageIamPolicy } from "./imageIamPolicy";
import { Instance } from "./instance";
import { InstanceFromMachineImage } from "./instanceFromMachineImage";
import { InstanceFromTemplate } from "./instanceFromTemplate";
import { InstanceGroup } from "./instanceGroup";
import { InstanceGroupManager } from "./instanceGroupManager";
import { InstanceGroupNamedPort } from "./instanceGroupNamedPort";
import { InstanceIAMBinding } from "./instanceIAMBinding";
import { InstanceIAMMember } from "./instanceIAMMember";
import { InstanceIAMPolicy } from "./instanceIAMPolicy";
import { InstanceTemplate } from "./instanceTemplate";
import { InterconnectAttachment } from "./interconnectAttachment";
import { MachineImage } from "./machineImage";
import { MachineImageIamBinding } from "./machineImageIamBinding";
import { MachineImageIamMember } from "./machineImageIamMember";
import { MachineImageIamPolicy } from "./machineImageIamPolicy";
import { ManagedSslCertificate } from "./managedSslCertificate";
import { MangedSslCertificate } from "./mangedSslCertificate";
import { Network } from "./network";
import { NetworkEndpoint } from "./networkEndpoint";
import { NetworkEndpointGroup } from "./networkEndpointGroup";
import { NetworkPeering } from "./networkPeering";
import { NetworkPeeringRoutesConfig } from "./networkPeeringRoutesConfig";
import { NodeGroup } from "./nodeGroup";
import { NodeTemplate } from "./nodeTemplate";
import { OrganizationSecurityPolicy } from "./organizationSecurityPolicy";
import { OrganizationSecurityPolicyAssociation } from "./organizationSecurityPolicyAssociation";
import { OrganizationSecurityPolicyRule } from "./organizationSecurityPolicyRule";
import { PacketMirroring } from "./packetMirroring";
import { PerInstanceConfig } from "./perInstanceConfig";
import { ProjectDefaultNetworkTier } from "./projectDefaultNetworkTier";
import { ProjectMetadata } from "./projectMetadata";
import { ProjectMetadataItem } from "./projectMetadataItem";
import { RegionAutoscaler } from "./regionAutoscaler";
import { RegionBackendService } from "./regionBackendService";
import { RegionDisk } from "./regionDisk";
import { RegionDiskIamBinding } from "./regionDiskIamBinding";
import { RegionDiskIamMember } from "./regionDiskIamMember";
import { RegionDiskIamPolicy } from "./regionDiskIamPolicy";
import { RegionDiskResourcePolicyAttachment } from "./regionDiskResourcePolicyAttachment";
import { RegionHealthCheck } from "./regionHealthCheck";
import { RegionInstanceGroupManager } from "./regionInstanceGroupManager";
import { RegionNetworkEndpointGroup } from "./regionNetworkEndpointGroup";
import { RegionPerInstanceConfig } from "./regionPerInstanceConfig";
import { RegionSslCertificate } from "./regionSslCertificate";
import { RegionTargetHttpProxy } from "./regionTargetHttpProxy";
import { RegionTargetHttpsProxy } from "./regionTargetHttpsProxy";
import { RegionUrlMap } from "./regionUrlMap";
import { Reservation } from "./reservation";
import { ResourcePolicy } from "./resourcePolicy";
import { Route } from "./route";
import { Router } from "./router";
import { RouterInterface } from "./routerInterface";
import { RouterNat } from "./routerNat";
import { RouterPeer } from "./routerPeer";
import { SSLCertificate } from "./sslcertificate";
import { SSLPolicy } from "./sslpolicy";
import { SecurityPolicy } from "./securityPolicy";
import { SecurityScanConfig } from "./securityScanConfig";
import { SharedVPCHostProject } from "./sharedVPCHostProject";
import { SharedVPCServiceProject } from "./sharedVPCServiceProject";
import { Snapshot } from "./snapshot";
import { Subnetwork } from "./subnetwork";
import { SubnetworkIAMBinding } from "./subnetworkIAMBinding";
import { SubnetworkIAMMember } from "./subnetworkIAMMember";
import { SubnetworkIAMPolicy } from "./subnetworkIAMPolicy";
import { TargetGrpcProxy } from "./targetGrpcProxy";
import { TargetHttpProxy } from "./targetHttpProxy";
import { TargetHttpsProxy } from "./targetHttpsProxy";
import { TargetInstance } from "./targetInstance";
import { TargetPool } from "./targetPool";
import { TargetSSLProxy } from "./targetSSLProxy";
import { TargetTCPProxy } from "./targetTCPProxy";
import { URLMap } from "./urlmap";
import { VPNGateway } from "./vpngateway";
import { VPNTunnel } from "./vpntunnel";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:compute/address:Address":
                return new Address(name, <any>undefined, { urn })
            case "gcp:compute/attachedDisk:AttachedDisk":
                return new AttachedDisk(name, <any>undefined, { urn })
            case "gcp:compute/autoscalar:Autoscalar":
                return new Autoscalar(name, <any>undefined, { urn })
            case "gcp:compute/autoscaler:Autoscaler":
                return new Autoscaler(name, <any>undefined, { urn })
            case "gcp:compute/backendBucket:BackendBucket":
                return new BackendBucket(name, <any>undefined, { urn })
            case "gcp:compute/backendBucketSignedUrlKey:BackendBucketSignedUrlKey":
                return new BackendBucketSignedUrlKey(name, <any>undefined, { urn })
            case "gcp:compute/backendService:BackendService":
                return new BackendService(name, <any>undefined, { urn })
            case "gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey":
                return new BackendServiceSignedUrlKey(name, <any>undefined, { urn })
            case "gcp:compute/disk:Disk":
                return new Disk(name, <any>undefined, { urn })
            case "gcp:compute/diskIamBinding:DiskIamBinding":
                return new DiskIamBinding(name, <any>undefined, { urn })
            case "gcp:compute/diskIamMember:DiskIamMember":
                return new DiskIamMember(name, <any>undefined, { urn })
            case "gcp:compute/diskIamPolicy:DiskIamPolicy":
                return new DiskIamPolicy(name, <any>undefined, { urn })
            case "gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment":
                return new DiskResourcePolicyAttachment(name, <any>undefined, { urn })
            case "gcp:compute/externalVpnGateway:ExternalVpnGateway":
                return new ExternalVpnGateway(name, <any>undefined, { urn })
            case "gcp:compute/firewall:Firewall":
                return new Firewall(name, <any>undefined, { urn })
            case "gcp:compute/forwardingRule:ForwardingRule":
                return new ForwardingRule(name, <any>undefined, { urn })
            case "gcp:compute/globalAddress:GlobalAddress":
                return new GlobalAddress(name, <any>undefined, { urn })
            case "gcp:compute/globalForwardingRule:GlobalForwardingRule":
                return new GlobalForwardingRule(name, <any>undefined, { urn })
            case "gcp:compute/globalNetworkEndpoint:GlobalNetworkEndpoint":
                return new GlobalNetworkEndpoint(name, <any>undefined, { urn })
            case "gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup":
                return new GlobalNetworkEndpointGroup(name, <any>undefined, { urn })
            case "gcp:compute/haVpnGateway:HaVpnGateway":
                return new HaVpnGateway(name, <any>undefined, { urn })
            case "gcp:compute/healthCheck:HealthCheck":
                return new HealthCheck(name, <any>undefined, { urn })
            case "gcp:compute/httpHealthCheck:HttpHealthCheck":
                return new HttpHealthCheck(name, <any>undefined, { urn })
            case "gcp:compute/httpsHealthCheck:HttpsHealthCheck":
                return new HttpsHealthCheck(name, <any>undefined, { urn })
            case "gcp:compute/image:Image":
                return new Image(name, <any>undefined, { urn })
            case "gcp:compute/imageIamBinding:ImageIamBinding":
                return new ImageIamBinding(name, <any>undefined, { urn })
            case "gcp:compute/imageIamMember:ImageIamMember":
                return new ImageIamMember(name, <any>undefined, { urn })
            case "gcp:compute/imageIamPolicy:ImageIamPolicy":
                return new ImageIamPolicy(name, <any>undefined, { urn })
            case "gcp:compute/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "gcp:compute/instanceFromMachineImage:InstanceFromMachineImage":
                return new InstanceFromMachineImage(name, <any>undefined, { urn })
            case "gcp:compute/instanceFromTemplate:InstanceFromTemplate":
                return new InstanceFromTemplate(name, <any>undefined, { urn })
            case "gcp:compute/instanceGroup:InstanceGroup":
                return new InstanceGroup(name, <any>undefined, { urn })
            case "gcp:compute/instanceGroupManager:InstanceGroupManager":
                return new InstanceGroupManager(name, <any>undefined, { urn })
            case "gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort":
                return new InstanceGroupNamedPort(name, <any>undefined, { urn })
            case "gcp:compute/instanceIAMBinding:InstanceIAMBinding":
                return new InstanceIAMBinding(name, <any>undefined, { urn })
            case "gcp:compute/instanceIAMMember:InstanceIAMMember":
                return new InstanceIAMMember(name, <any>undefined, { urn })
            case "gcp:compute/instanceIAMPolicy:InstanceIAMPolicy":
                return new InstanceIAMPolicy(name, <any>undefined, { urn })
            case "gcp:compute/instanceTemplate:InstanceTemplate":
                return new InstanceTemplate(name, <any>undefined, { urn })
            case "gcp:compute/interconnectAttachment:InterconnectAttachment":
                return new InterconnectAttachment(name, <any>undefined, { urn })
            case "gcp:compute/machineImage:MachineImage":
                return new MachineImage(name, <any>undefined, { urn })
            case "gcp:compute/machineImageIamBinding:MachineImageIamBinding":
                return new MachineImageIamBinding(name, <any>undefined, { urn })
            case "gcp:compute/machineImageIamMember:MachineImageIamMember":
                return new MachineImageIamMember(name, <any>undefined, { urn })
            case "gcp:compute/machineImageIamPolicy:MachineImageIamPolicy":
                return new MachineImageIamPolicy(name, <any>undefined, { urn })
            case "gcp:compute/managedSslCertificate:ManagedSslCertificate":
                return new ManagedSslCertificate(name, <any>undefined, { urn })
            case "gcp:compute/mangedSslCertificate:MangedSslCertificate":
                return new MangedSslCertificate(name, <any>undefined, { urn })
            case "gcp:compute/network:Network":
                return new Network(name, <any>undefined, { urn })
            case "gcp:compute/networkEndpoint:NetworkEndpoint":
                return new NetworkEndpoint(name, <any>undefined, { urn })
            case "gcp:compute/networkEndpointGroup:NetworkEndpointGroup":
                return new NetworkEndpointGroup(name, <any>undefined, { urn })
            case "gcp:compute/networkPeering:NetworkPeering":
                return new NetworkPeering(name, <any>undefined, { urn })
            case "gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig":
                return new NetworkPeeringRoutesConfig(name, <any>undefined, { urn })
            case "gcp:compute/nodeGroup:NodeGroup":
                return new NodeGroup(name, <any>undefined, { urn })
            case "gcp:compute/nodeTemplate:NodeTemplate":
                return new NodeTemplate(name, <any>undefined, { urn })
            case "gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy":
                return new OrganizationSecurityPolicy(name, <any>undefined, { urn })
            case "gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation":
                return new OrganizationSecurityPolicyAssociation(name, <any>undefined, { urn })
            case "gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule":
                return new OrganizationSecurityPolicyRule(name, <any>undefined, { urn })
            case "gcp:compute/packetMirroring:PacketMirroring":
                return new PacketMirroring(name, <any>undefined, { urn })
            case "gcp:compute/perInstanceConfig:PerInstanceConfig":
                return new PerInstanceConfig(name, <any>undefined, { urn })
            case "gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier":
                return new ProjectDefaultNetworkTier(name, <any>undefined, { urn })
            case "gcp:compute/projectMetadata:ProjectMetadata":
                return new ProjectMetadata(name, <any>undefined, { urn })
            case "gcp:compute/projectMetadataItem:ProjectMetadataItem":
                return new ProjectMetadataItem(name, <any>undefined, { urn })
            case "gcp:compute/regionAutoscaler:RegionAutoscaler":
                return new RegionAutoscaler(name, <any>undefined, { urn })
            case "gcp:compute/regionBackendService:RegionBackendService":
                return new RegionBackendService(name, <any>undefined, { urn })
            case "gcp:compute/regionDisk:RegionDisk":
                return new RegionDisk(name, <any>undefined, { urn })
            case "gcp:compute/regionDiskIamBinding:RegionDiskIamBinding":
                return new RegionDiskIamBinding(name, <any>undefined, { urn })
            case "gcp:compute/regionDiskIamMember:RegionDiskIamMember":
                return new RegionDiskIamMember(name, <any>undefined, { urn })
            case "gcp:compute/regionDiskIamPolicy:RegionDiskIamPolicy":
                return new RegionDiskIamPolicy(name, <any>undefined, { urn })
            case "gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment":
                return new RegionDiskResourcePolicyAttachment(name, <any>undefined, { urn })
            case "gcp:compute/regionHealthCheck:RegionHealthCheck":
                return new RegionHealthCheck(name, <any>undefined, { urn })
            case "gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager":
                return new RegionInstanceGroupManager(name, <any>undefined, { urn })
            case "gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup":
                return new RegionNetworkEndpointGroup(name, <any>undefined, { urn })
            case "gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig":
                return new RegionPerInstanceConfig(name, <any>undefined, { urn })
            case "gcp:compute/regionSslCertificate:RegionSslCertificate":
                return new RegionSslCertificate(name, <any>undefined, { urn })
            case "gcp:compute/regionTargetHttpProxy:RegionTargetHttpProxy":
                return new RegionTargetHttpProxy(name, <any>undefined, { urn })
            case "gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy":
                return new RegionTargetHttpsProxy(name, <any>undefined, { urn })
            case "gcp:compute/regionUrlMap:RegionUrlMap":
                return new RegionUrlMap(name, <any>undefined, { urn })
            case "gcp:compute/reservation:Reservation":
                return new Reservation(name, <any>undefined, { urn })
            case "gcp:compute/resourcePolicy:ResourcePolicy":
                return new ResourcePolicy(name, <any>undefined, { urn })
            case "gcp:compute/route:Route":
                return new Route(name, <any>undefined, { urn })
            case "gcp:compute/router:Router":
                return new Router(name, <any>undefined, { urn })
            case "gcp:compute/routerInterface:RouterInterface":
                return new RouterInterface(name, <any>undefined, { urn })
            case "gcp:compute/routerNat:RouterNat":
                return new RouterNat(name, <any>undefined, { urn })
            case "gcp:compute/routerPeer:RouterPeer":
                return new RouterPeer(name, <any>undefined, { urn })
            case "gcp:compute/sSLCertificate:SSLCertificate":
                return new SSLCertificate(name, <any>undefined, { urn })
            case "gcp:compute/sSLPolicy:SSLPolicy":
                return new SSLPolicy(name, <any>undefined, { urn })
            case "gcp:compute/securityPolicy:SecurityPolicy":
                return new SecurityPolicy(name, <any>undefined, { urn })
            case "gcp:compute/securityScanConfig:SecurityScanConfig":
                return new SecurityScanConfig(name, <any>undefined, { urn })
            case "gcp:compute/sharedVPCHostProject:SharedVPCHostProject":
                return new SharedVPCHostProject(name, <any>undefined, { urn })
            case "gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject":
                return new SharedVPCServiceProject(name, <any>undefined, { urn })
            case "gcp:compute/snapshot:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "gcp:compute/subnetwork:Subnetwork":
                return new Subnetwork(name, <any>undefined, { urn })
            case "gcp:compute/subnetworkIAMBinding:SubnetworkIAMBinding":
                return new SubnetworkIAMBinding(name, <any>undefined, { urn })
            case "gcp:compute/subnetworkIAMMember:SubnetworkIAMMember":
                return new SubnetworkIAMMember(name, <any>undefined, { urn })
            case "gcp:compute/subnetworkIAMPolicy:SubnetworkIAMPolicy":
                return new SubnetworkIAMPolicy(name, <any>undefined, { urn })
            case "gcp:compute/targetGrpcProxy:TargetGrpcProxy":
                return new TargetGrpcProxy(name, <any>undefined, { urn })
            case "gcp:compute/targetHttpProxy:TargetHttpProxy":
                return new TargetHttpProxy(name, <any>undefined, { urn })
            case "gcp:compute/targetHttpsProxy:TargetHttpsProxy":
                return new TargetHttpsProxy(name, <any>undefined, { urn })
            case "gcp:compute/targetInstance:TargetInstance":
                return new TargetInstance(name, <any>undefined, { urn })
            case "gcp:compute/targetPool:TargetPool":
                return new TargetPool(name, <any>undefined, { urn })
            case "gcp:compute/targetSSLProxy:TargetSSLProxy":
                return new TargetSSLProxy(name, <any>undefined, { urn })
            case "gcp:compute/targetTCPProxy:TargetTCPProxy":
                return new TargetTCPProxy(name, <any>undefined, { urn })
            case "gcp:compute/uRLMap:URLMap":
                return new URLMap(name, <any>undefined, { urn })
            case "gcp:compute/vPNGateway:VPNGateway":
                return new VPNGateway(name, <any>undefined, { urn })
            case "gcp:compute/vPNTunnel:VPNTunnel":
                return new VPNTunnel(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "compute/address", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/attachedDisk", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/autoscalar", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/autoscaler", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/backendBucket", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/backendBucketSignedUrlKey", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/backendService", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/backendServiceSignedUrlKey", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/disk", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/diskIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/diskIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/diskIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/diskResourcePolicyAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/externalVpnGateway", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/firewall", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/forwardingRule", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/globalAddress", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/globalForwardingRule", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/globalNetworkEndpoint", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/globalNetworkEndpointGroup", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/haVpnGateway", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/healthCheck", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/httpHealthCheck", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/httpsHealthCheck", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/image", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/imageIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/imageIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/imageIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instance", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceFromMachineImage", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceFromTemplate", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceGroup", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceGroupManager", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceGroupNamedPort", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceIAMBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceIAMMember", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceIAMPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/instanceTemplate", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/interconnectAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/machineImage", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/machineImageIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/machineImageIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/machineImageIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/managedSslCertificate", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/mangedSslCertificate", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/network", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/networkEndpoint", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/networkEndpointGroup", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/networkPeering", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/networkPeeringRoutesConfig", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/nodeGroup", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/nodeTemplate", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/organizationSecurityPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/organizationSecurityPolicyAssociation", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/organizationSecurityPolicyRule", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/packetMirroring", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/perInstanceConfig", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/projectDefaultNetworkTier", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/projectMetadata", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/projectMetadataItem", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionAutoscaler", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionBackendService", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionDisk", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionDiskIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionDiskIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionDiskIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionDiskResourcePolicyAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionHealthCheck", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionInstanceGroupManager", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionNetworkEndpointGroup", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionPerInstanceConfig", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionSslCertificate", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionTargetHttpProxy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionTargetHttpsProxy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/regionUrlMap", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/reservation", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/resourcePolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/route", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/router", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/routerInterface", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/routerNat", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/routerPeer", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/sSLCertificate", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/sSLPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/securityPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/securityScanConfig", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/sharedVPCHostProject", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/sharedVPCServiceProject", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/snapshot", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/subnetwork", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/subnetworkIAMBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/subnetworkIAMMember", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/subnetworkIAMPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/targetGrpcProxy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/targetHttpProxy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/targetHttpsProxy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/targetInstance", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/targetPool", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/targetSSLProxy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/targetTCPProxy", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/uRLMap", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/vPNGateway", _module)
pulumi.runtime.registerResourceModule("gcp", "compute/vPNTunnel", _module)
