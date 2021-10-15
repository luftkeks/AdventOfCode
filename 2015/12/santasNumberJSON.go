package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//commandLineArgs := os.Args
	//fileToRead := commandLineArgs[1]
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var result []interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		panic(err)
	}

	finalNumber := resolveStruct(result)
	fmt.Println("The sum of all the numbers is:", finalNumber)
}

func resolveStruct(input []interface{}) float64 {
	result := 0.
	for _, value := range input {
		result += switchInterface(value)
	}
	return result
}

func resolveMap(input map[string]interface{}) float64 {
	result := 0.
	for _, value := range input {
		str, err := value.(string)
		if err && str == "red" {
			return 0.
		}
		result += switchInterface(value)
	}
	return result
}

func switchInterface(value interface{}) float64 {
	result := 0.
	switch value.(type) {
	case []interface{}:
		cast := value.([]interface{})
		result += resolveStruct([]interface{}(cast))
	case map[string]interface{}:
		cast := value.(map[string]interface{})
		result += resolveMap(cast)
	case int:
		cast := value.(float64)
		result += cast
	case uint:
		cast := value.(float64)
		result += cast
	case float64:
		cast := value.(float64)
		result += cast
	}
	return result
}
