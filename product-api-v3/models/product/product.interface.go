package models

type iProduct interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
