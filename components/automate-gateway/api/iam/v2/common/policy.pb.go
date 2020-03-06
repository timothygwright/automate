// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/common/policy.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Type int32

const (
	Type_CHEF_MANAGED Type = 0
	Type_CUSTOM       Type = 1
)

var Type_name = map[int32]string{
	0: "CHEF_MANAGED",
	1: "CUSTOM",
}

var Type_value = map[string]int32{
	"CHEF_MANAGED": 0,
	"CUSTOM":       1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{0}
}

type Flag int32

const (
	Flag_VERSION_2_0 Flag = 0
	Flag_VERSION_2_1 Flag = 1
)

var Flag_name = map[int32]string{
	0: "VERSION_2_0",
	1: "VERSION_2_1",
}

var Flag_value = map[string]int32{
	"VERSION_2_0": 0,
	"VERSION_2_1": 1,
}

func (x Flag) String() string {
	return proto.EnumName(Flag_name, int32(x))
}

func (Flag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{1}
}

type Statement_Effect int32

const (
	Statement_ALLOW Statement_Effect = 0
	Statement_DENY  Statement_Effect = 1
)

var Statement_Effect_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var Statement_Effect_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x Statement_Effect) String() string {
	return proto.EnumName(Statement_Effect_name, int32(x))
}

func (Statement_Effect) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{1, 0}
}

type Version_VersionNumber int32

const (
	Version_V0 Version_VersionNumber = 0
	Version_V1 Version_VersionNumber = 1
	Version_V2 Version_VersionNumber = 2
)

var Version_VersionNumber_name = map[int32]string{
	0: "V0",
	1: "V1",
	2: "V2",
}

var Version_VersionNumber_value = map[string]int32{
	"V0": 0,
	"V1": 1,
	"V2": 2,
}

func (x Version_VersionNumber) String() string {
	return proto.EnumName(Version_VersionNumber_name, int32(x))
}

func (Version_VersionNumber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{4, 0}
}

type Policy struct {
	// Name for the policy.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique, user-specified ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// This doc-comment is ignored for an enum.
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// Members affected by this policy.
	Members []string `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	// Statements for the policy. Must contain one or more.
	Statements []*Statement `protobuf:"bytes,5,rep,name=statements,proto3" json:"statements,omitempty"`
	// The list of projects this policy belongs to.
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{0}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Policy) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Policy) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Policy) GetStatements() []*Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *Policy) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Statement struct {
	// This doc-comment is ignored for an enum.
	Effect Statement_Effect `protobuf:"varint,1,opt,name=effect,proto3,enum=chef.automate.api.iam.v2.Statement_Effect" json:"effect,omitempty"`
	// Defines the actions for the members of this policy. The best practice is to use chef-managed roles or to define custom roles for unique action sets.
	Actions []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	// The role defines a set of actions that the statement is scoped to.
	Role string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// DEPRECATED: Resources defined inline. Use projects instead.
	Resources []string `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty"`
	// The project list defines the set of resources that the statement is scoped to.
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{1}
}

func (m *Statement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement.Unmarshal(m, b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
}
func (m *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(m, src)
}
func (m *Statement) XXX_Size() int {
	return xxx_messageInfo_Statement.Size(m)
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetEffect() Statement_Effect {
	if m != nil {
		return m.Effect
	}
	return Statement_ALLOW
}

func (m *Statement) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Statement) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Statement) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Statement) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Role struct {
	// Name for the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique, user-specified ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Whether this policy is user created or chef managed.
	// One of `CUSTOM` or `CHEF_MANAGED`, respectively.
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// List of actions that this role scopes to.
	Actions []string `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	// The list of projects this role belongs to.
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{2}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Role) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Role) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Project struct {
	// Name for the project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique, user-specified ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Whether this policy is user created or chef managed.
	// One of `CUSTOM` or `CHEF_MANAGED`, respectively.
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// The current status of the rules for this project.
	Status               ProjectRulesStatus `protobuf:"varint,4,opt,name=status,proto3,enum=chef.automate.api.iam.v2.ProjectRulesStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{3}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Project) GetStatus() ProjectRulesStatus {
	if m != nil {
		return m.Status
	}
	return ProjectRulesStatus_PROJECT_RULES_STATUS_UNSET
}

type Version struct {
	Major                Version_VersionNumber `protobuf:"varint,1,opt,name=major,proto3,enum=chef.automate.api.iam.v2.Version_VersionNumber" json:"major,omitempty"`
	Minor                Version_VersionNumber `protobuf:"varint,2,opt,name=minor,proto3,enum=chef.automate.api.iam.v2.Version_VersionNumber" json:"minor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{4}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() Version_VersionNumber {
	if m != nil {
		return m.Major
	}
	return Version_V0
}

func (m *Version) GetMinor() Version_VersionNumber {
	if m != nil {
		return m.Minor
	}
	return Version_V0
}

func init() {
	proto.RegisterEnum("chef.automate.api.iam.v2.Type", Type_name, Type_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.Flag", Flag_name, Flag_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.Statement_Effect", Statement_Effect_name, Statement_Effect_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.Version_VersionNumber", Version_VersionNumber_name, Version_VersionNumber_value)
	proto.RegisterType((*Policy)(nil), "chef.automate.api.iam.v2.Policy")
	proto.RegisterType((*Statement)(nil), "chef.automate.api.iam.v2.Statement")
	proto.RegisterType((*Role)(nil), "chef.automate.api.iam.v2.Role")
	proto.RegisterType((*Project)(nil), "chef.automate.api.iam.v2.Project")
	proto.RegisterType((*Version)(nil), "chef.automate.api.iam.v2.Version")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/common/policy.proto", fileDescriptor_96f942825c854ad5)
}

var fileDescriptor_96f942825c854ad5 = []byte{
	// 867 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5f, 0x6b, 0xec, 0x44,
	0x14, 0x6f, 0xb2, 0x69, 0x7a, 0x77, 0xae, 0xd6, 0x30, 0x4f, 0xa1, 0xa8, 0x84, 0x51, 0x70, 0x29,
	0xd9, 0xc4, 0xc6, 0x27, 0xf3, 0xb6, 0xb7, 0xdd, 0x7a, 0x85, 0xdb, 0xf6, 0x92, 0xd6, 0x55, 0x6f,
	0x09, 0x65, 0x36, 0x3b, 0xcd, 0xcd, 0x25, 0x93, 0x89, 0xc9, 0xa4, 0x65, 0x59, 0x16, 0x7c, 0xd2,
	0xf7, 0xe2, 0x47, 0x10, 0x04, 0x1f, 0xfc, 0x0c, 0x7e, 0x12, 0x3f, 0x86, 0x20, 0xf8, 0x20, 0x33,
	0x49, 0xf6, 0x1f, 0xac, 0x78, 0x85, 0x3e, 0xcd, 0xcc, 0x2f, 0xe7, 0x77, 0xe6, 0x9c, 0xf3, 0x3b,
	0x39, 0x03, 0xfc, 0x88, 0xd1, 0x9c, 0x65, 0x24, 0xe3, 0xa5, 0x8b, 0x2b, 0xce, 0x28, 0xe6, 0xa4,
	0x1f, 0x63, 0x4e, 0xee, 0xf1, 0xd4, 0xc5, 0x79, 0xe2, 0x26, 0x98, 0xba, 0x77, 0x9e, 0x1b, 0x31,
	0x4a, 0x59, 0xe6, 0xe6, 0x2c, 0x4d, 0xa2, 0xa9, 0x93, 0x17, 0x8c, 0x33, 0x68, 0x46, 0xaf, 0xc9,
	0xad, 0xd3, 0xb2, 0x1c, 0x9c, 0x27, 0x4e, 0x82, 0xa9, 0x73, 0xe7, 0x1d, 0x7c, 0xfe, 0x76, 0x5e,
	0x8b, 0x2a, 0x25, 0x65, 0xed, 0xf4, 0xc0, 0x96, 0x4b, 0xd4, 0x8f, 0x49, 0xd6, 0x2f, 0xef, 0x71,
	0x1c, 0x93, 0xc2, 0x65, 0x39, 0x4f, 0x58, 0x56, 0xba, 0x38, 0xcb, 0x18, 0xc7, 0x72, 0x5f, 0x5b,
	0xa3, 0x3f, 0x3a, 0x40, 0x7f, 0x29, 0x63, 0x82, 0x10, 0x68, 0x19, 0xa6, 0xc4, 0x54, 0x2c, 0xa5,
	0xd7, 0x0d, 0xe4, 0x1e, 0xee, 0x03, 0x35, 0x99, 0x98, 0xaa, 0x44, 0xd4, 0x64, 0x02, 0x3d, 0xa0,
	0xf1, 0x69, 0x4e, 0xcc, 0x8e, 0xa5, 0xf4, 0xf6, 0xbd, 0x0f, 0x9d, 0x6d, 0x09, 0x38, 0x57, 0xd3,
	0x9c, 0x04, 0xd2, 0x16, 0x9a, 0x60, 0x8f, 0x12, 0x3a, 0x26, 0x45, 0x69, 0x6a, 0x56, 0xa7, 0xd7,
	0x0d, 0xda, 0x23, 0x3c, 0x06, 0xa0, 0xe4, 0x98, 0x13, 0x2a, 0xf2, 0x34, 0x77, 0xad, 0x4e, 0xef,
	0xa9, 0xf7, 0xd1, 0x76, 0x9f, 0x97, 0xad, 0x6d, 0xb0, 0x42, 0x83, 0x07, 0xe0, 0x49, 0x5e, 0xb0,
	0x37, 0x24, 0xe2, 0xa5, 0xa9, 0x4b, 0xff, 0x8b, 0xb3, 0xff, 0xbd, 0xfa, 0x30, 0xf8, 0x5b, 0xf1,
	0xfe, 0x52, 0xe0, 0x9f, 0xca, 0x0c, 0x89, 0x8c, 0x90, 0x6f, 0xa1, 0xb3, 0xa9, 0x35, 0x4a, 0xc8,
	0x3d, 0x29, 0xac, 0x3a, 0x73, 0x64, 0xa3, 0x64, 0x22, 0x3e, 0xdc, 0x49, 0xb4, 0x9f, 0xb7, 0xa8,
	0x08, 0x1f, 0xf9, 0xd6, 0x35, 0x3a, 0xfe, 0xea, 0xf2, 0xea, 0xe2, 0x0c, 0xd9, 0x16, 0x3a, 0x7e,
	0x3e, 0x3c, 0xbd, 0x39, 0x1b, 0x9c, 0x0f, 0xbe, 0x18, 0x9e, 0xa0, 0xd0, 0x46, 0x4d, 0x2a, 0xd2,
	0x8a, 0x13, 0x4c, 0xfd, 0x74, 0x82, 0x73, 0xff, 0x50, 0x7c, 0x5a, 0xc4, 0x28, 0x3e, 0xce, 0x50,
	0xc1, 0x52, 0xb2, 0xbc, 0x04, 0xd9, 0xa8, 0x8d, 0x12, 0xf9, 0xd7, 0xed, 0xfe, 0x48, 0x5c, 0xd1,
	0xec, 0x3d, 0x14, 0xce, 0xed, 0x25, 0xed, 0x3b, 0xbc, 0x46, 0xb1, 0xae, 0x11, 0x8e, 0x22, 0x92,
	0x73, 0x9c, 0x45, 0xe4, 0x65, 0x0d, 0xa3, 0x70, 0x1e, 0xae, 0xdb, 0x84, 0x73, 0xf4, 0xa3, 0x06,
	0xba, 0x8b, 0xc2, 0xc1, 0x67, 0x40, 0x27, 0xb7, 0xb7, 0x24, 0xe2, 0x52, 0xe5, 0x7d, 0xef, 0xf0,
	0x3f, 0x54, 0xdb, 0x19, 0x4a, 0x46, 0xd0, 0x30, 0x85, 0x9e, 0x38, 0x92, 0x3d, 0x64, 0x76, 0x6a,
	0x3d, 0x9b, 0xa3, 0xe8, 0x20, 0x11, 0xae, 0xa9, 0xd5, 0x1d, 0x24, 0xf6, 0xf0, 0x7d, 0xd0, 0x2d,
	0x48, 0xc9, 0xaa, 0x22, 0x22, 0xb5, 0xc4, 0xdd, 0x60, 0x09, 0xfc, 0x9b, 0x78, 0xe8, 0x03, 0xa0,
	0xd7, 0x37, 0xc3, 0x2e, 0xd8, 0x1d, 0xbc, 0x78, 0x71, 0xf1, 0xb5, 0xb1, 0x03, 0x9f, 0x00, 0xed,
	0x64, 0x78, 0xfe, 0xad, 0xa1, 0xf8, 0xbf, 0xa9, 0x0f, 0x83, 0x5f, 0x55, 0xef, 0x17, 0x15, 0xfe,
	0xac, 0xce, 0x1a, 0x01, 0x85, 0x02, 0x7d, 0x8a, 0x33, 0x1c, 0x93, 0xa2, 0xec, 0x4f, 0xc8, 0x1d,
	0xcb, 0x4b, 0xb4, 0x22, 0xd2, 0x35, 0xaa, 0x4a, 0x52, 0xf8, 0x29, 0x8b, 0x70, 0xea, 0x8f, 0xd9,
	0x18, 0xd9, 0xab, 0x40, 0x8c, 0x29, 0xc5, 0x42, 0xb9, 0xb6, 0x51, 0xae, 0x08, 0xa6, 0xd6, 0x89,
	0xf4, 0x62, 0x9d, 0x35, 0x5e, 0x37, 0x55, 0x68, 0x0e, 0xed, 0x65, 0xa1, 0x6d, 0x2d, 0x85, 0x17,
	0x16, 0x33, 0x54, 0x97, 0x4d, 0xf8, 0x93, 0x49, 0x20, 0x1b, 0x35, 0xf5, 0x92, 0xfc, 0x04, 0x53,
	0x5f, 0xc4, 0x50, 0xfa, 0x55, 0x3e, 0xc1, 0x9c, 0x88, 0x66, 0x5c, 0x40, 0x69, 0x52, 0xf2, 0x35,
	0x20, 0x26, 0xed, 0x59, 0xa4, 0xba, 0xc1, 0xa9, 0xa1, 0x15, 0x4e, 0x0d, 0x08, 0x4e, 0x38, 0x9f,
	0xa3, 0x9f, 0x14, 0xa0, 0x05, 0x42, 0x92, 0x47, 0xfc, 0xd1, 0xdb, 0xc6, 0xd0, 0xd6, 0x1b, 0x63,
	0x55, 0xe6, 0xdd, 0x0d, 0x99, 0x7f, 0x50, 0xc1, 0x5e, 0xd3, 0xc0, 0x8f, 0x16, 0xd9, 0x09, 0xd0,
	0x85, 0x28, 0x55, 0x29, 0x5b, 0x73, 0xdf, 0xb3, 0xb7, 0xb3, 0x9a, 0x50, 0x02, 0x31, 0x51, 0x2f,
	0x25, 0x27, 0x68, 0xb8, 0x7e, 0xf8, 0x30, 0x78, 0xe5, 0x7d, 0x03, 0x47, 0xcb, 0x49, 0xb2, 0x21,
	0x7e, 0x3b, 0x47, 0x36, 0x61, 0xab, 0x99, 0x24, 0xeb, 0xd3, 0xa3, 0xe9, 0x94, 0xaa, 0xfd, 0x53,
	0x7f, 0x57, 0xc0, 0xde, 0x88, 0x14, 0x65, 0xc2, 0x32, 0x38, 0x04, 0xbb, 0x14, 0xbf, 0x61, 0x45,
	0xf3, 0x9b, 0xba, 0xdb, 0xe3, 0x6d, 0x18, 0xed, 0x7a, 0x5e, 0x89, 0x56, 0x0f, 0x6a, 0xb6, 0x74,
	0x93, 0x64, 0xac, 0x90, 0xe5, 0xfb, 0x5f, 0x6e, 0x04, 0x1b, 0x7d, 0x02, 0xde, 0x5d, 0xc3, 0xa1,
	0x0e, 0xd4, 0xd1, 0xa7, 0xc6, 0x8e, 0x5c, 0x8f, 0x0c, 0x45, 0xae, 0x9e, 0xa1, 0x1e, 0x7e, 0x0c,
	0x34, 0x51, 0x75, 0x68, 0x80, 0x77, 0x56, 0x53, 0x35, 0x76, 0x20, 0x00, 0x7a, 0x3d, 0x44, 0x0d,
	0xe5, 0xb0, 0x07, 0xb4, 0xd3, 0x14, 0xc7, 0xf0, 0x3d, 0xf0, 0x74, 0x34, 0x0c, 0x2e, 0xbf, 0xbc,
	0x38, 0xbf, 0xf1, 0x6e, 0x84, 0xbb, 0x35, 0xe0, 0xc8, 0x50, 0x9e, 0x3d, 0x7f, 0x75, 0x1a, 0x27,
	0xfc, 0x75, 0x35, 0x76, 0x22, 0x46, 0x5d, 0x11, 0xfc, 0xe2, 0x35, 0x74, 0xdf, 0xea, 0x85, 0x1c,
	0xeb, 0xf2, 0xb9, 0xfb, 0xec, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xb6, 0xf5, 0x2f, 0xaf,
	0x07, 0x00, 0x00,
}
