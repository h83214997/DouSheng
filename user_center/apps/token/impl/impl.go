// @Author: Ciusyan 2023/2/6
package impl

import (
	"context"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/Go-To-Byte/DouSheng/dou_common/ioc"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
	"github.com/Go-To-Byte/DouSheng/user_center/conf"
)

var (
	impl = &tokenServiceImpl{}
)

type tokenServiceImpl struct {
	col *mongo.Collection
	log logger.Logger

	token.UnimplementedServiceServer
}

func (s *tokenServiceImpl) Init() error {
	db := conf.C().Mongodb.GetDB()

	// 定义 document 名称
	s.col = db.Collection(token.AppName)
	s.log = zap.L().Named(token.AppName)

	// 定义一个索引
	index := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "refresh_token", Value: bsonx.Int32(-1)},
			},
		},
	}

	_, err := s.col.Indexes().CreateMany(context.Background(), index)

	return err
}

func (s *tokenServiceImpl) Name() string {
	return token.AppName
}

func (s *tokenServiceImpl) Registry(server *grpc.Server) {
	token.RegisterServiceServer(server, impl)
}

func init() {
	ioc.GrpcDI(impl)
}
