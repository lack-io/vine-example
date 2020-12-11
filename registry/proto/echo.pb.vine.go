// Code generated by proto-gen-vine. DO NOT EDIT.
// source: registry/proto/echo.proto

package echo

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/lack-io/vine/api"
	client "github.com/lack-io/vine/client"
	server "github.com/lack-io/vine/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Echo service
func NewEchoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Echo service
type EchoService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type echoService struct {
	c    client.Client
	name string
}

func NewEchoService(name string, c client.Client) EchoService {
	return &echoService{
		c:    c,
		name: name,
	}
}

func (c *echoService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Echo.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service
type EchoHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterEchoHandler(s server.Server, hdlr EchoHandler, opts ...server.HandlerOption) error {
	type echo interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type Echo struct {
		echo
	}
	h := &echoHandler{hdlr}
	return s.Handle(s.NewHandler(&Echo{h}, opts...))
}

type echoHandler struct {
	EchoHandler
}

func (h *echoHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.EchoHandler.Call(ctx, in, out)
}
