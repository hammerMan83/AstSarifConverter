package main

type sastQuery struct {
	Metadata metadata
	Results results
}

type sastMetadata struct {
	Id int64
	QueryName string
	GroupName string
	Severity string
	CweId string
}

type sastQueryResult struct {
	PathId uint32
	SimilarityId int64
	UniqueId int64
	//Nodes
}