package main

import (
	"encoding/json"
)

type sastResult struct {
	Version string
	ScanId string
	CreateTime string
}

func newSastResult(jsonSastResult string) sastResult {
	var unmarshaledSastResult sastResult
	json.Unmarshal([]byte(jsonSastResult), &unmarshaledSastResult)
	return unmarshaledSastResult
}

func (r sastResult) toJSONByteArray() ([]byte) { //, error){
	jsonResult, _ := json.Marshal(r)
	return jsonResult
}

func (r sastResult) toJSON() (string) { //, error){
	jsonResult, _ := json.Marshal(r)
	return string(jsonResult)
}

