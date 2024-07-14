package service

import (
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/repository"
)

type IRoomObjectService interface {
	Calculate(roomID string) error
}

type RoomObjectService struct {
	roomObjRepo repository.IRoomObjectRepository
}

func NewRoomObjectService(roomObjRepo repository.IRoomObjectRepository) IRoomObjectService {
	return &RoomObjectService{
		roomObjRepo: roomObjRepo,
	}
}

func (s *RoomObjectService) Calculate(roomID string) error {
	return nil
}
