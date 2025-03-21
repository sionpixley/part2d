package part2d

type Particle struct {
	Position Vector `json:"position"`
	Velocity Vector `json:"velocity"`
	Mass     int64  `json:"mass"`
}

func (particle *Particle) ApplyForce(force *Vector, dt int64) {
	acceleration := Vector{X: force.X / particle.Mass, Y: force.Y / particle.Mass}
	particle.Velocity.X += acceleration.X * dt
	particle.Velocity.Y += acceleration.Y * dt
	particle.Position.X += particle.Velocity.X * dt
	particle.Position.Y += particle.Velocity.Y * dt
}

type Vector struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}
