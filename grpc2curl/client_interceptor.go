package grpc2curl

import (
	"context"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
)

func GRPCCliInterceptor(addr string) func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		span, ctx := opentracing.StartSpanFromContext(ctx, "grpc2curl:GRPCCliInterceptor:"+method)
		defer span.Finish()

		m := method
		if len(m) > 0 && m[0] == '/' {
			m = m[1:]
		}

		mar := jsonpb.Marshaler{}
		reqBuilder, callBuilder := strings.Builder{}, strings.Builder{}

		_ = mar.Marshal(&reqBuilder, req.(proto.Message))

		callBuilder.WriteString("grpc_cli call --json_input --json_output ")
		callBuilder.WriteString(addr)
		callBuilder.WriteRune(' ')
		callBuilder.WriteString(m)
		callBuilder.WriteRune(' ')
		callBuilder.WriteRune('\'')
		callBuilder.WriteString(reqBuilder.String())
		callBuilder.WriteRune('\'')

		curl := callBuilder.String()

		span.LogFields(
			log.String("grpc_method", m),
			log.String("grpc_call", curl),
		)

		return invoker(ctx, method, req, reply, cc, opts...)
	}

}
