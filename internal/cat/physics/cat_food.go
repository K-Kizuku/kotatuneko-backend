package physics

import (
	"math"

	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/constants"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

// ベクトルの大きさを計算
func Magnitude(v *entity.Vector3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// 運動エネルギーの計算
func KineticEnergy(v *entity.Vector3) float64 {
	speed := Magnitude(v)
	return 0.5 * constants.BlockMass * speed * speed
}
