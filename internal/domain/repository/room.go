package repository

import "github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"

type IRoomObjectRepository interface {
	Get(roomID string) (*[]entity.Object, error)
	Set(roomID string, objects *[]entity.Object) error
	Resister(id string) error
	Unregister(id string) error
}
