package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	turns := DivRoundedUp(trackDistance, car.speed)
	return car.battery >= turns*car.batteryDrain
}

func DivRoundedUp(m, n int) int {
	// https://stackoverflow.com/questions/17944/
	return (m-1)/n + 1
}
