package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
	see https://gist.github.com/alexmcroberts/219127816e7a16c7bd70
	Example for unmarshalling JSON containing an epoch timestamp with the value in milliseconds,
	stored as a string:
		{"timestamp": "1436150027000"}

	Note the three zeroes at the end of the value, and also note that the value is stored as a string
	– not as a number
		string: {"timestamp": "1436150027000"}
		number: {"timestamp": 1436150027000}

	The example code above will process both samples listed where epoch time in milliseconds is stored
	as either string or number
*/

func main() {
	b := []byte(`{"timestamp": 1436150027000}`)
	var ep EpochParent
	err := json.Unmarshal(b, &ep)
	fmt.Println(ep, err)

	fmt.Println(ep.Epoch)
	fmt.Println("=======")
	c := []byte(`{"timestamp": "1436150027000"}`)
	var dp EpochParent
	err = json.Unmarshal(c, &dp)
	fmt.Println(dp, err)

	fmt.Println(dp.Epoch)
	fmt.Println("=======")
}

type EpochParent struct {
	Epoch jsonTime `json:"timestamp"`
}

type jsonTime time.Time

func (t jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *jsonTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)

	// *** maybe redundant
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q/1000, 0)
	// ***

	return
}

func (t jsonTime) String() string { return time.Time(t).String() }
