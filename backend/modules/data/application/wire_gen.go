// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"github.com/google/wire"

	"github.com/coze-dev/cozeloop/backend/infra/db"
	"github.com/coze-dev/cozeloop/backend/infra/external/audit"
	"github.com/coze-dev/cozeloop/backend/infra/fileserver"
	"github.com/coze-dev/cozeloop/backend/infra/idgen"
	"github.com/coze-dev/cozeloop/backend/infra/lock"
	"github.com/coze-dev/cozeloop/backend/infra/mq"
	"github.com/coze-dev/cozeloop/backend/infra/redis"
	"github.com/coze-dev/cozeloop/backend/kitex_gen/coze/loop/foundation/auth/authservice"
	"github.com/coze-dev/cozeloop/backend/modules/data/domain/dataset/service"
	"github.com/coze-dev/cozeloop/backend/modules/data/domain/entity"
	conf2 "github.com/coze-dev/cozeloop/backend/modules/data/infra/conf"
	"github.com/coze-dev/cozeloop/backend/modules/data/infra/mq/producer"
	"github.com/coze-dev/cozeloop/backend/modules/data/infra/repo/dataset"
	"github.com/coze-dev/cozeloop/backend/modules/data/infra/repo/dataset/item_dao"
	"github.com/coze-dev/cozeloop/backend/modules/data/infra/repo/dataset/mysql"
	oss2 "github.com/coze-dev/cozeloop/backend/modules/data/infra/repo/dataset/oss"
	redis2 "github.com/coze-dev/cozeloop/backend/modules/data/infra/repo/dataset/redis"
	"github.com/coze-dev/cozeloop/backend/modules/data/infra/rpc/foundation"
	"github.com/coze-dev/cozeloop/backend/modules/data/infra/vfs/oss"
	"github.com/coze-dev/cozeloop/backend/modules/data/infra/vfs/unionfs"
	"github.com/coze-dev/cozeloop/backend/pkg/conf"
)

// Injectors from wire.go:

func InitDatasetApplication(idgen2 idgen.IIDGenerator, db2 db.Provider, cmdable redis.Cmdable, configFactory conf.IConfigLoaderFactory, mqFactory mq.IFactory, objectStorage fileserver.ObjectStorage, batchObjectStorage fileserver.BatchObjectStorage, auditClient audit.IAuditService, authClient authservice.Client) (IDatasetApplication, error) {
	iAuthProvider := foundation.NewAuthRPCProvider(authClient)
	iDatasetDAO := mysql.NewDatasetDAO(db2, cmdable)
	iSchemaDAO := mysql.NewDatasetSchemaDAO(db2, cmdable)
	datasetDAO := redis2.NewDatasetDAO(cmdable)
	iVersionDAO := mysql.NewDatasetVersionDAO(db2, cmdable)
	versionDAO := redis2.NewVersionDAO(cmdable)
	operationDAO := redis2.NewOperationDAO(cmdable)
	iItemDAO := mysql.NewDatasetItemDAO(db2, cmdable)
	iItemSnapshotDAO := mysql.NewDatasetItemSnapshotDAO(db2, cmdable)
	iioJobDAO := mysql.NewDatasetIOJobDAO(db2, cmdable)
	v := NewItemProviderDAO(batchObjectStorage)
	iDatasetAPI := dataset.NewDatasetRepo(idgen2, db2, iDatasetDAO, iSchemaDAO, datasetDAO, iVersionDAO, versionDAO, operationDAO, iItemDAO, iItemSnapshotDAO, iioJobDAO, v)
	iConfig, err := conf2.NewConfiger(configFactory)
	if err != nil {
		return nil, err
	}
	iDatasetJobPublisher, err := producer.NewDatasetJobPublisher(iConfig, mqFactory)
	if err != nil {
		return nil, err
	}
	client := oss.NewClient(objectStorage)
	iUnionFS := unionfs.NewUnionFS(client)
	iLocker := lock.NewRedisLocker(cmdable)
	serviceIDatasetAPI := service.NewDatasetServiceImpl(db2, idgen2, iDatasetAPI, iConfig, iDatasetJobPublisher, iUnionFS, iLocker)
	iDatasetApplication := NewDatasetApplicationImpl(iAuthProvider, serviceIDatasetAPI, iDatasetAPI, auditClient)
	return iDatasetApplication, nil
}

// wire.go:

var (
	datasetSet = wire.NewSet(
		NewDatasetApplicationImpl, service.NewDatasetServiceImpl, dataset.NewDatasetRepo, mysql.NewDatasetDAO, mysql.NewDatasetItemDAO, mysql.NewDatasetVersionDAO, mysql.NewDatasetItemSnapshotDAO, mysql.NewDatasetSchemaDAO, mysql.NewDatasetIOJobDAO, redis2.NewOperationDAO, redis2.NewDatasetDAO, redis2.NewVersionDAO, conf2.NewConfiger, producer.NewDatasetJobPublisher, foundation.NewAuthRPCProvider, oss.NewClient, unionfs.NewUnionFS, lock.NewRedisLocker, NewItemProviderDAO,
	)
)

func NewItemProviderDAO(batchObjectStorage fileserver.BatchObjectStorage) map[entity.Provider]item_dao.ItemDAO {
	return map[entity.Provider]item_dao.ItemDAO{entity.ProviderS3: oss2.NewDatasetItemDAO(batchObjectStorage)}
}
