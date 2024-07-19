package entity

type Nikukyu struct {
	UserID             string
	State              HandState
	CenterPosition     Vector3
	ActionPosition     Vector3
	Velocity           Vector3
	PrevCenterPosition Vector3
	PrevActionPosition Vector3
	PrevVelocity       Vector3
}

func (n *Nikukyu) DeepCopy() *Nikukyu {
	return &Nikukyu{
		UserID: n.UserID,
		State:  n.State,
		CenterPosition: Vector3{
			X: n.CenterPosition.X,
			Y: n.CenterPosition.Y,
			Z: n.CenterPosition.Z,
		},
		ActionPosition: Vector3{
			X: n.ActionPosition.X,
			Y: n.ActionPosition.Y,
			Z: n.ActionPosition.Z,
		},
		Velocity: Vector3{
			X: n.Velocity.X,
			Y: n.Velocity.Y,
			Z: n.Velocity.Z,
		},
		PrevCenterPosition: Vector3{
			X: n.PrevCenterPosition.X,
			Y: n.PrevCenterPosition.Y,
			Z: n.PrevCenterPosition.Z,
		},
		PrevActionPosition: Vector3{
			X: n.PrevActionPosition.X,
			Y: n.PrevActionPosition.Y,
			Z: n.PrevActionPosition.Z,
		},
		PrevVelocity: Vector3{
			X: n.PrevVelocity.X,
			Y: n.PrevVelocity.Y,
			Z: n.PrevVelocity.Z,
		},
	}
}
