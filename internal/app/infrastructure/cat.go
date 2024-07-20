package infrastructure

import (
	"context"
	"sync"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/repository"
	domainService "github.com/K-Kizuku/kotatuneko-backend/internal/domain/service"
)

type Cat struct {
	cat   domainService.ICatService
	world map[string][]*entity.Object
	mux   sync.RWMutex
}

func NewCat(cat domainService.ICatService) repository.ICatRepository {
	return &Cat{
		cat:   cat,
		world: make(map[string][]*entity.Object),
		mux:   sync.RWMutex{},
	}
}

func (c *Cat) Calculate(ctx context.Context, roomID string, hand *entity.Hand) error {
	obj, err := c.cat.Do(ctx, hand)
	if err != nil {
		return err
	}
	if err := c.set(roomID, obj); err != nil {
		return err
	}
	return nil
}

func (c *Cat) Get(ctx context.Context, roomID string) ([]*entity.Object, error) {
	return c.cat.Get(ctx)
}

func (c *Cat) Init(ctx context.Context, roomID string) error {
	obj := make([]*entity.Object, 20)
	c.set(roomID, obj)
	return c.cat.Init(ctx)
}

func (c *Cat) set(roomID string, objects []*entity.Object) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.world[roomID] = append(c.world[roomID], objects...)
	return nil
}

func (c *Cat) get(roomID string) ([]*entity.Object, error) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	originalSlice := c.world[roomID]
	slice := make([]*entity.Object, len(originalSlice))
	for i, obj := range originalSlice {
		slice[i] = obj.DeepCopy()
	}
	return slice, nil
}
