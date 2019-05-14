// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/slipway/slipway.proto

package slipway

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MembersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MembersRequest) Reset()         { *m = MembersRequest{} }
func (m *MembersRequest) String() string { return proto.CompactTextString(m) }
func (*MembersRequest) ProtoMessage()    {}
func (*MembersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{0}
}

func (m *MembersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MembersRequest.Unmarshal(m, b)
}
func (m *MembersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MembersRequest.Marshal(b, m, deterministic)
}
func (m *MembersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MembersRequest.Merge(m, src)
}
func (m *MembersRequest) XXX_Size() int {
	return xxx_messageInfo_MembersRequest.Size(m)
}
func (m *MembersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MembersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MembersRequest proto.InternalMessageInfo

type MembersResponse struct {
	Hanlder              *HandlerMember   `protobuf:"bytes,1,opt,name=hanlder,proto3" json:"hanlder,omitempty"`
	Result               []*ClusterMember `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MembersResponse) Reset()         { *m = MembersResponse{} }
func (m *MembersResponse) String() string { return proto.CompactTextString(m) }
func (*MembersResponse) ProtoMessage()    {}
func (*MembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{1}
}

func (m *MembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MembersResponse.Unmarshal(m, b)
}
func (m *MembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MembersResponse.Marshal(b, m, deterministic)
}
func (m *MembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MembersResponse.Merge(m, src)
}
func (m *MembersResponse) XXX_Size() int {
	return xxx_messageInfo_MembersResponse.Size(m)
}
func (m *MembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MembersResponse proto.InternalMessageInfo

func (m *MembersResponse) GetHanlder() *HandlerMember {
	if m != nil {
		return m.Hanlder
	}
	return nil
}

func (m *MembersResponse) GetResult() []*ClusterMember {
	if m != nil {
		return m.Result
	}
	return nil
}

type ClusterMember struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Labels               []string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	IpAddress            string   `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterMember) Reset()         { *m = ClusterMember{} }
func (m *ClusterMember) String() string { return proto.CompactTextString(m) }
func (*ClusterMember) ProtoMessage()    {}
func (*ClusterMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{2}
}

func (m *ClusterMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterMember.Unmarshal(m, b)
}
func (m *ClusterMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterMember.Marshal(b, m, deterministic)
}
func (m *ClusterMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterMember.Merge(m, src)
}
func (m *ClusterMember) XXX_Size() int {
	return xxx_messageInfo_ClusterMember.Size(m)
}
func (m *ClusterMember) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterMember.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterMember proto.InternalMessageInfo

func (m *ClusterMember) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ClusterMember) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ClusterMember) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ClusterMember) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type InitRequest struct {
	// 1 - 15: Core Cluster Settings
	// 16 - 50: Extra Cluster Settings
	// 51 - 90: Config File Yamls, Cert and other files
	// 91 - 60: Run Params Other Extas
	DryRun                   bool     `protobuf:"varint,1,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	ClusterDnsDomain         string   `protobuf:"bytes,2,opt,name=cluster_dns_domain,json=clusterDnsDomain,proto3" json:"cluster_dns_domain,omitempty"`
	ControlPlaneUrl          string   `protobuf:"bytes,3,opt,name=control_plane_url,json=controlPlaneUrl,proto3" json:"control_plane_url,omitempty"`
	ClusterCni               string   `protobuf:"bytes,4,opt,name=cluster_cni,json=clusterCni,proto3" json:"cluster_cni,omitempty"`
	PodCidr                  string   `protobuf:"bytes,5,opt,name=pod_cidr,json=podCidr,proto3" json:"pod_cidr,omitempty"`
	ServiceCidr              string   `protobuf:"bytes,6,opt,name=service_cidr,json=serviceCidr,proto3" json:"service_cidr,omitempty"`
	ExtraKubeadmFlags        []string `protobuf:"bytes,15,rep,name=extra_kubeadm_flags,json=extraKubeadmFlags,proto3" json:"extra_kubeadm_flags,omitempty"`
	YamlClusterConfiguration []byte   `protobuf:"bytes,51,opt,name=yaml_cluster_configuration,json=yamlClusterConfiguration,proto3" json:"yaml_cluster_configuration,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{3}
}

func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (m *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(m, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

func (m *InitRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *InitRequest) GetClusterDnsDomain() string {
	if m != nil {
		return m.ClusterDnsDomain
	}
	return ""
}

func (m *InitRequest) GetControlPlaneUrl() string {
	if m != nil {
		return m.ControlPlaneUrl
	}
	return ""
}

func (m *InitRequest) GetClusterCni() string {
	if m != nil {
		return m.ClusterCni
	}
	return ""
}

func (m *InitRequest) GetPodCidr() string {
	if m != nil {
		return m.PodCidr
	}
	return ""
}

func (m *InitRequest) GetServiceCidr() string {
	if m != nil {
		return m.ServiceCidr
	}
	return ""
}

func (m *InitRequest) GetExtraKubeadmFlags() []string {
	if m != nil {
		return m.ExtraKubeadmFlags
	}
	return nil
}

func (m *InitRequest) GetYamlClusterConfiguration() []byte {
	if m != nil {
		return m.YamlClusterConfiguration
	}
	return nil
}

type InitResponse struct {
	Hanlder              *HandlerMember `protobuf:"bytes,1,opt,name=hanlder,proto3" json:"hanlder,omitempty"`
	HandledBy            string         `protobuf:"bytes,2,opt,name=handled_by,json=handledBy,proto3" json:"handled_by,omitempty"`
	LogStream            []byte         `protobuf:"bytes,3,opt,name=log_stream,json=logStream,proto3" json:"log_stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitResponse) Reset()         { *m = InitResponse{} }
func (m *InitResponse) String() string { return proto.CompactTextString(m) }
func (*InitResponse) ProtoMessage()    {}
func (*InitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{4}
}

func (m *InitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitResponse.Unmarshal(m, b)
}
func (m *InitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitResponse.Marshal(b, m, deterministic)
}
func (m *InitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitResponse.Merge(m, src)
}
func (m *InitResponse) XXX_Size() int {
	return xxx_messageInfo_InitResponse.Size(m)
}
func (m *InitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitResponse proto.InternalMessageInfo

func (m *InitResponse) GetHanlder() *HandlerMember {
	if m != nil {
		return m.Hanlder
	}
	return nil
}

func (m *InitResponse) GetHandledBy() string {
	if m != nil {
		return m.HandledBy
	}
	return ""
}

func (m *InitResponse) GetLogStream() []byte {
	if m != nil {
		return m.LogStream
	}
	return nil
}

type BootstrapTokenRequest struct {
	DryRun               string   `protobuf:"bytes,1,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	TimeToLive           string   `protobuf:"bytes,2,opt,name=time_to_live,json=timeToLive,proto3" json:"time_to_live,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	GroupsOverride       string   `protobuf:"bytes,4,opt,name=groups_override,json=groupsOverride,proto3" json:"groups_override,omitempty"`
	UsagesOverride       string   `protobuf:"bytes,5,opt,name=usages_override,json=usagesOverride,proto3" json:"usages_override,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BootstrapTokenRequest) Reset()         { *m = BootstrapTokenRequest{} }
func (m *BootstrapTokenRequest) String() string { return proto.CompactTextString(m) }
func (*BootstrapTokenRequest) ProtoMessage()    {}
func (*BootstrapTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{5}
}

func (m *BootstrapTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapTokenRequest.Unmarshal(m, b)
}
func (m *BootstrapTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapTokenRequest.Marshal(b, m, deterministic)
}
func (m *BootstrapTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapTokenRequest.Merge(m, src)
}
func (m *BootstrapTokenRequest) XXX_Size() int {
	return xxx_messageInfo_BootstrapTokenRequest.Size(m)
}
func (m *BootstrapTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapTokenRequest proto.InternalMessageInfo

func (m *BootstrapTokenRequest) GetDryRun() string {
	if m != nil {
		return m.DryRun
	}
	return ""
}

func (m *BootstrapTokenRequest) GetTimeToLive() string {
	if m != nil {
		return m.TimeToLive
	}
	return ""
}

func (m *BootstrapTokenRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BootstrapTokenRequest) GetGroupsOverride() string {
	if m != nil {
		return m.GroupsOverride
	}
	return ""
}

func (m *BootstrapTokenRequest) GetUsagesOverride() string {
	if m != nil {
		return m.UsagesOverride
	}
	return ""
}

type BootstrapTokenResponse struct {
	Hanlder              *HandlerMember `protobuf:"bytes,1,opt,name=hanlder,proto3" json:"hanlder,omitempty"`
	RawJoinCommand       string         `protobuf:"bytes,2,opt,name=raw_join_command,json=rawJoinCommand,proto3" json:"raw_join_command,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BootstrapTokenResponse) Reset()         { *m = BootstrapTokenResponse{} }
func (m *BootstrapTokenResponse) String() string { return proto.CompactTextString(m) }
func (*BootstrapTokenResponse) ProtoMessage()    {}
func (*BootstrapTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{6}
}

func (m *BootstrapTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapTokenResponse.Unmarshal(m, b)
}
func (m *BootstrapTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapTokenResponse.Marshal(b, m, deterministic)
}
func (m *BootstrapTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapTokenResponse.Merge(m, src)
}
func (m *BootstrapTokenResponse) XXX_Size() int {
	return xxx_messageInfo_BootstrapTokenResponse.Size(m)
}
func (m *BootstrapTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapTokenResponse proto.InternalMessageInfo

func (m *BootstrapTokenResponse) GetHanlder() *HandlerMember {
	if m != nil {
		return m.Hanlder
	}
	return nil
}

func (m *BootstrapTokenResponse) GetRawJoinCommand() string {
	if m != nil {
		return m.RawJoinCommand
	}
	return ""
}

type BootstrapMasterRequest struct {
	DryRun                 string   `protobuf:"bytes,1,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	TimeToLive             string   `protobuf:"bytes,2,opt,name=time_to_live,json=timeToLive,proto3" json:"time_to_live,omitempty"`
	Description            string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	EncryptionKey          string   `protobuf:"bytes,4,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
	SkipUploadCertificates bool     `protobuf:"varint,5,opt,name=skip_upload_certificates,json=skipUploadCertificates,proto3" json:"skip_upload_certificates,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *BootstrapMasterRequest) Reset()         { *m = BootstrapMasterRequest{} }
func (m *BootstrapMasterRequest) String() string { return proto.CompactTextString(m) }
func (*BootstrapMasterRequest) ProtoMessage()    {}
func (*BootstrapMasterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{7}
}

func (m *BootstrapMasterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapMasterRequest.Unmarshal(m, b)
}
func (m *BootstrapMasterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapMasterRequest.Marshal(b, m, deterministic)
}
func (m *BootstrapMasterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapMasterRequest.Merge(m, src)
}
func (m *BootstrapMasterRequest) XXX_Size() int {
	return xxx_messageInfo_BootstrapMasterRequest.Size(m)
}
func (m *BootstrapMasterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapMasterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapMasterRequest proto.InternalMessageInfo

func (m *BootstrapMasterRequest) GetDryRun() string {
	if m != nil {
		return m.DryRun
	}
	return ""
}

func (m *BootstrapMasterRequest) GetTimeToLive() string {
	if m != nil {
		return m.TimeToLive
	}
	return ""
}

func (m *BootstrapMasterRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BootstrapMasterRequest) GetEncryptionKey() string {
	if m != nil {
		return m.EncryptionKey
	}
	return ""
}

func (m *BootstrapMasterRequest) GetSkipUploadCertificates() bool {
	if m != nil {
		return m.SkipUploadCertificates
	}
	return false
}

type UploadCertsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadCertsRequest) Reset()         { *m = UploadCertsRequest{} }
func (m *UploadCertsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadCertsRequest) ProtoMessage()    {}
func (*UploadCertsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{8}
}

func (m *UploadCertsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadCertsRequest.Unmarshal(m, b)
}
func (m *UploadCertsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadCertsRequest.Marshal(b, m, deterministic)
}
func (m *UploadCertsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadCertsRequest.Merge(m, src)
}
func (m *UploadCertsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadCertsRequest.Size(m)
}
func (m *UploadCertsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadCertsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadCertsRequest proto.InternalMessageInfo

type UploadCertsResponse struct {
	Hanlder              *HandlerMember `protobuf:"bytes,1,opt,name=hanlder,proto3" json:"hanlder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UploadCertsResponse) Reset()         { *m = UploadCertsResponse{} }
func (m *UploadCertsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadCertsResponse) ProtoMessage()    {}
func (*UploadCertsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{9}
}

func (m *UploadCertsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadCertsResponse.Unmarshal(m, b)
}
func (m *UploadCertsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadCertsResponse.Marshal(b, m, deterministic)
}
func (m *UploadCertsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadCertsResponse.Merge(m, src)
}
func (m *UploadCertsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadCertsResponse.Size(m)
}
func (m *UploadCertsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadCertsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadCertsResponse proto.InternalMessageInfo

func (m *UploadCertsResponse) GetHanlder() *HandlerMember {
	if m != nil {
		return m.Hanlder
	}
	return nil
}

type HandlerMember struct {
	RequestId            string         `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Member               *ClusterMember `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HandlerMember) Reset()         { *m = HandlerMember{} }
func (m *HandlerMember) String() string { return proto.CompactTextString(m) }
func (*HandlerMember) ProtoMessage()    {}
func (*HandlerMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ddd0b3792332164, []int{10}
}

func (m *HandlerMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandlerMember.Unmarshal(m, b)
}
func (m *HandlerMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandlerMember.Marshal(b, m, deterministic)
}
func (m *HandlerMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandlerMember.Merge(m, src)
}
func (m *HandlerMember) XXX_Size() int {
	return xxx_messageInfo_HandlerMember.Size(m)
}
func (m *HandlerMember) XXX_DiscardUnknown() {
	xxx_messageInfo_HandlerMember.DiscardUnknown(m)
}

var xxx_messageInfo_HandlerMember proto.InternalMessageInfo

func (m *HandlerMember) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *HandlerMember) GetMember() *ClusterMember {
	if m != nil {
		return m.Member
	}
	return nil
}

func init() {
	proto.RegisterType((*MembersRequest)(nil), "slipway.MembersRequest")
	proto.RegisterType((*MembersResponse)(nil), "slipway.MembersResponse")
	proto.RegisterType((*ClusterMember)(nil), "slipway.ClusterMember")
	proto.RegisterType((*InitRequest)(nil), "slipway.InitRequest")
	proto.RegisterType((*InitResponse)(nil), "slipway.InitResponse")
	proto.RegisterType((*BootstrapTokenRequest)(nil), "slipway.BootstrapTokenRequest")
	proto.RegisterType((*BootstrapTokenResponse)(nil), "slipway.BootstrapTokenResponse")
	proto.RegisterType((*BootstrapMasterRequest)(nil), "slipway.BootstrapMasterRequest")
	proto.RegisterType((*UploadCertsRequest)(nil), "slipway.UploadCertsRequest")
	proto.RegisterType((*UploadCertsResponse)(nil), "slipway.UploadCertsResponse")
	proto.RegisterType((*HandlerMember)(nil), "slipway.HandlerMember")
}

func init() { proto.RegisterFile("protobuf/slipway/slipway.proto", fileDescriptor_1ddd0b3792332164) }

var fileDescriptor_1ddd0b3792332164 = []byte{
	// 814 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0x96, 0x1d, 0xb0, 0xe3, 0x67, 0xc7, 0x4e, 0xa6, 0x6d, 0xba, 0x18, 0xda, 0x9a, 0x95, 0x10,
	0x11, 0x42, 0xa1, 0x4a, 0x2f, 0x48, 0x70, 0x80, 0xba, 0xa2, 0x94, 0xb4, 0x14, 0x6d, 0xd3, 0x2b,
	0xa3, 0xf1, 0xce, 0x8b, 0x33, 0x64, 0x76, 0x66, 0x99, 0x99, 0x4d, 0x58, 0x0e, 0xfc, 0x31, 0xee,
	0xdc, 0xb8, 0xf2, 0x57, 0xb8, 0xa2, 0x9d, 0x9d, 0xdd, 0xd8, 0x69, 0x02, 0x52, 0x24, 0x7a, 0xb2,
	0xf6, 0xfb, 0xbe, 0x37, 0x33, 0xef, 0xfb, 0x66, 0x9e, 0xe1, 0x7e, 0x6e, 0xb4, 0xd3, 0x8b, 0xe2,
	0xf8, 0x33, 0x2b, 0x45, 0x7e, 0xce, 0xca, 0xe6, 0x77, 0xdf, 0x13, 0xa4, 0x1f, 0x3e, 0xe3, 0x6d,
	0x18, 0xbf, 0xc0, 0x6c, 0x81, 0xc6, 0x26, 0xf8, 0x73, 0x81, 0xd6, 0xc5, 0x16, 0x26, 0x2d, 0x62,
	0x73, 0xad, 0x2c, 0x92, 0x87, 0xd0, 0x3f, 0x61, 0x4a, 0x72, 0x34, 0x51, 0x67, 0xd6, 0xd9, 0x1b,
	0x1e, 0xec, 0xee, 0x37, 0xcb, 0x7d, 0xcb, 0x14, 0x97, 0x68, 0xea, 0x8a, 0xa4, 0x91, 0x91, 0x7d,
	0xe8, 0x19, 0xb4, 0x85, 0x74, 0x51, 0x77, 0xb6, 0xb1, 0x56, 0x30, 0x97, 0x85, 0x75, 0x6d, 0x41,
	0x50, 0xc5, 0x06, 0xb6, 0xd6, 0x08, 0x32, 0x86, 0xae, 0xe0, 0x7e, 0xb7, 0x41, 0xd2, 0x15, 0x9c,
	0x4c, 0x61, 0xf3, 0x44, 0x5b, 0xa7, 0x58, 0x86, 0x51, 0xd7, 0xa3, 0xed, 0x37, 0xd9, 0x85, 0x9e,
	0x64, 0x0b, 0x94, 0x36, 0xda, 0x98, 0x6d, 0xec, 0x0d, 0x92, 0xf0, 0x45, 0xee, 0x01, 0x88, 0x9c,
	0x32, 0xce, 0x0d, 0x5a, 0x1b, 0xbd, 0xe3, 0xab, 0x06, 0x22, 0xff, 0xba, 0x06, 0xe2, 0x3f, 0xbb,
	0x30, 0x7c, 0xa6, 0x84, 0x0b, 0x8d, 0x93, 0xbb, 0xd0, 0xe7, 0xa6, 0xa4, 0xa6, 0x50, 0x7e, 0xdf,
	0xcd, 0xa4, 0xc7, 0x4d, 0x99, 0x14, 0x8a, 0x7c, 0x0a, 0x24, 0xad, 0x0f, 0x47, 0xb9, 0xb2, 0x94,
	0xeb, 0x8c, 0x09, 0x15, 0x4e, 0xb1, 0x1d, 0x98, 0x27, 0xca, 0x3e, 0xf1, 0x38, 0xf9, 0x04, 0x76,
	0x52, 0xad, 0x9c, 0xd1, 0x92, 0xe6, 0x92, 0x29, 0xa4, 0x85, 0x91, 0xd1, 0x86, 0x17, 0x4f, 0x02,
	0xf1, 0x43, 0x85, 0xbf, 0x36, 0x92, 0x3c, 0x80, 0x61, 0xb3, 0x72, 0xaa, 0x44, 0x38, 0x22, 0x04,
	0x68, 0xae, 0x04, 0x79, 0x0f, 0x36, 0x73, 0xcd, 0x69, 0x2a, 0xb8, 0x89, 0xde, 0xf5, 0x6c, 0x3f,
	0xd7, 0x7c, 0x2e, 0xb8, 0x21, 0x1f, 0xc2, 0xc8, 0xa2, 0x39, 0x13, 0x29, 0xd6, 0x74, 0xcf, 0xd3,
	0xc3, 0x80, 0x79, 0xc9, 0x3e, 0xdc, 0xc2, 0x5f, 0x9c, 0x61, 0xf4, 0xb4, 0x58, 0x20, 0xe3, 0x19,
	0x3d, 0x96, 0x6c, 0x69, 0xa3, 0x89, 0x77, 0x69, 0xc7, 0x53, 0x87, 0x35, 0xf3, 0x4d, 0x45, 0x90,
	0x2f, 0x61, 0x5a, 0xb2, 0x4c, 0xd2, 0xf6, 0x4c, 0x5a, 0x1d, 0x8b, 0x65, 0x61, 0x98, 0x13, 0x5a,
	0x45, 0x8f, 0x66, 0x9d, 0xbd, 0x51, 0x12, 0x55, 0x8a, 0x90, 0xd5, 0x7c, 0x95, 0x8f, 0x7f, 0x83,
	0x51, 0x6d, 0xe7, 0x8d, 0x6f, 0xcd, 0x3d, 0x80, 0x13, 0xcf, 0x70, 0xba, 0x28, 0x83, 0xc1, 0x83,
	0x80, 0x3c, 0x2e, 0x2b, 0x5a, 0xea, 0x25, 0xb5, 0xce, 0x20, 0xcb, 0xbc, 0xa5, 0xa3, 0x64, 0x20,
	0xf5, 0xf2, 0x95, 0x07, 0xe2, 0x3f, 0x3a, 0x70, 0xe7, 0xb1, 0xd6, 0xce, 0x3a, 0xc3, 0xf2, 0x23,
	0x7d, 0x8a, 0xea, 0x9a, 0x64, 0x07, 0x6d, 0xb2, 0x33, 0x18, 0x39, 0x91, 0x21, 0x75, 0x9a, 0x4a,
	0x71, 0xd6, 0xdc, 0x2c, 0xa8, 0xb0, 0x23, 0xfd, 0x5c, 0x9c, 0x21, 0x99, 0xc1, 0x90, 0xa3, 0x4d,
	0x8d, 0xc8, 0xbd, 0x07, 0x75, 0x8e, 0xab, 0x10, 0xf9, 0x18, 0x26, 0x4b, 0xa3, 0x8b, 0xdc, 0x52,
	0x7d, 0x86, 0xc6, 0x08, 0x8e, 0x21, 0xc7, 0x71, 0x0d, 0xbf, 0x0c, 0x68, 0x25, 0x2c, 0x2c, 0x5b,
	0xe2, 0x8a, 0xb0, 0x8e, 0x74, 0x5c, 0xc3, 0x8d, 0x30, 0x76, 0xb0, 0x7b, 0xb9, 0x8f, 0x1b, 0x5b,
	0xba, 0x07, 0xdb, 0x86, 0x9d, 0xd3, 0x9f, 0xb4, 0x50, 0x34, 0xd5, 0x59, 0xc6, 0x14, 0x0f, 0x5d,
	0x8e, 0x0d, 0x3b, 0xff, 0x4e, 0x0b, 0x35, 0xaf, 0xd1, 0xf8, 0xaf, 0xce, 0xca, 0xb6, 0x2f, 0x58,
	0x95, 0xef, 0x5b, 0xf1, 0xef, 0x23, 0x18, 0xa3, 0x4a, 0x4d, 0xe9, 0xbf, 0xe8, 0x29, 0x96, 0xc1,
	0xbe, 0xad, 0x0b, 0xf4, 0x10, 0x4b, 0xf2, 0x39, 0x44, 0xf6, 0x54, 0xe4, 0xb4, 0xc8, 0xa5, 0x66,
	0x9c, 0xa6, 0x68, 0x9c, 0x38, 0x16, 0x29, 0x73, 0x68, 0xbd, 0x8d, 0x9b, 0xc9, 0x6e, 0xc5, 0xbf,
	0xf6, 0xf4, 0x7c, 0x85, 0x8d, 0x6f, 0x03, 0xb9, 0x40, 0xdb, 0x31, 0xf7, 0x14, 0x6e, 0xad, 0xa1,
	0x37, 0x75, 0x38, 0xfe, 0x11, 0xb6, 0xd6, 0x98, 0xea, 0x9a, 0x9a, 0x7a, 0x13, 0xda, 0x8e, 0xb0,
	0x41, 0x40, 0x9e, 0xf1, 0x6a, 0x34, 0x66, 0x5e, 0xe8, 0xdd, 0xfa, 0x97, 0xd1, 0x58, 0xab, 0x0e,
	0x0e, 0xa1, 0x1f, 0x08, 0xf2, 0x15, 0x0c, 0x9f, 0x0b, 0xeb, 0xc2, 0x78, 0x26, 0x77, 0xdb, 0xca,
	0xf5, 0x11, 0x3e, 0x8d, 0xde, 0x24, 0xea, 0xf6, 0x0e, 0xfe, 0xee, 0xc0, 0x4e, 0xf5, 0xe4, 0x8d,
	0x42, 0x87, 0xb6, 0x59, 0xf7, 0x0b, 0x80, 0xea, 0xe5, 0x0a, 0x26, 0xc5, 0xaf, 0x48, 0x6e, 0xb7,
	0xd5, 0x2b, 0xd3, 0x71, 0x7a, 0xe7, 0x12, 0x5a, 0x2f, 0xf8, 0xb0, 0x43, 0x8e, 0x60, 0xe7, 0x29,
	0xba, 0xf5, 0x0b, 0x4b, 0xee, 0xb7, 0xea, 0x2b, 0x5f, 0xe4, 0xf4, 0xc1, 0xb5, 0x7c, 0xc8, 0xe1,
	0xe5, 0x6a, 0x68, 0x4d, 0x94, 0xe4, 0xfd, 0xb6, 0xec, 0xcd, 0x44, 0xa7, 0x1f, 0x5c, 0x4d, 0x86,
	0xce, 0x7f, 0xef, 0xc0, 0xf8, 0xa2, 0xf3, 0xef, 0x35, 0xc7, 0xff, 0xe9, 0xe4, 0xaf, 0x60, 0x72,
	0xe9, 0x19, 0x91, 0x2b, 0x6a, 0xd6, 0x1e, 0xd8, 0x7f, 0x2e, 0xba, 0xe8, 0xf9, 0xbf, 0xed, 0x47,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x96, 0x7b, 0xd7, 0xb5, 0xd8, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterClient interface {
	ListMembers(ctx context.Context, in *MembersRequest, opts ...grpc.CallOption) (*MembersResponse, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) ListMembers(ctx context.Context, in *MembersRequest, opts ...grpc.CallOption) (*MembersResponse, error) {
	out := new(MembersResponse)
	err := c.cc.Invoke(ctx, "/slipway.Cluster/ListMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
type ClusterServer interface {
	ListMembers(context.Context, *MembersRequest) (*MembersResponse, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slipway.Cluster/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ListMembers(ctx, req.(*MembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slipway.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMembers",
			Handler:    _Cluster_ListMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/slipway/slipway.proto",
}

// KubernetesClusterClient is the client API for KubernetesCluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KubernetesClusterClient interface {
	Initialize(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (KubernetesCluster_InitializeClient, error)
	GetBootstrapToken(ctx context.Context, in *BootstrapTokenRequest, opts ...grpc.CallOption) (*BootstrapTokenResponse, error)
	UploadCertificates(ctx context.Context, in *UploadCertsRequest, opts ...grpc.CallOption) (*UploadCertsResponse, error)
}

type kubernetesClusterClient struct {
	cc *grpc.ClientConn
}

func NewKubernetesClusterClient(cc *grpc.ClientConn) KubernetesClusterClient {
	return &kubernetesClusterClient{cc}
}

func (c *kubernetesClusterClient) Initialize(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (KubernetesCluster_InitializeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KubernetesCluster_serviceDesc.Streams[0], "/slipway.KubernetesCluster/Initialize", opts...)
	if err != nil {
		return nil, err
	}
	x := &kubernetesClusterInitializeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KubernetesCluster_InitializeClient interface {
	Recv() (*InitResponse, error)
	grpc.ClientStream
}

type kubernetesClusterInitializeClient struct {
	grpc.ClientStream
}

func (x *kubernetesClusterInitializeClient) Recv() (*InitResponse, error) {
	m := new(InitResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kubernetesClusterClient) GetBootstrapToken(ctx context.Context, in *BootstrapTokenRequest, opts ...grpc.CallOption) (*BootstrapTokenResponse, error) {
	out := new(BootstrapTokenResponse)
	err := c.cc.Invoke(ctx, "/slipway.KubernetesCluster/GetBootstrapToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterClient) UploadCertificates(ctx context.Context, in *UploadCertsRequest, opts ...grpc.CallOption) (*UploadCertsResponse, error) {
	out := new(UploadCertsResponse)
	err := c.cc.Invoke(ctx, "/slipway.KubernetesCluster/UploadCertificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesClusterServer is the server API for KubernetesCluster service.
type KubernetesClusterServer interface {
	Initialize(*InitRequest, KubernetesCluster_InitializeServer) error
	GetBootstrapToken(context.Context, *BootstrapTokenRequest) (*BootstrapTokenResponse, error)
	UploadCertificates(context.Context, *UploadCertsRequest) (*UploadCertsResponse, error)
}

func RegisterKubernetesClusterServer(s *grpc.Server, srv KubernetesClusterServer) {
	s.RegisterService(&_KubernetesCluster_serviceDesc, srv)
}

func _KubernetesCluster_Initialize_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InitRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KubernetesClusterServer).Initialize(m, &kubernetesClusterInitializeServer{stream})
}

type KubernetesCluster_InitializeServer interface {
	Send(*InitResponse) error
	grpc.ServerStream
}

type kubernetesClusterInitializeServer struct {
	grpc.ServerStream
}

func (x *kubernetesClusterInitializeServer) Send(m *InitResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KubernetesCluster_GetBootstrapToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterServer).GetBootstrapToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slipway.KubernetesCluster/GetBootstrapToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterServer).GetBootstrapToken(ctx, req.(*BootstrapTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesCluster_UploadCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterServer).UploadCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slipway.KubernetesCluster/UploadCertificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterServer).UploadCertificates(ctx, req.(*UploadCertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KubernetesCluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slipway.KubernetesCluster",
	HandlerType: (*KubernetesClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBootstrapToken",
			Handler:    _KubernetesCluster_GetBootstrapToken_Handler,
		},
		{
			MethodName: "UploadCertificates",
			Handler:    _KubernetesCluster_UploadCertificates_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Initialize",
			Handler:       _KubernetesCluster_Initialize_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/slipway/slipway.proto",
}

// KubernetesNodeClient is the client API for KubernetesNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KubernetesNodeClient interface {
	GetBootstrapToken(ctx context.Context, in *BootstrapTokenRequest, opts ...grpc.CallOption) (*BootstrapTokenResponse, error)
	BootstrapMaster(ctx context.Context, in *BootstrapMasterRequest, opts ...grpc.CallOption) (*BootstrapTokenResponse, error)
}

type kubernetesNodeClient struct {
	cc *grpc.ClientConn
}

func NewKubernetesNodeClient(cc *grpc.ClientConn) KubernetesNodeClient {
	return &kubernetesNodeClient{cc}
}

func (c *kubernetesNodeClient) GetBootstrapToken(ctx context.Context, in *BootstrapTokenRequest, opts ...grpc.CallOption) (*BootstrapTokenResponse, error) {
	out := new(BootstrapTokenResponse)
	err := c.cc.Invoke(ctx, "/slipway.KubernetesNode/GetBootstrapToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesNodeClient) BootstrapMaster(ctx context.Context, in *BootstrapMasterRequest, opts ...grpc.CallOption) (*BootstrapTokenResponse, error) {
	out := new(BootstrapTokenResponse)
	err := c.cc.Invoke(ctx, "/slipway.KubernetesNode/BootstrapMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesNodeServer is the server API for KubernetesNode service.
type KubernetesNodeServer interface {
	GetBootstrapToken(context.Context, *BootstrapTokenRequest) (*BootstrapTokenResponse, error)
	BootstrapMaster(context.Context, *BootstrapMasterRequest) (*BootstrapTokenResponse, error)
}

func RegisterKubernetesNodeServer(s *grpc.Server, srv KubernetesNodeServer) {
	s.RegisterService(&_KubernetesNode_serviceDesc, srv)
}

func _KubernetesNode_GetBootstrapToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesNodeServer).GetBootstrapToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slipway.KubernetesNode/GetBootstrapToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesNodeServer).GetBootstrapToken(ctx, req.(*BootstrapTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesNode_BootstrapMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesNodeServer).BootstrapMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slipway.KubernetesNode/BootstrapMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesNodeServer).BootstrapMaster(ctx, req.(*BootstrapMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KubernetesNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slipway.KubernetesNode",
	HandlerType: (*KubernetesNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBootstrapToken",
			Handler:    _KubernetesNode_GetBootstrapToken_Handler,
		},
		{
			MethodName: "BootstrapMaster",
			Handler:    _KubernetesNode_BootstrapMaster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/slipway/slipway.proto",
}