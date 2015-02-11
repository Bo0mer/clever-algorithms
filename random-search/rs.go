package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func F(args []float64) float64 {
	sum := 0.0
	for _, arg := range args {
		sum += (arg * arg)
	}
	return sum
}

type Constraint struct {
	Min float64
	Max float64
}

func RandomVector(constraints []Constraint) []float64 {
	vector := make([]float64, len(constraints))
	for i, constraint := range constraints {
		vector[i] = constraint.Min + ((constraint.Max - constraint.Min) * rand.Float64())
	}
	return vector
}

func RandomSearch(iterations int, constraints []Constraint) []float64 {

	var best []float64

	for iterations >= 0 {
		current := RandomVector(constraints)

		if best == nil || F(current) < F(best) {
			best = current
		}
		iterations--
	}
	return best
}

func main() {
	if len(os.Args) < 2 {
		panic("Iterations CLA not specified.")
	}

	iterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Iterations CLA must be integer.")
	}

	rand.Seed(time.Now().UnixNano())

	constraints := make([]Constraint, 2)

	for i := 0; i < len(constraints); i++ {
		constraints[i].Min = -5.0
		constraints[i].Max = 5.0
	}

	best := RandomSearch(iterations, constraints)

	fmt.Println("Best solution:", best)
	fmt.Println("Function value:", F(best))
}
