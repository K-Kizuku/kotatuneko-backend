package cat

import (
	"context"

	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/service"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
	domainService "github.com/K-Kizuku/kotatuneko-backend/internal/domain/service"
)

type Cat struct {
	hs service.IHandService
	os service.IObjectService
}

func New(hs service.IHandService, os service.IObjectService) domainService.ICatService {
	return &Cat{
		hs: hs,
		os: os,
	}
}

// 1フレームの間に行う処理
func (c *Cat) Do(ctx context.Context, hand *entity.Hand) ([]*entity.Object, error) {
	// nikukyu := c.hs.TransferHandToNikukyu(hand)
	// // 手の当たり判定を求める
	// if id := c.hs.CollideWithObj(nikukyu); id != nil {
	// 	// オブジェクトに与える影響を計算
	// 	force := c.hs.CalculateHandForce(nikukyu)
	// 	c.hs.ApplyForceToObj(*id, force)
	// }
	// for sourceObjID, targetObjID := range c.os.CollideWithObj() {
	// 	// オブジェクトに与える影響を計算
	// 	for _, id := range targetObjID {
	// 		c.os.ApplyForceToObj(sourceObjID, id)
	// 	}
	// }
	// オブジェクトの当たり判定を求める
	// trueの場合、オブジェクトに与える影響を計算
	// 状態の更新
	obj := make([]*entity.Object, 20)
	for _, nekojarashi := range c.os.GetObjectsSlice() {
		obj = append(obj, &entity.Object{
			ID:       nekojarashi.ID,
			Layer:    nekojarashi.Layer,
			Kinds:    nekojarashi.Kinds,
			State:    nekojarashi.State,
			Position: nekojarashi.Position,
			Size:     nekojarashi.Size,
		})
	}

	return obj, nil
}

func (c *Cat) Get(ctx context.Context) ([]*entity.Object, error) {
	obj := make([]*entity.Object, 20)
	for _, nekojarashi := range c.os.GetObjectsSlice() {
		obj = append(obj, &entity.Object{
			ID:       nekojarashi.ID,
			Layer:    nekojarashi.Layer,
			Kinds:    nekojarashi.Kinds,
			State:    nekojarashi.State,
			Position: nekojarashi.Position,
			Size:     nekojarashi.Size,
		})
	}
	return obj, nil
}

func (c *Cat) Init(ctx context.Context) error {

	return nil
}
