package physics

import (
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

// 速度の計算 v = (x2 - x1) / dt
func CalculateVelocity(prevPos, currPos entity.Vector3, deltaTime float64) *entity.Vector3 {
	return &entity.Vector3{
		X: (currPos.X - prevPos.X) / deltaTime,
		Y: (currPos.Y - prevPos.Y) / deltaTime,
		Z: (currPos.Z - prevPos.Z) / deltaTime,
	}
}

// 加速度の計算 a = (v2 - v1) / dt
func CalculateAcceleration(prevVel, currVel entity.Vector3, deltaTime float64) *entity.Vector3 {
	return &entity.Vector3{
		X: (currVel.X - prevVel.X) / deltaTime,
		Y: (currVel.Y - prevVel.Y) / deltaTime,
		Z: (currVel.Z - prevVel.Z) / deltaTime,
	}
}

// 力の計算 f = ma
func CalculateForce(mass float64, acceleration entity.Vector3) *entity.Vector3 {
	return &entity.Vector3{
		X: mass * acceleration.X,
		Y: mass * acceleration.Y,
		Z: mass * acceleration.Z,
	}
}
