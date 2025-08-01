// Code generated by Kitex v0.13.1. DO NOT EDIT.
package promptdebugservice

import (
	server "github.com/cloudwego/kitex/server"
	debug "github.com/coze-dev/cozeloop/backend/kitex_gen/coze/loop/prompt/debug"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler debug.PromptDebugService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler debug.PromptDebugService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
