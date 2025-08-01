// Code generated by Kitex v0.13.1. DO NOT EDIT.
package fileservice

import (
	server "github.com/cloudwego/kitex/server"
	foundation "github.com/coze-dev/cozeloop/backend/kitex_gen/coze/loop/foundation"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler foundation.FileService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler foundation.FileService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
