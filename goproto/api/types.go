package api

// +gogo:genproto=true
type TestStruct1 struct {
	Name      string            `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Map       map[string]string `json:"map,omitempty" protobuf:"bytes,2,rep,name=map"`
	Age       uint64            `json:"age,omitempty" protobuf:"varint,3,opt,name=age"`
	Raise     float64           `json:"raise,omitempty" protobuf:"fixed64,4,opt,name=raise"`
	Address   string            `json:"address,omitempty" protobuf:"bytes,5,opt,name=address"`
	Languages []string          `json:"languages,omitempty" protobuf:"bytes,6,rep,name=languages"`
	Others    Others            `json:"others,omitempty" protobuf:"bytes,7,opt,name=others"`
}

// +gogo:genproto=true
type Others struct {
	Info1 string `json:"info1,omitempty" protobuf:"bytes,1,opt,name=info1"`
}

type IgnoreStruct struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type App struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Type int32  `json:"type" protobuf:"varint,2,opt,name=type"`
}
