package main

import "fmt"

func main() {
	testMinParticleIndex := findPart1Particle(testInput)
	fmt.Printf("Test part 1 - the particle that will stay closest to 0,0,0 is %d\n", testMinParticleIndex)

	minParticleIndex := findPart1Particle(input)
	fmt.Printf("Part 1 - the particle that will stay closest to 0,0,0 is %d\n", minParticleIndex)

	testWorld := world{testInputPart2}
	for !testWorld.isDivergent() {
		testWorld.tick()
	}
	fmt.Printf("Test part 2 - after all collisions there are %d particles left\n", len(testWorld.particles))

	part2World := world{input}
	for !part2World.isDivergent() {
		part2World.tick()
	}
	fmt.Printf("Part 2 - after all collisions there are %d particles left\n", len(part2World.particles))

}
func findPart1Particle(particles []particle) int {
	if len(particles) == 0 {
		return -1
	}
	minParticleIndex := 0
	minAccel := particles[minParticleIndex].a.manhattanMagnitude()
	for i, p := range particles {
		if p.a.manhattanMagnitude() < minAccel {
			minAccel = p.a.manhattanMagnitude()
			minParticleIndex = i
		}
	}
	return minParticleIndex
}
