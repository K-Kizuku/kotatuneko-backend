package service

import (
	"context"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

type ICatService interface {
	Do(ctx context.Context, hand *entity.Hand, roomID string) (*[]entity.Object, error)
	Get(ctx context.Context, roomID string) (*[]entity.Object, error)
	Init(ctx context.Context, roomID string) error
}
