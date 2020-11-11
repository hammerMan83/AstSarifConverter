package main

import (
	"time"
)

type Result_Severity int32

const (
	Result_HIGH   Result_Severity = 0 // Engine[Critical]
	Result_MEDIUM Result_Severity = 1 // Engine[Serious]
	Result_LOW    Result_Severity = 2 // Engine[Warning]
	Result_INFO   Result_Severity = 3 // Engine[Information, Metrics]
)

type Result_Status int32

const (
	Result_NEW       Result_Status = 0
	Result_RECURRENT Result_Status = 1
	Result_FIXED     Result_Status = 2
)

type Result_Classification int32

const (
	Result_A Result_Classification = 0 //All of the result nodes are inside the changed files - GOOD
	Result_B Result_Classification = 1 //All of the result nodes are inside the closure files - should ignore
	Result_C Result_Classification = 2 //All of the result nodes are outside the closure files - can't be found by inc scan
	Result_D Result_Classification = 3 //The result nodes are inside both the changed files and the closure files - GOOD
	//        E = 4; //The result nodes are both inside and outside of the closure files - can't be found by inc scan! CAN'T BE FOUND ANYWAY! will look like C!
	Result_F       Result_Classification = 5 //The result nodes are inside the changed files, inside the closure files, and also outside the closure files. - can't be found by inc scan
	Result_DELETED Result_Classification = 6 // Some of the result nodes deleted from the source code.
)

type Result struct {
	// identity
	TenantID      string `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`           //ID of the customer tenant
	ProjectID     string `protobuf:"bytes,2,opt,name=projectID,proto3" json:"projectID,omitempty"`         //project ID
	ScanID        string `protobuf:"bytes,3,opt,name=scanID,proto3" json:"scanID,omitempty"`               //ID of the scan
	CorrelationID string `protobuf:"bytes,4,opt,name=correlationID,proto3" json:"correlationID,omitempty"` //correlation ID
	// general
	IsIncremental bool                 `protobuf:"varint,7,opt,name=isIncremental,proto3" json:"isIncremental,omitempty"`
	CreatedAt     time.Time // *timestamp.Timestamp `protobuf:"bytes,1001,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// query
	QueryID      uint64          `protobuf:"varint,11,opt,name=queryID,proto3" json:"queryID,omitempty"`           //Query ID
	QueryVersion int64           `protobuf:"varint,12,opt,name=queryVersion,proto3" json:"queryVersion,omitempty"` //Query Version
	QueryName    string          `protobuf:"bytes,13,opt,name=queryName,proto3" json:"queryName,omitempty"`        //Query name
	LanguageName string          `protobuf:"bytes,14,opt,name=languageName,proto3" json:"languageName,omitempty"`
	PackageName  string          `protobuf:"bytes,15,opt,name=packageName,proto3" json:"packageName,omitempty"`
	GroupName    string          `protobuf:"bytes,16,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Severity     Result_Severity `protobuf:"varint,17,opt,name=severity,proto3,enum=results.Result_Severity" json:"severity,omitempty"`
	CweID        int64           `protobuf:"varint,18,opt,name=cweID,proto3" json:"cweID,omitempty"`
	// path
	PathID          int64          `protobuf:"varint,31,opt,name=pathID,proto3" json:"pathID,omitempty"`                    //ID of the path
	SimilarityID    int64          `protobuf:"varint,32,opt,name=similarityID,proto3" json:"similarityID,omitempty"`        //ID of the Similarity feature (Indicator to identify a result by its first and last nodes)
	UniqueID        int64          `protobuf:"varint,33,opt,name=uniqueID,proto3" json:"uniqueID,omitempty"`                //ID of the Similarity feature (Indicator to identify a result by its first and last nodes)
	ConfidenceLevel float32        `protobuf:"fixed32,34,opt,name=confidenceLevel,proto3" json:"confidenceLevel,omitempty"` //Confidence Level of the existence of the result
	Nodes           []*Result_Node `protobuf:"bytes,35,rep,name=nodes,proto3" json:"nodes,omitempty"`                       //List of nodes that represents the flow of the result
	PathSystemID    string         `protobuf:"bytes,1021,opt,name=pathSystemID,proto3" json:"pathSystemID,omitempty"`       //ID of the customer tenant
	// Inc scan
	Classification Result_Classification `protobuf:"varint,1031,opt,name=classification,proto3,enum=results.Result_Classification" json:"classification,omitempty"` //inc scan result classification - take a look at the option for better understanding.
	// enrich
	Status Result_Status `protobuf:"varint,1032,opt,name=status,proto3,enum=results.Result_Status" json:"status,omitempty"` //enum of the current state(new,old,fixed)
	// additional fields
	MetadataJSON []byte `protobuf:"bytes,1101,opt,name=metadataJSON,proto3" json:"metadataJSON,omitempty"` //TBD
	ExtraJSON    []byte `protobuf:"bytes,1102,opt,name=extraJSON,proto3" json:"extraJSON,omitempty"`       //TBD
}



type Result_Node struct {
	// state         protoimpl.MessageState
	// sizeCache     protoimpl.SizeCache
	// unknownFields protoimpl.UnknownFields

	Column               int32  `protobuf:"varint,1,opt,name=column,proto3" json:"column,omitempty"`         //Column position of the node
	FileName             string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`      //Full file name of the containing source file
	FullName             string `protobuf:"bytes,3,opt,name=fullName,proto3" json:"fullName,omitempty"`      //FQN of the node
	Length               int32  `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`         //Length of the node
	Line                 int32  `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`             //Line position of the node
	MethodLine           int32  `protobuf:"varint,6,opt,name=methodLine,proto3" json:"methodLine,omitempty"` //Line position of the containing method
	Name                 string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`              //node name
	NodeID               int32  `protobuf:"varint,8,opt,name=nodeID,proto3" json:"nodeID,omitempty"`         //ID of node
	DomType              string `protobuf:"bytes,9,opt,name=domType,proto3" json:"domType,omitempty"`        //node DomType
	Method               string `protobuf:"bytes,10,opt,name=method,proto3" json:"method,omitempty"`
	Definitions          string `protobuf:"bytes,11,opt,name=definitions,proto3" json:"definitions,omitempty"`
	TypeName             string `protobuf:"bytes,12,opt,name=typeName,proto3" json:"typeName,omitempty"`
	MethodFullName       string `protobuf:"bytes,13,opt,name=methodFullName,proto3" json:"methodFullName,omitempty"`
	MethodParameterTypes string `protobuf:"bytes,14,opt,name=methodParameterTypes,proto3" json:"methodParameterTypes,omitempty"`
	NodeSystemID         string `protobuf:"bytes,1001,opt,name=nodeSystemID,proto3" json:"nodeSystemID,omitempty"` // unique id of the node
}




