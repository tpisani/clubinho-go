package main

import (
	"fmt"
)

type BallThrower interface {
	ThrowBall() (int, error)
}

type Person struct {
	thrower BallThrower
}

type Arm struct {
	distance int
}

func (a Arm) ThrowBall() (int, error) {
	return a.distance, nil
}
