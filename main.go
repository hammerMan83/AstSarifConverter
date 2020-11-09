package main

 import (
 	"fmt"
 )

func main() {

	j := readJSONFromFile("cLangSarifOutput.json")
	m := jsonToMap(j) 
	//print the map
    //fmt.Println(m)

	s := buildSastResultMap(m)
	fmt.Println(s)

	//convertCLangSarifToSastResult(j)

	//fmt.Println(j)



	// jsonResultUnMarshal := `{"Version": "1.0.0",
	// 					"ScanId": "418ebe63-1ab2-4600-a637-b8f4676616c7",
	// 					"CreateTime": "2020-06-30T01:40:37.9046048+03:00"
	// 					}`

	// // sastResult struct object, converted from json
	// newSastResult := newSastResult(jsonResultUnMarshal)

	// fmt.Println("")
	// fmt.Println("Sast Result from Json: ")
	// fmt.Println("Version: " + newSastResult.Version)
	// fmt.Println("ScanId: " + newSastResult.ScanId)
	// fmt.Println("CreateTime: " + newSastResult.CreateTime)

	// // sastResult converted back to json
	// fmt.Println("")
	// fmt.Println("Sast Result back to Json: ")
	// fmt.Println(newSastResult.toJSON())
}

