package part2d

type Particle struct {
	Position Vector `json:"position"`
	Velocity Vector `json:"velocity"`
	Mass     int64  `json:"mass"`
}

func (particle *Particle) ApplyForce(dt int64, force Vector) {
	acceleration := Vector{X: force.X / particle.Mass, Y: force.Y / particle.Mass}
	particle.Velocity.X += acceleration.X * dt
	particle.Velocity.Y += acceleration.Y * dt
	particle.Position.X += particle.Velocity.X * dt
	particle.Position.Y += particle.Velocity.Y * dt
}

func (particle *Particle) ApplyForces(dt int64, forces ...Vector) {
	aggregateForce := CombineForces(forces...)
	particle.ApplyForce(dt, aggregateForce)
}

type Vector struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

func CombineForces(forces ...Vector) Vector {
	var aggregateForce Vector
	for _, force := range forces {
		aggregateForce.X += force.X
		aggregateForce.Y += force.Y
	}
	return aggregateForce
}
