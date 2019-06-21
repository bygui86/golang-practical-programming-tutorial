package main

import (
	"fmt"
	"strconv"
	"time"
)

// See https://medium.com/@kpbird/golang-parse-and-format-unix-timestamp-6f467389cfe1
func main() {
	unixTimestampAsString()
	unixTimestampAsInt()
	timestampToRFC3339()
}

// The time.Parse function does not do Unix timestamps.
// Instead you can use strconv.ParseIntto parse the string to int64 and create the timestamp with time.Unix
func unixTimestampAsString() {
	unixTimestamp := "1518328047"
	i, err := strconv.ParseInt(unixTimestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	time := time.Unix(i, 0)
	fmt.Println(time)
}

func unixTimestampAsInt() {
	var unixTimestamp int64 = 1518328047
	time := time.Unix(unixTimestamp, 0)
	fmt.Println(time)
}

func timestampToRFC3339() {
	var unixTimestamp int64 = 1405544146
	unixTimeUTC := time.Unix(unixTimestamp, 0) // gives unix time stamp in utc
	fmt.Println("UTC: ", unixTimeUTC)
	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format
	fmt.Println("RFC3339:", unitTimeInRFC3339)
}
