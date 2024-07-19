package service

import (
	"context"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

type ICatService interface {
	Do(ctx context.Context, hand *entity.Hand) error
	Get(ctx context.Context) error
	Init(ctx context.Context) error
}
