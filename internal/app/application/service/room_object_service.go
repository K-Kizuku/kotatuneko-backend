package service

import (
	"context"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/repository"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/service"
	"github.com/K-Kizuku/kotatuneko-protobuf/gen/game/rpc"
	"google.golang.org/protobuf/proto"
)

type IRoomObjectService interface {
	Calculate(ctx context.Context, roomID string) error
}

type RoomObjectService struct {
	roomObjRepo repository.IRoomObjectRepository
	msgSender   service.IMessageSender
}

func NewRoomObjectService(roomObjRepo repository.IRoomObjectRepository, msgSender service.IMessageSender) IRoomObjectService {
	return &RoomObjectService{
		roomObjRepo: roomObjRepo,
		msgSender:   msgSender,
	}
}

func (s *RoomObjectService) Calculate(ctx context.Context, roomID string) error {
	physics := rpc.PhysicsResponse{
		RoomId: "yamato",
	}

	b, err := proto.Marshal(&physics)
	if err != nil {
		return err
	}

	if err := s.msgSender.Send(ctx, roomID, b); err != nil {
		return err
	}
	return nil
}
