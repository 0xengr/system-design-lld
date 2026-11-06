package main

import "fmt"

// Let's achieve the abstraction in GO

// Car is the interface. It says WHAT a car can do, not HOW.
// // Go interfaces are satisfied implicitly, so SportsCar never has to
// declare that it implement Car.
type Car interface {
	StartEngine()
	ShiftGear(gear int)
	Accelerate()
	Brake()
	StopEngine()
}

// SportsCar is the concrete type that provides the NOW.
type SportsCar struct {
	brand        string
	model        string
	isEngineOn   bool
	currentSpeed int
	currentGear  int
}

// NewSportsCar is the constructor by convention. It returns the concrete
// type; callers decide whether to treat it as a Car.
func NewSportsCar(brand, model string) *SportsCar {
	return &SportsCar{brand: brand, model: model}
}

func (c *SportsCar) logf(format string, args ...any) {
	fmt.Printf("%s %s : %s\n", c.brand, c.model, fmt.Sprintf(format, args...))
}

func (c *SportsCar) StartEngine() {
	c.isEngineOn = true
	c.logf("Engine starts with a roar")
}

func (c *SportsCar) ShiftGear(gear int) {
	if !c.isEngineOn {
		c.logf("Engine is Off! Cannot shift gear")
		return
	}
	c.currentGear = gear
	c.logf("Shifted to gear: %d", c.currentGear)
}

func (c *SportsCar) Accelerate() {
	if !c.isEngineOn {
		c.logf("Engine is Off! cannot accelerate")
		return
	}
	c.currentSpeed += 20
	c.logf("Accelerating to %d km/h", c.currentSpeed)
}

func (c *SportsCar) Brake() {
	c.currentSpeed = max(0, c.currentSpeed-20)
	c.logf("Braking! Speed is now %d km/h", c.currentSpeed)
}

func (c *SportsCar) StopEngine() {
	c.isEngineOn = false
	c.currentGear = 0
	c.currentSpeed = 0
	c.logf("Engine turned off.")
}

func main() {
	var myCar Car = NewSportsCar("Ford", "Mustang")

	myCar.StartEngine()
	myCar.ShiftGear(1)
	myCar.Accelerate()
	myCar.ShiftGear(2)
	myCar.Accelerate()
	myCar.Brake()
	myCar.StopEngine()
}
