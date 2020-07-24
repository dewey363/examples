// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/services/helloworld/proto/helloworld.proto

package go_micro_srv_helloworld

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

// Api Endpoints for Helloworld service

func NewHelloworldEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Helloworld service

type HelloworldService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type helloworldService struct {
	c    client.Client
	name string
}

func NewHelloworldService(name string, c client.Client) HelloworldService {
	return &helloworldService{
		c:    c,
		name: name,
	}
}

func (c *helloworldService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Helloworld.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Helloworld service

type HelloworldHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterHelloworldHandler(s server.Server, hdlr HelloworldHandler, opts ...server.HandlerOption) error {
	type helloworld interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type Helloworld struct {
		helloworld
	}
	h := &helloworldHandler{hdlr}
	return s.Handle(s.NewHandler(&Helloworld{h}, opts...))
}

type helloworldHandler struct {
	HelloworldHandler
}

func (h *helloworldHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.HelloworldHandler.Call(ctx, in, out)
}
