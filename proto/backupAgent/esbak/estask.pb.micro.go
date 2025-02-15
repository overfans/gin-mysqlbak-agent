// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/backupAgent/esbak/estask.proto

package esbak

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for EsService service

func NewEsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for EsService service

type EsService interface {
	TaskAdd(ctx context.Context, in *EsBakTaskADDInput, opts ...client.CallOption) (*EsOneMessage, error)
	TaskDelete(ctx context.Context, in *EsTaskIDInput, opts ...client.CallOption) (*EsOneMessage, error)
	TaskUpdate(ctx context.Context, in *EsBakTaskUpdateInput, opts ...client.CallOption) (*EsOneMessage, error)
	GetTaskList(ctx context.Context, in *EsTaskListInput, opts ...client.CallOption) (*EsTaskListOutPut, error)
	GetTaskDetail(ctx context.Context, in *EsTaskIDInput, opts ...client.CallOption) (*EsTaskDetailOutPut, error)
}

type esService struct {
	c    client.Client
	name string
}

func NewEsService(name string, c client.Client) EsService {
	return &esService{
		c:    c,
		name: name,
	}
}

func (c *esService) TaskAdd(ctx context.Context, in *EsBakTaskADDInput, opts ...client.CallOption) (*EsOneMessage, error) {
	req := c.c.NewRequest(c.name, "EsService.TaskAdd", in)
	out := new(EsOneMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esService) TaskDelete(ctx context.Context, in *EsTaskIDInput, opts ...client.CallOption) (*EsOneMessage, error) {
	req := c.c.NewRequest(c.name, "EsService.TaskDelete", in)
	out := new(EsOneMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esService) TaskUpdate(ctx context.Context, in *EsBakTaskUpdateInput, opts ...client.CallOption) (*EsOneMessage, error) {
	req := c.c.NewRequest(c.name, "EsService.TaskUpdate", in)
	out := new(EsOneMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esService) GetTaskList(ctx context.Context, in *EsTaskListInput, opts ...client.CallOption) (*EsTaskListOutPut, error) {
	req := c.c.NewRequest(c.name, "EsService.GetTaskList", in)
	out := new(EsTaskListOutPut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *esService) GetTaskDetail(ctx context.Context, in *EsTaskIDInput, opts ...client.CallOption) (*EsTaskDetailOutPut, error) {
	req := c.c.NewRequest(c.name, "EsService.GetTaskDetail", in)
	out := new(EsTaskDetailOutPut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EsService service

type EsServiceHandler interface {
	TaskAdd(context.Context, *EsBakTaskADDInput, *EsOneMessage) error
	TaskDelete(context.Context, *EsTaskIDInput, *EsOneMessage) error
	TaskUpdate(context.Context, *EsBakTaskUpdateInput, *EsOneMessage) error
	GetTaskList(context.Context, *EsTaskListInput, *EsTaskListOutPut) error
	GetTaskDetail(context.Context, *EsTaskIDInput, *EsTaskDetailOutPut) error
}

func RegisterEsServiceHandler(s server.Server, hdlr EsServiceHandler, opts ...server.HandlerOption) error {
	type esService interface {
		TaskAdd(ctx context.Context, in *EsBakTaskADDInput, out *EsOneMessage) error
		TaskDelete(ctx context.Context, in *EsTaskIDInput, out *EsOneMessage) error
		TaskUpdate(ctx context.Context, in *EsBakTaskUpdateInput, out *EsOneMessage) error
		GetTaskList(ctx context.Context, in *EsTaskListInput, out *EsTaskListOutPut) error
		GetTaskDetail(ctx context.Context, in *EsTaskIDInput, out *EsTaskDetailOutPut) error
	}
	type EsService struct {
		esService
	}
	h := &esServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&EsService{h}, opts...))
}

type esServiceHandler struct {
	EsServiceHandler
}

func (h *esServiceHandler) TaskAdd(ctx context.Context, in *EsBakTaskADDInput, out *EsOneMessage) error {
	return h.EsServiceHandler.TaskAdd(ctx, in, out)
}

func (h *esServiceHandler) TaskDelete(ctx context.Context, in *EsTaskIDInput, out *EsOneMessage) error {
	return h.EsServiceHandler.TaskDelete(ctx, in, out)
}

func (h *esServiceHandler) TaskUpdate(ctx context.Context, in *EsBakTaskUpdateInput, out *EsOneMessage) error {
	return h.EsServiceHandler.TaskUpdate(ctx, in, out)
}

func (h *esServiceHandler) GetTaskList(ctx context.Context, in *EsTaskListInput, out *EsTaskListOutPut) error {
	return h.EsServiceHandler.GetTaskList(ctx, in, out)
}

func (h *esServiceHandler) GetTaskDetail(ctx context.Context, in *EsTaskIDInput, out *EsTaskDetailOutPut) error {
	return h.EsServiceHandler.GetTaskDetail(ctx, in, out)
}
