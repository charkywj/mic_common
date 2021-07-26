// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package BankService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BankServiceClient is the client API for BankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankServiceClient interface {
	//银行账号相关
	ListBankAccount(ctx context.Context, in *ListBankAccountReq, opts ...grpc.CallOption) (*ListBankAccountResp, error)
	CreateOneAccountInfo(ctx context.Context, in *CreateOneAccountInfoReq, opts ...grpc.CallOption) (*CreateOneAccountInfoResp, error)
	UpdateBankAccount(ctx context.Context, in *UpdateBankAccountReq, opts ...grpc.CallOption) (*UpdateBankAccountResp, error)
	UpdateBankAccountList(ctx context.Context, in *UpdateBankAccountListReq, opts ...grpc.CallOption) (*UpdateBankAccountListResp, error)
	//配送企业授信额度设置
	GetCompanyCreditConfigs(ctx context.Context, in *GetCompanyCreditConfigsReq, opts ...grpc.CallOption) (*GetCompanyCreditConfigsResp, error)
	UpdateCompanyCreditConfig(ctx context.Context, in *UpdateCompanyCreditConfigReq, opts ...grpc.CallOption) (*UpdateCompanyCreditConfigResp, error)
	//企业账户操作相关
	GetCompanyBankAccountDetail(ctx context.Context, in *GetCompanyBankAccountDetailReq, opts ...grpc.CallOption) (*GetCompanyBankAccountDetailResp, error)
	CompanyWithdraw(ctx context.Context, in *CompanyWithdrawReq, opts ...grpc.CallOption) (*CompanyWithdrawResp, error)
	GetWithdrawReceipt(ctx context.Context, in *GetWithdrawReceiptReq, opts ...grpc.CallOption) (*GetWithdrawReceiptResp, error)
	ListCompanyWithdraw(ctx context.Context, in *ListCompanyWithdrawReq, opts ...grpc.CallOption) (*ListCompanyWithdrawResp, error)
}

type bankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankServiceClient(cc grpc.ClientConnInterface) BankServiceClient {
	return &bankServiceClient{cc}
}

func (c *bankServiceClient) ListBankAccount(ctx context.Context, in *ListBankAccountReq, opts ...grpc.CallOption) (*ListBankAccountResp, error) {
	out := new(ListBankAccountResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/ListBankAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) CreateOneAccountInfo(ctx context.Context, in *CreateOneAccountInfoReq, opts ...grpc.CallOption) (*CreateOneAccountInfoResp, error) {
	out := new(CreateOneAccountInfoResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/CreateOneAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) UpdateBankAccount(ctx context.Context, in *UpdateBankAccountReq, opts ...grpc.CallOption) (*UpdateBankAccountResp, error) {
	out := new(UpdateBankAccountResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/UpdateBankAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) UpdateBankAccountList(ctx context.Context, in *UpdateBankAccountListReq, opts ...grpc.CallOption) (*UpdateBankAccountListResp, error) {
	out := new(UpdateBankAccountListResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/UpdateBankAccountList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) GetCompanyCreditConfigs(ctx context.Context, in *GetCompanyCreditConfigsReq, opts ...grpc.CallOption) (*GetCompanyCreditConfigsResp, error) {
	out := new(GetCompanyCreditConfigsResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/GetCompanyCreditConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) UpdateCompanyCreditConfig(ctx context.Context, in *UpdateCompanyCreditConfigReq, opts ...grpc.CallOption) (*UpdateCompanyCreditConfigResp, error) {
	out := new(UpdateCompanyCreditConfigResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/UpdateCompanyCreditConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) GetCompanyBankAccountDetail(ctx context.Context, in *GetCompanyBankAccountDetailReq, opts ...grpc.CallOption) (*GetCompanyBankAccountDetailResp, error) {
	out := new(GetCompanyBankAccountDetailResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/GetCompanyBankAccountDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) CompanyWithdraw(ctx context.Context, in *CompanyWithdrawReq, opts ...grpc.CallOption) (*CompanyWithdrawResp, error) {
	out := new(CompanyWithdrawResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/CompanyWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) GetWithdrawReceipt(ctx context.Context, in *GetWithdrawReceiptReq, opts ...grpc.CallOption) (*GetWithdrawReceiptResp, error) {
	out := new(GetWithdrawReceiptResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/GetWithdrawReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) ListCompanyWithdraw(ctx context.Context, in *ListCompanyWithdrawReq, opts ...grpc.CallOption) (*ListCompanyWithdrawResp, error) {
	out := new(ListCompanyWithdrawResp)
	err := c.cc.Invoke(ctx, "/BankService.BankService/ListCompanyWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServiceServer is the server API for BankService service.
// All implementations must embed UnimplementedBankServiceServer
// for forward compatibility
type BankServiceServer interface {
	//银行账号相关
	ListBankAccount(context.Context, *ListBankAccountReq) (*ListBankAccountResp, error)
	CreateOneAccountInfo(context.Context, *CreateOneAccountInfoReq) (*CreateOneAccountInfoResp, error)
	UpdateBankAccount(context.Context, *UpdateBankAccountReq) (*UpdateBankAccountResp, error)
	UpdateBankAccountList(context.Context, *UpdateBankAccountListReq) (*UpdateBankAccountListResp, error)
	//配送企业授信额度设置
	GetCompanyCreditConfigs(context.Context, *GetCompanyCreditConfigsReq) (*GetCompanyCreditConfigsResp, error)
	UpdateCompanyCreditConfig(context.Context, *UpdateCompanyCreditConfigReq) (*UpdateCompanyCreditConfigResp, error)
	//企业账户操作相关
	GetCompanyBankAccountDetail(context.Context, *GetCompanyBankAccountDetailReq) (*GetCompanyBankAccountDetailResp, error)
	CompanyWithdraw(context.Context, *CompanyWithdrawReq) (*CompanyWithdrawResp, error)
	GetWithdrawReceipt(context.Context, *GetWithdrawReceiptReq) (*GetWithdrawReceiptResp, error)
	ListCompanyWithdraw(context.Context, *ListCompanyWithdrawReq) (*ListCompanyWithdrawResp, error)
	mustEmbedUnimplementedBankServiceServer()
}

// UnimplementedBankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankServiceServer struct {
}

func (UnimplementedBankServiceServer) ListBankAccount(context.Context, *ListBankAccountReq) (*ListBankAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBankAccount not implemented")
}
func (UnimplementedBankServiceServer) CreateOneAccountInfo(context.Context, *CreateOneAccountInfoReq) (*CreateOneAccountInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOneAccountInfo not implemented")
}
func (UnimplementedBankServiceServer) UpdateBankAccount(context.Context, *UpdateBankAccountReq) (*UpdateBankAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBankAccount not implemented")
}
func (UnimplementedBankServiceServer) UpdateBankAccountList(context.Context, *UpdateBankAccountListReq) (*UpdateBankAccountListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBankAccountList not implemented")
}
func (UnimplementedBankServiceServer) GetCompanyCreditConfigs(context.Context, *GetCompanyCreditConfigsReq) (*GetCompanyCreditConfigsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyCreditConfigs not implemented")
}
func (UnimplementedBankServiceServer) UpdateCompanyCreditConfig(context.Context, *UpdateCompanyCreditConfigReq) (*UpdateCompanyCreditConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompanyCreditConfig not implemented")
}
func (UnimplementedBankServiceServer) GetCompanyBankAccountDetail(context.Context, *GetCompanyBankAccountDetailReq) (*GetCompanyBankAccountDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyBankAccountDetail not implemented")
}
func (UnimplementedBankServiceServer) CompanyWithdraw(context.Context, *CompanyWithdrawReq) (*CompanyWithdrawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompanyWithdraw not implemented")
}
func (UnimplementedBankServiceServer) GetWithdrawReceipt(context.Context, *GetWithdrawReceiptReq) (*GetWithdrawReceiptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawReceipt not implemented")
}
func (UnimplementedBankServiceServer) ListCompanyWithdraw(context.Context, *ListCompanyWithdrawReq) (*ListCompanyWithdrawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanyWithdraw not implemented")
}
func (UnimplementedBankServiceServer) mustEmbedUnimplementedBankServiceServer() {}

// UnsafeBankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServiceServer will
// result in compilation errors.
type UnsafeBankServiceServer interface {
	mustEmbedUnimplementedBankServiceServer()
}

func RegisterBankServiceServer(s grpc.ServiceRegistrar, srv BankServiceServer) {
	s.RegisterService(&BankService_ServiceDesc, srv)
}

func _BankService_ListBankAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBankAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).ListBankAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/ListBankAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ListBankAccount(ctx, req.(*ListBankAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_CreateOneAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOneAccountInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).CreateOneAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/CreateOneAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).CreateOneAccountInfo(ctx, req.(*CreateOneAccountInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_UpdateBankAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBankAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).UpdateBankAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/UpdateBankAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).UpdateBankAccount(ctx, req.(*UpdateBankAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_UpdateBankAccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBankAccountListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).UpdateBankAccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/UpdateBankAccountList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).UpdateBankAccountList(ctx, req.(*UpdateBankAccountListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_GetCompanyCreditConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyCreditConfigsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).GetCompanyCreditConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/GetCompanyCreditConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).GetCompanyCreditConfigs(ctx, req.(*GetCompanyCreditConfigsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_UpdateCompanyCreditConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyCreditConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).UpdateCompanyCreditConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/UpdateCompanyCreditConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).UpdateCompanyCreditConfig(ctx, req.(*UpdateCompanyCreditConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_GetCompanyBankAccountDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyBankAccountDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).GetCompanyBankAccountDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/GetCompanyBankAccountDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).GetCompanyBankAccountDetail(ctx, req.(*GetCompanyBankAccountDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_CompanyWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyWithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).CompanyWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/CompanyWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).CompanyWithdraw(ctx, req.(*CompanyWithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_GetWithdrawReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawReceiptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).GetWithdrawReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/GetWithdrawReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).GetWithdrawReceipt(ctx, req.(*GetWithdrawReceiptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_ListCompanyWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompanyWithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).ListCompanyWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.BankService/ListCompanyWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ListCompanyWithdraw(ctx, req.(*ListCompanyWithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BankService_ServiceDesc is the grpc.ServiceDesc for BankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BankService.BankService",
	HandlerType: (*BankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBankAccount",
			Handler:    _BankService_ListBankAccount_Handler,
		},
		{
			MethodName: "CreateOneAccountInfo",
			Handler:    _BankService_CreateOneAccountInfo_Handler,
		},
		{
			MethodName: "UpdateBankAccount",
			Handler:    _BankService_UpdateBankAccount_Handler,
		},
		{
			MethodName: "UpdateBankAccountList",
			Handler:    _BankService_UpdateBankAccountList_Handler,
		},
		{
			MethodName: "GetCompanyCreditConfigs",
			Handler:    _BankService_GetCompanyCreditConfigs_Handler,
		},
		{
			MethodName: "UpdateCompanyCreditConfig",
			Handler:    _BankService_UpdateCompanyCreditConfig_Handler,
		},
		{
			MethodName: "GetCompanyBankAccountDetail",
			Handler:    _BankService_GetCompanyBankAccountDetail_Handler,
		},
		{
			MethodName: "CompanyWithdraw",
			Handler:    _BankService_CompanyWithdraw_Handler,
		},
		{
			MethodName: "GetWithdrawReceipt",
			Handler:    _BankService_GetWithdrawReceipt_Handler,
		},
		{
			MethodName: "ListCompanyWithdraw",
			Handler:    _BankService_ListCompanyWithdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BankService.proto",
}

// ShipmentPayOrderServiceClient is the client API for ShipmentPayOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShipmentPayOrderServiceClient interface {
	//获取支付相关的支付单
	ListPayOrder(ctx context.Context, in *ListPayOrderReq, opts ...grpc.CallOption) (*ListPayOrderResp, error)
	//获取支付相关的支付计划信息
	ListPayOrderPlan(ctx context.Context, in *ListPayOrderPlanReq, opts ...grpc.CallOption) (*ListPayOrderPlanResp, error)
	//线下处置支付单
	FinishPayOrderError(ctx context.Context, in *FinishPayOrderErrorReq, opts ...grpc.CallOption) (*FinishPayOrderErrorResp, error)
	//获取支付回单
	GetTransferReceipt(ctx context.Context, in *GetTransferReceiptReq, opts ...grpc.CallOption) (*GetTransferReceiptResp, error)
	//医院主动批量支付订单，自动按照配送企业进行拆单
	PayShipmentOrders(ctx context.Context, in *PayShipmentOrdersReq, opts ...grpc.CallOption) (*PayShipmentOrdersResp, error)
}

type shipmentPayOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShipmentPayOrderServiceClient(cc grpc.ClientConnInterface) ShipmentPayOrderServiceClient {
	return &shipmentPayOrderServiceClient{cc}
}

func (c *shipmentPayOrderServiceClient) ListPayOrder(ctx context.Context, in *ListPayOrderReq, opts ...grpc.CallOption) (*ListPayOrderResp, error) {
	out := new(ListPayOrderResp)
	err := c.cc.Invoke(ctx, "/BankService.ShipmentPayOrderService/ListPayOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentPayOrderServiceClient) ListPayOrderPlan(ctx context.Context, in *ListPayOrderPlanReq, opts ...grpc.CallOption) (*ListPayOrderPlanResp, error) {
	out := new(ListPayOrderPlanResp)
	err := c.cc.Invoke(ctx, "/BankService.ShipmentPayOrderService/ListPayOrderPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentPayOrderServiceClient) FinishPayOrderError(ctx context.Context, in *FinishPayOrderErrorReq, opts ...grpc.CallOption) (*FinishPayOrderErrorResp, error) {
	out := new(FinishPayOrderErrorResp)
	err := c.cc.Invoke(ctx, "/BankService.ShipmentPayOrderService/FinishPayOrderError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentPayOrderServiceClient) GetTransferReceipt(ctx context.Context, in *GetTransferReceiptReq, opts ...grpc.CallOption) (*GetTransferReceiptResp, error) {
	out := new(GetTransferReceiptResp)
	err := c.cc.Invoke(ctx, "/BankService.ShipmentPayOrderService/GetTransferReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentPayOrderServiceClient) PayShipmentOrders(ctx context.Context, in *PayShipmentOrdersReq, opts ...grpc.CallOption) (*PayShipmentOrdersResp, error) {
	out := new(PayShipmentOrdersResp)
	err := c.cc.Invoke(ctx, "/BankService.ShipmentPayOrderService/PayShipmentOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipmentPayOrderServiceServer is the server API for ShipmentPayOrderService service.
// All implementations must embed UnimplementedShipmentPayOrderServiceServer
// for forward compatibility
type ShipmentPayOrderServiceServer interface {
	//获取支付相关的支付单
	ListPayOrder(context.Context, *ListPayOrderReq) (*ListPayOrderResp, error)
	//获取支付相关的支付计划信息
	ListPayOrderPlan(context.Context, *ListPayOrderPlanReq) (*ListPayOrderPlanResp, error)
	//线下处置支付单
	FinishPayOrderError(context.Context, *FinishPayOrderErrorReq) (*FinishPayOrderErrorResp, error)
	//获取支付回单
	GetTransferReceipt(context.Context, *GetTransferReceiptReq) (*GetTransferReceiptResp, error)
	//医院主动批量支付订单，自动按照配送企业进行拆单
	PayShipmentOrders(context.Context, *PayShipmentOrdersReq) (*PayShipmentOrdersResp, error)
	mustEmbedUnimplementedShipmentPayOrderServiceServer()
}

// UnimplementedShipmentPayOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShipmentPayOrderServiceServer struct {
}

func (UnimplementedShipmentPayOrderServiceServer) ListPayOrder(context.Context, *ListPayOrderReq) (*ListPayOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayOrder not implemented")
}
func (UnimplementedShipmentPayOrderServiceServer) ListPayOrderPlan(context.Context, *ListPayOrderPlanReq) (*ListPayOrderPlanResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayOrderPlan not implemented")
}
func (UnimplementedShipmentPayOrderServiceServer) FinishPayOrderError(context.Context, *FinishPayOrderErrorReq) (*FinishPayOrderErrorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishPayOrderError not implemented")
}
func (UnimplementedShipmentPayOrderServiceServer) GetTransferReceipt(context.Context, *GetTransferReceiptReq) (*GetTransferReceiptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferReceipt not implemented")
}
func (UnimplementedShipmentPayOrderServiceServer) PayShipmentOrders(context.Context, *PayShipmentOrdersReq) (*PayShipmentOrdersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayShipmentOrders not implemented")
}
func (UnimplementedShipmentPayOrderServiceServer) mustEmbedUnimplementedShipmentPayOrderServiceServer() {
}

// UnsafeShipmentPayOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShipmentPayOrderServiceServer will
// result in compilation errors.
type UnsafeShipmentPayOrderServiceServer interface {
	mustEmbedUnimplementedShipmentPayOrderServiceServer()
}

func RegisterShipmentPayOrderServiceServer(s grpc.ServiceRegistrar, srv ShipmentPayOrderServiceServer) {
	s.RegisterService(&ShipmentPayOrderService_ServiceDesc, srv)
}

func _ShipmentPayOrderService_ListPayOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPayOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentPayOrderServiceServer).ListPayOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.ShipmentPayOrderService/ListPayOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentPayOrderServiceServer).ListPayOrder(ctx, req.(*ListPayOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentPayOrderService_ListPayOrderPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPayOrderPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentPayOrderServiceServer).ListPayOrderPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.ShipmentPayOrderService/ListPayOrderPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentPayOrderServiceServer).ListPayOrderPlan(ctx, req.(*ListPayOrderPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentPayOrderService_FinishPayOrderError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishPayOrderErrorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentPayOrderServiceServer).FinishPayOrderError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.ShipmentPayOrderService/FinishPayOrderError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentPayOrderServiceServer).FinishPayOrderError(ctx, req.(*FinishPayOrderErrorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentPayOrderService_GetTransferReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransferReceiptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentPayOrderServiceServer).GetTransferReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.ShipmentPayOrderService/GetTransferReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentPayOrderServiceServer).GetTransferReceipt(ctx, req.(*GetTransferReceiptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentPayOrderService_PayShipmentOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayShipmentOrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentPayOrderServiceServer).PayShipmentOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.ShipmentPayOrderService/PayShipmentOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentPayOrderServiceServer).PayShipmentOrders(ctx, req.(*PayShipmentOrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ShipmentPayOrderService_ServiceDesc is the grpc.ServiceDesc for ShipmentPayOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShipmentPayOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BankService.ShipmentPayOrderService",
	HandlerType: (*ShipmentPayOrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPayOrder",
			Handler:    _ShipmentPayOrderService_ListPayOrder_Handler,
		},
		{
			MethodName: "ListPayOrderPlan",
			Handler:    _ShipmentPayOrderService_ListPayOrderPlan_Handler,
		},
		{
			MethodName: "FinishPayOrderError",
			Handler:    _ShipmentPayOrderService_FinishPayOrderError_Handler,
		},
		{
			MethodName: "GetTransferReceipt",
			Handler:    _ShipmentPayOrderService_GetTransferReceipt_Handler,
		},
		{
			MethodName: "PayShipmentOrders",
			Handler:    _ShipmentPayOrderService_PayShipmentOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BankService.proto",
}

// FactoringOrderServiceClient is the client API for FactoringOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FactoringOrderServiceClient interface {
	//获取保理单一般信息
	ListFactoringOrder(ctx context.Context, in *ListFactoringOrderReq, opts ...grpc.CallOption) (*ListFactoringOrderResp, error)
	//获取保理单下的支付单信息
	ListFactoringPayOrder(ctx context.Context, in *ListFactoringPayOrderReq, opts ...grpc.CallOption) (*ListFactoringPayOrderResp, error)
	//获取保理单下的配送计划信息
	ListFactoringOrderPlan(ctx context.Context, in *ListFactoringOrderPlanReq, opts ...grpc.CallOption) (*ListFactoringOrderPlanResp, error)
	//线下处置保理单
	FinishFactoringOrderError(ctx context.Context, in *FinishFactoringOrderErrorReq, opts ...grpc.CallOption) (*FinishFactoringOrderErrorResp, error)
	//发起保理请求
	Apply(ctx context.Context, in *ApplyReq, opts ...grpc.CallOption) (*ApplyResp, error)
}

type factoringOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFactoringOrderServiceClient(cc grpc.ClientConnInterface) FactoringOrderServiceClient {
	return &factoringOrderServiceClient{cc}
}

func (c *factoringOrderServiceClient) ListFactoringOrder(ctx context.Context, in *ListFactoringOrderReq, opts ...grpc.CallOption) (*ListFactoringOrderResp, error) {
	out := new(ListFactoringOrderResp)
	err := c.cc.Invoke(ctx, "/BankService.FactoringOrderService/ListFactoringOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factoringOrderServiceClient) ListFactoringPayOrder(ctx context.Context, in *ListFactoringPayOrderReq, opts ...grpc.CallOption) (*ListFactoringPayOrderResp, error) {
	out := new(ListFactoringPayOrderResp)
	err := c.cc.Invoke(ctx, "/BankService.FactoringOrderService/ListFactoringPayOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factoringOrderServiceClient) ListFactoringOrderPlan(ctx context.Context, in *ListFactoringOrderPlanReq, opts ...grpc.CallOption) (*ListFactoringOrderPlanResp, error) {
	out := new(ListFactoringOrderPlanResp)
	err := c.cc.Invoke(ctx, "/BankService.FactoringOrderService/ListFactoringOrderPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factoringOrderServiceClient) FinishFactoringOrderError(ctx context.Context, in *FinishFactoringOrderErrorReq, opts ...grpc.CallOption) (*FinishFactoringOrderErrorResp, error) {
	out := new(FinishFactoringOrderErrorResp)
	err := c.cc.Invoke(ctx, "/BankService.FactoringOrderService/FinishFactoringOrderError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *factoringOrderServiceClient) Apply(ctx context.Context, in *ApplyReq, opts ...grpc.CallOption) (*ApplyResp, error) {
	out := new(ApplyResp)
	err := c.cc.Invoke(ctx, "/BankService.FactoringOrderService/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FactoringOrderServiceServer is the server API for FactoringOrderService service.
// All implementations must embed UnimplementedFactoringOrderServiceServer
// for forward compatibility
type FactoringOrderServiceServer interface {
	//获取保理单一般信息
	ListFactoringOrder(context.Context, *ListFactoringOrderReq) (*ListFactoringOrderResp, error)
	//获取保理单下的支付单信息
	ListFactoringPayOrder(context.Context, *ListFactoringPayOrderReq) (*ListFactoringPayOrderResp, error)
	//获取保理单下的配送计划信息
	ListFactoringOrderPlan(context.Context, *ListFactoringOrderPlanReq) (*ListFactoringOrderPlanResp, error)
	//线下处置保理单
	FinishFactoringOrderError(context.Context, *FinishFactoringOrderErrorReq) (*FinishFactoringOrderErrorResp, error)
	//发起保理请求
	Apply(context.Context, *ApplyReq) (*ApplyResp, error)
	mustEmbedUnimplementedFactoringOrderServiceServer()
}

// UnimplementedFactoringOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFactoringOrderServiceServer struct {
}

func (UnimplementedFactoringOrderServiceServer) ListFactoringOrder(context.Context, *ListFactoringOrderReq) (*ListFactoringOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFactoringOrder not implemented")
}
func (UnimplementedFactoringOrderServiceServer) ListFactoringPayOrder(context.Context, *ListFactoringPayOrderReq) (*ListFactoringPayOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFactoringPayOrder not implemented")
}
func (UnimplementedFactoringOrderServiceServer) ListFactoringOrderPlan(context.Context, *ListFactoringOrderPlanReq) (*ListFactoringOrderPlanResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFactoringOrderPlan not implemented")
}
func (UnimplementedFactoringOrderServiceServer) FinishFactoringOrderError(context.Context, *FinishFactoringOrderErrorReq) (*FinishFactoringOrderErrorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishFactoringOrderError not implemented")
}
func (UnimplementedFactoringOrderServiceServer) Apply(context.Context, *ApplyReq) (*ApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedFactoringOrderServiceServer) mustEmbedUnimplementedFactoringOrderServiceServer() {}

// UnsafeFactoringOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FactoringOrderServiceServer will
// result in compilation errors.
type UnsafeFactoringOrderServiceServer interface {
	mustEmbedUnimplementedFactoringOrderServiceServer()
}

func RegisterFactoringOrderServiceServer(s grpc.ServiceRegistrar, srv FactoringOrderServiceServer) {
	s.RegisterService(&FactoringOrderService_ServiceDesc, srv)
}

func _FactoringOrderService_ListFactoringOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFactoringOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoringOrderServiceServer).ListFactoringOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.FactoringOrderService/ListFactoringOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoringOrderServiceServer).ListFactoringOrder(ctx, req.(*ListFactoringOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FactoringOrderService_ListFactoringPayOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFactoringPayOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoringOrderServiceServer).ListFactoringPayOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.FactoringOrderService/ListFactoringPayOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoringOrderServiceServer).ListFactoringPayOrder(ctx, req.(*ListFactoringPayOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FactoringOrderService_ListFactoringOrderPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFactoringOrderPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoringOrderServiceServer).ListFactoringOrderPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.FactoringOrderService/ListFactoringOrderPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoringOrderServiceServer).ListFactoringOrderPlan(ctx, req.(*ListFactoringOrderPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FactoringOrderService_FinishFactoringOrderError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishFactoringOrderErrorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoringOrderServiceServer).FinishFactoringOrderError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.FactoringOrderService/FinishFactoringOrderError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoringOrderServiceServer).FinishFactoringOrderError(ctx, req.(*FinishFactoringOrderErrorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FactoringOrderService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoringOrderServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BankService.FactoringOrderService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoringOrderServiceServer).Apply(ctx, req.(*ApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FactoringOrderService_ServiceDesc is the grpc.ServiceDesc for FactoringOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FactoringOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BankService.FactoringOrderService",
	HandlerType: (*FactoringOrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFactoringOrder",
			Handler:    _FactoringOrderService_ListFactoringOrder_Handler,
		},
		{
			MethodName: "ListFactoringPayOrder",
			Handler:    _FactoringOrderService_ListFactoringPayOrder_Handler,
		},
		{
			MethodName: "ListFactoringOrderPlan",
			Handler:    _FactoringOrderService_ListFactoringOrderPlan_Handler,
		},
		{
			MethodName: "FinishFactoringOrderError",
			Handler:    _FactoringOrderService_FinishFactoringOrderError_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _FactoringOrderService_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BankService.proto",
}