package farm

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

type GopherFarm struct {
	Gophers   []*Gopher `json:"gophers"`
	TotalFood int       `json:"totalFood"`
	sync.Mutex
	radio chan string
}

func Build() *GopherFarm {
	farm := &GopherFarm{}
	err := json.Unmarshal([]byte(inputJson), farm)
	if err != nil {
		log.Panicln("Error parsing json input")
	}
	return farm
}

func (f *GopherFarm) Open() {
	f.radio = make(chan string)
	for _, gopher := range f.Gophers {
		log.Printf("Gopher %s joins the farm", gopher.Name)
		go gopher.Live(f, f.radio)
	}
}

func (f *GopherFarm) LiveFreeOrDieHard() {
	for range f.Gophers {
		msg, ok := <-f.radio
		if !ok {
			log.Panicln("Radio interference!!!")
		}
		log.Println(msg)
	}
	log.Println("Gopher farm is empty, all Gophers died :( I'm so sad!")
}

func (f *GopherFarm) LunchTime(gopher *Gopher) error {
	f.Lock()
	defer f.Unlock()

	if f.TotalFood < gopher.Eat {
		log.Printf("Gopher %s wants to eat %d, but not enough food!", gopher.Name, gopher.Eat)
		return errors.New("Not enough food")
	}

	log.Printf("Gopher %s eats %d, gnummy!", gopher.Name, gopher.Eat)
	f.TotalFood -= gopher.Eat
	return nil
}
