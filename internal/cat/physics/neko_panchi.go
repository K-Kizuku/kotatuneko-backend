package physics

import (
	"math"

	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/constants"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

// 当たり判定
func IsColliding(center1, center2 entity.Vector3) bool {
	dx := math.Abs(center1.X - center2.X)
	dy := math.Abs(center1.Y - center2.Y)
	dz := math.Abs(center1.Z - center2.Z)

	return dx <= constants.BlockSizeX && dy <= constants.BlockSizeY && dz <= constants.BlockSizeZ
}

// 衝突後の速度の更新
func CollidedVelocity(obj1, obj2 *entity.Nekojarashi) {
	obj1.Velocity = entity.Vector3{
		X: (obj1.Velocity.X*(1+constants.Restitution) + obj2.Velocity.X*(1-constants.Restitution)) / 2,
		Y: (obj1.Velocity.Y*(1+constants.Restitution) + obj2.Velocity.Y*(1-constants.Restitution)) / 2,
		Z: (obj1.Velocity.Z*(1+constants.Restitution) + obj2.Velocity.Z*(1-constants.Restitution)) / 2,
	}
	obj2.Velocity = entity.Vector3{
		X: (obj2.Velocity.X*(1+constants.Restitution) + obj1.Velocity.X*(1-constants.Restitution)) / 2,
		Y: (obj2.Velocity.Y*(1+constants.Restitution) + obj1.Velocity.Y*(1-constants.Restitution)) / 2,
		Z: (obj2.Velocity.Z*(1+constants.Restitution) + obj1.Velocity.Z*(1-constants.Restitution)) / 2,
	}
}

// 衝突後の運動エネルギーの計算
// func EnergyAfterCollisionAndFriction(v1, v2 *entity.Vector3) (float64, float64) {
// 	v1After, v2After := CollidedVelocity(v1, v2)

// 	v1After.X *= (1 - constants.Friction)
// 	v1After.Y *= (1 - constants.Friction)
// 	v1After.Z *= (1 - constants.Friction)
// 	v2After.X *= (1 - constants.Friction)
// 	v2After.Y *= (1 - constants.Friction)
// 	v2After.Z *= (1 - constants.Friction)

// 	e1 := KineticEnergy(v1After)
// 	e2 := KineticEnergy(v2After)

// 	return e1, e2
// }
