package main

import "fmt"

type car struct {
	gas_pedal      uint16  //min: 0,      max: 65535    16bit
	brake_pedal    uint16  //min: 0,      max: 65535
	steering_wheel int16   //min: -32768  max: 32768
	top_speed_kmh  float64 //what's our top speed?
}

func main() {
	extended_car := car{gas_pedal: 22314, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}

	short_car := car{22314, 0, 12562, 225.0}

	var (
		gas_pedal      uint16  = 22314
		brake_pedal    uint16  = 0
		steering_wheel int16   = 12562
		top_speed_kmh  float64 = 225.0
	)
	var_car := car{gas_pedal, brake_pedal, steering_wheel, top_speed_kmh}

	fmt.Println("[Extended car values] gas_pedal:", extended_car.gas_pedal, "brake_pedal:", extended_car.brake_pedal, "steering_wheel:", extended_car.steering_wheel, "top_speed_kmh:", extended_car.top_speed_kmh)
	fmt.Println("[Short car values] gas_pedal:", short_car.gas_pedal, "brake_pedal:", short_car.brake_pedal, "steering_wheel:", short_car.steering_wheel, "top_speed_kmh:", short_car.top_speed_kmh)
	fmt.Println("[Var car values] gas_pedal:", var_car.gas_pedal, "brake_pedal:", var_car.brake_pedal, "steering_wheel:", var_car.steering_wheel, "top_speed_kmh:", var_car.top_speed_kmh)
	fmt.Println("Var car:", var_car)

	var_car.gas_pedal = 0
	var_car.brake_pedal = 22314
	fmt.Println("New Var car:", var_car)
}
