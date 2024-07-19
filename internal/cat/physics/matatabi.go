package physics

import (
	"math"

	"github.com/K-Kizuku/kotatuneko-backend/internal/cat/constants"
	"github.com/K-Kizuku/kotatuneko-backend/internal/domain/entity"
)

// 力による速度の更新
func ApplyForce(obj *entity.Nekojarashi, force *entity.Vector3) {
	acceleration := entity.Vector3{
		X: force.X / obj.Mass,
		Y: force.Y / obj.Mass,
		Z: force.Z / obj.Mass,
	}
	obj.Velocity.X += acceleration.X * constants.TimeStep
	obj.Velocity.Y += acceleration.Y * constants.TimeStep
	obj.Velocity.Z += acceleration.Z * constants.TimeStep
}

// 重力による速度の更新
func ApplyGravity(obj *entity.Nekojarashi) {
	force := entity.Vector3{
		X: 0,
		Y: 0,
		Z: -obj.Mass * constants.Gravity,
	}
	ApplyForce(obj, &force)
}

// 摩擦力による速度の更新
func ApplyFriction(obj *entity.Nekojarashi) {
	normalForce := obj.Mass * constants.Gravity
	frictionForce := constants.KineticFriction * normalForce
	speed := math.Sqrt(obj.Velocity.X*obj.Velocity.X + obj.Velocity.Y*obj.Velocity.Y + obj.Velocity.Z*obj.Velocity.Z)
	if speed > 0 {
		friction := entity.Vector3{
			X: -obj.Velocity.X / speed * frictionForce,
			Y: -obj.Velocity.Y / speed * frictionForce,
			Z: -obj.Velocity.Z / speed * frictionForce,
		}
		obj.Velocity.X += friction.X * constants.TimeStep
		obj.Velocity.Y += friction.Y * constants.TimeStep
		obj.Velocity.Z += friction.Z * constants.TimeStep
	}
}

// 位置の更新
func UpdatePosition(obj *entity.Nekojarashi) {
	obj.Position.X += obj.Velocity.X * constants.TimeStep
	obj.Position.Y += obj.Velocity.Y * constants.TimeStep
	obj.Position.Z += obj.Velocity.Z * constants.TimeStep
}
