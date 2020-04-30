package main

import (
	"fmt"
)

// calculation: (v2 - v1)/v1, (v3 - v2)/v2, ... (vN - vN-1)/vN-1
func percentageChange(old, new float64) float64 {
	diff := new - old
	return (diff / old) * 100
}

// see https://www.socketloop.com/tutorials/golang-calculate-percentage-change-of-two-values
func main() {

	fmt.Printf("Old : 500, New : 500 Change : %0.2f %% \n", percentageChange(500, 500))

	// anything divided by 0 will become infinity
	fmt.Printf("Old : 0, New : 50 Change : %0.2f %% \n", percentageChange(0, 50))

	fmt.Printf("Old : 100, New : 30 Change : %0.2f %% \n", percentageChange(100, 30))

	fmt.Printf("Old : 1000, New : 830 Change : %0.2f %% \n", percentageChange(1000, 830))

	fmt.Printf("Old : 5, New : 3 Change : %0.2f %% \n", percentageChange(5, 3))

	fmt.Printf("Old : -15, New : 3 Change : %0.2f %% \n", percentageChange(-15, 3))

	fmt.Printf("Old : 100, New : 130 Change : %0.2f %% \n", percentageChange(100, 130))

	fmt.Printf("Old : 42.42, New : 24.24 Change : %0.2f %% \n", percentageChange(42.42, 24.24))
}
