package main

import (
	//"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
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
func cLangSarifToMap(j string) map[string]interface{} { //SastResult {

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

func buildSastResultMap() {



}




