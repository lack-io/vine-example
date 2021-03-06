// Code generated by proto-gen-vine. DO NOT EDIT.
// source: github.com/lack-io/vine-example/helloworld/proto/helloworld.proto

package testdata

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/lack-io/vine-example/goproto/api"
	math "math"
)

import (
	context "context"
	registry "github.com/lack-io/vine/proto/registry"
	api "github.com/lack-io/vine/service/api"
	client "github.com/lack-io/vine/service/client"
	server "github.com/lack-io/vine/service/server"
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
var _ registry.OpenAPI

// API Endpoints for Helloworld service
func NewHelloworldEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Helloworld.Call",
			Path:    []string{"/api/v1/call"},
			Method:  []string{"GET"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Swagger OpenAPI 3.0 for Helloworld service
func NewHelloworldOpenAPI() *registry.OpenAPI {
	return &registry.OpenAPI{
		Openapi: "3.0.1",
		Info: &registry.OpenAPIInfo{
			Title:       "HelloworldService",
			Description: "'OpenAPI3.0 for Helloworld'",
		},
		Version: "v1.0.0",
		Servers: []string{},
		Tags: []*registry.OpenAPITag{
			&registry.OpenAPITag{
				Name:        "Helloworld",
				Description: "OpenAPI3.0 for Helloworld",
			},
		},
		Paths: map[string]*registry.OpenAPIPath{
			"/api/v1/call": &registry.OpenAPIPath{
				Get: &registry.OpenAPIPathDocs{
					Tags:        []string{"Helloworld"},
					Description: "Helloworld Call",
					OperationId: "HelloworldCall",
					Parameters: []*registry.PathParameters{
						&registry.PathParameters{
							Name:        "name",
							In:          "query",
							Description: "HelloWorldRequest field name",
							Required:    true,
							Style:       "form",
							Explode:     true,
							Schema: &registry.Schema{
								Type: "string",
							},
						},
						&registry.PathParameters{
							Name:        "age",
							In:          "query",
							Description: "HelloWorldRequest field age",
							Style:       "form",
							Explode:     true,
							Schema: &registry.Schema{
								Type:   "integer",
								Format: "int32",
							},
						},
					},
					Responses: map[string]*registry.PathResponse{
						"200": &registry.PathResponse{
							Description: "successful response (stream response)",
							Content: &registry.PathRequestBodyContent{
								ApplicationJson: &registry.ApplicationContent{
									Schema: &registry.Schema{Ref: "#/components/errors/HelloWorldResponse"},
								},
							},
						},
					},
					Security: []*registry.PathSecurity{},
				},
			},
		},
		Components: &registry.OpenAPIComponents{
			SecuritySchemas: &registry.SecuritySchemas{},
			Schemas: map[string]*registry.Model{
				"proto.HelloWorldRequest": &registry.Model{
					Type: "object",
					Properties: map[string]*registry.Schema{
						"name": &registry.Schema{
							Type: "string",
						},
						"age": &registry.Schema{
							Type:   "integer",
							Format: "int32",
						},
					},
					Required: []string{"name"},
				},
				"proto.HelloWorldResponse": &registry.Model{
					Type: "object",
					Properties: map[string]*registry.Schema{
						"reply": &registry.Schema{
							Type: "string",
						},
					},
				},
			},
			Errors: map[string]*registry.Model{},
		},
	}
}

// Client API for Helloworld service
// +gen:openapi
type HelloworldService interface {
	// +gen:get=/api/v1/call
	// +gen:body=*
	Call(ctx context.Context, in *HelloWorldRequest, opts ...client.CallOption) (*HelloWorldResponse, error)
	MulPath(ctx context.Context, in *MulPathRequest, opts ...client.CallOption) (*MulPathResponse, error)
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

func (c *helloworldService) Call(ctx context.Context, in *HelloWorldRequest, opts ...client.CallOption) (*HelloWorldResponse, error) {
	req := c.c.NewRequest(c.name, "Helloworld.Call", in)
	out := new(HelloWorldResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldService) MulPath(ctx context.Context, in *MulPathRequest, opts ...client.CallOption) (*MulPathResponse, error) {
	req := c.c.NewRequest(c.name, "Helloworld.MulPath", in)
	out := new(MulPathResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Helloworld service
// +gen:openapi
type HelloworldHandler interface {
	// +gen:get=/api/v1/call
	// +gen:body=*
	Call(context.Context, *HelloWorldRequest, *HelloWorldResponse) error
	MulPath(context.Context, *MulPathRequest, *MulPathResponse) error
}

func RegisterHelloworldHandler(s server.Server, hdlr HelloworldHandler, opts ...server.HandlerOption) error {
	type helloworldImpl interface {
		Call(ctx context.Context, in *HelloWorldRequest, out *HelloWorldResponse) error
		MulPath(ctx context.Context, in *MulPathRequest, out *MulPathResponse) error
	}
	type Helloworld struct {
		helloworldImpl
	}
	h := &helloworldHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Helloworld.Call",
		Path:    []string{"/api/v1/call"},
		Method:  []string{"GET"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, server.OpenAPIHandler(NewHelloworldOpenAPI()))
	return s.Handle(s.NewHandler(&Helloworld{h}, opts...))
}

type helloworldHandler struct {
	HelloworldHandler
}

func (h *helloworldHandler) Call(ctx context.Context, in *HelloWorldRequest, out *HelloWorldResponse) error {
	return h.HelloworldHandler.Call(ctx, in, out)
}

func (h *helloworldHandler) MulPath(ctx context.Context, in *MulPathRequest, out *MulPathResponse) error {
	return h.HelloworldHandler.MulPath(ctx, in, out)
}
