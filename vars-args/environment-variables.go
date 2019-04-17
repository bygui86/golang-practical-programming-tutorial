package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	getVarKey   string = "GOPATH"
	setVarKey   string = "GO_ENV_VAR"
	setVarValue string = "new-var-from-go"
)

func main() {
	getEnvVar(&getVarKey)
	fmt.Println()
	setEnvVar(&setVarKey, &setVarValue)
	fmt.Println()
	printAllEnvVars()
}

func getEnvVar(envVarKey *string) {
	fmt.Println("Get Environment Variable")
	envVarValue := os.Getenv(*envVarKey)
	fmt.Println("\t", *envVarKey, envVarValue)
}

func setEnvVar(envVarKey *string, envVarValue *string) {
	fmt.Println("Set Environment Variable")
	err := os.Setenv(*envVarKey, *envVarValue)
	if err != nil {
		fmt.Println(err)
	}
	envVarValueTemp := os.Getenv(*envVarKey)
	fmt.Println("\t", *envVarKey, envVarValueTemp)
}

func printAllEnvVars() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
