// Code generated by Kitex v0.13.1. DO NOT EDIT.
package traceservice

import (
	server "github.com/cloudwego/kitex/server"
	trace "github.com/coze-dev/cozeloop/backend/kitex_gen/coze/loop/observability/trace"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler trace.TraceService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler trace.TraceService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
