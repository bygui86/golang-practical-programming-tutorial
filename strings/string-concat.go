package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

func main() {

	strs := []string{}
	for i := 0; i < 100; i++ {
		strs = append(strs, "a")
	}

	log.Println("[SLOWEST] Naive concat...")
	naiveConcat(strs)

	log.Println("[AVERAGE BUT MEMORY-EFFICIENT] Strings join...")
	stringsJoin(strs)

	log.Println("[PRETTY FAST] String builder...")
	strBuild(strs)

	log.Println("[FASTEST] Bytes buffer...")
	buffer(strs)

	log.Println("String builder with IO Writer...")
	strBuildIo()
}

func naiveConcat(strs ...string) string {
	var naiveConcat string
	for _, str := range strs {
		naiveConcat += str
	}
	return naiveConcat
}

func stringsJoin(strs ...string) string {
	return strings.Join(strs, ",")
}

func strBuild(strs ...string) string {
	var strBuild strings.Builder
	for _, str := range strs {
		strBuild.WriteString(str)
	}
	return strBuild.String()
}

func buffer(strs ...string) string {
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func strBuildIo() string {
	var strBuild strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&strBuild, "%d...", i)
	}
	strBuild.WriteString("ignition")
	return strBuild.String()
}
