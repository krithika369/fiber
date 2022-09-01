package grpc

import (
	"strings"

	"github.com/gojek/fiber"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Response struct {
	Metadata metadata.MD
	Message  proto.Message
	Status   status.Status
}

func (r *Response) IsSuccess() bool {
	return r.StatusCode() == int(codes.OK)
}

func (r *Response) Payload() interface{} {
	return r.Message
}

func (r *Response) StatusCode() int {
	return int(r.Status.Code())
}

func (r *Response) BackendName() string {
	return strings.Join(r.Metadata.Get("backend"), ",")
}

func (r *Response) WithBackendName(backendName string) fiber.Response {
	r.Metadata.Set("backend", backendName)
	return r
}
