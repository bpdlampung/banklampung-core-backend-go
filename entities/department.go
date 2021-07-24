package entities

type Department struct {
	Code     string  `bson:"code,omitempty" json:"code,omitempty"`
	Name     string  `bson:"name,omitempty" json:"name,omitempty"`
	Type     string  `bson:"type,omitempty" json:"type,omitempty"`
	TypeCode *string `bson:"typeCode,omitempty" json:"type_code,omitempty"`
}
