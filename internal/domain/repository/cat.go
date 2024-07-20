package repository

import (
	"context"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

type ICatRepository interface {
	Calculate(ctx context.Context, roomID string, hand *entity.Hand) error // 1フレームの間に行う処理
	Get(ctx context.Context, roomID string) ([]*entity.Object, error)      // Objectの座標を返す
	Init(ctx context.Context, roomID string) error
}
