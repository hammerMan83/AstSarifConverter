package main

import (
	"encoding/json"
)

//https://stackoverflow.com/questions/32708717/go-when-will-json-unmarshal-to-struct-return-error

type CLangSarif struct {
	schema  string `json:"$schema"`
	runs    []CLangRun
	version string
}

type CLangRun struct {
	artifacts  []CLangRunArtifact
	columnKind string
	results    []CLangRunResult
	tool       CLangRunTool
}

type CLangRunArtifact struct {
}

type CLangRunTool struct {
}

type CLangRunResult struct {
}

func newCLangSarif(jsonCLangSarif string) CLangSarif {
	var unmarshaledCLangSarif CLangSarif
	json.Unmarshal([]byte(jsonCLangSarif), &unmarshaledCLangSarif)
	return unmarshaledCLangSarif
}
