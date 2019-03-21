package main

import "fmt"

const uint16_max float64 = 65535
const kmh_mph_factor float64 = 1.60934

type car struct {
	gas_pedal      uint16  //min: 0,      max: 65535    16bit
	brake_pedal    uint16  //min: 0,      max: 65535
	steering_wheel int16   //min: -32768  max: 32768
	top_speed_kmh  float64 //what's our top speed?
}

// The method gets a copy of the object (receiver type), so you cannot actually modify it here, you can only take actions or do something like coming up with a calculation.
func (c *car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / uint16_max)
}
func (c *car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / kmh_mph_factor / uint16_max)
}

// The method MODIFIES THE STRUCT ITSELF VIA POINTER
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{22314, 0, 12562, 225.0}
	fmt.Println("[Car specs] gas_pedal:", a_car.gas_pedal, "brake_pedal:", a_car.brake_pedal, "steering_wheel:", a_car.steering_wheel)
	fmt.Println("Car is going", a_car.mph(), "MPH,", a_car.kmh(), "KMH, and top speed is", a_car.top_speed_kmh)
	a_car.new_top_speed(500)
	fmt.Println("Car is now going", a_car.mph(), "MPH,", a_car.kmh(), "KMH, and top speed is", a_car.top_speed_kmh)
}
