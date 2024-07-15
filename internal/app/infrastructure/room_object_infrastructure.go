package infrastructure

import (
	"context"

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

func (r *RoomObjectRepository) Get(ctx context.Context, roomID string) (*[]entity.Object, error) {
	data, exist := r.cache.Get(ctx, roomID)
	if !exist {
		return &[]entity.Object{}, nil
	}
	resp, ok := data.(*[]entity.Object)
	if !ok {
		return &[]entity.Object{}, nil
	}
	return resp, nil
}

func (r *RoomObjectRepository) Set(ctx context.Context, roomID string, objects *[]entity.Object) error {
	if !r.cache.Set(ctx, roomID, objects) {
		// TODO: error handling
		return nil
	}
	return nil
}

func (r *RoomObjectRepository) Resister(ctx context.Context, id string) error {
	return nil
}

func (r *RoomObjectRepository) Unregister(ctx context.Context, id string) error {
	r.cache.Delete(ctx, id)
	return nil
}
