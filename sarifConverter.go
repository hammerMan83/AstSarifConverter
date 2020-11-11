package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

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

func convertCLangSarifJSONToStruct(clangJSONSarifResult string) []Result { //CLangResult {
	var c CLangSarif
	json.Unmarshal([]byte(clangJSONSarifResult), &c)

	var sastResults []Result

	for _, run := range c.Runs {
		rules := run.Tool.Driver.Rules
		fmt.Println("Results")
		for j, result := range run.Results {
			var sastResult Result
			sastResult.ScanID = uuid.New().String()
			sastResult.ProjectID = "1" // TODO: Make sure we use project id "1"
			sastResult.IsIncremental = false
			sastResult.CreatedAt = time.Now()
			sastResult.QueryID = 1               // TODO: Make sure we use query id 1
			sastResult.QueryVersion = 1          // TODO: Make sure we use QueryVersion 1
			sastResult.QueryName = rules[j].Name // Rule is equivalent to a query.
			sastResult.Nodes = createNodes(result.Locations, result)

			if strings.HasSuffix(sastResult.Nodes[0].FileName, "cpp") {
				sastResult.LanguageName = "c++"
			}

			if strings.HasSuffix(sastResult.Nodes[0].FileName, "py") {
				sastResult.LanguageName = "python"
			}

			sastResult.PackageName = "My Package"
			sastResult.GroupName = "My Group"
			sastResult.Severity = Result_Severity(rand.Intn(4))
			sastResult.CweID = int64(rand.Intn(100000000))
			sastResult.PathID = int64(j)
			sastResult.UniqueID = int64(j)
			sastResult.SimilarityID = int64(rand.Intn(9223372036854775807))
			sastResult.ConfidenceLevel = 100
			sastResult.Classification = Result_Classification(rand.Intn(7))
			sastResult.Status = Result_Status(rand.Intn(3))

			sastResults = append(sastResults, sastResult)
		}
	}

	return sastResults	
}

func createNodes(locations []ResultLocation, clResult CLResult) []*Result_Node {

	var nodes []*Result_Node

	for i, location := range locations {
		var node Result_Node
		node.Name = clResult.Message.Text
		node.FileName = location.PhysicalLocation.ArtifactLocation.URI
		node.FullName = location.PhysicalLocation.ArtifactLocation.URI
		node.Line = location.PhysicalLocation.Region.StartLine
		node.Column = location.PhysicalLocation.Region.StartColumn
		node.NodeID = int32(i)

		fmt.Println("****Node: ", node)
		fmt.Println("\n")

		nodes = append(nodes, &node)
	}

	return nodes
}

// func buildSastResultMap(c string) { //map[string]interface{} {

// }
