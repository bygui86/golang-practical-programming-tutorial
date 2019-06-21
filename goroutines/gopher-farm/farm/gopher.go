package farm

import (
	"fmt"
	"log"
	"time"
)

type Gopher struct {
	Name  string        `json:"name"`
	Sleep time.Duration `json:"sleep"`
	Eat   int           `json:"eat"`
}

func (g *Gopher) Live(farm *GopherFarm, radio chan string) {
	for {
		time.Sleep(g.Sleep * time.Second)
		err := farm.LunchTime(g)
		if err != nil {
			log.Printf("No more food for Gopher %s...", g.Name)
			radio <- fmt.Sprintf("Gopher %s died", g.Name)
			return
		}
	}
}
