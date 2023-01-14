package server

import (
	"encoding/json"
	"golang.org/x/net/context"
	"gorm_pro/blog-service/pkg/errcode"
	"gorm_pro/tag-service/pkg/bapi"
	pb "gorm_pro/tag-service/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8080")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError
	}

	return &tagList, nil
}
