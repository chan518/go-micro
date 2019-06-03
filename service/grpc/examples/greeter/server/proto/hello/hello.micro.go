// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/go-micro/service/grpc/examples/greeter/server/proto/hello/hello.proto

/*
Package go_micro_srv_greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-micro/service/grpc/examples/greeter/server/proto/hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Say service

type SayService interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type sayService struct {
	c           client.Client
	serviceName string
}

func NewSayService(serviceName string, c client.Client) SayService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.greeter"
	}
	return &sayService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *sayService) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Say.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterSayHandler(s server.Server, hdlr SayHandler, opts ...server.HandlerOption) {
	type say interface {
		Hello(ctx context.Context, in *Request, out *Response) error
	}
	type Say struct {
		say
	}
	h := &sayHandler{hdlr}
	s.Handle(s.NewHandler(&Say{h}, opts...))
}

type sayHandler struct {
	SayHandler
}

func (h *sayHandler) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.SayHandler.Hello(ctx, in, out)
}
