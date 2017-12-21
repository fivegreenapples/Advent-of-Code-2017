package main

import "sort"

type position [3]int
type velocity [3]int
type acceleration [3]int

func (p *position) applyVelocityForOneTick(v velocity) {
	p[0] += v[0]
	p[1] += v[1]
	p[2] += v[2]
}

func (v *velocity) applyAccelerationForOneTick(a acceleration) {
	v[0] += a[0]
	v[1] += a[1]
	v[2] += a[2]
}

type particle struct {
	p position
	v velocity
	a acceleration
}

func (p *particle) tick() {
	p.v.applyAccelerationForOneTick(p.a)
	p.p.applyVelocityForOneTick(p.v)
}

func (p position) manhattanMagnitude() int {
	return manhattanMagnitude(p[0], p[1], p[2])
}
func (v velocity) manhattanMagnitude() int {
	return manhattanMagnitude(v[0], v[1], v[2])
}
func (a acceleration) manhattanMagnitude() int {
	return manhattanMagnitude(a[0], a[1], a[2])
}

func manhattanMagnitude(x, y, z int) int {
	absX := x
	if x < 0 {
		absX = -x
	}
	absY := y
	if y < 0 {
		absY = -y
	}
	absZ := z
	if z < 0 {
		absZ = -z
	}
	return absX + absY + absZ
}

type world struct {
	particles []particle
}

func (w *world) tick() {
	// records first particle to be in this position
	particlePositions := map[position]int{}
	// records particles we wish to keep for the next tick
	particlesToKeep := map[int]bool{}
	// if we find a particle that moves into a position that a previous particle has
	// already got to, then we ignore the new particle and unmark the first particle as one we
	// want to keep

	for i := range w.particles {
		w.particles[i].tick()
		if particleIndex, seen := particlePositions[w.particles[i].p]; seen {
			delete(particlesToKeep, particleIndex)
		} else {
			particlePositions[w.particles[i].p] = i
			particlesToKeep[i] = true
		}
	}

	particlesNew := []particle{}
	for keepParticle := range particlesToKeep {
		particlesNew = append(particlesNew, w.particles[keepParticle])
	}
	w.particles = particlesNew
}

func (w world) isDivergent() bool {
	// Do all the particles diverge from each other?
	// We do this independently for each axis.

	if len(w.particles) <= 1 {
		return true
	}

	return w.isDivergentInAxis(0) || w.isDivergentInAxis(1) || w.isDivergentInAxis(2)

}

func (w world) isDivergentInAxis(axis int) bool {

	if len(w.particles) <= 1 {
		return true
	}

	// we sort by position and check whether the accelerations are in the same order
	// when accelerations are the same, we check whether velocities are in the correct order
	// and then whether positions are different

	sort.Slice(w.particles, func(i, j int) bool {
		return w.particles[i].p[axis] < w.particles[j].p[axis]
	})
	previousParticle := w.particles[0]
	for _, p := range w.particles[1:] {
		if p.a[axis] < previousParticle.a[axis] {
			return false
		} else if p.a[axis] == previousParticle.a[axis] {
			if p.v[axis] < previousParticle.v[axis] {
				return false
			} else if p.v[axis] == previousParticle.v[axis] {
				if p.p[axis] == previousParticle.p[axis] {
					return false
				}
			}
		}
		previousParticle = p
	}

	return true
}
