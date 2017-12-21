package main

type Jumper interface {
	Jump() string
}

type Runner interface {
	Run() string
}

type Athlete interface {
	Jumper
	Runner
}
