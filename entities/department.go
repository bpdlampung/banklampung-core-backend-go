package entities

type Department struct {
	Code     string `bson:"code" json:"code"`
	Name     string `bson:"name" json:"name"`
	Type     string `bson:"type" json:"type"`
	TypeCode string `bson:"typeCode" json:"type_code"`
}
