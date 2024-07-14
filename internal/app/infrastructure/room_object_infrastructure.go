package infrastructure

import (
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/repository"
	"github.com/K-Kizuku/kotatuneko-backend/pkg/cache"
)

type RoomObjectRepository struct {
	cache *cache.Client
}

func NewRoomObjectRepository(cache *cache.Client) repository.IRoomObjectRepository {
	return &RoomObjectRepository{
		cache: cache,
	}
}

func (r *RoomObjectRepository) Get(roomID string) (*[]entity.Object, error) {
	return nil, nil
}

func (r *RoomObjectRepository) Set(roomID string, objects *[]entity.Object) error {
	return nil
}

func (r *RoomObjectRepository) Resister(id string) error {
	return nil
}

func (r *RoomObjectRepository) Unregister(id string) error {
	return nil
}
