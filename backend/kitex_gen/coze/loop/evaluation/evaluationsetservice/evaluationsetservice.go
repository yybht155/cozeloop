// Code generated by Kitex v0.13.1. DO NOT EDIT.

package evaluationsetservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	evaluation "github.com/coze-dev/cozeloop/backend/kitex_gen/coze/loop/evaluation"
	eval_set "github.com/coze-dev/cozeloop/backend/kitex_gen/coze/loop/evaluation/eval_set"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateEvaluationSet": kitex.NewMethodInfo(
		createEvaluationSetHandler,
		newEvaluationSetServiceCreateEvaluationSetArgs,
		newEvaluationSetServiceCreateEvaluationSetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateEvaluationSet": kitex.NewMethodInfo(
		updateEvaluationSetHandler,
		newEvaluationSetServiceUpdateEvaluationSetArgs,
		newEvaluationSetServiceUpdateEvaluationSetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteEvaluationSet": kitex.NewMethodInfo(
		deleteEvaluationSetHandler,
		newEvaluationSetServiceDeleteEvaluationSetArgs,
		newEvaluationSetServiceDeleteEvaluationSetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetEvaluationSet": kitex.NewMethodInfo(
		getEvaluationSetHandler,
		newEvaluationSetServiceGetEvaluationSetArgs,
		newEvaluationSetServiceGetEvaluationSetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListEvaluationSets": kitex.NewMethodInfo(
		listEvaluationSetsHandler,
		newEvaluationSetServiceListEvaluationSetsArgs,
		newEvaluationSetServiceListEvaluationSetsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateEvaluationSetVersion": kitex.NewMethodInfo(
		createEvaluationSetVersionHandler,
		newEvaluationSetServiceCreateEvaluationSetVersionArgs,
		newEvaluationSetServiceCreateEvaluationSetVersionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetEvaluationSetVersion": kitex.NewMethodInfo(
		getEvaluationSetVersionHandler,
		newEvaluationSetServiceGetEvaluationSetVersionArgs,
		newEvaluationSetServiceGetEvaluationSetVersionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListEvaluationSetVersions": kitex.NewMethodInfo(
		listEvaluationSetVersionsHandler,
		newEvaluationSetServiceListEvaluationSetVersionsArgs,
		newEvaluationSetServiceListEvaluationSetVersionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"BatchGetEvaluationSetVersions": kitex.NewMethodInfo(
		batchGetEvaluationSetVersionsHandler,
		newEvaluationSetServiceBatchGetEvaluationSetVersionsArgs,
		newEvaluationSetServiceBatchGetEvaluationSetVersionsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateEvaluationSetSchema": kitex.NewMethodInfo(
		updateEvaluationSetSchemaHandler,
		newEvaluationSetServiceUpdateEvaluationSetSchemaArgs,
		newEvaluationSetServiceUpdateEvaluationSetSchemaResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"BatchCreateEvaluationSetItems": kitex.NewMethodInfo(
		batchCreateEvaluationSetItemsHandler,
		newEvaluationSetServiceBatchCreateEvaluationSetItemsArgs,
		newEvaluationSetServiceBatchCreateEvaluationSetItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateEvaluationSetItem": kitex.NewMethodInfo(
		updateEvaluationSetItemHandler,
		newEvaluationSetServiceUpdateEvaluationSetItemArgs,
		newEvaluationSetServiceUpdateEvaluationSetItemResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"BatchDeleteEvaluationSetItems": kitex.NewMethodInfo(
		batchDeleteEvaluationSetItemsHandler,
		newEvaluationSetServiceBatchDeleteEvaluationSetItemsArgs,
		newEvaluationSetServiceBatchDeleteEvaluationSetItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListEvaluationSetItems": kitex.NewMethodInfo(
		listEvaluationSetItemsHandler,
		newEvaluationSetServiceListEvaluationSetItemsArgs,
		newEvaluationSetServiceListEvaluationSetItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"BatchGetEvaluationSetItems": kitex.NewMethodInfo(
		batchGetEvaluationSetItemsHandler,
		newEvaluationSetServiceBatchGetEvaluationSetItemsArgs,
		newEvaluationSetServiceBatchGetEvaluationSetItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ClearEvaluationSetDraftItem": kitex.NewMethodInfo(
		clearEvaluationSetDraftItemHandler,
		newEvaluationSetServiceClearEvaluationSetDraftItemArgs,
		newEvaluationSetServiceClearEvaluationSetDraftItemResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	evaluationSetServiceServiceInfo = NewServiceInfo()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return evaluationSetServiceServiceInfo
}

// NewServiceInfo creates a new ServiceInfo
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo()
}

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "EvaluationSetService"
	handlerType := (*evaluation.EvaluationSetService)(nil)
	extra := map[string]interface{}{
		"PackageName": "evaluation",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         serviceMethods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.13.1",
		Extra:           extra,
	}
	return svcInfo
}

func createEvaluationSetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceCreateEvaluationSetArgs)
	realResult := result.(*eval_set.EvaluationSetServiceCreateEvaluationSetResult)
	success, err := handler.(eval_set.EvaluationSetService).CreateEvaluationSet(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceCreateEvaluationSetArgs() interface{} {
	return eval_set.NewEvaluationSetServiceCreateEvaluationSetArgs()
}

func newEvaluationSetServiceCreateEvaluationSetResult() interface{} {
	return eval_set.NewEvaluationSetServiceCreateEvaluationSetResult()
}

func updateEvaluationSetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceUpdateEvaluationSetArgs)
	realResult := result.(*eval_set.EvaluationSetServiceUpdateEvaluationSetResult)
	success, err := handler.(eval_set.EvaluationSetService).UpdateEvaluationSet(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceUpdateEvaluationSetArgs() interface{} {
	return eval_set.NewEvaluationSetServiceUpdateEvaluationSetArgs()
}

func newEvaluationSetServiceUpdateEvaluationSetResult() interface{} {
	return eval_set.NewEvaluationSetServiceUpdateEvaluationSetResult()
}

func deleteEvaluationSetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceDeleteEvaluationSetArgs)
	realResult := result.(*eval_set.EvaluationSetServiceDeleteEvaluationSetResult)
	success, err := handler.(eval_set.EvaluationSetService).DeleteEvaluationSet(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceDeleteEvaluationSetArgs() interface{} {
	return eval_set.NewEvaluationSetServiceDeleteEvaluationSetArgs()
}

func newEvaluationSetServiceDeleteEvaluationSetResult() interface{} {
	return eval_set.NewEvaluationSetServiceDeleteEvaluationSetResult()
}

func getEvaluationSetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceGetEvaluationSetArgs)
	realResult := result.(*eval_set.EvaluationSetServiceGetEvaluationSetResult)
	success, err := handler.(eval_set.EvaluationSetService).GetEvaluationSet(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceGetEvaluationSetArgs() interface{} {
	return eval_set.NewEvaluationSetServiceGetEvaluationSetArgs()
}

func newEvaluationSetServiceGetEvaluationSetResult() interface{} {
	return eval_set.NewEvaluationSetServiceGetEvaluationSetResult()
}

func listEvaluationSetsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceListEvaluationSetsArgs)
	realResult := result.(*eval_set.EvaluationSetServiceListEvaluationSetsResult)
	success, err := handler.(eval_set.EvaluationSetService).ListEvaluationSets(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceListEvaluationSetsArgs() interface{} {
	return eval_set.NewEvaluationSetServiceListEvaluationSetsArgs()
}

func newEvaluationSetServiceListEvaluationSetsResult() interface{} {
	return eval_set.NewEvaluationSetServiceListEvaluationSetsResult()
}

func createEvaluationSetVersionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceCreateEvaluationSetVersionArgs)
	realResult := result.(*eval_set.EvaluationSetServiceCreateEvaluationSetVersionResult)
	success, err := handler.(eval_set.EvaluationSetService).CreateEvaluationSetVersion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceCreateEvaluationSetVersionArgs() interface{} {
	return eval_set.NewEvaluationSetServiceCreateEvaluationSetVersionArgs()
}

func newEvaluationSetServiceCreateEvaluationSetVersionResult() interface{} {
	return eval_set.NewEvaluationSetServiceCreateEvaluationSetVersionResult()
}

func getEvaluationSetVersionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceGetEvaluationSetVersionArgs)
	realResult := result.(*eval_set.EvaluationSetServiceGetEvaluationSetVersionResult)
	success, err := handler.(eval_set.EvaluationSetService).GetEvaluationSetVersion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceGetEvaluationSetVersionArgs() interface{} {
	return eval_set.NewEvaluationSetServiceGetEvaluationSetVersionArgs()
}

func newEvaluationSetServiceGetEvaluationSetVersionResult() interface{} {
	return eval_set.NewEvaluationSetServiceGetEvaluationSetVersionResult()
}

func listEvaluationSetVersionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceListEvaluationSetVersionsArgs)
	realResult := result.(*eval_set.EvaluationSetServiceListEvaluationSetVersionsResult)
	success, err := handler.(eval_set.EvaluationSetService).ListEvaluationSetVersions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceListEvaluationSetVersionsArgs() interface{} {
	return eval_set.NewEvaluationSetServiceListEvaluationSetVersionsArgs()
}

func newEvaluationSetServiceListEvaluationSetVersionsResult() interface{} {
	return eval_set.NewEvaluationSetServiceListEvaluationSetVersionsResult()
}

func batchGetEvaluationSetVersionsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceBatchGetEvaluationSetVersionsArgs)
	realResult := result.(*eval_set.EvaluationSetServiceBatchGetEvaluationSetVersionsResult)
	success, err := handler.(eval_set.EvaluationSetService).BatchGetEvaluationSetVersions(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceBatchGetEvaluationSetVersionsArgs() interface{} {
	return eval_set.NewEvaluationSetServiceBatchGetEvaluationSetVersionsArgs()
}

func newEvaluationSetServiceBatchGetEvaluationSetVersionsResult() interface{} {
	return eval_set.NewEvaluationSetServiceBatchGetEvaluationSetVersionsResult()
}

func updateEvaluationSetSchemaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceUpdateEvaluationSetSchemaArgs)
	realResult := result.(*eval_set.EvaluationSetServiceUpdateEvaluationSetSchemaResult)
	success, err := handler.(eval_set.EvaluationSetService).UpdateEvaluationSetSchema(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceUpdateEvaluationSetSchemaArgs() interface{} {
	return eval_set.NewEvaluationSetServiceUpdateEvaluationSetSchemaArgs()
}

func newEvaluationSetServiceUpdateEvaluationSetSchemaResult() interface{} {
	return eval_set.NewEvaluationSetServiceUpdateEvaluationSetSchemaResult()
}

func batchCreateEvaluationSetItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceBatchCreateEvaluationSetItemsArgs)
	realResult := result.(*eval_set.EvaluationSetServiceBatchCreateEvaluationSetItemsResult)
	success, err := handler.(eval_set.EvaluationSetService).BatchCreateEvaluationSetItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceBatchCreateEvaluationSetItemsArgs() interface{} {
	return eval_set.NewEvaluationSetServiceBatchCreateEvaluationSetItemsArgs()
}

func newEvaluationSetServiceBatchCreateEvaluationSetItemsResult() interface{} {
	return eval_set.NewEvaluationSetServiceBatchCreateEvaluationSetItemsResult()
}

func updateEvaluationSetItemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceUpdateEvaluationSetItemArgs)
	realResult := result.(*eval_set.EvaluationSetServiceUpdateEvaluationSetItemResult)
	success, err := handler.(eval_set.EvaluationSetService).UpdateEvaluationSetItem(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceUpdateEvaluationSetItemArgs() interface{} {
	return eval_set.NewEvaluationSetServiceUpdateEvaluationSetItemArgs()
}

func newEvaluationSetServiceUpdateEvaluationSetItemResult() interface{} {
	return eval_set.NewEvaluationSetServiceUpdateEvaluationSetItemResult()
}

func batchDeleteEvaluationSetItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceBatchDeleteEvaluationSetItemsArgs)
	realResult := result.(*eval_set.EvaluationSetServiceBatchDeleteEvaluationSetItemsResult)
	success, err := handler.(eval_set.EvaluationSetService).BatchDeleteEvaluationSetItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceBatchDeleteEvaluationSetItemsArgs() interface{} {
	return eval_set.NewEvaluationSetServiceBatchDeleteEvaluationSetItemsArgs()
}

func newEvaluationSetServiceBatchDeleteEvaluationSetItemsResult() interface{} {
	return eval_set.NewEvaluationSetServiceBatchDeleteEvaluationSetItemsResult()
}

func listEvaluationSetItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceListEvaluationSetItemsArgs)
	realResult := result.(*eval_set.EvaluationSetServiceListEvaluationSetItemsResult)
	success, err := handler.(eval_set.EvaluationSetService).ListEvaluationSetItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceListEvaluationSetItemsArgs() interface{} {
	return eval_set.NewEvaluationSetServiceListEvaluationSetItemsArgs()
}

func newEvaluationSetServiceListEvaluationSetItemsResult() interface{} {
	return eval_set.NewEvaluationSetServiceListEvaluationSetItemsResult()
}

func batchGetEvaluationSetItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceBatchGetEvaluationSetItemsArgs)
	realResult := result.(*eval_set.EvaluationSetServiceBatchGetEvaluationSetItemsResult)
	success, err := handler.(eval_set.EvaluationSetService).BatchGetEvaluationSetItems(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceBatchGetEvaluationSetItemsArgs() interface{} {
	return eval_set.NewEvaluationSetServiceBatchGetEvaluationSetItemsArgs()
}

func newEvaluationSetServiceBatchGetEvaluationSetItemsResult() interface{} {
	return eval_set.NewEvaluationSetServiceBatchGetEvaluationSetItemsResult()
}

func clearEvaluationSetDraftItemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eval_set.EvaluationSetServiceClearEvaluationSetDraftItemArgs)
	realResult := result.(*eval_set.EvaluationSetServiceClearEvaluationSetDraftItemResult)
	success, err := handler.(eval_set.EvaluationSetService).ClearEvaluationSetDraftItem(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEvaluationSetServiceClearEvaluationSetDraftItemArgs() interface{} {
	return eval_set.NewEvaluationSetServiceClearEvaluationSetDraftItemArgs()
}

func newEvaluationSetServiceClearEvaluationSetDraftItemResult() interface{} {
	return eval_set.NewEvaluationSetServiceClearEvaluationSetDraftItemResult()
}

type kClient struct {
	c  client.Client
	sc client.Streaming
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c:  c,
		sc: c.(client.Streaming),
	}
}

func (p *kClient) CreateEvaluationSet(ctx context.Context, req *eval_set.CreateEvaluationSetRequest) (r *eval_set.CreateEvaluationSetResponse, err error) {
	var _args eval_set.EvaluationSetServiceCreateEvaluationSetArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceCreateEvaluationSetResult
	if err = p.c.Call(ctx, "CreateEvaluationSet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateEvaluationSet(ctx context.Context, req *eval_set.UpdateEvaluationSetRequest) (r *eval_set.UpdateEvaluationSetResponse, err error) {
	var _args eval_set.EvaluationSetServiceUpdateEvaluationSetArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceUpdateEvaluationSetResult
	if err = p.c.Call(ctx, "UpdateEvaluationSet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteEvaluationSet(ctx context.Context, req *eval_set.DeleteEvaluationSetRequest) (r *eval_set.DeleteEvaluationSetResponse, err error) {
	var _args eval_set.EvaluationSetServiceDeleteEvaluationSetArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceDeleteEvaluationSetResult
	if err = p.c.Call(ctx, "DeleteEvaluationSet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetEvaluationSet(ctx context.Context, req *eval_set.GetEvaluationSetRequest) (r *eval_set.GetEvaluationSetResponse, err error) {
	var _args eval_set.EvaluationSetServiceGetEvaluationSetArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceGetEvaluationSetResult
	if err = p.c.Call(ctx, "GetEvaluationSet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListEvaluationSets(ctx context.Context, req *eval_set.ListEvaluationSetsRequest) (r *eval_set.ListEvaluationSetsResponse, err error) {
	var _args eval_set.EvaluationSetServiceListEvaluationSetsArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceListEvaluationSetsResult
	if err = p.c.Call(ctx, "ListEvaluationSets", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateEvaluationSetVersion(ctx context.Context, req *eval_set.CreateEvaluationSetVersionRequest) (r *eval_set.CreateEvaluationSetVersionResponse, err error) {
	var _args eval_set.EvaluationSetServiceCreateEvaluationSetVersionArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceCreateEvaluationSetVersionResult
	if err = p.c.Call(ctx, "CreateEvaluationSetVersion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetEvaluationSetVersion(ctx context.Context, req *eval_set.GetEvaluationSetVersionRequest) (r *eval_set.GetEvaluationSetVersionResponse, err error) {
	var _args eval_set.EvaluationSetServiceGetEvaluationSetVersionArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceGetEvaluationSetVersionResult
	if err = p.c.Call(ctx, "GetEvaluationSetVersion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListEvaluationSetVersions(ctx context.Context, req *eval_set.ListEvaluationSetVersionsRequest) (r *eval_set.ListEvaluationSetVersionsResponse, err error) {
	var _args eval_set.EvaluationSetServiceListEvaluationSetVersionsArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceListEvaluationSetVersionsResult
	if err = p.c.Call(ctx, "ListEvaluationSetVersions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetEvaluationSetVersions(ctx context.Context, req *eval_set.BatchGetEvaluationSetVersionsRequest) (r *eval_set.BatchGetEvaluationSetVersionsResponse, err error) {
	var _args eval_set.EvaluationSetServiceBatchGetEvaluationSetVersionsArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceBatchGetEvaluationSetVersionsResult
	if err = p.c.Call(ctx, "BatchGetEvaluationSetVersions", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateEvaluationSetSchema(ctx context.Context, req *eval_set.UpdateEvaluationSetSchemaRequest) (r *eval_set.UpdateEvaluationSetSchemaResponse, err error) {
	var _args eval_set.EvaluationSetServiceUpdateEvaluationSetSchemaArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceUpdateEvaluationSetSchemaResult
	if err = p.c.Call(ctx, "UpdateEvaluationSetSchema", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchCreateEvaluationSetItems(ctx context.Context, req *eval_set.BatchCreateEvaluationSetItemsRequest) (r *eval_set.BatchCreateEvaluationSetItemsResponse, err error) {
	var _args eval_set.EvaluationSetServiceBatchCreateEvaluationSetItemsArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceBatchCreateEvaluationSetItemsResult
	if err = p.c.Call(ctx, "BatchCreateEvaluationSetItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateEvaluationSetItem(ctx context.Context, req *eval_set.UpdateEvaluationSetItemRequest) (r *eval_set.UpdateEvaluationSetItemResponse, err error) {
	var _args eval_set.EvaluationSetServiceUpdateEvaluationSetItemArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceUpdateEvaluationSetItemResult
	if err = p.c.Call(ctx, "UpdateEvaluationSetItem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchDeleteEvaluationSetItems(ctx context.Context, req *eval_set.BatchDeleteEvaluationSetItemsRequest) (r *eval_set.BatchDeleteEvaluationSetItemsResponse, err error) {
	var _args eval_set.EvaluationSetServiceBatchDeleteEvaluationSetItemsArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceBatchDeleteEvaluationSetItemsResult
	if err = p.c.Call(ctx, "BatchDeleteEvaluationSetItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListEvaluationSetItems(ctx context.Context, req *eval_set.ListEvaluationSetItemsRequest) (r *eval_set.ListEvaluationSetItemsResponse, err error) {
	var _args eval_set.EvaluationSetServiceListEvaluationSetItemsArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceListEvaluationSetItemsResult
	if err = p.c.Call(ctx, "ListEvaluationSetItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetEvaluationSetItems(ctx context.Context, req *eval_set.BatchGetEvaluationSetItemsRequest) (r *eval_set.BatchGetEvaluationSetItemsResponse, err error) {
	var _args eval_set.EvaluationSetServiceBatchGetEvaluationSetItemsArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceBatchGetEvaluationSetItemsResult
	if err = p.c.Call(ctx, "BatchGetEvaluationSetItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ClearEvaluationSetDraftItem(ctx context.Context, req *eval_set.ClearEvaluationSetDraftItemRequest) (r *eval_set.ClearEvaluationSetDraftItemResponse, err error) {
	var _args eval_set.EvaluationSetServiceClearEvaluationSetDraftItemArgs
	_args.Req = req
	var _result eval_set.EvaluationSetServiceClearEvaluationSetDraftItemResult
	if err = p.c.Call(ctx, "ClearEvaluationSetDraftItem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
