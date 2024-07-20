package service

import (
	"context"
	"fmt"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/repository"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/service"
	"github.com/K-Kizuku/kotatuneko-protobuf/gen/game/resources"
	"github.com/K-Kizuku/kotatuneko-protobuf/gen/game/rpc"
	"google.golang.org/protobuf/proto"
)

type IRoomObjectService interface {
	Calculate(ctx context.Context, senderID, roomID string, hand *entity.Hand) error
}

type RoomObjectService struct {
	roomObjRepo repository.IRoomObjectRepository
	msgSender   service.IMessageSender
	catRepo     repository.ICatRepository
}

func NewRoomObjectService(roomObjRepo repository.IRoomObjectRepository, msgSender service.IMessageSender, catRepo repository.ICatRepository) IRoomObjectService {
	return &RoomObjectService{
		roomObjRepo: roomObjRepo,
		msgSender:   msgSender,
		catRepo:     catRepo,
	}
}

func (s *RoomObjectService) Calculate(ctx context.Context, senderID, roomID string, hand *entity.Hand) error {
	fmt.Println("before calculate")

	if err := s.catRepo.Calculate(ctx, roomID, hand); err != nil {
		return err
	}
	fmt.Println("calculate")
	// ここで1フレーム待つ？
	objs, err := s.catRepo.Get(ctx, roomID)
	fmt.Println("get", objs)
	resourceObj := make([]*resources.Object, 0, len(objs))
	for _, obj := range objs {
		resourceObj = append(resourceObj, &resources.Object{
			ObjectId: obj.ID,
			Layer:    obj.Layer,
			Kinds:    resources.ObjectKind(obj.Kinds),
			State:    resources.ObjectState(obj.State),
			Position: &resources.Vector3{
				X: obj.Position.X,
				Y: obj.Position.Y,
				Z: obj.Position.Z,
			},
			Size: &resources.Vector3{
				X: obj.Size.X,
				Y: obj.Size.Y,
				Z: obj.Size.Z,
			},
		})
	}
	fmt.Println("resourceObj", resourceObj)

	physics := rpc.PhysicsResponse{
		RoomId:   roomID,
		SenderId: senderID,
		Objects:  resourceObj,
	}

	b, err := proto.Marshal(&physics)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("send physics")
	if err := s.msgSender.Broadcast(ctx, roomID, b); err != nil {
		return err
	}
	return nil
}
