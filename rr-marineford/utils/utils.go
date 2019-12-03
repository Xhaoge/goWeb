package utils

import (
	"context"
	"math/big"
	"time"
	Nrand "crypto/rand"

	conrectx "rr-factory.gloryholiday.com/yuetu/golang-core/context"
	"rr-factory.gloryholiday.com/yuetu/yuetu/golang-core/location"
	pb "rr-factory.gloryholiday.com/yuetu/marineford/proto/marineford"
)

func NewSimpleUpdateResponse(id string) * pb.SimpleUpdateResponse {
	return &pb.SimpleUpdateResponse{Id: id}
}

func NewFailedSimpleUpdateResponse(id string,msg string) *pb.SimpleUpdateResponse {
	return &pb.SimpleUpdateResponse{Id:id, Code:pb.ResponseStatusCode_FAILED,ErrMsg:msg}
}
