package main

import (
	//"os"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"
)

func readJSONFromFile(fileName string) string {
	s, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Can't read file:", fileName)
		panic(err)
	}

	return string(s)
}

//https://stackoverflow.com/questions/32708717/go-when-will-json-unmarshal-to-struct-return-error
func jsonToMap(j string) map[string]interface{} { //SastResult {

	//Marshal the json to a map
	var i interface{}
	json.Unmarshal([]byte(j), &i)

	m := i.(map[string]interface{})

	return m

	// //print the map
	// fmt.Println(m)

	//unmarshal the map to json
	//result,_:= json.Marshal(m)

	//print the json
	//os.Stdout.Write(result)
}

func buildSastResultMap(m map[string]interface{}) map[string]interface{} {
	// Create a new Cx Sast Result Map

	// TODO: Need to build the json in another way. Map does not maintain order,
	// so it's not suitable here.

	r := make(map[string]interface{})
	r["Version"] = "1.0.0"
	r["ScanId"] = uuid.New()
	return r
}
