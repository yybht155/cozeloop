// Code generated by cozeloop. DO NOT EDIT.
package lodataset // import github.com/coze-dev/cozeloop/backend/lodataset

import (
	"context"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/coze-dev/cozeloop/backend/kitex_gen/coze/loop/data/dataset"
)

type LocalDatasetService struct {
	impl dataset.DatasetService // the service implementation
	mds  endpoint.Middleware
}

func NewLocalDatasetService(impl dataset.DatasetService, mds ...endpoint.Middleware) *LocalDatasetService {
	return &LocalDatasetService{
		impl: impl,
		mds:  endpoint.Chain(mds...),
	}
}

// CreateDataset
/* Dataset */
// 新增数据集
func (l *LocalDatasetService) CreateDataset(ctx context.Context, req *dataset.CreateDatasetRequest, callOptions ...callopt.Option) (*dataset.CreateDatasetResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceCreateDatasetArgs)
		result := out.(*dataset.DatasetServiceCreateDatasetResult)
		resp, err := l.impl.CreateDataset(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceCreateDatasetArgs{Req: req}
	result := &dataset.DatasetServiceCreateDatasetResult{}
	ctx = l.injectRPCInfo(ctx, "CreateDataset")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// UpdateDataset
// 修改数据集
func (l *LocalDatasetService) UpdateDataset(ctx context.Context, req *dataset.UpdateDatasetRequest, callOptions ...callopt.Option) (*dataset.UpdateDatasetResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceUpdateDatasetArgs)
		result := out.(*dataset.DatasetServiceUpdateDatasetResult)
		resp, err := l.impl.UpdateDataset(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceUpdateDatasetArgs{Req: req}
	result := &dataset.DatasetServiceUpdateDatasetResult{}
	ctx = l.injectRPCInfo(ctx, "UpdateDataset")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// DeleteDataset
// 删除数据集
func (l *LocalDatasetService) DeleteDataset(ctx context.Context, req *dataset.DeleteDatasetRequest, callOptions ...callopt.Option) (*dataset.DeleteDatasetResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceDeleteDatasetArgs)
		result := out.(*dataset.DatasetServiceDeleteDatasetResult)
		resp, err := l.impl.DeleteDataset(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceDeleteDatasetArgs{Req: req}
	result := &dataset.DatasetServiceDeleteDatasetResult{}
	ctx = l.injectRPCInfo(ctx, "DeleteDataset")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// ListDatasets
// 获取数据集列表
func (l *LocalDatasetService) ListDatasets(ctx context.Context, req *dataset.ListDatasetsRequest, callOptions ...callopt.Option) (*dataset.ListDatasetsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceListDatasetsArgs)
		result := out.(*dataset.DatasetServiceListDatasetsResult)
		resp, err := l.impl.ListDatasets(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceListDatasetsArgs{Req: req}
	result := &dataset.DatasetServiceListDatasetsResult{}
	ctx = l.injectRPCInfo(ctx, "ListDatasets")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// GetDataset
// 数据集当前信息（不包括数据）
func (l *LocalDatasetService) GetDataset(ctx context.Context, req *dataset.GetDatasetRequest, callOptions ...callopt.Option) (*dataset.GetDatasetResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceGetDatasetArgs)
		result := out.(*dataset.DatasetServiceGetDatasetResult)
		resp, err := l.impl.GetDataset(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceGetDatasetArgs{Req: req}
	result := &dataset.DatasetServiceGetDatasetResult{}
	ctx = l.injectRPCInfo(ctx, "GetDataset")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// BatchGetDatasets
// 批量获取数据集
func (l *LocalDatasetService) BatchGetDatasets(ctx context.Context, req *dataset.BatchGetDatasetsRequest, callOptions ...callopt.Option) (*dataset.BatchGetDatasetsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceBatchGetDatasetsArgs)
		result := out.(*dataset.DatasetServiceBatchGetDatasetsResult)
		resp, err := l.impl.BatchGetDatasets(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceBatchGetDatasetsArgs{Req: req}
	result := &dataset.DatasetServiceBatchGetDatasetsResult{}
	ctx = l.injectRPCInfo(ctx, "BatchGetDatasets")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// ImportDataset
// 导入数据
func (l *LocalDatasetService) ImportDataset(ctx context.Context, req *dataset.ImportDatasetRequest, callOptions ...callopt.Option) (*dataset.ImportDatasetResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceImportDatasetArgs)
		result := out.(*dataset.DatasetServiceImportDatasetResult)
		resp, err := l.impl.ImportDataset(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceImportDatasetArgs{Req: req}
	result := &dataset.DatasetServiceImportDatasetResult{}
	ctx = l.injectRPCInfo(ctx, "ImportDataset")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// GetDatasetIOJob
// 任务(导入、导出、转换)详情
func (l *LocalDatasetService) GetDatasetIOJob(ctx context.Context, req *dataset.GetDatasetIOJobRequest, callOptions ...callopt.Option) (*dataset.GetDatasetIOJobResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceGetDatasetIOJobArgs)
		result := out.(*dataset.DatasetServiceGetDatasetIOJobResult)
		resp, err := l.impl.GetDatasetIOJob(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceGetDatasetIOJobArgs{Req: req}
	result := &dataset.DatasetServiceGetDatasetIOJobResult{}
	ctx = l.injectRPCInfo(ctx, "GetDatasetIOJob")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// ListDatasetIOJobs
// 数据集任务列表
func (l *LocalDatasetService) ListDatasetIOJobs(ctx context.Context, req *dataset.ListDatasetIOJobsRequest, callOptions ...callopt.Option) (*dataset.ListDatasetIOJobsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceListDatasetIOJobsArgs)
		result := out.(*dataset.DatasetServiceListDatasetIOJobsResult)
		resp, err := l.impl.ListDatasetIOJobs(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceListDatasetIOJobsArgs{Req: req}
	result := &dataset.DatasetServiceListDatasetIOJobsResult{}
	ctx = l.injectRPCInfo(ctx, "ListDatasetIOJobs")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// CreateDatasetVersion
/* Dataset Version */
// 生成一个新版本
func (l *LocalDatasetService) CreateDatasetVersion(ctx context.Context, req *dataset.CreateDatasetVersionRequest, callOptions ...callopt.Option) (*dataset.CreateDatasetVersionResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceCreateDatasetVersionArgs)
		result := out.(*dataset.DatasetServiceCreateDatasetVersionResult)
		resp, err := l.impl.CreateDatasetVersion(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceCreateDatasetVersionArgs{Req: req}
	result := &dataset.DatasetServiceCreateDatasetVersionResult{}
	ctx = l.injectRPCInfo(ctx, "CreateDatasetVersion")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// ListDatasetVersions
// 版本列表
func (l *LocalDatasetService) ListDatasetVersions(ctx context.Context, req *dataset.ListDatasetVersionsRequest, callOptions ...callopt.Option) (*dataset.ListDatasetVersionsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceListDatasetVersionsArgs)
		result := out.(*dataset.DatasetServiceListDatasetVersionsResult)
		resp, err := l.impl.ListDatasetVersions(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceListDatasetVersionsArgs{Req: req}
	result := &dataset.DatasetServiceListDatasetVersionsResult{}
	ctx = l.injectRPCInfo(ctx, "ListDatasetVersions")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// GetDatasetVersion
// 获取指定版本的数据集详情
func (l *LocalDatasetService) GetDatasetVersion(ctx context.Context, req *dataset.GetDatasetVersionRequest, callOptions ...callopt.Option) (*dataset.GetDatasetVersionResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceGetDatasetVersionArgs)
		result := out.(*dataset.DatasetServiceGetDatasetVersionResult)
		resp, err := l.impl.GetDatasetVersion(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceGetDatasetVersionArgs{Req: req}
	result := &dataset.DatasetServiceGetDatasetVersionResult{}
	ctx = l.injectRPCInfo(ctx, "GetDatasetVersion")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// BatchGetDatasetVersions
// 批量获取指定版本的数据集详情
func (l *LocalDatasetService) BatchGetDatasetVersions(ctx context.Context, req *dataset.BatchGetDatasetVersionsRequest, callOptions ...callopt.Option) (*dataset.BatchGetDatasetVersionsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceBatchGetDatasetVersionsArgs)
		result := out.(*dataset.DatasetServiceBatchGetDatasetVersionsResult)
		resp, err := l.impl.BatchGetDatasetVersions(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceBatchGetDatasetVersionsArgs{Req: req}
	result := &dataset.DatasetServiceBatchGetDatasetVersionsResult{}
	ctx = l.injectRPCInfo(ctx, "BatchGetDatasetVersions")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// GetDatasetSchema
/* Dataset Schema */
// 获取数据集当前的 schema
func (l *LocalDatasetService) GetDatasetSchema(ctx context.Context, req *dataset.GetDatasetSchemaRequest, callOptions ...callopt.Option) (*dataset.GetDatasetSchemaResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceGetDatasetSchemaArgs)
		result := out.(*dataset.DatasetServiceGetDatasetSchemaResult)
		resp, err := l.impl.GetDatasetSchema(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceGetDatasetSchemaArgs{Req: req}
	result := &dataset.DatasetServiceGetDatasetSchemaResult{}
	ctx = l.injectRPCInfo(ctx, "GetDatasetSchema")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// UpdateDatasetSchema
// 覆盖更新 schema
func (l *LocalDatasetService) UpdateDatasetSchema(ctx context.Context, req *dataset.UpdateDatasetSchemaRequest, callOptions ...callopt.Option) (*dataset.UpdateDatasetSchemaResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceUpdateDatasetSchemaArgs)
		result := out.(*dataset.DatasetServiceUpdateDatasetSchemaResult)
		resp, err := l.impl.UpdateDatasetSchema(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceUpdateDatasetSchemaArgs{Req: req}
	result := &dataset.DatasetServiceUpdateDatasetSchemaResult{}
	ctx = l.injectRPCInfo(ctx, "UpdateDatasetSchema")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// BatchCreateDatasetItems
/* Dataset Item */
// 批量新增数据
func (l *LocalDatasetService) BatchCreateDatasetItems(ctx context.Context, req *dataset.BatchCreateDatasetItemsRequest, callOptions ...callopt.Option) (*dataset.BatchCreateDatasetItemsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceBatchCreateDatasetItemsArgs)
		result := out.(*dataset.DatasetServiceBatchCreateDatasetItemsResult)
		resp, err := l.impl.BatchCreateDatasetItems(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceBatchCreateDatasetItemsArgs{Req: req}
	result := &dataset.DatasetServiceBatchCreateDatasetItemsResult{}
	ctx = l.injectRPCInfo(ctx, "BatchCreateDatasetItems")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// UpdateDatasetItem
// 更新数据
func (l *LocalDatasetService) UpdateDatasetItem(ctx context.Context, req *dataset.UpdateDatasetItemRequest, callOptions ...callopt.Option) (*dataset.UpdateDatasetItemResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceUpdateDatasetItemArgs)
		result := out.(*dataset.DatasetServiceUpdateDatasetItemResult)
		resp, err := l.impl.UpdateDatasetItem(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceUpdateDatasetItemArgs{Req: req}
	result := &dataset.DatasetServiceUpdateDatasetItemResult{}
	ctx = l.injectRPCInfo(ctx, "UpdateDatasetItem")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// DeleteDatasetItem
// 删除数据
func (l *LocalDatasetService) DeleteDatasetItem(ctx context.Context, req *dataset.DeleteDatasetItemRequest, callOptions ...callopt.Option) (*dataset.DeleteDatasetItemResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceDeleteDatasetItemArgs)
		result := out.(*dataset.DatasetServiceDeleteDatasetItemResult)
		resp, err := l.impl.DeleteDatasetItem(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceDeleteDatasetItemArgs{Req: req}
	result := &dataset.DatasetServiceDeleteDatasetItemResult{}
	ctx = l.injectRPCInfo(ctx, "DeleteDatasetItem")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// BatchDeleteDatasetItems
// 批量删除数据
func (l *LocalDatasetService) BatchDeleteDatasetItems(ctx context.Context, req *dataset.BatchDeleteDatasetItemsRequest, callOptions ...callopt.Option) (*dataset.BatchDeleteDatasetItemsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceBatchDeleteDatasetItemsArgs)
		result := out.(*dataset.DatasetServiceBatchDeleteDatasetItemsResult)
		resp, err := l.impl.BatchDeleteDatasetItems(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceBatchDeleteDatasetItemsArgs{Req: req}
	result := &dataset.DatasetServiceBatchDeleteDatasetItemsResult{}
	ctx = l.injectRPCInfo(ctx, "BatchDeleteDatasetItems")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// ListDatasetItems
// 分页查询当前数据
func (l *LocalDatasetService) ListDatasetItems(ctx context.Context, req *dataset.ListDatasetItemsRequest, callOptions ...callopt.Option) (*dataset.ListDatasetItemsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceListDatasetItemsArgs)
		result := out.(*dataset.DatasetServiceListDatasetItemsResult)
		resp, err := l.impl.ListDatasetItems(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceListDatasetItemsArgs{Req: req}
	result := &dataset.DatasetServiceListDatasetItemsResult{}
	ctx = l.injectRPCInfo(ctx, "ListDatasetItems")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// ListDatasetItemsByVersion
// 分页查询指定版本的数据
func (l *LocalDatasetService) ListDatasetItemsByVersion(ctx context.Context, req *dataset.ListDatasetItemsByVersionRequest, callOptions ...callopt.Option) (*dataset.ListDatasetItemsByVersionResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceListDatasetItemsByVersionArgs)
		result := out.(*dataset.DatasetServiceListDatasetItemsByVersionResult)
		resp, err := l.impl.ListDatasetItemsByVersion(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceListDatasetItemsByVersionArgs{Req: req}
	result := &dataset.DatasetServiceListDatasetItemsByVersionResult{}
	ctx = l.injectRPCInfo(ctx, "ListDatasetItemsByVersion")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// GetDatasetItem
// 获取一行数据
func (l *LocalDatasetService) GetDatasetItem(ctx context.Context, req *dataset.GetDatasetItemRequest, callOptions ...callopt.Option) (*dataset.GetDatasetItemResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceGetDatasetItemArgs)
		result := out.(*dataset.DatasetServiceGetDatasetItemResult)
		resp, err := l.impl.GetDatasetItem(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceGetDatasetItemArgs{Req: req}
	result := &dataset.DatasetServiceGetDatasetItemResult{}
	ctx = l.injectRPCInfo(ctx, "GetDatasetItem")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// BatchGetDatasetItems
// 批量获取数据
func (l *LocalDatasetService) BatchGetDatasetItems(ctx context.Context, req *dataset.BatchGetDatasetItemsRequest, callOptions ...callopt.Option) (*dataset.BatchGetDatasetItemsResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceBatchGetDatasetItemsArgs)
		result := out.(*dataset.DatasetServiceBatchGetDatasetItemsResult)
		resp, err := l.impl.BatchGetDatasetItems(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceBatchGetDatasetItemsArgs{Req: req}
	result := &dataset.DatasetServiceBatchGetDatasetItemsResult{}
	ctx = l.injectRPCInfo(ctx, "BatchGetDatasetItems")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// BatchGetDatasetItemsByVersion
// 批量获取指定版本的数据
func (l *LocalDatasetService) BatchGetDatasetItemsByVersion(ctx context.Context, req *dataset.BatchGetDatasetItemsByVersionRequest, callOptions ...callopt.Option) (*dataset.BatchGetDatasetItemsByVersionResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceBatchGetDatasetItemsByVersionArgs)
		result := out.(*dataset.DatasetServiceBatchGetDatasetItemsByVersionResult)
		resp, err := l.impl.BatchGetDatasetItemsByVersion(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceBatchGetDatasetItemsByVersionArgs{Req: req}
	result := &dataset.DatasetServiceBatchGetDatasetItemsByVersionResult{}
	ctx = l.injectRPCInfo(ctx, "BatchGetDatasetItemsByVersion")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

// ClearDatasetItem
// 清除(草稿)数据项
func (l *LocalDatasetService) ClearDatasetItem(ctx context.Context, req *dataset.ClearDatasetItemRequest, callOptions ...callopt.Option) (*dataset.ClearDatasetItemResponse, error) {
	chain := l.mds(func(ctx context.Context, in, out interface{}) error {
		arg := in.(*dataset.DatasetServiceClearDatasetItemArgs)
		result := out.(*dataset.DatasetServiceClearDatasetItemResult)
		resp, err := l.impl.ClearDatasetItem(ctx, arg.Req)
		if err != nil {
			return err
		}
		result.SetSuccess(resp)
		return nil
	})

	arg := &dataset.DatasetServiceClearDatasetItemArgs{Req: req}
	result := &dataset.DatasetServiceClearDatasetItemResult{}
	ctx = l.injectRPCInfo(ctx, "ClearDatasetItem")
	if err := chain(ctx, arg, result); err != nil {
		return nil, err
	}
	return result.GetSuccess(), nil
}

func (l *LocalDatasetService) injectRPCInfo(ctx context.Context, method string) context.Context {
	rpcStats := rpcinfo.AsMutableRPCStats(rpcinfo.NewRPCStats())
	ri := rpcinfo.NewRPCInfo(
		rpcinfo.NewEndpointInfo("DatasetService", method, nil, nil),
		rpcinfo.NewEndpointInfo("DatasetService", method, nil, nil),
		rpcinfo.NewServerInvocation(),
		nil,
		rpcStats.ImmutableView(),
	)
	return rpcinfo.NewCtxWithRPCInfo(ctx, ri)
}
