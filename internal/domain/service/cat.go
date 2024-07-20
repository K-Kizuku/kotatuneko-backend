package service

import (
	"context"

	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

type ICatService interface {
	Do(ctx context.Context, hand *entity.Hand) ([]*entity.Object, error) // 1フレームの間に行う処理
	Get(ctx context.Context) ([]*entity.Object, error)                   // Objectの座標を返す
	Init(ctx context.Context) error                                      // 初期化処理
}
