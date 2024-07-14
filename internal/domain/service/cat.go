package service

import "github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"

type ICatService interface {
	Do(hand *entity.Hand, roomID string) (*[]entity.Object, error)
	Get(roomID string) (*[]entity.Object, error)
	Init(roomID string) error
}
