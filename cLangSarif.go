package main

import "encoding/json"

func UnmarshalWelcome(data []byte) (CLangSarif, error) {
	var r CLangSarif
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CLangSarif) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CLangSarif struct {
	Schema  string `json:"$schema"`
	Runs    []Run  `json:"runs"`   
	Version string `json:"version"`
}

type Run struct {
	Artifacts  []Artifact `json:"artifacts"` 
	ColumnKind string     `json:"columnKind"`
	Results    []CLResult   `json:"results"`   
	Tool       Tool       `json:"tool"`      
}

type Artifact struct {
	Length   int64            `json:"length"`  
	Location ArtifactLocation `json:"location"`
	MIMEType string           `json:"mimeType"`
	Roles    []string         `json:"roles"`   
}

type ArtifactLocation struct {
	URI string `json:"uri"`
}

type CLResult struct {
	CodeFlows []CodeFlow       `json:"codeFlows"`
	Locations []ResultLocation `json:"locations"`
	Message   Message          `json:"message"`  
	RuleID    string           `json:"ruleId"`   
	RuleIndex int64            `json:"ruleIndex"`
}

type CodeFlow struct {
	ThreadFlows []ThreadFlow `json:"threadFlows"`
}

type ThreadFlow struct {
	Locations []ThreadFlowLocation `json:"locations"`
}

type ThreadFlowLocation struct {
	Importance string           `json:"importance"`
	Location   LocationLocation `json:"location"`  
}

type LocationLocation struct {
	Message          Message          `json:"message"`         
	PhysicalLocation PhysicalLocation `json:"physicalLocation"`
}

type Message struct {
	Text string `json:"text"`
}

type PhysicalLocation struct {
	ArtifactLocation ArtifactLocationClass `json:"artifactLocation"`
	Region           Region                `json:"region"`          
}

type ArtifactLocationClass struct {
	Index int64  `json:"index"`
	URI   string `json:"uri"`  
}

type Region struct {
	EndColumn   int32  `json:"endColumn"`        
	EndLine     int32 `json:"endLine,omitempty"`
	StartColumn int32  `json:"startColumn"`      
	StartLine   int32  `json:"startLine"`        
}

type ResultLocation struct {
	PhysicalLocation PhysicalLocation `json:"physicalLocation"`
}

type Tool struct {
	Driver Driver `json:"driver"`
}

type Driver struct {
	FullName string `json:"fullName"`
	Language string `json:"language"`
	Name     string `json:"name"`    
	Rules    []Rule `json:"rules"`   
	Version  string `json:"version"` 
}

type Rule struct {
	FullDescription Message `json:"fullDescription"`
	HelpURI         string  `json:"helpUri"`        
	ID              string  `json:"id"`             
	Name            string  `json:"name"`           
}
