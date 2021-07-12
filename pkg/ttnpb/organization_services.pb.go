// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/organization_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("lorawan-stack/api/organization_services.proto", fileDescriptor_1a990e3af7846fd3)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/organization_services.proto", fileDescriptor_1a990e3af7846fd3)
}

var fileDescriptor_1a990e3af7846fd3 = []byte{
	// 814 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcf, 0x6f, 0xd3, 0x48,
	0x14, 0xc7, 0x33, 0xfb, 0x23, 0xd5, 0xce, 0x56, 0xad, 0x76, 0xb4, 0xda, 0x95, 0xbc, 0x5d, 0x0b,
	0x19, 0xd4, 0xa6, 0x11, 0x99, 0x81, 0xa6, 0xaa, 0xd4, 0xaa, 0x52, 0xd5, 0x5f, 0x0a, 0xa5, 0x15,
	0x54, 0xa9, 0xb8, 0xf4, 0x52, 0x39, 0xe9, 0xd4, 0xb1, 0x92, 0x7a, 0xcc, 0xcc, 0xa4, 0x25, 0x94,
	0x22, 0xc4, 0x81, 0x2b, 0x08, 0x6e, 0x1c, 0x80, 0x23, 0x88, 0x13, 0xff, 0x01, 0x47, 0x8e, 0x48,
	0x08, 0x89, 0x0b, 0x12, 0x4d, 0x39, 0x70, 0xe4, 0xc8, 0x11, 0x79, 0xec, 0x80, 0x9d, 0xa4, 0x71,
	0xd3, 0xf4, 0x16, 0x8f, 0xdf, 0xbc, 0xef, 0xe7, 0xbd, 0x99, 0xf7, 0x75, 0x60, 0xa6, 0xc2, 0xb8,
	0xb9, 0x6b, 0x3a, 0x19, 0x21, 0xcd, 0x62, 0x99, 0x98, 0xae, 0x4d, 0x18, 0xb7, 0x4c, 0xc7, 0xbe,
	0x69, 0x4a, 0x9b, 0x39, 0x1b, 0x82, 0xf2, 0x1d, 0xbb, 0x48, 0x05, 0x76, 0x39, 0x93, 0x0c, 0x0d,
	0x48, 0xe9, 0xe0, 0x60, 0x0b, 0xde, 0xc9, 0x6a, 0x19, 0xcb, 0x96, 0xa5, 0x6a, 0x01, 0x17, 0xd9,
	0x36, 0xb1, 0x98, 0xc5, 0x88, 0x0a, 0x2b, 0x54, 0xb7, 0xd4, 0x93, 0x7a, 0x50, 0xbf, 0xfc, 0xed,
	0xda, 0x90, 0xc5, 0x98, 0x55, 0xa1, 0x4a, 0xc6, 0x74, 0x1c, 0x26, 0x95, 0x48, 0x90, 0x5c, 0xfb,
	0x2f, 0x78, 0xfb, 0x23, 0x07, 0xdd, 0x76, 0x65, 0x2d, 0x78, 0x79, 0xb6, 0x15, 0xd4, 0xde, 0xa4,
	0x8e, 0xb4, 0xb7, 0x6c, 0xca, 0x1b, 0x19, 0xce, 0x75, 0xae, 0x26, 0x88, 0xd2, 0x5b, 0xa3, 0xb8,
	0x6d, 0x95, 0x64, 0x90, 0x65, 0xec, 0x7d, 0x1f, 0xfc, 0xfb, 0x6a, 0x68, 0x5b, 0x9e, 0x5a, 0xb6,
	0x90, 0xbc, 0x86, 0x1e, 0x02, 0x98, 0x9c, 0xe7, 0xd4, 0x94, 0x14, 0x8d, 0xe2, 0x68, 0x27, 0xb0,
	0xbf, 0x1e, 0xdd, 0x76, 0xbd, 0x4a, 0x85, 0xd4, 0x86, 0x9a, 0x43, 0xc3, 0x41, 0xc6, 0xcc, 0xdd,
	0x77, 0x9f, 0x1f, 0xfd, 0x32, 0x69, 0x8c, 0x93, 0xaa, 0xa0, 0x5c, 0x90, 0xbd, 0x22, 0xab, 0x54,
	0xcc, 0x02, 0xe3, 0xa6, 0x64, 0x1c, 0x7b, 0x6b, 0x1b, 0xf6, 0xa6, 0x68, 0xfc, 0xd8, 0x8f, 0xd4,
	0x23, 0xa6, 0x40, 0x1a, 0xdd, 0x03, 0xf0, 0xd7, 0x1c, 0x95, 0x68, 0xb8, 0x59, 0x26, 0x47, 0x65,
	0xf7, 0x38, 0x93, 0x0a, 0x27, 0x8b, 0x2e, 0x46, 0x85, 0xc8, 0x5e, 0xe4, 0x56, 0x78, 0x44, 0x4d,
	0x0b, 0xfb, 0xe8, 0x29, 0x80, 0xbf, 0xad, 0xd8, 0x42, 0xa2, 0x54, 0xb3, 0x82, 0xb7, 0x1a, 0x56,
	0x11, 0x0d, 0x96, 0xff, 0x3b, 0xb1, 0x08, 0xe3, 0x8a, 0x82, 0xb9, 0x84, 0x06, 0xa2, 0x30, 0xeb,
	0x13, 0xe8, 0x44, 0xdd, 0x42, 0xf7, 0x01, 0x4c, 0x5e, 0x73, 0x37, 0xdb, 0x9e, 0x9f, 0xbf, 0xde,
	0x7d, 0xc3, 0xa6, 0x15, 0xe3, 0x84, 0xd6, 0xb1, 0x61, 0xb8, 0x5d, 0xc3, 0xbc, 0xc3, 0x13, 0x30,
	0xb9, 0x40, 0x2b, 0x54, 0x52, 0x34, 0xd2, 0x49, 0x65, 0xe9, 0xe7, 0x4d, 0xd7, 0xfe, 0xc1, 0xfe,
	0x98, 0xe0, 0xc6, 0x98, 0xe0, 0x45, 0x6f, 0x4c, 0x8c, 0x94, 0x02, 0x31, 0xd2, 0x67, 0x62, 0x4e,
	0x6e, 0x1f, 0xdd, 0x82, 0x7d, 0x79, 0x2a, 0x24, 0xe3, 0xa7, 0xa0, 0x7a, 0x41, 0xa9, 0xa6, 0x8d,
	0x54, 0x9c, 0x2a, 0xe1, 0x81, 0xe4, 0x0d, 0xf8, 0xfb, 0x6a, 0x95, 0x5b, 0xa7, 0xa0, 0x8d, 0x95,
	0x76, 0x2a, 0x3d, 0x1c, 0xab, 0xed, 0x7a, 0x82, 0x63, 0x07, 0x10, 0xa2, 0xb0, 0xc6, 0x6c, 0xb1,
	0x48, 0x85, 0x40, 0xb7, 0x21, 0xf4, 0x2e, 0x68, 0x5e, 0x59, 0x40, 0x37, 0x54, 0x4d, 0x81, 0x7e,
	0x02, 0x83, 0x28, 0xaa, 0x51, 0x34, 0x12, 0xdf, 0x11, 0x5f, 0xf1, 0x09, 0x80, 0xfd, 0xbe, 0x7b,
	0xcc, 0xae, 0x2e, 0x2d, 0xd3, 0x1a, 0x22, 0xf1, 0xde, 0xe2, 0x47, 0x36, 0x6e, 0x68, 0x0b, 0x8a,
	0xff, 0xda, 0x58, 0x54, 0x28, 0x33, 0xc6, 0x54, 0xd7, 0xc3, 0xec, 0x99, 0x62, 0xa6, 0x4c, 0x6b,
	0xca, 0x61, 0x1e, 0x03, 0xf8, 0xa7, 0xd7, 0x21, 0x3f, 0xab, 0x40, 0x38, 0x6e, 0xbe, 0x83, 0xc0,
	0x06, 0xde, 0xbf, 0xed, 0xf1, 0x84, 0x31, 0xa7, 0xf8, 0xa6, 0x51, 0x0f, 0x7c, 0x5e, 0xf7, 0xfe,
	0xc8, 0xd1, 0x80, 0x0d, 0x9d, 0x8f, 0x31, 0xc1, 0xe3, 0xf5, 0x6d, 0x59, 0x71, 0x2d, 0xa2, 0xf9,
	0x93, 0x73, 0x91, 0xbd, 0x32, 0xad, 0xa9, 0x69, 0x7b, 0x09, 0x60, 0xbf, 0x6f, 0x2e, 0x47, 0x1d,
	0x6f, 0xab, 0xf5, 0x1c, 0x0f, 0x33, 0xaf, 0x30, 0x57, 0xb4, 0x5c, 0x2f, 0x98, 0xa6, 0x6b, 0x6f,
	0x94, 0x69, 0x0d, 0x07, 0x86, 0xf4, 0x11, 0xc0, 0xc1, 0x1c, 0x95, 0xf3, 0x21, 0x5b, 0x45, 0x63,
	0x31, 0x4d, 0x0d, 0x07, 0x37, 0x98, 0x47, 0xda, 0xec, 0x89, 0xc6, 0x09, 0x97, 0x39, 0x82, 0x1a,
	0xdb, 0xaa, 0x08, 0x6b, 0x9d, 0xa2, 0x62, 0xf7, 0x65, 0x84, 0xdd, 0x5f, 0x7d, 0x11, 0xe2, 0x3e,
	0x08, 0xe8, 0x05, 0x80, 0x83, 0x6b, 0x71, 0xf5, 0xad, 0xc5, 0xd7, 0x77, 0x94, 0x27, 0x5d, 0x56,
	0xe5, 0x2c, 0x68, 0x33, 0xbd, 0x15, 0xa3, 0xe6, 0xee, 0x15, 0x80, 0x7f, 0x79, 0xa3, 0x15, 0xd6,
	0x17, 0x68, 0x3c, 0x6e, 0xfa, 0x22, 0xe1, 0x47, 0x7e, 0x69, 0x23, 0x51, 0x46, 0x4e, 0x61, 0xcf,
	0xa2, 0x5e, 0xb1, 0xe7, 0xcc, 0x0f, 0x07, 0x7a, 0xe2, 0xdb, 0x81, 0x0e, 0xee, 0xd4, 0x75, 0xf0,
	0xbc, 0xae, 0x83, 0x37, 0x75, 0x1d, 0xbc, 0xad, 0xeb, 0xe0, 0x53, 0x5d, 0x07, 0x5f, 0xea, 0x7a,
	0xe2, 0x6b, 0x5d, 0x07, 0x0f, 0x0e, 0xf5, 0xc4, 0xb3, 0x43, 0x3d, 0xf1, 0xfa, 0x50, 0x07, 0xeb,
	0xc4, 0x62, 0x58, 0x96, 0xa8, 0x2c, 0xd9, 0x8e, 0x25, 0xb0, 0x43, 0xe5, 0x2e, 0xe3, 0x65, 0x12,
	0xfd, 0x8b, 0xb6, 0x93, 0x25, 0x6e, 0xd9, 0x22, 0x52, 0x3a, 0x6e, 0xa1, 0x90, 0x54, 0x2d, 0xcf,
	0x7e, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x44, 0x56, 0x4c, 0x6a, 0xbb, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrganizationRegistryClient is the client API for OrganizationRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrganizationRegistryClient interface {
	// Create a new organization. This also sets the given user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// Get the organization with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// List organizations where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the organizations the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*Organizations, error)
	// Update the organization, changing the fields specified by the field mask to the provided values.
	Update(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error)
	// Delete the organization. This may not release the organization ID for reuse.
	Delete(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	// Restore a recently deleted organization.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	// Purge the organization. This will release the organization ID for reuse.
	// The user is responsible for clearing data from any (external) integrations
	// that may store and expose data by user or organization ID.
	Purge(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type organizationRegistryClient struct {
	cc *grpc.ClientConn
}

func NewOrganizationRegistryClient(cc *grpc.ClientConn) OrganizationRegistryClient {
	return &organizationRegistryClient{cc}
}

func (c *organizationRegistryClient) Create(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Get(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) List(ctx context.Context, in *ListOrganizationsRequest, opts ...grpc.CallOption) (*Organizations, error) {
	out := new(Organizations)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Update(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*Organization, error) {
	out := new(Organization)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Delete(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Restore(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationRegistryClient) Purge(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationRegistry/Purge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationRegistryServer is the server API for OrganizationRegistry service.
type OrganizationRegistryServer interface {
	// Create a new organization. This also sets the given user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateOrganizationRequest) (*Organization, error)
	// Get the organization with the given identifiers, selecting the fields specified
	// in the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	Get(context.Context, *GetOrganizationRequest) (*Organization, error)
	// List organizations where the given user or organization is a direct collaborator.
	// If no user or organization is given, this returns the organizations the caller
	// has access to.
	// Similar to Get, this selects the fields given by the field mask.
	// More or less fields may be returned, depending on the rights of the caller.
	List(context.Context, *ListOrganizationsRequest) (*Organizations, error)
	// Update the organization, changing the fields specified by the field mask to the provided values.
	Update(context.Context, *UpdateOrganizationRequest) (*Organization, error)
	// Delete the organization. This may not release the organization ID for reuse.
	Delete(context.Context, *OrganizationIdentifiers) (*types.Empty, error)
	// Restore a recently deleted organization.
	//
	// Deployment configuration may specify if, and for how long after deletion,
	// entities can be restored.
	Restore(context.Context, *OrganizationIdentifiers) (*types.Empty, error)
	// Purge the organization. This will release the organization ID for reuse.
	// The user is responsible for clearing data from any (external) integrations
	// that may store and expose data by user or organization ID.
	Purge(context.Context, *OrganizationIdentifiers) (*types.Empty, error)
}

// UnimplementedOrganizationRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedOrganizationRegistryServer struct {
}

func (*UnimplementedOrganizationRegistryServer) Create(ctx context.Context, req *CreateOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Get(ctx context.Context, req *GetOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedOrganizationRegistryServer) List(ctx context.Context, req *ListOrganizationsRequest) (*Organizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Update(ctx context.Context, req *UpdateOrganizationRequest) (*Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Delete(ctx context.Context, req *OrganizationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Restore(ctx context.Context, req *OrganizationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (*UnimplementedOrganizationRegistryServer) Purge(ctx context.Context, req *OrganizationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}

func RegisterOrganizationRegistryServer(s *grpc.Server, srv OrganizationRegistryServer) {
	s.RegisterService(&_OrganizationRegistry_serviceDesc, srv)
}

func _OrganizationRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Create(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Get(ctx, req.(*GetOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).List(ctx, req.(*ListOrganizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Update(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Delete(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Restore(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationRegistry_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationRegistryServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationRegistry/Purge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationRegistryServer).Purge(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrganizationRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.OrganizationRegistry",
	HandlerType: (*OrganizationRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrganizationRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _OrganizationRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OrganizationRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrganizationRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrganizationRegistry_Delete_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _OrganizationRegistry_Restore_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _OrganizationRegistry_Purge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/organization_services.proto",
}

// OrganizationAccessClient is the client API for OrganizationAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrganizationAccessClient interface {
	// List the rights the caller has on this organization.
	ListRights(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	// Create an API key scoped to this organization.
	// Organization API keys can give access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	CreateAPIKey(ctx context.Context, in *CreateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// List the API keys for this organization.
	ListAPIKeys(ctx context.Context, in *ListOrganizationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error)
	// Get a single API key of this organization.
	GetAPIKey(ctx context.Context, in *GetOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Update the rights of an API key of the organization.
	// This method can also be used to delete the API key, by giving it no rights.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(ctx context.Context, in *UpdateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Get the rights of a collaborator (member) of the organization.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the organization.
	// Organization collaborators can get access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(ctx context.Context, in *SetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error)
	// List the collaborators on this organization.
	ListCollaborators(ctx context.Context, in *ListOrganizationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
}

type organizationAccessClient struct {
	cc *grpc.ClientConn
}

func NewOrganizationAccessClient(cc *grpc.ClientConn) OrganizationAccessClient {
	return &organizationAccessClient{cc}
}

func (c *organizationAccessClient) ListRights(ctx context.Context, in *OrganizationIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) CreateAPIKey(ctx context.Context, in *CreateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) ListAPIKeys(ctx context.Context, in *ListOrganizationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) GetAPIKey(ctx context.Context, in *GetOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/GetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateOrganizationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) GetCollaborator(ctx context.Context, in *GetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/GetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) SetCollaborator(ctx context.Context, in *SetOrganizationCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAccessClient) ListCollaborators(ctx context.Context, in *ListOrganizationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OrganizationAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationAccessServer is the server API for OrganizationAccess service.
type OrganizationAccessServer interface {
	// List the rights the caller has on this organization.
	ListRights(context.Context, *OrganizationIdentifiers) (*Rights, error)
	// Create an API key scoped to this organization.
	// Organization API keys can give access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	CreateAPIKey(context.Context, *CreateOrganizationAPIKeyRequest) (*APIKey, error)
	// List the API keys for this organization.
	ListAPIKeys(context.Context, *ListOrganizationAPIKeysRequest) (*APIKeys, error)
	// Get a single API key of this organization.
	GetAPIKey(context.Context, *GetOrganizationAPIKeyRequest) (*APIKey, error)
	// Update the rights of an API key of the organization.
	// This method can also be used to delete the API key, by giving it no rights.
	// The caller is required to have all assigned or/and removed rights.
	UpdateAPIKey(context.Context, *UpdateOrganizationAPIKeyRequest) (*APIKey, error)
	// Get the rights of a collaborator (member) of the organization.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetOrganizationCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the organization.
	// Organization collaborators can get access to the organization itself, as well as
	// any application, gateway and OAuth client this organization is a collaborator of.
	// This method can also be used to delete the collaborator, by giving them no rights.
	// The caller is required to have all assigned or/and removed rights.
	SetCollaborator(context.Context, *SetOrganizationCollaboratorRequest) (*types.Empty, error)
	// List the collaborators on this organization.
	ListCollaborators(context.Context, *ListOrganizationCollaboratorsRequest) (*Collaborators, error)
}

// UnimplementedOrganizationAccessServer can be embedded to have forward compatible implementations.
type UnimplementedOrganizationAccessServer struct {
}

func (*UnimplementedOrganizationAccessServer) ListRights(ctx context.Context, req *OrganizationIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (*UnimplementedOrganizationAccessServer) CreateAPIKey(ctx context.Context, req *CreateOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIKey not implemented")
}
func (*UnimplementedOrganizationAccessServer) ListAPIKeys(ctx context.Context, req *ListOrganizationAPIKeysRequest) (*APIKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAPIKeys not implemented")
}
func (*UnimplementedOrganizationAccessServer) GetAPIKey(ctx context.Context, req *GetOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIKey not implemented")
}
func (*UnimplementedOrganizationAccessServer) UpdateAPIKey(ctx context.Context, req *UpdateOrganizationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPIKey not implemented")
}
func (*UnimplementedOrganizationAccessServer) GetCollaborator(ctx context.Context, req *GetOrganizationCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (*UnimplementedOrganizationAccessServer) SetCollaborator(ctx context.Context, req *SetOrganizationCollaboratorRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (*UnimplementedOrganizationAccessServer) ListCollaborators(ctx context.Context, req *ListOrganizationCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}

func RegisterOrganizationAccessServer(s *grpc.Server, srv OrganizationAccessServer) {
	s.RegisterService(&_OrganizationAccess_serviceDesc, srv)
}

func _OrganizationAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListRights(ctx, req.(*OrganizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).CreateAPIKey(ctx, req.(*CreateOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationAPIKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListAPIKeys(ctx, req.(*ListOrganizationAPIKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/GetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).GetAPIKey(ctx, req.(*GetOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).UpdateAPIKey(ctx, req.(*UpdateOrganizationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/GetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).GetCollaborator(ctx, req.(*GetOrganizationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOrganizationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).SetCollaborator(ctx, req.(*SetOrganizationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OrganizationAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAccessServer).ListCollaborators(ctx, req.(*ListOrganizationCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrganizationAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.OrganizationAccess",
	HandlerType: (*OrganizationAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _OrganizationAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _OrganizationAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _OrganizationAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "GetAPIKey",
			Handler:    _OrganizationAccess_GetAPIKey_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _OrganizationAccess_UpdateAPIKey_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _OrganizationAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _OrganizationAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _OrganizationAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/organization_services.proto",
}
