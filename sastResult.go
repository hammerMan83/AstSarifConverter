package main

import (
	"encoding/json"
)

type SastResult struct {
	Version string
	ScanId string
	CreateTime string
	Queries []SastQuery //`json:"first_name"`
}

func newSastResult(jsonSastResult string) SastResult {
	var unmarshaledSastResult SastResult
	json.Unmarshal([]byte(jsonSastResult), &unmarshaledSastResult)
	return unmarshaledSastResult
}

func (r SastResult) toJSONByteArray() ([]byte) { //, error){
	jsonResult, _ := json.Marshal(r)
	return jsonResult
}

func (r SastResult) toJSON() (string) { //, error){
	jsonResult, _ := json.Marshal(r)
	return string(jsonResult)
}

type SastQuery struct {
	Metadata SastMetadata
	Results []SastQueryResult
}

type SastMetadata struct {
	Id uint64
	QueryName string
	GroupName string
	Severity string
	CweId int64
}

type SastQueryResult struct {
	PathId uint32
	SimilarityId int64
	UniqueId int64
	Nodes []SastResultNode
}

type SastResultNode struct {
	Column uint32
	FileName string
	FullName string
	Length uint32
	Line uint64
	MethodLine int64
	Name string
	NodeId uint64
	DomType string
	//Nodes
}