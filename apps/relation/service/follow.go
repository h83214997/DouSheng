// Author: BeYoung
// Date: 2023/1/30 17:30
// Software: GoLand

package service

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao"
	"github.com/Go-To-Byte/DouSheng/apps/relation/dao/dal/model"
	"github.com/Go-To-Byte/DouSheng/apps/relation/models"
	"github.com/Go-To-Byte/DouSheng/apps/relation/proto"
	"go.uber.org/zap"
)

func (r *Relation) Follow(ctx context.Context, req *proto.FollowRequest) (*proto.FollowResponse, error) {
	relation := model.Relation{
		ID:       models.Node.Generate().Int64(),
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
		Flag:     1,
	}
	if err := dao.Add(relation); err != nil {
		return &proto.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed add follow",
		}, err
	}
	return &proto.FollowResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func (r *Relation) FollowDelete(ctx context.Context, req *proto.FollowRequest) (*proto.FollowResponse, error) {
	relation := model.Relation{
		ID:       0,
		UserID:   req.UserId,
		ToUserID: req.ToUserId,
		Flag:     0,
	}

	if err := dao.Delete(relation); err != nil {
		return &proto.FollowResponse{
			StatusCode: 1,
			StatusMsg:  "failed delete follow",
		}, err
	}
	return &proto.FollowResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func (r *Relation) FollowList(ctx context.Context, req *proto.FollowListRequest) (*proto.FollowListResponse, error) {
	relation := model.Relation{
		ID:       0,
		UserID:   req.UserId,
		ToUserID: 0,
		Flag:     1,
	}
	result := dao.RelationFindByUserID(relation)
	list := make([]int64, len(result))
	for i := range result {
		list = append(list, result[i].ToUserID)
	}
	zap.S().Debugf("id:%v ==> len(follow):%v", req.UserId, len(list))

	return &proto.FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserList:   list,
	}, nil
}
